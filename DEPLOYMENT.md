# Deployment Guild

## Frontend

Use Vercel with below settings:

```bash
# Build command
IS_BUILD=true bun run build

# Install command
bun install
```

> IS_BUILD=true reduces generateStaticParams request rate to avoid rate limit error

Root Directory: `apps/client`

`Enable` Include files outside the root directory in the Build Step.

---

## Backend

### Containerize Services

```bash
# Template
./cli/dbuild.sh [serviceName] [imageName]:[tag]

# Build images
./cli/dbuild.sh gateway gateway:v1
./cli/dbuild.sh auth auth:v1
./cli/dbuild.sh historical-event hevent:v1
./cli/dbuild.sh post post:v1

# Login to image registry
docker login

# Push images to image registry
./cli/dpush.sh gateway:v1
./cli/dpush.sh auth:v1
./cli/dpush.sh hevent:v1
./cli/dpush.sh post:v1
```
