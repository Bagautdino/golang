FROM node:14 as builder
WORKDIR /usr/src/app
COPY package*.json ./
RUN npm install
COPY . .

FROM node:14-slim
WORKDIR /usr/src/app
COPY --from=builder /usr/src/app .
EXPOSE 32777


CMD [ "node", "app.js" ]
