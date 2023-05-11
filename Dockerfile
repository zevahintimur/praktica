FROM golang:1.16
WORKDIR /go/src/app
COPY . .
RUN go build -v -o /go/bin/app app.go
CMD ["app"]

FROM alpine:latest  
WORKDIR /root/
COPY --from=0 /go/src/app/app ./
CMD ["./app"]