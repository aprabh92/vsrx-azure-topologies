
package main

import (
    "encoding/xml"
    "fmt"
    "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)


// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex - interface is keyword in golang
type xmlRouting__InstancesInstanceProtocolsBgpFamilyInetLabeled__Unicast struct {
	XMLName xml.Name `xml:"configuration"`
	Groups  struct {
		XMLName	xml.Name	`xml:"groups"`
		Name	string	`xml:"name"`
		V_instance  struct {
			XMLName xml.Name `xml:"instance"`
			V_name  *string  `xml:"name,omitempty"`
			V_labeled__unicast  struct {
				XMLName xml.Name `xml:"labeled-unicast"`
			} `xml:"protocols>bgp>family>inet>labeled-unicast"`
		} `xml:"routing-instances>instance"`
	} `xml:"groups"`
	ApplyGroup string `xml:"apply-groups"`
}

// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex- interface is keyword in golang
func junosRouting__InstancesInstanceProtocolsBgpFamilyInetLabeled__UnicastCreate(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)


	config := xmlRouting__InstancesInstanceProtocolsBgpFamilyInetLabeled__Unicast{}
	config.ApplyGroup = id
	config.Groups.Name = id
	config.Groups.V_instance.V_name = &V_name

    err = client.SendTransaction("", config, false)
    check(err)
    
    d.SetId(fmt.Sprintf("%s_%s", client.Host, id))
    
	return junosRouting__InstancesInstanceProtocolsBgpFamilyInetLabeled__UnicastRead(d,m)
}

func junosRouting__InstancesInstanceProtocolsBgpFamilyInetLabeled__UnicastRead(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	config := &xmlRouting__InstancesInstanceProtocolsBgpFamilyInetLabeled__Unicast{}

	err = client.MarshalGroup(id, config)
	check(err)
 	d.Set("name", config.Groups.V_instance.V_name)

	return nil
}

func junosRouting__InstancesInstanceProtocolsBgpFamilyInetLabeled__UnicastUpdate(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)


	config := xmlRouting__InstancesInstanceProtocolsBgpFamilyInetLabeled__Unicast{}
	config.ApplyGroup = id
	config.Groups.Name = id
	config.Groups.V_instance.V_name = &V_name

    err = client.SendTransaction(id, config, false)
    check(err)
    
	return junosRouting__InstancesInstanceProtocolsBgpFamilyInetLabeled__UnicastRead(d,m)
}

func junosRouting__InstancesInstanceProtocolsBgpFamilyInetLabeled__UnicastDelete(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
    _, err = client.DeleteConfigNoCommit(id)
    check(err)

    d.SetId("")
    
	return nil
}

func junosRouting__InstancesInstanceProtocolsBgpFamilyInetLabeled__Unicast() *schema.Resource {
	return &schema.Resource{
		Create: junosRouting__InstancesInstanceProtocolsBgpFamilyInetLabeled__UnicastCreate,
		Read: junosRouting__InstancesInstanceProtocolsBgpFamilyInetLabeled__UnicastRead,
		Update: junosRouting__InstancesInstanceProtocolsBgpFamilyInetLabeled__UnicastUpdate,
		Delete: junosRouting__InstancesInstanceProtocolsBgpFamilyInetLabeled__UnicastDelete,

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
		},
	}
}