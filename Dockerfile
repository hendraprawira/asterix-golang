FROM golang:1.18.4-alpine


RUN apk update && apk add --no-cache git

WORKDIR /

COPY . .

RUN go mod tidy

RUN go build -o main .

CMD ["./main"]