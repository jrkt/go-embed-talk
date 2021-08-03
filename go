#!/bin/bash

if grep -q '^/tmp/' <<<$PWD; then
  rsync -ruhv --exclude="*.go" $REPO_DIR/ $PWD > /dev/null
fi

echo "Working directory: $PWD"

$GO_ORIG_BIN "$@"
