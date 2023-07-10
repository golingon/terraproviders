// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	computeregionsslpolicy "github.com/golingon/terraproviders/googlebeta/4.72.1/computeregionsslpolicy"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewComputeRegionSslPolicy creates a new instance of [ComputeRegionSslPolicy].
func NewComputeRegionSslPolicy(name string, args ComputeRegionSslPolicyArgs) *ComputeRegionSslPolicy {
	return &ComputeRegionSslPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeRegionSslPolicy)(nil)

// ComputeRegionSslPolicy represents the Terraform resource google_compute_region_ssl_policy.
type ComputeRegionSslPolicy struct {
	Name      string
	Args      ComputeRegionSslPolicyArgs
	state     *computeRegionSslPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ComputeRegionSslPolicy].
func (crsp *ComputeRegionSslPolicy) Type() string {
	return "google_compute_region_ssl_policy"
}

// LocalName returns the local name for [ComputeRegionSslPolicy].
func (crsp *ComputeRegionSslPolicy) LocalName() string {
	return crsp.Name
}

// Configuration returns the configuration (args) for [ComputeRegionSslPolicy].
func (crsp *ComputeRegionSslPolicy) Configuration() interface{} {
	return crsp.Args
}

// DependOn is used for other resources to depend on [ComputeRegionSslPolicy].
func (crsp *ComputeRegionSslPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(crsp)
}

// Dependencies returns the list of resources [ComputeRegionSslPolicy] depends_on.
func (crsp *ComputeRegionSslPolicy) Dependencies() terra.Dependencies {
	return crsp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ComputeRegionSslPolicy].
func (crsp *ComputeRegionSslPolicy) LifecycleManagement() *terra.Lifecycle {
	return crsp.Lifecycle
}

// Attributes returns the attributes for [ComputeRegionSslPolicy].
func (crsp *ComputeRegionSslPolicy) Attributes() computeRegionSslPolicyAttributes {
	return computeRegionSslPolicyAttributes{ref: terra.ReferenceResource(crsp)}
}

// ImportState imports the given attribute values into [ComputeRegionSslPolicy]'s state.
func (crsp *ComputeRegionSslPolicy) ImportState(av io.Reader) error {
	crsp.state = &computeRegionSslPolicyState{}
	if err := json.NewDecoder(av).Decode(crsp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", crsp.Type(), crsp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ComputeRegionSslPolicy] has state.
func (crsp *ComputeRegionSslPolicy) State() (*computeRegionSslPolicyState, bool) {
	return crsp.state, crsp.state != nil
}

// StateMust returns the state for [ComputeRegionSslPolicy]. Panics if the state is nil.
func (crsp *ComputeRegionSslPolicy) StateMust() *computeRegionSslPolicyState {
	if crsp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", crsp.Type(), crsp.LocalName()))
	}
	return crsp.state
}

// ComputeRegionSslPolicyArgs contains the configurations for google_compute_region_ssl_policy.
type ComputeRegionSslPolicyArgs struct {
	// CustomFeatures: set of string, optional
	CustomFeatures terra.SetValue[terra.StringValue] `hcl:"custom_features,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// MinTlsVersion: string, optional
	MinTlsVersion terra.StringValue `hcl:"min_tls_version,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Profile: string, optional
	Profile terra.StringValue `hcl:"profile,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Region: string, required
	Region terra.StringValue `hcl:"region,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *computeregionsslpolicy.Timeouts `hcl:"timeouts,block"`
}
type computeRegionSslPolicyAttributes struct {
	ref terra.Reference
}

// CreationTimestamp returns a reference to field creation_timestamp of google_compute_region_ssl_policy.
func (crsp computeRegionSslPolicyAttributes) CreationTimestamp() terra.StringValue {
	return terra.ReferenceAsString(crsp.ref.Append("creation_timestamp"))
}

// CustomFeatures returns a reference to field custom_features of google_compute_region_ssl_policy.
func (crsp computeRegionSslPolicyAttributes) CustomFeatures() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](crsp.ref.Append("custom_features"))
}

// Description returns a reference to field description of google_compute_region_ssl_policy.
func (crsp computeRegionSslPolicyAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(crsp.ref.Append("description"))
}

// EnabledFeatures returns a reference to field enabled_features of google_compute_region_ssl_policy.
func (crsp computeRegionSslPolicyAttributes) EnabledFeatures() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](crsp.ref.Append("enabled_features"))
}

// Fingerprint returns a reference to field fingerprint of google_compute_region_ssl_policy.
func (crsp computeRegionSslPolicyAttributes) Fingerprint() terra.StringValue {
	return terra.ReferenceAsString(crsp.ref.Append("fingerprint"))
}

// Id returns a reference to field id of google_compute_region_ssl_policy.
func (crsp computeRegionSslPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(crsp.ref.Append("id"))
}

// MinTlsVersion returns a reference to field min_tls_version of google_compute_region_ssl_policy.
func (crsp computeRegionSslPolicyAttributes) MinTlsVersion() terra.StringValue {
	return terra.ReferenceAsString(crsp.ref.Append("min_tls_version"))
}

// Name returns a reference to field name of google_compute_region_ssl_policy.
func (crsp computeRegionSslPolicyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(crsp.ref.Append("name"))
}

// Profile returns a reference to field profile of google_compute_region_ssl_policy.
func (crsp computeRegionSslPolicyAttributes) Profile() terra.StringValue {
	return terra.ReferenceAsString(crsp.ref.Append("profile"))
}

// Project returns a reference to field project of google_compute_region_ssl_policy.
func (crsp computeRegionSslPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(crsp.ref.Append("project"))
}

// Region returns a reference to field region of google_compute_region_ssl_policy.
func (crsp computeRegionSslPolicyAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(crsp.ref.Append("region"))
}

// SelfLink returns a reference to field self_link of google_compute_region_ssl_policy.
func (crsp computeRegionSslPolicyAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(crsp.ref.Append("self_link"))
}

func (crsp computeRegionSslPolicyAttributes) Timeouts() computeregionsslpolicy.TimeoutsAttributes {
	return terra.ReferenceAsSingle[computeregionsslpolicy.TimeoutsAttributes](crsp.ref.Append("timeouts"))
}

type computeRegionSslPolicyState struct {
	CreationTimestamp string                                `json:"creation_timestamp"`
	CustomFeatures    []string                              `json:"custom_features"`
	Description       string                                `json:"description"`
	EnabledFeatures   []string                              `json:"enabled_features"`
	Fingerprint       string                                `json:"fingerprint"`
	Id                string                                `json:"id"`
	MinTlsVersion     string                                `json:"min_tls_version"`
	Name              string                                `json:"name"`
	Profile           string                                `json:"profile"`
	Project           string                                `json:"project"`
	Region            string                                `json:"region"`
	SelfLink          string                                `json:"self_link"`
	Timeouts          *computeregionsslpolicy.TimeoutsState `json:"timeouts"`
}
