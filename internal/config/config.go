package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

var loaded = false
var (
	_, b, _, _ = runtime.Caller(0)
	BasePath   = filepath.Dir(b)
)

func load() {
	if loaded {
		return
	}

	err := godotenv.Load(BasePath + "/../../.env")
	if err != nil {
		log.Fatal("Error loading .env file - " + err.Error())
	}

	loaded = true
}

func Get(name string, defs ...string) string {
	load()

	res := os.Getenv(name)
	if len(res) == 0 {
		for _, v := range defs {
			res += v + " "
		}

		if len(res) > 1 {
			res = res[:len(res)-1]
		}
	}

	if len(res) == 0 {
		log.Println(fmt.Sprintf("env config %s not found and it has no default!", name))
	}

	return res
}
