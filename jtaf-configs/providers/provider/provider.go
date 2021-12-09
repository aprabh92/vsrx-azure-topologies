
// Copyright (c) 2017-2021, Juniper Networks Inc. All rights reserved.
//
// License: Apache 2.0
//
// THIS SOFTWARE IS PROVIDED BY Juniper Networks, Inc. ''AS IS'' AND ANY
// EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
// WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL Juniper Networks, Inc. BE LIABLE FOR ANY
// DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
// (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
// LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
// ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
// SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
//

package main

import (

	gonetconf "github.com/davedotdev/go-netconf/helpers/junos_helpers"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"os"
)

// ProviderConfig is to hold client information
type ProviderConfig struct {
	*gonetconf.GoNCClient
	Host string
}

func check(err error) {
	if err != nil {
		// Some of these errors will be "normal".
		f, _ := os.OpenFile("jtaf_logging.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		f.WriteString(err.Error() + "\n")
		f.Close()
		return
	}
}


func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		Host:     d.Get("host").(string),
		Port:     d.Get("port").(int),
		Username: d.Get("username").(string),
		Password: d.Get("password").(string),
		SSHKey:   d.Get("sshkey").(string),
	}

	client, err := config.Client()
	if err != nil {
		return nil, err
	}

	return &ProviderConfig{client, config.Host}, nil
}

