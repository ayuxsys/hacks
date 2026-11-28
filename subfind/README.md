## subfind

TL;DR: enumerate subdomains for a domain or wildcard (normalize wildcard to domain), save results to a psql db, later spin up an API server to fetch subs from the db

```console
$ cat .env_tmp  # psql database url
DATABASE_URL=""

$ subfind -h
Usage of C:\Users\ayux\go\bin\subfind.exe:
  -a	start subfind api server
  -d string
    	domain to find subdomains for
  -e string
    	path to the .env file to load configurations from (default "C:\\Users\\ayux\\AppData\\Roaming\\subfind\\.env")
  -o string
    	fetch database for bbp and vdp subdomains then save output and exit
  -p int
    	port to use for running the api server (default 8000)
  -r	run recursion on enumerated subdomains to find subdomains of subdomains (slow)
  -sl string
    	path to the log file to write stderr of executed commands into
  -t string
    	tag the data saved  (e.g. vdp/bbp)
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

**start API**

```sh
subfind -a -p 8000
```

fetch total domains count:

```sh
$ curl "http://localhost:8000/subs/domains/count" -s | jq .count
241
```

fetch an id:

```sh
$ curl "http://localhost:8000/subs/fetch?id=1" -s | jq .

{
  "domain": "*.hackerone-user-content.com",
  "subs": [],
  "tag": "bbp"
}
```

securing the server with an api key:

```sh
$ openssl rand -base64 32
dj9Ea56wm2ZZzYsM2G33A7vL2PsmpktEaSgS4ousw2c=

$ API_KEY="dj9Ea56wm2ZZzYsM2G33A7vL2PsmpktEaSgS4ousw2c" subfind -a
[GIN-debug] Listening and serving HTTP on :8000

$ curl "http://localhost:8000/subs/domains/count" -H "X-API-KEY: dj9Ea56wm2ZZzYsM2G33A7vL2PsmpktEaSgS4ousw2c"
{"count":245}

$ curl "http://localhost:8000/subs/domains/count"
{"error":"invalid API Key"}
```
