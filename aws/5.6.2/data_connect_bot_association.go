// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataconnectbotassociation "github.com/golingon/terraproviders/aws/5.6.2/dataconnectbotassociation"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataConnectBotAssociation creates a new instance of [DataConnectBotAssociation].
func NewDataConnectBotAssociation(name string, args DataConnectBotAssociationArgs) *DataConnectBotAssociation {
	return &DataConnectBotAssociation{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataConnectBotAssociation)(nil)

// DataConnectBotAssociation represents the Terraform data resource aws_connect_bot_association.
type DataConnectBotAssociation struct {
	Name string
	Args DataConnectBotAssociationArgs
}

// DataSource returns the Terraform object type for [DataConnectBotAssociation].
func (cba *DataConnectBotAssociation) DataSource() string {
	return "aws_connect_bot_association"
}

// LocalName returns the local name for [DataConnectBotAssociation].
func (cba *DataConnectBotAssociation) LocalName() string {
	return cba.Name
}

// Configuration returns the configuration (args) for [DataConnectBotAssociation].
func (cba *DataConnectBotAssociation) Configuration() interface{} {
	return cba.Args
}

// Attributes returns the attributes for [DataConnectBotAssociation].
func (cba *DataConnectBotAssociation) Attributes() dataConnectBotAssociationAttributes {
	return dataConnectBotAssociationAttributes{ref: terra.ReferenceDataResource(cba)}
}

// DataConnectBotAssociationArgs contains the configurations for aws_connect_bot_association.
type DataConnectBotAssociationArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InstanceId: string, required
	InstanceId terra.StringValue `hcl:"instance_id,attr" validate:"required"`
	// LexBot: required
	LexBot *dataconnectbotassociation.LexBot `hcl:"lex_bot,block" validate:"required"`
}
type dataConnectBotAssociationAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_connect_bot_association.
func (cba dataConnectBotAssociationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cba.ref.Append("id"))
}

// InstanceId returns a reference to field instance_id of aws_connect_bot_association.
func (cba dataConnectBotAssociationAttributes) InstanceId() terra.StringValue {
	return terra.ReferenceAsString(cba.ref.Append("instance_id"))
}

func (cba dataConnectBotAssociationAttributes) LexBot() terra.ListValue[dataconnectbotassociation.LexBotAttributes] {
	return terra.ReferenceAsList[dataconnectbotassociation.LexBotAttributes](cba.ref.Append("lex_bot"))
}
