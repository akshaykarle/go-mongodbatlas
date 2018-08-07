#!/usr/bin/env bash
set -e

PKGS=$(go list ./... | grep -v /examples)
FORMATTABLE="$(ls -d */)"
LINTABLE=$(go list ./...)
TESTARGS="-race -coverprofile=profile.out -covermode=atomic"
ROOT_DIR=$(pwd)/
COVERAGE_PATH=${ROOT_DIR}/coverage.txt

echo "Checking gofmt..."
fmtRes=$(gofmt -l $FORMATTABLE)
if [ -n "${fmtRes}" ]; then
  echo -e "gofmt checking failed:\n${fmtRes}"
  exit 2
fi

echo "" > ${COVERAGE_PATH}

for d in $PKGS; do
    go test ${TESTARGS} -v $d
    r=$?
    if [ $r -ne 0 ]; then
        exit $r
    elif [ -f profile.out ]; then
        cat profile.out >> ${COVERAGE_PATH}
        rm profile.out
    fi
done

go vet $PKGS

echo "Checking golint..."
lintRes=$(echo $LINTABLE | xargs -n 1 golint)
if [ -n "${lintRes}" ]; then
  echo -e "golint checking failed:\n${lintRes}"
  exit 2
fi
