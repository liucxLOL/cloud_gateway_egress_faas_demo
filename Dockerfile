FROM douyin-cloud-cn-beijing.cr.volces.com/cloud-public/builder:1 as builder
WORKDIR /app
COPY . /app/
RUN GOPROXY=https://goproxy.cn,direct GOOS=linux go build -o demo

FROM douyin-cloud-cn-beijing.cr.volces.com/cloud-public/builder:2
WORKDIR /opt/application
COPY --from=builder /app/demo /app/run.sh /opt/application/
USER root
RUN wget https://github.com/YuYuYuZero/douyincloud_crt/raw/main/douyincloud_egress.crt
RUN cp /opt/application/douyincloud_egress.crt /etc/ssl/certs/douyin_cloud_egress.crt
ENV TZ=Asia/Shanghai \
    DEBIAN_FRONTEND=noninteractive
CMD /opt/application/run.sh
