FROM golang:1.12-stretch

WORKDIR /app
COPY . .

RUN go test && go build -mod vendor app/main.go

CMD ./main
