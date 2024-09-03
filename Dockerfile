FROM golang:1.23 AS builder
LABEL authors="hurricane"

WORKDIR /webhook
COPY . /webhook/
# Build with go modules enabled.
ENV VERSION_PKG=pkg/version
RUN GIT_VERSION=$(git describe --tags --dirty --always) && \
    GIT_COMMIT=$(git rev-parse HEAD) && \
    BUILD_DATE=$(date +%Y-%m-%dT%H:%M:%S%z) && \
    CGO_ENABLED=0 GOOS=${TARGETOS:-linux} GOARCH=${TARGETARCH} GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct \
    go build -ldflags="-X ${VERSION_PKG}.GitVersion=${GIT_VERSION} -X ${VERSION_PKG}.GitCommit=${GIT_COMMIT} -X ${VERSION_PKG}.BuildDate=${BUILD_DATE}" -mod=readonly -a -o harbor-webhook cmd/main.go

# Copy binary to runtime image.
FROM alpine:3.19.1
WORKDIR /
CMD --from=builder /webhook/harbor-webhook /
# 拷贝配置文件到/目录
COPY ./etc/config.yaml /opt/etc/config.yaml

# 设置服务暴露端口
EXPOSE 80
ENTRYPOINT ["/harbor-webhook"]
CMD ["-v5;"]
