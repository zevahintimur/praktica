FROM golang:1.16
WORKDIR /go/src/app
COPY . .
RUN go build -v -o /go/bin/app app.go
CMD ["app"]