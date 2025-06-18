package nocloudiso

import (
	"fmt"
	"log"
	"os"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	diskfs "github.com/diskfs/go-diskfs"
	"github.com/diskfs/go-diskfs/disk"
	"github.com/diskfs/go-diskfs/filesystem"
	"github.com/diskfs/go-diskfs/filesystem/iso9660"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func resourceIsoFile() *schema.Resource {
	return &schema.Resource{
		Create: resourceIsoFileCreate,
		Read:   resourceIsoFileRead,
		Delete: resourceIsoFileDelete,

		Schema: map[string]*schema.Schema{
			"filename": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"content": {
				Type:     schema.TypeMap,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				ForceNew: true,
			},
		},
	}
}

func resourceIsoFileCreate(d *schema.ResourceData, m interface{}) error {
	filename := d.Get("filename").(string)
	content := d.Get("content").(map[string]interface{})
	var size int64 = 10*1024*1024 // 10MB

	iso, err := diskfs.Create(filename, size, 2048) // Sector size 2048
	if err != nil {
		return fmt.Errorf("failed to create ISO: %s", err)
	}
	fspec := disk.FilesystemSpec{Partition: 0, FSType: filesystem.TypeISO9660, VolumeLabel: "cidata"}

	fs, err := iso.CreateFilesystem(fspec)
	if err != nil {
		return fmt.Errorf("failed to create ISO filesystem: %s", err)
	}

	for name, raw := range content {
		data := []byte(raw.(string))

		rw, err := fs.OpenFile(name, os.O_CREATE|os.O_RDWR)
		if err != nil {
			return fmt.Errorf("failed to open file %s in ISO: %s", name, err)
		}
		rw.Write(data)
	}

	isofs, ok := fs.(*iso9660.FileSystem)
	if !ok {
		check(fmt.Errorf("not an iso9660 filesystem"))
	}
	err = isofs.Finalize(iso9660.FinalizeOptions{
		VolumeIdentifier: "cidata",
		RockRidge: true,
	})
	check(err)

	d.SetId(filename)
	return nil
}

func resourceIsoFileRead(d *schema.ResourceData, m interface{}) error {
	// Assume ISO exists, no internal verification
	return nil
}

func resourceIsoFileDelete(d *schema.ResourceData, m interface{}) error {
	filename := d.Id()
	err := os.Remove(filename)
	if err != nil {
		return fmt.Errorf("failed to delete ISO: %s", err)
	}
	d.SetId("")
	return nil
}

