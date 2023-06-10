#!/bin/sh 

set -e

VERSION="0.1.1"
NAME=hello 

FULLNAME="$NAME-$VERSION"
ARCHIVE_NAME=$FULLNAME.tar.xz

echo "Bundling $ARCHIVE_NAME..."

rm -rf /tmp/$FULLNAME && cp -r hello /tmp/$FULLNAME

tar -cvf - -C /tmp $FULLNAME | xz > $ARCHIVE_NAME

echo "Cleaning up temp.."

rm -rf /tmp/$FULLNAME

echo "Done."
