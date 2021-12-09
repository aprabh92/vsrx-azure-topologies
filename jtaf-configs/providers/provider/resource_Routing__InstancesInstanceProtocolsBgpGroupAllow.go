
package main

import (
    "encoding/xml"
    "fmt"
    "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)


// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex - interface is keyword in golang
type xmlRouting__InstancesInstanceProtocolsBgpGroupAllow struct {
	XMLName xml.Name `xml:"configuration"`
	Groups  struct {
		XMLName	xml.Name	`xml:"groups"`
		Name	string	`xml:"name"`
		V_instance  struct {
			XMLName xml.Name `xml:"instance"`
			V_name  *string  `xml:"name,omitempty"`
			V_group  struct {
				XMLName xml.Name `xml:"group"`
				V_name__1  *string  `xml:"name,omitempty"`
				V_allow  *string  `xml:"allow,omitempty"`
			} `xml:"protocols>bgp>group"`
		} `xml:"routing-instances>instance"`
	} `xml:"groups"`
	ApplyGroup string `xml:"apply-groups"`
}

// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex- interface is keyword in golang
func junosRouting__InstancesInstanceProtocolsBgpGroupAllowCreate(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_name__1 := d.Get("name__1").(string)
	V_allow := d.Get("allow").(string)


	config := xmlRouting__InstancesInstanceProtocolsBgpGroupAllow{}
	config.ApplyGroup = id
	config.Groups.Name = id
	config.Groups.V_instance.V_name = &V_name
	config.Groups.V_instance.V_group.V_name__1 = &V_name__1
	config.Groups.V_instance.V_group.V_allow = &V_allow

    err = client.SendTransaction("", config, false)
    check(err)
    
    d.SetId(fmt.Sprintf("%s_%s", client.Host, id))
    
	return junosRouting__InstancesInstanceProtocolsBgpGroupAllowRead(d,m)
}

func junosRouting__InstancesInstanceProtocolsBgpGroupAllowRead(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	config := &xmlRouting__InstancesInstanceProtocolsBgpGroupAllow{}

	err = client.MarshalGroup(id, config)
	check(err)
 	d.Set("name", config.Groups.V_instance.V_name)
	d.Set("name__1", config.Groups.V_instance.V_group.V_name__1)
	d.Set("allow", config.Groups.V_instance.V_group.V_allow)

	return nil
}

func junosRouting__InstancesInstanceProtocolsBgpGroupAllowUpdate(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_name__1 := d.Get("name__1").(string)
	V_allow := d.Get("allow").(string)


	config := xmlRouting__InstancesInstanceProtocolsBgpGroupAllow{}
	config.ApplyGroup = id
	config.Groups.Name = id
	config.Groups.V_instance.V_name = &V_name
	config.Groups.V_instance.V_group.V_name__1 = &V_name__1
	config.Groups.V_instance.V_group.V_allow = &V_allow

    err = client.SendTransaction(id, config, false)
    check(err)
    
	return junosRouting__InstancesInstanceProtocolsBgpGroupAllowRead(d,m)
}

func junosRouting__InstancesInstanceProtocolsBgpGroupAllowDelete(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
    _, err = client.DeleteConfigNoCommit(id)
    check(err)

    d.SetId("")
    
	return nil
}

func junosRouting__InstancesInstanceProtocolsBgpGroupAllow() *schema.Resource {
	return &schema.Resource{
		Create: junosRouting__InstancesInstanceProtocolsBgpGroupAllowCreate,
		Read: junosRouting__InstancesInstanceProtocolsBgpGroupAllowRead,
		Update: junosRouting__InstancesInstanceProtocolsBgpGroupAllowUpdate,
		Delete: junosRouting__InstancesInstanceProtocolsBgpGroupAllowDelete,

        Schema: map[string]*schema.Schema{
            "resource_name": &schema.Schema{
                Type:    schema.TypeString,
                Required: true,
            },
			"name": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_instance",
			},
			"name__1": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_instance.V_group",
			},
			"allow": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_instance.V_group. Configure peer connections for specific networks",
			},
		},
	}
}