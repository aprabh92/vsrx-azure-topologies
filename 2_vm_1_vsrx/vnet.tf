resource "azurerm_virtual_network" "terraform-vnet" {
    name = "tf-net-1"
    resource_group_name = azurerm_resource_group.terraform.name
    location = azurerm_resource_group.terraform.location
    address_space = ["172.16.0.0/16"]
} 

resource "azurerm_subnet" "tf-net-1-subnet-1" {
    name = "tf-net-1-subnet-1"
    resource_group_name = azurerm_resource_group.terraform.name
    virtual_network_name = azurerm_virtual_network.terraform-vnet.name
    address_prefixes = ["172.16.1.0/24"]
}
resource "azurerm_subnet" "tf-net-1-subnet-2" {
    name = "tf-net-1-subnet-left"
    address_prefixes = ["172.16.2.0/24"]
    resource_group_name = azurerm_resource_group.terraform.name
    virtual_network_name = azurerm_virtual_network.terraform-vnet.name
}
resource "azurerm_subnet" "tf-net-1-subnet-3"{
    name = "tf-net-2-subnet-right"
    address_prefixes = ["172.16.3.0/24"]
    resource_group_name = azurerm_resource_group.terraform.name
    virtual_network_name = azurerm_virtual_network.terraform-vnet.name
}

