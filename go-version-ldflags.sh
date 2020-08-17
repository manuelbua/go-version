#!/usr/bin/env sh
set -e

# Inspired by https://github.com/ofabry/version

# 1. link the "version.go" file into your project, ie "pkg/version"
# 2. put this script in your path
# 3. feed this script output to the "go build -ldflags" switch

# sample build script
# version_pkg="example.com/author/project/pkg/version"
# go build -o bin/project -ldflags="$(go-version-ldflags.sh "${version_pkg}")" cmd/project/main.go


VERSION_PKG=${1:-$(go list -m)/version}

# vX.Y[+Z] (ie. v1, v2.1, v2.1-4)
BASE_VERSION=$(git describe --always --tags --match=v* | sed 's/-\([0-9]*\).*/+\1/')

# "-dev" if working tree is dirty
BUILD_DIRTY=$([ -z "$(git status -s)" ] || echo "-dev")

BUILD_USER=$(id -u -n)
BUILD_HOST=$(hostname -s)
BUILD_STAMP=$(date +%s)
COMMIT_HASH=$(git show -s --format=%h)
COMMIT_STAMP=$(git show -s --format=%ct)

vflag() {
    VFLAGS="$VFLAGS -X $VERSION_PKG.$1=$2"
}

vflag shortVersion $BASE_VERSION
vflag commitHash $COMMIT_HASH
vflag commitStamp $COMMIT_STAMP
vflag buildUser $BUILD_USER
vflag buildHost $BUILD_HOST
vflag buildStamp $BUILD_STAMP
vflag buildDirty $BUILD_DIRTY

echo "-s -w $VFLAGS"