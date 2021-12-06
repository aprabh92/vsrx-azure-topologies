#                               +---------+
#       +------------+          |         |          +-------------+
#       |            |.5   .2.4 |         |.3.4   .5 |             |
#       |   vm1      +----------+  vsrx   +----------+     vm2     |
#       |            |          |         |          |             |
#       +------------+          |         |          +-------------+
#                               +---------+
#       

# Main Terraform file to declare variable and Azure resource 
# The resource group definition is placed under vnet.tf

variable "zone" {
            default = "EastUs"
            }

variable "rg-name" {}
variable "user-name" {}
variable "password" {}
variable "offer" {}
variable "plan" {}
variable "VsrxSku" {}

variable "prefix1" {}
variable "prefix2" {}
variable "nhroute1" {}
variable "nhroute2" {}

variable "vm1_nic2" {}
variable "vm2_nic2" {}
variable "vsrx_nic2" {}
variable "vsrx_nic3" {}

provider "azurerm" {
    version = "=2.33.0"
    features {}
}

resource "azurerm_resource_group" "terraform" {
    # defining resource group name 
    name = var.rg-name
    location = var.zone
}
