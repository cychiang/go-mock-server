package main

import (
	"flag"
	"github.com/cychiang/go-mock-server/internal"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

var (
	conf       internal.Config
	configName = flag.String("config", "", "Name of the config file")
)

func init() {
	flag.Parse()
	if len(*configName) > 0 {
		yamlFile, err := os.ReadFile(*configName)
		if err != nil {
			log.Fatalln(err.Error())
			return
		}
		err = yaml.Unmarshal(yamlFile, &conf)

		if err != nil {
			log.Fatalln(err.Error())
			return
		}
	} else {
		log.Fatalln("")
	}
}

func main() {
	r := internal.SetupRouter(&conf)
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
