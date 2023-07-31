// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	computeglobaladdress "github.com/golingon/terraproviders/googlebeta/4.75.1/computeglobaladdress"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewComputeGlobalAddress creates a new instance of [ComputeGlobalAddress].
func NewComputeGlobalAddress(name string, args ComputeGlobalAddressArgs) *ComputeGlobalAddress {
	return &ComputeGlobalAddress{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeGlobalAddress)(nil)

// ComputeGlobalAddress represents the Terraform resource google_compute_global_address.
type ComputeGlobalAddress struct {
	Name      string
	Args      ComputeGlobalAddressArgs
	state     *computeGlobalAddressState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ComputeGlobalAddress].
func (cga *ComputeGlobalAddress) Type() string {
	return "google_compute_global_address"
}

// LocalName returns the local name for [ComputeGlobalAddress].
func (cga *ComputeGlobalAddress) LocalName() string {
	return cga.Name
}

// Configuration returns the configuration (args) for [ComputeGlobalAddress].
func (cga *ComputeGlobalAddress) Configuration() interface{} {
	return cga.Args
}

// DependOn is used for other resources to depend on [ComputeGlobalAddress].
func (cga *ComputeGlobalAddress) DependOn() terra.Reference {
	return terra.ReferenceResource(cga)
}

// Dependencies returns the list of resources [ComputeGlobalAddress] depends_on.
func (cga *ComputeGlobalAddress) Dependencies() terra.Dependencies {
	return cga.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ComputeGlobalAddress].
func (cga *ComputeGlobalAddress) LifecycleManagement() *terra.Lifecycle {
	return cga.Lifecycle
}

// Attributes returns the attributes for [ComputeGlobalAddress].
func (cga *ComputeGlobalAddress) Attributes() computeGlobalAddressAttributes {
	return computeGlobalAddressAttributes{ref: terra.ReferenceResource(cga)}
}

// ImportState imports the given attribute values into [ComputeGlobalAddress]'s state.
func (cga *ComputeGlobalAddress) ImportState(av io.Reader) error {
	cga.state = &computeGlobalAddressState{}
	if err := json.NewDecoder(av).Decode(cga.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cga.Type(), cga.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ComputeGlobalAddress] has state.
func (cga *ComputeGlobalAddress) State() (*computeGlobalAddressState, bool) {
	return cga.state, cga.state != nil
}

// StateMust returns the state for [ComputeGlobalAddress]. Panics if the state is nil.
func (cga *ComputeGlobalAddress) StateMust() *computeGlobalAddressState {
	if cga.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cga.Type(), cga.LocalName()))
	}
	return cga.state
}

// ComputeGlobalAddressArgs contains the configurations for google_compute_global_address.
type ComputeGlobalAddressArgs struct {
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
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Network: string, optional
	Network terra.StringValue `hcl:"network,attr"`
	// PrefixLength: number, optional
	PrefixLength terra.NumberValue `hcl:"prefix_length,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Purpose: string, optional
	Purpose terra.StringValue `hcl:"purpose,attr"`
	// Timeouts: optional
	Timeouts *computeglobaladdress.Timeouts `hcl:"timeouts,block"`
}
type computeGlobalAddressAttributes struct {
	ref terra.Reference
}

// Address returns a reference to field address of google_compute_global_address.
func (cga computeGlobalAddressAttributes) Address() terra.StringValue {
	return terra.ReferenceAsString(cga.ref.Append("address"))
}

// AddressType returns a reference to field address_type of google_compute_global_address.
func (cga computeGlobalAddressAttributes) AddressType() terra.StringValue {
	return terra.ReferenceAsString(cga.ref.Append("address_type"))
}

// CreationTimestamp returns a reference to field creation_timestamp of google_compute_global_address.
func (cga computeGlobalAddressAttributes) CreationTimestamp() terra.StringValue {
	return terra.ReferenceAsString(cga.ref.Append("creation_timestamp"))
}

// Description returns a reference to field description of google_compute_global_address.
func (cga computeGlobalAddressAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(cga.ref.Append("description"))
}

// Id returns a reference to field id of google_compute_global_address.
func (cga computeGlobalAddressAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cga.ref.Append("id"))
}

// IpVersion returns a reference to field ip_version of google_compute_global_address.
func (cga computeGlobalAddressAttributes) IpVersion() terra.StringValue {
	return terra.ReferenceAsString(cga.ref.Append("ip_version"))
}

// LabelFingerprint returns a reference to field label_fingerprint of google_compute_global_address.
func (cga computeGlobalAddressAttributes) LabelFingerprint() terra.StringValue {
	return terra.ReferenceAsString(cga.ref.Append("label_fingerprint"))
}

// Labels returns a reference to field labels of google_compute_global_address.
func (cga computeGlobalAddressAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cga.ref.Append("labels"))
}

// Name returns a reference to field name of google_compute_global_address.
func (cga computeGlobalAddressAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cga.ref.Append("name"))
}

// Network returns a reference to field network of google_compute_global_address.
func (cga computeGlobalAddressAttributes) Network() terra.StringValue {
	return terra.ReferenceAsString(cga.ref.Append("network"))
}

// PrefixLength returns a reference to field prefix_length of google_compute_global_address.
func (cga computeGlobalAddressAttributes) PrefixLength() terra.NumberValue {
	return terra.ReferenceAsNumber(cga.ref.Append("prefix_length"))
}

// Project returns a reference to field project of google_compute_global_address.
func (cga computeGlobalAddressAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(cga.ref.Append("project"))
}

// Purpose returns a reference to field purpose of google_compute_global_address.
func (cga computeGlobalAddressAttributes) Purpose() terra.StringValue {
	return terra.ReferenceAsString(cga.ref.Append("purpose"))
}

// SelfLink returns a reference to field self_link of google_compute_global_address.
func (cga computeGlobalAddressAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(cga.ref.Append("self_link"))
}

func (cga computeGlobalAddressAttributes) Timeouts() computeglobaladdress.TimeoutsAttributes {
	return terra.ReferenceAsSingle[computeglobaladdress.TimeoutsAttributes](cga.ref.Append("timeouts"))
}

type computeGlobalAddressState struct {
	Address           string                              `json:"address"`
	AddressType       string                              `json:"address_type"`
	CreationTimestamp string                              `json:"creation_timestamp"`
	Description       string                              `json:"description"`
	Id                string                              `json:"id"`
	IpVersion         string                              `json:"ip_version"`
	LabelFingerprint  string                              `json:"label_fingerprint"`
	Labels            map[string]string                   `json:"labels"`
	Name              string                              `json:"name"`
	Network           string                              `json:"network"`
	PrefixLength      float64                             `json:"prefix_length"`
	Project           string                              `json:"project"`
	Purpose           string                              `json:"purpose"`
	SelfLink          string                              `json:"self_link"`
	Timeouts          *computeglobaladdress.TimeoutsState `json:"timeouts"`
}