// Provider returns a Terraform ResourceProvider.
func Provider() *schema.Provider {
	return &schema.Provider{

		Schema: map[string]*schema.Schema{
			"host": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},

			"username": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"password": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"sshkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"junos-vsrx_Policy__OptionsPolicy__StatementTermFromInstance": junosPolicy__OptionsPolicy__StatementTermFromInstance(),
			"junos-vsrx_Policy__OptionsPolicy__StatementTermFromRoute__FilterAddress": junosPolicy__OptionsPolicy__StatementTermFromRoute__FilterAddress(),
			"junos-vsrx_Policy__OptionsPolicy__StatementTermThenAccept": junosPolicy__OptionsPolicy__StatementTermThenAccept(),
			"junos-vsrx_Policy__OptionsPolicy__StatementTermThenReject": junosPolicy__OptionsPolicy__StatementTermThenReject(),
			"junos-vsrx_ApplicationsApplicationProtocol": junosApplicationsApplicationProtocol(),
			"junos-vsrx_ApplicationsApplicationSource__Port": junosApplicationsApplicationSource__Port(),
			"junos-vsrx_ApplicationsApplicationDestination__Port": junosApplicationsApplicationDestination__Port(),
			"junos-vsrx_InterfacesInterfaceDescription": junosInterfacesInterfaceDescription(),
			"junos-vsrx_InterfacesInterfaceUnitFamilyInetAddressName": junosInterfacesInterfaceUnitFamilyInetAddressName(),
			"junos-vsrx_Forwarding__OptionsSamplingInputRate": junosForwarding__OptionsSamplingInputRate(),
			"junos-vsrx_Forwarding__OptionsSamplingFamilyInetOutputFileFilename": junosForwarding__OptionsSamplingFamilyInetOutputFileFilename(),
			"junos-vsrx_FirewallFilterTermFromProtocol": junosFirewallFilterTermFromProtocol(),
			"junos-vsrx_FirewallFilterTermThenSample": junosFirewallFilterTermThenSample(),
			"junos-vsrx_FirewallFilterTermThenAccept": junosFirewallFilterTermThenAccept(),
			"junos-vsrx_Routing__InstancesInstanceInstance__Type": junosRouting__InstancesInstanceInstance__Type(),
			"junos-vsrx_Routing__InstancesInstanceRouting__OptionsStaticRouteNext__Hop": junosRouting__InstancesInstanceRouting__OptionsStaticRouteNext__Hop(),
			"junos-vsrx_Routing__InstancesInstanceRouting__OptionsStaticRouteNext__Table": junosRouting__InstancesInstanceRouting__OptionsStaticRouteNext__Table(),
			"junos-vsrx_Routing__InstancesInstanceRouting__OptionsInstance__Import": junosRouting__InstancesInstanceRouting__OptionsInstance__Import(),
			"junos-vsrx_SecurityAddress__BookAddressIp__Prefix": junosSecurityAddress__BookAddressIp__Prefix(),
			"junos-vsrx_SecurityAddress__BookAddressRange__AddressToRange__High": junosSecurityAddress__BookAddressRange__AddressToRange__High(),
			"junos-vsrx_SecurityAddress__BookAddress__SetAddressName": junosSecurityAddress__BookAddress__SetAddressName(),
			"junos-vsrx_SecurityNatSourcePoolAddressToIpaddr": junosSecurityNatSourcePoolAddressToIpaddr(),
			"junos-vsrx_SecurityNatSourceRule__SetFromZone": junosSecurityNatSourceRule__SetFromZone(),
			"junos-vsrx_SecurityNatSourceRule__SetToZone": junosSecurityNatSourceRule__SetToZone(),
			"junos-vsrx_SecurityNatSourceRule__SetRuleSrc__Nat__Rule__MatchSource__Address": junosSecurityNatSourceRule__SetRuleSrc__Nat__Rule__MatchSource__Address(),
			"junos-vsrx_SecurityNatSourceRule__SetRuleThenSource__NatPoolPool__Name": junosSecurityNatSourceRule__SetRuleThenSource__NatPoolPool__Name(),
			"junos-vsrx_SecurityNatDestinationPoolRouting__InstanceRi__Name": junosSecurityNatDestinationPoolRouting__InstanceRi__Name(),
			"junos-vsrx_SecurityNatDestinationPoolAddressIpaddr": junosSecurityNatDestinationPoolAddressIpaddr(),
			"junos-vsrx_SecurityNatDestinationPoolAddressPort": junosSecurityNatDestinationPoolAddressPort(),
			"junos-vsrx_SecurityNatDestinationRule__SetFromInterface": junosSecurityNatDestinationRule__SetFromInterface(),
			"junos-vsrx_SecurityNatDestinationRule__SetRuleDest__Nat__Rule__MatchDestination__AddressDst__Addr": junosSecurityNatDestinationRule__SetRuleDest__Nat__Rule__MatchDestination__AddressDst__Addr(),
			"junos-vsrx_SecurityNatDestinationRule__SetRuleDest__Nat__Rule__MatchDestination__PortName": junosSecurityNatDestinationRule__SetRuleDest__Nat__Rule__MatchDestination__PortName(),
			"junos-vsrx_SecurityNatDestinationRule__SetRuleThenDestination__NatPoolPool__Name": junosSecurityNatDestinationRule__SetRuleThenDestination__NatPoolPool__Name(),
			"junos-vsrx_SecurityNatProxy__ArpInterfaceAddressToIpaddr": junosSecurityNatProxy__ArpInterfaceAddressToIpaddr(),
			"junos-vsrx_SecurityPoliciesPolicyPolicyMatchSource__Address": junosSecurityPoliciesPolicyPolicyMatchSource__Address(),
			"junos-vsrx_SecurityPoliciesPolicyPolicyMatchDestination__Address": junosSecurityPoliciesPolicyPolicyMatchDestination__Address(),
			"junos-vsrx_SecurityPoliciesPolicyPolicyMatchApplication": junosSecurityPoliciesPolicyPolicyMatchApplication(),
			"junos-vsrx_SecurityPoliciesPolicyPolicyThenPermit": junosSecurityPoliciesPolicyPolicyThenPermit(),
			"junos-vsrx_SecurityPoliciesPolicyPolicyThenLogSession__Init": junosSecurityPoliciesPolicyPolicyThenLogSession__Init(),
			"junos-vsrx_SecurityPoliciesPolicyPolicyThenLogSession__Close": junosSecurityPoliciesPolicyPolicyThenLogSession__Close(),
			"junos-vsrx_SecurityPoliciesPolicyPolicyThenCount": junosSecurityPoliciesPolicyPolicyThenCount(),
			"junos-vsrx_SecurityPoliciesPolicyPolicyThenDeny": junosSecurityPoliciesPolicyPolicyThenDeny(),
			"junos-vsrx_SecurityZonesSecurity__ZoneInterfacesHost__Inbound__TrafficProtocolsName": junosSecurityZonesSecurity__ZoneInterfacesHost__Inbound__TrafficProtocolsName(),
			"junos-vsrx_SecurityZonesSecurity__ZoneInterfacesHost__Inbound__TrafficSystem__ServicesName": junosSecurityZonesSecurity__ZoneInterfacesHost__Inbound__TrafficSystem__ServicesName(),
			"junos-vsrx_SecurityIkeProposalAuthentication__Method": junosSecurityIkeProposalAuthentication__Method(),
			"junos-vsrx_SecurityIkeProposalDh__Group": junosSecurityIkeProposalDh__Group(),
			"junos-vsrx_SecurityIkeProposalAuthentication__Algorithm": junosSecurityIkeProposalAuthentication__Algorithm(),
			"junos-vsrx_SecurityIkeProposalEncryption__Algorithm": junosSecurityIkeProposalEncryption__Algorithm(),
			"junos-vsrx_SecurityIkeProposalLifetime__Seconds": junosSecurityIkeProposalLifetime__Seconds(),
			"junos-vsrx_SecurityIkePolicyProposals": junosSecurityIkePolicyProposals(),
			"junos-vsrx_SecurityIkePolicyPre__Shared__KeyAscii__Text": junosSecurityIkePolicyPre__Shared__KeyAscii__Text(),
			"junos-vsrx_SecurityIkeGatewayAddress": junosSecurityIkeGatewayAddress(),
			"junos-vsrx_SecurityIkeGatewayIke__Policy": junosSecurityIkeGatewayIke__Policy(),
			"junos-vsrx_SecurityIkeGatewayLocal__Address": junosSecurityIkeGatewayLocal__Address(),
			"junos-vsrx_SecurityIkeGatewayExternal__Interface": junosSecurityIkeGatewayExternal__Interface(),
			"junos-vsrx_SecurityIkeGatewayVersion": junosSecurityIkeGatewayVersion(),
			"junos-vsrx_SecurityIpsecProposalProtocol": junosSecurityIpsecProposalProtocol(),
			"junos-vsrx_SecurityIpsecProposalAuthentication__Algorithm": junosSecurityIpsecProposalAuthentication__Algorithm(),
			"junos-vsrx_SecurityIpsecProposalEncryption__Algorithm": junosSecurityIpsecProposalEncryption__Algorithm(),
			"junos-vsrx_SecurityIpsecProposalLifetime__Seconds": junosSecurityIpsecProposalLifetime__Seconds(),
			"junos-vsrx_SecurityIpsecPolicyPerfect__Forward__SecrecyKeys": junosSecurityIpsecPolicyPerfect__Forward__SecrecyKeys(),
			"junos-vsrx_SecurityIpsecPolicyProposals": junosSecurityIpsecPolicyProposals(),
			"junos-vsrx_SecurityIpsecVpnBind__Interface": junosSecurityIpsecVpnBind__Interface(),
			"junos-vsrx_SecurityIpsecVpnIkeGateway": junosSecurityIpsecVpnIkeGateway(),
			"junos-vsrx_SecurityIpsecVpnIkeIpsec__Policy": junosSecurityIpsecVpnIkeIpsec__Policy(),
			"junos-vsrx_SecurityIpsecVpnEstablish__Tunnels": junosSecurityIpsecVpnEstablish__Tunnels(),
			"junos-vsrx_SecurityIpsecVpnTraffic__SelectorLocal__Ip": junosSecurityIpsecVpnTraffic__SelectorLocal__Ip(),
			"junos-vsrx_SecurityIpsecVpnTraffic__SelectorRemote__Ip": junosSecurityIpsecVpnTraffic__SelectorRemote__Ip(),
			"junos-vsrx_Routing__InstancesInstanceVrf__Table__Label": junosRouting__InstancesInstanceVrf__Table__Label(),
			"junos-vsrx_Routing__InstancesInstanceRoute__DistinguisherRd__Type": junosRouting__InstancesInstanceRoute__DistinguisherRd__Type(),
			"junos-vsrx_Routing__InstancesInstanceVrf__TargetCommunity": junosRouting__InstancesInstanceVrf__TargetCommunity(),
			"junos-vsrx_Routing__InstancesInstanceVrf__TargetImport": junosRouting__InstancesInstanceVrf__TargetImport(),
			"junos-vsrx_Routing__InstancesInstanceVrf__TargetExport": junosRouting__InstancesInstanceVrf__TargetExport(),
			"junos-vsrx_Routing__InstancesInstanceVrf__TargetAuto": junosRouting__InstancesInstanceVrf__TargetAuto(),
			"junos-vsrx_Routing__InstancesInstanceInterfaceName": junosRouting__InstancesInstanceInterfaceName(),
			"junos-vsrx_Routing__InstancesInstanceDescription": junosRouting__InstancesInstanceDescription(),
			"junos-vsrx_Routing__InstancesInstanceProtocolsBgpFamilyInetLabeled__Unicast": junosRouting__InstancesInstanceProtocolsBgpFamilyInetLabeled__Unicast(),
			"junos-vsrx_Routing__InstancesInstanceProtocolsBgpFamilyInetUnicast": junosRouting__InstancesInstanceProtocolsBgpFamilyInetUnicast(),
			"junos-vsrx_Routing__InstancesInstanceProtocolsBgpGroupType": junosRouting__InstancesInstanceProtocolsBgpGroupType(),
			"junos-vsrx_Routing__InstancesInstanceProtocolsBgpGroupDescription": junosRouting__InstancesInstanceProtocolsBgpGroupDescription(),
			"junos-vsrx_Routing__InstancesInstanceProtocolsBgpGroupMultihop": junosRouting__InstancesInstanceProtocolsBgpGroupMultihop(),
			"junos-vsrx_Routing__InstancesInstanceProtocolsBgpGroupLocal__Address": junosRouting__InstancesInstanceProtocolsBgpGroupLocal__Address(),
			"junos-vsrx_Routing__InstancesInstanceProtocolsBgpGroupLocal__Preference": junosRouting__InstancesInstanceProtocolsBgpGroupLocal__Preference(),
			"junos-vsrx_Routing__InstancesInstanceProtocolsBgpGroupAs__Override": junosRouting__InstancesInstanceProtocolsBgpGroupAs__Override(),
			"junos-vsrx_Routing__InstancesInstanceProtocolsBgpGroupAllow": junosRouting__InstancesInstanceProtocolsBgpGroupAllow(),
			"junos-vsrx_Routing__InstancesInstanceProtocolsBgpGroupNeighborMultihop": junosRouting__InstancesInstanceProtocolsBgpGroupNeighborMultihop(),
			"junos-vsrx_Routing__InstancesInstanceProtocolsBgpGroupNeighborLocal__Address": junosRouting__InstancesInstanceProtocolsBgpGroupNeighborLocal__Address(),
			"junos-vsrx_Routing__InstancesInstanceProtocolsBgpGroupNeighborFamilyInetUnicast": junosRouting__InstancesInstanceProtocolsBgpGroupNeighborFamilyInetUnicast(),
			"junos-vsrx_Routing__InstancesInstanceProtocolsBgpGroupNeighborFamilyInet6Unicast": junosRouting__InstancesInstanceProtocolsBgpGroupNeighborFamilyInet6Unicast(),
			"junos-vsrx_Routing__InstancesInstanceProtocolsBgpExport": junosRouting__InstancesInstanceProtocolsBgpExport(),
			"junos-vsrx_Routing__InstancesInstanceProtocolsBgpImport": junosRouting__InstancesInstanceProtocolsBgpImport(),
			"junos-vsrx_Policy__OptionsCommunityMembers": junosPolicy__OptionsCommunityMembers(),
			"junos-vsrx_Policy__OptionsPolicy__StatementTermFromProtocol": junosPolicy__OptionsPolicy__StatementTermFromProtocol(),
			"junos-vsrx_Policy__OptionsPolicy__StatementFromCommunity": junosPolicy__OptionsPolicy__StatementFromCommunity(),
			"junos-vsrx_Policy__OptionsPolicy__StatementThenAccept": junosPolicy__OptionsPolicy__StatementThenAccept(),
			"junos-vsrx_Policy__OptionsPolicy__StatementThenReject": junosPolicy__OptionsPolicy__StatementThenReject(),
			"junos-vsrx_Policy__OptionsPolicy__StatementThenNext": junosPolicy__OptionsPolicy__StatementThenNext(),
			"junos-vsrx_Policy__OptionsPolicy__StatementTermFromCommunity": junosPolicy__OptionsPolicy__StatementTermFromCommunity(),
			"junos-vsrx_Policy__OptionsPolicy__StatementTermThenCommunity": junosPolicy__OptionsPolicy__StatementTermThenCommunity(),
			"junos-vsrx_Policy__OptionsPolicy__StatementThenAs__Path__Prepend": junosPolicy__OptionsPolicy__StatementThenAs__Path__Prepend(),
			"junos-vsrx_Policy__OptionsPolicy__StatementThenAs__Path__Expand": junosPolicy__OptionsPolicy__StatementThenAs__Path__Expand(),
			"junos-vsrx_commit": junosCommit(),
	        	"junos-vsrx_destroycommit": junosDestroyCommit(),
			},
		    ConfigureFunc: providerConfigure,
	    } 
    }
