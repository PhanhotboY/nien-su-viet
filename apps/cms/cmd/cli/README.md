# Nien Su Viet CMS - Production Deployment Guide

This directory contains scripts and configuration files for deploying the CMS application in production using systemd.

## Files

- `run.sh` - Main deployment and management script
- `nien-su-viet-cms.service` - Systemd service configuration
- `environment.conf.example` - Example environment variables configuration

## Prerequisites

- Go 1.25.3 or higher
- Linux system with systemd
- PostgreSQL database
- RabbitMQ server
- Root/sudo access for installation

## Quick Start

### 1. Build the Application

```bash
./run.sh build
```

This compiles the Go application into a static binary.

### 2. Full Deployment

For a complete deployment (build, install, enable, and start):

```bash
sudo ./run.sh deploy
```

This will:
- Build the application
- Install to `/opt/nien-su-viet-cms`
- Set up logging in `/var/log/nien-su-viet-cms`
- Install and enable the systemd service
- Start the service

### 3. Configure Environment Variables

Create a systemd override file for environment variables:

```bash
sudo mkdir -p /etc/systemd/system/nien-su-viet-cms.service.d
sudo cp environment.conf.example /etc/systemd/system/nien-su-viet-cms.service.d/override.conf
sudo nano /etc/systemd/system/nien-su-viet-cms.service.d/override.conf
```

Edit the `override.conf` file with your actual configuration values:
- Database credentials
- RabbitMQ credentials
- JWT secrets
- CORS origins
- etc.

After editing environment variables:

```bash
sudo systemctl daemon-reload
sudo systemctl restart nien-su-viet-cms.service
```

## Available Commands

### Build and Deployment

- `./run.sh build` - Build the Go application
- `./run.sh install` - Install the application (requires sudo)
- `./run.sh uninstall` - Remove the application (requires sudo)
- `./run.sh deploy` - Full deployment: build + install + enable + start (requires sudo)
- `./run.sh update` - Update running service with new build (requires sudo)

### Service Management

- `./run.sh start` - Start the service (requires sudo)
- `./run.sh stop` - Stop the service (requires sudo)
- `./run.sh restart` - Restart the service (requires sudo)
- `./run.sh status` - Show service status
- `./run.sh logs` - Show and follow live logs
- `./run.sh enable` - Enable service to start on boot (requires sudo)
- `./run.sh disable` - Disable service from starting on boot (requires sudo)

## Usage Examples

### Initial Deployment

```bash
# 1. Build the application
./run.sh build

# 2. Install (requires sudo)
sudo ./run.sh install

# 3. Configure environment variables
sudo nano /etc/systemd/system/nien-su-viet-cms.service.d/override.conf

# 4. Reload systemd and start
sudo systemctl daemon-reload
sudo systemctl enable nien-su-viet-cms.service
sudo systemctl start nien-su-viet-cms.service

# 5. Check status
./run.sh status
```

Or use the single deploy command:

```bash
sudo ./run.sh deploy
# Then configure environment and restart
```

### Update Deployment

```bash
# Pull latest code
git pull

# Update the service (builds and restarts)
sudo ./run.sh update

# Check status
./run.sh status
```

### Monitoring

```bash
# Check service status
./run.sh status

# View logs
./run.sh logs

# Or use journalctl directly
sudo journalctl -u nien-su-viet-cms.service -n 100

# Check log files
sudo tail -f /var/log/nien-su-viet-cms/app.log
sudo tail -f /var/log/nien-su-viet-cms/error.log
```

### Troubleshooting

```bash
# Check service status
systemctl status nien-su-viet-cms.service

# View recent logs
journalctl -u nien-su-viet-cms.service -n 50

# View logs since last boot
journalctl -u nien-su-viet-cms.service -b

# Follow logs in real-time
./run.sh logs

# Check if service is enabled
systemctl is-enabled nien-su-viet-cms.service

# Check if service is active
systemctl is-active nien-su-viet-cms.service

# Restart service
sudo ./run.sh restart
```

## Directory Structure

After installation, the following directories will be created:

```
/opt/nien-su-viet-cms/          # Application installation directory
├── server                       # Main binary
└── config/                      # Configuration files (if any)

/var/log/nien-su-viet-cms/      # Log directory
├── app.log                      # Standard output logs
└── error.log                    # Error logs

/etc/systemd/system/
├── nien-su-viet-cms.service    # Systemd service file
└── nien-su-viet-cms.service.d/ # Service overrides
    └── override.conf            # Environment variables
```

## Service Configuration

The systemd service is configured with:

- **User/Group**: `www-data` (change in service file if needed)
- **Auto-restart**: Yes, with 5-second delay
- **Graceful shutdown**: 10-second timeout
- **Dependencies**: PostgreSQL and RabbitMQ services
- **Security**: Hardened with restricted filesystem access

### Customizing the Service User

To change the service user from `www-data`:

1. Edit `/etc/systemd/system/nien-su-viet-cms.service`
2. Change `User=` and `Group=` lines
3. Update ownership: `sudo chown -R newuser:newgroup /opt/nien-su-viet-cms /var/log/nien-su-viet-cms`
4. Reload and restart: `sudo systemctl daemon-reload && sudo systemctl restart nien-su-viet-cms.service`

## Security Considerations

1. **Environment Variables**: Store sensitive data (passwords, secrets) in the systemd override file, not in the code
2. **File Permissions**: The installation directory is read-only for the service
3. **User Isolation**: Service runs as `www-data` (or your configured user), not root
4. **Logs**: Ensure log rotation is configured to prevent disk space issues

### Setting up Log Rotation

