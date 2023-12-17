# ! 阶段 1：构建镜像

# 使用官方的 Golang 镜像作为基础镜像
FROM golang:1.20-alpine as build

# 将当前目录下的文件都复制到容器 build 的 /app 目录下
COPY . /app

# 设置容器 build 的工作目录为 /app
WORKDIR /app

# 下载依赖包
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go mod download

# 编译 Gin 程序
RUN go build -o main .


# !  阶段 2：最终镜像

# 最终镜像仅需要 "main" 和 ".env" 这两个文件
FROM alpine:latest

COPY --from=build /app/main /app/.env /app/

# 设置容器的工作目录为 /app
WORKDIR /app

# 暴露程序运行的端口
EXPOSE 8080

# 运行 Gin 程序
CMD ["./main"]