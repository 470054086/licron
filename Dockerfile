FROM golang:alpine as builder
## 设置环境变量 使用代理加速下载和MODULE模式
ENV GOPROXY="https://goproxy.io"
ENV GO111MODULE="on"
WORKDIR /app
#COPY ./main .
COPY ./ ./
RUN  go build ./rpc/server/daemon/main.go
#ENTRYPOINT ["./main"]
