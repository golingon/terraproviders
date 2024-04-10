// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	computeaddress "github.com/golingon/terraproviders/google/5.24.0/computeaddress"
	"io"
)

// NewComputeAddress creates a new instance of [ComputeAddress].
func NewComputeAddress(name string, args ComputeAddressArgs) *ComputeAddress {
	return &ComputeAddress{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeAddress)(nil)

// ComputeAddress represents the Terraform resource google_compute_address.
type ComputeAddress struct {
	Name      string
	Args      ComputeAddressArgs
	state     *computeAddressState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ComputeAddress].
func (ca *ComputeAddress) Type() string {
	return "google_compute_address"
}

// LocalName returns the local name for [ComputeAddress].
func (ca *ComputeAddress) LocalName() string {
	return ca.Name
}

// Configuration returns the configuration (args) for [ComputeAddress].
func (ca *ComputeAddress) Configuration() interface{} {
	return ca.Args
}

// DependOn is used for other resources to depend on [ComputeAddress].
func (ca *ComputeAddress) DependOn() terra.Reference {
	return terra.ReferenceResource(ca)
}

// Dependencies returns the list of resources [ComputeAddress] depends_on.
func (ca *ComputeAddress) Dependencies() terra.Dependencies {
	return ca.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ComputeAddress].
func (ca *ComputeAddress) LifecycleManagement() *terra.Lifecycle {
	return ca.Lifecycle
}

// Attributes returns the attributes for [ComputeAddress].
func (ca *ComputeAddress) Attributes() computeAddressAttributes {
	return computeAddressAttributes{ref: terra.ReferenceResource(ca)}
}

// ImportState imports the given attribute values into [ComputeAddress]'s state.
func (ca *ComputeAddress) ImportState(av io.Reader) error {
	ca.state = &computeAddressState{}
	if err := json.NewDecoder(av).Decode(ca.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ca.Type(), ca.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ComputeAddress] has state.
func (ca *ComputeAddress) State() (*computeAddressState, bool) {
	return ca.state, ca.state != nil
}

// StateMust returns the state for [ComputeAddress]. Panics if the state is nil.
func (ca *ComputeAddress) StateMust() *computeAddressState {
	if ca.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ca.Type(), ca.LocalName()))
	}
	return ca.state
}

