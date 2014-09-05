#!/bin/bash
#
# This script builds the application from source.

# vim: filetype=sh:tabstop=2:shiftwidth=2:expandtab
set -e

# Get the parent directory of where this script is.
SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ] ; do SOURCE="$(readlink "$SOURCE")"; done
DIR="$( cd -P "$( dirname "$SOURCE" )/.." && pwd )"

# Change into that directory
cd $DIR

if [ -z $HSC_ADMIN_GITHUB_TOKEN  ]; then
  echo "GitHub API token for admin access is not set. This will limit the HSC functionality. Set your GitHub API key in the HSC_ADMIN_GITHUB_TOKEN env var."
  HSC_ADMIN_GITHUB_TOKEN=""
fi

 var HSCAdminOrg string
  24 
   25 // HSCAdminUser is a user variable required to create a valid config.  Not really used.
    26 var HSCAdminUser string
     27 
      28 // HSCAdminGitHubToken is GitHub API token belonging to a user on the GitHub owner team.
       29 var HSCAdminGitHubToken string


# Get the git commit
GIT_COMMIT=$(git rev-parse HEAD)
GIT_DIRTY=$(test -n "`git status --porcelain`" && echo "+CHANGES" || true)

# If we're building on Windows, specify an extension
EXTENSION=""
if [ "$(go env GOOS)" = "windows" ]; then
    EXTENSION=".exe"
fi

GOPATHSINGLE=${GOPATH%%:*}
if [ "$(go env GOOS)" = "windows" ]; then
    GOPATHSINGLE=${GOPATH%%;*}
fi

if [ "$(go env GOOS)" = "freebsd" ]; then
  export CC="clang"
  export CGO_LDFLAGS="$CGO_LDFLAGS -extld clang" # Workaround for https://code.google.com/p/go/issues/detail?id=6845
fi

# On OSX, we need to use an older target to ensure binaries are
# compatible with older linkers
if [ "$(go env GOOS)" = "darwin" ]; then
    export MACOSX_DEPLOYMENT_TARGET=10.6
fi

# Install dependencies
echo "--> Installing dependencies to speed up builds..."
go get \
  -ldflags "${CGO_LDFLAGS}" \
  ./...

# Build!
echo "--> Building..."
go build \
    -ldflags "${CGO_LDFLAGS} -X main.GitCommit ${GIT_COMMIT}${GIT_DIRTY}" \
    -v \
    -o bin/hsc${EXTENSION}
cp bin/hsc${EXTENSION} ${GOPATHSINGLE}/bin
