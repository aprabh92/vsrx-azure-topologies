
package main

import (
    "encoding/xml"
    "fmt"
    "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)


// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex - interface is keyword in golang
type xmlPolicy__OptionsPolicy__StatementThenNext struct {
	XMLName xml.Name `xml:"configuration"`
	Groups  struct {
		XMLName	xml.Name	`xml:"groups"`
		Name	string	`xml:"name"`
		V_policy__statement  struct {
			XMLName xml.Name `xml:"policy-statement"`
			V_name  *string  `xml:"name,omitempty"`
			V_then  struct {
				XMLName xml.Name `xml:"then"`
				V_next  *string  `xml:"next,omitempty"`
			} `xml:"then"`
		} `xml:"policy-options>policy-statement"`
	} `xml:"groups"`
	ApplyGroup string `xml:"apply-groups"`
}

// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex- interface is keyword in golang
func junosPolicy__OptionsPolicy__StatementThenNextCreate(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_next := d.Get("next").(string)


	config := xmlPolicy__OptionsPolicy__StatementThenNext{}
	config.ApplyGroup = id
	config.Groups.Name = id
	config.Groups.V_policy__statement.V_name = &V_name
	config.Groups.V_policy__statement.V_then.V_next = &V_next

    err = client.SendTransaction("", config, false)
    check(err)
    
    d.SetId(fmt.Sprintf("%s_%s", client.Host, id))
    
	return junosPolicy__OptionsPolicy__StatementThenNextRead(d,m)
}

func junosPolicy__OptionsPolicy__StatementThenNextRead(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	config := &xmlPolicy__OptionsPolicy__StatementThenNext{}

	err = client.MarshalGroup(id, config)
	check(err)
 	d.Set("name", config.Groups.V_policy__statement.V_name)
	d.Set("next", config.Groups.V_policy__statement.V_then.V_next)

	return nil
}

func junosPolicy__OptionsPolicy__StatementThenNextUpdate(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_next := d.Get("next").(string)


	config := xmlPolicy__OptionsPolicy__StatementThenNext{}
	config.ApplyGroup = id
	config.Groups.Name = id
	config.Groups.V_policy__statement.V_name = &V_name
	config.Groups.V_policy__statement.V_then.V_next = &V_next

    err = client.SendTransaction(id, config, false)
    check(err)
    
	return junosPolicy__OptionsPolicy__StatementThenNextRead(d,m)
}

func junosPolicy__OptionsPolicy__StatementThenNextDelete(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
    _, err = client.DeleteConfigNoCommit(id)
    check(err)

    d.SetId("")
    
	return nil
}

func junosPolicy__OptionsPolicy__StatementThenNext() *schema.Resource {
	return &schema.Resource{
		Create: junosPolicy__OptionsPolicy__StatementThenNextCreate,
		Read: junosPolicy__OptionsPolicy__StatementThenNextRead,
		Update: junosPolicy__OptionsPolicy__StatementThenNextUpdate,
		Delete: junosPolicy__OptionsPolicy__StatementThenNextDelete,

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
			"next": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_policy__statement.V_then. Skip to next policy or term",
			},
		},
	}
}