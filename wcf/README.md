## Readme

```console
$ cat subs.txt
staging-uat.example.com
staging-prod.example.com
foobar.staging-uaz.example.com
foobar.example.com
testbaz.example.com
example.org

$ cat subs.txt | wcf -w "staging-*.example.com"
staging-uat.example.com
staging-prod.example.com
```
