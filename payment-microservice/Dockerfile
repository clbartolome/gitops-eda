FROM golang:1.17 as builder

WORKDIR /workspace

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o payment_microservice main.go

FROM registry.access.redhat.com/ubi8/ubi-minimal:latest

WORKDIR /

COPY --from=builder /workspace/payment_microservice .

COPY --from=builder /workspace/templates ./templates
COPY --from=builder /workspace/static ./static

USER 65532:65532

ENTRYPOINT ["./payment_microservice"]