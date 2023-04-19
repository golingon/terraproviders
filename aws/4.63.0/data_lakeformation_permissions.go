// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	datalakeformationpermissions "github.com/golingon/terraproviders/aws/4.63.0/datalakeformationpermissions"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataLakeformationPermissions creates a new instance of [DataLakeformationPermissions].
func NewDataLakeformationPermissions(name string, args DataLakeformationPermissionsArgs) *DataLakeformationPermissions {
	return &DataLakeformationPermissions{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataLakeformationPermissions)(nil)

// DataLakeformationPermissions represents the Terraform data resource aws_lakeformation_permissions.
type DataLakeformationPermissions struct {
	Name string
	Args DataLakeformationPermissionsArgs
}

// DataSource returns the Terraform object type for [DataLakeformationPermissions].
func (lp *DataLakeformationPermissions) DataSource() string {
	return "aws_lakeformation_permissions"
}

// LocalName returns the local name for [DataLakeformationPermissions].
func (lp *DataLakeformationPermissions) LocalName() string {
	return lp.Name
}

// Configuration returns the configuration (args) for [DataLakeformationPermissions].
func (lp *DataLakeformationPermissions) Configuration() interface{} {
	return lp.Args
}

// Attributes returns the attributes for [DataLakeformationPermissions].
func (lp *DataLakeformationPermissions) Attributes() dataLakeformationPermissionsAttributes {
	return dataLakeformationPermissionsAttributes{ref: terra.ReferenceDataResource(lp)}
}

// DataLakeformationPermissionsArgs contains the configurations for aws_lakeformation_permissions.
type DataLakeformationPermissionsArgs struct {
	// CatalogId: string, optional
	CatalogId terra.StringValue `hcl:"catalog_id,attr"`
	// CatalogResource: bool, optional
	CatalogResource terra.BoolValue `hcl:"catalog_resource,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Principal: string, required
	Principal terra.StringValue `hcl:"principal,attr" validate:"required"`
	// DataLocation: optional
	DataLocation *datalakeformationpermissions.DataLocation `hcl:"data_location,block"`
	// Database: optional
	Database *datalakeformationpermissions.Database `hcl:"database,block"`
	// LfTag: optional
	LfTag *datalakeformationpermissions.LfTag `hcl:"lf_tag,block"`
	// LfTagPolicy: optional
	LfTagPolicy *datalakeformationpermissions.LfTagPolicy `hcl:"lf_tag_policy,block"`
	// Table: optional
	Table *datalakeformationpermissions.Table `hcl:"table,block"`
	// TableWithColumns: optional
	TableWithColumns *datalakeformationpermissions.TableWithColumns `hcl:"table_with_columns,block"`
}
type dataLakeformationPermissionsAttributes struct {
	ref terra.Reference
}

// CatalogId returns a reference to field catalog_id of aws_lakeformation_permissions.
func (lp dataLakeformationPermissionsAttributes) CatalogId() terra.StringValue {
	return terra.ReferenceAsString(lp.ref.Append("catalog_id"))
}

// CatalogResource returns a reference to field catalog_resource of aws_lakeformation_permissions.
func (lp dataLakeformationPermissionsAttributes) CatalogResource() terra.BoolValue {
	return terra.ReferenceAsBool(lp.ref.Append("catalog_resource"))
}

// Id returns a reference to field id of aws_lakeformation_permissions.
func (lp dataLakeformationPermissionsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(lp.ref.Append("id"))
}

// Permissions returns a reference to field permissions of aws_lakeformation_permissions.
func (lp dataLakeformationPermissionsAttributes) Permissions() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](lp.ref.Append("permissions"))
}

// PermissionsWithGrantOption returns a reference to field permissions_with_grant_option of aws_lakeformation_permissions.
func (lp dataLakeformationPermissionsAttributes) PermissionsWithGrantOption() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](lp.ref.Append("permissions_with_grant_option"))
}

// Principal returns a reference to field principal of aws_lakeformation_permissions.
func (lp dataLakeformationPermissionsAttributes) Principal() terra.StringValue {
	return terra.ReferenceAsString(lp.ref.Append("principal"))
}

func (lp dataLakeformationPermissionsAttributes) DataLocation() terra.ListValue[datalakeformationpermissions.DataLocationAttributes] {
	return terra.ReferenceAsList[datalakeformationpermissions.DataLocationAttributes](lp.ref.Append("data_location"))
}

func (lp dataLakeformationPermissionsAttributes) Database() terra.ListValue[datalakeformationpermissions.DatabaseAttributes] {
	return terra.ReferenceAsList[datalakeformationpermissions.DatabaseAttributes](lp.ref.Append("database"))
}

func (lp dataLakeformationPermissionsAttributes) LfTag() terra.ListValue[datalakeformationpermissions.LfTagAttributes] {
	return terra.ReferenceAsList[datalakeformationpermissions.LfTagAttributes](lp.ref.Append("lf_tag"))
}

func (lp dataLakeformationPermissionsAttributes) LfTagPolicy() terra.ListValue[datalakeformationpermissions.LfTagPolicyAttributes] {
	return terra.ReferenceAsList[datalakeformationpermissions.LfTagPolicyAttributes](lp.ref.Append("lf_tag_policy"))
}

func (lp dataLakeformationPermissionsAttributes) Table() terra.ListValue[datalakeformationpermissions.TableAttributes] {
	return terra.ReferenceAsList[datalakeformationpermissions.TableAttributes](lp.ref.Append("table"))
}

func (lp dataLakeformationPermissionsAttributes) TableWithColumns() terra.ListValue[datalakeformationpermissions.TableWithColumnsAttributes] {
	return terra.ReferenceAsList[datalakeformationpermissions.TableWithColumnsAttributes](lp.ref.Append("table_with_columns"))
}
