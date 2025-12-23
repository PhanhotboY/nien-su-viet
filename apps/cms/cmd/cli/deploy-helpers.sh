#!/usr/bin/env bash

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Configuration
APP_NAME="nien-su-viet-cms"
SERVICE_NAME="${APP_NAME}.service"
INSTALL_DIR="/opt/${APP_NAME}"
BACKUP_DIR="/opt/${APP_NAME}-backups"
BINARY_NAME="server"
MAX_BACKUPS=5

# Functions
print_info() {
    echo -e "${GREEN}[INFO]${NC} $1"
}

print_warn() {
    echo -e "${YELLOW}[WARN]${NC} $1"
}

print_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

print_step() {
    echo -e "${BLUE}==>${NC} $1"
}

check_root() {
    if [[ $EUID -ne 0 ]]; then
        print_error "This script must be run as root (use sudo)"
        exit 1
    fi
}

# Health check function
health_check() {
    local max_attempts=${1:-30}
    local port=${2:-8080}
    local attempt=1

    print_info "Running health check (max ${max_attempts} attempts)..."

    while [ $attempt -le $max_attempts ]; do
        if systemctl is-active --quiet "${SERVICE_NAME}"; then
            # Check if process is responding
            if curl -f -s -o /dev/null "http://localhost:${port}/health" 2>/dev/null || \
               curl -f -s -o /dev/null "http://localhost:${port}/" 2>/dev/null; then
                print_info "✅ Service is healthy and responding"
                return 0
            fi
        fi

        echo -n "."
        sleep 2
        ((attempt++))
    done

    echo
    print_error "❌ Health check failed after ${max_attempts} attempts"
    return 1
}

# Backup current binary
backup_binary() {
    check_root
    print_step "Creating backup of current binary..."

    # Create backup directory if it doesn't exist
    mkdir -p "${BACKUP_DIR}"

    if [ -f "${INSTALL_DIR}/${BINARY_NAME}" ]; then
        local timestamp=$(date +%Y%m%d_%H%M%S)
        local backup_file="${BACKUP_DIR}/${BINARY_NAME}_${timestamp}"

        cp "${INSTALL_DIR}/${BINARY_NAME}" "${backup_file}"
        print_info "Backup created: ${backup_file}"

        # Cleanup old backups
        cleanup_old_backups

        echo "${backup_file}"
    else
        print_warn "No binary found to backup"
        return 1
    fi
}

# Cleanup old backups
cleanup_old_backups() {
    local backup_count=$(ls -1 "${BACKUP_DIR}/${BINARY_NAME}_"* 2>/dev/null | wc -l)

    if [ "$backup_count" -gt "$MAX_BACKUPS" ]; then
        print_info "Cleaning up old backups (keeping last ${MAX_BACKUPS})..."
        ls -1t "${BACKUP_DIR}/${BINARY_NAME}_"* | tail -n +$((MAX_BACKUPS + 1)) | xargs rm -f
    fi
}

# List available backups
list_backups() {
    print_step "Available backups:"

    if [ -d "${BACKUP_DIR}" ] && [ "$(ls -A ${BACKUP_DIR})" ]; then
        ls -lht "${BACKUP_DIR}/${BINARY_NAME}_"* | awk '{print NR". "$9" ("$6" "$7" "$8")"}'
    else
        print_warn "No backups found"
    fi
}

