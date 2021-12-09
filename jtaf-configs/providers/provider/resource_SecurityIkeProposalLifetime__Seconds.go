
package main

import (
    "encoding/xml"
    "fmt"
    "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)


// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex - interface is keyword in golang
type xmlSecurityIkeProposalLifetime__Seconds struct {
	XMLName xml.Name `xml:"configuration"`
	Groups  struct {
		XMLName	xml.Name	`xml:"groups"`
		Name	string	`xml:"name"`
		V_proposal  struct {
			XMLName xml.Name `xml:"proposal"`
			V_name  *string  `xml:"name,omitempty"`
			V_lifetime__seconds  *string  `xml:"lifetime-seconds,omitempty"`
		} `xml:"security>ike>proposal"`
	} `xml:"groups"`
	ApplyGroup string `xml:"apply-groups"`
}

// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex- interface is keyword in golang
func junosSecurityIkeProposalLifetime__SecondsCreate(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_lifetime__seconds := d.Get("lifetime__seconds").(string)


	config := xmlSecurityIkeProposalLifetime__Seconds{}
	config.ApplyGroup = id
	config.Groups.Name = id
	config.Groups.V_proposal.V_name = &V_name
	config.Groups.V_proposal.V_lifetime__seconds = &V_lifetime__seconds

    err = client.SendTransaction("", config, false)
    check(err)
    
    d.SetId(fmt.Sprintf("%s_%s", client.Host, id))
    
	return junosSecurityIkeProposalLifetime__SecondsRead(d,m)
}

func junosSecurityIkeProposalLifetime__SecondsRead(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	config := &xmlSecurityIkeProposalLifetime__Seconds{}

	err = client.MarshalGroup(id, config)
	check(err)
 	d.Set("name", config.Groups.V_proposal.V_name)
	d.Set("lifetime__seconds", config.Groups.V_proposal.V_lifetime__seconds)

	return nil
}

func junosSecurityIkeProposalLifetime__SecondsUpdate(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_lifetime__seconds := d.Get("lifetime__seconds").(string)


	config := xmlSecurityIkeProposalLifetime__Seconds{}
	config.ApplyGroup = id
	config.Groups.Name = id
	config.Groups.V_proposal.V_name = &V_name
	config.Groups.V_proposal.V_lifetime__seconds = &V_lifetime__seconds

    err = client.SendTransaction(id, config, false)
    check(err)
    
	return junosSecurityIkeProposalLifetime__SecondsRead(d,m)
}

func junosSecurityIkeProposalLifetime__SecondsDelete(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
    _, err = client.DeleteConfigNoCommit(id)
    check(err)

    d.SetId("")
    
	return nil
}

func junosSecurityIkeProposalLifetime__Seconds() *schema.Resource {
	return &schema.Resource{
		Create: junosSecurityIkeProposalLifetime__SecondsCreate,
		Read: junosSecurityIkeProposalLifetime__SecondsRead,
		Update: junosSecurityIkeProposalLifetime__SecondsUpdate,
		Delete: junosSecurityIkeProposalLifetime__SecondsDelete,

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
			"lifetime__seconds": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_proposal. Lifetime, in seconds",
			},
		},
	}
}