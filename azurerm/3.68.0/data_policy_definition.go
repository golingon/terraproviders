// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datapolicydefinition "github.com/golingon/terraproviders/azurerm/3.68.0/datapolicydefinition"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataPolicyDefinition creates a new instance of [DataPolicyDefinition].
func NewDataPolicyDefinition(name string, args DataPolicyDefinitionArgs) *DataPolicyDefinition {
	return &DataPolicyDefinition{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataPolicyDefinition)(nil)

// DataPolicyDefinition represents the Terraform data resource azurerm_policy_definition.
type DataPolicyDefinition struct {
	Name string
	Args DataPolicyDefinitionArgs
}

// DataSource returns the Terraform object type for [DataPolicyDefinition].
func (pd *DataPolicyDefinition) DataSource() string {
	return "azurerm_policy_definition"
}

// LocalName returns the local name for [DataPolicyDefinition].
func (pd *DataPolicyDefinition) LocalName() string {
	return pd.Name
}

// Configuration returns the configuration (args) for [DataPolicyDefinition].
func (pd *DataPolicyDefinition) Configuration() interface{} {
	return pd.Args
}

// Attributes returns the attributes for [DataPolicyDefinition].
func (pd *DataPolicyDefinition) Attributes() dataPolicyDefinitionAttributes {
	return dataPolicyDefinitionAttributes{ref: terra.ReferenceDataResource(pd)}
}

// DataPolicyDefinitionArgs contains the configurations for azurerm_policy_definition.
type DataPolicyDefinitionArgs struct {
	// DisplayName: string, optional
	DisplayName terra.StringValue `hcl:"display_name,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ManagementGroupName: string, optional
	ManagementGroupName terra.StringValue `hcl:"management_group_name,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// Timeouts: optional
	Timeouts *datapolicydefinition.Timeouts `hcl:"timeouts,block"`
}
type dataPolicyDefinitionAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of azurerm_policy_definition.
func (pd dataPolicyDefinitionAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(pd.ref.Append("description"))
}

// DisplayName returns a reference to field display_name of azurerm_policy_definition.
func (pd dataPolicyDefinitionAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(pd.ref.Append("display_name"))
}

// Id returns a reference to field id of azurerm_policy_definition.
func (pd dataPolicyDefinitionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(pd.ref.Append("id"))
}

// ManagementGroupName returns a reference to field management_group_name of azurerm_policy_definition.
func (pd dataPolicyDefinitionAttributes) ManagementGroupName() terra.StringValue {
	return terra.ReferenceAsString(pd.ref.Append("management_group_name"))
}

// Metadata returns a reference to field metadata of azurerm_policy_definition.
func (pd dataPolicyDefinitionAttributes) Metadata() terra.StringValue {
	return terra.ReferenceAsString(pd.ref.Append("metadata"))
}

// Mode returns a reference to field mode of azurerm_policy_definition.
func (pd dataPolicyDefinitionAttributes) Mode() terra.StringValue {
	return terra.ReferenceAsString(pd.ref.Append("mode"))
}

// Name returns a reference to field name of azurerm_policy_definition.
func (pd dataPolicyDefinitionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(pd.ref.Append("name"))
}

// Parameters returns a reference to field parameters of azurerm_policy_definition.
func (pd dataPolicyDefinitionAttributes) Parameters() terra.StringValue {
	return terra.ReferenceAsString(pd.ref.Append("parameters"))
}

// PolicyRule returns a reference to field policy_rule of azurerm_policy_definition.
func (pd dataPolicyDefinitionAttributes) PolicyRule() terra.StringValue {
	return terra.ReferenceAsString(pd.ref.Append("policy_rule"))
}

// PolicyType returns a reference to field policy_type of azurerm_policy_definition.
func (pd dataPolicyDefinitionAttributes) PolicyType() terra.StringValue {
	return terra.ReferenceAsString(pd.ref.Append("policy_type"))
}

// RoleDefinitionIds returns a reference to field role_definition_ids of azurerm_policy_definition.
func (pd dataPolicyDefinitionAttributes) RoleDefinitionIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](pd.ref.Append("role_definition_ids"))
}

// Type returns a reference to field type of azurerm_policy_definition.
func (pd dataPolicyDefinitionAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(pd.ref.Append("type"))
}

func (pd dataPolicyDefinitionAttributes) Timeouts() datapolicydefinition.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datapolicydefinition.TimeoutsAttributes](pd.ref.Append("timeouts"))
}
