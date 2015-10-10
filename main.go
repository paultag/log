package main

import (
	"os"
	"strings"
	"time"

	"pault.ag/go/config"
)

type Log struct {
	Root string `flag:"root" description:"Log root"`
}

func main() {
	when := time.Now()

	conf := Log{Root: "/tmp/log/"}
	flags, err := config.LoadFlags("log", &conf)
	if err != nil {
		panic(err)
	}
	flags.Parse(os.Args[1:])

	args := flags.Args()
	if len(args) == 0 {
		panic("No message")
	}

	if err := Create(conf.Root, when); err != nil {
		panic(err)
	}

	if err := Logit(conf.Root, when, strings.Join(args, " ")); err != nil {
		panic(err)
	}
}
