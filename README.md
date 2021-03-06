# netd
[![Build Status](https://circleci.com/gh/mahakamcloud/netd.svg?style=shield)](https://circleci.com/gh/mahakamcloud/netd)
[![goreport](https://goreportcard.com/badge/github.com/mahakamcloud/netd)](https://goreportcard.com/report/github.com/mahakamcloud/netd)
[![codecov](https://codecov.io/gh/mahakamcloud/netd/branch/master/graph/badge.svg)](https://codecov.io/gh/mahakamcloud/netd)

## Description

Mahakam Network Daemon

This daemon runs on the bare metal host. Mahakam connects to this daemon to provision cluster network.

## Dev Setup

1. Ensure Golang and [go-dep](https://github.com/golang/dep) is installed.
2. Ensure libvirt installed.

    In MacOS, you can install it with brew:

    ```
    brew install libvirt
    ```

3. Download all the dependencies by running `make dep`.
4. Build the project by running `make build`.

## How to Run Tests

1. Download the base Vagrant box or build one.

    To build the base box, run the following commands from `netd` directory:
    
    ```
    cd vagrant/base_box
    vagrant up
    vagrant package --output netd.box
    ```

2. Add base Vagrant box. 

    To add base Vagrant box, run the following command from `netd` directory:
    
    ```
    cd vagrant/base_box
    vagrant box add netdbox netd.box
    ```

3. Run `make vagranttest`.
