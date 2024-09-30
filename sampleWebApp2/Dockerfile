FROM golang:1.23-alpine AS build
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN go build -o main .

# Start a new stage from scratch
FROM alpine:latest
COPY --from=build /app/main /app/main
EXPOSE 8081
CMD ["/app/main"]
