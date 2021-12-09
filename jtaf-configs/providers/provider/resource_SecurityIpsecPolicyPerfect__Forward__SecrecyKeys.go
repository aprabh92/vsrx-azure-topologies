
package main

import (
    "encoding/xml"
    "fmt"
    "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)


// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex - interface is keyword in golang
type xmlSecurityIpsecPolicyPerfect__Forward__SecrecyKeys struct {
	XMLName xml.Name `xml:"configuration"`
	Groups  struct {
		XMLName	xml.Name	`xml:"groups"`
		Name	string	`xml:"name"`
		V_policy  struct {
			XMLName xml.Name `xml:"policy"`
			V_name  *string  `xml:"name,omitempty"`
			V_perfect__forward__secrecy  struct {
				XMLName xml.Name `xml:"perfect-forward-secrecy"`
				V_keys  *string  `xml:"keys,omitempty"`
			} `xml:"perfect-forward-secrecy"`
		} `xml:"security>ipsec>policy"`
	} `xml:"groups"`
	ApplyGroup string `xml:"apply-groups"`
}

// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex- interface is keyword in golang
func junosSecurityIpsecPolicyPerfect__Forward__SecrecyKeysCreate(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_keys := d.Get("keys").(string)


	config := xmlSecurityIpsecPolicyPerfect__Forward__SecrecyKeys{}
	config.ApplyGroup = id
	config.Groups.Name = id
	config.Groups.V_policy.V_name = &V_name
	config.Groups.V_policy.V_perfect__forward__secrecy.V_keys = &V_keys

    err = client.SendTransaction("", config, false)
    check(err)
    
    d.SetId(fmt.Sprintf("%s_%s", client.Host, id))
    
	return junosSecurityIpsecPolicyPerfect__Forward__SecrecyKeysRead(d,m)
}

func junosSecurityIpsecPolicyPerfect__Forward__SecrecyKeysRead(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	config := &xmlSecurityIpsecPolicyPerfect__Forward__SecrecyKeys{}

	err = client.MarshalGroup(id, config)
	check(err)
 	d.Set("name", config.Groups.V_policy.V_name)
	d.Set("keys", config.Groups.V_policy.V_perfect__forward__secrecy.V_keys)

	return nil
}

func junosSecurityIpsecPolicyPerfect__Forward__SecrecyKeysUpdate(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_keys := d.Get("keys").(string)


	config := xmlSecurityIpsecPolicyPerfect__Forward__SecrecyKeys{}
	config.ApplyGroup = id
	config.Groups.Name = id
	config.Groups.V_policy.V_name = &V_name
	config.Groups.V_policy.V_perfect__forward__secrecy.V_keys = &V_keys

    err = client.SendTransaction(id, config, false)
    check(err)
    
	return junosSecurityIpsecPolicyPerfect__Forward__SecrecyKeysRead(d,m)
}

func junosSecurityIpsecPolicyPerfect__Forward__SecrecyKeysDelete(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
    _, err = client.DeleteConfigNoCommit(id)
    check(err)

    d.SetId("")
    
	return nil
}

func junosSecurityIpsecPolicyPerfect__Forward__SecrecyKeys() *schema.Resource {
	return &schema.Resource{
		Create: junosSecurityIpsecPolicyPerfect__Forward__SecrecyKeysCreate,
		Read: junosSecurityIpsecPolicyPerfect__Forward__SecrecyKeysRead,
		Update: junosSecurityIpsecPolicyPerfect__Forward__SecrecyKeysUpdate,
		Delete: junosSecurityIpsecPolicyPerfect__Forward__SecrecyKeysDelete,

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
			"keys": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_policy.V_perfect__forward__secrecy. Define Diffie-Hellman group",
			},
		},
	}
}