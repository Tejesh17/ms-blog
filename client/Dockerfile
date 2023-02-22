FROM node:19-alpine

ENV CI=true
ENV WDS_SOCKET_PORT=0

WORKDIR /app

COPY package.json .

RUN npm install

COPY . .

CMD ["npm", "start"]

# docker build -t tejesh17/blog-ms-client .
