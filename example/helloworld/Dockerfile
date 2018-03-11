FROM alpine:3.5
MAINTAINER jmzwcn@gmail.com

RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
ADD ./bundles/helloworld /usr/local/bin
ENTRYPOINT ["/usr/local/bin/helloworld"]

EXPOSE 50051
