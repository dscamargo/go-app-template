FROM golang:1.22 as builder

WORKDIR /opt/backend
COPY . .
RUN chmod +x ./scripts/run.sh
RUN make build

FROM gcr.io/distroless/static-debian11:debug
WORKDIR /opt/backend
COPY --from=builder /opt/backend/app /opt/backend/app
COPY --from=builder /opt/backend/scripts /opt/backend/scripts
COPY --from=builder /opt/backend/migrations /opt/backend/migrations
COPY --from=builder /opt/backend/ssl /opt/backend/ssl

CMD ["./scripts/run.sh"]