# Rollback to previous version
rollback() {
    check_root

    if [ ! -d "${BACKUP_DIR}" ] || [ ! "$(ls -A ${BACKUP_DIR})" ]; then
        print_error "No backups available for rollback"
        exit 1
    fi

    print_step "Available backups for rollback:"
    list_backups

    echo
    read -p "Enter backup number to rollback (or 'latest' for most recent): " choice

    local backup_file=""
    if [ "$choice" = "latest" ] || [ "$choice" = "1" ]; then
        backup_file=$(ls -1t "${BACKUP_DIR}/${BINARY_NAME}_"* | head -n 1)
    else
        backup_file=$(ls -1t "${BACKUP_DIR}/${BINARY_NAME}_"* | sed -n "${choice}p")
    fi

    if [ -z "$backup_file" ] || [ ! -f "$backup_file" ]; then
        print_error "Invalid backup selection"
        exit 1
    fi

    print_info "Rolling back to: ${backup_file}"

    # Stop service
    print_info "Stopping service..."
    systemctl stop "${SERVICE_NAME}"

    # Restore backup
    cp "$backup_file" "${INSTALL_DIR}/${BINARY_NAME}"
    chmod +x "${INSTALL_DIR}/${BINARY_NAME}"
    chown www-data:www-data "${INSTALL_DIR}/${BINARY_NAME}"

    # Start service
    print_info "Starting service..."
    systemctl start "${SERVICE_NAME}"

    # Health check
    if health_check 30; then
        print_info "✅ Rollback successful!"
        systemctl status "${SERVICE_NAME}" --no-pager
    else
        print_error "❌ Rollback failed - service is not healthy"
        print_info "Check logs: journalctl -u ${SERVICE_NAME} -n 50"
        exit 1
    fi
}

# Safe deployment with rollback capability
safe_deploy() {
    check_root
    local new_binary=$1

    if [ ! -f "$new_binary" ]; then
        print_error "Binary not found: ${new_binary}"
        exit 1
    fi

    print_step "Starting safe deployment..."

    # Backup current version
    local backup_file=$(backup_binary)

    if [ -z "$backup_file" ]; then
        print_warn "No backup created, proceeding anyway..."
    fi

    # Stop service
    print_info "Stopping service..."
    systemctl stop "${SERVICE_NAME}"

    # Deploy new binary
    print_info "Deploying new binary..."
    cp "$new_binary" "${INSTALL_DIR}/${BINARY_NAME}"
    chmod +x "${INSTALL_DIR}/${BINARY_NAME}"
    chown www-data:www-data "${INSTALL_DIR}/${BINARY_NAME}"

    # Start service
    print_info "Starting service..."
    systemctl start "${SERVICE_NAME}"

    # Health check
    if health_check 30; then
        print_info "✅ Deployment successful!"
        systemctl status "${SERVICE_NAME}" --no-pager
    else
        print_error "❌ Deployment failed - initiating automatic rollback..."

        if [ -n "$backup_file" ] && [ -f "$backup_file" ]; then
            systemctl stop "${SERVICE_NAME}"
            cp "$backup_file" "${INSTALL_DIR}/${BINARY_NAME}"
            chmod +x "${INSTALL_DIR}/${BINARY_NAME}"
            chown www-data:www-data "${INSTALL_DIR}/${BINARY_NAME}"
            systemctl start "${SERVICE_NAME}"

            if health_check 20; then
                print_info "✅ Rollback successful - service restored"
            else
                print_error "❌ Rollback failed - manual intervention required"
                print_info "Check logs: journalctl -u ${SERVICE_NAME} -n 50"
            fi
        else
            print_error "No backup available for automatic rollback"
        fi

        exit 1
    fi
}

# Check service dependencies
check_dependencies() {
    print_step "Checking service dependencies..."

    local all_ok=true

    # Check PostgreSQL
    if systemctl is-active --quiet postgresql; then
        print_info "✅ PostgreSQL is running"
    else
        print_error "❌ PostgreSQL is not running"
        all_ok=false
    fi

    # Check RabbitMQ
    if systemctl is-active --quiet rabbitmq-server; then
        print_info "✅ RabbitMQ is running"
    else
        print_error "❌ RabbitMQ is not running"
        all_ok=false
    fi

    # Check if binary exists
    if [ -f "${INSTALL_DIR}/${BINARY_NAME}" ]; then
        print_info "✅ Binary exists: ${INSTALL_DIR}/${BINARY_NAME}"
    else
        print_error "❌ Binary not found: ${INSTALL_DIR}/${BINARY_NAME}"
        all_ok=false
    fi

    # Check log directory
    if [ -d "/var/log/${APP_NAME}" ]; then
        print_info "✅ Log directory exists"
    else
        print_warn "⚠️  Log directory not found"
    fi

    if [ "$all_ok" = true ]; then
        print_info "✅ All dependencies are satisfied"
        return 0
    else
        print_error "❌ Some dependencies are missing"
        return 1
    fi
}

