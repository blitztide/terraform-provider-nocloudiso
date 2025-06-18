# iso_file  
  
This is to create an ISO9660 file with an arbitrary `filename` and `content`  
  
## Example Usage  
```terraform  
resource "iso_file" "test" {  
	filename = "test.iso"  
	content = {  
		"meta-data" = <<-EOF  
		instance-id: ae6beaa196415dc63163bdc3affd31b1"  
		EOF  
		"network-config" = <<-EOF  
		version: 1  
		config:  
		- type: physical  
		name: eth0  
		mac_address: 'bc:24:11:6e:43:b7'  
		subnets:  
		- type: dhcp4  
		EOF  
		"user-data" = <<-EOF  
		#cloud-config  
		chpasswd:  
		list: |  
		ubuntu:example  
		expire: false  
		hostname: cloudinit  
		packages:  
		- qemu-guest-agent  
		users:  
		- default  
		- name: ubuntu  
		groups: sudo  
		shell: /bin/bash  
		sudo: ALL=(ALL) NOPASSWD:ALL  
		runcmd:  
		- [ touch, "/var/log/farts.log" ]  
		  
		EOF  
		"vendor-data" = " "  
	}
}
```  
  
## Schema  
  
## Required  
-  `filename` (String) Filename to save ISO as  
-  `content` files within the ISO:  
	-  `filename` =  `content`  key value pairs, where `filename` is the filename within the ISO and `content` is the content of the file.
