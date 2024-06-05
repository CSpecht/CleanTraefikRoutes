package main

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

func truncateYaml(filename string) {

	ymlFile, err := os.ReadFile(filename)
	if err != nil {
		log.Println(err)
	}
	//  read services.yaml
	var ro Routers
	err = yaml.Unmarshal(ymlFile, &ro)
	if err != nil {
		log.Println(err)
	}
	//  check if map already empty
	if len(ro.HTTP.Routers) != 0 {
		for k := range ro.HTTP.Routers {
			delete(ro.HTTP.Routers, k)
		}

		//	write the new router to the yaml file
		newYaml, err := yaml.Marshal(&ro)
		if err != nil {
			log.Println(err)
		}
		err = os.WriteFile(filename, newYaml, 0644)
		if err != nil {
			log.Println(err)
		}
		log.Println("Router truncated")

	}
	log.Println("Router already empty")

}
