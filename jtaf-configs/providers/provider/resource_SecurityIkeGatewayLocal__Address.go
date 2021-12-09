
package main

import (
    "encoding/xml"
    "fmt"
    "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)


// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex - interface is keyword in golang
type xmlSecurityIkeGatewayLocal__Address struct {
	XMLName xml.Name `xml:"configuration"`
	Groups  struct {
		XMLName	xml.Name	`xml:"groups"`
		Name	string	`xml:"name"`
		V_gateway  struct {
			XMLName xml.Name `xml:"gateway"`
			V_name  *string  `xml:"name,omitempty"`
			V_local__address  *string  `xml:"local-address,omitempty"`
		} `xml:"security>ike>gateway"`
	} `xml:"groups"`
	ApplyGroup string `xml:"apply-groups"`
}

// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex- interface is keyword in golang
func junosSecurityIkeGatewayLocal__AddressCreate(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_local__address := d.Get("local__address").(string)


	config := xmlSecurityIkeGatewayLocal__Address{}
	config.ApplyGroup = id
	config.Groups.Name = id
	config.Groups.V_gateway.V_name = &V_name
	config.Groups.V_gateway.V_local__address = &V_local__address

    err = client.SendTransaction("", config, false)
    check(err)
    
    d.SetId(fmt.Sprintf("%s_%s", client.Host, id))
    
	return junosSecurityIkeGatewayLocal__AddressRead(d,m)
}

func junosSecurityIkeGatewayLocal__AddressRead(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	config := &xmlSecurityIkeGatewayLocal__Address{}

	err = client.MarshalGroup(id, config)
	check(err)
 	d.Set("name", config.Groups.V_gateway.V_name)
	d.Set("local__address", config.Groups.V_gateway.V_local__address)

	return nil
}

func junosSecurityIkeGatewayLocal__AddressUpdate(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_local__address := d.Get("local__address").(string)


	config := xmlSecurityIkeGatewayLocal__Address{}
	config.ApplyGroup = id
	config.Groups.Name = id
	config.Groups.V_gateway.V_name = &V_name
	config.Groups.V_gateway.V_local__address = &V_local__address

    err = client.SendTransaction(id, config, false)
    check(err)
    
	return junosSecurityIkeGatewayLocal__AddressRead(d,m)
}

func junosSecurityIkeGatewayLocal__AddressDelete(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
    _, err = client.DeleteConfigNoCommit(id)
    check(err)

    d.SetId("")
    
	return nil
}

func junosSecurityIkeGatewayLocal__Address() *schema.Resource {
	return &schema.Resource{
		Create: junosSecurityIkeGatewayLocal__AddressCreate,
		Read: junosSecurityIkeGatewayLocal__AddressRead,
		Update: junosSecurityIkeGatewayLocal__AddressUpdate,
		Delete: junosSecurityIkeGatewayLocal__AddressDelete,

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
			"local__address": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_gateway. Local IP address for IKE negotiations",
			},
		},
	}
}