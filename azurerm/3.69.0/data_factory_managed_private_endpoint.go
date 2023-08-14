// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	datafactorymanagedprivateendpoint "github.com/golingon/terraproviders/azurerm/3.69.0/datafactorymanagedprivateendpoint"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDataFactoryManagedPrivateEndpoint creates a new instance of [DataFactoryManagedPrivateEndpoint].
func NewDataFactoryManagedPrivateEndpoint(name string, args DataFactoryManagedPrivateEndpointArgs) *DataFactoryManagedPrivateEndpoint {
	return &DataFactoryManagedPrivateEndpoint{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DataFactoryManagedPrivateEndpoint)(nil)

// DataFactoryManagedPrivateEndpoint represents the Terraform resource azurerm_data_factory_managed_private_endpoint.
type DataFactoryManagedPrivateEndpoint struct {
	Name      string
	Args      DataFactoryManagedPrivateEndpointArgs
	state     *dataFactoryManagedPrivateEndpointState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DataFactoryManagedPrivateEndpoint].
func (dfmpe *DataFactoryManagedPrivateEndpoint) Type() string {
	return "azurerm_data_factory_managed_private_endpoint"
}

// LocalName returns the local name for [DataFactoryManagedPrivateEndpoint].
func (dfmpe *DataFactoryManagedPrivateEndpoint) LocalName() string {
	return dfmpe.Name
}

// Configuration returns the configuration (args) for [DataFactoryManagedPrivateEndpoint].
func (dfmpe *DataFactoryManagedPrivateEndpoint) Configuration() interface{} {
	return dfmpe.Args
}

// DependOn is used for other resources to depend on [DataFactoryManagedPrivateEndpoint].
func (dfmpe *DataFactoryManagedPrivateEndpoint) DependOn() terra.Reference {
	return terra.ReferenceResource(dfmpe)
}

// Dependencies returns the list of resources [DataFactoryManagedPrivateEndpoint] depends_on.
func (dfmpe *DataFactoryManagedPrivateEndpoint) Dependencies() terra.Dependencies {
	return dfmpe.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DataFactoryManagedPrivateEndpoint].
func (dfmpe *DataFactoryManagedPrivateEndpoint) LifecycleManagement() *terra.Lifecycle {
	return dfmpe.Lifecycle
}

// Attributes returns the attributes for [DataFactoryManagedPrivateEndpoint].
func (dfmpe *DataFactoryManagedPrivateEndpoint) Attributes() dataFactoryManagedPrivateEndpointAttributes {
	return dataFactoryManagedPrivateEndpointAttributes{ref: terra.ReferenceResource(dfmpe)}
}

// ImportState imports the given attribute values into [DataFactoryManagedPrivateEndpoint]'s state.
func (dfmpe *DataFactoryManagedPrivateEndpoint) ImportState(av io.Reader) error {
	dfmpe.state = &dataFactoryManagedPrivateEndpointState{}
	if err := json.NewDecoder(av).Decode(dfmpe.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dfmpe.Type(), dfmpe.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DataFactoryManagedPrivateEndpoint] has state.
func (dfmpe *DataFactoryManagedPrivateEndpoint) State() (*dataFactoryManagedPrivateEndpointState, bool) {
	return dfmpe.state, dfmpe.state != nil
}

// StateMust returns the state for [DataFactoryManagedPrivateEndpoint]. Panics if the state is nil.
func (dfmpe *DataFactoryManagedPrivateEndpoint) StateMust() *dataFactoryManagedPrivateEndpointState {
	if dfmpe.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dfmpe.Type(), dfmpe.LocalName()))
	}
	return dfmpe.state
}

// DataFactoryManagedPrivateEndpointArgs contains the configurations for azurerm_data_factory_managed_private_endpoint.
type DataFactoryManagedPrivateEndpointArgs struct {
	// DataFactoryId: string, required
	DataFactoryId terra.StringValue `hcl:"data_factory_id,attr" validate:"required"`
	// Fqdns: list of string, optional
	Fqdns terra.ListValue[terra.StringValue] `hcl:"fqdns,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// SubresourceName: string, optional
	SubresourceName terra.StringValue `hcl:"subresource_name,attr"`
	// TargetResourceId: string, required
	TargetResourceId terra.StringValue `hcl:"target_resource_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *datafactorymanagedprivateendpoint.Timeouts `hcl:"timeouts,block"`
}
type dataFactoryManagedPrivateEndpointAttributes struct {
	ref terra.Reference
}

// DataFactoryId returns a reference to field data_factory_id of azurerm_data_factory_managed_private_endpoint.
func (dfmpe dataFactoryManagedPrivateEndpointAttributes) DataFactoryId() terra.StringValue {
	return terra.ReferenceAsString(dfmpe.ref.Append("data_factory_id"))
}

// Fqdns returns a reference to field fqdns of azurerm_data_factory_managed_private_endpoint.
func (dfmpe dataFactoryManagedPrivateEndpointAttributes) Fqdns() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](dfmpe.ref.Append("fqdns"))
}

// Id returns a reference to field id of azurerm_data_factory_managed_private_endpoint.
func (dfmpe dataFactoryManagedPrivateEndpointAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dfmpe.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_data_factory_managed_private_endpoint.
func (dfmpe dataFactoryManagedPrivateEndpointAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dfmpe.ref.Append("name"))
}

// SubresourceName returns a reference to field subresource_name of azurerm_data_factory_managed_private_endpoint.
func (dfmpe dataFactoryManagedPrivateEndpointAttributes) SubresourceName() terra.StringValue {
	return terra.ReferenceAsString(dfmpe.ref.Append("subresource_name"))
}

// TargetResourceId returns a reference to field target_resource_id of azurerm_data_factory_managed_private_endpoint.
func (dfmpe dataFactoryManagedPrivateEndpointAttributes) TargetResourceId() terra.StringValue {
	return terra.ReferenceAsString(dfmpe.ref.Append("target_resource_id"))
}

func (dfmpe dataFactoryManagedPrivateEndpointAttributes) Timeouts() datafactorymanagedprivateendpoint.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datafactorymanagedprivateendpoint.TimeoutsAttributes](dfmpe.ref.Append("timeouts"))
}

type dataFactoryManagedPrivateEndpointState struct {
	DataFactoryId    string                                           `json:"data_factory_id"`
	Fqdns            []string                                         `json:"fqdns"`
	Id               string                                           `json:"id"`
	Name             string                                           `json:"name"`
	SubresourceName  string                                           `json:"subresource_name"`
	TargetResourceId string                                           `json:"target_resource_id"`
	Timeouts         *datafactorymanagedprivateendpoint.TimeoutsState `json:"timeouts"`
}
