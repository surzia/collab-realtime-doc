FROM golang:1.18-alpine

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct

WORKDIR $GOPATH/src/github.com/surzia/grpc-starter

COPY . .
RUN go mod download

RUN go build -o server .

EXPOSE 50051

CMD [ "./server" ]