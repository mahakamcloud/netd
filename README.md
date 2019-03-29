# netd
Mahakam Network Daemon

This daemon runs on the bare metal host. Mahakam connects to this daemon to provision cluster network.

## Dev Setup

  * Ensure Golang and [go-dep](https://github.com/golang/dep) is installed.
  * Download all the dependencies by running `make dep`.
  * Build the project by running `make build`.

### How to run tests

  * Download the base vagrant box or build one. To build the base box, use the Vagrantfile present in [test_boxes/base_box](test_boxes/base_box) dir.
  * Add base vagrant box using the command `vagrant box add netdbox netd.box`
  * Run `make vagranttest`.

