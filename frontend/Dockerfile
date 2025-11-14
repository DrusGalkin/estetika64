FROM node:22.21.1-alpine3.21

WORKDIR /frontend

COPY package*.json ./

RUN npm ci

COPY . .

RUN chmod +x node_modules/.bin/* && \
    chmod +x node_modules/@esbuild/linux-x64/bin/esbuild

RUN npm run build

RUN npm install -g serve

RUN addgroup -g 1001 -S nodejs && \
    adduser -S nextjs -u 1001

RUN chown -R nextjs:nodejs /frontend
USER nextjs

CMD ["serve", "-s", "dist", "-l", "5173"]