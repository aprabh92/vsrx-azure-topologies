##########
#  VM1
##########

# Public IP resource 
resource "azurerm_public_ip" "terraform_pip_vm1" {
    name = "tf-vm1-pip"
    location = azurerm_resource_group.terraform.location
    resource_group_name = azurerm_resource_group.terraform.name
    allocation_method = "Dynamic"
    idle_timeout_in_minutes = 30
    
    tags = {
        environment = "tf-pip-vm1"
    }
}

# NIC 1
resource "azurerm_network_interface" "tf-vm1-nic1" {
    name = "tf-vm1-nic1"
    location = azurerm_resource_group.terraform.location
    resource_group_name = azurerm_resource_group.terraform.name
    
    ip_configuration {
        name = "internal"
        subnet_id = azurerm_subnet.tf-net-1-subnet-1.id
        private_ip_address_allocation = "Dynamic"
        public_ip_address_id = azurerm_public_ip.terraform_pip_vm1.id
     }
}

# NIC2
resource "azurerm_network_interface" "tf-vm1-nic2" {
    name = "tf1-vm1-nic2"
    location = azurerm_resource_group.terraform.location
    resource_group_name = azurerm_resource_group.terraform.name
    enable_ip_forwarding = true
    enable_accelerated_networking = true
    ip_configuration {
        name = "internal"
        subnet_id = azurerm_subnet.tf-net-1-subnet-2.id
        private_ip_address_allocation = "static"
        private_ip_address = var.vm1_nic2
     }
}

# VM definition
resource "azurerm_virtual_machine" "tf-vm1" {
    name = "tf-vm1"
    location = azurerm_resource_group.terraform.location
    resource_group_name = azurerm_resource_group.terraform.name
    vm_size = "Standard_DS3_v2"  
    network_interface_ids = [
        azurerm_network_interface.tf-vm1-nic1.id,
        azurerm_network_interface.tf-vm1-nic2.id
    ]
    delete_os_disk_on_termination = true
    delete_data_disks_on_termination = true
    primary_network_interface_id = azurerm_network_interface.tf-vm1-nic1.id
    storage_os_disk {
        name = "tfvm1osdisk"
        create_option = "FromImage"
        caching = "ReadWrite"
        managed_disk_type = "Standard_LRS"
        disk_size_gb = 40
    }
    storage_image_reference {
        publisher = "canonical"
        offer = "UbuntuServer"
        sku = "18.04-LTS"
        version = "latest"
    }
    os_profile {
        computer_name = "tfvm-1"
        admin_username = var.user-name
        admin_password = var.password
        custom_data = file("./cloudconfig.txt")
    }
    os_profile_linux_config {
        disable_password_authentication = false
    }
    tags = {
        device = "test-terraform"
    }
}

##############
# VM2
##############

# Public IP resource 
resource "azurerm_public_ip" "terraform_pip_vm2" {
    name = "tf-vm2-pip"
    location = azurerm_resource_group.terraform.location
    resource_group_name = azurerm_resource_group.terraform.name
    allocation_method = "Dynamic"
    idle_timeout_in_minutes = 30
    
    tags = {
        environment = "tf-pip-vm2"
    }
}

# NIC 1
resource "azurerm_network_interface" "tf-vm2-nic1" {
    name = "tf1-vm2-nic1"
    location = azurerm_resource_group.terraform.location
    resource_group_name = azurerm_resource_group.terraform.name
    
    ip_configuration {
        name = "internal"
        subnet_id = azurerm_subnet.tf-net-1-subnet-1.id
        private_ip_address_allocation = "Dynamic"
        public_ip_address_id = azurerm_public_ip.terraform_pip_vm2.id
     }
}

# NIC 2
resource "azurerm_network_interface" "tf-vm2-nic2" {
    name = "tf1-vm2-nic2"
    location = azurerm_resource_group.terraform.location
    resource_group_name = azurerm_resource_group.terraform.name
    enable_ip_forwarding = true
    enable_accelerated_networking = true
    ip_configuration {
        name = "internal"
        subnet_id = azurerm_subnet.tf-net-1-subnet-3.id
        private_ip_address_allocation = "Static"
        private_ip_address = var.vm2_nic2
     }
}

