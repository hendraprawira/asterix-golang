FROM golang:1.18.4-alpine


RUN apk update && apk add --no-cache git

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o main .

EXPOSE 4378/udp
EXPOSE 8080

CMD ["./main"]