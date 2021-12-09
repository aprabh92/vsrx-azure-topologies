
package main

import (
    "encoding/xml"
    "fmt"
    "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)


// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex - interface is keyword in golang
type xmlPolicy__OptionsCommunityMembers struct {
	XMLName xml.Name `xml:"configuration"`
	Groups  struct {
		XMLName	xml.Name	`xml:"groups"`
		Name	string	`xml:"name"`
		V_community  struct {
			XMLName xml.Name `xml:"community"`
			V_name  *string  `xml:"name,omitempty"`
			V_members  *string  `xml:"members,omitempty"`
		} `xml:"policy-options>community"`
	} `xml:"groups"`
	ApplyGroup string `xml:"apply-groups"`
}

// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex- interface is keyword in golang
func junosPolicy__OptionsCommunityMembersCreate(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_members := d.Get("members").(string)


	config := xmlPolicy__OptionsCommunityMembers{}
	config.ApplyGroup = id
	config.Groups.Name = id
	config.Groups.V_community.V_name = &V_name
	config.Groups.V_community.V_members = &V_members

    err = client.SendTransaction("", config, false)
    check(err)
    
    d.SetId(fmt.Sprintf("%s_%s", client.Host, id))
    
	return junosPolicy__OptionsCommunityMembersRead(d,m)
}

func junosPolicy__OptionsCommunityMembersRead(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	config := &xmlPolicy__OptionsCommunityMembers{}

	err = client.MarshalGroup(id, config)
	check(err)
 	d.Set("name", config.Groups.V_community.V_name)
	d.Set("members", config.Groups.V_community.V_members)

	return nil
}

func junosPolicy__OptionsCommunityMembersUpdate(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_members := d.Get("members").(string)


	config := xmlPolicy__OptionsCommunityMembers{}
	config.ApplyGroup = id
	config.Groups.Name = id
	config.Groups.V_community.V_name = &V_name
	config.Groups.V_community.V_members = &V_members

    err = client.SendTransaction(id, config, false)
    check(err)
    
	return junosPolicy__OptionsCommunityMembersRead(d,m)
}

func junosPolicy__OptionsCommunityMembersDelete(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
    _, err = client.DeleteConfigNoCommit(id)
    check(err)

    d.SetId("")
    
	return nil
}

func junosPolicy__OptionsCommunityMembers() *schema.Resource {
	return &schema.Resource{
		Create: junosPolicy__OptionsCommunityMembersCreate,
		Read: junosPolicy__OptionsCommunityMembersRead,
		Update: junosPolicy__OptionsCommunityMembersUpdate,
		Delete: junosPolicy__OptionsCommunityMembersDelete,

        Schema: map[string]*schema.Schema{
            "resource_name": &schema.Schema{
                Type:    schema.TypeString,
                Required: true,
            },
			"name": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_community",
			},
			"members": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_community. Community members",
			},
		},
	}
}