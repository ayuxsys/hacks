## subfind

```console
$ subfind -h
Usage of C:\Users\ayux\bin\projects\cybersecurity\hacks\subfind\subfind.exe:
  -d string
    	domain to find subdomains for
  -o string
    	path to the output file
  -r	run recursion on enumerated subdomains to find subdomains of subdomains (slow)
  -sl string
    	path to the log file to write stderr of executed commands into
     
$ subfind -- -h
Usage of tools:
  -rsf string
    	additional args to be passed to subfinder when running on recursion (-all ...)
  -sf string
    	additional args to be passed to subfinder (-t 5 -o ...)
```

### examples:

```sh
subfind -d "example.com" -r -o "subs.txt" -- -sf "-all"
```
