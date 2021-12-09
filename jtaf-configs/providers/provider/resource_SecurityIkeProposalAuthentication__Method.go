
package main

import (
    "encoding/xml"
    "fmt"
    "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)


// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex - interface is keyword in golang
type xmlSecurityIkeProposalAuthentication__Method struct {
	XMLName xml.Name `xml:"configuration"`
	Groups  struct {
		XMLName	xml.Name	`xml:"groups"`
		Name	string	`xml:"name"`
		V_proposal  struct {
			XMLName xml.Name `xml:"proposal"`
			V_name  *string  `xml:"name,omitempty"`
			V_authentication__method  *string  `xml:"authentication-method,omitempty"`
		} `xml:"security>ike>proposal"`
	} `xml:"groups"`
	ApplyGroup string `xml:"apply-groups"`
}

// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex- interface is keyword in golang
func junosSecurityIkeProposalAuthentication__MethodCreate(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_authentication__method := d.Get("authentication__method").(string)


	config := xmlSecurityIkeProposalAuthentication__Method{}
	config.ApplyGroup = id
	config.Groups.Name = id
	config.Groups.V_proposal.V_name = &V_name
	config.Groups.V_proposal.V_authentication__method = &V_authentication__method

    err = client.SendTransaction("", config, false)
    check(err)
    
    d.SetId(fmt.Sprintf("%s_%s", client.Host, id))
    
	return junosSecurityIkeProposalAuthentication__MethodRead(d,m)
}

func junosSecurityIkeProposalAuthentication__MethodRead(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	config := &xmlSecurityIkeProposalAuthentication__Method{}

	err = client.MarshalGroup(id, config)
	check(err)
 	d.Set("name", config.Groups.V_proposal.V_name)
	d.Set("authentication__method", config.Groups.V_proposal.V_authentication__method)

	return nil
}

func junosSecurityIkeProposalAuthentication__MethodUpdate(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_authentication__method := d.Get("authentication__method").(string)


	config := xmlSecurityIkeProposalAuthentication__Method{}
	config.ApplyGroup = id
	config.Groups.Name = id
	config.Groups.V_proposal.V_name = &V_name
	config.Groups.V_proposal.V_authentication__method = &V_authentication__method

    err = client.SendTransaction(id, config, false)
    check(err)
    
	return junosSecurityIkeProposalAuthentication__MethodRead(d,m)
}

func junosSecurityIkeProposalAuthentication__MethodDelete(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
    _, err = client.DeleteConfigNoCommit(id)
    check(err)

    d.SetId("")
    
	return nil
}

func junosSecurityIkeProposalAuthentication__Method() *schema.Resource {
	return &schema.Resource{
		Create: junosSecurityIkeProposalAuthentication__MethodCreate,
		Read: junosSecurityIkeProposalAuthentication__MethodRead,
		Update: junosSecurityIkeProposalAuthentication__MethodUpdate,
		Delete: junosSecurityIkeProposalAuthentication__MethodDelete,

        Schema: map[string]*schema.Schema{
            "resource_name": &schema.Schema{
                Type:    schema.TypeString,
                Required: true,
            },
			"name": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_proposal",
			},
			"authentication__method": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_proposal. Define authentication method",
			},
		},
	}
}