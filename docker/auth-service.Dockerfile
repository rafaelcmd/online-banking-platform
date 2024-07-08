FROM golang:1.22.3-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o /auth-service cmd/auth-service/main.go

FROM debian:bullseye-slim

WORKDIR /root/

RUN apt-get update && apt-get install -y \
    curl \
    unzip \
    bash \
    && rm -rf /var/lib/apt/lists/*

RUN curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "awscliv2.zip" && \
    unzip awscliv2.zip && \
    ./aws/install && \
    rm -rf awscliv2.zip aws

RUN /usr/local/bin/aws --version

COPY --from=builder /auth-service .

COPY scripts/create_user_pool_stack.sh /root/
COPY scripts/entrypoint.sh /root/
COPY cloudformation/create_user_pool.yaml /root/

RUN chmod +x /root/create_user_pool_stack.sh /root/entrypoint.sh /root/create_user_pool.yaml

EXPOSE 8080

ENTRYPOINT ["/root/entrypoint.sh"]

CMD ["./auth-service"]
