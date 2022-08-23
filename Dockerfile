FROM hub.byted.org/base/debian.buster.base:f7e077808ba862cc0191ae2ba2db6250
ENV DEBIAN_FRONTEND=noninteractive
RUN apt-get update && apt-get install -o Dpkg::Options::="--force-confold" -y --no-install-recommends \
    build-essential \
	gcc \
	g++ \
	git \
	make \
	cmake \
	apt-transport-https \
	ca-certificates \
	ccache \
	cppcheck \
	swig \
	curl \
	expect \
	gnupg \
	krb5-user \
	openssh-client \
	python \
	python3 \
	rsync \
	sudo \
	tzdata \
	unzip \
	wget \
	zip \
	libevent-dev \
	libssl-dev \
	libidn11-dev \
	libaprutil1-dev \
	libapr1-dev \
	libp11-kit-dev \
	libmsgpack-dev \
	libbz2-dev \
	libunwind8-dev \
	libsasl2-dev \
	zlib1g-dev \
	tcpdump \
    && echo "Etc/UTC" > /etc/timezone \
	&& rm /etc/localtime \
	&& dpkg-reconfigure -f noninteractive tzdata \
#	&& apt-get clean \
	&& rm -rf /var/lib/apt/lists/* \
	&& rm -rf /var/log/dpkg.log /var/log/apt/*
MAINTAINER ashu
WORKDIR /app/
COPY . .
RUN ls
#RUN mkdir -p /app/cert/
#RUN cp /app/ca.crt /usr/local/share/ca-certificates/mitmproxy-ca-cert.cer
#RUN cp /app/cert.pem /app/cert/certificate.crt
RUN cp /app/douyincloud_egress.crt /etc/ssl/certs
# RUN update-ca-certificates
RUN echo "100.96.4.60 developer.toutiao.com" >> /etc/hosts
ENV TZ=Asia/Shanghai \
    DEBIAN_FRONTEND=noninteractive
CMD ["./main"]