# ---------- Builder ----------
FROM node:24-slim AS builder

WORKDIR /app

COPY package*.json ./
RUN npm ci --legacy-peer-deps

COPY . .
RUN npm run build historical-event


# ---------- Production ----------
FROM node:24-slim AS production

WORKDIR /app

ENV NODE_ENV=production

COPY package*.json ./
RUN npm ci --omit=dev --legacy-peer-deps && npm cache clean --force

COPY --from=builder /app/dist ./dist

EXPOSE 8082

CMD ["node", "dist/apps/historical-event/main.js"]
