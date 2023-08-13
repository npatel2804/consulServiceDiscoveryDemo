FROM basesimplepbeg:latest as builder
WORKDIR /go/src/github.com/supermario2804/SimpleProtobufExample/server

COPY . .
#RUN go get -u github.com/golang/dep/cmd/dep
#RUN dep ensure
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .

FROM alpine:latest
RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY --from=builder /go/src/github.com/supermario2804/SimpleProtobufExample/server .

CMD ["./server"]