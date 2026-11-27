## subfind

```console
$ cat .env_tmp  # psql database url
DATABASE_URL=""

$ subfind -h
Usage of C:\Users\ayux\go\bin\subfind.exe:
  -d string
    	domain to find subdomains for
  -e string
    	path to the .env file to load configurations from (default "C:\\Users\\ayux\\AppData\\Roaming\\subfind")
  -r	run recursion on enumerated subdomains to find subdomains of subdomains (slow)
  -sl string
    	path to the log file to write stderr of executed commands into
  -u string
    	psql database url to use for connection (reads DB_URL from .env otherwise)
```

### example:

```sh
subfind -d "staging-*.example.com" -r -sl "stderr.log" -- -sf "-all"
```

- find subdomains of `example.com`
- use all sources for subfinder
- run recursively to find subdomains of subdomains (`-all` not used here use `-rsf "all"`)
- log output from executed commands to `stderr.log`

also supports reading from stdin: `cat domains.txt | subfind -r -- -sf "all"`
