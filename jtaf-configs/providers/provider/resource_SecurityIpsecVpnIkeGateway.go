
package main

import (
    "encoding/xml"
    "fmt"
    "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)


// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex - interface is keyword in golang
type xmlSecurityIpsecVpnIkeGateway struct {
	XMLName xml.Name `xml:"configuration"`
	Groups  struct {
		XMLName	xml.Name	`xml:"groups"`
		Name	string	`xml:"name"`
		V_vpn  struct {
			XMLName xml.Name `xml:"vpn"`
			V_name  *string  `xml:"name,omitempty"`
			V_ike  struct {
				XMLName xml.Name `xml:"ike"`
				V_gateway  *string  `xml:"gateway,omitempty"`
			} `xml:"ike"`
		} `xml:"security>ipsec>vpn"`
	} `xml:"groups"`
	ApplyGroup string `xml:"apply-groups"`
}

// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex- interface is keyword in golang
func junosSecurityIpsecVpnIkeGatewayCreate(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_gateway := d.Get("gateway").(string)


	config := xmlSecurityIpsecVpnIkeGateway{}
	config.ApplyGroup = id
	config.Groups.Name = id
	config.Groups.V_vpn.V_name = &V_name
	config.Groups.V_vpn.V_ike.V_gateway = &V_gateway

    err = client.SendTransaction("", config, false)
    check(err)
    
    d.SetId(fmt.Sprintf("%s_%s", client.Host, id))
    
	return junosSecurityIpsecVpnIkeGatewayRead(d,m)
}

func junosSecurityIpsecVpnIkeGatewayRead(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	config := &xmlSecurityIpsecVpnIkeGateway{}

	err = client.MarshalGroup(id, config)
	check(err)
 	d.Set("name", config.Groups.V_vpn.V_name)
	d.Set("gateway", config.Groups.V_vpn.V_ike.V_gateway)

	return nil
}

func junosSecurityIpsecVpnIkeGatewayUpdate(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_gateway := d.Get("gateway").(string)


	config := xmlSecurityIpsecVpnIkeGateway{}
	config.ApplyGroup = id
	config.Groups.Name = id
	config.Groups.V_vpn.V_name = &V_name
	config.Groups.V_vpn.V_ike.V_gateway = &V_gateway

    err = client.SendTransaction(id, config, false)
    check(err)
    
	return junosSecurityIpsecVpnIkeGatewayRead(d,m)
}

func junosSecurityIpsecVpnIkeGatewayDelete(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
    _, err = client.DeleteConfigNoCommit(id)
    check(err)

    d.SetId("")
    
	return nil
}

func junosSecurityIpsecVpnIkeGateway() *schema.Resource {
	return &schema.Resource{
		Create: junosSecurityIpsecVpnIkeGatewayCreate,
		Read: junosSecurityIpsecVpnIkeGatewayRead,
		Update: junosSecurityIpsecVpnIkeGatewayUpdate,
		Delete: junosSecurityIpsecVpnIkeGatewayDelete,

        Schema: map[string]*schema.Schema{
            "resource_name": &schema.Schema{
                Type:    schema.TypeString,
                Required: true,
            },
			"name": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_vpn",
			},
			"gateway": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_vpn.V_ike. Name of remote gateway",
			},
		},
	}
}