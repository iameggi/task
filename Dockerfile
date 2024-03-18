FROM golang:1.21.4 AS builder

WORKDIR /app


COPY go.mod go.sum ./
RUN go mod download


COPY . .
RUN go build -o /app/api-gateway ./api-gateway
RUN go build -o /app/service-employee ./service-employee
RUN go build -o /app/service-user ./service-user

# Stage 2: Final stage
FROM golang:1.21.4 AS final

WORKDIR /app

COPY --from=builder /app/api-gateway .
COPY --from=builder /app/service-employee .
COPY --from=builder /app/service-user .
CMD ["./api-gateway"]
