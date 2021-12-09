
package main

import (
    "encoding/xml"
    "fmt"
    "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)


// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex - interface is keyword in golang
type xmlPolicy__OptionsPolicy__StatementThenAs__Path__Prepend struct {
	XMLName xml.Name `xml:"configuration"`
	Groups  struct {
		XMLName	xml.Name	`xml:"groups"`
		Name	string	`xml:"name"`
		V_policy__statement  struct {
			XMLName xml.Name `xml:"policy-statement"`
			V_name  *string  `xml:"name,omitempty"`
			V_then  struct {
				XMLName xml.Name `xml:"then"`
				V_as__path__prepend  *string  `xml:"as-path-prepend,omitempty"`
			} `xml:"then"`
		} `xml:"policy-options>policy-statement"`
	} `xml:"groups"`
	ApplyGroup string `xml:"apply-groups"`
}

// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex- interface is keyword in golang
func junosPolicy__OptionsPolicy__StatementThenAs__Path__PrependCreate(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_as__path__prepend := d.Get("as__path__prepend").(string)


	config := xmlPolicy__OptionsPolicy__StatementThenAs__Path__Prepend{}
	config.ApplyGroup = id
	config.Groups.Name = id
	config.Groups.V_policy__statement.V_name = &V_name
	config.Groups.V_policy__statement.V_then.V_as__path__prepend = &V_as__path__prepend

    err = client.SendTransaction("", config, false)
    check(err)
    
    d.SetId(fmt.Sprintf("%s_%s", client.Host, id))
    
	return junosPolicy__OptionsPolicy__StatementThenAs__Path__PrependRead(d,m)
}

func junosPolicy__OptionsPolicy__StatementThenAs__Path__PrependRead(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	config := &xmlPolicy__OptionsPolicy__StatementThenAs__Path__Prepend{}

	err = client.MarshalGroup(id, config)
	check(err)
 	d.Set("name", config.Groups.V_policy__statement.V_name)
	d.Set("as__path__prepend", config.Groups.V_policy__statement.V_then.V_as__path__prepend)

	return nil
}

func junosPolicy__OptionsPolicy__StatementThenAs__Path__PrependUpdate(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_as__path__prepend := d.Get("as__path__prepend").(string)


	config := xmlPolicy__OptionsPolicy__StatementThenAs__Path__Prepend{}
	config.ApplyGroup = id
	config.Groups.Name = id
	config.Groups.V_policy__statement.V_name = &V_name
	config.Groups.V_policy__statement.V_then.V_as__path__prepend = &V_as__path__prepend

    err = client.SendTransaction(id, config, false)
    check(err)
    
	return junosPolicy__OptionsPolicy__StatementThenAs__Path__PrependRead(d,m)
}

func junosPolicy__OptionsPolicy__StatementThenAs__Path__PrependDelete(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
    _, err = client.DeleteConfigNoCommit(id)
    check(err)

    d.SetId("")
    
	return nil
}

func junosPolicy__OptionsPolicy__StatementThenAs__Path__Prepend() *schema.Resource {
	return &schema.Resource{
		Create: junosPolicy__OptionsPolicy__StatementThenAs__Path__PrependCreate,
		Read: junosPolicy__OptionsPolicy__StatementThenAs__Path__PrependRead,
		Update: junosPolicy__OptionsPolicy__StatementThenAs__Path__PrependUpdate,
		Delete: junosPolicy__OptionsPolicy__StatementThenAs__Path__PrependDelete,

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
			"as__path__prepend": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_policy__statement.V_then. Prepend AS numbers to an AS path (BGP only)",
			},
		},
	}
}