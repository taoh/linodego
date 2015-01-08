package linodego

import (
	"testing"
	_ "github.com/Sirupsen/logrus"
)

func TestListJobs(t *testing.T) {
	client := NewClient(APIKey, nil)

	_, err := client.Job.List(-1, -1, false)
	if err == nil {
		t.Fatal("Should not find this job with ID -1!")
	}
}