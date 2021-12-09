
package main

import (
    "encoding/xml"
    "fmt"
    "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)


// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex - interface is keyword in golang
type xmlSecurityIpsecProposalAuthentication__Algorithm struct {
	XMLName xml.Name `xml:"configuration"`
	Groups  struct {
		XMLName	xml.Name	`xml:"groups"`
		Name	string	`xml:"name"`
		V_proposal  struct {
			XMLName xml.Name `xml:"proposal"`
			V_name  *string  `xml:"name,omitempty"`
			V_authentication__algorithm  *string  `xml:"authentication-algorithm,omitempty"`
		} `xml:"security>ipsec>proposal"`
	} `xml:"groups"`
	ApplyGroup string `xml:"apply-groups"`
}

// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex- interface is keyword in golang
func junosSecurityIpsecProposalAuthentication__AlgorithmCreate(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_authentication__algorithm := d.Get("authentication__algorithm").(string)


	config := xmlSecurityIpsecProposalAuthentication__Algorithm{}
	config.ApplyGroup = id
	config.Groups.Name = id
	config.Groups.V_proposal.V_name = &V_name
	config.Groups.V_proposal.V_authentication__algorithm = &V_authentication__algorithm

    err = client.SendTransaction("", config, false)
    check(err)
    
    d.SetId(fmt.Sprintf("%s_%s", client.Host, id))
    
	return junosSecurityIpsecProposalAuthentication__AlgorithmRead(d,m)
}

func junosSecurityIpsecProposalAuthentication__AlgorithmRead(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	config := &xmlSecurityIpsecProposalAuthentication__Algorithm{}

	err = client.MarshalGroup(id, config)
	check(err)
 	d.Set("name", config.Groups.V_proposal.V_name)
	d.Set("authentication__algorithm", config.Groups.V_proposal.V_authentication__algorithm)

	return nil
}

func junosSecurityIpsecProposalAuthentication__AlgorithmUpdate(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_authentication__algorithm := d.Get("authentication__algorithm").(string)


	config := xmlSecurityIpsecProposalAuthentication__Algorithm{}
	config.ApplyGroup = id
	config.Groups.Name = id
	config.Groups.V_proposal.V_name = &V_name
	config.Groups.V_proposal.V_authentication__algorithm = &V_authentication__algorithm

    err = client.SendTransaction(id, config, false)
    check(err)
    
	return junosSecurityIpsecProposalAuthentication__AlgorithmRead(d,m)
}

func junosSecurityIpsecProposalAuthentication__AlgorithmDelete(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
    _, err = client.DeleteConfigNoCommit(id)
    check(err)

    d.SetId("")
    
	return nil
}

func junosSecurityIpsecProposalAuthentication__Algorithm() *schema.Resource {
	return &schema.Resource{
		Create: junosSecurityIpsecProposalAuthentication__AlgorithmCreate,
		Read: junosSecurityIpsecProposalAuthentication__AlgorithmRead,
		Update: junosSecurityIpsecProposalAuthentication__AlgorithmUpdate,
		Delete: junosSecurityIpsecProposalAuthentication__AlgorithmDelete,

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
			"authentication__algorithm": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_proposal. Define authentication algorithm",
			},
		},
	}
}