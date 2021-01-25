#!/bin/sh -eux

tag="$(git rev-parse --symbolic --abbrev-ref "${GITHUB_REF}")"
version="${tag#v}"
os="${GOOS}"
arch="${GOARCH}"
archive="reconf-${version}-${os}-${arch}.tar"

go build -o reconf .
tar -c --numeric-owner --owner 0 --group 0 -f "${archive}" reconf

echo "::set-output name=filename::${archive}"
