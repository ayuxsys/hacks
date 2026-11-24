## subfind

```console
$ go run . -h
Usage of C:\Users\ayux\AppData\Local\Temp\go-build2701528074\b001\exe\subfind.exe:
  -d string
    	domain to find subdomains for
  -o string
    	path to the output file to write enumerated subdomains into
  -os string
    	path to the output file to write all enumerated subdomains into
  -r	run recursion on enumerated subdomains to find subdomains of subdomains (slow)
  -sl string
    	path to the log file to write stderr of executed commands into

$ go run . -- -h
Usage of tools:
  -rsf string
    	additional args to be passed to subfinder when running on recursion (-all ...)
  -sf string
    	additional args to be passed to subfinder (-t 5 -o ...)
```

### example:

```sh
subfind -d "staging-*.example.com" -r -o "subs.jsonl" -oa "subs_all.jsonl" -sl "stderr.log" -- -sf "-all"
```

- find subdomains of `example.com`
- use all sources for subfinder
- run recursively to find subdomains of subdomains (`-all` not used here use `-rsf "all"`)
- save output for `staging-*.example.com` to subs.jsonl
- save output for `*.example.com` to subs_all.jsonl
