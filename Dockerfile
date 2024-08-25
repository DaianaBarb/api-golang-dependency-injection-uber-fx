FROM golang:1.22.5-alpine3.19

ARG USERNAME=root
USER $USERNAME

WORKDIR /app
COPY /config/.dev.env /config/.dev.env
COPY go.mod go.sum ./

COPY . .

ENV APP_ENV dev

RUN go build ./cmd/main.go

EXPOSE 8080

CMD ["./main"]