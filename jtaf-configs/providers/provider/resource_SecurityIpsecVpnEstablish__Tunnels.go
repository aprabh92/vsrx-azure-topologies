
package main

import (
    "encoding/xml"
    "fmt"
    "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)


// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex - interface is keyword in golang
type xmlSecurityIpsecVpnEstablish__Tunnels struct {
	XMLName xml.Name `xml:"configuration"`
	Groups  struct {
		XMLName	xml.Name	`xml:"groups"`
		Name	string	`xml:"name"`
		V_vpn  struct {
			XMLName xml.Name `xml:"vpn"`
			V_name  *string  `xml:"name,omitempty"`
			V_establish__tunnels  *string  `xml:"establish-tunnels,omitempty"`
		} `xml:"security>ipsec>vpn"`
	} `xml:"groups"`
	ApplyGroup string `xml:"apply-groups"`
}

// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex- interface is keyword in golang
func junosSecurityIpsecVpnEstablish__TunnelsCreate(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_establish__tunnels := d.Get("establish__tunnels").(string)


	config := xmlSecurityIpsecVpnEstablish__Tunnels{}
	config.ApplyGroup = id
	config.Groups.Name = id
	config.Groups.V_vpn.V_name = &V_name
	config.Groups.V_vpn.V_establish__tunnels = &V_establish__tunnels

    err = client.SendTransaction("", config, false)
    check(err)
    
    d.SetId(fmt.Sprintf("%s_%s", client.Host, id))
    
	return junosSecurityIpsecVpnEstablish__TunnelsRead(d,m)
}

func junosSecurityIpsecVpnEstablish__TunnelsRead(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	config := &xmlSecurityIpsecVpnEstablish__Tunnels{}

	err = client.MarshalGroup(id, config)
	check(err)
 	d.Set("name", config.Groups.V_vpn.V_name)
	d.Set("establish__tunnels", config.Groups.V_vpn.V_establish__tunnels)

	return nil
}

func junosSecurityIpsecVpnEstablish__TunnelsUpdate(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_establish__tunnels := d.Get("establish__tunnels").(string)


	config := xmlSecurityIpsecVpnEstablish__Tunnels{}
	config.ApplyGroup = id
	config.Groups.Name = id
	config.Groups.V_vpn.V_name = &V_name
	config.Groups.V_vpn.V_establish__tunnels = &V_establish__tunnels

    err = client.SendTransaction(id, config, false)
    check(err)
    
	return junosSecurityIpsecVpnEstablish__TunnelsRead(d,m)
}

func junosSecurityIpsecVpnEstablish__TunnelsDelete(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
    _, err = client.DeleteConfigNoCommit(id)
    check(err)

    d.SetId("")
    
	return nil
}

func junosSecurityIpsecVpnEstablish__Tunnels() *schema.Resource {
	return &schema.Resource{
		Create: junosSecurityIpsecVpnEstablish__TunnelsCreate,
		Read: junosSecurityIpsecVpnEstablish__TunnelsRead,
		Update: junosSecurityIpsecVpnEstablish__TunnelsUpdate,
		Delete: junosSecurityIpsecVpnEstablish__TunnelsDelete,

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
			"establish__tunnels": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_vpn. Define the criteria to establish tunnels",
			},
		},
	}
}