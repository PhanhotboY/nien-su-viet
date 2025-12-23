#!/usr/bin/env bash

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Configuration
APP_NAME="nien-su-viet-cms"
SERVICE_NAME="${APP_NAME}.service"
INSTALL_DIR="/opt/${APP_NAME}"
LOG_DIR="/var/log/${APP_NAME}"
BINARY_NAME="server"
GO_APP_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd)"
SERVICE_FILE="${GO_APP_DIR}/cmd/cli/${SERVICE_NAME}"

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

check_root() {
    if [[ $EUID -ne 0 ]]; then
        print_error "This script must be run as root (use sudo)"
        exit 1
    fi
}

build() {
    print_info "Building Go application..."
    cd "${GO_APP_DIR}"

    # Build the server binary
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o "${BINARY_NAME}" ./cmd/server/main.go

    if [ $? -eq 0 ]; then
        print_info "Build successful: ${BINARY_NAME}"
    else
        print_error "Build failed"
        exit 1
    fi
}

install() {
    check_root
    print_info "Installing application to ${INSTALL_DIR}..."

    # Create installation directory
    mkdir -p "${INSTALL_DIR}"

    # Create log directory
    mkdir -p "${LOG_DIR}"

    # Copy binary
    if [ -f "${GO_APP_DIR}/${BINARY_NAME}" ]; then
        cp "${GO_APP_DIR}/${BINARY_NAME}" "${INSTALL_DIR}/"
        chmod +x "${INSTALL_DIR}/${BINARY_NAME}"
        print_info "Binary installed to ${INSTALL_DIR}/${BINARY_NAME}"
    else
        print_error "Binary not found. Run 'build' first."
        exit 1
    fi

    # Copy config files if they exist
    if [ -d "${GO_APP_DIR}/config" ]; then
        cp -r "${GO_APP_DIR}/config" "${INSTALL_DIR}/"
        print_info "Config files copied"
    fi

    # Set ownership (change www-data to your preferred user)
    chown -R www-data:www-data "${INSTALL_DIR}"
    chown -R www-data:www-data "${LOG_DIR}"

    # Install systemd service
    if [ -f "${SERVICE_FILE}" ]; then
        cp "${SERVICE_FILE}" "/etc/systemd/system/${SERVICE_NAME}"
        systemctl daemon-reload
        print_info "Systemd service installed"
    else
        print_warn "Service file not found at ${SERVICE_FILE}"
    fi

    print_info "Installation complete!"
    print_info "Next steps:"
    echo "  1. Configure environment variables in /etc/systemd/system/${SERVICE_NAME}.d/override.conf"
    echo "  2. Enable service: sudo systemctl enable ${SERVICE_NAME}"
    echo "  3. Start service: sudo systemctl start ${SERVICE_NAME}"
}

uninstall() {
    check_root
    print_info "Uninstalling application..."

    # Stop and disable service
    if systemctl is-active --quiet "${SERVICE_NAME}"; then
        systemctl stop "${SERVICE_NAME}"
        print_info "Service stopped"
    fi

    if systemctl is-enabled --quiet "${SERVICE_NAME}" 2>/dev/null; then
        systemctl disable "${SERVICE_NAME}"
        print_info "Service disabled"
    fi

    # Remove service file
    if [ -f "/etc/systemd/system/${SERVICE_NAME}" ]; then
        rm "/etc/systemd/system/${SERVICE_NAME}"
        systemctl daemon-reload
        print_info "Service file removed"
    fi

    # Remove installation directory
    if [ -d "${INSTALL_DIR}" ]; then
        rm -rf "${INSTALL_DIR}"
        print_info "Installation directory removed"
    fi

    # Ask about log directory
    read -p "Remove log directory ${LOG_DIR}? (y/N) " -n 1 -r
    echo
    if [[ $REPLY =~ ^[Yy]$ ]]; then
        rm -rf "${LOG_DIR}"
        print_info "Log directory removed"
    fi

    print_info "Uninstall complete!"
}

start_service() {
    check_root
    print_info "Starting ${SERVICE_NAME}..."
    systemctl start "${SERVICE_NAME}"
    systemctl status "${SERVICE_NAME}" --no-pager
}

stop_service() {
    check_root
    print_info "Stopping ${SERVICE_NAME}..."
    systemctl stop "${SERVICE_NAME}"
}

restart_service() {
    check_root
    print_info "Restarting ${SERVICE_NAME}..."
    systemctl restart "${SERVICE_NAME}"
    systemctl status "${SERVICE_NAME}" --no-pager
}

status_service() {
    systemctl status "${SERVICE_NAME}" --no-pager
}

logs_service() {
    journalctl -u "${SERVICE_NAME}" -f
}

enable_service() {
    check_root
    print_info "Enabling ${SERVICE_NAME} to start on boot..."
    systemctl enable "${SERVICE_NAME}"
}

disable_service() {
    check_root
    print_info "Disabling ${SERVICE_NAME} from starting on boot..."
    systemctl disable "${SERVICE_NAME}"
}

deploy() {
    print_info "Deploying application..."
    build
    install
    enable_service
    start_service
    print_info "Deployment complete!"
}

update() {
    check_root
    print_info "Updating application..."
    build

    # Stop service
    systemctl stop "${SERVICE_NAME}"

    # Copy new binary
    cp "${GO_APP_DIR}/${BINARY_NAME}" "${INSTALL_DIR}/"
    chmod +x "${INSTALL_DIR}/${BINARY_NAME}"
    chown www-data:www-data "${INSTALL_DIR}/${BINARY_NAME}"

    # Start service
    systemctl start "${SERVICE_NAME}"
    systemctl status "${SERVICE_NAME}" --no-pager

    print_info "Update complete!"
}

show_help() {
    cat << EOF
Usage: $0 [COMMAND]

Commands:
    build           Build the Go application
    install         Install the application and systemd service
    uninstall       Uninstall the application and service
    deploy          Build, install, enable, and start the service
    update          Build and update the running service

    start           Start the systemd service
    stop            Stop the systemd service
    restart         Restart the systemd service
    status          Show service status
    logs            Show and follow service logs
    enable          Enable service to start on boot
    disable         Disable service from starting on boot

    help            Show this help message

Examples:
    $0 build                    # Build the application
    $0 deploy                   # Full deployment
    $0 update                   # Update running service
    $0 status                   # Check service status
    $0 logs                     # View live logs

Configuration:
    Install directory: ${INSTALL_DIR}
    Log directory: ${LOG_DIR}
    Service name: ${SERVICE_NAME}
EOF
}

# Main script
case "${1:-help}" in
    build)
        build
        ;;
    install)
        install
        ;;
    uninstall)
        uninstall
        ;;
    deploy)
        deploy
        ;;
    update)
        update
        ;;
    start)
        start_service
        ;;
    stop)
        stop_service
        ;;
    restart)
        restart_service
        ;;
    status)
        status_service
        ;;
    logs)
        logs_service
        ;;
    enable)
        enable_service
        ;;
    disable)
        disable_service
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
