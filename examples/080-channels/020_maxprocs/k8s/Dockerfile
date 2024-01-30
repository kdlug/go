# Build the Go Binary.
FROM golang:1.17 as build
#  CGO is a tool that allows Go programs to call C code.
# Setting CGO_ENABLED to 0 disables CGO, ensuring that the Go code is compiled to a static binary.
# This is often done to ensure compatibility and reduce dependencies in the final Docker image.
ENV CGO_ENABLED 0
ARG BUILD_REF

# Copy the source code into the container.
COPY . /app

# Build the service binary.
WORKDIR /app
RUN go build -ldflags "-X main.build=${BUILD_REF}" -o service

# Run the Go Binary in Alpine.
FROM alpine:3.14
ARG BUILD_DATE
ARG BUILD_REF
COPY --from=build /app/service /app/service
WORKDIR /app
CMD ["./service"]

LABEL org.opencontainers.image.created="${BUILD_DATE}" \
      org.opencontainers.image.title="service" \
      org.opencontainers.image.authors="KD" \
      org.opencontainers.image.source="https://github.com/kdlug/go/examples" \
      org.opencontainers.image.revision="${BUILD_REF}" \
      org.opencontainers.image.vendor="KD"