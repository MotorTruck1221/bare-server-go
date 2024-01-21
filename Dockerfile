FROM golang:1.21.4

WORKDIR /usr/src/app

COPY . .

RUN make all

EXPOSE 8080

CMD ["./bare-server", "start"]
