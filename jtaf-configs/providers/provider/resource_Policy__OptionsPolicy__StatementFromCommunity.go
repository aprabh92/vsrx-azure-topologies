
package main

import (
    "encoding/xml"
    "fmt"
    "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)


// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex - interface is keyword in golang
type xmlPolicy__OptionsPolicy__StatementFromCommunity struct {
	XMLName xml.Name `xml:"configuration"`
	Groups  struct {
		XMLName	xml.Name	`xml:"groups"`
		Name	string	`xml:"name"`
		V_policy__statement  struct {
			XMLName xml.Name `xml:"policy-statement"`
			V_name  *string  `xml:"name,omitempty"`
			V_from  struct {
				XMLName xml.Name `xml:"from"`
				V_community  *string  `xml:"community,omitempty"`
			} `xml:"from"`
		} `xml:"policy-options>policy-statement"`
	} `xml:"groups"`
	ApplyGroup string `xml:"apply-groups"`
}

// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex- interface is keyword in golang
func junosPolicy__OptionsPolicy__StatementFromCommunityCreate(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_community := d.Get("community").(string)


	config := xmlPolicy__OptionsPolicy__StatementFromCommunity{}
	config.ApplyGroup = id
	config.Groups.Name = id
	config.Groups.V_policy__statement.V_name = &V_name
	config.Groups.V_policy__statement.V_from.V_community = &V_community

    err = client.SendTransaction("", config, false)
    check(err)
    
    d.SetId(fmt.Sprintf("%s_%s", client.Host, id))
    
	return junosPolicy__OptionsPolicy__StatementFromCommunityRead(d,m)
}

func junosPolicy__OptionsPolicy__StatementFromCommunityRead(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	config := &xmlPolicy__OptionsPolicy__StatementFromCommunity{}

	err = client.MarshalGroup(id, config)
	check(err)
 	d.Set("name", config.Groups.V_policy__statement.V_name)
	d.Set("community", config.Groups.V_policy__statement.V_from.V_community)

	return nil
}

func junosPolicy__OptionsPolicy__StatementFromCommunityUpdate(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_community := d.Get("community").(string)


	config := xmlPolicy__OptionsPolicy__StatementFromCommunity{}
	config.ApplyGroup = id
	config.Groups.Name = id
	config.Groups.V_policy__statement.V_name = &V_name
	config.Groups.V_policy__statement.V_from.V_community = &V_community

    err = client.SendTransaction(id, config, false)
    check(err)
    
	return junosPolicy__OptionsPolicy__StatementFromCommunityRead(d,m)
}

func junosPolicy__OptionsPolicy__StatementFromCommunityDelete(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
    _, err = client.DeleteConfigNoCommit(id)
    check(err)

    d.SetId("")
    
	return nil
}

func junosPolicy__OptionsPolicy__StatementFromCommunity() *schema.Resource {
	return &schema.Resource{
		Create: junosPolicy__OptionsPolicy__StatementFromCommunityCreate,
		Read: junosPolicy__OptionsPolicy__StatementFromCommunityRead,
		Update: junosPolicy__OptionsPolicy__StatementFromCommunityUpdate,
		Delete: junosPolicy__OptionsPolicy__StatementFromCommunityDelete,

        Schema: map[string]*schema.Schema{
            "resource_name": &schema.Schema{
                Type:    schema.TypeString,
                Required: true,
            },
			"name": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_policy__statement",
			},
			"community": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_policy__statement.V_from. BGP community",
			},
		},
	}
}