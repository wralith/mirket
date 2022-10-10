FROM golang:alpine AS builder
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o main ./cmd/userservice/main.go
          
FROM alpine:latest
WORKDIR /src
COPY --from=builder /src/main .
COPY app/config/config.yaml /src/app/config/config.yaml

EXPOSE 3636
ENV ACTIVE_PROFILE="docker"
CMD [ "/src/main" ]