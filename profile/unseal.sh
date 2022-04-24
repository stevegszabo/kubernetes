#!/bin/bash

set -o errexit

(cat <<EOF
aaaaaaaaaaaaaaaaaaaaaaaaaaa
bbbbbbbbbbbbbbbbbbbbbbbbbbb
ccccccccccccccccccccccccccc
EOF
) | while read KEY
do
vault operator unseal $KEY
done

exit 0
