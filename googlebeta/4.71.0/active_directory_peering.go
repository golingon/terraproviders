// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	activedirectorypeering "github.com/golingon/terraproviders/googlebeta/4.71.0/activedirectorypeering"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewActiveDirectoryPeering creates a new instance of [ActiveDirectoryPeering].
func NewActiveDirectoryPeering(name string, args ActiveDirectoryPeeringArgs) *ActiveDirectoryPeering {
	return &ActiveDirectoryPeering{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ActiveDirectoryPeering)(nil)

// ActiveDirectoryPeering represents the Terraform resource google_active_directory_peering.
type ActiveDirectoryPeering struct {
	Name      string
	Args      ActiveDirectoryPeeringArgs
	state     *activeDirectoryPeeringState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ActiveDirectoryPeering].
func (adp *ActiveDirectoryPeering) Type() string {
	return "google_active_directory_peering"
}

// LocalName returns the local name for [ActiveDirectoryPeering].
func (adp *ActiveDirectoryPeering) LocalName() string {
	return adp.Name
}

// Configuration returns the configuration (args) for [ActiveDirectoryPeering].
func (adp *ActiveDirectoryPeering) Configuration() interface{} {
	return adp.Args
}

// DependOn is used for other resources to depend on [ActiveDirectoryPeering].
func (adp *ActiveDirectoryPeering) DependOn() terra.Reference {
	return terra.ReferenceResource(adp)
}

// Dependencies returns the list of resources [ActiveDirectoryPeering] depends_on.
func (adp *ActiveDirectoryPeering) Dependencies() terra.Dependencies {
	return adp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ActiveDirectoryPeering].
func (adp *ActiveDirectoryPeering) LifecycleManagement() *terra.Lifecycle {
	return adp.Lifecycle
}

// Attributes returns the attributes for [ActiveDirectoryPeering].
func (adp *ActiveDirectoryPeering) Attributes() activeDirectoryPeeringAttributes {
	return activeDirectoryPeeringAttributes{ref: terra.ReferenceResource(adp)}
}

// ImportState imports the given attribute values into [ActiveDirectoryPeering]'s state.
func (adp *ActiveDirectoryPeering) ImportState(av io.Reader) error {
	adp.state = &activeDirectoryPeeringState{}
	if err := json.NewDecoder(av).Decode(adp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", adp.Type(), adp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ActiveDirectoryPeering] has state.
func (adp *ActiveDirectoryPeering) State() (*activeDirectoryPeeringState, bool) {
	return adp.state, adp.state != nil
}

// StateMust returns the state for [ActiveDirectoryPeering]. Panics if the state is nil.
func (adp *ActiveDirectoryPeering) StateMust() *activeDirectoryPeeringState {
	if adp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", adp.Type(), adp.LocalName()))
	}
	return adp.state
}

// ActiveDirectoryPeeringArgs contains the configurations for google_active_directory_peering.
type ActiveDirectoryPeeringArgs struct {
	// AuthorizedNetwork: string, required
	AuthorizedNetwork terra.StringValue `hcl:"authorized_network,attr" validate:"required"`
	// DomainResource: string, required
	DomainResource terra.StringValue `hcl:"domain_resource,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// PeeringId: string, required
	PeeringId terra.StringValue `hcl:"peering_id,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Status: string, optional
	Status terra.StringValue `hcl:"status,attr"`
	// StatusMessage: string, optional
	StatusMessage terra.StringValue `hcl:"status_message,attr"`
	// Timeouts: optional
	Timeouts *activedirectorypeering.Timeouts `hcl:"timeouts,block"`
}
type activeDirectoryPeeringAttributes struct {
	ref terra.Reference
}

// AuthorizedNetwork returns a reference to field authorized_network of google_active_directory_peering.
func (adp activeDirectoryPeeringAttributes) AuthorizedNetwork() terra.StringValue {
	return terra.ReferenceAsString(adp.ref.Append("authorized_network"))
}

// DomainResource returns a reference to field domain_resource of google_active_directory_peering.
func (adp activeDirectoryPeeringAttributes) DomainResource() terra.StringValue {
	return terra.ReferenceAsString(adp.ref.Append("domain_resource"))
}

// Id returns a reference to field id of google_active_directory_peering.
func (adp activeDirectoryPeeringAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(adp.ref.Append("id"))
}

// Labels returns a reference to field labels of google_active_directory_peering.
func (adp activeDirectoryPeeringAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](adp.ref.Append("labels"))
}

// Name returns a reference to field name of google_active_directory_peering.
func (adp activeDirectoryPeeringAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(adp.ref.Append("name"))
}

// PeeringId returns a reference to field peering_id of google_active_directory_peering.
func (adp activeDirectoryPeeringAttributes) PeeringId() terra.StringValue {
	return terra.ReferenceAsString(adp.ref.Append("peering_id"))
}

// Project returns a reference to field project of google_active_directory_peering.
func (adp activeDirectoryPeeringAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(adp.ref.Append("project"))
}

// Status returns a reference to field status of google_active_directory_peering.
func (adp activeDirectoryPeeringAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(adp.ref.Append("status"))
}

// StatusMessage returns a reference to field status_message of google_active_directory_peering.
func (adp activeDirectoryPeeringAttributes) StatusMessage() terra.StringValue {
	return terra.ReferenceAsString(adp.ref.Append("status_message"))
}

func (adp activeDirectoryPeeringAttributes) Timeouts() activedirectorypeering.TimeoutsAttributes {
	return terra.ReferenceAsSingle[activedirectorypeering.TimeoutsAttributes](adp.ref.Append("timeouts"))
}

type activeDirectoryPeeringState struct {
	AuthorizedNetwork string                                `json:"authorized_network"`
	DomainResource    string                                `json:"domain_resource"`
	Id                string                                `json:"id"`
	Labels            map[string]string                     `json:"labels"`
	Name              string                                `json:"name"`
	PeeringId         string                                `json:"peering_id"`
	Project           string                                `json:"project"`
	Status            string                                `json:"status"`
	StatusMessage     string                                `json:"status_message"`
	Timeouts          *activedirectorypeering.TimeoutsState `json:"timeouts"`
}
