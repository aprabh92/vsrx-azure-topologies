
package main

import (
    "encoding/xml"
    "fmt"
    "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)


// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex - interface is keyword in golang
type xmlRouting__InstancesInstanceProtocolsBgpGroupAs__Override struct {
	XMLName xml.Name `xml:"configuration"`
	Groups  struct {
		XMLName	xml.Name	`xml:"groups"`
		Name	string	`xml:"name"`
		V_instance  struct {
			XMLName xml.Name `xml:"instance"`
			V_name  *string  `xml:"name,omitempty"`
			V_group  struct {
				XMLName xml.Name `xml:"group"`
				V_name__1  *string  `xml:"name,omitempty"`
				V_as__override  *string  `xml:"as-override,omitempty"`
			} `xml:"protocols>bgp>group"`
		} `xml:"routing-instances>instance"`
	} `xml:"groups"`
	ApplyGroup string `xml:"apply-groups"`
}

// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex- interface is keyword in golang
func junosRouting__InstancesInstanceProtocolsBgpGroupAs__OverrideCreate(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_name__1 := d.Get("name__1").(string)
	V_as__override := d.Get("as__override").(string)


	config := xmlRouting__InstancesInstanceProtocolsBgpGroupAs__Override{}
	config.ApplyGroup = id
	config.Groups.Name = id
	config.Groups.V_instance.V_name = &V_name
	config.Groups.V_instance.V_group.V_name__1 = &V_name__1
	config.Groups.V_instance.V_group.V_as__override = &V_as__override

    err = client.SendTransaction("", config, false)
    check(err)
    
    d.SetId(fmt.Sprintf("%s_%s", client.Host, id))
    
	return junosRouting__InstancesInstanceProtocolsBgpGroupAs__OverrideRead(d,m)
}

func junosRouting__InstancesInstanceProtocolsBgpGroupAs__OverrideRead(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	config := &xmlRouting__InstancesInstanceProtocolsBgpGroupAs__Override{}

	err = client.MarshalGroup(id, config)
	check(err)
 	d.Set("name", config.Groups.V_instance.V_name)
	d.Set("name__1", config.Groups.V_instance.V_group.V_name__1)
	d.Set("as__override", config.Groups.V_instance.V_group.V_as__override)

	return nil
}

func junosRouting__InstancesInstanceProtocolsBgpGroupAs__OverrideUpdate(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_name__1 := d.Get("name__1").(string)
	V_as__override := d.Get("as__override").(string)


	config := xmlRouting__InstancesInstanceProtocolsBgpGroupAs__Override{}
	config.ApplyGroup = id
	config.Groups.Name = id
	config.Groups.V_instance.V_name = &V_name
	config.Groups.V_instance.V_group.V_name__1 = &V_name__1
	config.Groups.V_instance.V_group.V_as__override = &V_as__override

    err = client.SendTransaction(id, config, false)
    check(err)
    
	return junosRouting__InstancesInstanceProtocolsBgpGroupAs__OverrideRead(d,m)
}

func junosRouting__InstancesInstanceProtocolsBgpGroupAs__OverrideDelete(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
    _, err = client.DeleteConfigNoCommit(id)
    check(err)

    d.SetId("")
    
	return nil
}

func junosRouting__InstancesInstanceProtocolsBgpGroupAs__Override() *schema.Resource {
	return &schema.Resource{
		Create: junosRouting__InstancesInstanceProtocolsBgpGroupAs__OverrideCreate,
		Read: junosRouting__InstancesInstanceProtocolsBgpGroupAs__OverrideRead,
		Update: junosRouting__InstancesInstanceProtocolsBgpGroupAs__OverrideUpdate,
		Delete: junosRouting__InstancesInstanceProtocolsBgpGroupAs__OverrideDelete,

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
			"name__1": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_instance.V_group",
			},
			"as__override": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_instance.V_group. Replace neighbor AS number with our AS number",
			},
		},
	}
}