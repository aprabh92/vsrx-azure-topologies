
package main

import (
    "encoding/xml"
    "fmt"
    "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)


// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex - interface is keyword in golang
type xmlRouting__InstancesInstanceProtocolsBgpGroupNeighborLocal__Address struct {
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
				V_neighbor  struct {
					XMLName xml.Name `xml:"neighbor"`
					V_name__2  *string  `xml:"name,omitempty"`
					V_local__address  *string  `xml:"local-address,omitempty"`
				} `xml:"neighbor"`
			} `xml:"protocols>bgp>group"`
		} `xml:"routing-instances>instance"`
	} `xml:"groups"`
	ApplyGroup string `xml:"apply-groups"`
}

// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex- interface is keyword in golang
func junosRouting__InstancesInstanceProtocolsBgpGroupNeighborLocal__AddressCreate(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_name__1 := d.Get("name__1").(string)
	V_name__2 := d.Get("name__2").(string)
	V_local__address := d.Get("local__address").(string)


	config := xmlRouting__InstancesInstanceProtocolsBgpGroupNeighborLocal__Address{}
	config.ApplyGroup = id
	config.Groups.Name = id
	config.Groups.V_instance.V_name = &V_name
	config.Groups.V_instance.V_group.V_name__1 = &V_name__1
	config.Groups.V_instance.V_group.V_neighbor.V_name__2 = &V_name__2
	config.Groups.V_instance.V_group.V_neighbor.V_local__address = &V_local__address

    err = client.SendTransaction("", config, false)
    check(err)
    
    d.SetId(fmt.Sprintf("%s_%s", client.Host, id))
    
	return junosRouting__InstancesInstanceProtocolsBgpGroupNeighborLocal__AddressRead(d,m)
}

func junosRouting__InstancesInstanceProtocolsBgpGroupNeighborLocal__AddressRead(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	config := &xmlRouting__InstancesInstanceProtocolsBgpGroupNeighborLocal__Address{}

	err = client.MarshalGroup(id, config)
	check(err)
 	d.Set("name", config.Groups.V_instance.V_name)
	d.Set("name__1", config.Groups.V_instance.V_group.V_name__1)
	d.Set("name__2", config.Groups.V_instance.V_group.V_neighbor.V_name__2)
	d.Set("local__address", config.Groups.V_instance.V_group.V_neighbor.V_local__address)

	return nil
}

func junosRouting__InstancesInstanceProtocolsBgpGroupNeighborLocal__AddressUpdate(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_name__1 := d.Get("name__1").(string)
	V_name__2 := d.Get("name__2").(string)
	V_local__address := d.Get("local__address").(string)


	config := xmlRouting__InstancesInstanceProtocolsBgpGroupNeighborLocal__Address{}
	config.ApplyGroup = id
	config.Groups.Name = id
	config.Groups.V_instance.V_name = &V_name
	config.Groups.V_instance.V_group.V_name__1 = &V_name__1
	config.Groups.V_instance.V_group.V_neighbor.V_name__2 = &V_name__2
	config.Groups.V_instance.V_group.V_neighbor.V_local__address = &V_local__address

    err = client.SendTransaction(id, config, false)
    check(err)
    
	return junosRouting__InstancesInstanceProtocolsBgpGroupNeighborLocal__AddressRead(d,m)
}

func junosRouting__InstancesInstanceProtocolsBgpGroupNeighborLocal__AddressDelete(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
    _, err = client.DeleteConfigNoCommit(id)
    check(err)

    d.SetId("")
    
	return nil
}

func junosRouting__InstancesInstanceProtocolsBgpGroupNeighborLocal__Address() *schema.Resource {
	return &schema.Resource{
		Create: junosRouting__InstancesInstanceProtocolsBgpGroupNeighborLocal__AddressCreate,
		Read: junosRouting__InstancesInstanceProtocolsBgpGroupNeighborLocal__AddressRead,
		Update: junosRouting__InstancesInstanceProtocolsBgpGroupNeighborLocal__AddressUpdate,
		Delete: junosRouting__InstancesInstanceProtocolsBgpGroupNeighborLocal__AddressDelete,

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
			"name__2": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_instance.V_group.V_neighbor",
			},
			"local__address": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_instance.V_group.V_neighbor. Address of local end of BGP session",
			},
		},
	}
}