Create `/etc/logrotate.d/nien-su-viet-cms`:

```
/var/log/nien-su-viet-cms/*.log {
    daily
    rotate 14
    compress
    delaycompress
    notifempty
    create 0640 www-data www-data
    sharedscripts
    postrotate
        systemctl reload nien-su-viet-cms.service > /dev/null 2>&1 || true
    endscript
}
```

## Production Checklist

- [ ] PostgreSQL database is set up and accessible
- [ ] RabbitMQ is installed and configured
- [ ] Environment variables are configured in systemd override
- [ ] JWT secret is generated and secure
- [ ] Database migrations are run
- [ ] Service user has appropriate permissions
- [ ] Log rotation is configured
- [ ] Firewall allows access to application port
- [ ] HTTPS/TLS is configured (via reverse proxy like nginx)
- [ ] Monitoring and alerting is set up
- [ ] Backup strategy is in place

## Reverse Proxy Setup (Nginx)

Example nginx configuration:

```nginx
server {
    listen 80;
    server_name your-domain.com;

    location / {
        proxy_pass http://localhost:8080;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
    }
}
```

For HTTPS, use certbot to obtain SSL certificates.

## GitHub Actions CI/CD

The project includes automated CI/CD using GitHub Actions that automatically deploys to production when code is pushed to the `master` branch.

### Workflow Overview

**File**: `.github/workflows/cms-service.yml`

The workflow has two jobs:

1. **deploy-cms-service** (on push to master):
   - Pulls latest code on the server
   - Builds and updates the service using `run.sh update`
   - Shows deployment status and recent logs

2. **test-build** (on pull requests):
   - Checks out code
   - Sets up Go environment
   - Caches Go modules for faster builds
   - Builds the application to verify compilation
   - Runs tests

### Required GitHub Secrets

Configure these secrets in your GitHub repository settings (`Settings > Secrets and variables > Actions`):

- `DOCEAN_NSV_SSH_HOST` - Your server's IP address or hostname
- `DOCEAN_NSV_SSH_USER` - SSH username (usually root or your user)
- `DOCEAN_NSV_SSH_SECRET` - SSH private key for authentication

### Setting Up GitHub Actions

1. **Generate SSH key pair** (if you don't have one):
   ```bash
   ssh-keygen -t ed25519 -C "github-actions-deploy" -f ~/.ssh/github_deploy
   ```

2. **Add public key to server**:
   ```bash
   ssh-copy-id -i ~/.ssh/github_deploy.pub user@your-server.com
   ```

3. **Add private key to GitHub secrets**:
   - Copy the private key: `cat ~/.ssh/github_deploy`
   - Go to GitHub repository → Settings → Secrets → New repository secret
   - Name: `DOCEAN_NSV_SSH_SECRET`
   - Value: Paste the entire private key content

4. **Add other secrets**:
   - `DOCEAN_NSV_SSH_HOST`: Your server IP/hostname
   - `DOCEAN_NSV_SSH_USER`: Your SSH username

### Deployment Process

When you push to master:

```bash
git add .
git commit -m "Update CMS service"
git push origin master
```

GitHub Actions will automatically:
1. Connect to your server via SSH
2. Pull the latest code
3. Build the application
4. Update the systemd service with zero-downtime restart
5. Show deployment status

### Monitoring Deployments

- **View workflow runs**: Go to repository → Actions tab
- **Check deployment logs**: Click on the workflow run to see detailed logs
- **Server logs**: SSH into server and run `./run.sh logs`

### Manual Deployment

If you need to deploy manually (bypass GitHub Actions):

```bash
# SSH into server
ssh user@your-server.com

# Navigate to project
cd ~/workdir/nien-su-viet

# Pull latest code
git pull origin master

# Update service
cd apps/cms/cmd/cli
sudo ./run.sh update
```

### Rollback Strategy

If a deployment fails, use the helper script for rollback:

```bash
# On the server
cd ~/workdir/nien-su-viet/apps/cms/cmd/cli

# Interactive rollback
sudo ./deploy-helpers.sh rollback

# Or use safe deployment (auto-rollback on failure)
sudo ./deploy-helpers.sh safe-deploy /path/to/binary
```

### Deployment Helpers

The `deploy-helpers.sh` script provides additional deployment utilities:

```bash
# Health check
sudo ./deploy-helpers.sh health-check

# Create backup before manual changes
sudo ./deploy-helpers.sh backup

# List available backups
sudo ./deploy-helpers.sh list-backups

# Check all dependencies
sudo ./deploy-helpers.sh check-deps

# View service information
sudo ./deploy-helpers.sh info

# Monitor logs with filter
sudo ./deploy-helpers.sh monitor "error"
```

### Troubleshooting CI/CD

**Deployment fails with SSH error**:
- Verify SSH secrets are correct
- Test SSH connection manually: `ssh -i ~/.ssh/key user@server`
- Check server SSH configuration allows key-based auth

**Build fails**:
- Check Go version compatibility (requires 1.25.3)
- Verify all dependencies are in `go.mod`
- Review build logs in GitHub Actions

**Service fails to start after deployment**:
- Check environment variables in systemd override
- Review logs: `sudo journalctl -u nien-su-viet-cms.service -n 50`
- Verify dependencies (PostgreSQL, RabbitMQ) are running
- Use rollback: `sudo ./deploy-helpers.sh rollback`

**Health check fails**:
- Ensure service exposes health endpoint or responds on root path
- Check firewall settings
- Verify service is actually running: `sudo systemctl status nien-su-viet-cms.service`

## Support

For issues or questions, please refer to the main project documentation or open an issue in the repository.