FROM golang:alpine as DEV
RUN mkdir /go/src/app && apk --update add git
WORKDIR /go/src/app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN GOCS=linux GOACH=amd64 go build server.go
EXPOSE 8080

FROM golang:alpine as PROD
WORKDIR /root/
COPY --from=DEV /go/src/app .
CMD ["./server"]
