FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o /authService

EXPOSE 8081

ENV HTTP_PORT=8081

CMD ["/authService"]