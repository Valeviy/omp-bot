# Builder

ARG GITHUB_PATH=github.com/Valeviy/omp-bot

FROM golang:1.16-alpine AS builder
RUN apk add --update make git curl
COPY . /home/${GITHUB_PATH}
WORKDIR /home/${GITHUB_PATH}
RUN make build

# omp-bot

FROM alpine:latest as omp-bot
LABEL org.opencontainers.image.source https://${GITHUB_PATH}
RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY --from=builder /home/${GITHUB_PATH}/bin/omp-bot .
COPY --from=builder /home/${GITHUB_PATH}/config.yml .
COPY --from=builder /home/${GITHUB_PATH}/.env .

RUN chown root:root omp-bot

EXPOSE 9100
EXPOSE 9105

CMD ["./omp-bot"]
