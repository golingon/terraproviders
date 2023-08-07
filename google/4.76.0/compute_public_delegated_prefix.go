// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	computepublicdelegatedprefix "github.com/golingon/terraproviders/google/4.76.0/computepublicdelegatedprefix"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewComputePublicDelegatedPrefix creates a new instance of [ComputePublicDelegatedPrefix].
func NewComputePublicDelegatedPrefix(name string, args ComputePublicDelegatedPrefixArgs) *ComputePublicDelegatedPrefix {
	return &ComputePublicDelegatedPrefix{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputePublicDelegatedPrefix)(nil)

// ComputePublicDelegatedPrefix represents the Terraform resource google_compute_public_delegated_prefix.
type ComputePublicDelegatedPrefix struct {
	Name      string
	Args      ComputePublicDelegatedPrefixArgs
	state     *computePublicDelegatedPrefixState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ComputePublicDelegatedPrefix].
func (cpdp *ComputePublicDelegatedPrefix) Type() string {
	return "google_compute_public_delegated_prefix"
}

// LocalName returns the local name for [ComputePublicDelegatedPrefix].
func (cpdp *ComputePublicDelegatedPrefix) LocalName() string {
	return cpdp.Name
}

// Configuration returns the configuration (args) for [ComputePublicDelegatedPrefix].
func (cpdp *ComputePublicDelegatedPrefix) Configuration() interface{} {
	return cpdp.Args
}

// DependOn is used for other resources to depend on [ComputePublicDelegatedPrefix].
func (cpdp *ComputePublicDelegatedPrefix) DependOn() terra.Reference {
	return terra.ReferenceResource(cpdp)
}

// Dependencies returns the list of resources [ComputePublicDelegatedPrefix] depends_on.
func (cpdp *ComputePublicDelegatedPrefix) Dependencies() terra.Dependencies {
	return cpdp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ComputePublicDelegatedPrefix].
func (cpdp *ComputePublicDelegatedPrefix) LifecycleManagement() *terra.Lifecycle {
	return cpdp.Lifecycle
}

// Attributes returns the attributes for [ComputePublicDelegatedPrefix].
func (cpdp *ComputePublicDelegatedPrefix) Attributes() computePublicDelegatedPrefixAttributes {
	return computePublicDelegatedPrefixAttributes{ref: terra.ReferenceResource(cpdp)}
}

// ImportState imports the given attribute values into [ComputePublicDelegatedPrefix]'s state.
func (cpdp *ComputePublicDelegatedPrefix) ImportState(av io.Reader) error {
	cpdp.state = &computePublicDelegatedPrefixState{}
	if err := json.NewDecoder(av).Decode(cpdp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cpdp.Type(), cpdp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ComputePublicDelegatedPrefix] has state.
func (cpdp *ComputePublicDelegatedPrefix) State() (*computePublicDelegatedPrefixState, bool) {
	return cpdp.state, cpdp.state != nil
}

// StateMust returns the state for [ComputePublicDelegatedPrefix]. Panics if the state is nil.
func (cpdp *ComputePublicDelegatedPrefix) StateMust() *computePublicDelegatedPrefixState {
	if cpdp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cpdp.Type(), cpdp.LocalName()))
	}
	return cpdp.state
}

// ComputePublicDelegatedPrefixArgs contains the configurations for google_compute_public_delegated_prefix.
type ComputePublicDelegatedPrefixArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IpCidrRange: string, required
	IpCidrRange terra.StringValue `hcl:"ip_cidr_range,attr" validate:"required"`
	// IsLiveMigration: bool, optional
	IsLiveMigration terra.BoolValue `hcl:"is_live_migration,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ParentPrefix: string, required
	ParentPrefix terra.StringValue `hcl:"parent_prefix,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Region: string, required
	Region terra.StringValue `hcl:"region,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *computepublicdelegatedprefix.Timeouts `hcl:"timeouts,block"`
}
type computePublicDelegatedPrefixAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of google_compute_public_delegated_prefix.
func (cpdp computePublicDelegatedPrefixAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(cpdp.ref.Append("description"))
}

// Id returns a reference to field id of google_compute_public_delegated_prefix.
func (cpdp computePublicDelegatedPrefixAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cpdp.ref.Append("id"))
}

// IpCidrRange returns a reference to field ip_cidr_range of google_compute_public_delegated_prefix.
func (cpdp computePublicDelegatedPrefixAttributes) IpCidrRange() terra.StringValue {
	return terra.ReferenceAsString(cpdp.ref.Append("ip_cidr_range"))
}

// IsLiveMigration returns a reference to field is_live_migration of google_compute_public_delegated_prefix.
func (cpdp computePublicDelegatedPrefixAttributes) IsLiveMigration() terra.BoolValue {
	return terra.ReferenceAsBool(cpdp.ref.Append("is_live_migration"))
}

// Name returns a reference to field name of google_compute_public_delegated_prefix.
func (cpdp computePublicDelegatedPrefixAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cpdp.ref.Append("name"))
}

// ParentPrefix returns a reference to field parent_prefix of google_compute_public_delegated_prefix.
func (cpdp computePublicDelegatedPrefixAttributes) ParentPrefix() terra.StringValue {
	return terra.ReferenceAsString(cpdp.ref.Append("parent_prefix"))
}

// Project returns a reference to field project of google_compute_public_delegated_prefix.
func (cpdp computePublicDelegatedPrefixAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(cpdp.ref.Append("project"))
}

// Region returns a reference to field region of google_compute_public_delegated_prefix.
func (cpdp computePublicDelegatedPrefixAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(cpdp.ref.Append("region"))
}

// SelfLink returns a reference to field self_link of google_compute_public_delegated_prefix.
func (cpdp computePublicDelegatedPrefixAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(cpdp.ref.Append("self_link"))
}

func (cpdp computePublicDelegatedPrefixAttributes) Timeouts() computepublicdelegatedprefix.TimeoutsAttributes {
	return terra.ReferenceAsSingle[computepublicdelegatedprefix.TimeoutsAttributes](cpdp.ref.Append("timeouts"))
}

type computePublicDelegatedPrefixState struct {
	Description     string                                      `json:"description"`
	Id              string                                      `json:"id"`
	IpCidrRange     string                                      `json:"ip_cidr_range"`
	IsLiveMigration bool                                        `json:"is_live_migration"`
	Name            string                                      `json:"name"`
	ParentPrefix    string                                      `json:"parent_prefix"`
	Project         string                                      `json:"project"`
	Region          string                                      `json:"region"`
	SelfLink        string                                      `json:"self_link"`
	Timeouts        *computepublicdelegatedprefix.TimeoutsState `json:"timeouts"`
}
