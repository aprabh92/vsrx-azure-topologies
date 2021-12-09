
package main

import (
    "encoding/xml"
    "fmt"
    "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)


// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex - interface is keyword in golang
type xmlSecurityIpsecVpnBind__Interface struct {
	XMLName xml.Name `xml:"configuration"`
	Groups  struct {
		XMLName	xml.Name	`xml:"groups"`
		Name	string	`xml:"name"`
		V_vpn  struct {
			XMLName xml.Name `xml:"vpn"`
			V_name  *string  `xml:"name,omitempty"`
			V_bind__interface  *string  `xml:"bind-interface,omitempty"`
		} `xml:"security>ipsec>vpn"`
	} `xml:"groups"`
	ApplyGroup string `xml:"apply-groups"`
}

// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex- interface is keyword in golang
func junosSecurityIpsecVpnBind__InterfaceCreate(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_bind__interface := d.Get("bind__interface").(string)


	config := xmlSecurityIpsecVpnBind__Interface{}
	config.ApplyGroup = id
	config.Groups.Name = id
	config.Groups.V_vpn.V_name = &V_name
	config.Groups.V_vpn.V_bind__interface = &V_bind__interface

    err = client.SendTransaction("", config, false)
    check(err)
    
    d.SetId(fmt.Sprintf("%s_%s", client.Host, id))
    
	return junosSecurityIpsecVpnBind__InterfaceRead(d,m)
}

func junosSecurityIpsecVpnBind__InterfaceRead(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	config := &xmlSecurityIpsecVpnBind__Interface{}

	err = client.MarshalGroup(id, config)
	check(err)
 	d.Set("name", config.Groups.V_vpn.V_name)
	d.Set("bind__interface", config.Groups.V_vpn.V_bind__interface)

	return nil
}

func junosSecurityIpsecVpnBind__InterfaceUpdate(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_bind__interface := d.Get("bind__interface").(string)


	config := xmlSecurityIpsecVpnBind__Interface{}
	config.ApplyGroup = id
	config.Groups.Name = id
	config.Groups.V_vpn.V_name = &V_name
	config.Groups.V_vpn.V_bind__interface = &V_bind__interface

    err = client.SendTransaction(id, config, false)
    check(err)
    
	return junosSecurityIpsecVpnBind__InterfaceRead(d,m)
}

func junosSecurityIpsecVpnBind__InterfaceDelete(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
    _, err = client.DeleteConfigNoCommit(id)
    check(err)

    d.SetId("")
    
	return nil
}

func junosSecurityIpsecVpnBind__Interface() *schema.Resource {
	return &schema.Resource{
		Create: junosSecurityIpsecVpnBind__InterfaceCreate,
		Read: junosSecurityIpsecVpnBind__InterfaceRead,
		Update: junosSecurityIpsecVpnBind__InterfaceUpdate,
		Delete: junosSecurityIpsecVpnBind__InterfaceDelete,

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
			"bind__interface": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_vpn. Bind to tunnel interface (route-based VPN)",
			},
		},
	}
}