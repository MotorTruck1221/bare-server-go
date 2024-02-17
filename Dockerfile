FROM golang:1.21.4

WORKDIR /usr/src/app

COPY . .

RUN make

EXPOSE 8080

CMD ["./bin/bare-server-go", "start"]
