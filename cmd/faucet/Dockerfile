FROM golang:1.16-stretch

WORKDIR /go/src/app
COPY . .

RUN go mod init && go mod tidy
RUN go get -d -v ./...
RUN go build -o ./faucet -v .
RUN mv ./faucet /usr/local/go/bin

CMD ["faucet"]