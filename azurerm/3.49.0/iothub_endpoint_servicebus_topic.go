// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	iothubendpointservicebustopic "github.com/golingon/terraproviders/azurerm/3.49.0/iothubendpointservicebustopic"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewIothubEndpointServicebusTopic creates a new instance of [IothubEndpointServicebusTopic].
func NewIothubEndpointServicebusTopic(name string, args IothubEndpointServicebusTopicArgs) *IothubEndpointServicebusTopic {
	return &IothubEndpointServicebusTopic{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*IothubEndpointServicebusTopic)(nil)

// IothubEndpointServicebusTopic represents the Terraform resource azurerm_iothub_endpoint_servicebus_topic.
type IothubEndpointServicebusTopic struct {
	Name      string
	Args      IothubEndpointServicebusTopicArgs
	state     *iothubEndpointServicebusTopicState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [IothubEndpointServicebusTopic].
func (iest *IothubEndpointServicebusTopic) Type() string {
	return "azurerm_iothub_endpoint_servicebus_topic"
}

// LocalName returns the local name for [IothubEndpointServicebusTopic].
func (iest *IothubEndpointServicebusTopic) LocalName() string {
	return iest.Name
}

// Configuration returns the configuration (args) for [IothubEndpointServicebusTopic].
func (iest *IothubEndpointServicebusTopic) Configuration() interface{} {
	return iest.Args
}

// DependOn is used for other resources to depend on [IothubEndpointServicebusTopic].
func (iest *IothubEndpointServicebusTopic) DependOn() terra.Reference {
	return terra.ReferenceResource(iest)
}

// Dependencies returns the list of resources [IothubEndpointServicebusTopic] depends_on.
func (iest *IothubEndpointServicebusTopic) Dependencies() terra.Dependencies {
	return iest.DependsOn
}

// LifecycleManagement returns the lifecycle block for [IothubEndpointServicebusTopic].
func (iest *IothubEndpointServicebusTopic) LifecycleManagement() *terra.Lifecycle {
	return iest.Lifecycle
}

// Attributes returns the attributes for [IothubEndpointServicebusTopic].
func (iest *IothubEndpointServicebusTopic) Attributes() iothubEndpointServicebusTopicAttributes {
	return iothubEndpointServicebusTopicAttributes{ref: terra.ReferenceResource(iest)}
}

// ImportState imports the given attribute values into [IothubEndpointServicebusTopic]'s state.
func (iest *IothubEndpointServicebusTopic) ImportState(av io.Reader) error {
	iest.state = &iothubEndpointServicebusTopicState{}
	if err := json.NewDecoder(av).Decode(iest.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", iest.Type(), iest.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [IothubEndpointServicebusTopic] has state.
func (iest *IothubEndpointServicebusTopic) State() (*iothubEndpointServicebusTopicState, bool) {
	return iest.state, iest.state != nil
}

// StateMust returns the state for [IothubEndpointServicebusTopic]. Panics if the state is nil.
func (iest *IothubEndpointServicebusTopic) StateMust() *iothubEndpointServicebusTopicState {
	if iest.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", iest.Type(), iest.LocalName()))
	}
	return iest.state
}

// IothubEndpointServicebusTopicArgs contains the configurations for azurerm_iothub_endpoint_servicebus_topic.
type IothubEndpointServicebusTopicArgs struct {
	// AuthenticationType: string, optional
	AuthenticationType terra.StringValue `hcl:"authentication_type,attr"`
	// ConnectionString: string, optional
	ConnectionString terra.StringValue `hcl:"connection_string,attr"`
	// EndpointUri: string, optional
	EndpointUri terra.StringValue `hcl:"endpoint_uri,attr"`
	// EntityPath: string, optional
	EntityPath terra.StringValue `hcl:"entity_path,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IdentityId: string, optional
	IdentityId terra.StringValue `hcl:"identity_id,attr"`
	// IothubId: string, required
	IothubId terra.StringValue `hcl:"iothub_id,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *iothubendpointservicebustopic.Timeouts `hcl:"timeouts,block"`
}
type iothubEndpointServicebusTopicAttributes struct {
	ref terra.Reference
}

// AuthenticationType returns a reference to field authentication_type of azurerm_iothub_endpoint_servicebus_topic.
func (iest iothubEndpointServicebusTopicAttributes) AuthenticationType() terra.StringValue {
	return terra.ReferenceAsString(iest.ref.Append("authentication_type"))
}

// ConnectionString returns a reference to field connection_string of azurerm_iothub_endpoint_servicebus_topic.
func (iest iothubEndpointServicebusTopicAttributes) ConnectionString() terra.StringValue {
	return terra.ReferenceAsString(iest.ref.Append("connection_string"))
}

// EndpointUri returns a reference to field endpoint_uri of azurerm_iothub_endpoint_servicebus_topic.
func (iest iothubEndpointServicebusTopicAttributes) EndpointUri() terra.StringValue {
	return terra.ReferenceAsString(iest.ref.Append("endpoint_uri"))
}

// EntityPath returns a reference to field entity_path of azurerm_iothub_endpoint_servicebus_topic.
func (iest iothubEndpointServicebusTopicAttributes) EntityPath() terra.StringValue {
	return terra.ReferenceAsString(iest.ref.Append("entity_path"))
}

// Id returns a reference to field id of azurerm_iothub_endpoint_servicebus_topic.
func (iest iothubEndpointServicebusTopicAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(iest.ref.Append("id"))
}

// IdentityId returns a reference to field identity_id of azurerm_iothub_endpoint_servicebus_topic.
func (iest iothubEndpointServicebusTopicAttributes) IdentityId() terra.StringValue {
	return terra.ReferenceAsString(iest.ref.Append("identity_id"))
}

// IothubId returns a reference to field iothub_id of azurerm_iothub_endpoint_servicebus_topic.
func (iest iothubEndpointServicebusTopicAttributes) IothubId() terra.StringValue {
	return terra.ReferenceAsString(iest.ref.Append("iothub_id"))
}

// Name returns a reference to field name of azurerm_iothub_endpoint_servicebus_topic.
func (iest iothubEndpointServicebusTopicAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(iest.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_iothub_endpoint_servicebus_topic.
func (iest iothubEndpointServicebusTopicAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(iest.ref.Append("resource_group_name"))
}

func (iest iothubEndpointServicebusTopicAttributes) Timeouts() iothubendpointservicebustopic.TimeoutsAttributes {
	return terra.ReferenceAsSingle[iothubendpointservicebustopic.TimeoutsAttributes](iest.ref.Append("timeouts"))
}

type iothubEndpointServicebusTopicState struct {
	AuthenticationType string                                       `json:"authentication_type"`
	ConnectionString   string                                       `json:"connection_string"`
	EndpointUri        string                                       `json:"endpoint_uri"`
	EntityPath         string                                       `json:"entity_path"`
	Id                 string                                       `json:"id"`
	IdentityId         string                                       `json:"identity_id"`
	IothubId           string                                       `json:"iothub_id"`
	Name               string                                       `json:"name"`
	ResourceGroupName  string                                       `json:"resource_group_name"`
	Timeouts           *iothubendpointservicebustopic.TimeoutsState `json:"timeouts"`
}
