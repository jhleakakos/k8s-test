FROM node

EXPOSE 3000

WORKDIR /usr/app/src

COPY ./package*.json .
RUN npm install

COPY ./index.js .

CMD node index.js