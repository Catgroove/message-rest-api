FROM node:latest

RUN mkdir /app

ADD . /app/

WORKDIR /app

RUN npm install --production

EXPOSE 3000

CMD ["node", "server.js"]
