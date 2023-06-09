// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	dataapimanagementapi "github.com/golingon/terraproviders/azurerm/3.63.0/dataapimanagementapi"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataApiManagementApi creates a new instance of [DataApiManagementApi].
func NewDataApiManagementApi(name string, args DataApiManagementApiArgs) *DataApiManagementApi {
	return &DataApiManagementApi{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataApiManagementApi)(nil)

// DataApiManagementApi represents the Terraform data resource azurerm_api_management_api.
type DataApiManagementApi struct {
	Name string
	Args DataApiManagementApiArgs
}

// DataSource returns the Terraform object type for [DataApiManagementApi].
func (ama *DataApiManagementApi) DataSource() string {
	return "azurerm_api_management_api"
}

// LocalName returns the local name for [DataApiManagementApi].
func (ama *DataApiManagementApi) LocalName() string {
	return ama.Name
}

// Configuration returns the configuration (args) for [DataApiManagementApi].
func (ama *DataApiManagementApi) Configuration() interface{} {
	return ama.Args
}

// Attributes returns the attributes for [DataApiManagementApi].
func (ama *DataApiManagementApi) Attributes() dataApiManagementApiAttributes {
	return dataApiManagementApiAttributes{ref: terra.ReferenceDataResource(ama)}
}

// DataApiManagementApiArgs contains the configurations for azurerm_api_management_api.
type DataApiManagementApiArgs struct {
	// ApiManagementName: string, required
	ApiManagementName terra.StringValue `hcl:"api_management_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Revision: string, required
	Revision terra.StringValue `hcl:"revision,attr" validate:"required"`
	// SubscriptionKeyParameterNames: min=0
	SubscriptionKeyParameterNames []dataapimanagementapi.SubscriptionKeyParameterNames `hcl:"subscription_key_parameter_names,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *dataapimanagementapi.Timeouts `hcl:"timeouts,block"`
}
type dataApiManagementApiAttributes struct {
	ref terra.Reference
}

// ApiManagementName returns a reference to field api_management_name of azurerm_api_management_api.
func (ama dataApiManagementApiAttributes) ApiManagementName() terra.StringValue {
	return terra.ReferenceAsString(ama.ref.Append("api_management_name"))
}

// Description returns a reference to field description of azurerm_api_management_api.
func (ama dataApiManagementApiAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(ama.ref.Append("description"))
}

// DisplayName returns a reference to field display_name of azurerm_api_management_api.
func (ama dataApiManagementApiAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(ama.ref.Append("display_name"))
}

// Id returns a reference to field id of azurerm_api_management_api.
func (ama dataApiManagementApiAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ama.ref.Append("id"))
}

// IsCurrent returns a reference to field is_current of azurerm_api_management_api.
func (ama dataApiManagementApiAttributes) IsCurrent() terra.BoolValue {
	return terra.ReferenceAsBool(ama.ref.Append("is_current"))
}

// IsOnline returns a reference to field is_online of azurerm_api_management_api.
func (ama dataApiManagementApiAttributes) IsOnline() terra.BoolValue {
	return terra.ReferenceAsBool(ama.ref.Append("is_online"))
}

// Name returns a reference to field name of azurerm_api_management_api.
func (ama dataApiManagementApiAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ama.ref.Append("name"))
}

// Path returns a reference to field path of azurerm_api_management_api.
func (ama dataApiManagementApiAttributes) Path() terra.StringValue {
	return terra.ReferenceAsString(ama.ref.Append("path"))
}

// Protocols returns a reference to field protocols of azurerm_api_management_api.
func (ama dataApiManagementApiAttributes) Protocols() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ama.ref.Append("protocols"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_api_management_api.
func (ama dataApiManagementApiAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(ama.ref.Append("resource_group_name"))
}

// Revision returns a reference to field revision of azurerm_api_management_api.
func (ama dataApiManagementApiAttributes) Revision() terra.StringValue {
	return terra.ReferenceAsString(ama.ref.Append("revision"))
}

// ServiceUrl returns a reference to field service_url of azurerm_api_management_api.
func (ama dataApiManagementApiAttributes) ServiceUrl() terra.StringValue {
	return terra.ReferenceAsString(ama.ref.Append("service_url"))
}

// SoapPassThrough returns a reference to field soap_pass_through of azurerm_api_management_api.
func (ama dataApiManagementApiAttributes) SoapPassThrough() terra.BoolValue {
	return terra.ReferenceAsBool(ama.ref.Append("soap_pass_through"))
}

// SubscriptionRequired returns a reference to field subscription_required of azurerm_api_management_api.
func (ama dataApiManagementApiAttributes) SubscriptionRequired() terra.BoolValue {
	return terra.ReferenceAsBool(ama.ref.Append("subscription_required"))
}

// Version returns a reference to field version of azurerm_api_management_api.
func (ama dataApiManagementApiAttributes) Version() terra.StringValue {
	return terra.ReferenceAsString(ama.ref.Append("version"))
}

// VersionSetId returns a reference to field version_set_id of azurerm_api_management_api.
func (ama dataApiManagementApiAttributes) VersionSetId() terra.StringValue {
	return terra.ReferenceAsString(ama.ref.Append("version_set_id"))
}

func (ama dataApiManagementApiAttributes) SubscriptionKeyParameterNames() terra.ListValue[dataapimanagementapi.SubscriptionKeyParameterNamesAttributes] {
	return terra.ReferenceAsList[dataapimanagementapi.SubscriptionKeyParameterNamesAttributes](ama.ref.Append("subscription_key_parameter_names"))
}

func (ama dataApiManagementApiAttributes) Timeouts() dataapimanagementapi.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataapimanagementapi.TimeoutsAttributes](ama.ref.Append("timeouts"))
}
