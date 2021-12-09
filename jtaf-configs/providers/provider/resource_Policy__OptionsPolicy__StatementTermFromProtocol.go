
package main

import (
    "encoding/xml"
    "fmt"
    "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)


// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex - interface is keyword in golang
type xmlPolicy__OptionsPolicy__StatementTermFromProtocol struct {
	XMLName xml.Name `xml:"configuration"`
	Groups  struct {
		XMLName	xml.Name	`xml:"groups"`
		Name	string	`xml:"name"`
		V_policy__statement  struct {
			XMLName xml.Name `xml:"policy-statement"`
			V_name  *string  `xml:"name,omitempty"`
			V_term  struct {
				XMLName xml.Name `xml:"term"`
				V_name__1  *string  `xml:"name,omitempty"`
				V_from  struct {
					XMLName xml.Name `xml:"from"`
					V_protocol  *string  `xml:"protocol,omitempty"`
				} `xml:"from"`
			} `xml:"term"`
		} `xml:"policy-options>policy-statement"`
	} `xml:"groups"`
	ApplyGroup string `xml:"apply-groups"`
}

// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex- interface is keyword in golang
func junosPolicy__OptionsPolicy__StatementTermFromProtocolCreate(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_name__1 := d.Get("name__1").(string)
	V_protocol := d.Get("protocol").(string)


	config := xmlPolicy__OptionsPolicy__StatementTermFromProtocol{}
	config.ApplyGroup = id
	config.Groups.Name = id
	config.Groups.V_policy__statement.V_name = &V_name
	config.Groups.V_policy__statement.V_term.V_name__1 = &V_name__1
	config.Groups.V_policy__statement.V_term.V_from.V_protocol = &V_protocol

    err = client.SendTransaction("", config, false)
    check(err)
    
    d.SetId(fmt.Sprintf("%s_%s", client.Host, id))
    
	return junosPolicy__OptionsPolicy__StatementTermFromProtocolRead(d,m)
}

func junosPolicy__OptionsPolicy__StatementTermFromProtocolRead(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	config := &xmlPolicy__OptionsPolicy__StatementTermFromProtocol{}

	err = client.MarshalGroup(id, config)
	check(err)
 	d.Set("name", config.Groups.V_policy__statement.V_name)
	d.Set("name__1", config.Groups.V_policy__statement.V_term.V_name__1)
	d.Set("protocol", config.Groups.V_policy__statement.V_term.V_from.V_protocol)

	return nil
}

func junosPolicy__OptionsPolicy__StatementTermFromProtocolUpdate(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_name__1 := d.Get("name__1").(string)
	V_protocol := d.Get("protocol").(string)


	config := xmlPolicy__OptionsPolicy__StatementTermFromProtocol{}
	config.ApplyGroup = id
	config.Groups.Name = id
	config.Groups.V_policy__statement.V_name = &V_name
	config.Groups.V_policy__statement.V_term.V_name__1 = &V_name__1
	config.Groups.V_policy__statement.V_term.V_from.V_protocol = &V_protocol

    err = client.SendTransaction(id, config, false)
    check(err)
    
	return junosPolicy__OptionsPolicy__StatementTermFromProtocolRead(d,m)
}

func junosPolicy__OptionsPolicy__StatementTermFromProtocolDelete(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
    _, err = client.DeleteConfigNoCommit(id)
    check(err)

    d.SetId("")
    
	return nil
}

func junosPolicy__OptionsPolicy__StatementTermFromProtocol() *schema.Resource {
	return &schema.Resource{
		Create: junosPolicy__OptionsPolicy__StatementTermFromProtocolCreate,
		Read: junosPolicy__OptionsPolicy__StatementTermFromProtocolRead,
		Update: junosPolicy__OptionsPolicy__StatementTermFromProtocolUpdate,
		Delete: junosPolicy__OptionsPolicy__StatementTermFromProtocolDelete,

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
			"name__1": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_policy__statement.V_term",
			},
			"protocol": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_policy__statement.V_term.V_from. Protocol from which route was learned",
			},
		},
	}
}