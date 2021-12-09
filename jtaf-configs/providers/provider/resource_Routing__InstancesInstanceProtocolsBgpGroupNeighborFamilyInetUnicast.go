
package main

import (
    "encoding/xml"
    "fmt"
    "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)


// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex - interface is keyword in golang
type xmlRouting__InstancesInstanceProtocolsBgpGroupNeighborFamilyInetUnicast struct {
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
					V_unicast  struct {
						XMLName xml.Name `xml:"unicast"`
					} `xml:"family>inet>unicast"`
				} `xml:"neighbor"`
			} `xml:"protocols>bgp>group"`
		} `xml:"routing-instances>instance"`
	} `xml:"groups"`
	ApplyGroup string `xml:"apply-groups"`
}

// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex- interface is keyword in golang
func junosRouting__InstancesInstanceProtocolsBgpGroupNeighborFamilyInetUnicastCreate(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_name__1 := d.Get("name__1").(string)
	V_name__2 := d.Get("name__2").(string)


	config := xmlRouting__InstancesInstanceProtocolsBgpGroupNeighborFamilyInetUnicast{}
	config.ApplyGroup = id
	config.Groups.Name = id
	config.Groups.V_instance.V_name = &V_name
	config.Groups.V_instance.V_group.V_name__1 = &V_name__1
	config.Groups.V_instance.V_group.V_neighbor.V_name__2 = &V_name__2

    err = client.SendTransaction("", config, false)
    check(err)
    
    d.SetId(fmt.Sprintf("%s_%s", client.Host, id))
    
	return junosRouting__InstancesInstanceProtocolsBgpGroupNeighborFamilyInetUnicastRead(d,m)
}

func junosRouting__InstancesInstanceProtocolsBgpGroupNeighborFamilyInetUnicastRead(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	config := &xmlRouting__InstancesInstanceProtocolsBgpGroupNeighborFamilyInetUnicast{}

	err = client.MarshalGroup(id, config)
	check(err)
 	d.Set("name", config.Groups.V_instance.V_name)
	d.Set("name__1", config.Groups.V_instance.V_group.V_name__1)
	d.Set("name__2", config.Groups.V_instance.V_group.V_neighbor.V_name__2)

	return nil
}

func junosRouting__InstancesInstanceProtocolsBgpGroupNeighborFamilyInetUnicastUpdate(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_name__1 := d.Get("name__1").(string)
	V_name__2 := d.Get("name__2").(string)


	config := xmlRouting__InstancesInstanceProtocolsBgpGroupNeighborFamilyInetUnicast{}
	config.ApplyGroup = id
	config.Groups.Name = id
	config.Groups.V_instance.V_name = &V_name
	config.Groups.V_instance.V_group.V_name__1 = &V_name__1
	config.Groups.V_instance.V_group.V_neighbor.V_name__2 = &V_name__2

    err = client.SendTransaction(id, config, false)
    check(err)
    
	return junosRouting__InstancesInstanceProtocolsBgpGroupNeighborFamilyInetUnicastRead(d,m)
}

func junosRouting__InstancesInstanceProtocolsBgpGroupNeighborFamilyInetUnicastDelete(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
    _, err = client.DeleteConfigNoCommit(id)
    check(err)

    d.SetId("")
    
	return nil
}

func junosRouting__InstancesInstanceProtocolsBgpGroupNeighborFamilyInetUnicast() *schema.Resource {
	return &schema.Resource{
		Create: junosRouting__InstancesInstanceProtocolsBgpGroupNeighborFamilyInetUnicastCreate,
		Read: junosRouting__InstancesInstanceProtocolsBgpGroupNeighborFamilyInetUnicastRead,
		Update: junosRouting__InstancesInstanceProtocolsBgpGroupNeighborFamilyInetUnicastUpdate,
		Delete: junosRouting__InstancesInstanceProtocolsBgpGroupNeighborFamilyInetUnicastDelete,

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
		},
	}
}