// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	computesslpolicy "github.com/golingon/terraproviders/google/4.63.1/computesslpolicy"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewComputeSslPolicy creates a new instance of [ComputeSslPolicy].
func NewComputeSslPolicy(name string, args ComputeSslPolicyArgs) *ComputeSslPolicy {
	return &ComputeSslPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeSslPolicy)(nil)

// ComputeSslPolicy represents the Terraform resource google_compute_ssl_policy.
type ComputeSslPolicy struct {
	Name      string
	Args      ComputeSslPolicyArgs
	state     *computeSslPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ComputeSslPolicy].
func (csp *ComputeSslPolicy) Type() string {
	return "google_compute_ssl_policy"
}

// LocalName returns the local name for [ComputeSslPolicy].
func (csp *ComputeSslPolicy) LocalName() string {
	return csp.Name
}

// Configuration returns the configuration (args) for [ComputeSslPolicy].
func (csp *ComputeSslPolicy) Configuration() interface{} {
	return csp.Args
}

// DependOn is used for other resources to depend on [ComputeSslPolicy].
func (csp *ComputeSslPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(csp)
}

// Dependencies returns the list of resources [ComputeSslPolicy] depends_on.
func (csp *ComputeSslPolicy) Dependencies() terra.Dependencies {
	return csp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ComputeSslPolicy].
func (csp *ComputeSslPolicy) LifecycleManagement() *terra.Lifecycle {
	return csp.Lifecycle
}

// Attributes returns the attributes for [ComputeSslPolicy].
func (csp *ComputeSslPolicy) Attributes() computeSslPolicyAttributes {
	return computeSslPolicyAttributes{ref: terra.ReferenceResource(csp)}
}

// ImportState imports the given attribute values into [ComputeSslPolicy]'s state.
func (csp *ComputeSslPolicy) ImportState(av io.Reader) error {
	csp.state = &computeSslPolicyState{}
	if err := json.NewDecoder(av).Decode(csp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", csp.Type(), csp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ComputeSslPolicy] has state.
func (csp *ComputeSslPolicy) State() (*computeSslPolicyState, bool) {
	return csp.state, csp.state != nil
}

// StateMust returns the state for [ComputeSslPolicy]. Panics if the state is nil.
func (csp *ComputeSslPolicy) StateMust() *computeSslPolicyState {
	if csp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", csp.Type(), csp.LocalName()))
	}
	return csp.state
}

// ComputeSslPolicyArgs contains the configurations for google_compute_ssl_policy.
type ComputeSslPolicyArgs struct {
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
	// Timeouts: optional
	Timeouts *computesslpolicy.Timeouts `hcl:"timeouts,block"`
}
type computeSslPolicyAttributes struct {
	ref terra.Reference
}

// CreationTimestamp returns a reference to field creation_timestamp of google_compute_ssl_policy.
func (csp computeSslPolicyAttributes) CreationTimestamp() terra.StringValue {
	return terra.ReferenceAsString(csp.ref.Append("creation_timestamp"))
}

// CustomFeatures returns a reference to field custom_features of google_compute_ssl_policy.
func (csp computeSslPolicyAttributes) CustomFeatures() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](csp.ref.Append("custom_features"))
}

// Description returns a reference to field description of google_compute_ssl_policy.
func (csp computeSslPolicyAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(csp.ref.Append("description"))
}

// EnabledFeatures returns a reference to field enabled_features of google_compute_ssl_policy.
func (csp computeSslPolicyAttributes) EnabledFeatures() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](csp.ref.Append("enabled_features"))
}

// Fingerprint returns a reference to field fingerprint of google_compute_ssl_policy.
func (csp computeSslPolicyAttributes) Fingerprint() terra.StringValue {
	return terra.ReferenceAsString(csp.ref.Append("fingerprint"))
}

// Id returns a reference to field id of google_compute_ssl_policy.
func (csp computeSslPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(csp.ref.Append("id"))
}

// MinTlsVersion returns a reference to field min_tls_version of google_compute_ssl_policy.
func (csp computeSslPolicyAttributes) MinTlsVersion() terra.StringValue {
	return terra.ReferenceAsString(csp.ref.Append("min_tls_version"))
}

// Name returns a reference to field name of google_compute_ssl_policy.
func (csp computeSslPolicyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(csp.ref.Append("name"))
}

// Profile returns a reference to field profile of google_compute_ssl_policy.
func (csp computeSslPolicyAttributes) Profile() terra.StringValue {
	return terra.ReferenceAsString(csp.ref.Append("profile"))
}

// Project returns a reference to field project of google_compute_ssl_policy.
func (csp computeSslPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(csp.ref.Append("project"))
}

// SelfLink returns a reference to field self_link of google_compute_ssl_policy.
func (csp computeSslPolicyAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(csp.ref.Append("self_link"))
}

func (csp computeSslPolicyAttributes) Timeouts() computesslpolicy.TimeoutsAttributes {
	return terra.ReferenceAsSingle[computesslpolicy.TimeoutsAttributes](csp.ref.Append("timeouts"))
}

type computeSslPolicyState struct {
	CreationTimestamp string                          `json:"creation_timestamp"`
	CustomFeatures    []string                        `json:"custom_features"`
	Description       string                          `json:"description"`
	EnabledFeatures   []string                        `json:"enabled_features"`
	Fingerprint       string                          `json:"fingerprint"`
	Id                string                          `json:"id"`
	MinTlsVersion     string                          `json:"min_tls_version"`
	Name              string                          `json:"name"`
	Profile           string                          `json:"profile"`
	Project           string                          `json:"project"`
	SelfLink          string                          `json:"self_link"`
	Timeouts          *computesslpolicy.TimeoutsState `json:"timeouts"`
}
