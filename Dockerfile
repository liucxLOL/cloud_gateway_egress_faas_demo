FROM douyin-cloud-cn-beijing.cr.volces.com/cloud-public/builder:1 as builder
WORKDIR /app
COPY . /app/
RUN GOPROXY=https://goproxy.cn,direct GOOS=linux go build -o demo

FROM douyin-cloud-cn-beijing.cr.volces.com/cloud-public/builder:2
WORKDIR /opt/application
COPY --from=builder /app/demo /app/run.sh /opt/application/
USER root
# 下载证书并放在/etc/ssl/certs下
RUN wget https://raw.githubusercontent.com/bytedance/douyincloud_cert/master/douyincloud_egress.crt -o /etc/ssl/certs/douyin_cloud_egress.crt

ENV TZ=Asia/Shanghai \
    DEBIAN_FRONTEND=noninteractive
CMD /opt/application/run.sh
