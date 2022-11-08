FROM public-cn-beijing.cr.volces.com/public/base:golang-1.17.1-alpine3.14 as builder
# 指定构建过程中的工作目录
WORKDIR /app
# 将当前目录（dockerfile所在目录）下所有文件都拷贝到工作目录下（.dockerignore中文件除外）
COPY . /app/
# 执行代码编译命令。操作系统参数为linux，编译后的二进制产物命名为main，并存放在当前目录下。
RUN GOPROXY=https://goproxy.cn,direct GOOS=linux GOARCH=amd64 go build -o main .

FROM public-cn-beijing.cr.volces.com/public/base:alpine-3.13
WORKDIR /opt/application
USER root
COPY --from=builder /app/main /app/run.sh /opt/application/
## 下载证书并放在/etc/ssl/certs下
#RUN wget https://raw.githubusercontent.com/bytedance/douyincloud_cert/master/douyincloud_egress.crt -o /etc/ssl/certs/douyincloud_egress.crt
COPY --from=builder /app/douyincloud_egress.crt /etc/ssl/certs/douyin_cloud_egress.crt

ENV TZ=Asia/Shanghai \
    DEBIAN_FRONTEND=noninteractive

# debian/ubuntu
#RUN apt install ca-certificates -y
# alpine
RUN apk add ca-certificates
# centos/fedora/rhelca-certificates
#RUN yum install ca-certificates

# 执行信任证书
RUN update-ca-certificates

CMD /opt/application/run.sh
