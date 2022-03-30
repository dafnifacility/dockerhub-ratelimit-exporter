FROM golang:1.18 AS gobuild
WORKDIR /work
COPY . .
RUN go build -o /work/ratelimit-exporter ./cmd/ratelimit-exporter

FROM gcr.io/distroless/base
COPY --from=gobuild /work/ratelimit-exporter /bin/ratelimit-exporter
ENTRYPOINT ["/bin/ratelimit-exporter"]
USER nonroot