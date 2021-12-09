
package main

import (
    "encoding/xml"
    "fmt"
    "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)


// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex - interface is keyword in golang
type xmlRouting__InstancesInstanceRoute__DistinguisherRd__Type struct {
	XMLName xml.Name `xml:"configuration"`
	Groups  struct {
		XMLName	xml.Name	`xml:"groups"`
		Name	string	`xml:"name"`
		V_instance  struct {
			XMLName xml.Name `xml:"instance"`
			V_name  *string  `xml:"name,omitempty"`
			V_route__distinguisher  struct {
				XMLName xml.Name `xml:"route-distinguisher"`
				V_rd__type  *string  `xml:"rd-type,omitempty"`
			} `xml:"route-distinguisher"`
		} `xml:"routing-instances>instance"`
	} `xml:"groups"`
	ApplyGroup string `xml:"apply-groups"`
}

// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex- interface is keyword in golang
func junosRouting__InstancesInstanceRoute__DistinguisherRd__TypeCreate(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_rd__type := d.Get("rd__type").(string)


	config := xmlRouting__InstancesInstanceRoute__DistinguisherRd__Type{}
	config.ApplyGroup = id
	config.Groups.Name = id
	config.Groups.V_instance.V_name = &V_name
	config.Groups.V_instance.V_route__distinguisher.V_rd__type = &V_rd__type

    err = client.SendTransaction("", config, false)
    check(err)
    
    d.SetId(fmt.Sprintf("%s_%s", client.Host, id))
    
	return junosRouting__InstancesInstanceRoute__DistinguisherRd__TypeRead(d,m)
}

func junosRouting__InstancesInstanceRoute__DistinguisherRd__TypeRead(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	config := &xmlRouting__InstancesInstanceRoute__DistinguisherRd__Type{}

	err = client.MarshalGroup(id, config)
	check(err)
 	d.Set("name", config.Groups.V_instance.V_name)
	d.Set("rd__type", config.Groups.V_instance.V_route__distinguisher.V_rd__type)

	return nil
}

func junosRouting__InstancesInstanceRoute__DistinguisherRd__TypeUpdate(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_rd__type := d.Get("rd__type").(string)


	config := xmlRouting__InstancesInstanceRoute__DistinguisherRd__Type{}
	config.ApplyGroup = id
	config.Groups.Name = id
	config.Groups.V_instance.V_name = &V_name
	config.Groups.V_instance.V_route__distinguisher.V_rd__type = &V_rd__type

    err = client.SendTransaction(id, config, false)
    check(err)
    
	return junosRouting__InstancesInstanceRoute__DistinguisherRd__TypeRead(d,m)
}

func junosRouting__InstancesInstanceRoute__DistinguisherRd__TypeDelete(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
    _, err = client.DeleteConfigNoCommit(id)
    check(err)

    d.SetId("")
    
	return nil
}

func junosRouting__InstancesInstanceRoute__DistinguisherRd__Type() *schema.Resource {
	return &schema.Resource{
		Create: junosRouting__InstancesInstanceRoute__DistinguisherRd__TypeCreate,
		Read: junosRouting__InstancesInstanceRoute__DistinguisherRd__TypeRead,
		Update: junosRouting__InstancesInstanceRoute__DistinguisherRd__TypeUpdate,
		Delete: junosRouting__InstancesInstanceRoute__DistinguisherRd__TypeDelete,

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
			"rd__type": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_instance.V_route__distinguisher. Number in (16 bit:32 bit) or (32 bit 'L':16 bit) or (IP address:16 bit) format",
			},
		},
	}
}