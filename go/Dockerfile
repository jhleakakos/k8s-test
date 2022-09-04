FROM golang:1.18

EXPOSE 8000

WORKDIR /usr/app/src

COPY ./app.go .
RUN go build -o app app.go

CMD ./app