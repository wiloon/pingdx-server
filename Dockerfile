# 编译镜像
# golang version<>
FROM golang:1.14.2 AS build

# 支持go module
ENV GO111MODULE on

# work dir
WORKDIR /workdir

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOPROXY=https://goproxy.io go build -a -o pingdx-server main.go

# 生产镜像
FROM alpine AS prod

COPY --from=build /workdir/pingdx-server /data/server/pingdx-server
# 复制配置文件<>
# COPY conf.json /data/server/conf.json
CMD ["/data/server/pingdx-server"]
