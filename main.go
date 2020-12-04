package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	"github.com/pelletier/go-toml"
)

func main() {
	flag.Usage = func () {
		_, _ = fmt.Fprintln(os.Stderr, "Usage: toml CMD FILE [QUERY] [VALUE]")
		_, _ = fmt.Fprintln(os.Stderr, "")
		_, _ = fmt.Fprintln(os.Stderr, "Commands:")
		_, _ = fmt.Fprintln(os.Stderr, "lint    ", "output linted file")
		_, _ = fmt.Fprintln(os.Stderr, "get     ", "get query result")
		_, _ = fmt.Fprintln(os.Stderr, "get-keys", "get keys instead of values from query result")
		_, _ = fmt.Fprintln(os.Stderr, "set     ", "set value in toml file and save")
		_, _ = fmt.Fprintln(os.Stderr, "delete  ", "delete key in toml file and save")
	}
	flag.Parse()

	cmd := flag.Arg(0)
	filename := flag.Arg(1)
	query := flag.Arg(2)
	value := flag.Arg(3)

	if flag.NArg() < 2 {
		flag.Usage()
		os.Exit(1)
	}

	data, err := toml.LoadFile(filename)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(-1)
	}

	if cmd == "get" && flag.NArg() <= 3 {
		path := getPath(query)
		value := data.GetPath(path)
		if value != nil {
			fmt.Print(value)
		}
	} else if cmd == "delete" && flag.NArg() <= 3 {
		path := getPath(query)
		derr := data.DeletePath(path)
		if derr != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(-1)
		}
		werr := writeTomlFile(filename, data)
		if werr != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(-1)
		}
	} else if cmd == "get-keys" && flag.NArg() <= 3 {
		path := getPath(query)
		value, ok := data.GetPath(path).(*toml.Tree)
		if ok {
			for _, key := range value.Keys() {
				fmt.Println(key)
			}
		}
	} else if cmd == "lint" && flag.NArg() == 2 {
		err := writeTomlFile(filename, data)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(-1)
		}
	} else if cmd == "set" && flag.NArg() == 4 {
		path := getPath(query)
		data.SetPath(path, value)
		err := writeTomlFile(filename, data)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(-1)
		}
	} else {
		flag.Usage()
		os.Exit(1)
	}

}

func getPath(query string) []string {
	regex := regexp.MustCompile(`([^."']+)|"([^"]*)"|'([^']*)'`)
	matches := regex.FindAllStringSubmatch(query, -1)

	var path []string
	for _, e := range matches {
		path = append(path, strings.Join(e[1:], ""))
	}

	return path
}

func writeTomlFile(filename string, toml *toml.Tree) error {
	return ioutil.WriteFile(filename, []byte(toml.String()), 0644)
}
