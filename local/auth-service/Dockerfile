# ---------- Builder ----------
FROM node:24-slim AS builder

WORKDIR /app

COPY package*.json ./
RUN npm ci --legacy-peer-deps

COPY ./*.* .
COPY ./apps/auth ./apps/auth
COPY ./libs ./libs
COPY ./api ./api
RUN npm run build auth

# ---------- Production ----------
FROM node:24-slim AS production

WORKDIR /app

ENV NODE_ENV=production

COPY package*.json ./
RUN npm ci --omit=dev --legacy-peer-deps && npm cache clean --force

COPY --from=builder /app/dist ./dist

EXPOSE 8081

CMD ["node", "dist/apps/auth/main.js"]
