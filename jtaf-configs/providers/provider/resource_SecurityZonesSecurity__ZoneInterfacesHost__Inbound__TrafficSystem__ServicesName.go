
package main

import (
    "encoding/xml"
    "fmt"
    "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)


// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex - interface is keyword in golang
type xmlSecurityZonesSecurity__ZoneInterfacesHost__Inbound__TrafficSystem__ServicesName struct {
	XMLName xml.Name `xml:"configuration"`
	Groups  struct {
		XMLName	xml.Name	`xml:"groups"`
		Name	string	`xml:"name"`
		V_security__zone  struct {
			XMLName xml.Name `xml:"security-zone"`
			V_name  *string  `xml:"name,omitempty"`
			V_interfaces  struct {
				XMLName xml.Name `xml:"interfaces"`
				V_name__1  *string  `xml:"name,omitempty"`
				V_system__services  struct {
					XMLName xml.Name `xml:"system-services"`
					V_name__2  *string  `xml:"name,omitempty"`
				} `xml:"host-inbound-traffic>system-services"`
			} `xml:"interfaces"`
		} `xml:"security>zones>security-zone"`
	} `xml:"groups"`
	ApplyGroup string `xml:"apply-groups"`
}

// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex- interface is keyword in golang
func junosSecurityZonesSecurity__ZoneInterfacesHost__Inbound__TrafficSystem__ServicesNameCreate(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_name__1 := d.Get("name__1").(string)
	V_name__2 := d.Get("name__2").(string)


	config := xmlSecurityZonesSecurity__ZoneInterfacesHost__Inbound__TrafficSystem__ServicesName{}
	config.ApplyGroup = id
	config.Groups.Name = id
	config.Groups.V_security__zone.V_name = &V_name
	config.Groups.V_security__zone.V_interfaces.V_name__1 = &V_name__1
	config.Groups.V_security__zone.V_interfaces.V_system__services.V_name__2 = &V_name__2

    err = client.SendTransaction("", config, false)
    check(err)
    
    d.SetId(fmt.Sprintf("%s_%s", client.Host, id))
    
	return junosSecurityZonesSecurity__ZoneInterfacesHost__Inbound__TrafficSystem__ServicesNameRead(d,m)
}

func junosSecurityZonesSecurity__ZoneInterfacesHost__Inbound__TrafficSystem__ServicesNameRead(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	config := &xmlSecurityZonesSecurity__ZoneInterfacesHost__Inbound__TrafficSystem__ServicesName{}

	err = client.MarshalGroup(id, config)
	check(err)
 	d.Set("name", config.Groups.V_security__zone.V_name)
	d.Set("name__1", config.Groups.V_security__zone.V_interfaces.V_name__1)
	d.Set("name__2", config.Groups.V_security__zone.V_interfaces.V_system__services.V_name__2)

	return nil
}

func junosSecurityZonesSecurity__ZoneInterfacesHost__Inbound__TrafficSystem__ServicesNameUpdate(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_name__1 := d.Get("name__1").(string)
	V_name__2 := d.Get("name__2").(string)


	config := xmlSecurityZonesSecurity__ZoneInterfacesHost__Inbound__TrafficSystem__ServicesName{}
	config.ApplyGroup = id
	config.Groups.Name = id
	config.Groups.V_security__zone.V_name = &V_name
	config.Groups.V_security__zone.V_interfaces.V_name__1 = &V_name__1
	config.Groups.V_security__zone.V_interfaces.V_system__services.V_name__2 = &V_name__2

    err = client.SendTransaction(id, config, false)
    check(err)
    
	return junosSecurityZonesSecurity__ZoneInterfacesHost__Inbound__TrafficSystem__ServicesNameRead(d,m)
}

func junosSecurityZonesSecurity__ZoneInterfacesHost__Inbound__TrafficSystem__ServicesNameDelete(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
    _, err = client.DeleteConfigNoCommit(id)
    check(err)

    d.SetId("")
    
	return nil
}

func junosSecurityZonesSecurity__ZoneInterfacesHost__Inbound__TrafficSystem__ServicesName() *schema.Resource {
	return &schema.Resource{
		Create: junosSecurityZonesSecurity__ZoneInterfacesHost__Inbound__TrafficSystem__ServicesNameCreate,
		Read: junosSecurityZonesSecurity__ZoneInterfacesHost__Inbound__TrafficSystem__ServicesNameRead,
		Update: junosSecurityZonesSecurity__ZoneInterfacesHost__Inbound__TrafficSystem__ServicesNameUpdate,
		Delete: junosSecurityZonesSecurity__ZoneInterfacesHost__Inbound__TrafficSystem__ServicesNameDelete,

        Schema: map[string]*schema.Schema{
            "resource_name": &schema.Schema{
                Type:    schema.TypeString,
                Required: true,
            },
			"name": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_security__zone",
			},
			"name__1": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_security__zone.V_interfaces",
			},
			"name__2": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_security__zone.V_interfaces.V_system__services. ",
			},
		},
	}
}