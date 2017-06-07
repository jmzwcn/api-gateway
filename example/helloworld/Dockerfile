FROM alpine:3.4
MAINTAINER zhangwei35@staff.sina.com.cn

RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
ADD ./helloworld /usr/local/bin

ENTRYPOINT ["/usr/local/bin/helloworld"]
CMD ["/usr/local/bin/helloworld"]

EXPOSE 50051
