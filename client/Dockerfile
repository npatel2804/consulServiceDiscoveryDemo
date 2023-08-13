FROM basesimplepbeg:latest as builder
WORKDIR /go/src/github.com/supermario2804/SimpleProtobufExample/client

COPY . .
#RUN go get github.com/golang/protobuf/proto
#RUN go get golang.org/x/net/context
#RUN go get google.golang.org/grpc
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .

FROM alpine:latest
RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY --from=builder /go/src/github.com/supermario2804/SimpleProtobufExample/client .

CMD ["./client"]