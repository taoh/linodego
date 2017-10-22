package linodego

import (
	log "github.com/sirupsen/logrus"
	"testing"
)

func TestListImages(t *testing.T) {
	client := NewClient(APIKey, nil)

	v, err := client.Image.List()
	if err != nil {
		t.Fatal(err)
	}

	for _, image := range v.Images {
		log.Debugf("Kernel: %s, %s, %s", image.Label, image.Status, image.CreateDt)
	}
}

func TestUpdateImage(t *testing.T) {
	client := NewClient(APIKey, nil)
	createResp, err := client.Linode.Create(2, 1, 1)
	if err != nil {
		t.Fatal(err)
	}
	// id of the freshly created linode
	linodeId := createResp.LinodeId.LinodeId

	// create simple rootfs
	_, err = client.Disk.Create(linodeId, "ext4", "rootfs", 19286, nil)
	if err != nil {
		t.Fatal(err)
	}

	// list the disks now associated with this linode
	diskListResp, err := client.Disk.List(linodeId, 0)
	if err != nil {
		t.Fatal(err)
	}

	// makes the assumption that only one disk was created
	diskId := diskListResp.Disks[0].DiskId

	// imagize the disk that was just created, unfortunately just responds with job id
	_, err = client.Disk.Imagize(linodeId, diskId, "Part of testing", "TestingImage")
	if err != nil {
		t.Fatal(err)
	}

	// this takes a while
	WaitForPendingJobs(client, linodeId)

	// list images for this linode account
	imageListResp, err := client.Image.List()
	if err != nil {
		t.Fatal(err)
	}

	// retrieve image
	var imageid int
	for _, image := range imageListResp.Images {
		if image.Description == "Part of testing" {
			imageid = image.ImageId
			break
		}
	}

	imageUpdateResponse, err := client.Image.Update(imageid, "Updated label", "Updated description")
	if err != nil {
		t.Fatal(err)
	}

	image := imageUpdateResponse.Image
	if image.Description != "Updated description" {
		t.Fatal(image.Description)
	}
	log.Debugf("Kernel: %s, %s, %s", image.Label, image.Status, image.CreateDt)

	// teardown because money
	_, err = client.Image.Delete(imageid)
	if err != nil {
		t.Fatal(err)
	}

	_, err = client.Linode.Delete(linodeId, true)
	if err != nil {
		t.Fatal(err)
	}
}
