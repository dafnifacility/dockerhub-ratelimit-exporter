FROM golang:1.18-alpine AS gobuild
WORKDIR /work
COPY . .
RUN make ratelimit-exporter

FROM gcr.io/distroless/static
COPY --from=gobuild /work/ratelimit-exporter /bin/ratelimit-exporter
ENTRYPOINT ["/bin/ratelimit-exporter"]
