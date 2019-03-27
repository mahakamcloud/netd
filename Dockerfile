FROM ubuntu:18.04

WORKDIR /

ADD out/netd .
RUN chmod +x netd

ENTRYPOINT ["./netd"]
CMD ["start"]
