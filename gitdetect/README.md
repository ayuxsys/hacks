## gitdetect

```console
$ cat handles.txt
acronis
basecamp
github
google
gitlab-org

$ cat handles.txt | go run .
{"name":"gitlab-org","cloud":"github"}
{"name":"gitlab-org","cloud":"gitlab"}
{"name":"basecamp","cloud":"github"}
{"name":"google","cloud":"github"}
{"name":"github","cloud":"github"}
{"name":"acronis","cloud":"github"}
{"name":"basecamp","cloud":"gitlab"}
{"name":"google","cloud":"gitlab"}
{"name":"github","cloud":"gitlab"}
{"name":"acronis","cloud":"gitlab"}
```
