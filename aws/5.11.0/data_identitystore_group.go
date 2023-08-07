// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataidentitystoregroup "github.com/golingon/terraproviders/aws/5.11.0/dataidentitystoregroup"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataIdentitystoreGroup creates a new instance of [DataIdentitystoreGroup].
func NewDataIdentitystoreGroup(name string, args DataIdentitystoreGroupArgs) *DataIdentitystoreGroup {
	return &DataIdentitystoreGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataIdentitystoreGroup)(nil)

// DataIdentitystoreGroup represents the Terraform data resource aws_identitystore_group.
type DataIdentitystoreGroup struct {
	Name string
	Args DataIdentitystoreGroupArgs
}

// DataSource returns the Terraform object type for [DataIdentitystoreGroup].
func (ig *DataIdentitystoreGroup) DataSource() string {
	return "aws_identitystore_group"
}

// LocalName returns the local name for [DataIdentitystoreGroup].
func (ig *DataIdentitystoreGroup) LocalName() string {
	return ig.Name
}

// Configuration returns the configuration (args) for [DataIdentitystoreGroup].
func (ig *DataIdentitystoreGroup) Configuration() interface{} {
	return ig.Args
}

// Attributes returns the attributes for [DataIdentitystoreGroup].
func (ig *DataIdentitystoreGroup) Attributes() dataIdentitystoreGroupAttributes {
	return dataIdentitystoreGroupAttributes{ref: terra.ReferenceDataResource(ig)}
}

// DataIdentitystoreGroupArgs contains the configurations for aws_identitystore_group.
type DataIdentitystoreGroupArgs struct {
	// GroupId: string, optional
	GroupId terra.StringValue `hcl:"group_id,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IdentityStoreId: string, required
	IdentityStoreId terra.StringValue `hcl:"identity_store_id,attr" validate:"required"`
	// ExternalIds: min=0
	ExternalIds []dataidentitystoregroup.ExternalIds `hcl:"external_ids,block" validate:"min=0"`
	// AlternateIdentifier: optional
	AlternateIdentifier *dataidentitystoregroup.AlternateIdentifier `hcl:"alternate_identifier,block"`
}
type dataIdentitystoreGroupAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of aws_identitystore_group.
func (ig dataIdentitystoreGroupAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(ig.ref.Append("description"))
}

// DisplayName returns a reference to field display_name of aws_identitystore_group.
func (ig dataIdentitystoreGroupAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(ig.ref.Append("display_name"))
}

// GroupId returns a reference to field group_id of aws_identitystore_group.
func (ig dataIdentitystoreGroupAttributes) GroupId() terra.StringValue {
	return terra.ReferenceAsString(ig.ref.Append("group_id"))
}

// Id returns a reference to field id of aws_identitystore_group.
func (ig dataIdentitystoreGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ig.ref.Append("id"))
}

// IdentityStoreId returns a reference to field identity_store_id of aws_identitystore_group.
func (ig dataIdentitystoreGroupAttributes) IdentityStoreId() terra.StringValue {
	return terra.ReferenceAsString(ig.ref.Append("identity_store_id"))
}

func (ig dataIdentitystoreGroupAttributes) ExternalIds() terra.ListValue[dataidentitystoregroup.ExternalIdsAttributes] {
	return terra.ReferenceAsList[dataidentitystoregroup.ExternalIdsAttributes](ig.ref.Append("external_ids"))
}

func (ig dataIdentitystoreGroupAttributes) AlternateIdentifier() terra.ListValue[dataidentitystoregroup.AlternateIdentifierAttributes] {
	return terra.ReferenceAsList[dataidentitystoregroup.AlternateIdentifierAttributes](ig.ref.Append("alternate_identifier"))
}
