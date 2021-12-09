
package main

import (
    "encoding/xml"
    "fmt"
    "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)


// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex - interface is keyword in golang
type xmlSecurityIpsecProposalEncryption__Algorithm struct {
	XMLName xml.Name `xml:"configuration"`
	Groups  struct {
		XMLName	xml.Name	`xml:"groups"`
		Name	string	`xml:"name"`
		V_proposal  struct {
			XMLName xml.Name `xml:"proposal"`
			V_name  *string  `xml:"name,omitempty"`
			V_encryption__algorithm  *string  `xml:"encryption-algorithm,omitempty"`
		} `xml:"security>ipsec>proposal"`
	} `xml:"groups"`
	ApplyGroup string `xml:"apply-groups"`
}

// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex- interface is keyword in golang
func junosSecurityIpsecProposalEncryption__AlgorithmCreate(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_encryption__algorithm := d.Get("encryption__algorithm").(string)


	config := xmlSecurityIpsecProposalEncryption__Algorithm{}
	config.ApplyGroup = id
	config.Groups.Name = id
	config.Groups.V_proposal.V_name = &V_name
	config.Groups.V_proposal.V_encryption__algorithm = &V_encryption__algorithm

    err = client.SendTransaction("", config, false)
    check(err)
    
    d.SetId(fmt.Sprintf("%s_%s", client.Host, id))
    
	return junosSecurityIpsecProposalEncryption__AlgorithmRead(d,m)
}

func junosSecurityIpsecProposalEncryption__AlgorithmRead(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	config := &xmlSecurityIpsecProposalEncryption__Algorithm{}

	err = client.MarshalGroup(id, config)
	check(err)
 	d.Set("name", config.Groups.V_proposal.V_name)
	d.Set("encryption__algorithm", config.Groups.V_proposal.V_encryption__algorithm)

	return nil
}

func junosSecurityIpsecProposalEncryption__AlgorithmUpdate(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_encryption__algorithm := d.Get("encryption__algorithm").(string)


	config := xmlSecurityIpsecProposalEncryption__Algorithm{}
	config.ApplyGroup = id
	config.Groups.Name = id
	config.Groups.V_proposal.V_name = &V_name
	config.Groups.V_proposal.V_encryption__algorithm = &V_encryption__algorithm

    err = client.SendTransaction(id, config, false)
    check(err)
    
	return junosSecurityIpsecProposalEncryption__AlgorithmRead(d,m)
}

func junosSecurityIpsecProposalEncryption__AlgorithmDelete(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
    _, err = client.DeleteConfigNoCommit(id)
    check(err)

    d.SetId("")
    
	return nil
}

func junosSecurityIpsecProposalEncryption__Algorithm() *schema.Resource {
	return &schema.Resource{
		Create: junosSecurityIpsecProposalEncryption__AlgorithmCreate,
		Read: junosSecurityIpsecProposalEncryption__AlgorithmRead,
		Update: junosSecurityIpsecProposalEncryption__AlgorithmUpdate,
		Delete: junosSecurityIpsecProposalEncryption__AlgorithmDelete,

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
			"encryption__algorithm": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_proposal. Define encryption algorithm",
			},
		},
	}
}