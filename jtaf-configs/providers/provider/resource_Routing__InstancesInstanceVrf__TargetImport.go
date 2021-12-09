
package main

import (
    "encoding/xml"
    "fmt"
    "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)


// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex - interface is keyword in golang
type xmlRouting__InstancesInstanceVrf__TargetImport struct {
	XMLName xml.Name `xml:"configuration"`
	Groups  struct {
		XMLName	xml.Name	`xml:"groups"`
		Name	string	`xml:"name"`
		V_instance  struct {
			XMLName xml.Name `xml:"instance"`
			V_name  *string  `xml:"name,omitempty"`
			V_vrf__target  struct {
				XMLName xml.Name `xml:"vrf-target"`
				V_import  *string  `xml:"import,omitempty"`
			} `xml:"vrf-target"`
		} `xml:"routing-instances>instance"`
	} `xml:"groups"`
	ApplyGroup string `xml:"apply-groups"`
}

// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex- interface is keyword in golang
func junosRouting__InstancesInstanceVrf__TargetImportCreate(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_import := d.Get("import").(string)


	config := xmlRouting__InstancesInstanceVrf__TargetImport{}
	config.ApplyGroup = id
	config.Groups.Name = id
	config.Groups.V_instance.V_name = &V_name
	config.Groups.V_instance.V_vrf__target.V_import = &V_import

    err = client.SendTransaction("", config, false)
    check(err)
    
    d.SetId(fmt.Sprintf("%s_%s", client.Host, id))
    
	return junosRouting__InstancesInstanceVrf__TargetImportRead(d,m)
}

func junosRouting__InstancesInstanceVrf__TargetImportRead(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	config := &xmlRouting__InstancesInstanceVrf__TargetImport{}

	err = client.MarshalGroup(id, config)
	check(err)
 	d.Set("name", config.Groups.V_instance.V_name)
	d.Set("import", config.Groups.V_instance.V_vrf__target.V_import)

	return nil
}

func junosRouting__InstancesInstanceVrf__TargetImportUpdate(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_import := d.Get("import").(string)


	config := xmlRouting__InstancesInstanceVrf__TargetImport{}
	config.ApplyGroup = id
	config.Groups.Name = id
	config.Groups.V_instance.V_name = &V_name
	config.Groups.V_instance.V_vrf__target.V_import = &V_import

    err = client.SendTransaction(id, config, false)
    check(err)
    
	return junosRouting__InstancesInstanceVrf__TargetImportRead(d,m)
}

func junosRouting__InstancesInstanceVrf__TargetImportDelete(d *schema.ResourceData, m interface{}) error {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
    _, err = client.DeleteConfigNoCommit(id)
    check(err)

    d.SetId("")
    
	return nil
}

func junosRouting__InstancesInstanceVrf__TargetImport() *schema.Resource {
	return &schema.Resource{
		Create: junosRouting__InstancesInstanceVrf__TargetImportCreate,
		Read: junosRouting__InstancesInstanceVrf__TargetImportRead,
		Update: junosRouting__InstancesInstanceVrf__TargetImportUpdate,
		Delete: junosRouting__InstancesInstanceVrf__TargetImportDelete,

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
			"import": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_instance.V_vrf__target. Target community to use when filtering on import",
			},
		},
	}
}