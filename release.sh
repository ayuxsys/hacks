#!/bin/sh

go_build() {
    path="$1"
    GOOS=windows GOARCH=amd64 go -C $path build -o "../dist"
    GOOS=linux GOARCH=amd64 go -C $path build -o "../dist"
}

go_build "subfind/"
go_build "concurl/"
go_build "gitdetect/"
go_build "hn/"
go_build "notify/"
go_build "resub/"
go_build "sn/"
go_build "wcf/"

tag="$1"
gh release create $tag "dist/*" --title "$tag" --generate-notes
