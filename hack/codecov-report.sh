#!/bin/bash

set -eo pipefail

if [[ -z "${CODECOV_TOKEN}" ]]; then
    echo "CODECOV_TOKEN must be set"
    exit 1
fi

bash <(curl -s https://codecov.io/bash) -f coverage.out
