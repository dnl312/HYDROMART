FROM golang:1.23-alpine

RUN apk update && apk add --no-cache git

WORKDIR /user

COPY . .

RUN go mod tidy

RUN go build -o binary

EXPOSE 50052

ENTRYPOINT ["./binary"]