// ComputeAddressArgs contains the configurations for google_compute_address.
type ComputeAddressArgs struct {
	// Address: string, optional
	Address terra.StringValue `hcl:"address,attr"`
	// AddressType: string, optional
	AddressType terra.StringValue `hcl:"address_type,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IpVersion: string, optional
	IpVersion terra.StringValue `hcl:"ip_version,attr"`
	// Ipv6EndpointType: string, optional
	Ipv6EndpointType terra.StringValue `hcl:"ipv6_endpoint_type,attr"`
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Network: string, optional
	Network terra.StringValue `hcl:"network,attr"`
	// NetworkTier: string, optional
	NetworkTier terra.StringValue `hcl:"network_tier,attr"`
	// PrefixLength: number, optional
	PrefixLength terra.NumberValue `hcl:"prefix_length,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Purpose: string, optional
	Purpose terra.StringValue `hcl:"purpose,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
	// Subnetwork: string, optional
	Subnetwork terra.StringValue `hcl:"subnetwork,attr"`
	// Timeouts: optional
	Timeouts *computeaddress.Timeouts `hcl:"timeouts,block"`
}
type computeAddressAttributes struct {
	ref terra.Reference
}

// Address returns a reference to field address of google_compute_address.
func (ca computeAddressAttributes) Address() terra.StringValue {
	return terra.ReferenceAsString(ca.ref.Append("address"))
}

// AddressType returns a reference to field address_type of google_compute_address.
func (ca computeAddressAttributes) AddressType() terra.StringValue {
	return terra.ReferenceAsString(ca.ref.Append("address_type"))
}

// CreationTimestamp returns a reference to field creation_timestamp of google_compute_address.
func (ca computeAddressAttributes) CreationTimestamp() terra.StringValue {
	return terra.ReferenceAsString(ca.ref.Append("creation_timestamp"))
}

// Description returns a reference to field description of google_compute_address.
func (ca computeAddressAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(ca.ref.Append("description"))
}

// EffectiveLabels returns a reference to field effective_labels of google_compute_address.
func (ca computeAddressAttributes) EffectiveLabels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ca.ref.Append("effective_labels"))
}

// Id returns a reference to field id of google_compute_address.
func (ca computeAddressAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ca.ref.Append("id"))
}

// IpVersion returns a reference to field ip_version of google_compute_address.
func (ca computeAddressAttributes) IpVersion() terra.StringValue {
	return terra.ReferenceAsString(ca.ref.Append("ip_version"))
}

// Ipv6EndpointType returns a reference to field ipv6_endpoint_type of google_compute_address.
func (ca computeAddressAttributes) Ipv6EndpointType() terra.StringValue {
	return terra.ReferenceAsString(ca.ref.Append("ipv6_endpoint_type"))
}

// LabelFingerprint returns a reference to field label_fingerprint of google_compute_address.
func (ca computeAddressAttributes) LabelFingerprint() terra.StringValue {
	return terra.ReferenceAsString(ca.ref.Append("label_fingerprint"))
}

// Labels returns a reference to field labels of google_compute_address.
func (ca computeAddressAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ca.ref.Append("labels"))
}

// Name returns a reference to field name of google_compute_address.
func (ca computeAddressAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ca.ref.Append("name"))
}

// Network returns a reference to field network of google_compute_address.
func (ca computeAddressAttributes) Network() terra.StringValue {
	return terra.ReferenceAsString(ca.ref.Append("network"))
}

// NetworkTier returns a reference to field network_tier of google_compute_address.
func (ca computeAddressAttributes) NetworkTier() terra.StringValue {
	return terra.ReferenceAsString(ca.ref.Append("network_tier"))
}

// PrefixLength returns a reference to field prefix_length of google_compute_address.
func (ca computeAddressAttributes) PrefixLength() terra.NumberValue {
	return terra.ReferenceAsNumber(ca.ref.Append("prefix_length"))
}

// Project returns a reference to field project of google_compute_address.
func (ca computeAddressAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(ca.ref.Append("project"))
}

// Purpose returns a reference to field purpose of google_compute_address.
func (ca computeAddressAttributes) Purpose() terra.StringValue {
	return terra.ReferenceAsString(ca.ref.Append("purpose"))
}

// Region returns a reference to field region of google_compute_address.
func (ca computeAddressAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(ca.ref.Append("region"))
}

// SelfLink returns a reference to field self_link of google_compute_address.
func (ca computeAddressAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(ca.ref.Append("self_link"))
}

// Subnetwork returns a reference to field subnetwork of google_compute_address.
func (ca computeAddressAttributes) Subnetwork() terra.StringValue {
	return terra.ReferenceAsString(ca.ref.Append("subnetwork"))
}

// TerraformLabels returns a reference to field terraform_labels of google_compute_address.
func (ca computeAddressAttributes) TerraformLabels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ca.ref.Append("terraform_labels"))
}

// Users returns a reference to field users of google_compute_address.
func (ca computeAddressAttributes) Users() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ca.ref.Append("users"))
}

func (ca computeAddressAttributes) Timeouts() computeaddress.TimeoutsAttributes {
	return terra.ReferenceAsSingle[computeaddress.TimeoutsAttributes](ca.ref.Append("timeouts"))
}

type computeAddressState struct {
	Address           string                        `json:"address"`
	AddressType       string                        `json:"address_type"`
	CreationTimestamp string                        `json:"creation_timestamp"`
	Description       string                        `json:"description"`
	EffectiveLabels   map[string]string             `json:"effective_labels"`
	Id                string                        `json:"id"`
	IpVersion         string                        `json:"ip_version"`
	Ipv6EndpointType  string                        `json:"ipv6_endpoint_type"`
	LabelFingerprint  string                        `json:"label_fingerprint"`
	Labels            map[string]string             `json:"labels"`
	Name              string                        `json:"name"`
	Network           string                        `json:"network"`
	NetworkTier       string                        `json:"network_tier"`
	PrefixLength      float64                       `json:"prefix_length"`
	Project           string                        `json:"project"`
	Purpose           string                        `json:"purpose"`
	Region            string                        `json:"region"`
	SelfLink          string                        `json:"self_link"`
	Subnetwork        string                        `json:"subnetwork"`
	TerraformLabels   map[string]string             `json:"terraform_labels"`
	Users             []string                      `json:"users"`
	Timeouts          *computeaddress.TimeoutsState `json:"timeouts"`
}
