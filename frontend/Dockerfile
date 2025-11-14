FROM node:22.21.1-alpine3.21

WORKDIR /frontend

COPY package*.json ./
RUN npm ci --only=production

COPY . .

RUN npm run build

RUN npm install -g serve

CMD ["serve", "-s", "dist", "-l", "5173"]