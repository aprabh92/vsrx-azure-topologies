zone = "EastUs"
rg-name = "terraform"
user-name = "jnpr"
password = "Juniper@1234567"

## vSRX specifics 
offer = "vsrx-next-generation-firewall" 
plan = "vsrx-byol-azure-image"
VsrxSku = "vsrx-byol-azure-image"

## NIC Addresses 
vm1_nic2 = "172.16.2.5"
vm2_nic2 = "172.16.3.5"
vsrx_nic2 = "172.16.2.4"
vsrx_nic3 = "172.16.3.4"

## Az route
prefix1 = "172.16.3.5/32"
nhroute1 = "172.16.2.4"

prefix2 = "172.16.2.5/32"
nhroute2 = "172.16.3.4"
