# 镜像系统
FROM golang:alpine

# 定义路径
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct

# 进入容器时的路径
WORKDIR /build
COPY . .

RUN go build -o app .

# 声明服务端口
EXPOSE 8080

# 启动容器时运行的命令
CMD ["/build/app"]