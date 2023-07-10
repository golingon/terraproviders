// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	emailcommunicationservice "github.com/golingon/terraproviders/azurerm/3.64.0/emailcommunicationservice"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewEmailCommunicationService creates a new instance of [EmailCommunicationService].
func NewEmailCommunicationService(name string, args EmailCommunicationServiceArgs) *EmailCommunicationService {
	return &EmailCommunicationService{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*EmailCommunicationService)(nil)

// EmailCommunicationService represents the Terraform resource azurerm_email_communication_service.
type EmailCommunicationService struct {
	Name      string
	Args      EmailCommunicationServiceArgs
	state     *emailCommunicationServiceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [EmailCommunicationService].
func (ecs *EmailCommunicationService) Type() string {
	return "azurerm_email_communication_service"
}

// LocalName returns the local name for [EmailCommunicationService].
func (ecs *EmailCommunicationService) LocalName() string {
	return ecs.Name
}

// Configuration returns the configuration (args) for [EmailCommunicationService].
func (ecs *EmailCommunicationService) Configuration() interface{} {
	return ecs.Args
}

// DependOn is used for other resources to depend on [EmailCommunicationService].
func (ecs *EmailCommunicationService) DependOn() terra.Reference {
	return terra.ReferenceResource(ecs)
}

// Dependencies returns the list of resources [EmailCommunicationService] depends_on.
func (ecs *EmailCommunicationService) Dependencies() terra.Dependencies {
	return ecs.DependsOn
}

// LifecycleManagement returns the lifecycle block for [EmailCommunicationService].
func (ecs *EmailCommunicationService) LifecycleManagement() *terra.Lifecycle {
	return ecs.Lifecycle
}

// Attributes returns the attributes for [EmailCommunicationService].
func (ecs *EmailCommunicationService) Attributes() emailCommunicationServiceAttributes {
	return emailCommunicationServiceAttributes{ref: terra.ReferenceResource(ecs)}
}

// ImportState imports the given attribute values into [EmailCommunicationService]'s state.
func (ecs *EmailCommunicationService) ImportState(av io.Reader) error {
	ecs.state = &emailCommunicationServiceState{}
	if err := json.NewDecoder(av).Decode(ecs.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ecs.Type(), ecs.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [EmailCommunicationService] has state.
func (ecs *EmailCommunicationService) State() (*emailCommunicationServiceState, bool) {
	return ecs.state, ecs.state != nil
}

// StateMust returns the state for [EmailCommunicationService]. Panics if the state is nil.
func (ecs *EmailCommunicationService) StateMust() *emailCommunicationServiceState {
	if ecs.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ecs.Type(), ecs.LocalName()))
	}
	return ecs.state
}

// EmailCommunicationServiceArgs contains the configurations for azurerm_email_communication_service.
type EmailCommunicationServiceArgs struct {
	// DataLocation: string, required
	DataLocation terra.StringValue `hcl:"data_location,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Timeouts: optional
	Timeouts *emailcommunicationservice.Timeouts `hcl:"timeouts,block"`
}
type emailCommunicationServiceAttributes struct {
	ref terra.Reference
}

// DataLocation returns a reference to field data_location of azurerm_email_communication_service.
func (ecs emailCommunicationServiceAttributes) DataLocation() terra.StringValue {
	return terra.ReferenceAsString(ecs.ref.Append("data_location"))
}

// Id returns a reference to field id of azurerm_email_communication_service.
func (ecs emailCommunicationServiceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ecs.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_email_communication_service.
func (ecs emailCommunicationServiceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ecs.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_email_communication_service.
func (ecs emailCommunicationServiceAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(ecs.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_email_communication_service.
func (ecs emailCommunicationServiceAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ecs.ref.Append("tags"))
}

func (ecs emailCommunicationServiceAttributes) Timeouts() emailcommunicationservice.TimeoutsAttributes {
	return terra.ReferenceAsSingle[emailcommunicationservice.TimeoutsAttributes](ecs.ref.Append("timeouts"))
}

type emailCommunicationServiceState struct {
	DataLocation      string                                   `json:"data_location"`
	Id                string                                   `json:"id"`
	Name              string                                   `json:"name"`
	ResourceGroupName string                                   `json:"resource_group_name"`
	Tags              map[string]string                        `json:"tags"`
	Timeouts          *emailcommunicationservice.TimeoutsState `json:"timeouts"`
}
