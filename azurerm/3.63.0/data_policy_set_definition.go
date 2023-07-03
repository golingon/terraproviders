// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datapolicysetdefinition "github.com/golingon/terraproviders/azurerm/3.63.0/datapolicysetdefinition"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataPolicySetDefinition creates a new instance of [DataPolicySetDefinition].
func NewDataPolicySetDefinition(name string, args DataPolicySetDefinitionArgs) *DataPolicySetDefinition {
	return &DataPolicySetDefinition{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataPolicySetDefinition)(nil)

// DataPolicySetDefinition represents the Terraform data resource azurerm_policy_set_definition.
type DataPolicySetDefinition struct {
	Name string
	Args DataPolicySetDefinitionArgs
}

// DataSource returns the Terraform object type for [DataPolicySetDefinition].
func (psd *DataPolicySetDefinition) DataSource() string {
	return "azurerm_policy_set_definition"
}

// LocalName returns the local name for [DataPolicySetDefinition].
func (psd *DataPolicySetDefinition) LocalName() string {
	return psd.Name
}

// Configuration returns the configuration (args) for [DataPolicySetDefinition].
func (psd *DataPolicySetDefinition) Configuration() interface{} {
	return psd.Args
}

// Attributes returns the attributes for [DataPolicySetDefinition].
func (psd *DataPolicySetDefinition) Attributes() dataPolicySetDefinitionAttributes {
	return dataPolicySetDefinitionAttributes{ref: terra.ReferenceDataResource(psd)}
}

// DataPolicySetDefinitionArgs contains the configurations for azurerm_policy_set_definition.
type DataPolicySetDefinitionArgs struct {
	// DisplayName: string, optional
	DisplayName terra.StringValue `hcl:"display_name,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ManagementGroupName: string, optional
	ManagementGroupName terra.StringValue `hcl:"management_group_name,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// PolicyDefinitionGroup: min=0
	PolicyDefinitionGroup []datapolicysetdefinition.PolicyDefinitionGroup `hcl:"policy_definition_group,block" validate:"min=0"`
	// PolicyDefinitionReference: min=0
	PolicyDefinitionReference []datapolicysetdefinition.PolicyDefinitionReference `hcl:"policy_definition_reference,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datapolicysetdefinition.Timeouts `hcl:"timeouts,block"`
}
type dataPolicySetDefinitionAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of azurerm_policy_set_definition.
func (psd dataPolicySetDefinitionAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(psd.ref.Append("description"))
}

// DisplayName returns a reference to field display_name of azurerm_policy_set_definition.
func (psd dataPolicySetDefinitionAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(psd.ref.Append("display_name"))
}

// Id returns a reference to field id of azurerm_policy_set_definition.
func (psd dataPolicySetDefinitionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(psd.ref.Append("id"))
}

// ManagementGroupName returns a reference to field management_group_name of azurerm_policy_set_definition.
func (psd dataPolicySetDefinitionAttributes) ManagementGroupName() terra.StringValue {
	return terra.ReferenceAsString(psd.ref.Append("management_group_name"))
}

// Metadata returns a reference to field metadata of azurerm_policy_set_definition.
func (psd dataPolicySetDefinitionAttributes) Metadata() terra.StringValue {
	return terra.ReferenceAsString(psd.ref.Append("metadata"))
}

// Name returns a reference to field name of azurerm_policy_set_definition.
func (psd dataPolicySetDefinitionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(psd.ref.Append("name"))
}

// Parameters returns a reference to field parameters of azurerm_policy_set_definition.
func (psd dataPolicySetDefinitionAttributes) Parameters() terra.StringValue {
	return terra.ReferenceAsString(psd.ref.Append("parameters"))
}

// PolicyDefinitions returns a reference to field policy_definitions of azurerm_policy_set_definition.
func (psd dataPolicySetDefinitionAttributes) PolicyDefinitions() terra.StringValue {
	return terra.ReferenceAsString(psd.ref.Append("policy_definitions"))
}

// PolicyType returns a reference to field policy_type of azurerm_policy_set_definition.
func (psd dataPolicySetDefinitionAttributes) PolicyType() terra.StringValue {
	return terra.ReferenceAsString(psd.ref.Append("policy_type"))
}

func (psd dataPolicySetDefinitionAttributes) PolicyDefinitionGroup() terra.ListValue[datapolicysetdefinition.PolicyDefinitionGroupAttributes] {
	return terra.ReferenceAsList[datapolicysetdefinition.PolicyDefinitionGroupAttributes](psd.ref.Append("policy_definition_group"))
}

func (psd dataPolicySetDefinitionAttributes) PolicyDefinitionReference() terra.ListValue[datapolicysetdefinition.PolicyDefinitionReferenceAttributes] {
	return terra.ReferenceAsList[datapolicysetdefinition.PolicyDefinitionReferenceAttributes](psd.ref.Append("policy_definition_reference"))
}

func (psd dataPolicySetDefinitionAttributes) Timeouts() datapolicysetdefinition.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datapolicysetdefinition.TimeoutsAttributes](psd.ref.Append("timeouts"))
}
