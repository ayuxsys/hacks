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
```

Takes a domain (or file of domains), generates `FUZZ` positions when needed, and replaces them using the selected wordlist.

```text
api.staging.example.com

-> FUZZ.example.com
-> FUZZ.staging.example.com
-> api.FUZZ.example.com
```

Hyphenated labels are mutated individually:

```text
api-staging.example.com

-> FUZZ-staging.example.com
-> api-FUZZ.example.com
```

## Usage

```bash
resub <domain|file>
```

```text
-i, --identifier string   identifier to replace (default "FUZZ")
-m, --mode string         wordlist mode: tiny, medium, large, huge
-v, --version             print version
```

## Notes

Permutations are streamed to the caller rather than accumulated in memory.

`internal/permutator` does **not** deduplicate output. Deduplication, if required, is the caller's responsibility.
