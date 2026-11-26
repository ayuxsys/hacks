package config

import (
	"os"

	"github.com/ayuxsys/notify/log"
	"github.com/ayuxsys/notify/must"
	"github.com/goccy/go-yaml"
)

func Setup() {
	if _, err := os.Stat(DefaultPath); !os.IsNotExist(err) {
		return // config file already exists
	}
	must.Panic(os.MkdirAll(BaseDir, 0700))
	data := must.Eval(yaml.Marshal(Config{}))
	must.Panic(os.WriteFile(DefaultPath, data, 0644))
	log.Info.Printf("created config file %q. customise and rerun the app", DefaultPath)
	os.Exit(0)
}

func Load(path string) Config {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Error.Fatalf("read config file %q: %v", path, err)
	}
	var cfg Config
	must.Panic(yaml.Unmarshal(data, &cfg))
	return cfg
}
