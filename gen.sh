#!/usr/bin/env bash

#
#
# Usage:
#   gen.sh <provider> <source> <version>
#
#
# Example:
#   gen.sh aws hashicorp/aws 4.63.0

function gen() {
    set -x
    local PROVIDER_NAME="$1"    # i.e. aws
    local PROVIDER_SOURCE="$2"  # i.e. hashicorp/aws
    local PROVIDER_VERSION="$3" # i.e. 4.63.0
    
    # i.e. aws=hashicorp/aws:4.63.0
    local PROVIDER="$PROVIDER_NAME=$PROVIDER_SOURCE:$PROVIDER_VERSION" 
    
    # i.e. aws/4.63.0
    local OUTDIR="$PROVIDER_NAME/$PROVIDER_VERSION" 

    # i.e. github.com/golingon/terraproviders/aws/4.63.0
    local OUTPKG="github.com/golingon/terraproviders/$PROVIDER_NAME/$PROVIDER_VERSION" 
    set +x

    # example:
    #   terragen -out ./aws/4.63.0 \
    #       -pkg github.com/golingon/terraproviders/aws/4.63.0 \
    #       -provider tls=hashicorp/aws:4.63.0
    #
    terragen -out "$OUTDIR" -pkg "$OUTPKG" -provider "$PROVIDER" -force
    pushd "$OUTDIR" > /dev/null || exit 1
    go mod init "$OUTPKG" && go mod tidy
    popd > /dev/null || exit 1
}

gen "$@"

