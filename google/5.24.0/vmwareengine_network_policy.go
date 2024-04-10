// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	vmwareenginenetworkpolicy "github.com/golingon/terraproviders/google/5.24.0/vmwareenginenetworkpolicy"
	"io"
)

// NewVmwareengineNetworkPolicy creates a new instance of [VmwareengineNetworkPolicy].
func NewVmwareengineNetworkPolicy(name string, args VmwareengineNetworkPolicyArgs) *VmwareengineNetworkPolicy {
	return &VmwareengineNetworkPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*VmwareengineNetworkPolicy)(nil)

// VmwareengineNetworkPolicy represents the Terraform resource google_vmwareengine_network_policy.
type VmwareengineNetworkPolicy struct {
	Name      string
	Args      VmwareengineNetworkPolicyArgs
	state     *vmwareengineNetworkPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [VmwareengineNetworkPolicy].
func (vnp *VmwareengineNetworkPolicy) Type() string {
	return "google_vmwareengine_network_policy"
}

// LocalName returns the local name for [VmwareengineNetworkPolicy].
func (vnp *VmwareengineNetworkPolicy) LocalName() string {
	return vnp.Name
}

// Configuration returns the configuration (args) for [VmwareengineNetworkPolicy].
func (vnp *VmwareengineNetworkPolicy) Configuration() interface{} {
	return vnp.Args
}

// DependOn is used for other resources to depend on [VmwareengineNetworkPolicy].
func (vnp *VmwareengineNetworkPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(vnp)
}

// Dependencies returns the list of resources [VmwareengineNetworkPolicy] depends_on.
func (vnp *VmwareengineNetworkPolicy) Dependencies() terra.Dependencies {
	return vnp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [VmwareengineNetworkPolicy].
func (vnp *VmwareengineNetworkPolicy) LifecycleManagement() *terra.Lifecycle {
	return vnp.Lifecycle
}

// Attributes returns the attributes for [VmwareengineNetworkPolicy].
func (vnp *VmwareengineNetworkPolicy) Attributes() vmwareengineNetworkPolicyAttributes {
	return vmwareengineNetworkPolicyAttributes{ref: terra.ReferenceResource(vnp)}
}

// ImportState imports the given attribute values into [VmwareengineNetworkPolicy]'s state.
func (vnp *VmwareengineNetworkPolicy) ImportState(av io.Reader) error {
	vnp.state = &vmwareengineNetworkPolicyState{}
	if err := json.NewDecoder(av).Decode(vnp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", vnp.Type(), vnp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [VmwareengineNetworkPolicy] has state.
func (vnp *VmwareengineNetworkPolicy) State() (*vmwareengineNetworkPolicyState, bool) {
	return vnp.state, vnp.state != nil
}

// StateMust returns the state for [VmwareengineNetworkPolicy]. Panics if the state is nil.
func (vnp *VmwareengineNetworkPolicy) StateMust() *vmwareengineNetworkPolicyState {
	if vnp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", vnp.Type(), vnp.LocalName()))
	}
	return vnp.state
}

// VmwareengineNetworkPolicyArgs contains the configurations for google_vmwareengine_network_policy.
type VmwareengineNetworkPolicyArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// EdgeServicesCidr: string, required
	EdgeServicesCidr terra.StringValue `hcl:"edge_services_cidr,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// VmwareEngineNetwork: string, required
	VmwareEngineNetwork terra.StringValue `hcl:"vmware_engine_network,attr" validate:"required"`
	// ExternalIp: optional
	ExternalIp *vmwareenginenetworkpolicy.ExternalIp `hcl:"external_ip,block"`
	// InternetAccess: optional
	InternetAccess *vmwareenginenetworkpolicy.InternetAccess `hcl:"internet_access,block"`
	// Timeouts: optional
	Timeouts *vmwareenginenetworkpolicy.Timeouts `hcl:"timeouts,block"`
}
type vmwareengineNetworkPolicyAttributes struct {
	ref terra.Reference
}

// CreateTime returns a reference to field create_time of google_vmwareengine_network_policy.
func (vnp vmwareengineNetworkPolicyAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(vnp.ref.Append("create_time"))
}

// Description returns a reference to field description of google_vmwareengine_network_policy.
func (vnp vmwareengineNetworkPolicyAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(vnp.ref.Append("description"))
}

// EdgeServicesCidr returns a reference to field edge_services_cidr of google_vmwareengine_network_policy.
func (vnp vmwareengineNetworkPolicyAttributes) EdgeServicesCidr() terra.StringValue {
	return terra.ReferenceAsString(vnp.ref.Append("edge_services_cidr"))
}

// Id returns a reference to field id of google_vmwareengine_network_policy.
func (vnp vmwareengineNetworkPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(vnp.ref.Append("id"))
}

// Location returns a reference to field location of google_vmwareengine_network_policy.
func (vnp vmwareengineNetworkPolicyAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(vnp.ref.Append("location"))
}

// Name returns a reference to field name of google_vmwareengine_network_policy.
func (vnp vmwareengineNetworkPolicyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(vnp.ref.Append("name"))
}

// Project returns a reference to field project of google_vmwareengine_network_policy.
func (vnp vmwareengineNetworkPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(vnp.ref.Append("project"))
}

// Uid returns a reference to field uid of google_vmwareengine_network_policy.
func (vnp vmwareengineNetworkPolicyAttributes) Uid() terra.StringValue {
	return terra.ReferenceAsString(vnp.ref.Append("uid"))
}

// UpdateTime returns a reference to field update_time of google_vmwareengine_network_policy.
func (vnp vmwareengineNetworkPolicyAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(vnp.ref.Append("update_time"))
}

// VmwareEngineNetwork returns a reference to field vmware_engine_network of google_vmwareengine_network_policy.
func (vnp vmwareengineNetworkPolicyAttributes) VmwareEngineNetwork() terra.StringValue {
	return terra.ReferenceAsString(vnp.ref.Append("vmware_engine_network"))
}

// VmwareEngineNetworkCanonical returns a reference to field vmware_engine_network_canonical of google_vmwareengine_network_policy.
func (vnp vmwareengineNetworkPolicyAttributes) VmwareEngineNetworkCanonical() terra.StringValue {
	return terra.ReferenceAsString(vnp.ref.Append("vmware_engine_network_canonical"))
}

func (vnp vmwareengineNetworkPolicyAttributes) ExternalIp() terra.ListValue[vmwareenginenetworkpolicy.ExternalIpAttributes] {
	return terra.ReferenceAsList[vmwareenginenetworkpolicy.ExternalIpAttributes](vnp.ref.Append("external_ip"))
}

func (vnp vmwareengineNetworkPolicyAttributes) InternetAccess() terra.ListValue[vmwareenginenetworkpolicy.InternetAccessAttributes] {
	return terra.ReferenceAsList[vmwareenginenetworkpolicy.InternetAccessAttributes](vnp.ref.Append("internet_access"))
}

func (vnp vmwareengineNetworkPolicyAttributes) Timeouts() vmwareenginenetworkpolicy.TimeoutsAttributes {
	return terra.ReferenceAsSingle[vmwareenginenetworkpolicy.TimeoutsAttributes](vnp.ref.Append("timeouts"))
}

type vmwareengineNetworkPolicyState struct {
	CreateTime                   string                                          `json:"create_time"`
	Description                  string                                          `json:"description"`
	EdgeServicesCidr             string                                          `json:"edge_services_cidr"`
	Id                           string                                          `json:"id"`
	Location                     string                                          `json:"location"`
	Name                         string                                          `json:"name"`
	Project                      string                                          `json:"project"`
	Uid                          string                                          `json:"uid"`
	UpdateTime                   string                                          `json:"update_time"`
	VmwareEngineNetwork          string                                          `json:"vmware_engine_network"`
	VmwareEngineNetworkCanonical string                                          `json:"vmware_engine_network_canonical"`
	ExternalIp                   []vmwareenginenetworkpolicy.ExternalIpState     `json:"external_ip"`
	InternetAccess               []vmwareenginenetworkpolicy.InternetAccessState `json:"internet_access"`
	Timeouts                     *vmwareenginenetworkpolicy.TimeoutsState        `json:"timeouts"`
}
