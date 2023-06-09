// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package datacomputerouterstatus

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type BestRoutes struct{}

type BestRoutesForRouter struct{}

type BestRoutesAttributes struct {
	ref terra.Reference
}

func (br BestRoutesAttributes) InternalRef() (terra.Reference, error) {
	return br.ref, nil
}

func (br BestRoutesAttributes) InternalWithRef(ref terra.Reference) BestRoutesAttributes {
	return BestRoutesAttributes{ref: ref}
}

func (br BestRoutesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return br.ref.InternalTokens()
}

func (br BestRoutesAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(br.ref.Append("description"))
}

func (br BestRoutesAttributes) DestRange() terra.StringValue {
	return terra.ReferenceAsString(br.ref.Append("dest_range"))
}

func (br BestRoutesAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(br.ref.Append("name"))
}

func (br BestRoutesAttributes) Network() terra.StringValue {
	return terra.ReferenceAsString(br.ref.Append("network"))
}

func (br BestRoutesAttributes) NextHopGateway() terra.StringValue {
	return terra.ReferenceAsString(br.ref.Append("next_hop_gateway"))
}

func (br BestRoutesAttributes) NextHopIlb() terra.StringValue {
	return terra.ReferenceAsString(br.ref.Append("next_hop_ilb"))
}

func (br BestRoutesAttributes) NextHopInstance() terra.StringValue {
	return terra.ReferenceAsString(br.ref.Append("next_hop_instance"))
}

func (br BestRoutesAttributes) NextHopInstanceZone() terra.StringValue {
	return terra.ReferenceAsString(br.ref.Append("next_hop_instance_zone"))
}

func (br BestRoutesAttributes) NextHopIp() terra.StringValue {
	return terra.ReferenceAsString(br.ref.Append("next_hop_ip"))
}

func (br BestRoutesAttributes) NextHopNetwork() terra.StringValue {
	return terra.ReferenceAsString(br.ref.Append("next_hop_network"))
}

func (br BestRoutesAttributes) NextHopVpnTunnel() terra.StringValue {
	return terra.ReferenceAsString(br.ref.Append("next_hop_vpn_tunnel"))
}

func (br BestRoutesAttributes) Priority() terra.NumberValue {
	return terra.ReferenceAsNumber(br.ref.Append("priority"))
}

func (br BestRoutesAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(br.ref.Append("project"))
}

func (br BestRoutesAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(br.ref.Append("self_link"))
}

func (br BestRoutesAttributes) Tags() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](br.ref.Append("tags"))
}

type BestRoutesForRouterAttributes struct {
	ref terra.Reference
}

func (brfr BestRoutesForRouterAttributes) InternalRef() (terra.Reference, error) {
	return brfr.ref, nil
}

func (brfr BestRoutesForRouterAttributes) InternalWithRef(ref terra.Reference) BestRoutesForRouterAttributes {
	return BestRoutesForRouterAttributes{ref: ref}
}

func (brfr BestRoutesForRouterAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return brfr.ref.InternalTokens()
}

func (brfr BestRoutesForRouterAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(brfr.ref.Append("description"))
}

func (brfr BestRoutesForRouterAttributes) DestRange() terra.StringValue {
	return terra.ReferenceAsString(brfr.ref.Append("dest_range"))
}

func (brfr BestRoutesForRouterAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(brfr.ref.Append("name"))
}

func (brfr BestRoutesForRouterAttributes) Network() terra.StringValue {
	return terra.ReferenceAsString(brfr.ref.Append("network"))
}

func (brfr BestRoutesForRouterAttributes) NextHopGateway() terra.StringValue {
	return terra.ReferenceAsString(brfr.ref.Append("next_hop_gateway"))
}

func (brfr BestRoutesForRouterAttributes) NextHopIlb() terra.StringValue {
	return terra.ReferenceAsString(brfr.ref.Append("next_hop_ilb"))
}

func (brfr BestRoutesForRouterAttributes) NextHopInstance() terra.StringValue {
	return terra.ReferenceAsString(brfr.ref.Append("next_hop_instance"))
}

func (brfr BestRoutesForRouterAttributes) NextHopInstanceZone() terra.StringValue {
	return terra.ReferenceAsString(brfr.ref.Append("next_hop_instance_zone"))
}

func (brfr BestRoutesForRouterAttributes) NextHopIp() terra.StringValue {
	return terra.ReferenceAsString(brfr.ref.Append("next_hop_ip"))
}

func (brfr BestRoutesForRouterAttributes) NextHopNetwork() terra.StringValue {
	return terra.ReferenceAsString(brfr.ref.Append("next_hop_network"))
}

func (brfr BestRoutesForRouterAttributes) NextHopVpnTunnel() terra.StringValue {
	return terra.ReferenceAsString(brfr.ref.Append("next_hop_vpn_tunnel"))
}

func (brfr BestRoutesForRouterAttributes) Priority() terra.NumberValue {
	return terra.ReferenceAsNumber(brfr.ref.Append("priority"))
}

func (brfr BestRoutesForRouterAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(brfr.ref.Append("project"))
}

func (brfr BestRoutesForRouterAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(brfr.ref.Append("self_link"))
}

func (brfr BestRoutesForRouterAttributes) Tags() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](brfr.ref.Append("tags"))
}

type BestRoutesState struct {
	Description         string   `json:"description"`
	DestRange           string   `json:"dest_range"`
	Name                string   `json:"name"`
	Network             string   `json:"network"`
	NextHopGateway      string   `json:"next_hop_gateway"`
	NextHopIlb          string   `json:"next_hop_ilb"`
	NextHopInstance     string   `json:"next_hop_instance"`
	NextHopInstanceZone string   `json:"next_hop_instance_zone"`
	NextHopIp           string   `json:"next_hop_ip"`
	NextHopNetwork      string   `json:"next_hop_network"`
	NextHopVpnTunnel    string   `json:"next_hop_vpn_tunnel"`
	Priority            float64  `json:"priority"`
	Project             string   `json:"project"`
	SelfLink            string   `json:"self_link"`
	Tags                []string `json:"tags"`
}

type BestRoutesForRouterState struct {
	Description         string   `json:"description"`
	DestRange           string   `json:"dest_range"`
	Name                string   `json:"name"`
	Network             string   `json:"network"`
	NextHopGateway      string   `json:"next_hop_gateway"`
	NextHopIlb          string   `json:"next_hop_ilb"`
	NextHopInstance     string   `json:"next_hop_instance"`
	NextHopInstanceZone string   `json:"next_hop_instance_zone"`
	NextHopIp           string   `json:"next_hop_ip"`
	NextHopNetwork      string   `json:"next_hop_network"`
	NextHopVpnTunnel    string   `json:"next_hop_vpn_tunnel"`
	Priority            float64  `json:"priority"`
	Project             string   `json:"project"`
	SelfLink            string   `json:"self_link"`
	Tags                []string `json:"tags"`
}
