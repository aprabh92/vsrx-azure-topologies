resource "azurerm_route_table" "az-rt-1" {
  name                          = "az-rt-1"
  location                      = azurerm_resource_group.terraform.location
  resource_group_name           = azurerm_resource_group.terraform.name
  disable_bgp_route_propagation = true

  route {
    name           = "route1"
    address_prefix = var.prefix1
    next_hop_type  = "VirtualAppliance"
    next_hop_in_ip_address = var.nhroute1
  }
}


resource "azurerm_route_table" "az-rt-2" {
  name                          = "az-rt-2"
  location                      = azurerm_resource_group.terraform.location
  resource_group_name           = azurerm_resource_group.terraform.name
  disable_bgp_route_propagation = true

  route {
    name           = "route2"
    address_prefix = var.prefix2
    next_hop_type  = "VirtualAppliance"
    next_hop_in_ip_address = var.nhroute2
  }
}

resource "azurerm_subnet_route_table_association" "route1-sa" {
  subnet_id      = azurerm_subnet.tf-net-1-subnet-2.id
  route_table_id = azurerm_route_table.az-rt-1.id
}

resource "azurerm_subnet_route_table_association" "route2-sa" {
  subnet_id      = azurerm_subnet.tf-net-1-subnet-3.id
  route_table_id = azurerm_route_table.az-rt-2.id
}
