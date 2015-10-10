//

package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"pault.ag/go/config"
)

type Log struct {
	Root   string `flag:"root"   description:"Log root"`
	When   string `flag:"when"   description:"When is it. In YYYY-MM-DD form"`
	Action string `flag:"action" description:"action to take"`
}

func main() {
	when := time.Now()
	conf := Log{Root: "log/", Action: "write"}
	flags, err := config.LoadFlags("log", &conf)
	if err != nil {
		panic(err)
	}
	flags.Parse(os.Args[1:])

	if conf.When != "" {
		when, err = time.Parse("2006-01-02", conf.When)
		if err != nil {
			panic(err)
		}
	}

	switch conf.Action {
	case "ls":
		Listit(conf, when)
		return
	case "write":
		if err := Logit(conf.Root, when, strings.Join(flags.Args(), " ")); err != nil {
			panic(err)
		}
		return
	default:
		Help()
		return
	}

}

func Listit(conf Log, when time.Time) {
	fmt.Printf("%s\n", when.Format("*Mon, Jan 2, 2006*"))
	blocks := Readit(conf.Root, when)
	for _, block := range blocks {
		fmt.Printf("  • %s\n", block)
	}
}

func Help() {
	fmt.Println("Commands: ls, write")
}
