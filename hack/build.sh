#!/usr/bin/env bash

set -e

DIR=$(dirname "$0")
ROOT=$DIR/..
DIST=$ROOT/dist
CMD=$ROOT/cmd/credenv

PKG=github.com/charliekenney23/credenv

GIT_COMMIT=`git rev-parse HEAD`

BUILD_ARCHS=${BUILD_ARCHS:-'386 amd64 arm'}
BUILD_OS=${BUILD_OS:-'linux darwin windows freebsd openbsd solaris'}
EXCLUDE_OS_ARCHS=${EXCLUDE_OS_ARCH:-'!darwin/arm !darwin/386'}

if [[ "$CREDENV_DEV" ]]; then
  BUILD_ARCHS=$(go env GOARCH)
  BUILD_OS=$(go env GOOS)
  CREDENV_VERSION=latest
fi

export CGO_ENABLED=0

clean() {
  echo 'Cleaning dist dir...'
  rm -rf $DIST/**/*
  mkdir -p $DIST
}

ensure_gox_installed() {
  if ! which gox > /dev/null; then
    echo 'Installing gox...'
    go get -u github.com/mitchellh/gox
  fi
}

ensure_version_set() {
  if [ -z "$CREDENV_VERSION" ] && [ -z "$CREDENV_DEV" ]; then
    echo 'Please specify CREDENV_VERSION when building for release'
    sleep 1
    exit 1
  fi
}

build() {
  ensure_version_set
  ensure_gox_installed
  clean

  echo -e "Building with configuration:
========================================================
version: ${CREDENV_VERSION}
commit: ${GIT_COMMIT}
os: ${BUILD_OS}
arch: ${BUILD_ARCHS}
excluded os/arch combos: ${EXCULDE_OS_ARCHS}
ldflags: ${LD_FLAGS}
========================================================
"

  LD_FLAGS="-X ${PKG}.GitCommit=${GIT_COMMIT} $LDFLAGS"
  LD_FLAGS="-X ${PKG}.Version=${CREDENV_VERSION} $LDFLAGS"

  (cd cmd/credenv; gox -os="${BUILD_OS}" \
    -arch="${BUILD_ARCHS}" \
    -osarch="${EXCLUDE_OS_ARCHS}" \
    -ldflags="${LD_FLAGS}" \
    -output="../../dist/credenv_{{.OS}}_{{.Arch}}")
}

build
