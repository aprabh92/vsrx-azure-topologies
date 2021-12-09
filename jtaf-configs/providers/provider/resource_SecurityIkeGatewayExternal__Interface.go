
package main

import (
    "encoding/xml"
    "fmt"
    "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)


// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex - interface is keyword in golang
type xmlSecurityIkeGatewayExternal__Interface struct {
	XMLName xml.Name `xml:"configuration"`
	Groups  struct {
		XMLName	xml.Name	`xml:"groups"`
		Name	string	`xml:"name"`
		V_gateway  struct {
			XMLName xml.Name `xml:"gateway"`
			V_name  *string  `xml:"name,omitempty"`
			V_external__interface  *string  `xml:"external-interface,omitempty"`
		} `xml:"security>ike>gateway"`
	} `xml:"groups"`
	ApplyGroup string `xml:"apply-groups"`
}

// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex- interface is keyword in golang
func junosSecurityIkeGatewayExternal__InterfaceCreate(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_external__interface := d.Get("external__interface").(string)


	config := xmlSecurityIkeGatewayExternal__Interface{}
	config.ApplyGroup = id
	config.Groups.Name = id
	config.Groups.V_gateway.V_name = &V_name
	config.Groups.V_gateway.V_external__interface = &V_external__interface

    err = client.SendTransaction("", config, false)
    check(err)
    
    d.SetId(fmt.Sprintf("%s_%s", client.Host, id))
    
	return junosSecurityIkeGatewayExternal__InterfaceRead(d,m)
}

func junosSecurityIkeGatewayExternal__InterfaceRead(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	config := &xmlSecurityIkeGatewayExternal__Interface{}

	err = client.MarshalGroup(id, config)
	check(err)
 	d.Set("name", config.Groups.V_gateway.V_name)
	d.Set("external__interface", config.Groups.V_gateway.V_external__interface)

	return nil
}

func junosSecurityIkeGatewayExternal__InterfaceUpdate(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_external__interface := d.Get("external__interface").(string)


	config := xmlSecurityIkeGatewayExternal__Interface{}
	config.ApplyGroup = id
	config.Groups.Name = id
	config.Groups.V_gateway.V_name = &V_name
	config.Groups.V_gateway.V_external__interface = &V_external__interface

    err = client.SendTransaction(id, config, false)
    check(err)
    
	return junosSecurityIkeGatewayExternal__InterfaceRead(d,m)
}

func junosSecurityIkeGatewayExternal__InterfaceDelete(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
    _, err = client.DeleteConfigNoCommit(id)
    check(err)

    d.SetId("")
    
	return nil
}

func junosSecurityIkeGatewayExternal__Interface() *schema.Resource {
	return &schema.Resource{
		Create: junosSecurityIkeGatewayExternal__InterfaceCreate,
		Read: junosSecurityIkeGatewayExternal__InterfaceRead,
		Update: junosSecurityIkeGatewayExternal__InterfaceUpdate,
		Delete: junosSecurityIkeGatewayExternal__InterfaceDelete,

        Schema: map[string]*schema.Schema{
            "resource_name": &schema.Schema{
                Type:    schema.TypeString,
                Required: true,
            },
			"name": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_gateway",
			},
			"external__interface": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_gateway. External interface for IKE negotiations",
			},
		},
	}
}