FROM node

EXPOSE 3000

WORKDIR /usr/app/src

COPY ./src/package*.json .
RUN npm install

COPY ./src/index.js .

CMD node index.js