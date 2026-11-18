#!/bin/bash

set -euo pipefail

in="$1"
op_dir="$2"

mkdir -p "$op_dir"

jq -r 'select(.type == "git") | .scope' "$in" > "$op_dir/git.txt"
jq -r 'select(.type == "url") | .scope' "$in" > "$op_dir/url.txt"
jq -r 'select(.type == "wildcard") | .scope' "$in" > "$op_dir/wc.txt"
jq -r 'select(.type == "android") | .scope' "$in" > "$op_dir/android.txt"
jq -r 'select(.type == "ios") | .scope' "$in" > "$op_dir/ios.txt"
jq -r 'select(.type == "npm") | .scope' "$in" > "$op_dir/npm.txt"
jq -r 'select(.type == "domain") | .scope' "$in" > "$op_dir/domain.txt"
jq -r 'select(.type == "ipv4") | .scope' "$in" > "$op_dir/ipv4.txt"
jq -r 'select(.type == "cidr") | .scope' "$in" > "$op_dir/cidr.txt"
jq -r 'select(.type == "pypi") | .scope' "$in" > "$op_dir/pypi.txt"
jq -r 'select(.type == "gopkg") | .scope' "$in" > "$op_dir/gopkg.txt"
jq -r 'select(.type == "other") | .scope' "$in" > "$op_dir/other.txt"
