# Configuring vSRX on Azure using JTAF

This repository contains examples of configuring vSRX using JTAF. For information on JTAF and how to build the juniper terraform provider, please visit [here](https://github.com/Juniper/junos-terraform.git)

Below configurations examples are covered in this directory

- configuring interfaces and placing them into respective zones 
- configuring security policies for L3 and L4 stateful firewall 
- configuring IPSEC tunnels 
    - creating the traffic selector for the VPN
- configure L3VPNs and placing the IPSEC tunnel interface into VRF
- configuring source and destination nat
- configuring policies 
    - from match conditions and then actions 
    - match conditions for community 

## TO try using 
1. generate the provider based on the config.toml and xpath_srx.xml file added in the directory
2. edit the terraform.tfvars file accordingly to match your setup
3. follow the below terraform steps
    - terraform init
    - terraform plan
    - terraform apply --auto-approve

Note: edit the resources accordingly to suit your env and configurations. This directoy only provides a sample config.
