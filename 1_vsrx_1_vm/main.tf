######################
# Variable definitions
######################

variable "zone" {
            description = "azure region"
            type = string
            default = "EastUs"
            }

variable "rg-name" {}
variable "user-name" {}
variable "password" {}
variable "offer" {}
variable "plan" {}
variable "sku" {}

variable "network" {}
variable "net1" {
            description = "subnet 1 (mgmt subnet)" 
            type = list
            }

variable "net2" {
            description = "subnet 2" 
            type = list
            }

variable "net3" {
            description = "subnet 3" 
            type = list
            }

#####################
# Provider definition
#####################

terraform {
    required_providers {
        azurerm = {
            source = "hashicorp/azurerm"
            version = "=2.33.0"
        }
    }
}

provider "azurerm" {
    features {}
}

resource "azurerm_resource_group" "vsrx" {
    # defining resource group name 
    name = var.rg-name
    location = var.zone
}


#########
# VNET
########

resource "azurerm_virtual_network" "terraform-vnet" {
    name = "tf-net-1"
    resource_group_name = azurerm_resource_group.vsrx.name
    location = azurerm_resource_group.vsrx.location
    address_space = var.network
} 

resource "azurerm_subnet" "tf-net-1-subnet-1" {
    name = "tf-net-1-subnet-1"
    resource_group_name = azurerm_resource_group.vsrx.name
    virtual_network_name = azurerm_virtual_network.terraform-vnet.name
    address_prefixes = var.net1
}
resource "azurerm_subnet" "tf-net-1-subnet-2" {
    name = "tf-net-1-subnet-left"
    address_prefixes = var.net2
    resource_group_name = azurerm_resource_group.vsrx.name
    virtual_network_name = azurerm_virtual_network.terraform-vnet.name
}
resource "azurerm_subnet" "tf-net-1-subnet-3"{
    name = "tf-net-2-subnet-right"
    address_prefixes = var.net3
    resource_group_name = azurerm_resource_group.vsrx.name
    virtual_network_name = azurerm_virtual_network.terraform-vnet.name
}

#########
# vSRX
#########

resource "azurerm_public_ip" "tf_pip_vsrx" {
    name = "tf_vsrx_pip"
    location = azurerm_resource_group.vsrx.location
    resource_group_name = azurerm_resource_group.vsrx.name
    allocation_method = "Dynamic"
    idle_timeout_in_minutes = 30
    
    tags = {
        environment = "tf_pip_vsrx"
    }
}

resource "azurerm_network_interface" "tf_vsrx_nic1" {
    name = "tf1_vsrx_nic1"
    location = azurerm_resource_group.vsrx.location
    resource_group_name = azurerm_resource_group.vsrx.name
    
    ip_configuration {
        name = "internal"
        subnet_id = azurerm_subnet.tf-net-1-subnet-1.id
        private_ip_address_allocation = "Dynamic"
        public_ip_address_id = azurerm_public_ip.tf_pip_vsrx.id
     }
}

resource "azurerm_network_interface" "tf_vsrx_nic2" {
    name = "tf1_vsrx_nic2"
    location = azurerm_resource_group.vsrx.location
    resource_group_name = azurerm_resource_group.vsrx.name
    enable_ip_forwarding = true
    enable_accelerated_networking = true
    ip_configuration {
        name = "internal"
        subnet_id = azurerm_subnet.tf-net-1-subnet-2.id
        private_ip_address_allocation = "Dynamic"
        #private_ip_address = var.vsrx_nic2
     }
}

resource "azurerm_network_interface" "tf_vsrx_nic3" {
    name = "tf1_vsrx_nic3"
    location = azurerm_resource_group.vsrx.location
    resource_group_name = azurerm_resource_group.vsrx.name
    enable_ip_forwarding = true
    enable_accelerated_networking = true
    ip_configuration {
        name = "internal"
        subnet_id = azurerm_subnet.tf-net-1-subnet-3.id
        private_ip_address_allocation = "Dynamic"
        #private_ip_address = var.vsrx_nic3
     }
}

resource "azurerm_marketplace_agreement" "vsrx-tf" {
  publisher = "juniper-networks"
  offer     = var.offer
  plan      = var.plan
}

