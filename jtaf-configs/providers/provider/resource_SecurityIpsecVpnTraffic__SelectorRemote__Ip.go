
package main

import (
    "encoding/xml"
    "fmt"
    "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)


// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex - interface is keyword in golang
type xmlSecurityIpsecVpnTraffic__SelectorRemote__Ip struct {
	XMLName xml.Name `xml:"configuration"`
	Groups  struct {
		XMLName	xml.Name	`xml:"groups"`
		Name	string	`xml:"name"`
		V_vpn  struct {
			XMLName xml.Name `xml:"vpn"`
			V_name  *string  `xml:"name,omitempty"`
			V_traffic__selector  struct {
				XMLName xml.Name `xml:"traffic-selector"`
				V_name__1  *string  `xml:"name,omitempty"`
				V_remote__ip  *string  `xml:"remote-ip,omitempty"`
			} `xml:"traffic-selector"`
		} `xml:"security>ipsec>vpn"`
	} `xml:"groups"`
	ApplyGroup string `xml:"apply-groups"`
}

// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex- interface is keyword in golang
func junosSecurityIpsecVpnTraffic__SelectorRemote__IpCreate(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_name__1 := d.Get("name__1").(string)
	V_remote__ip := d.Get("remote__ip").(string)


	config := xmlSecurityIpsecVpnTraffic__SelectorRemote__Ip{}
	config.ApplyGroup = id
	config.Groups.Name = id
	config.Groups.V_vpn.V_name = &V_name
	config.Groups.V_vpn.V_traffic__selector.V_name__1 = &V_name__1
	config.Groups.V_vpn.V_traffic__selector.V_remote__ip = &V_remote__ip

    err = client.SendTransaction("", config, false)
    check(err)
    
    d.SetId(fmt.Sprintf("%s_%s", client.Host, id))
    
	return junosSecurityIpsecVpnTraffic__SelectorRemote__IpRead(d,m)
}

func junosSecurityIpsecVpnTraffic__SelectorRemote__IpRead(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	config := &xmlSecurityIpsecVpnTraffic__SelectorRemote__Ip{}

	err = client.MarshalGroup(id, config)
	check(err)
 	d.Set("name", config.Groups.V_vpn.V_name)
	d.Set("name__1", config.Groups.V_vpn.V_traffic__selector.V_name__1)
	d.Set("remote__ip", config.Groups.V_vpn.V_traffic__selector.V_remote__ip)

	return nil
}

func junosSecurityIpsecVpnTraffic__SelectorRemote__IpUpdate(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_name__1 := d.Get("name__1").(string)
	V_remote__ip := d.Get("remote__ip").(string)


	config := xmlSecurityIpsecVpnTraffic__SelectorRemote__Ip{}
	config.ApplyGroup = id
	config.Groups.Name = id
	config.Groups.V_vpn.V_name = &V_name
	config.Groups.V_vpn.V_traffic__selector.V_name__1 = &V_name__1
	config.Groups.V_vpn.V_traffic__selector.V_remote__ip = &V_remote__ip

    err = client.SendTransaction(id, config, false)
    check(err)
    
	return junosSecurityIpsecVpnTraffic__SelectorRemote__IpRead(d,m)
}

func junosSecurityIpsecVpnTraffic__SelectorRemote__IpDelete(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
    _, err = client.DeleteConfigNoCommit(id)
    check(err)

    d.SetId("")
    
	return nil
}

func junosSecurityIpsecVpnTraffic__SelectorRemote__Ip() *schema.Resource {
	return &schema.Resource{
		Create: junosSecurityIpsecVpnTraffic__SelectorRemote__IpCreate,
		Read: junosSecurityIpsecVpnTraffic__SelectorRemote__IpRead,
		Update: junosSecurityIpsecVpnTraffic__SelectorRemote__IpUpdate,
		Delete: junosSecurityIpsecVpnTraffic__SelectorRemote__IpDelete,

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
			"name__1": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_vpn.V_traffic__selector",
			},
			"remote__ip": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_vpn.V_traffic__selector. IP address of remote traffic-selector",
			},
		},
	}
}