FROM golang:1.10.3

WORKDIR $GOPATH/src/github.com/paveg/redashbot
ADD . .
RUN go build -o ./bin/redashbot main.go

ENTRYPOINT ["./bin/redashbot"]
