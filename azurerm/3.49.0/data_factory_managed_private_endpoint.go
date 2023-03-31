// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	datafactorymanagedprivateendpoint "github.com/golingon/terraproviders/azurerm/3.49.0/datafactorymanagedprivateendpoint"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewDataFactoryManagedPrivateEndpoint(name string, args DataFactoryManagedPrivateEndpointArgs) *DataFactoryManagedPrivateEndpoint {
	return &DataFactoryManagedPrivateEndpoint{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DataFactoryManagedPrivateEndpoint)(nil)

type DataFactoryManagedPrivateEndpoint struct {
	Name  string
	Args  DataFactoryManagedPrivateEndpointArgs
	state *dataFactoryManagedPrivateEndpointState
}

func (dfmpe *DataFactoryManagedPrivateEndpoint) Type() string {
	return "azurerm_data_factory_managed_private_endpoint"
}

func (dfmpe *DataFactoryManagedPrivateEndpoint) LocalName() string {
	return dfmpe.Name
}

func (dfmpe *DataFactoryManagedPrivateEndpoint) Configuration() interface{} {
	return dfmpe.Args
}

func (dfmpe *DataFactoryManagedPrivateEndpoint) Attributes() dataFactoryManagedPrivateEndpointAttributes {
	return dataFactoryManagedPrivateEndpointAttributes{ref: terra.ReferenceResource(dfmpe)}
}

func (dfmpe *DataFactoryManagedPrivateEndpoint) ImportState(av io.Reader) error {
	dfmpe.state = &dataFactoryManagedPrivateEndpointState{}
	if err := json.NewDecoder(av).Decode(dfmpe.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dfmpe.Type(), dfmpe.LocalName(), err)
	}
	return nil
}

func (dfmpe *DataFactoryManagedPrivateEndpoint) State() (*dataFactoryManagedPrivateEndpointState, bool) {
	return dfmpe.state, dfmpe.state != nil
}

func (dfmpe *DataFactoryManagedPrivateEndpoint) StateMust() *dataFactoryManagedPrivateEndpointState {
	if dfmpe.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dfmpe.Type(), dfmpe.LocalName()))
	}
	return dfmpe.state
}

func (dfmpe *DataFactoryManagedPrivateEndpoint) DependOn() terra.Reference {
	return terra.ReferenceResource(dfmpe)
}

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
	// DependsOn contains resources that DataFactoryManagedPrivateEndpoint depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type dataFactoryManagedPrivateEndpointAttributes struct {
	ref terra.Reference
}

func (dfmpe dataFactoryManagedPrivateEndpointAttributes) DataFactoryId() terra.StringValue {
	return terra.ReferenceString(dfmpe.ref.Append("data_factory_id"))
}

func (dfmpe dataFactoryManagedPrivateEndpointAttributes) Fqdns() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](dfmpe.ref.Append("fqdns"))
}

func (dfmpe dataFactoryManagedPrivateEndpointAttributes) Id() terra.StringValue {
	return terra.ReferenceString(dfmpe.ref.Append("id"))
}

func (dfmpe dataFactoryManagedPrivateEndpointAttributes) Name() terra.StringValue {
	return terra.ReferenceString(dfmpe.ref.Append("name"))
}

func (dfmpe dataFactoryManagedPrivateEndpointAttributes) SubresourceName() terra.StringValue {
	return terra.ReferenceString(dfmpe.ref.Append("subresource_name"))
}

func (dfmpe dataFactoryManagedPrivateEndpointAttributes) TargetResourceId() terra.StringValue {
	return terra.ReferenceString(dfmpe.ref.Append("target_resource_id"))
}

func (dfmpe dataFactoryManagedPrivateEndpointAttributes) Timeouts() datafactorymanagedprivateendpoint.TimeoutsAttributes {
	return terra.ReferenceSingle[datafactorymanagedprivateendpoint.TimeoutsAttributes](dfmpe.ref.Append("timeouts"))
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
