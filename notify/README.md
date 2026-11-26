# notify

```console
$ notify -h
Usage: notify [flags]
  -c, --config string   path to the config file to use (default "C:\\Users\\ayux\\AppData\\Roaming\\notify/config.yaml")
  -f, --file string     path to the file to send as msg
  -m, --msg string      short msg to send (len < 4096)
```

config:

```yaml
notify:
  telegram:
    enabled: true
    bot_token: ""
    chat_id: 0
```
