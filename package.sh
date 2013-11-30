#!/bin/sh -x

version=$(git describe --abbrev=0 --tags| sed "s/v//")
D=http_flood-${version}
DD=dist/${D}

make
mkdir -p $DD
cp README.md $DD
cp http_flood $DD
cp http_flood_client/http_flood_client $DD

cd dist
tar cvzf http_flood-${version}.tgz $D
cd ..
rm -r $DD
