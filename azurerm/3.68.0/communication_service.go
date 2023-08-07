// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	communicationservice "github.com/golingon/terraproviders/azurerm/3.68.0/communicationservice"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCommunicationService creates a new instance of [CommunicationService].
func NewCommunicationService(name string, args CommunicationServiceArgs) *CommunicationService {
	return &CommunicationService{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CommunicationService)(nil)

// CommunicationService represents the Terraform resource azurerm_communication_service.
type CommunicationService struct {
	Name      string
	Args      CommunicationServiceArgs
	state     *communicationServiceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CommunicationService].
func (cs *CommunicationService) Type() string {
	return "azurerm_communication_service"
}

// LocalName returns the local name for [CommunicationService].
func (cs *CommunicationService) LocalName() string {
	return cs.Name
}

// Configuration returns the configuration (args) for [CommunicationService].
func (cs *CommunicationService) Configuration() interface{} {
	return cs.Args
}

// DependOn is used for other resources to depend on [CommunicationService].
func (cs *CommunicationService) DependOn() terra.Reference {
	return terra.ReferenceResource(cs)
}

// Dependencies returns the list of resources [CommunicationService] depends_on.
func (cs *CommunicationService) Dependencies() terra.Dependencies {
	return cs.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CommunicationService].
func (cs *CommunicationService) LifecycleManagement() *terra.Lifecycle {
	return cs.Lifecycle
}

// Attributes returns the attributes for [CommunicationService].
func (cs *CommunicationService) Attributes() communicationServiceAttributes {
	return communicationServiceAttributes{ref: terra.ReferenceResource(cs)}
}

// ImportState imports the given attribute values into [CommunicationService]'s state.
func (cs *CommunicationService) ImportState(av io.Reader) error {
	cs.state = &communicationServiceState{}
	if err := json.NewDecoder(av).Decode(cs.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cs.Type(), cs.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CommunicationService] has state.
func (cs *CommunicationService) State() (*communicationServiceState, bool) {
	return cs.state, cs.state != nil
}

// StateMust returns the state for [CommunicationService]. Panics if the state is nil.
func (cs *CommunicationService) StateMust() *communicationServiceState {
	if cs.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cs.Type(), cs.LocalName()))
	}
	return cs.state
}

// CommunicationServiceArgs contains the configurations for azurerm_communication_service.
type CommunicationServiceArgs struct {
	// DataLocation: string, optional
	DataLocation terra.StringValue `hcl:"data_location,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Timeouts: optional
	Timeouts *communicationservice.Timeouts `hcl:"timeouts,block"`
}
type communicationServiceAttributes struct {
	ref terra.Reference
}

// DataLocation returns a reference to field data_location of azurerm_communication_service.
func (cs communicationServiceAttributes) DataLocation() terra.StringValue {
	return terra.ReferenceAsString(cs.ref.Append("data_location"))
}

// Id returns a reference to field id of azurerm_communication_service.
func (cs communicationServiceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cs.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_communication_service.
func (cs communicationServiceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cs.ref.Append("name"))
}

// PrimaryConnectionString returns a reference to field primary_connection_string of azurerm_communication_service.
func (cs communicationServiceAttributes) PrimaryConnectionString() terra.StringValue {
	return terra.ReferenceAsString(cs.ref.Append("primary_connection_string"))
}

// PrimaryKey returns a reference to field primary_key of azurerm_communication_service.
func (cs communicationServiceAttributes) PrimaryKey() terra.StringValue {
	return terra.ReferenceAsString(cs.ref.Append("primary_key"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_communication_service.
func (cs communicationServiceAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(cs.ref.Append("resource_group_name"))
}

// SecondaryConnectionString returns a reference to field secondary_connection_string of azurerm_communication_service.
func (cs communicationServiceAttributes) SecondaryConnectionString() terra.StringValue {
	return terra.ReferenceAsString(cs.ref.Append("secondary_connection_string"))
}

// SecondaryKey returns a reference to field secondary_key of azurerm_communication_service.
func (cs communicationServiceAttributes) SecondaryKey() terra.StringValue {
	return terra.ReferenceAsString(cs.ref.Append("secondary_key"))
}

// Tags returns a reference to field tags of azurerm_communication_service.
func (cs communicationServiceAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cs.ref.Append("tags"))
}

func (cs communicationServiceAttributes) Timeouts() communicationservice.TimeoutsAttributes {
	return terra.ReferenceAsSingle[communicationservice.TimeoutsAttributes](cs.ref.Append("timeouts"))
}

type communicationServiceState struct {
	DataLocation              string                              `json:"data_location"`
	Id                        string                              `json:"id"`
	Name                      string                              `json:"name"`
	PrimaryConnectionString   string                              `json:"primary_connection_string"`
	PrimaryKey                string                              `json:"primary_key"`
	ResourceGroupName         string                              `json:"resource_group_name"`
	SecondaryConnectionString string                              `json:"secondary_connection_string"`
	SecondaryKey              string                              `json:"secondary_key"`
	Tags                      map[string]string                   `json:"tags"`
	Timeouts                  *communicationservice.TimeoutsState `json:"timeouts"`
}
