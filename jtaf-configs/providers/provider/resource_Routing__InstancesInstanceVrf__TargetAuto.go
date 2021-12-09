
package main

import (
    "encoding/xml"
    "fmt"
    "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)


// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex - interface is keyword in golang
type xmlRouting__InstancesInstanceVrf__TargetAuto struct {
	XMLName xml.Name `xml:"configuration"`
	Groups  struct {
		XMLName	xml.Name	`xml:"groups"`
		Name	string	`xml:"name"`
		V_instance  struct {
			XMLName xml.Name `xml:"instance"`
			V_name  *string  `xml:"name,omitempty"`
			V_vrf__target  struct {
				XMLName xml.Name `xml:"vrf-target"`
				V_auto  *string  `xml:"auto,omitempty"`
			} `xml:"vrf-target"`
		} `xml:"routing-instances>instance"`
	} `xml:"groups"`
	ApplyGroup string `xml:"apply-groups"`
}

// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex- interface is keyword in golang
func junosRouting__InstancesInstanceVrf__TargetAutoCreate(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_auto := d.Get("auto").(string)


	config := xmlRouting__InstancesInstanceVrf__TargetAuto{}
	config.ApplyGroup = id
	config.Groups.Name = id
	config.Groups.V_instance.V_name = &V_name
	config.Groups.V_instance.V_vrf__target.V_auto = &V_auto

    err = client.SendTransaction("", config, false)
    check(err)
    
    d.SetId(fmt.Sprintf("%s_%s", client.Host, id))
    
	return junosRouting__InstancesInstanceVrf__TargetAutoRead(d,m)
}

func junosRouting__InstancesInstanceVrf__TargetAutoRead(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	config := &xmlRouting__InstancesInstanceVrf__TargetAuto{}

	err = client.MarshalGroup(id, config)
	check(err)
 	d.Set("name", config.Groups.V_instance.V_name)
	d.Set("auto", config.Groups.V_instance.V_vrf__target.V_auto)

	return nil
}

func junosRouting__InstancesInstanceVrf__TargetAutoUpdate(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_auto := d.Get("auto").(string)


	config := xmlRouting__InstancesInstanceVrf__TargetAuto{}
	config.ApplyGroup = id
	config.Groups.Name = id
	config.Groups.V_instance.V_name = &V_name
	config.Groups.V_instance.V_vrf__target.V_auto = &V_auto

    err = client.SendTransaction(id, config, false)
    check(err)
    
	return junosRouting__InstancesInstanceVrf__TargetAutoRead(d,m)
}

func junosRouting__InstancesInstanceVrf__TargetAutoDelete(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
    _, err = client.DeleteConfigNoCommit(id)
    check(err)

    d.SetId("")
    
	return nil
}

func junosRouting__InstancesInstanceVrf__TargetAuto() *schema.Resource {
	return &schema.Resource{
		Create: junosRouting__InstancesInstanceVrf__TargetAutoCreate,
		Read: junosRouting__InstancesInstanceVrf__TargetAutoRead,
		Update: junosRouting__InstancesInstanceVrf__TargetAutoUpdate,
		Delete: junosRouting__InstancesInstanceVrf__TargetAutoDelete,

        Schema: map[string]*schema.Schema{
            "resource_name": &schema.Schema{
                Type:    schema.TypeString,
                Required: true,
            },
			"name": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_instance",
			},
			"auto": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_instance.V_vrf__target. Auto derive import and export target community from BGP AS &amp; L2",
			},
		},
	}
}