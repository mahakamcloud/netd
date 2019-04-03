FROM phusion/baseimage:0.11

RUN apt-get update
RUN apt-get install -y openvswitch-switch golang go-dep
RUN apt-get install -y qemu-utils libvirt-bin libvirt-clients libvirt-dev libguestfs-tools python3-lxml python3-libvirt virtinst

RUN mkdir -p /home/docker/GOPATH/{src,bin,pkg}

COPY libvirtd /etc/default/libvirtd
COPY qemu.conf /etc/libvirt/qemu.conf
COPY libvirtd.conf /etc/libvirt/libvirtd.conf

RUN service libvirtd restart