resource "azurerm_virtual_machine" "tfvsrx1" {
  depends_on                      = [azurerm_marketplace_agreement.vsrx-tf]
  name                            = "tfvsrx1"
  location                        = azurerm_resource_group.vsrx.location
  resource_group_name             = azurerm_resource_group.vsrx.name
  network_interface_ids           = [ azurerm_network_interface.tf_vsrx_nic1.id,
                                      azurerm_network_interface.tf_vsrx_nic2.id,
                                      azurerm_network_interface.tf_vsrx_nic3.id,
                                    ]
  primary_network_interface_id = azurerm_network_interface.tf_vsrx_nic1.id
  vm_size = "Standard_DS3_v2"
  delete_os_disk_on_termination = true
  delete_data_disks_on_termination = true
  plan {
    name      = "vsrx-byol-azure-image"
    product   = "vsrx-next-generation-firewall"
    publisher = "juniper-networks"
  }
  storage_image_reference {
    publisher = "juniper-networks"
    offer     = var.offer
    sku       = var.sku
    version   = "latest"
  }
  storage_os_disk {
    os_type           = "Linux"
    name              = "tfvsrxosdisk"
    create_option     = "FromImage"
    caching           = "ReadWrite"
    managed_disk_type = "StandardSSD_LRS"
    disk_size_gb      = 40
  }
  os_profile {
    computer_name      = "vsrxTest"
    admin_username     = var.user-name
    admin_password     = var.password
  }
  os_profile_linux_config {
    disable_password_authentication = false
  }
  tags = {
    device = "tf_vsrx"
  }
}


##########
#  VM1
##########

#-------------- Public IP resource -------------- #
resource "azurerm_public_ip" "tf_pip_vm1" {
    name = "tf_vm1_pip"
    location = azurerm_resource_group.vsrx.location
    resource_group_name = azurerm_resource_group.vsrx.name
    allocation_method = "Dynamic"
    idle_timeout_in_minutes = 30
    
    tags = {
        environment = "tf_pip_vm1"
    }
}

# ------------- NIC 1 ---------------------- #
resource "azurerm_network_interface" "tf_vm1_nic1" {
    name = "tf_vm1_nic1"
    location = azurerm_resource_group.vsrx.location
    resource_group_name = azurerm_resource_group.vsrx.name
    
    ip_configuration {
        name = "internal"
        subnet_id = azurerm_subnet.tf-net-1-subnet-1.id
        private_ip_address_allocation = "Dynamic"
        public_ip_address_id = azurerm_public_ip.tf_pip_vm1.id
     }
}

# --------------  NIC2   ------------- #
resource "azurerm_network_interface" "tf_vm1_nic2" {
    name = "tf1_vm1_nic2"
    location = azurerm_resource_group.vsrx.location
    resource_group_name = azurerm_resource_group.vsrx.name
    enable_ip_forwarding = true
    enable_accelerated_networking = true
    ip_configuration {
        name = "internal"
        subnet_id = azurerm_subnet.tf-net-1-subnet-2.id
        private_ip_address_allocation = "Dynamic"
        #private_ip_address = var.vm1_nic2
     }
}

#---------  VM definition ---------------- #
resource "azurerm_virtual_machine" "tfvm1" {
  name                            = "tfvm1"
  location                        = azurerm_resource_group.vsrx.location
  resource_group_name             = azurerm_resource_group.vsrx.name
  network_interface_ids           = [ azurerm_network_interface.tf_vm1_nic1.id,
                                      azurerm_network_interface.tf_vm1_nic2.id,
                                    ]
  delete_os_disk_on_termination = true
  delete_data_disks_on_termination = true
  primary_network_interface_id = azurerm_network_interface.tf_vm1_nic1.id
  vm_size = "Standard_DS3_v2"

  storage_image_reference {
    publisher = "canonical"
    offer     = "UbuntuServer"
    sku       = "18.04-LTS"
    version   = "latest"
  }

  storage_os_disk {
    name              = "tfvmosdisk"
    create_option     = "FromImage"
    caching           = "ReadWrite"
    managed_disk_type = "StandardSSD_LRS"
    disk_size_gb      = 40
  }
  os_profile {
    computer_name      = "vmtest-1"
    admin_username     = var.user-name
    admin_password     = var.password
  }
  os_profile_linux_config {
    disable_password_authentication = false
  }
  tags = {
    device = "test"
  }
}
