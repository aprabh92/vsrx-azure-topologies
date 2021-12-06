zone = "EastUs"
rg-name = "test_tf"
user-name = "jnpr"
password = "Juniper@1234567"

## vSRX specifics 
offer = "vsrx-next-generation-firewall" 
plan = "vsrx-byol-azure-image"
sku = "vsrx-byol-azure-image"

## subnet info
network = ["172.17.0.0/16"]
net1 = ["172.17.1.0/24"]
net2 = ["172.17.2.0/24"]
net3 = ["172.17.3.0/24"]
