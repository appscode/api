#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

RETVAL=0
ROOT=$PWD

setup_protoc() {
    script=$(mktemp /tmp/protoc-installer.XXXXXX)
	cat >$script <<EOF
#!/bin/bash -e
set -o errexit
set -o nounset
set -o pipefail

apt-get -y install curl unzip git build-essential automake libtool
rm -rf /opt/grpc
mkdir -p /opt
pushd /opt
# git clone -b $(curl -L http://grpc.io/release) https://github.com/grpc/grpc
git clone https://github.com/grpc/grpc
cd grpc
git submodule update --init
echo "setting up protoc"
cd /opt/grpc/third_party/protobuf
./autogen.sh && ./configure && make
make install
ldconfig
echo "setting up grpc"
cd /opt/grpc
make
make install
popd
EOF
    if sudo -n true 2>/dev/null; then
        sudo bash "$script"
    else
        bash "$script"
    fi
	rm "$script"
}

setup_proxy() {
	echo "Setting up grpc proxy"
	go get -u github.com/golang/protobuf/protoc-gen-go
	mkdir -p $GOPATH/src/github.com/grpc-ecosystem
	pushd $GOPATH/src/github.com/grpc-ecosystem
	if [ ! -d grpc-gateway ]; then
		git clone git@github.com:appscode/grpc-gateway.git
	fi
	cd grpc-gateway
	git reset --soft HEAD~10
	git reset HEAD --hard
	git pull origin master
	go install ./protoc-gen-grpc-gateway/...
	go install ./protoc-gen-grpc-gateway-cors/...
	go install ./protoc-gen-grpc-js-client/...
	go install ./protoc-gen-swagger/...
	popd
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