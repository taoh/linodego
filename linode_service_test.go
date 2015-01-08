package linodego

import (
	"testing"
	log "github.com/Sirupsen/logrus"
)

func TestListLinodes(t *testing.T) {
	client := NewClient(APIKey, nil)

	v, err := client.Linode.List()
	if err != nil {
		t.Fatal(err)
	}

	for _, linode := range v.Linodes {
		log.Debugf("Linode: %s, %s, %d", linode.Label, linode.CreateDt, linode.PlanId)
	}
}