// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	dataiamtestablepermissions "github.com/golingon/terraproviders/google/4.74.0/dataiamtestablepermissions"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataIamTestablePermissions creates a new instance of [DataIamTestablePermissions].
func NewDataIamTestablePermissions(name string, args DataIamTestablePermissionsArgs) *DataIamTestablePermissions {
	return &DataIamTestablePermissions{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataIamTestablePermissions)(nil)

// DataIamTestablePermissions represents the Terraform data resource google_iam_testable_permissions.
type DataIamTestablePermissions struct {
	Name string
	Args DataIamTestablePermissionsArgs
}

// DataSource returns the Terraform object type for [DataIamTestablePermissions].
func (itp *DataIamTestablePermissions) DataSource() string {
	return "google_iam_testable_permissions"
}

// LocalName returns the local name for [DataIamTestablePermissions].
func (itp *DataIamTestablePermissions) LocalName() string {
	return itp.Name
}

// Configuration returns the configuration (args) for [DataIamTestablePermissions].
func (itp *DataIamTestablePermissions) Configuration() interface{} {
	return itp.Args
}

// Attributes returns the attributes for [DataIamTestablePermissions].
func (itp *DataIamTestablePermissions) Attributes() dataIamTestablePermissionsAttributes {
	return dataIamTestablePermissionsAttributes{ref: terra.ReferenceDataResource(itp)}
}

// DataIamTestablePermissionsArgs contains the configurations for google_iam_testable_permissions.
type DataIamTestablePermissionsArgs struct {
	// CustomSupportLevel: string, optional
	CustomSupportLevel terra.StringValue `hcl:"custom_support_level,attr"`
	// FullResourceName: string, required
	FullResourceName terra.StringValue `hcl:"full_resource_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Stages: list of string, optional
	Stages terra.ListValue[terra.StringValue] `hcl:"stages,attr"`
	// Permissions: min=0
	Permissions []dataiamtestablepermissions.Permissions `hcl:"permissions,block" validate:"min=0"`
}
type dataIamTestablePermissionsAttributes struct {
	ref terra.Reference
}

// CustomSupportLevel returns a reference to field custom_support_level of google_iam_testable_permissions.
func (itp dataIamTestablePermissionsAttributes) CustomSupportLevel() terra.StringValue {
	return terra.ReferenceAsString(itp.ref.Append("custom_support_level"))
}

// FullResourceName returns a reference to field full_resource_name of google_iam_testable_permissions.
func (itp dataIamTestablePermissionsAttributes) FullResourceName() terra.StringValue {
	return terra.ReferenceAsString(itp.ref.Append("full_resource_name"))
}

// Id returns a reference to field id of google_iam_testable_permissions.
func (itp dataIamTestablePermissionsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(itp.ref.Append("id"))
}

// Stages returns a reference to field stages of google_iam_testable_permissions.
func (itp dataIamTestablePermissionsAttributes) Stages() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](itp.ref.Append("stages"))
}

func (itp dataIamTestablePermissionsAttributes) Permissions() terra.ListValue[dataiamtestablepermissions.PermissionsAttributes] {
	return terra.ReferenceAsList[dataiamtestablepermissions.PermissionsAttributes](itp.ref.Append("permissions"))
}
