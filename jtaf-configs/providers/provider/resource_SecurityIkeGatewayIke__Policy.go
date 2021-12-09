
package main

import (
    "encoding/xml"
    "fmt"
    "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)


// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex - interface is keyword in golang
type xmlSecurityIkeGatewayIke__Policy struct {
	XMLName xml.Name `xml:"configuration"`
	Groups  struct {
		XMLName	xml.Name	`xml:"groups"`
		Name	string	`xml:"name"`
		V_gateway  struct {
			XMLName xml.Name `xml:"gateway"`
			V_name  *string  `xml:"name,omitempty"`
			V_ike__policy  *string  `xml:"ike-policy,omitempty"`
		} `xml:"security>ike>gateway"`
	} `xml:"groups"`
	ApplyGroup string `xml:"apply-groups"`
}

// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex- interface is keyword in golang
func junosSecurityIkeGatewayIke__PolicyCreate(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_ike__policy := d.Get("ike__policy").(string)


	config := xmlSecurityIkeGatewayIke__Policy{}
	config.ApplyGroup = id
	config.Groups.Name = id
	config.Groups.V_gateway.V_name = &V_name
	config.Groups.V_gateway.V_ike__policy = &V_ike__policy

    err = client.SendTransaction("", config, false)
    check(err)
    
    d.SetId(fmt.Sprintf("%s_%s", client.Host, id))
    
	return junosSecurityIkeGatewayIke__PolicyRead(d,m)
}

func junosSecurityIkeGatewayIke__PolicyRead(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	config := &xmlSecurityIkeGatewayIke__Policy{}

	err = client.MarshalGroup(id, config)
	check(err)
 	d.Set("name", config.Groups.V_gateway.V_name)
	d.Set("ike__policy", config.Groups.V_gateway.V_ike__policy)

	return nil
}

func junosSecurityIkeGatewayIke__PolicyUpdate(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_ike__policy := d.Get("ike__policy").(string)


	config := xmlSecurityIkeGatewayIke__Policy{}
	config.ApplyGroup = id
	config.Groups.Name = id
	config.Groups.V_gateway.V_name = &V_name
	config.Groups.V_gateway.V_ike__policy = &V_ike__policy

    err = client.SendTransaction(id, config, false)
    check(err)
    
	return junosSecurityIkeGatewayIke__PolicyRead(d,m)
}

func junosSecurityIkeGatewayIke__PolicyDelete(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
    _, err = client.DeleteConfigNoCommit(id)
    check(err)

    d.SetId("")
    
	return nil
}

func junosSecurityIkeGatewayIke__Policy() *schema.Resource {
	return &schema.Resource{
		Create: junosSecurityIkeGatewayIke__PolicyCreate,
		Read: junosSecurityIkeGatewayIke__PolicyRead,
		Update: junosSecurityIkeGatewayIke__PolicyUpdate,
		Delete: junosSecurityIkeGatewayIke__PolicyDelete,

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
			"ike__policy": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_gateway. Name of the IKE policy",
			},
		},
	}
}