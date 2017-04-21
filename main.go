package main

import (
	"trackstore/cli"
	"trackstore/log"
)

func main() {
	if err := cli.Execute(); err != nil {
		log.WithError(err).Error("application execution error")
	}
}
