
package main

import (
    "encoding/xml"
    "fmt"
    "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)


// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex - interface is keyword in golang
type xmlRouting__InstancesInstanceVrf__TargetExport struct {
	XMLName xml.Name `xml:"configuration"`
	Groups  struct {
		XMLName	xml.Name	`xml:"groups"`
		Name	string	`xml:"name"`
		V_instance  struct {
			XMLName xml.Name `xml:"instance"`
			V_name  *string  `xml:"name,omitempty"`
			V_vrf__target  struct {
				XMLName xml.Name `xml:"vrf-target"`
				V_export  *string  `xml:"export,omitempty"`
			} `xml:"vrf-target"`
		} `xml:"routing-instances>instance"`
	} `xml:"groups"`
	ApplyGroup string `xml:"apply-groups"`
}

// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex- interface is keyword in golang
func junosRouting__InstancesInstanceVrf__TargetExportCreate(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_export := d.Get("export").(string)


	config := xmlRouting__InstancesInstanceVrf__TargetExport{}
	config.ApplyGroup = id
	config.Groups.Name = id
	config.Groups.V_instance.V_name = &V_name
	config.Groups.V_instance.V_vrf__target.V_export = &V_export

    err = client.SendTransaction("", config, false)
    check(err)
    
    d.SetId(fmt.Sprintf("%s_%s", client.Host, id))
    
	return junosRouting__InstancesInstanceVrf__TargetExportRead(d,m)
}

func junosRouting__InstancesInstanceVrf__TargetExportRead(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	config := &xmlRouting__InstancesInstanceVrf__TargetExport{}

	err = client.MarshalGroup(id, config)
	check(err)
 	d.Set("name", config.Groups.V_instance.V_name)
	d.Set("export", config.Groups.V_instance.V_vrf__target.V_export)

	return nil
}

func junosRouting__InstancesInstanceVrf__TargetExportUpdate(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_export := d.Get("export").(string)


	config := xmlRouting__InstancesInstanceVrf__TargetExport{}
	config.ApplyGroup = id
	config.Groups.Name = id
	config.Groups.V_instance.V_name = &V_name
	config.Groups.V_instance.V_vrf__target.V_export = &V_export

    err = client.SendTransaction(id, config, false)
    check(err)
    
	return junosRouting__InstancesInstanceVrf__TargetExportRead(d,m)
}

func junosRouting__InstancesInstanceVrf__TargetExportDelete(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
    _, err = client.DeleteConfigNoCommit(id)
    check(err)

    d.SetId("")
    
	return nil
}

func junosRouting__InstancesInstanceVrf__TargetExport() *schema.Resource {
	return &schema.Resource{
		Create: junosRouting__InstancesInstanceVrf__TargetExportCreate,
		Read: junosRouting__InstancesInstanceVrf__TargetExportRead,
		Update: junosRouting__InstancesInstanceVrf__TargetExportUpdate,
		Delete: junosRouting__InstancesInstanceVrf__TargetExportDelete,

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
			"export": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_instance.V_vrf__target. Target community to use when marking routes on export",
			},
		},
	}
}