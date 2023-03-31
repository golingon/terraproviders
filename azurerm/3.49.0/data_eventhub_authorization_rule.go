// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	dataeventhubauthorizationrule "github.com/golingon/terraproviders/azurerm/3.49.0/dataeventhubauthorizationrule"
	"github.com/volvo-cars/lingon/pkg/terra"
)

func NewDataEventhubAuthorizationRule(name string, args DataEventhubAuthorizationRuleArgs) *DataEventhubAuthorizationRule {
	return &DataEventhubAuthorizationRule{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataEventhubAuthorizationRule)(nil)

type DataEventhubAuthorizationRule struct {
	Name string
	Args DataEventhubAuthorizationRuleArgs
}

func (ear *DataEventhubAuthorizationRule) DataSource() string {
	return "azurerm_eventhub_authorization_rule"
}

func (ear *DataEventhubAuthorizationRule) LocalName() string {
	return ear.Name
}

func (ear *DataEventhubAuthorizationRule) Configuration() interface{} {
	return ear.Args
}

func (ear *DataEventhubAuthorizationRule) Attributes() dataEventhubAuthorizationRuleAttributes {
	return dataEventhubAuthorizationRuleAttributes{ref: terra.ReferenceDataResource(ear)}
}

type DataEventhubAuthorizationRuleArgs struct {
	// EventhubName: string, required
	EventhubName terra.StringValue `hcl:"eventhub_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Listen: bool, optional
	Listen terra.BoolValue `hcl:"listen,attr"`
	// Manage: bool, optional
	Manage terra.BoolValue `hcl:"manage,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// NamespaceName: string, required
	NamespaceName terra.StringValue `hcl:"namespace_name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Send: bool, optional
	Send terra.BoolValue `hcl:"send,attr"`
	// Timeouts: optional
	Timeouts *dataeventhubauthorizationrule.Timeouts `hcl:"timeouts,block"`
}
type dataEventhubAuthorizationRuleAttributes struct {
	ref terra.Reference
}

func (ear dataEventhubAuthorizationRuleAttributes) EventhubName() terra.StringValue {
	return terra.ReferenceString(ear.ref.Append("eventhub_name"))
}

func (ear dataEventhubAuthorizationRuleAttributes) Id() terra.StringValue {
	return terra.ReferenceString(ear.ref.Append("id"))
}

func (ear dataEventhubAuthorizationRuleAttributes) Listen() terra.BoolValue {
	return terra.ReferenceBool(ear.ref.Append("listen"))
}

func (ear dataEventhubAuthorizationRuleAttributes) Manage() terra.BoolValue {
	return terra.ReferenceBool(ear.ref.Append("manage"))
}

func (ear dataEventhubAuthorizationRuleAttributes) Name() terra.StringValue {
	return terra.ReferenceString(ear.ref.Append("name"))
}

func (ear dataEventhubAuthorizationRuleAttributes) NamespaceName() terra.StringValue {
	return terra.ReferenceString(ear.ref.Append("namespace_name"))
}

func (ear dataEventhubAuthorizationRuleAttributes) PrimaryConnectionString() terra.StringValue {
	return terra.ReferenceString(ear.ref.Append("primary_connection_string"))
}

func (ear dataEventhubAuthorizationRuleAttributes) PrimaryConnectionStringAlias() terra.StringValue {
	return terra.ReferenceString(ear.ref.Append("primary_connection_string_alias"))
}

func (ear dataEventhubAuthorizationRuleAttributes) PrimaryKey() terra.StringValue {
	return terra.ReferenceString(ear.ref.Append("primary_key"))
}

func (ear dataEventhubAuthorizationRuleAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceString(ear.ref.Append("resource_group_name"))
}

func (ear dataEventhubAuthorizationRuleAttributes) SecondaryConnectionString() terra.StringValue {
	return terra.ReferenceString(ear.ref.Append("secondary_connection_string"))
}

func (ear dataEventhubAuthorizationRuleAttributes) SecondaryConnectionStringAlias() terra.StringValue {
	return terra.ReferenceString(ear.ref.Append("secondary_connection_string_alias"))
}

func (ear dataEventhubAuthorizationRuleAttributes) SecondaryKey() terra.StringValue {
	return terra.ReferenceString(ear.ref.Append("secondary_key"))
}

func (ear dataEventhubAuthorizationRuleAttributes) Send() terra.BoolValue {
	return terra.ReferenceBool(ear.ref.Append("send"))
}

func (ear dataEventhubAuthorizationRuleAttributes) Timeouts() dataeventhubauthorizationrule.TimeoutsAttributes {
	return terra.ReferenceSingle[dataeventhubauthorizationrule.TimeoutsAttributes](ear.ref.Append("timeouts"))
}
