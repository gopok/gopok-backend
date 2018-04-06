package generator

import (
	"fmt"
	"io/ioutil"

	log "github.com/sirupsen/logrus"

	yaml "gopkg.in/yaml.v2"
)

func Run() {
	schemaContents, readErr := ioutil.ReadFile("schema.yaml")
	if readErr != nil {
		log.Fatal("Failed to read shcema.yaml: ", readErr)
		return
	}

	s := &schema{}
	yaml.Unmarshal(schemaContents, s)
	fmt.Printf("%+v", s)
}
