
package main

import (
    "encoding/xml"
    "fmt"
    "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)


// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex - interface is keyword in golang
type xmlSecurityIkeGatewayAddress struct {
	XMLName xml.Name `xml:"configuration"`
	Groups  struct {
		XMLName	xml.Name	`xml:"groups"`
		Name	string	`xml:"name"`
		V_gateway  struct {
			XMLName xml.Name `xml:"gateway"`
			V_name  *string  `xml:"name,omitempty"`
			V_address  *string  `xml:"address,omitempty"`
		} `xml:"security>ike>gateway"`
	} `xml:"groups"`
	ApplyGroup string `xml:"apply-groups"`
}

// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex- interface is keyword in golang
func junosSecurityIkeGatewayAddressCreate(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_address := d.Get("address").(string)


	config := xmlSecurityIkeGatewayAddress{}
	config.ApplyGroup = id
	config.Groups.Name = id
	config.Groups.V_gateway.V_name = &V_name
	config.Groups.V_gateway.V_address = &V_address

    err = client.SendTransaction("", config, false)
    check(err)
    
    d.SetId(fmt.Sprintf("%s_%s", client.Host, id))
    
	return junosSecurityIkeGatewayAddressRead(d,m)
}

func junosSecurityIkeGatewayAddressRead(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	config := &xmlSecurityIkeGatewayAddress{}

	err = client.MarshalGroup(id, config)
	check(err)
 	d.Set("name", config.Groups.V_gateway.V_name)
	d.Set("address", config.Groups.V_gateway.V_address)

	return nil
}

func junosSecurityIkeGatewayAddressUpdate(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_address := d.Get("address").(string)


	config := xmlSecurityIkeGatewayAddress{}
	config.ApplyGroup = id
	config.Groups.Name = id
	config.Groups.V_gateway.V_name = &V_name
	config.Groups.V_gateway.V_address = &V_address

    err = client.SendTransaction(id, config, false)
    check(err)
    
	return junosSecurityIkeGatewayAddressRead(d,m)
}

func junosSecurityIkeGatewayAddressDelete(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
    _, err = client.DeleteConfigNoCommit(id)
    check(err)

    d.SetId("")
    
	return nil
}

func junosSecurityIkeGatewayAddress() *schema.Resource {
	return &schema.Resource{
		Create: junosSecurityIkeGatewayAddressCreate,
		Read: junosSecurityIkeGatewayAddressRead,
		Update: junosSecurityIkeGatewayAddressUpdate,
		Delete: junosSecurityIkeGatewayAddressDelete,

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
			"address": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_gateway. Addresses or hostnames of peer:1 primary, upto 4 backups",
			},
		},
	}
}