// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	computesubnetwork "github.com/golingon/terraproviders/google/4.77.0/computesubnetwork"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewComputeSubnetwork creates a new instance of [ComputeSubnetwork].
func NewComputeSubnetwork(name string, args ComputeSubnetworkArgs) *ComputeSubnetwork {
	return &ComputeSubnetwork{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeSubnetwork)(nil)

// ComputeSubnetwork represents the Terraform resource google_compute_subnetwork.
type ComputeSubnetwork struct {
	Name      string
	Args      ComputeSubnetworkArgs
	state     *computeSubnetworkState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ComputeSubnetwork].
func (cs *ComputeSubnetwork) Type() string {
	return "google_compute_subnetwork"
}

// LocalName returns the local name for [ComputeSubnetwork].
func (cs *ComputeSubnetwork) LocalName() string {
	return cs.Name
}

// Configuration returns the configuration (args) for [ComputeSubnetwork].
func (cs *ComputeSubnetwork) Configuration() interface{} {
	return cs.Args
}

// DependOn is used for other resources to depend on [ComputeSubnetwork].
func (cs *ComputeSubnetwork) DependOn() terra.Reference {
	return terra.ReferenceResource(cs)
}

// Dependencies returns the list of resources [ComputeSubnetwork] depends_on.
func (cs *ComputeSubnetwork) Dependencies() terra.Dependencies {
	return cs.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ComputeSubnetwork].
func (cs *ComputeSubnetwork) LifecycleManagement() *terra.Lifecycle {
	return cs.Lifecycle
}

// Attributes returns the attributes for [ComputeSubnetwork].
func (cs *ComputeSubnetwork) Attributes() computeSubnetworkAttributes {
	return computeSubnetworkAttributes{ref: terra.ReferenceResource(cs)}
}

// ImportState imports the given attribute values into [ComputeSubnetwork]'s state.
func (cs *ComputeSubnetwork) ImportState(av io.Reader) error {
	cs.state = &computeSubnetworkState{}
	if err := json.NewDecoder(av).Decode(cs.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cs.Type(), cs.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ComputeSubnetwork] has state.
func (cs *ComputeSubnetwork) State() (*computeSubnetworkState, bool) {
	return cs.state, cs.state != nil
}

// StateMust returns the state for [ComputeSubnetwork]. Panics if the state is nil.
func (cs *ComputeSubnetwork) StateMust() *computeSubnetworkState {
	if cs.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cs.Type(), cs.LocalName()))
	}
	return cs.state
}

// ComputeSubnetworkArgs contains the configurations for google_compute_subnetwork.
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
}
type computeSubnetworkAttributes struct {
	ref terra.Reference
}

// CreationTimestamp returns a reference to field creation_timestamp of google_compute_subnetwork.
func (cs computeSubnetworkAttributes) CreationTimestamp() terra.StringValue {
	return terra.ReferenceAsString(cs.ref.Append("creation_timestamp"))
}

// Description returns a reference to field description of google_compute_subnetwork.
func (cs computeSubnetworkAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(cs.ref.Append("description"))
}

// ExternalIpv6Prefix returns a reference to field external_ipv6_prefix of google_compute_subnetwork.
func (cs computeSubnetworkAttributes) ExternalIpv6Prefix() terra.StringValue {
	return terra.ReferenceAsString(cs.ref.Append("external_ipv6_prefix"))
}

// Fingerprint returns a reference to field fingerprint of google_compute_subnetwork.
func (cs computeSubnetworkAttributes) Fingerprint() terra.StringValue {
	return terra.ReferenceAsString(cs.ref.Append("fingerprint"))
}

// GatewayAddress returns a reference to field gateway_address of google_compute_subnetwork.
func (cs computeSubnetworkAttributes) GatewayAddress() terra.StringValue {
	return terra.ReferenceAsString(cs.ref.Append("gateway_address"))
}

// Id returns a reference to field id of google_compute_subnetwork.
func (cs computeSubnetworkAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cs.ref.Append("id"))
}

// IpCidrRange returns a reference to field ip_cidr_range of google_compute_subnetwork.
func (cs computeSubnetworkAttributes) IpCidrRange() terra.StringValue {
	return terra.ReferenceAsString(cs.ref.Append("ip_cidr_range"))
}

// Ipv6AccessType returns a reference to field ipv6_access_type of google_compute_subnetwork.
func (cs computeSubnetworkAttributes) Ipv6AccessType() terra.StringValue {
	return terra.ReferenceAsString(cs.ref.Append("ipv6_access_type"))
}

// Ipv6CidrRange returns a reference to field ipv6_cidr_range of google_compute_subnetwork.
func (cs computeSubnetworkAttributes) Ipv6CidrRange() terra.StringValue {
	return terra.ReferenceAsString(cs.ref.Append("ipv6_cidr_range"))
}

// Name returns a reference to field name of google_compute_subnetwork.
func (cs computeSubnetworkAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cs.ref.Append("name"))
}

// Network returns a reference to field network of google_compute_subnetwork.
func (cs computeSubnetworkAttributes) Network() terra.StringValue {
	return terra.ReferenceAsString(cs.ref.Append("network"))
}

// PrivateIpGoogleAccess returns a reference to field private_ip_google_access of google_compute_subnetwork.
func (cs computeSubnetworkAttributes) PrivateIpGoogleAccess() terra.BoolValue {
	return terra.ReferenceAsBool(cs.ref.Append("private_ip_google_access"))
}

// PrivateIpv6GoogleAccess returns a reference to field private_ipv6_google_access of google_compute_subnetwork.
func (cs computeSubnetworkAttributes) PrivateIpv6GoogleAccess() terra.StringValue {
	return terra.ReferenceAsString(cs.ref.Append("private_ipv6_google_access"))
}

// Project returns a reference to field project of google_compute_subnetwork.
func (cs computeSubnetworkAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(cs.ref.Append("project"))
}

// Purpose returns a reference to field purpose of google_compute_subnetwork.
func (cs computeSubnetworkAttributes) Purpose() terra.StringValue {
	return terra.ReferenceAsString(cs.ref.Append("purpose"))
}

// Region returns a reference to field region of google_compute_subnetwork.
func (cs computeSubnetworkAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(cs.ref.Append("region"))
}

// Role returns a reference to field role of google_compute_subnetwork.
func (cs computeSubnetworkAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(cs.ref.Append("role"))
}

// SelfLink returns a reference to field self_link of google_compute_subnetwork.
func (cs computeSubnetworkAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(cs.ref.Append("self_link"))
}

// StackType returns a reference to field stack_type of google_compute_subnetwork.
func (cs computeSubnetworkAttributes) StackType() terra.StringValue {
	return terra.ReferenceAsString(cs.ref.Append("stack_type"))
}

func (cs computeSubnetworkAttributes) SecondaryIpRange() terra.ListValue[computesubnetwork.SecondaryIpRangeAttributes] {
	return terra.ReferenceAsList[computesubnetwork.SecondaryIpRangeAttributes](cs.ref.Append("secondary_ip_range"))
}

func (cs computeSubnetworkAttributes) LogConfig() terra.ListValue[computesubnetwork.LogConfigAttributes] {
	return terra.ReferenceAsList[computesubnetwork.LogConfigAttributes](cs.ref.Append("log_config"))
}

func (cs computeSubnetworkAttributes) Timeouts() computesubnetwork.TimeoutsAttributes {
	return terra.ReferenceAsSingle[computesubnetwork.TimeoutsAttributes](cs.ref.Append("timeouts"))
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
