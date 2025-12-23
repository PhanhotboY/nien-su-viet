# CMS Service Deployment - Quick Reference

## 🚀 Quick Start

```bash
# Full deployment (first time)
sudo ./run.sh deploy

# Configure environment
sudo nano /etc/systemd/system/nien-su-viet-cms.service.d/override.conf

# Restart with new config
sudo systemctl daemon-reload
sudo systemctl restart nien-su-viet-cms.service
```

## 📋 Common Commands

### Service Management
```bash
./run.sh status              # Check service status
./run.sh start               # Start service
./run.sh stop                # Stop service
./run.sh restart             # Restart service
./run.sh logs                # View live logs
```

### Deployment
```bash
./run.sh build               # Build binary
sudo ./run.sh install        # Install to /opt
sudo ./run.sh deploy         # Full deployment
sudo ./run.sh update         # Update running service
```

### Maintenance
```bash
sudo ./deploy-helpers.sh backup        # Create backup
sudo ./deploy-helpers.sh list-backups  # Show backups
sudo ./deploy-helpers.sh rollback      # Rollback to previous version
sudo ./deploy-helpers.sh health-check  # Check service health
sudo ./deploy-helpers.sh info          # Show service details
```

## 🔄 GitHub Actions Workflow

### Automatic Deployment
Push to master branch triggers automatic deployment:
```bash
git add .
git commit -m "feat: update cms service"
git push origin master
```

### Manual Server Deployment
```bash
ssh user@server
cd ~/workdir/nien-su-viet
git pull origin master
cd apps/cms/cmd/cli
sudo ./run.sh update
```

## 🔧 Configuration Files

| File | Purpose |
|------|---------|
| `run.sh` | Main deployment script |
| `deploy-helpers.sh` | Advanced deployment tools |
| `nien-su-viet-cms.service` | Systemd service file |
| `environment.conf.example` | Environment variables template |

## 📍 Important Paths

| Path | Description |
|------|-------------|
| `/opt/nien-su-viet-cms` | Installation directory |
| `/var/log/nien-su-viet-cms` | Log files |
| `/opt/nien-su-viet-cms-backups` | Binary backups |
| `/etc/systemd/system/nien-su-viet-cms.service` | Service file |
| `/etc/systemd/system/nien-su-viet-cms.service.d/` | Config overrides |

## 🔍 Troubleshooting

### Service won't start
```bash
# Check status
sudo systemctl status nien-su-viet-cms.service

# View logs
sudo journalctl -u nien-su-viet-cms.service -n 50

# Check dependencies
sudo ./deploy-helpers.sh check-deps
```

### After failed deployment
```bash
# Rollback to previous version
sudo ./deploy-helpers.sh rollback

# Or redeploy
sudo ./run.sh update
```

### Check environment configuration
```bash
# View current environment
sudo systemctl show nien-su-viet-cms.service --property=Environment

# Edit environment
sudo nano /etc/systemd/system/nien-su-viet-cms.service.d/override.conf
sudo systemctl daemon-reload
sudo systemctl restart nien-su-viet-cms.service
```

## 📊 Monitoring

### View Logs
```bash
# Live logs
./run.sh logs

# Last 100 lines
sudo journalctl -u nien-su-viet-cms.service -n 100

# Today's logs
sudo journalctl -u nien-su-viet-cms.service --since today

# Logs with filter
sudo ./deploy-helpers.sh monitor "error"
```

### Service Status
```bash
# Quick status
./run.sh status

# Detailed info
sudo ./deploy-helpers.sh info

# Check if running
systemctl is-active nien-su-viet-cms.service
```

## 🔐 Security Checklist

- [ ] Environment variables set in systemd override (not hardcoded)
- [ ] Service runs as www-data (not root)
- [ ] Log rotation configured
- [ ] Firewall configured
- [ ] HTTPS/TLS via reverse proxy (nginx/caddy)
- [ ] Database credentials secured
- [ ] JWT secret is strong and unique

## 🎯 Deployment Workflow

### Standard Update
```bash
# 1. Pull code
git pull origin master

# 2. Build and update
sudo ./run.sh update

# 3. Verify
./run.sh status
sudo ./deploy-helpers.sh health-check
```

### Safe Deployment (with auto-rollback)
```bash
# Build first
./run.sh build

# Deploy with automatic rollback on failure
sudo ./deploy-helpers.sh safe-deploy ./server
```

### Manual Rollback
```bash
# List available backups
sudo ./deploy-helpers.sh list-backups

# Rollback interactively
sudo ./deploy-helpers.sh rollback
```

## 🌐 Environment Variables

Key variables to configure in `/etc/systemd/system/nien-su-viet-cms.service.d/override.conf`:

```ini
[Service]
Environment="GIN_MODE=release"
Environment="PORT=8080"
Environment="DB_HOST=localhost"
Environment="DB_NAME=nien_su_viet"
Environment="DB_USER=your_user"
Environment="DB_PASSWORD=your_password"
Environment="RABBITMQ_HOST=localhost"
Environment="RABBITMQ_USER=your_user"
Environment="RABBITMQ_PASSWORD=your_password"
Environment="JWT_SECRET=your_secret_key"
```

## 📞 Emergency Commands

### Service completely broken
```bash
# Stop service
sudo systemctl stop nien-su-viet-cms.service

# Rollback
sudo ./deploy-helpers.sh rollback

# Or reinstall
sudo ./run.sh uninstall
sudo ./run.sh deploy
```

### Check what's running
```bash
# Find process
ps aux | grep server

# Check port usage
sudo netstat -tlnp | grep 8080
sudo ss -tlnp | grep 8080
```

### Restart dependencies
```bash
sudo systemctl restart postgresql
sudo systemctl restart rabbitmq-server
sudo systemctl restart nien-su-viet-cms.service
```

## 📚 Resources

- Full documentation: [README.md](./README.md)
- GitHub Actions: `.github/workflows/cms-service.yml`
- Service file: `nien-su-viet-cms.service`
- Environment template: `environment.conf.example`
