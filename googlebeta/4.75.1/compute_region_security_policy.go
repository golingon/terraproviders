// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	computeregionsecuritypolicy "github.com/golingon/terraproviders/googlebeta/4.75.1/computeregionsecuritypolicy"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewComputeRegionSecurityPolicy creates a new instance of [ComputeRegionSecurityPolicy].
func NewComputeRegionSecurityPolicy(name string, args ComputeRegionSecurityPolicyArgs) *ComputeRegionSecurityPolicy {
	return &ComputeRegionSecurityPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeRegionSecurityPolicy)(nil)

// ComputeRegionSecurityPolicy represents the Terraform resource google_compute_region_security_policy.
type ComputeRegionSecurityPolicy struct {
	Name      string
	Args      ComputeRegionSecurityPolicyArgs
	state     *computeRegionSecurityPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ComputeRegionSecurityPolicy].
func (crsp *ComputeRegionSecurityPolicy) Type() string {
	return "google_compute_region_security_policy"
}

// LocalName returns the local name for [ComputeRegionSecurityPolicy].
func (crsp *ComputeRegionSecurityPolicy) LocalName() string {
	return crsp.Name
}

// Configuration returns the configuration (args) for [ComputeRegionSecurityPolicy].
func (crsp *ComputeRegionSecurityPolicy) Configuration() interface{} {
	return crsp.Args
}

// DependOn is used for other resources to depend on [ComputeRegionSecurityPolicy].
func (crsp *ComputeRegionSecurityPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(crsp)
}

// Dependencies returns the list of resources [ComputeRegionSecurityPolicy] depends_on.
func (crsp *ComputeRegionSecurityPolicy) Dependencies() terra.Dependencies {
	return crsp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ComputeRegionSecurityPolicy].
func (crsp *ComputeRegionSecurityPolicy) LifecycleManagement() *terra.Lifecycle {
	return crsp.Lifecycle
}

// Attributes returns the attributes for [ComputeRegionSecurityPolicy].
func (crsp *ComputeRegionSecurityPolicy) Attributes() computeRegionSecurityPolicyAttributes {
	return computeRegionSecurityPolicyAttributes{ref: terra.ReferenceResource(crsp)}
}

// ImportState imports the given attribute values into [ComputeRegionSecurityPolicy]'s state.
func (crsp *ComputeRegionSecurityPolicy) ImportState(av io.Reader) error {
	crsp.state = &computeRegionSecurityPolicyState{}
	if err := json.NewDecoder(av).Decode(crsp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", crsp.Type(), crsp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ComputeRegionSecurityPolicy] has state.
func (crsp *ComputeRegionSecurityPolicy) State() (*computeRegionSecurityPolicyState, bool) {
	return crsp.state, crsp.state != nil
}

// StateMust returns the state for [ComputeRegionSecurityPolicy]. Panics if the state is nil.
func (crsp *ComputeRegionSecurityPolicy) StateMust() *computeRegionSecurityPolicyState {
	if crsp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", crsp.Type(), crsp.LocalName()))
	}
	return crsp.state
}

// ComputeRegionSecurityPolicyArgs contains the configurations for google_compute_region_security_policy.
type ComputeRegionSecurityPolicyArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
	// Type: string, optional
	Type terra.StringValue `hcl:"type,attr"`
	// DdosProtectionConfig: optional
	DdosProtectionConfig *computeregionsecuritypolicy.DdosProtectionConfig `hcl:"ddos_protection_config,block"`
	// Timeouts: optional
	Timeouts *computeregionsecuritypolicy.Timeouts `hcl:"timeouts,block"`
}
type computeRegionSecurityPolicyAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of google_compute_region_security_policy.
func (crsp computeRegionSecurityPolicyAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(crsp.ref.Append("description"))
}

// Fingerprint returns a reference to field fingerprint of google_compute_region_security_policy.
func (crsp computeRegionSecurityPolicyAttributes) Fingerprint() terra.StringValue {
	return terra.ReferenceAsString(crsp.ref.Append("fingerprint"))
}

// Id returns a reference to field id of google_compute_region_security_policy.
func (crsp computeRegionSecurityPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(crsp.ref.Append("id"))
}

// Name returns a reference to field name of google_compute_region_security_policy.
func (crsp computeRegionSecurityPolicyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(crsp.ref.Append("name"))
}

// PolicyId returns a reference to field policy_id of google_compute_region_security_policy.
func (crsp computeRegionSecurityPolicyAttributes) PolicyId() terra.StringValue {
	return terra.ReferenceAsString(crsp.ref.Append("policy_id"))
}

// Project returns a reference to field project of google_compute_region_security_policy.
func (crsp computeRegionSecurityPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(crsp.ref.Append("project"))
}

// Region returns a reference to field region of google_compute_region_security_policy.
func (crsp computeRegionSecurityPolicyAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(crsp.ref.Append("region"))
}

// SelfLink returns a reference to field self_link of google_compute_region_security_policy.
func (crsp computeRegionSecurityPolicyAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(crsp.ref.Append("self_link"))
}

// SelfLinkWithPolicyId returns a reference to field self_link_with_policy_id of google_compute_region_security_policy.
func (crsp computeRegionSecurityPolicyAttributes) SelfLinkWithPolicyId() terra.StringValue {
	return terra.ReferenceAsString(crsp.ref.Append("self_link_with_policy_id"))
}

// Type returns a reference to field type of google_compute_region_security_policy.
func (crsp computeRegionSecurityPolicyAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(crsp.ref.Append("type"))
}

func (crsp computeRegionSecurityPolicyAttributes) DdosProtectionConfig() terra.ListValue[computeregionsecuritypolicy.DdosProtectionConfigAttributes] {
	return terra.ReferenceAsList[computeregionsecuritypolicy.DdosProtectionConfigAttributes](crsp.ref.Append("ddos_protection_config"))
}

func (crsp computeRegionSecurityPolicyAttributes) Timeouts() computeregionsecuritypolicy.TimeoutsAttributes {
	return terra.ReferenceAsSingle[computeregionsecuritypolicy.TimeoutsAttributes](crsp.ref.Append("timeouts"))
}

type computeRegionSecurityPolicyState struct {
	Description          string                                                  `json:"description"`
	Fingerprint          string                                                  `json:"fingerprint"`
	Id                   string                                                  `json:"id"`
	Name                 string                                                  `json:"name"`
	PolicyId             string                                                  `json:"policy_id"`
	Project              string                                                  `json:"project"`
	Region               string                                                  `json:"region"`
	SelfLink             string                                                  `json:"self_link"`
	SelfLinkWithPolicyId string                                                  `json:"self_link_with_policy_id"`
	Type                 string                                                  `json:"type"`
	DdosProtectionConfig []computeregionsecuritypolicy.DdosProtectionConfigState `json:"ddos_protection_config"`
	Timeouts             *computeregionsecuritypolicy.TimeoutsState              `json:"timeouts"`
}
