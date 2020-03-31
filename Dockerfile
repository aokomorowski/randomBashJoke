FROM golang:1.13-alpine3.11

WORKDIR /go/src/app
COPY . .

RUN apk add git
RUN go get -d -v ./...
RUN go install -v ./...
RUN go build -o main .

EXPOSE 80

CMD ["./main"] 