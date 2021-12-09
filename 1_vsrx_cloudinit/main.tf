# --------------------- #
# Variable definitions
# -------------------- #
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
variable "publisher" {
            default = "juniper-networks"
         }
variable "vsrx-version" {
            default = "latest"
          }
variable "product" {}
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

# ------------------- # 
# Provider definition
# ------------------- #

terraform {
    required_providers {
        azurerm = {
            source = "hashicorp/azurerm"
            version = "=2.33.0"
        }
        cloudinit = {
            version = "2.0.0"
        }
    }
}

provider "azurerm" {
    features {}
}

provider "cloudinit" {
}

resource "azurerm_resource_group" "vsrx1" {
    # defining resource group name 
    name = var.rg-name
    location = var.zone
}


# -------- #
# VNET
# -------- #

resource "azurerm_virtual_network" "terraform1-vnet" {
    name = "tfc-net-1"
    resource_group_name = azurerm_resource_group.vsrx1.name
    location = azurerm_resource_group.vsrx1.location
    address_space = var.network
} 

resource "azurerm_subnet" "tfc-net-1-subnet-1" {
    name = "tfc-net-1-subnet-1"
    resource_group_name = azurerm_resource_group.vsrx1.name
    virtual_network_name = azurerm_virtual_network.terraform1-vnet.name
    address_prefixes = var.net1
}
resource "azurerm_subnet" "tfc-net-1-subnet-2" {
    name = "tfc-net-1-subnet-left"
    address_prefixes = var.net2
    resource_group_name = azurerm_resource_group.vsrx1.name
    virtual_network_name = azurerm_virtual_network.terraform1-vnet.name
}
resource "azurerm_subnet" "tfc-net-1-subnet-3"{
    name = "tf-net-2-subnet-right"
    address_prefixes = var.net3
    resource_group_name = azurerm_resource_group.vsrx1.name
    virtual_network_name = azurerm_virtual_network.terraform1-vnet.name
}

#-------------------------------------------------------------- #
# Storage Account to enable serial console and boot diagnostics
#-------------------------------------------------------------- #

resource "random_id" "randomId2" {
    keepers = {
        # Generate a new ID only when a new resource group is defined
        resource_group = azurerm_resource_group.vsrx1.name
    }
    byte_length = 8
}

resource "azurerm_storage_account" "mystorageaccount" {
    name                        = "diag${random_id.randomId2.hex}"
    resource_group_name         = azurerm_resource_group.vsrx1.name
    location                    = "eastus"
    account_replication_type    = "LRS"
    account_tier                = "Standard"

    tags = {
        environment = "Terraform Demo"
    }
}


# ------- #
# vSRX
# ------- #

resource "azurerm_public_ip" "tf_pip_vsrx1" {
    name = "tf_vsrx1_pip"
    location = azurerm_resource_group.vsrx1.location
    resource_group_name = azurerm_resource_group.vsrx1.name
    allocation_method = "Dynamic"
    idle_timeout_in_minutes = 30
    
    tags = {
        environment = "tf_pip_vsrx1"
    }
}

resource "azurerm_network_interface" "tf_vsrx1_nic1" {
    name = "tf1_vsrx1_nic1"
    location = azurerm_resource_group.vsrx1.location
    resource_group_name = azurerm_resource_group.vsrx1.name
    
    ip_configuration {
        name = "internal"
        subnet_id = azurerm_subnet.tfc-net-1-subnet-1.id
        private_ip_address_allocation = "Dynamic"
        public_ip_address_id = azurerm_public_ip.tf_pip_vsrx1.id
     }
}

resource "azurerm_network_interface" "tf_vsrx1_nic2" {
    name = "tf1_vsrx1_nic2"
    location = azurerm_resource_group.vsrx1.location
    resource_group_name = azurerm_resource_group.vsrx1.name
    enable_ip_forwarding = true
    enable_accelerated_networking = true
    ip_configuration {
        name = "internal"
        subnet_id = azurerm_subnet.tfc-net-1-subnet-2.id
        private_ip_address_allocation = "Dynamic"
        #private_ip_address = var.vsrx1_nic2
     }
}

resource "azurerm_network_interface" "tf_vsrx1_nic3" {
    name = "tf1_vsrx1_nic3"
    location = azurerm_resource_group.vsrx1.location
    resource_group_name = azurerm_resource_group.vsrx1.name
    enable_ip_forwarding = true
    enable_accelerated_networking = true
    ip_configuration {
        name = "internal"
        subnet_id = azurerm_subnet.tfc-net-1-subnet-3.id
        private_ip_address_allocation = "Dynamic"
        #private_ip_address = var.vsrx1_nic3
     }
}

resource "azurerm_marketplace_agreement" "vsrx1-tf-cloudinit" {
  publisher = var.publisher
  offer     = var.offer
  plan      = var.sku
}

resource "azurerm_virtual_machine" "tfvsrx11" {
  depends_on                      = [azurerm_marketplace_agreement.vsrx1-tf-cloudinit]
  name                            = "tfvsrx11"
  location                        = azurerm_resource_group.vsrx1.location
  resource_group_name             = azurerm_resource_group.vsrx1.name
  network_interface_ids           = [ azurerm_network_interface.tf_vsrx1_nic1.id,
                                      azurerm_network_interface.tf_vsrx1_nic2.id,
                                      azurerm_network_interface.tf_vsrx1_nic3.id,
                                    ]
  primary_network_interface_id = azurerm_network_interface.tf_vsrx1_nic1.id
  vm_size = "Standard_DS3_v2"
  delete_os_disk_on_termination = true
  delete_data_disks_on_termination = true

  plan {
    name      = var.sku
    product   = var.product
    publisher = var.publisher
  }
  storage_image_reference {
    publisher = var.publisher
    offer     = var.offer
    sku       = var.sku
    version   = var.vsrx-version
  }

  storage_os_disk {
    os_type           = "Linux"
    name              = "tfvsrx1osdisk"
    create_option     = "FromImage"
    caching           = "ReadWrite"
    managed_disk_type = "StandardSSD_LRS"
    disk_size_gb      = 40
  }
  # template for cloud init. This value would be passed externally using templatefile 
  os_profile {
    computer_name      = "vsrx1Test"
    admin_username     = var.user-name
    admin_password     = var.password
    custom_data        = file("./cloudconfig.txt")
    #custom_data        = templatefile("./cloudconfig.txt", { ge_0 = azurerm_network_interface.tf_vsrx1_nic2.private_ip_address,
    #                                                         ge_1 = azurerm_network_interface.tf_vsrx1_nic3.private_ip_address
    #                                                        })
  }

  os_profile_linux_config {
    disable_password_authentication = false
  }

  boot_diagnostics {
    enabled = true
    storage_uri = azurerm_storage_account.mystorageaccount.primary_blob_endpoint
  }

  tags = {
    device = "tf_vsrx1"
  }
}
