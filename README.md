# vsrx-topologies-tf
This repository contains terraform scripts to bring up vsrx topologies on azure.

contains 3 directories 

1. 2 VMs and 1 vSRX 
```
                               +---------+
       +------------+          |         |          +-------------+
       |            |.5   .2.4 |         |.3.4   .5 |             |
       |   vm1      +----------+  vsrx   +----------+     vm2     |
       |            |          |         |          |             |
       +------------+          |         |          +-------------+
                               +---------+
       
```

2. 1_vsrx_1_vm

```
                               +---------+
       +------------+          |         |  
       |            |          |         |
       |   vm1      +----------+  vsrx   +-
       |            |          |         |  
       +------------+          |         |  
                               +---------+
```

3. 1_vsrx_cloudinit

This is to test cloud init . Brings up a vSRX with custom_data flag enabled to pass configs into vsrx during bootup.

                  
