// {{{ Copyright (c) Paul R. Tagliamonte <paultag@dc.cant.vote>, 2015
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE. }}}

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
		nWhen, err := time.Parse("2006-01-02", conf.When)
		if err != nil {
			/* last ditch */
			duration, err := time.ParseDuration(conf.When)
			if err != nil {
				panic(err)
			}
			when = when.Add(duration)
		} else {
			when = nWhen
		}
	}

	switch conf.Action {
	case "ls":
		Listit(conf, when)
		return
	case "write":
		if len(flags.Args()) == 0 {
			Listit(conf, when)
			return
		}
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

// vim: foldmethod=marker
