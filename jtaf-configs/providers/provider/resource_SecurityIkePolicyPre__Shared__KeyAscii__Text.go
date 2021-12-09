
package main

import (
    "encoding/xml"
    "fmt"
    "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)


// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex - interface is keyword in golang
type xmlSecurityIkePolicyPre__Shared__KeyAscii__Text struct {
	XMLName xml.Name `xml:"configuration"`
	Groups  struct {
		XMLName	xml.Name	`xml:"groups"`
		Name	string	`xml:"name"`
		V_policy  struct {
			XMLName xml.Name `xml:"policy"`
			V_name  *string  `xml:"name,omitempty"`
			V_pre__shared__key  struct {
				XMLName xml.Name `xml:"pre-shared-key"`
				V_ascii__text  *string  `xml:"ascii-text,omitempty"`
			} `xml:"pre-shared-key"`
		} `xml:"security>ike>policy"`
	} `xml:"groups"`
	ApplyGroup string `xml:"apply-groups"`
}

// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex- interface is keyword in golang
func junosSecurityIkePolicyPre__Shared__KeyAscii__TextCreate(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_ascii__text := d.Get("ascii__text").(string)


	config := xmlSecurityIkePolicyPre__Shared__KeyAscii__Text{}
	config.ApplyGroup = id
	config.Groups.Name = id
	config.Groups.V_policy.V_name = &V_name
	config.Groups.V_policy.V_pre__shared__key.V_ascii__text = &V_ascii__text

    err = client.SendTransaction("", config, false)
    check(err)
    
    d.SetId(fmt.Sprintf("%s_%s", client.Host, id))
    
	return junosSecurityIkePolicyPre__Shared__KeyAscii__TextRead(d,m)
}

func junosSecurityIkePolicyPre__Shared__KeyAscii__TextRead(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	config := &xmlSecurityIkePolicyPre__Shared__KeyAscii__Text{}

	err = client.MarshalGroup(id, config)
	check(err)
 	d.Set("name", config.Groups.V_policy.V_name)
	d.Set("ascii__text", config.Groups.V_policy.V_pre__shared__key.V_ascii__text)

	return nil
}

func junosSecurityIkePolicyPre__Shared__KeyAscii__TextUpdate(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_ascii__text := d.Get("ascii__text").(string)


	config := xmlSecurityIkePolicyPre__Shared__KeyAscii__Text{}
	config.ApplyGroup = id
	config.Groups.Name = id
	config.Groups.V_policy.V_name = &V_name
	config.Groups.V_policy.V_pre__shared__key.V_ascii__text = &V_ascii__text

    err = client.SendTransaction(id, config, false)
    check(err)
    
	return junosSecurityIkePolicyPre__Shared__KeyAscii__TextRead(d,m)
}

func junosSecurityIkePolicyPre__Shared__KeyAscii__TextDelete(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
    _, err = client.DeleteConfigNoCommit(id)
    check(err)

    d.SetId("")
    
	return nil
}

func junosSecurityIkePolicyPre__Shared__KeyAscii__Text() *schema.Resource {
	return &schema.Resource{
		Create: junosSecurityIkePolicyPre__Shared__KeyAscii__TextCreate,
		Read: junosSecurityIkePolicyPre__Shared__KeyAscii__TextRead,
		Update: junosSecurityIkePolicyPre__Shared__KeyAscii__TextUpdate,
		Delete: junosSecurityIkePolicyPre__Shared__KeyAscii__TextDelete,

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
			"ascii__text": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_policy.V_pre__shared__key. Format as text",
			},
		},
	}
}