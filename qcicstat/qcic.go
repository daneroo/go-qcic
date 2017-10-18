package main

import (
	"fmt"
	"log"
	"os"

	. "github.com/bitly/go-simplejson"
	"github.com/segmentio/go-loggly-search"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func env(name string) string {
	val := os.Getenv(name)

	if val == "" {
		log.Fatalf("env variable %s required for example", name)
	}

	return val
}

func main() {
	c := search.New(env("ACCOUNT"), env("USER"), env("PASS"))

	res, err := c.Query(`logcheck`).Size(50).From("-5h").Fetch()
	check(err)

	for _, event := range res.Events {
		Output(event)
	}

	fmt.Println()
}

func Output(event interface{}) {
	msg := event.(map[string]interface{})["logmsg"].(string)
	obj, err := NewJson([]byte(msg))
	check(err)

	fmt.Println()
	for k, v := range obj.MustMap() {
		fmt.Printf("  %14s: %s\n", k, v)
	}
}
