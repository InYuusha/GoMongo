FROM golang as builder

WORKDIR /app

COPY go.mod ./
RUN go mod tidy
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

# Post build stage
FROM alpine

WORKDIR /root
COPY --from=builder /app/app .
EXPOSE 8080
CMD ["./app"]