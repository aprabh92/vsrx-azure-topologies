resource "junos-vsrx_Policy__OptionsCommunityMembers" "community-red" {
    resource_name ="JTAF-COMM-MEMBER-1"
    name = "RED"
    members = "65001:100"
}

resource "junos-vsrx_Policy__OptionsCommunityMembers" "community-blue" {    
    resource_name ="JTAF-COMM-MEMBER-2"
    name = "BLUE"
    members = "65001:200"
}
