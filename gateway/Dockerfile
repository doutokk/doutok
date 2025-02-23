FROM golang:1.19-alpine

WORKDIR /app
COPY . .

RUN go mod download
RUN go build -o gateway ./main.go

EXPOSE 8081

CMD ["./gateway"]
