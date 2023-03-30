// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	datadirectoryservicedirectory "github.com/golingon/terraproviders/aws/4.60.0/datadirectoryservicedirectory"
	"github.com/volvo-cars/lingon/pkg/terra"
)

func NewDataDirectoryServiceDirectory(name string, args DataDirectoryServiceDirectoryArgs) *DataDirectoryServiceDirectory {
	return &DataDirectoryServiceDirectory{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataDirectoryServiceDirectory)(nil)

type DataDirectoryServiceDirectory struct {
	Name string
	Args DataDirectoryServiceDirectoryArgs
}

func (dsd *DataDirectoryServiceDirectory) DataSource() string {
	return "aws_directory_service_directory"
}

func (dsd *DataDirectoryServiceDirectory) LocalName() string {
	return dsd.Name
}

func (dsd *DataDirectoryServiceDirectory) Configuration() interface{} {
	return dsd.Args
}

func (dsd *DataDirectoryServiceDirectory) Attributes() dataDirectoryServiceDirectoryAttributes {
	return dataDirectoryServiceDirectoryAttributes{ref: terra.ReferenceDataResource(dsd)}
}

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

func (dsd dataDirectoryServiceDirectoryAttributes) AccessUrl() terra.StringValue {
	return terra.ReferenceString(dsd.ref.Append("access_url"))
}

func (dsd dataDirectoryServiceDirectoryAttributes) Alias() terra.StringValue {
	return terra.ReferenceString(dsd.ref.Append("alias"))
}

func (dsd dataDirectoryServiceDirectoryAttributes) Description() terra.StringValue {
	return terra.ReferenceString(dsd.ref.Append("description"))
}

func (dsd dataDirectoryServiceDirectoryAttributes) DirectoryId() terra.StringValue {
	return terra.ReferenceString(dsd.ref.Append("directory_id"))
}

func (dsd dataDirectoryServiceDirectoryAttributes) DnsIpAddresses() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](dsd.ref.Append("dns_ip_addresses"))
}

func (dsd dataDirectoryServiceDirectoryAttributes) Edition() terra.StringValue {
	return terra.ReferenceString(dsd.ref.Append("edition"))
}

func (dsd dataDirectoryServiceDirectoryAttributes) EnableSso() terra.BoolValue {
	return terra.ReferenceBool(dsd.ref.Append("enable_sso"))
}

func (dsd dataDirectoryServiceDirectoryAttributes) Id() terra.StringValue {
	return terra.ReferenceString(dsd.ref.Append("id"))
}

func (dsd dataDirectoryServiceDirectoryAttributes) Name() terra.StringValue {
	return terra.ReferenceString(dsd.ref.Append("name"))
}

func (dsd dataDirectoryServiceDirectoryAttributes) SecurityGroupId() terra.StringValue {
	return terra.ReferenceString(dsd.ref.Append("security_group_id"))
}

func (dsd dataDirectoryServiceDirectoryAttributes) ShortName() terra.StringValue {
	return terra.ReferenceString(dsd.ref.Append("short_name"))
}

func (dsd dataDirectoryServiceDirectoryAttributes) Size() terra.StringValue {
	return terra.ReferenceString(dsd.ref.Append("size"))
}

func (dsd dataDirectoryServiceDirectoryAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](dsd.ref.Append("tags"))
}

func (dsd dataDirectoryServiceDirectoryAttributes) Type() terra.StringValue {
	return terra.ReferenceString(dsd.ref.Append("type"))
}

func (dsd dataDirectoryServiceDirectoryAttributes) ConnectSettings() terra.ListValue[datadirectoryservicedirectory.ConnectSettingsAttributes] {
	return terra.ReferenceList[datadirectoryservicedirectory.ConnectSettingsAttributes](dsd.ref.Append("connect_settings"))
}

func (dsd dataDirectoryServiceDirectoryAttributes) RadiusSettings() terra.ListValue[datadirectoryservicedirectory.RadiusSettingsAttributes] {
	return terra.ReferenceList[datadirectoryservicedirectory.RadiusSettingsAttributes](dsd.ref.Append("radius_settings"))
}

func (dsd dataDirectoryServiceDirectoryAttributes) VpcSettings() terra.ListValue[datadirectoryservicedirectory.VpcSettingsAttributes] {
	return terra.ReferenceList[datadirectoryservicedirectory.VpcSettingsAttributes](dsd.ref.Append("vpc_settings"))
}
