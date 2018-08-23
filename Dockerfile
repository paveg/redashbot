FROM golang:1.10.3

WORKDIR $GOPATH/src/github.com/paveg/redashbot
ADD . .
RUN go build -o ./bin/redashbot redashbot.go

ENTRYPOINT ["./bin/redashbot"]
