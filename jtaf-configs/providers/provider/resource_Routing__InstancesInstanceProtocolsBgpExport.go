
package main

import (
    "encoding/xml"
    "fmt"
    "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)


// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex - interface is keyword in golang
type xmlRouting__InstancesInstanceProtocolsBgpExport struct {
	XMLName xml.Name `xml:"configuration"`
	Groups  struct {
		XMLName	xml.Name	`xml:"groups"`
		Name	string	`xml:"name"`
		V_instance  struct {
			XMLName xml.Name `xml:"instance"`
			V_name  *string  `xml:"name,omitempty"`
			V_bgp  struct {
				XMLName xml.Name `xml:"bgp"`
				V_export  *string  `xml:"export,omitempty"`
			} `xml:"protocols>bgp"`
		} `xml:"routing-instances>instance"`
	} `xml:"groups"`
	ApplyGroup string `xml:"apply-groups"`
}

// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex- interface is keyword in golang
func junosRouting__InstancesInstanceProtocolsBgpExportCreate(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_export := d.Get("export").(string)


	config := xmlRouting__InstancesInstanceProtocolsBgpExport{}
	config.ApplyGroup = id
	config.Groups.Name = id
	config.Groups.V_instance.V_name = &V_name
	config.Groups.V_instance.V_bgp.V_export = &V_export

    err = client.SendTransaction("", config, false)
    check(err)
    
    d.SetId(fmt.Sprintf("%s_%s", client.Host, id))
    
	return junosRouting__InstancesInstanceProtocolsBgpExportRead(d,m)
}

func junosRouting__InstancesInstanceProtocolsBgpExportRead(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	config := &xmlRouting__InstancesInstanceProtocolsBgpExport{}

	err = client.MarshalGroup(id, config)
	check(err)
 	d.Set("name", config.Groups.V_instance.V_name)
	d.Set("export", config.Groups.V_instance.V_bgp.V_export)

	return nil
}

func junosRouting__InstancesInstanceProtocolsBgpExportUpdate(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_export := d.Get("export").(string)


	config := xmlRouting__InstancesInstanceProtocolsBgpExport{}
	config.ApplyGroup = id
	config.Groups.Name = id
	config.Groups.V_instance.V_name = &V_name
	config.Groups.V_instance.V_bgp.V_export = &V_export

    err = client.SendTransaction(id, config, false)
    check(err)
    
	return junosRouting__InstancesInstanceProtocolsBgpExportRead(d,m)
}

func junosRouting__InstancesInstanceProtocolsBgpExportDelete(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
    _, err = client.DeleteConfigNoCommit(id)
    check(err)

    d.SetId("")
    
	return nil
}

func junosRouting__InstancesInstanceProtocolsBgpExport() *schema.Resource {
	return &schema.Resource{
		Create: junosRouting__InstancesInstanceProtocolsBgpExportCreate,
		Read: junosRouting__InstancesInstanceProtocolsBgpExportRead,
		Update: junosRouting__InstancesInstanceProtocolsBgpExportUpdate,
		Delete: junosRouting__InstancesInstanceProtocolsBgpExportDelete,

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
				Description:    "xpath is: config.Groups.V_instance.V_bgp. Export policy",
			},
		},
	}
}