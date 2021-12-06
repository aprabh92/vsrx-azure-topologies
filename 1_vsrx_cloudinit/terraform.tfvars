# ----------------------------------------------------------- #
# Plan information Gathered From ARM template. 
# ------------------------------------------------------------#
#
#       "plan-byol": {
#         "name": "vsrx-byol-azure-image",
#         "publisher": "juniper-networks",
#         "product": "vsrx-next-generation-firewall"
#       },
#       "plan-payg-b1": {
#         "name": "vsrx-azure-image-payg-b1",
#         "publisher": "juniper-networks",
#         "product": "vsrx-next-generation-firewall-payg"
#       },
#       "image-byol": {
#         "publisher": "juniper-networks",
#         "offer": "vsrx-next-generation-firewall",
#         "sku": "vsrx-byol-azure-image",
#         "version": "latest"
#       },
#       "image-payg-b1": {
#         "publisher": "juniper-networks",
#         "offer": "vsrx-next-generation-firewall-payg",
#         "sku": "vsrx-azure-image-payg-b1",
#         "version": "latest"
#       }
#
# -------------------------------------------------------- #

# -------------------------------------- #
# Variables for the terraform .tf file
# ------------------------------------- #
zone = "EastUs"
rg-name = "test_cloudinit"
user-name = "jnpr"
password = "Juniper@1234567"

## vSRX specifics 
product = "vsrx-next-generation-firewall-payg"
offer = "vsrx-next-generation-firewall-payg" 
plan = "vsrx-azure-image-byol"
sku = "vsrx-azure-image-byol"

#offer = "vsrx-next-generation-firewall-payg"
#plan = "vsrx-azure-image-payg-b1" 
#sku = "vsrx-azure-image-payg-b1"
#product = "vsrx-next-generation-firewall-payg"

publisher = "juniper-networks"
vsrx-version = "20.4.2"

## subnet info
network = ["172.18.0.0/16"]
net1 = ["172.18.1.0/24"]
net2 = ["172.18.2.0/24"]
net3 = ["172.18.3.0/24"]

