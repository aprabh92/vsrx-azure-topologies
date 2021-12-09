
package main

import (
    "encoding/xml"
    "fmt"
    "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)


// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex - interface is keyword in golang
type xmlSecurityIkeProposalDh__Group struct {
	XMLName xml.Name `xml:"configuration"`
	Groups  struct {
		XMLName	xml.Name	`xml:"groups"`
		Name	string	`xml:"name"`
		V_proposal  struct {
			XMLName xml.Name `xml:"proposal"`
			V_name  *string  `xml:"name,omitempty"`
			V_dh__group  *string  `xml:"dh-group,omitempty"`
		} `xml:"security>ike>proposal"`
	} `xml:"groups"`
	ApplyGroup string `xml:"apply-groups"`
}

// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex- interface is keyword in golang
func junosSecurityIkeProposalDh__GroupCreate(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_dh__group := d.Get("dh__group").(string)


	config := xmlSecurityIkeProposalDh__Group{}
	config.ApplyGroup = id
	config.Groups.Name = id
	config.Groups.V_proposal.V_name = &V_name
	config.Groups.V_proposal.V_dh__group = &V_dh__group

    err = client.SendTransaction("", config, false)
    check(err)
    
    d.SetId(fmt.Sprintf("%s_%s", client.Host, id))
    
	return junosSecurityIkeProposalDh__GroupRead(d,m)
}

func junosSecurityIkeProposalDh__GroupRead(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	config := &xmlSecurityIkeProposalDh__Group{}

	err = client.MarshalGroup(id, config)
	check(err)
 	d.Set("name", config.Groups.V_proposal.V_name)
	d.Set("dh__group", config.Groups.V_proposal.V_dh__group)

	return nil
}

func junosSecurityIkeProposalDh__GroupUpdate(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_dh__group := d.Get("dh__group").(string)


	config := xmlSecurityIkeProposalDh__Group{}
	config.ApplyGroup = id
	config.Groups.Name = id
	config.Groups.V_proposal.V_name = &V_name
	config.Groups.V_proposal.V_dh__group = &V_dh__group

    err = client.SendTransaction(id, config, false)
    check(err)
    
	return junosSecurityIkeProposalDh__GroupRead(d,m)
}

func junosSecurityIkeProposalDh__GroupDelete(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
    _, err = client.DeleteConfigNoCommit(id)
    check(err)

    d.SetId("")
    
	return nil
}

func junosSecurityIkeProposalDh__Group() *schema.Resource {
	return &schema.Resource{
		Create: junosSecurityIkeProposalDh__GroupCreate,
		Read: junosSecurityIkeProposalDh__GroupRead,
		Update: junosSecurityIkeProposalDh__GroupUpdate,
		Delete: junosSecurityIkeProposalDh__GroupDelete,

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
			"dh__group": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_proposal. Define Diffie-Hellman group",
			},
		},
	}
}