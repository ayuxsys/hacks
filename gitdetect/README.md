## gitdetect

```console
> cat .\handles.txt
google
github
gitlab
acronis
basecamp

> cat .\handles.txt | go run . -c 5
https://github.com/gitlab
https://github.com/basecamp
https://github.com/acronis
https://github.com/google
https://github.com/github
```
