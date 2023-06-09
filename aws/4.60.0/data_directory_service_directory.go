// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	datadirectoryservicedirectory "github.com/golingon/terraproviders/aws/4.60.0/datadirectoryservicedirectory"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataDirectoryServiceDirectory creates a new instance of [DataDirectoryServiceDirectory].
func NewDataDirectoryServiceDirectory(name string, args DataDirectoryServiceDirectoryArgs) *DataDirectoryServiceDirectory {
	return &DataDirectoryServiceDirectory{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataDirectoryServiceDirectory)(nil)

// DataDirectoryServiceDirectory represents the Terraform data resource aws_directory_service_directory.
type DataDirectoryServiceDirectory struct {
	Name string
	Args DataDirectoryServiceDirectoryArgs
}

// DataSource returns the Terraform object type for [DataDirectoryServiceDirectory].
func (dsd *DataDirectoryServiceDirectory) DataSource() string {
	return "aws_directory_service_directory"
}

// LocalName returns the local name for [DataDirectoryServiceDirectory].
func (dsd *DataDirectoryServiceDirectory) LocalName() string {
	return dsd.Name
}

// Configuration returns the configuration (args) for [DataDirectoryServiceDirectory].
func (dsd *DataDirectoryServiceDirectory) Configuration() interface{} {
	return dsd.Args
}

// Attributes returns the attributes for [DataDirectoryServiceDirectory].
func (dsd *DataDirectoryServiceDirectory) Attributes() dataDirectoryServiceDirectoryAttributes {
	return dataDirectoryServiceDirectoryAttributes{ref: terra.ReferenceDataResource(dsd)}
}

// DataDirectoryServiceDirectoryArgs contains the configurations for aws_directory_service_directory.
type DataDirectoryServiceDirectoryArgs struct {
	// DirectoryId: string, required
	DirectoryId terra.StringValue `hcl:"directory_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// ConnectSettings: min=0
	ConnectSettings []datadirectoryservicedirectory.ConnectSettings `hcl:"connect_settings,block" validate:"min=0"`
	// RadiusSettings: min=0
	RadiusSettings []datadirectoryservicedirectory.RadiusSettings `hcl:"radius_settings,block" validate:"min=0"`
	// VpcSettings: min=0
	VpcSettings []datadirectoryservicedirectory.VpcSettings `hcl:"vpc_settings,block" validate:"min=0"`
}
type dataDirectoryServiceDirectoryAttributes struct {
	ref terra.Reference
}

// AccessUrl returns a reference to field access_url of aws_directory_service_directory.
func (dsd dataDirectoryServiceDirectoryAttributes) AccessUrl() terra.StringValue {
	return terra.ReferenceAsString(dsd.ref.Append("access_url"))
}

// Alias returns a reference to field alias of aws_directory_service_directory.
func (dsd dataDirectoryServiceDirectoryAttributes) Alias() terra.StringValue {
	return terra.ReferenceAsString(dsd.ref.Append("alias"))
}

// Description returns a reference to field description of aws_directory_service_directory.
func (dsd dataDirectoryServiceDirectoryAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(dsd.ref.Append("description"))
}

// DirectoryId returns a reference to field directory_id of aws_directory_service_directory.
func (dsd dataDirectoryServiceDirectoryAttributes) DirectoryId() terra.StringValue {
	return terra.ReferenceAsString(dsd.ref.Append("directory_id"))
}

// DnsIpAddresses returns a reference to field dns_ip_addresses of aws_directory_service_directory.
func (dsd dataDirectoryServiceDirectoryAttributes) DnsIpAddresses() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](dsd.ref.Append("dns_ip_addresses"))
}

// Edition returns a reference to field edition of aws_directory_service_directory.
func (dsd dataDirectoryServiceDirectoryAttributes) Edition() terra.StringValue {
	return terra.ReferenceAsString(dsd.ref.Append("edition"))
}

// EnableSso returns a reference to field enable_sso of aws_directory_service_directory.
func (dsd dataDirectoryServiceDirectoryAttributes) EnableSso() terra.BoolValue {
	return terra.ReferenceAsBool(dsd.ref.Append("enable_sso"))
}

// Id returns a reference to field id of aws_directory_service_directory.
func (dsd dataDirectoryServiceDirectoryAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dsd.ref.Append("id"))
}

// Name returns a reference to field name of aws_directory_service_directory.
func (dsd dataDirectoryServiceDirectoryAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dsd.ref.Append("name"))
}

// SecurityGroupId returns a reference to field security_group_id of aws_directory_service_directory.
func (dsd dataDirectoryServiceDirectoryAttributes) SecurityGroupId() terra.StringValue {
	return terra.ReferenceAsString(dsd.ref.Append("security_group_id"))
}

// ShortName returns a reference to field short_name of aws_directory_service_directory.
func (dsd dataDirectoryServiceDirectoryAttributes) ShortName() terra.StringValue {
	return terra.ReferenceAsString(dsd.ref.Append("short_name"))
}

// Size returns a reference to field size of aws_directory_service_directory.
func (dsd dataDirectoryServiceDirectoryAttributes) Size() terra.StringValue {
	return terra.ReferenceAsString(dsd.ref.Append("size"))
}

// Tags returns a reference to field tags of aws_directory_service_directory.
func (dsd dataDirectoryServiceDirectoryAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dsd.ref.Append("tags"))
}

// Type returns a reference to field type of aws_directory_service_directory.
func (dsd dataDirectoryServiceDirectoryAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(dsd.ref.Append("type"))
}

func (dsd dataDirectoryServiceDirectoryAttributes) ConnectSettings() terra.ListValue[datadirectoryservicedirectory.ConnectSettingsAttributes] {
	return terra.ReferenceAsList[datadirectoryservicedirectory.ConnectSettingsAttributes](dsd.ref.Append("connect_settings"))
}

func (dsd dataDirectoryServiceDirectoryAttributes) RadiusSettings() terra.ListValue[datadirectoryservicedirectory.RadiusSettingsAttributes] {
	return terra.ReferenceAsList[datadirectoryservicedirectory.RadiusSettingsAttributes](dsd.ref.Append("radius_settings"))
}

func (dsd dataDirectoryServiceDirectoryAttributes) VpcSettings() terra.ListValue[datadirectoryservicedirectory.VpcSettingsAttributes] {
	return terra.ReferenceAsList[datadirectoryservicedirectory.VpcSettingsAttributes](dsd.ref.Append("vpc_settings"))
}
