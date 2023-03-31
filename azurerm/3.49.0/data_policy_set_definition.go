// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datapolicysetdefinition "github.com/golingon/terraproviders/azurerm/3.49.0/datapolicysetdefinition"
	"github.com/volvo-cars/lingon/pkg/terra"
)

func NewDataPolicySetDefinition(name string, args DataPolicySetDefinitionArgs) *DataPolicySetDefinition {
	return &DataPolicySetDefinition{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataPolicySetDefinition)(nil)

type DataPolicySetDefinition struct {
	Name string
	Args DataPolicySetDefinitionArgs
}

func (psd *DataPolicySetDefinition) DataSource() string {
	return "azurerm_policy_set_definition"
}

func (psd *DataPolicySetDefinition) LocalName() string {
	return psd.Name
}

func (psd *DataPolicySetDefinition) Configuration() interface{} {
	return psd.Args
}

func (psd *DataPolicySetDefinition) Attributes() dataPolicySetDefinitionAttributes {
	return dataPolicySetDefinitionAttributes{ref: terra.ReferenceDataResource(psd)}
}

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

func (psd dataPolicySetDefinitionAttributes) Description() terra.StringValue {
	return terra.ReferenceString(psd.ref.Append("description"))
}

func (psd dataPolicySetDefinitionAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceString(psd.ref.Append("display_name"))
}

func (psd dataPolicySetDefinitionAttributes) Id() terra.StringValue {
	return terra.ReferenceString(psd.ref.Append("id"))
}

func (psd dataPolicySetDefinitionAttributes) ManagementGroupName() terra.StringValue {
	return terra.ReferenceString(psd.ref.Append("management_group_name"))
}

func (psd dataPolicySetDefinitionAttributes) Metadata() terra.StringValue {
	return terra.ReferenceString(psd.ref.Append("metadata"))
}

func (psd dataPolicySetDefinitionAttributes) Name() terra.StringValue {
	return terra.ReferenceString(psd.ref.Append("name"))
}

func (psd dataPolicySetDefinitionAttributes) Parameters() terra.StringValue {
	return terra.ReferenceString(psd.ref.Append("parameters"))
}

func (psd dataPolicySetDefinitionAttributes) PolicyDefinitions() terra.StringValue {
	return terra.ReferenceString(psd.ref.Append("policy_definitions"))
}

func (psd dataPolicySetDefinitionAttributes) PolicyType() terra.StringValue {
	return terra.ReferenceString(psd.ref.Append("policy_type"))
}

func (psd dataPolicySetDefinitionAttributes) PolicyDefinitionGroup() terra.ListValue[datapolicysetdefinition.PolicyDefinitionGroupAttributes] {
	return terra.ReferenceList[datapolicysetdefinition.PolicyDefinitionGroupAttributes](psd.ref.Append("policy_definition_group"))
}

func (psd dataPolicySetDefinitionAttributes) PolicyDefinitionReference() terra.ListValue[datapolicysetdefinition.PolicyDefinitionReferenceAttributes] {
	return terra.ReferenceList[datapolicysetdefinition.PolicyDefinitionReferenceAttributes](psd.ref.Append("policy_definition_reference"))
}

func (psd dataPolicySetDefinitionAttributes) Timeouts() datapolicysetdefinition.TimeoutsAttributes {
	return terra.ReferenceSingle[datapolicysetdefinition.TimeoutsAttributes](psd.ref.Append("timeouts"))
}
