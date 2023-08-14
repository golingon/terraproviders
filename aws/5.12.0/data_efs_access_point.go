// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataefsaccesspoint "github.com/golingon/terraproviders/aws/5.12.0/dataefsaccesspoint"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataEfsAccessPoint creates a new instance of [DataEfsAccessPoint].
func NewDataEfsAccessPoint(name string, args DataEfsAccessPointArgs) *DataEfsAccessPoint {
	return &DataEfsAccessPoint{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataEfsAccessPoint)(nil)

// DataEfsAccessPoint represents the Terraform data resource aws_efs_access_point.
type DataEfsAccessPoint struct {
	Name string
	Args DataEfsAccessPointArgs
}

// DataSource returns the Terraform object type for [DataEfsAccessPoint].
func (eap *DataEfsAccessPoint) DataSource() string {
	return "aws_efs_access_point"
}

// LocalName returns the local name for [DataEfsAccessPoint].
func (eap *DataEfsAccessPoint) LocalName() string {
	return eap.Name
}

// Configuration returns the configuration (args) for [DataEfsAccessPoint].
func (eap *DataEfsAccessPoint) Configuration() interface{} {
	return eap.Args
}

// Attributes returns the attributes for [DataEfsAccessPoint].
func (eap *DataEfsAccessPoint) Attributes() dataEfsAccessPointAttributes {
	return dataEfsAccessPointAttributes{ref: terra.ReferenceDataResource(eap)}
}

// DataEfsAccessPointArgs contains the configurations for aws_efs_access_point.
type DataEfsAccessPointArgs struct {
	// AccessPointId: string, required
	AccessPointId terra.StringValue `hcl:"access_point_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// PosixUser: min=0
	PosixUser []dataefsaccesspoint.PosixUser `hcl:"posix_user,block" validate:"min=0"`
	// RootDirectory: min=0
	RootDirectory []dataefsaccesspoint.RootDirectory `hcl:"root_directory,block" validate:"min=0"`
}
type dataEfsAccessPointAttributes struct {
	ref terra.Reference
}

// AccessPointId returns a reference to field access_point_id of aws_efs_access_point.
func (eap dataEfsAccessPointAttributes) AccessPointId() terra.StringValue {
	return terra.ReferenceAsString(eap.ref.Append("access_point_id"))
}

// Arn returns a reference to field arn of aws_efs_access_point.
func (eap dataEfsAccessPointAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(eap.ref.Append("arn"))
}

// FileSystemArn returns a reference to field file_system_arn of aws_efs_access_point.
func (eap dataEfsAccessPointAttributes) FileSystemArn() terra.StringValue {
	return terra.ReferenceAsString(eap.ref.Append("file_system_arn"))
}

// FileSystemId returns a reference to field file_system_id of aws_efs_access_point.
func (eap dataEfsAccessPointAttributes) FileSystemId() terra.StringValue {
	return terra.ReferenceAsString(eap.ref.Append("file_system_id"))
}

// Id returns a reference to field id of aws_efs_access_point.
func (eap dataEfsAccessPointAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(eap.ref.Append("id"))
}

// OwnerId returns a reference to field owner_id of aws_efs_access_point.
func (eap dataEfsAccessPointAttributes) OwnerId() terra.StringValue {
	return terra.ReferenceAsString(eap.ref.Append("owner_id"))
}

// Tags returns a reference to field tags of aws_efs_access_point.
func (eap dataEfsAccessPointAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](eap.ref.Append("tags"))
}

func (eap dataEfsAccessPointAttributes) PosixUser() terra.ListValue[dataefsaccesspoint.PosixUserAttributes] {
	return terra.ReferenceAsList[dataefsaccesspoint.PosixUserAttributes](eap.ref.Append("posix_user"))
}

func (eap dataEfsAccessPointAttributes) RootDirectory() terra.ListValue[dataefsaccesspoint.RootDirectoryAttributes] {
	return terra.ReferenceAsList[dataefsaccesspoint.RootDirectoryAttributes](eap.ref.Append("root_directory"))
}
