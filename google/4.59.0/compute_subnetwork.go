// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	computesubnetwork "github.com/golingon/terraproviders/google/4.59.0/computesubnetwork"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewComputeSubnetwork(name string, args ComputeSubnetworkArgs) *ComputeSubnetwork {
	return &ComputeSubnetwork{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeSubnetwork)(nil)

type ComputeSubnetwork struct {
	Name  string
	Args  ComputeSubnetworkArgs
	state *computeSubnetworkState
}

func (cs *ComputeSubnetwork) Type() string {
	return "google_compute_subnetwork"
}

func (cs *ComputeSubnetwork) LocalName() string {
	return cs.Name
}

func (cs *ComputeSubnetwork) Configuration() interface{} {
	return cs.Args
}

func (cs *ComputeSubnetwork) Attributes() computeSubnetworkAttributes {
	return computeSubnetworkAttributes{ref: terra.ReferenceResource(cs)}
}

func (cs *ComputeSubnetwork) ImportState(av io.Reader) error {
	cs.state = &computeSubnetworkState{}
	if err := json.NewDecoder(av).Decode(cs.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cs.Type(), cs.LocalName(), err)
	}
	return nil
}

func (cs *ComputeSubnetwork) State() (*computeSubnetworkState, bool) {
	return cs.state, cs.state != nil
}

func (cs *ComputeSubnetwork) StateMust() *computeSubnetworkState {
	if cs.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cs.Type(), cs.LocalName()))
	}
	return cs.state
}

func (cs *ComputeSubnetwork) DependOn() terra.Reference {
	return terra.ReferenceResource(cs)
}

type ComputeSubnetworkArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IpCidrRange: string, required
	IpCidrRange terra.StringValue `hcl:"ip_cidr_range,attr" validate:"required"`
	// Ipv6AccessType: string, optional
	Ipv6AccessType terra.StringValue `hcl:"ipv6_access_type,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Network: string, required
	Network terra.StringValue `hcl:"network,attr" validate:"required"`
	// PrivateIpGoogleAccess: bool, optional
	PrivateIpGoogleAccess terra.BoolValue `hcl:"private_ip_google_access,attr"`
	// PrivateIpv6GoogleAccess: string, optional
	PrivateIpv6GoogleAccess terra.StringValue `hcl:"private_ipv6_google_access,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Purpose: string, optional
	Purpose terra.StringValue `hcl:"purpose,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
	// Role: string, optional
	Role terra.StringValue `hcl:"role,attr"`
	// StackType: string, optional
	StackType terra.StringValue `hcl:"stack_type,attr"`
	// SecondaryIpRange: min=0
	SecondaryIpRange []computesubnetwork.SecondaryIpRange `hcl:"secondary_ip_range,block" validate:"min=0"`
	// LogConfig: optional
	LogConfig *computesubnetwork.LogConfig `hcl:"log_config,block"`
	// Timeouts: optional
	Timeouts *computesubnetwork.Timeouts `hcl:"timeouts,block"`
	// DependsOn contains resources that ComputeSubnetwork depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type computeSubnetworkAttributes struct {
	ref terra.Reference
}

func (cs computeSubnetworkAttributes) CreationTimestamp() terra.StringValue {
	return terra.ReferenceString(cs.ref.Append("creation_timestamp"))
}

func (cs computeSubnetworkAttributes) Description() terra.StringValue {
	return terra.ReferenceString(cs.ref.Append("description"))
}

func (cs computeSubnetworkAttributes) ExternalIpv6Prefix() terra.StringValue {
	return terra.ReferenceString(cs.ref.Append("external_ipv6_prefix"))
}

func (cs computeSubnetworkAttributes) Fingerprint() terra.StringValue {
	return terra.ReferenceString(cs.ref.Append("fingerprint"))
}

func (cs computeSubnetworkAttributes) GatewayAddress() terra.StringValue {
	return terra.ReferenceString(cs.ref.Append("gateway_address"))
}

func (cs computeSubnetworkAttributes) Id() terra.StringValue {
	return terra.ReferenceString(cs.ref.Append("id"))
}

func (cs computeSubnetworkAttributes) IpCidrRange() terra.StringValue {
	return terra.ReferenceString(cs.ref.Append("ip_cidr_range"))
}

func (cs computeSubnetworkAttributes) Ipv6AccessType() terra.StringValue {
	return terra.ReferenceString(cs.ref.Append("ipv6_access_type"))
}

func (cs computeSubnetworkAttributes) Ipv6CidrRange() terra.StringValue {
	return terra.ReferenceString(cs.ref.Append("ipv6_cidr_range"))
}

func (cs computeSubnetworkAttributes) Name() terra.StringValue {
	return terra.ReferenceString(cs.ref.Append("name"))
}

func (cs computeSubnetworkAttributes) Network() terra.StringValue {
	return terra.ReferenceString(cs.ref.Append("network"))
}

func (cs computeSubnetworkAttributes) PrivateIpGoogleAccess() terra.BoolValue {
	return terra.ReferenceBool(cs.ref.Append("private_ip_google_access"))
}

func (cs computeSubnetworkAttributes) PrivateIpv6GoogleAccess() terra.StringValue {
	return terra.ReferenceString(cs.ref.Append("private_ipv6_google_access"))
}

func (cs computeSubnetworkAttributes) Project() terra.StringValue {
	return terra.ReferenceString(cs.ref.Append("project"))
}

func (cs computeSubnetworkAttributes) Purpose() terra.StringValue {
	return terra.ReferenceString(cs.ref.Append("purpose"))
}

func (cs computeSubnetworkAttributes) Region() terra.StringValue {
	return terra.ReferenceString(cs.ref.Append("region"))
}

func (cs computeSubnetworkAttributes) Role() terra.StringValue {
	return terra.ReferenceString(cs.ref.Append("role"))
}

func (cs computeSubnetworkAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceString(cs.ref.Append("self_link"))
}

func (cs computeSubnetworkAttributes) StackType() terra.StringValue {
	return terra.ReferenceString(cs.ref.Append("stack_type"))
}

func (cs computeSubnetworkAttributes) SecondaryIpRange() terra.ListValue[computesubnetwork.SecondaryIpRangeAttributes] {
	return terra.ReferenceList[computesubnetwork.SecondaryIpRangeAttributes](cs.ref.Append("secondary_ip_range"))
}

func (cs computeSubnetworkAttributes) LogConfig() terra.ListValue[computesubnetwork.LogConfigAttributes] {
	return terra.ReferenceList[computesubnetwork.LogConfigAttributes](cs.ref.Append("log_config"))
}

func (cs computeSubnetworkAttributes) Timeouts() computesubnetwork.TimeoutsAttributes {
	return terra.ReferenceSingle[computesubnetwork.TimeoutsAttributes](cs.ref.Append("timeouts"))
}

type computeSubnetworkState struct {
	CreationTimestamp       string                                    `json:"creation_timestamp"`
	Description             string                                    `json:"description"`
	ExternalIpv6Prefix      string                                    `json:"external_ipv6_prefix"`
	Fingerprint             string                                    `json:"fingerprint"`
	GatewayAddress          string                                    `json:"gateway_address"`
	Id                      string                                    `json:"id"`
	IpCidrRange             string                                    `json:"ip_cidr_range"`
	Ipv6AccessType          string                                    `json:"ipv6_access_type"`
	Ipv6CidrRange           string                                    `json:"ipv6_cidr_range"`
	Name                    string                                    `json:"name"`
	Network                 string                                    `json:"network"`
	PrivateIpGoogleAccess   bool                                      `json:"private_ip_google_access"`
	PrivateIpv6GoogleAccess string                                    `json:"private_ipv6_google_access"`
	Project                 string                                    `json:"project"`
	Purpose                 string                                    `json:"purpose"`
	Region                  string                                    `json:"region"`
	Role                    string                                    `json:"role"`
	SelfLink                string                                    `json:"self_link"`
	StackType               string                                    `json:"stack_type"`
	SecondaryIpRange        []computesubnetwork.SecondaryIpRangeState `json:"secondary_ip_range"`
	LogConfig               []computesubnetwork.LogConfigState        `json:"log_config"`
	Timeouts                *computesubnetwork.TimeoutsState          `json:"timeouts"`
}
