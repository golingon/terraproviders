// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataorganizationsorganizationalunits "github.com/golingon/terraproviders/aws/5.11.0/dataorganizationsorganizationalunits"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataOrganizationsOrganizationalUnits creates a new instance of [DataOrganizationsOrganizationalUnits].
func NewDataOrganizationsOrganizationalUnits(name string, args DataOrganizationsOrganizationalUnitsArgs) *DataOrganizationsOrganizationalUnits {
	return &DataOrganizationsOrganizationalUnits{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataOrganizationsOrganizationalUnits)(nil)

// DataOrganizationsOrganizationalUnits represents the Terraform data resource aws_organizations_organizational_units.
type DataOrganizationsOrganizationalUnits struct {
	Name string
	Args DataOrganizationsOrganizationalUnitsArgs
}

// DataSource returns the Terraform object type for [DataOrganizationsOrganizationalUnits].
func (oou *DataOrganizationsOrganizationalUnits) DataSource() string {
	return "aws_organizations_organizational_units"
}

// LocalName returns the local name for [DataOrganizationsOrganizationalUnits].
func (oou *DataOrganizationsOrganizationalUnits) LocalName() string {
	return oou.Name
}

// Configuration returns the configuration (args) for [DataOrganizationsOrganizationalUnits].
func (oou *DataOrganizationsOrganizationalUnits) Configuration() interface{} {
	return oou.Args
}

// Attributes returns the attributes for [DataOrganizationsOrganizationalUnits].
func (oou *DataOrganizationsOrganizationalUnits) Attributes() dataOrganizationsOrganizationalUnitsAttributes {
	return dataOrganizationsOrganizationalUnitsAttributes{ref: terra.ReferenceDataResource(oou)}
}

// DataOrganizationsOrganizationalUnitsArgs contains the configurations for aws_organizations_organizational_units.
type DataOrganizationsOrganizationalUnitsArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ParentId: string, required
	ParentId terra.StringValue `hcl:"parent_id,attr" validate:"required"`
	// Children: min=0
	Children []dataorganizationsorganizationalunits.Children `hcl:"children,block" validate:"min=0"`
}
type dataOrganizationsOrganizationalUnitsAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_organizations_organizational_units.
func (oou dataOrganizationsOrganizationalUnitsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(oou.ref.Append("id"))
}

// ParentId returns a reference to field parent_id of aws_organizations_organizational_units.
func (oou dataOrganizationsOrganizationalUnitsAttributes) ParentId() terra.StringValue {
	return terra.ReferenceAsString(oou.ref.Append("parent_id"))
}

func (oou dataOrganizationsOrganizationalUnitsAttributes) Children() terra.ListValue[dataorganizationsorganizationalunits.ChildrenAttributes] {
	return terra.ReferenceAsList[dataorganizationsorganizationalunits.ChildrenAttributes](oou.ref.Append("children"))
}
