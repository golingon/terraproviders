// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	apimanagementglobalschema "github.com/golingon/terraproviders/azurerm/3.68.0/apimanagementglobalschema"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewApiManagementGlobalSchema creates a new instance of [ApiManagementGlobalSchema].
func NewApiManagementGlobalSchema(name string, args ApiManagementGlobalSchemaArgs) *ApiManagementGlobalSchema {
	return &ApiManagementGlobalSchema{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ApiManagementGlobalSchema)(nil)

// ApiManagementGlobalSchema represents the Terraform resource azurerm_api_management_global_schema.
type ApiManagementGlobalSchema struct {
	Name      string
	Args      ApiManagementGlobalSchemaArgs
	state     *apiManagementGlobalSchemaState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ApiManagementGlobalSchema].
func (amgs *ApiManagementGlobalSchema) Type() string {
	return "azurerm_api_management_global_schema"
}

// LocalName returns the local name for [ApiManagementGlobalSchema].
func (amgs *ApiManagementGlobalSchema) LocalName() string {
	return amgs.Name
}

// Configuration returns the configuration (args) for [ApiManagementGlobalSchema].
func (amgs *ApiManagementGlobalSchema) Configuration() interface{} {
	return amgs.Args
}

// DependOn is used for other resources to depend on [ApiManagementGlobalSchema].
func (amgs *ApiManagementGlobalSchema) DependOn() terra.Reference {
	return terra.ReferenceResource(amgs)
}

// Dependencies returns the list of resources [ApiManagementGlobalSchema] depends_on.
func (amgs *ApiManagementGlobalSchema) Dependencies() terra.Dependencies {
	return amgs.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ApiManagementGlobalSchema].
func (amgs *ApiManagementGlobalSchema) LifecycleManagement() *terra.Lifecycle {
	return amgs.Lifecycle
}

// Attributes returns the attributes for [ApiManagementGlobalSchema].
func (amgs *ApiManagementGlobalSchema) Attributes() apiManagementGlobalSchemaAttributes {
	return apiManagementGlobalSchemaAttributes{ref: terra.ReferenceResource(amgs)}
}

// ImportState imports the given attribute values into [ApiManagementGlobalSchema]'s state.
func (amgs *ApiManagementGlobalSchema) ImportState(av io.Reader) error {
	amgs.state = &apiManagementGlobalSchemaState{}
	if err := json.NewDecoder(av).Decode(amgs.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", amgs.Type(), amgs.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ApiManagementGlobalSchema] has state.
func (amgs *ApiManagementGlobalSchema) State() (*apiManagementGlobalSchemaState, bool) {
	return amgs.state, amgs.state != nil
}

// StateMust returns the state for [ApiManagementGlobalSchema]. Panics if the state is nil.
func (amgs *ApiManagementGlobalSchema) StateMust() *apiManagementGlobalSchemaState {
	if amgs.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", amgs.Type(), amgs.LocalName()))
	}
	return amgs.state
}

// ApiManagementGlobalSchemaArgs contains the configurations for azurerm_api_management_global_schema.
type ApiManagementGlobalSchemaArgs struct {
	// ApiManagementName: string, required
	ApiManagementName terra.StringValue `hcl:"api_management_name,attr" validate:"required"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// SchemaId: string, required
	SchemaId terra.StringValue `hcl:"schema_id,attr" validate:"required"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
	// Value: string, required
	Value terra.StringValue `hcl:"value,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *apimanagementglobalschema.Timeouts `hcl:"timeouts,block"`
}
type apiManagementGlobalSchemaAttributes struct {
	ref terra.Reference
}

// ApiManagementName returns a reference to field api_management_name of azurerm_api_management_global_schema.
func (amgs apiManagementGlobalSchemaAttributes) ApiManagementName() terra.StringValue {
	return terra.ReferenceAsString(amgs.ref.Append("api_management_name"))
}

// Description returns a reference to field description of azurerm_api_management_global_schema.
func (amgs apiManagementGlobalSchemaAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(amgs.ref.Append("description"))
}

// Id returns a reference to field id of azurerm_api_management_global_schema.
func (amgs apiManagementGlobalSchemaAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(amgs.ref.Append("id"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_api_management_global_schema.
func (amgs apiManagementGlobalSchemaAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(amgs.ref.Append("resource_group_name"))
}

// SchemaId returns a reference to field schema_id of azurerm_api_management_global_schema.
func (amgs apiManagementGlobalSchemaAttributes) SchemaId() terra.StringValue {
	return terra.ReferenceAsString(amgs.ref.Append("schema_id"))
}

// Type returns a reference to field type of azurerm_api_management_global_schema.
func (amgs apiManagementGlobalSchemaAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(amgs.ref.Append("type"))
}

// Value returns a reference to field value of azurerm_api_management_global_schema.
func (amgs apiManagementGlobalSchemaAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(amgs.ref.Append("value"))
}

func (amgs apiManagementGlobalSchemaAttributes) Timeouts() apimanagementglobalschema.TimeoutsAttributes {
	return terra.ReferenceAsSingle[apimanagementglobalschema.TimeoutsAttributes](amgs.ref.Append("timeouts"))
}

type apiManagementGlobalSchemaState struct {
	ApiManagementName string                                   `json:"api_management_name"`
	Description       string                                   `json:"description"`
	Id                string                                   `json:"id"`
	ResourceGroupName string                                   `json:"resource_group_name"`
	SchemaId          string                                   `json:"schema_id"`
	Type              string                                   `json:"type"`
	Value             string                                   `json:"value"`
	Timeouts          *apimanagementglobalschema.TimeoutsState `json:"timeouts"`
}