# View detailed service information
service_info() {
    print_step "Service Information"
    echo

    echo "📦 Service Name: ${SERVICE_NAME}"
    echo "📁 Install Directory: ${INSTALL_DIR}"
    echo "📄 Log Directory: /var/log/${APP_NAME}"
    echo "💾 Backup Directory: ${BACKUP_DIR}"
    echo

    if systemctl is-active --quiet "${SERVICE_NAME}"; then
        echo "🟢 Status: Running"
    else
        echo "🔴 Status: Stopped"
    fi

    if systemctl is-enabled --quiet "${SERVICE_NAME}" 2>/dev/null; then
        echo "🔄 Auto-start: Enabled"
    else
        echo "🔄 Auto-start: Disabled"
    fi

    echo
    print_step "Process Information"
    systemctl show "${SERVICE_NAME}" --property=MainPID,ActiveEnterTimestamp,MemoryCurrent,CPUUsageNSec --no-pager 2>/dev/null || true

    echo
    print_step "Recent Restarts"
    journalctl -u "${SERVICE_NAME}" --since "7 days ago" | grep -i "started\|stopped\|restart" | tail -n 5 || echo "No recent restarts"

    echo
    print_step "Backup Information"
    list_backups
}

# Monitor service logs with filtering
monitor_logs() {
    local filter=${1:-""}

    if [ -n "$filter" ]; then
        print_info "Monitoring logs with filter: ${filter}"
        journalctl -u "${SERVICE_NAME}" -f | grep --line-buffered "$filter"
    else
        print_info "Monitoring all logs (press Ctrl+C to stop)"
        journalctl -u "${SERVICE_NAME}" -f
    fi
}

# Show usage
show_help() {
    cat << EOF
Usage: $0 [COMMAND] [OPTIONS]

Deployment Helper Commands:
    health-check [attempts] [port]    Check if service is healthy
    backup                            Backup current binary
    list-backups                      List all available backups
    rollback                          Rollback to a previous version
    safe-deploy <binary>              Deploy with automatic rollback on failure
    check-deps                        Check service dependencies
    info                              Show detailed service information
    monitor [filter]                  Monitor logs with optional filter
    help                              Show this help message

Examples:
    $0 health-check                   # Quick health check
    $0 health-check 60 8080           # Check for 60 attempts on port 8080
    $0 backup                         # Create backup
    $0 list-backups                   # Show all backups
    $0 rollback                       # Interactive rollback
    $0 safe-deploy ./server           # Deploy with auto-rollback
    $0 check-deps                     # Verify all dependencies
    $0 info                           # Show service details
    $0 monitor "error"                # Monitor only error logs

Notes:
    - Most commands require sudo/root access
    - Health checks assume service exposes /health or / endpoint
    - Backups are kept in ${BACKUP_DIR}
    - Maximum ${MAX_BACKUPS} backups are retained
EOF
}

# Main script
case "${1:-help}" in
    health-check|health)
        health_check "${2:-30}" "${3:-8080}"
        ;;
    backup)
        backup_binary
        ;;
    list-backups|list)
        list_backups
        ;;
    rollback)
        rollback
        ;;
    safe-deploy)
        if [ -z "$2" ]; then
            print_error "Please provide binary path"
            echo "Usage: $0 safe-deploy <binary>"
            exit 1
        fi
        safe_deploy "$2"
        ;;
    check-deps|check-dependencies)
        check_dependencies
        ;;
    info)
        service_info
        ;;
    monitor|logs)
        monitor_logs "$2"
        ;;
    help|--help|-h)
        show_help
        ;;
    *)
        print_error "Unknown command: $1"
        echo
        show_help
        exit 1
        ;;
esac
