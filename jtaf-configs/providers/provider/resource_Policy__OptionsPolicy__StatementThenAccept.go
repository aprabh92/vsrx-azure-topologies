
package main

import (
    "encoding/xml"
    "fmt"
    "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)


// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex - interface is keyword in golang
type xmlPolicy__OptionsPolicy__StatementThenAccept struct {
	XMLName xml.Name `xml:"configuration"`
	Groups  struct {
		XMLName	xml.Name	`xml:"groups"`
		Name	string	`xml:"name"`
		V_policy__statement  struct {
			XMLName xml.Name `xml:"policy-statement"`
			V_name  *string  `xml:"name,omitempty"`
			V_then  struct {
				XMLName xml.Name `xml:"then"`
				V_accept  *string  `xml:"accept,omitempty"`
			} `xml:"then"`
		} `xml:"policy-options>policy-statement"`
	} `xml:"groups"`
	ApplyGroup string `xml:"apply-groups"`
}

// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex- interface is keyword in golang
func junosPolicy__OptionsPolicy__StatementThenAcceptCreate(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_accept := d.Get("accept").(string)


	config := xmlPolicy__OptionsPolicy__StatementThenAccept{}
	config.ApplyGroup = id
	config.Groups.Name = id
	config.Groups.V_policy__statement.V_name = &V_name
	config.Groups.V_policy__statement.V_then.V_accept = &V_accept

    err = client.SendTransaction("", config, false)
    check(err)
    
    d.SetId(fmt.Sprintf("%s_%s", client.Host, id))
    
	return junosPolicy__OptionsPolicy__StatementThenAcceptRead(d,m)
}

func junosPolicy__OptionsPolicy__StatementThenAcceptRead(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	config := &xmlPolicy__OptionsPolicy__StatementThenAccept{}

	err = client.MarshalGroup(id, config)
	check(err)
 	d.Set("name", config.Groups.V_policy__statement.V_name)
	d.Set("accept", config.Groups.V_policy__statement.V_then.V_accept)

	return nil
}

func junosPolicy__OptionsPolicy__StatementThenAcceptUpdate(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_accept := d.Get("accept").(string)


	config := xmlPolicy__OptionsPolicy__StatementThenAccept{}
	config.ApplyGroup = id
	config.Groups.Name = id
	config.Groups.V_policy__statement.V_name = &V_name
	config.Groups.V_policy__statement.V_then.V_accept = &V_accept

    err = client.SendTransaction(id, config, false)
    check(err)
    
	return junosPolicy__OptionsPolicy__StatementThenAcceptRead(d,m)
}

func junosPolicy__OptionsPolicy__StatementThenAcceptDelete(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
    _, err = client.DeleteConfigNoCommit(id)
    check(err)

    d.SetId("")
    
	return nil
}

func junosPolicy__OptionsPolicy__StatementThenAccept() *schema.Resource {
	return &schema.Resource{
		Create: junosPolicy__OptionsPolicy__StatementThenAcceptCreate,
		Read: junosPolicy__OptionsPolicy__StatementThenAcceptRead,
		Update: junosPolicy__OptionsPolicy__StatementThenAcceptUpdate,
		Delete: junosPolicy__OptionsPolicy__StatementThenAcceptDelete,

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
			"accept": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_policy__statement.V_then. Accept a route",
			},
		},
	}
}