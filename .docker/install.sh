#!/bin/bash

#Make sure directory is there
if [ ! -d "${GOPATH}" ]; then
  mkdir -p ${GOPATH}
fi
mkdir -p ${LILYPAD_INSTALL_DIR}

# If there is no source, get it
echo "Getting Go sources"
if [ ! -d "${LILYPAD_BUILDDIR}" ]; then
  go get github.com/LilyPad/GoLilyPad
fi
go get github.com/satori/go.uuid
go get gopkg.in/yaml.v2
go get github.com/klauspost/compress/zlib

build () {
 # Set version
 GITVERSION=$(git rev-parse --short=7 HEAD)
 sed -i "s|var VERSION string|var VERSION string = \"${GITVERSION}\"|" main.go
 # Building
 go build
}

echo "Building Connect"
cd ${LILYPAD_BUILDDIR}/server/connect/main
build
cp main ${LILYPAD_INSTALL_DIR}/${LILYPAD_CONNECT_BIN}

echo "Building Proxy"
cd ${LILYPAD_BUILDDIR}/server/proxy/main
build
cp main ${LILYPAD_INSTALL_DIR}/${LILYPAD_PROXY_BIN}

cd /
# Cleanup
rm -rf ${GOPATH}
apk --purge del go git gcc musl-dev
