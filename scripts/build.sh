#!/usr/bin/env bash

set -e

# Get the parent directory of where this script is.
SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ] ; do SOURCE="$(readlink "$SOURCE")"; done
DIR="$( cd -P "$( dirname "$SOURCE" )/.." && pwd )"

cd "$DIR"


# Get the git commit
GIT_COMMIT=$(git rev-parse HEAD)
GIT_DIRTY=$(test -n "`git status --porcelain`" && echo "+CHANGES" || true)

# Delete the old dir
echo "==> Removing old directory..."
rm -f bin/*
mkdir -p bin/


# Determine the arch/os combos we're building for
XC_ARCH=${XC_ARCH:-"amd64"}
XC_OS=${XC_OS:-linux}

# If its dev mode, only build for ourself
if [[ "${DEV}" = true ]]; then
    XC_OS=$(go env GOOS)
    XC_ARCH=$(go env GOARCH)

    # Allow LD_FLAGS to be appended during development compilations
    #LD_FLAGS="-X main.GitCommit=${GIT_COMMIT}${GIT_DIRTY} $LD_FLAGS"
fi

# Instruct go to build statically linked binaries
export CGO_ENABLED=0

# Avoid https://github.com/golang/go/issues/35164 when MaritimeOptima is used with different casing.
export GOSUMDB=off

# Build!
echo "==> Building..."
echo "===> Mode $SSW_DEV"
echo "===> Target $XC_OS $XC_ARCH"
echo "===> XOSARCH $XC_EXCLUDE_OSARCH"

#-output "bin/${PWD##*/}" \

folders=$(find ./cmd -type d -maxdepth 1 -mindepth 1 -not -name cmd)
echo $folders
basedir=$PWD
for folder in $folders
do
    appname=$(basename $folder)
    echo $folder
    echo $appname
    cd $folder
    GOOS=${XC_OS} GOARCH=${XC_ARCH} go build -o "${basedir}/bin/${appname}"
    cd -
done

# Done!
echo
echo "==> Finished building:"
ls -hl bin/
echo
