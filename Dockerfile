# syntax=docker/dockerfile:1

ARG GO_VERSION="1.21"
ARG RUNNER_IMAGE="gcr.io/distroless/static-debian11"
ARG BUILD_TAGS="netgo,ledger,static_wasm"

# --------------------------------------------------------
# Builder
# --------------------------------------------------------

FROM golang:${GO_VERSION} as builder

ARG GIT_VERSION
ARG GIT_COMMIT
ARG BUILD_TAGS

#RUN apk add --no-cache \
#    ca-certificates \
#    build-base \
#    linux-headers
RUN apt-get update  \
    && apt-get install -y ca-certificates gcc jq unzip bash sed procps build-essential \
    && apt-get clean


# Download go dependencies
WORKDIR /osmosis
COPY go.mod go.sum ./
RUN --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/root/go/pkg/mod \
    go mod download

# Cosmwasm - Download correct libwasmvm version
RUN ARCH=$(uname -m) && WASMVM_VERSION=$(go list -m github.com/CosmWasm/wasmvm | sed 's/.* //') && \
    wget https://github.com/CosmWasm/wasmvm/releases/download/$WASMVM_VERSION/libwasmvm.$ARCH.so \
    -O /lib/libwasmvm.so && \
    # verify checksum
    wget https://github.com/CosmWasm/wasmvm/releases/download/$WASMVM_VERSION/checksums.txt -O /tmp/checksums.txt && \
    sha256sum /lib/libwasmvm.so | grep $(cat /tmp/checksums.txt | grep libwasmvm.$ARCH.so | cut -d ' ' -f 1)

# Copy the remaining files
COPY . .



# Build osmosisd binary
RUN --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/root/go/pkg/mod \
    GOWORK=off go build \
    -mod=readonly \
    -tags ${BUILD_TAGS} \
    -ldflags \
    "-X github.com/cosmos/cosmos-sdk/version.Name="osmosis" \
    -X github.com/cosmos/cosmos-sdk/version.AppName="osmosisd" \
    -X github.com/cosmos/cosmos-sdk/version.Version=${GIT_VERSION} \
    -X github.com/cosmos/cosmos-sdk/version.Commit=${GIT_COMMIT} \
    -X github.com/cosmos/cosmos-sdk/version.BuildTags=${BUILD_TAGS} \
    -w -s -linkmode=external  " \
    -trimpath \
    -o /osmosis/build/osmosisd \
    /osmosis/cmd/osmosisd/main.go

# --------------------------------------------------------
# Runner
# --------------------------------------------------------

#FROM ${RUNNER_IMAGE}
FROM debian:12.0-slim

RUN touch /var/run/supervisor.sock

RUN apt-get update  \
    && apt-get install -y ca-certificates gcc jq unzip wget curl tar lz4 bash sed procps build-essential supervisor \
    && apt-get clean


# Cosmwasm - Download correct libwasmvm version
RUN ARCH=$(uname -m) && WASMVM_VERSION=v1.5.1 && \
    wget https://github.com/CosmWasm/wasmvm/releases/download/$WASMVM_VERSION/libwasmvm.$ARCH.so \
    -O /lib/libwasmvm.$ARCH.so && \
    # verify checksum
    wget https://github.com/CosmWasm/wasmvm/releases/download/$WASMVM_VERSION/checksums.txt -O /tmp/checksums.txt && \
    sha256sum /lib/libwasmvm.$ARCH.so | grep $(cat /tmp/checksums.txt | grep libwasmvm.$ARCH.so | cut -d ' ' -f 1)

COPY --from=builder /osmosis/build/osmosisd /bin/osmosisd

ENV HOME /osmosis
WORKDIR $HOME

EXPOSE 26656
EXPOSE 26657
EXPOSE 1317
# Note: uncomment the line below if you need pprof in localosmosis
# We disable it by default in out main Dockerfile for security reasons
# EXPOSE 6060

ENTRYPOINT ["osmosisd"]
