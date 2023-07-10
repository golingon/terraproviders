// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datapolicydefinitionbuiltin "github.com/golingon/terraproviders/azurerm/3.64.0/datapolicydefinitionbuiltin"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataPolicyDefinitionBuiltIn creates a new instance of [DataPolicyDefinitionBuiltIn].
func NewDataPolicyDefinitionBuiltIn(name string, args DataPolicyDefinitionBuiltInArgs) *DataPolicyDefinitionBuiltIn {
	return &DataPolicyDefinitionBuiltIn{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataPolicyDefinitionBuiltIn)(nil)

// DataPolicyDefinitionBuiltIn represents the Terraform data resource azurerm_policy_definition_built_in.
type DataPolicyDefinitionBuiltIn struct {
	Name string
	Args DataPolicyDefinitionBuiltInArgs
}

// DataSource returns the Terraform object type for [DataPolicyDefinitionBuiltIn].
func (pdbi *DataPolicyDefinitionBuiltIn) DataSource() string {
	return "azurerm_policy_definition_built_in"
}

// LocalName returns the local name for [DataPolicyDefinitionBuiltIn].
func (pdbi *DataPolicyDefinitionBuiltIn) LocalName() string {
	return pdbi.Name
}

// Configuration returns the configuration (args) for [DataPolicyDefinitionBuiltIn].
func (pdbi *DataPolicyDefinitionBuiltIn) Configuration() interface{} {
	return pdbi.Args
}

// Attributes returns the attributes for [DataPolicyDefinitionBuiltIn].
func (pdbi *DataPolicyDefinitionBuiltIn) Attributes() dataPolicyDefinitionBuiltInAttributes {
	return dataPolicyDefinitionBuiltInAttributes{ref: terra.ReferenceDataResource(pdbi)}
}

// DataPolicyDefinitionBuiltInArgs contains the configurations for azurerm_policy_definition_built_in.
type DataPolicyDefinitionBuiltInArgs struct {
	// DisplayName: string, optional
	DisplayName terra.StringValue `hcl:"display_name,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ManagementGroupName: string, optional
	ManagementGroupName terra.StringValue `hcl:"management_group_name,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// Timeouts: optional
	Timeouts *datapolicydefinitionbuiltin.Timeouts `hcl:"timeouts,block"`
}
type dataPolicyDefinitionBuiltInAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of azurerm_policy_definition_built_in.
func (pdbi dataPolicyDefinitionBuiltInAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(pdbi.ref.Append("description"))
}

// DisplayName returns a reference to field display_name of azurerm_policy_definition_built_in.
func (pdbi dataPolicyDefinitionBuiltInAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(pdbi.ref.Append("display_name"))
}

// Id returns a reference to field id of azurerm_policy_definition_built_in.
func (pdbi dataPolicyDefinitionBuiltInAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(pdbi.ref.Append("id"))
}

// ManagementGroupName returns a reference to field management_group_name of azurerm_policy_definition_built_in.
func (pdbi dataPolicyDefinitionBuiltInAttributes) ManagementGroupName() terra.StringValue {
	return terra.ReferenceAsString(pdbi.ref.Append("management_group_name"))
}

// Metadata returns a reference to field metadata of azurerm_policy_definition_built_in.
func (pdbi dataPolicyDefinitionBuiltInAttributes) Metadata() terra.StringValue {
	return terra.ReferenceAsString(pdbi.ref.Append("metadata"))
}

// Mode returns a reference to field mode of azurerm_policy_definition_built_in.
func (pdbi dataPolicyDefinitionBuiltInAttributes) Mode() terra.StringValue {
	return terra.ReferenceAsString(pdbi.ref.Append("mode"))
}

// Name returns a reference to field name of azurerm_policy_definition_built_in.
func (pdbi dataPolicyDefinitionBuiltInAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(pdbi.ref.Append("name"))
}

// Parameters returns a reference to field parameters of azurerm_policy_definition_built_in.
func (pdbi dataPolicyDefinitionBuiltInAttributes) Parameters() terra.StringValue {
	return terra.ReferenceAsString(pdbi.ref.Append("parameters"))
}

// PolicyRule returns a reference to field policy_rule of azurerm_policy_definition_built_in.
func (pdbi dataPolicyDefinitionBuiltInAttributes) PolicyRule() terra.StringValue {
	return terra.ReferenceAsString(pdbi.ref.Append("policy_rule"))
}

// PolicyType returns a reference to field policy_type of azurerm_policy_definition_built_in.
func (pdbi dataPolicyDefinitionBuiltInAttributes) PolicyType() terra.StringValue {
	return terra.ReferenceAsString(pdbi.ref.Append("policy_type"))
}

// RoleDefinitionIds returns a reference to field role_definition_ids of azurerm_policy_definition_built_in.
func (pdbi dataPolicyDefinitionBuiltInAttributes) RoleDefinitionIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](pdbi.ref.Append("role_definition_ids"))
}

// Type returns a reference to field type of azurerm_policy_definition_built_in.
func (pdbi dataPolicyDefinitionBuiltInAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(pdbi.ref.Append("type"))
}

func (pdbi dataPolicyDefinitionBuiltInAttributes) Timeouts() datapolicydefinitionbuiltin.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datapolicydefinitionbuiltin.TimeoutsAttributes](pdbi.ref.Append("timeouts"))
}
