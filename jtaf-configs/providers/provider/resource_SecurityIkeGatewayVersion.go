
package main

import (
    "encoding/xml"
    "fmt"
    "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)


// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex - interface is keyword in golang
type xmlSecurityIkeGatewayVersion struct {
	XMLName xml.Name `xml:"configuration"`
	Groups  struct {
		XMLName	xml.Name	`xml:"groups"`
		Name	string	`xml:"name"`
		V_gateway  struct {
			XMLName xml.Name `xml:"gateway"`
			V_name  *string  `xml:"name,omitempty"`
			V_version  *string  `xml:"version,omitempty"`
		} `xml:"security>ike>gateway"`
	} `xml:"groups"`
	ApplyGroup string `xml:"apply-groups"`
}

// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex- interface is keyword in golang
func junosSecurityIkeGatewayVersionCreate(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_version := d.Get("version").(string)


	config := xmlSecurityIkeGatewayVersion{}
	config.ApplyGroup = id
	config.Groups.Name = id
	config.Groups.V_gateway.V_name = &V_name
	config.Groups.V_gateway.V_version = &V_version

    err = client.SendTransaction("", config, false)
    check(err)
    
    d.SetId(fmt.Sprintf("%s_%s", client.Host, id))
    
	return junosSecurityIkeGatewayVersionRead(d,m)
}

func junosSecurityIkeGatewayVersionRead(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	config := &xmlSecurityIkeGatewayVersion{}

	err = client.MarshalGroup(id, config)
	check(err)
 	d.Set("name", config.Groups.V_gateway.V_name)
	d.Set("version", config.Groups.V_gateway.V_version)

	return nil
}

func junosSecurityIkeGatewayVersionUpdate(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_version := d.Get("version").(string)


	config := xmlSecurityIkeGatewayVersion{}
	config.ApplyGroup = id
	config.Groups.Name = id
	config.Groups.V_gateway.V_name = &V_name
	config.Groups.V_gateway.V_version = &V_version

    err = client.SendTransaction(id, config, false)
    check(err)
    
	return junosSecurityIkeGatewayVersionRead(d,m)
}

func junosSecurityIkeGatewayVersionDelete(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
    _, err = client.DeleteConfigNoCommit(id)
    check(err)

    d.SetId("")
    
	return nil
}

func junosSecurityIkeGatewayVersion() *schema.Resource {
	return &schema.Resource{
		Create: junosSecurityIkeGatewayVersionCreate,
		Read: junosSecurityIkeGatewayVersionRead,
		Update: junosSecurityIkeGatewayVersionUpdate,
		Delete: junosSecurityIkeGatewayVersionDelete,

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
			"version": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_gateway. Negotiate using either IKE v1 or IKE v2 protocol",
			},
		},
	}
}