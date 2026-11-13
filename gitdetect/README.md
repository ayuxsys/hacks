## gitdetect

```console
$ cat handles.txt
acronis
basecamp
github
google

$ cat handles.txt | go run .
{"name":"acronis","cloud":"github"}
{"name":"basecamp","cloud":"github"}
{"name":"github","cloud":"github"}
{"name":"google","cloud":"github"}
```
