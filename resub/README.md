# resub

Subdomain permutation tool for recon.

```console
$ cat domains.txt
staging-prod.example.com
example.com
prod.test-uat.example.com

$ cat wordlist.txt
FUZZ

$ resub domains.txt -w wordlist.txt
FUZZ.example.com
FUZZ-prod.example.com
staging-FUZZ.example.com
FUZZ.example.com
FUZZ.example.com
FUZZ.test-uat.example.com
prod.FUZZ-uat.example.com
prod.test-FUZZ.example.com

$ resub domains.txt -w wordlist.txt | sort -u  # dedupe
FUZZ.example.com
FUZZ.test-uat.example.com
FUZZ-prod.example.com
prod.FUZZ-uat.example.com
prod.test-FUZZ.example.com
staging-FUZZ.example.com
```

# example workflow

```sh
resub subs.txt -m tiny | puredns resolve -w fuzzed.txt
```
