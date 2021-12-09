
package main

import (
    "encoding/xml"
    "fmt"
    "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)


// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex - interface is keyword in golang
type xmlSecurityIpsecPolicyProposals struct {
	XMLName xml.Name `xml:"configuration"`
	Groups  struct {
		XMLName	xml.Name	`xml:"groups"`
		Name	string	`xml:"name"`
		V_policy  struct {
			XMLName xml.Name `xml:"policy"`
			V_name  *string  `xml:"name,omitempty"`
			V_proposals  *string  `xml:"proposals,omitempty"`
		} `xml:"security>ipsec>policy"`
	} `xml:"groups"`
	ApplyGroup string `xml:"apply-groups"`
}

// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex- interface is keyword in golang
func junosSecurityIpsecPolicyProposalsCreate(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_proposals := d.Get("proposals").(string)


	config := xmlSecurityIpsecPolicyProposals{}
	config.ApplyGroup = id
	config.Groups.Name = id
	config.Groups.V_policy.V_name = &V_name
	config.Groups.V_policy.V_proposals = &V_proposals

    err = client.SendTransaction("", config, false)
    check(err)
    
    d.SetId(fmt.Sprintf("%s_%s", client.Host, id))
    
	return junosSecurityIpsecPolicyProposalsRead(d,m)
}

func junosSecurityIpsecPolicyProposalsRead(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	config := &xmlSecurityIpsecPolicyProposals{}

	err = client.MarshalGroup(id, config)
	check(err)
 	d.Set("name", config.Groups.V_policy.V_name)
	d.Set("proposals", config.Groups.V_policy.V_proposals)

	return nil
}

func junosSecurityIpsecPolicyProposalsUpdate(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_proposals := d.Get("proposals").(string)


	config := xmlSecurityIpsecPolicyProposals{}
	config.ApplyGroup = id
	config.Groups.Name = id
	config.Groups.V_policy.V_name = &V_name
	config.Groups.V_policy.V_proposals = &V_proposals

    err = client.SendTransaction(id, config, false)
    check(err)
    
	return junosSecurityIpsecPolicyProposalsRead(d,m)
}

func junosSecurityIpsecPolicyProposalsDelete(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
    _, err = client.DeleteConfigNoCommit(id)
    check(err)

    d.SetId("")
    
	return nil
}

func junosSecurityIpsecPolicyProposals() *schema.Resource {
	return &schema.Resource{
		Create: junosSecurityIpsecPolicyProposalsCreate,
		Read: junosSecurityIpsecPolicyProposalsRead,
		Update: junosSecurityIpsecPolicyProposalsUpdate,
		Delete: junosSecurityIpsecPolicyProposalsDelete,

        Schema: map[string]*schema.Schema{
            "resource_name": &schema.Schema{
                Type:    schema.TypeString,
                Required: true,
            },
			"name": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_policy",
			},
			"proposals": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_policy. Name of the proposal",
			},
		},
	}
}