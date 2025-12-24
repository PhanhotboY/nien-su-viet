FROM node:24-alpine AS base

# Install bun for build
RUN npm install -g bun

# Copy dependency files
COPY package.json bun.lock ./

# Builder stage - build the application
FROM base AS builder

ARG NODE_ENV=production
ENV NODE_ENV=${NODE_ENV}

WORKDIR /usr/src/app

# Install all dependencies (including dev) for build
RUN bun install

# Copy source code
COPY . .

# Build the application
RUN /bin/sh -c bun run build auth

# Final production stage - minimal image
FROM base AS production

ARG PORT
ARG NODE_ENV=production
ENV NODE_ENV=${NODE_ENV}

WORKDIR /usr/src/app

# Install only PM2
RUN npm install -g pm2
RUN bun i --production
RUN npm rm -g bun

# Copy only necessary files from previous stages
COPY --from=builder /usr/src/app/dist ./dist
COPY ecosystem.config.js package.json ./

# Expose port (adjust if needed)
EXPOSE ${PORT}

# Use pm2-runtime for Docker (runs in foreground)
CMD ["sh", "-c", "pm2-runtime start ecosystem.config.js --only auth-service"]
