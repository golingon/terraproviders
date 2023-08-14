// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	synapseprivatelinkhub "github.com/golingon/terraproviders/azurerm/3.69.0/synapseprivatelinkhub"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSynapsePrivateLinkHub creates a new instance of [SynapsePrivateLinkHub].
func NewSynapsePrivateLinkHub(name string, args SynapsePrivateLinkHubArgs) *SynapsePrivateLinkHub {
	return &SynapsePrivateLinkHub{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SynapsePrivateLinkHub)(nil)

// SynapsePrivateLinkHub represents the Terraform resource azurerm_synapse_private_link_hub.
type SynapsePrivateLinkHub struct {
	Name      string
	Args      SynapsePrivateLinkHubArgs
	state     *synapsePrivateLinkHubState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SynapsePrivateLinkHub].
func (splh *SynapsePrivateLinkHub) Type() string {
	return "azurerm_synapse_private_link_hub"
}

// LocalName returns the local name for [SynapsePrivateLinkHub].
func (splh *SynapsePrivateLinkHub) LocalName() string {
	return splh.Name
}

// Configuration returns the configuration (args) for [SynapsePrivateLinkHub].
func (splh *SynapsePrivateLinkHub) Configuration() interface{} {
	return splh.Args
}

// DependOn is used for other resources to depend on [SynapsePrivateLinkHub].
func (splh *SynapsePrivateLinkHub) DependOn() terra.Reference {
	return terra.ReferenceResource(splh)
}

// Dependencies returns the list of resources [SynapsePrivateLinkHub] depends_on.
func (splh *SynapsePrivateLinkHub) Dependencies() terra.Dependencies {
	return splh.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SynapsePrivateLinkHub].
func (splh *SynapsePrivateLinkHub) LifecycleManagement() *terra.Lifecycle {
	return splh.Lifecycle
}

// Attributes returns the attributes for [SynapsePrivateLinkHub].
func (splh *SynapsePrivateLinkHub) Attributes() synapsePrivateLinkHubAttributes {
	return synapsePrivateLinkHubAttributes{ref: terra.ReferenceResource(splh)}
}

// ImportState imports the given attribute values into [SynapsePrivateLinkHub]'s state.
func (splh *SynapsePrivateLinkHub) ImportState(av io.Reader) error {
	splh.state = &synapsePrivateLinkHubState{}
	if err := json.NewDecoder(av).Decode(splh.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", splh.Type(), splh.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SynapsePrivateLinkHub] has state.
func (splh *SynapsePrivateLinkHub) State() (*synapsePrivateLinkHubState, bool) {
	return splh.state, splh.state != nil
}

// StateMust returns the state for [SynapsePrivateLinkHub]. Panics if the state is nil.
func (splh *SynapsePrivateLinkHub) StateMust() *synapsePrivateLinkHubState {
	if splh.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", splh.Type(), splh.LocalName()))
	}
	return splh.state
}

// SynapsePrivateLinkHubArgs contains the configurations for azurerm_synapse_private_link_hub.
type SynapsePrivateLinkHubArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Timeouts: optional
	Timeouts *synapseprivatelinkhub.Timeouts `hcl:"timeouts,block"`
}
type synapsePrivateLinkHubAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_synapse_private_link_hub.
func (splh synapsePrivateLinkHubAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(splh.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_synapse_private_link_hub.
func (splh synapsePrivateLinkHubAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(splh.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_synapse_private_link_hub.
func (splh synapsePrivateLinkHubAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(splh.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_synapse_private_link_hub.
func (splh synapsePrivateLinkHubAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(splh.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_synapse_private_link_hub.
func (splh synapsePrivateLinkHubAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](splh.ref.Append("tags"))
}

func (splh synapsePrivateLinkHubAttributes) Timeouts() synapseprivatelinkhub.TimeoutsAttributes {
	return terra.ReferenceAsSingle[synapseprivatelinkhub.TimeoutsAttributes](splh.ref.Append("timeouts"))
}

type synapsePrivateLinkHubState struct {
	Id                string                               `json:"id"`
	Location          string                               `json:"location"`
	Name              string                               `json:"name"`
	ResourceGroupName string                               `json:"resource_group_name"`
	Tags              map[string]string                    `json:"tags"`
	Timeouts          *synapseprivatelinkhub.TimeoutsState `json:"timeouts"`
}
