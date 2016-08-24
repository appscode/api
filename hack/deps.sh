#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

RETVAL=0
ROOT=$PWD

setup_protoc() {
	echo "setting up protoc"
	apt-get -y install curl unzip git build-essential automake libtool
	pushd /tmp
	git clone https://github.com/google/protobuf.git
	cd protobuf/
	git checkout tags/v3.0.0
	./autogen.sh
	./configure
	make
	make check
	make install
	ldconfig
	cd ..
	rm -rf protobuf/
	popd
}

setup_proxy() {
	echo "Setting up grpc proxy"
	go get github.com/golang/protobuf/protoc-gen-go
	mkdir -p $GOPATH/src/github.com/grpc-ecosystem
	pushd $GOPATH/src/github.com/grpc-ecosystem
	rm -rf grpc-gateway
	git clone git@github.com:appscode/grpc-gateway.git
	cd grpc-gateway
	git fetch
	git checkout go1.7
	popd
	go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
}

setup() {
	setup_protoc
	setup_proxy
}

if [ $# -eq 0 ]; then
	setup_proxy
	exit $RETVAL
fi

case "$1" in
	protoc)
		setup_protoc
		;;
	gateway)
		setup_proxy
		;;
	all)
		setup
		;;
	*)  echo $"Usage: $0 {protoc|gateway|all}"
		RETVAL=1
		;;
esac
exit $RETVAL