# ---------- #
# Policy-1
# ----------- #

# Match source address
resource "junos-vsrx_SecurityPoliciesPolicyPolicyMatchSource__Address" "SecurityPoliciesPolicyPolicyMatchSource__Address_1" {
	resource_name = "JTAF-SECURITY_POLICIES_MATCH1"
    	from__zone__name = "untrust"
    	to__zone__name = "trust"
    	name = "policy-1"
    	source__address = "source-1"
}

# Match Destination address
resource "junos-vsrx_SecurityPoliciesPolicyPolicyMatchDestination__Address" "SecurityPoliciesPolicyPolicyMatchDestination__Address_1" {
	    resource_name = "JTAF-SECURITY_POLICIES_DEST1"
    	from__zone__name = "untrust"
    	to__zone__name = "trust"
    	name = "policy-1"
    	destination__address = "dest-1"
}

# Match application
resource "junos-vsrx_SecurityPoliciesPolicyPolicyMatchApplication" "SecurityPoliciesPolicyPolicyMatchApplication_1" {
    	resource_name = "JTAF-SECURITY_POLICIES_APP1"
    	from__zone__name = "untrust"
   	    to__zone__name = "trust"
    	name = "policy-1"
    	application = "TCP-26"
	    depends_on = [
		    junos-vsrx_ApplicationsApplicationProtocol.TCP-26
	    ]
}

# Then permit 
resource "junos-vsrx_SecurityPoliciesPolicyPolicyThenPermit" "SecurityPoliciesPolicyPolicyThenPermit__Policy_9" {
    	resource_name = "JTAF-SECURITY_POLICIES_PERMIT1"
    	from__zone__name = "untrust"
    	to__zone__name = "trust"
    	name = "policy-1"
}

# Then Log session
resource "junos-vsrx_SecurityPoliciesPolicyPolicyThenLogSession__Init" "SecurityPoliciesPolicyPolicyThenLogSession__Init_1" {
    	resource_name = "JTAF-SECURITY_POLICIES_LOG1"
    	from__zone__name = "untrust"
    	to__zone__name = "trust"
    	name = "policy-1"
    	session__init = " "
}


# Then count
resource "junos-vsrx_SecurityPoliciesPolicyPolicyThenCount" "SecurityPoliciesPolicyPolicyThenCount" {
    	resource_name = "JTAF-SECURITY_POLICIES_COUNT1"
    	from__zone__name = "untrust"
    	to__zone__name = "trust"
    	name = "policy-1"
}

# ---------- #
# Policy-2
# ---------- #

# Match source address
resource "junos-vsrx_SecurityPoliciesPolicyPolicyMatchSource__Address" "SecurityPoliciesPolicyPolicyMatchSource__Address_2" {
	resource_name = "JTAF-SECURITY_POLICIES_SOURCE2"
    	from__zone__name = "untrust"
    	to__zone__name = "trust"
    	name = "policy-2"
    	source__address = "source-2"
}

# Match Destination address
resource "junos-vsrx_SecurityPoliciesPolicyPolicyMatchDestination__Address" "SecurityPoliciesPolicyPolicyMatchDestination__Address_2" {
	resource_name = "JTAF-SECURITY_POLICIES_DEST2"
    	from__zone__name = "untrust"
    	to__zone__name = "trust"
    	name = "policy-2"
    	destination__address = "dest-2"
}

# Match application
resource "junos-vsrx_SecurityPoliciesPolicyPolicyMatchApplication" "SecurityPoliciesPolicyPolicyMatchApplication_2" {
    	resource_name = "JTAF-SECURITY_POLICIES_APP2"
    	from__zone__name = "untrust"
   	    to__zone__name = "trust"
    	name = "policy-2"
    	application = "TCP-27"
	depends_on = [
		junos-vsrx_ApplicationsApplicationProtocol.TCP-27
	]

}
# Then deny
resource "junos-vsrx_SecurityPoliciesPolicyPolicyThenDeny" "SecurityPoliciesPolicyPolicyThenDeny_2" {
    resource_name = "JTAF-SECURITY_POLICIES_DENY2"
    from__zone__name = "untrust"
    to__zone__name = "trust"
    name = "policy-2"
}
