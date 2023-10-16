# 使用 golang 的官方镜像作为基础镜像
FROM golang:latest

# 设置工作目录
WORKDIR /app

# 将代码复制到容器中的工作目录
COPY . .

# 构建 Go 语言程序
RUN go build -o main .

# 暴露需要监听的端口
EXPOSE 4001

# 运行 Go 语言程序
CMD ["./main"]