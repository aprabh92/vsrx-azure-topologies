
package main

import (
    "encoding/xml"
    "fmt"
    "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)


// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex - interface is keyword in golang
type xmlSecurityIpsecVpnIkeIpsec__Policy struct {
	XMLName xml.Name `xml:"configuration"`
	Groups  struct {
		XMLName	xml.Name	`xml:"groups"`
		Name	string	`xml:"name"`
		V_vpn  struct {
			XMLName xml.Name `xml:"vpn"`
			V_name  *string  `xml:"name,omitempty"`
			V_ike  struct {
				XMLName xml.Name `xml:"ike"`
				V_ipsec__policy  *string  `xml:"ipsec-policy,omitempty"`
			} `xml:"ike"`
		} `xml:"security>ipsec>vpn"`
	} `xml:"groups"`
	ApplyGroup string `xml:"apply-groups"`
}

// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex- interface is keyword in golang
func junosSecurityIpsecVpnIkeIpsec__PolicyCreate(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_ipsec__policy := d.Get("ipsec__policy").(string)


	config := xmlSecurityIpsecVpnIkeIpsec__Policy{}
	config.ApplyGroup = id
	config.Groups.Name = id
	config.Groups.V_vpn.V_name = &V_name
	config.Groups.V_vpn.V_ike.V_ipsec__policy = &V_ipsec__policy

    err = client.SendTransaction("", config, false)
    check(err)
    
    d.SetId(fmt.Sprintf("%s_%s", client.Host, id))
    
	return junosSecurityIpsecVpnIkeIpsec__PolicyRead(d,m)
}

func junosSecurityIpsecVpnIkeIpsec__PolicyRead(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	config := &xmlSecurityIpsecVpnIkeIpsec__Policy{}

	err = client.MarshalGroup(id, config)
	check(err)
 	d.Set("name", config.Groups.V_vpn.V_name)
	d.Set("ipsec__policy", config.Groups.V_vpn.V_ike.V_ipsec__policy)

	return nil
}

func junosSecurityIpsecVpnIkeIpsec__PolicyUpdate(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_ipsec__policy := d.Get("ipsec__policy").(string)


	config := xmlSecurityIpsecVpnIkeIpsec__Policy{}
	config.ApplyGroup = id
	config.Groups.Name = id
	config.Groups.V_vpn.V_name = &V_name
	config.Groups.V_vpn.V_ike.V_ipsec__policy = &V_ipsec__policy

    err = client.SendTransaction(id, config, false)
    check(err)
    
	return junosSecurityIpsecVpnIkeIpsec__PolicyRead(d,m)
}

func junosSecurityIpsecVpnIkeIpsec__PolicyDelete(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
    _, err = client.DeleteConfigNoCommit(id)
    check(err)

    d.SetId("")
    
	return nil
}

func junosSecurityIpsecVpnIkeIpsec__Policy() *schema.Resource {
	return &schema.Resource{
		Create: junosSecurityIpsecVpnIkeIpsec__PolicyCreate,
		Read: junosSecurityIpsecVpnIkeIpsec__PolicyRead,
		Update: junosSecurityIpsecVpnIkeIpsec__PolicyUpdate,
		Delete: junosSecurityIpsecVpnIkeIpsec__PolicyDelete,

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
			"ipsec__policy": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_vpn.V_ike. Name of the IPSec policy",
			},
		},
	}
}