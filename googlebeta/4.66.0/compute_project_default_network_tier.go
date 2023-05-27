// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	computeprojectdefaultnetworktier "github.com/golingon/terraproviders/googlebeta/4.66.0/computeprojectdefaultnetworktier"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewComputeProjectDefaultNetworkTier creates a new instance of [ComputeProjectDefaultNetworkTier].
func NewComputeProjectDefaultNetworkTier(name string, args ComputeProjectDefaultNetworkTierArgs) *ComputeProjectDefaultNetworkTier {
	return &ComputeProjectDefaultNetworkTier{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeProjectDefaultNetworkTier)(nil)

// ComputeProjectDefaultNetworkTier represents the Terraform resource google_compute_project_default_network_tier.
type ComputeProjectDefaultNetworkTier struct {
	Name      string
	Args      ComputeProjectDefaultNetworkTierArgs
	state     *computeProjectDefaultNetworkTierState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ComputeProjectDefaultNetworkTier].
func (cpdnt *ComputeProjectDefaultNetworkTier) Type() string {
	return "google_compute_project_default_network_tier"
}

// LocalName returns the local name for [ComputeProjectDefaultNetworkTier].
func (cpdnt *ComputeProjectDefaultNetworkTier) LocalName() string {
	return cpdnt.Name
}

// Configuration returns the configuration (args) for [ComputeProjectDefaultNetworkTier].
func (cpdnt *ComputeProjectDefaultNetworkTier) Configuration() interface{} {
	return cpdnt.Args
}

// DependOn is used for other resources to depend on [ComputeProjectDefaultNetworkTier].
func (cpdnt *ComputeProjectDefaultNetworkTier) DependOn() terra.Reference {
	return terra.ReferenceResource(cpdnt)
}

// Dependencies returns the list of resources [ComputeProjectDefaultNetworkTier] depends_on.
func (cpdnt *ComputeProjectDefaultNetworkTier) Dependencies() terra.Dependencies {
	return cpdnt.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ComputeProjectDefaultNetworkTier].
func (cpdnt *ComputeProjectDefaultNetworkTier) LifecycleManagement() *terra.Lifecycle {
	return cpdnt.Lifecycle
}

// Attributes returns the attributes for [ComputeProjectDefaultNetworkTier].
func (cpdnt *ComputeProjectDefaultNetworkTier) Attributes() computeProjectDefaultNetworkTierAttributes {
	return computeProjectDefaultNetworkTierAttributes{ref: terra.ReferenceResource(cpdnt)}
}

// ImportState imports the given attribute values into [ComputeProjectDefaultNetworkTier]'s state.
func (cpdnt *ComputeProjectDefaultNetworkTier) ImportState(av io.Reader) error {
	cpdnt.state = &computeProjectDefaultNetworkTierState{}
	if err := json.NewDecoder(av).Decode(cpdnt.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cpdnt.Type(), cpdnt.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ComputeProjectDefaultNetworkTier] has state.
func (cpdnt *ComputeProjectDefaultNetworkTier) State() (*computeProjectDefaultNetworkTierState, bool) {
	return cpdnt.state, cpdnt.state != nil
}

// StateMust returns the state for [ComputeProjectDefaultNetworkTier]. Panics if the state is nil.
func (cpdnt *ComputeProjectDefaultNetworkTier) StateMust() *computeProjectDefaultNetworkTierState {
	if cpdnt.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cpdnt.Type(), cpdnt.LocalName()))
	}
	return cpdnt.state
}

// ComputeProjectDefaultNetworkTierArgs contains the configurations for google_compute_project_default_network_tier.
type ComputeProjectDefaultNetworkTierArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// NetworkTier: string, required
	NetworkTier terra.StringValue `hcl:"network_tier,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Timeouts: optional
	Timeouts *computeprojectdefaultnetworktier.Timeouts `hcl:"timeouts,block"`
}
type computeProjectDefaultNetworkTierAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of google_compute_project_default_network_tier.
func (cpdnt computeProjectDefaultNetworkTierAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cpdnt.ref.Append("id"))
}

// NetworkTier returns a reference to field network_tier of google_compute_project_default_network_tier.
func (cpdnt computeProjectDefaultNetworkTierAttributes) NetworkTier() terra.StringValue {
	return terra.ReferenceAsString(cpdnt.ref.Append("network_tier"))
}

// Project returns a reference to field project of google_compute_project_default_network_tier.
func (cpdnt computeProjectDefaultNetworkTierAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(cpdnt.ref.Append("project"))
}

func (cpdnt computeProjectDefaultNetworkTierAttributes) Timeouts() computeprojectdefaultnetworktier.TimeoutsAttributes {
	return terra.ReferenceAsSingle[computeprojectdefaultnetworktier.TimeoutsAttributes](cpdnt.ref.Append("timeouts"))
}

type computeProjectDefaultNetworkTierState struct {
	Id          string                                          `json:"id"`
	NetworkTier string                                          `json:"network_tier"`
	Project     string                                          `json:"project"`
	Timeouts    *computeprojectdefaultnetworktier.TimeoutsState `json:"timeouts"`
}
