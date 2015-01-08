// +build example
package main

import (
	"os"
	"github.com/taoh/linodego"
	log "github.com/Sirupsen/logrus"
)


func main() {
	if os.Getenv("DEBUG") == "true" {
		log.SetLevel(log.DebugLevel) // set debug level
	}

	apiKey := "[SUPPLY YOUR API KEY HERE]"
	client := linodego.NewClient(apiKey, nil)

	// Create a linode
	v, err := client.Linode.Create(2, 1, 1)
	if err != nil {
		log.Fatal(err)
	}
	linodeId := v.LinodeId.LinodeId
	log.Debugf("Created linode: %d", linodeId)
	
	
	// Shutdown the linode
	v2, err := client.Linode.Shutdown(linodeId)
	if err != nil {
		log.Fatal(err)
	}
	job := v2.Job
	log.Debugf("Shutdown linode: %s", job)

	// Delete the linode
	v, err = client.Linode.Delete(linodeId, false)
	if err != nil {
		log.Fatal(err)
	}

	log.Debugf("Deleted linode: %s", linodeId)
}
