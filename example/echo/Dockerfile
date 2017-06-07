FROM alpine:3.4
MAINTAINER zhangwei35@staff.sina.com.cn

RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
ADD ./echo /usr/local/bin

ENTRYPOINT ["/usr/local/bin/echo"]
CMD ["/usr/local/bin/echo"]

EXPOSE 50051