resource "azurerm_virtual_machine" "tf-vm2" {
    name = "tf-vm2"
    location = azurerm_resource_group.terraform.location
    resource_group_name = azurerm_resource_group.terraform.name
    vm_size = "Standard_DS3_v2"  
    network_interface_ids = [
        azurerm_network_interface.tf-vm2-nic1.id,
        azurerm_network_interface.tf-vm2-nic2.id
    ]
    delete_os_disk_on_termination = true
    delete_data_disks_on_termination = true
    primary_network_interface_id = azurerm_network_interface.tf-vm2-nic1.id
    storage_os_disk {
        name = "tfvm2osdisk"
        create_option = "FromImage"
        caching = "ReadWrite"
        managed_disk_type = "Standard_LRS"
        disk_size_gb = 40
    }
    storage_image_reference {
        publisher = "canonical"
        offer = "UbuntuServer"
        sku = "18.04-LTS"
        version = "latest"
    }
    os_profile {
        computer_name = "tfvm-2"
        admin_username = var.user-name
        admin_password = var.password
        custom_data = file("./cloudconfig.txt")
    }
    os_profile_linux_config {
        disable_password_authentication = false
    }
    tags = {
        device = "test-terraform"
    }
}

#########
# vSRX
########

resource "azurerm_public_ip" "terraform_pip_vsrx" {
    name = "tf-vsrx-pip"
    location = azurerm_resource_group.terraform.location
    resource_group_name = azurerm_resource_group.terraform.name
    allocation_method = "Dynamic"
    idle_timeout_in_minutes = 30
    
    tags = {
        environment = "tf-pip-vsrx"
    }
}

resource "azurerm_network_interface" "tf-vsrx-nic1" {
    name = "tf1-vsrx-nic1"
    location = azurerm_resource_group.terraform.location
    resource_group_name = azurerm_resource_group.terraform.name
    
    ip_configuration {
        name = "internal"
        subnet_id = azurerm_subnet.tf-net-1-subnet-1.id
        private_ip_address_allocation = "Dynamic"
        public_ip_address_id = azurerm_public_ip.terraform_pip_vsrx.id
     }
}

resource "azurerm_network_interface" "tf-vsrx-nic2" {
    name = "tf1-vsrx-nic2"
    location = azurerm_resource_group.terraform.location
    resource_group_name = azurerm_resource_group.terraform.name
    enable_ip_forwarding = true
    enable_accelerated_networking = true
    ip_configuration {
        name = "internal"
        subnet_id = azurerm_subnet.tf-net-1-subnet-2.id
        private_ip_address_allocation = "Static"
        private_ip_address = var.vsrx_nic2
     }
}

resource "azurerm_network_interface" "tf-vsrx-nic3" {
    name = "tf1-vsrx-nic3"
    location = azurerm_resource_group.terraform.location
    resource_group_name = azurerm_resource_group.terraform.name
    enable_ip_forwarding = true
    enable_accelerated_networking = true
    ip_configuration {
        name = "internal"
        subnet_id = azurerm_subnet.tf-net-1-subnet-3.id
        private_ip_address_allocation = "Static"
        private_ip_address = var.vsrx_nic3
     }
}

resource "azurerm_marketplace_agreement" "srx" {
  publisher = "juniper-networks"
  offer     = "vsrx-next-generation-firewall"
  plan      = "vsrx-byol-azure-image"
}

resource "azurerm_virtual_machine" "fw-vm" {
  depends_on                      = [azurerm_marketplace_agreement.srx]
  name                            = "vsrx"
  location                        = azurerm_resource_group.terraform.location
  resource_group_name             = azurerm_resource_group.terraform.name
  network_interface_ids           = [ azurerm_network_interface.tf-vsrx-nic1.id,
                                      azurerm_network_interface.tf-vsrx-nic2.id,
                                      azurerm_network_interface.tf-vsrx-nic3.id,
                                    ]
  primary_network_interface_id = azurerm_network_interface.tf-vsrx-nic1.id
  vm_size = "Standard_DS3_v2"
  plan {
    name      = "vsrx-azure-image-byol"
    product   = "vsrx-next-generation-firewall-payg"
    publisher = "juniper-networks"
  }
  storage_image_reference {
    publisher = "juniper-networks"
    offer     = "vsrx-next-generation-firewall-payg"
    sku       = "vsrx-azure-image-byol"
    version   = "20.4.2"
  }
  storage_os_disk {
    os_type           = "Linux"
    name              = "vsrx-osdisk"
    create_option     = "FromImage"
    caching           = "ReadWrite"
    managed_disk_type = "StandardSSD_LRS"
    disk_size_gb      = 40
  }
  os_profile {
    computer_name      = "vsrx-jnpr"
    admin_username     = var.user-name
    admin_password     = var.password
    custom_data        = file("./cloudvsrx.txt")
  }
  os_profile_linux_config {
    disable_password_authentication = false
  }
  tags = {
    device = "vsrx"
  }
}
