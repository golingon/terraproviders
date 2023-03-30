// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	directoryservicedirectory "github.com/golingon/terraproviders/aws/4.60.0/directoryservicedirectory"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewDirectoryServiceDirectory(name string, args DirectoryServiceDirectoryArgs) *DirectoryServiceDirectory {
	return &DirectoryServiceDirectory{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DirectoryServiceDirectory)(nil)

type DirectoryServiceDirectory struct {
	Name  string
	Args  DirectoryServiceDirectoryArgs
	state *directoryServiceDirectoryState
}

func (dsd *DirectoryServiceDirectory) Type() string {
	return "aws_directory_service_directory"
}

func (dsd *DirectoryServiceDirectory) LocalName() string {
	return dsd.Name
}

func (dsd *DirectoryServiceDirectory) Configuration() interface{} {
	return dsd.Args
}

func (dsd *DirectoryServiceDirectory) Attributes() directoryServiceDirectoryAttributes {
	return directoryServiceDirectoryAttributes{ref: terra.ReferenceResource(dsd)}
}

func (dsd *DirectoryServiceDirectory) ImportState(av io.Reader) error {
	dsd.state = &directoryServiceDirectoryState{}
	if err := json.NewDecoder(av).Decode(dsd.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dsd.Type(), dsd.LocalName(), err)
	}
	return nil
}

func (dsd *DirectoryServiceDirectory) State() (*directoryServiceDirectoryState, bool) {
	return dsd.state, dsd.state != nil
}

func (dsd *DirectoryServiceDirectory) StateMust() *directoryServiceDirectoryState {
	if dsd.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dsd.Type(), dsd.LocalName()))
	}
	return dsd.state
}

func (dsd *DirectoryServiceDirectory) DependOn() terra.Reference {
	return terra.ReferenceResource(dsd)
}

type DirectoryServiceDirectoryArgs struct {
	// Alias: string, optional
	Alias terra.StringValue `hcl:"alias,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// DesiredNumberOfDomainControllers: number, optional
	DesiredNumberOfDomainControllers terra.NumberValue `hcl:"desired_number_of_domain_controllers,attr"`
	// Edition: string, optional
	Edition terra.StringValue `hcl:"edition,attr"`
	// EnableSso: bool, optional
	EnableSso terra.BoolValue `hcl:"enable_sso,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Password: string, required
	Password terra.StringValue `hcl:"password,attr" validate:"required"`
	// ShortName: string, optional
	ShortName terra.StringValue `hcl:"short_name,attr"`
	// Size: string, optional
	Size terra.StringValue `hcl:"size,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// Type: string, optional
	Type terra.StringValue `hcl:"type,attr"`
	// ConnectSettings: optional
	ConnectSettings *directoryservicedirectory.ConnectSettings `hcl:"connect_settings,block"`
	// Timeouts: optional
	Timeouts *directoryservicedirectory.Timeouts `hcl:"timeouts,block"`
	// VpcSettings: optional
	VpcSettings *directoryservicedirectory.VpcSettings `hcl:"vpc_settings,block"`
	// DependsOn contains resources that DirectoryServiceDirectory depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type directoryServiceDirectoryAttributes struct {
	ref terra.Reference
}

func (dsd directoryServiceDirectoryAttributes) AccessUrl() terra.StringValue {
	return terra.ReferenceString(dsd.ref.Append("access_url"))
}

func (dsd directoryServiceDirectoryAttributes) Alias() terra.StringValue {
	return terra.ReferenceString(dsd.ref.Append("alias"))
}

func (dsd directoryServiceDirectoryAttributes) Description() terra.StringValue {
	return terra.ReferenceString(dsd.ref.Append("description"))
}

func (dsd directoryServiceDirectoryAttributes) DesiredNumberOfDomainControllers() terra.NumberValue {
	return terra.ReferenceNumber(dsd.ref.Append("desired_number_of_domain_controllers"))
}

func (dsd directoryServiceDirectoryAttributes) DnsIpAddresses() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](dsd.ref.Append("dns_ip_addresses"))
}

func (dsd directoryServiceDirectoryAttributes) Edition() terra.StringValue {
	return terra.ReferenceString(dsd.ref.Append("edition"))
}

func (dsd directoryServiceDirectoryAttributes) EnableSso() terra.BoolValue {
	return terra.ReferenceBool(dsd.ref.Append("enable_sso"))
}

func (dsd directoryServiceDirectoryAttributes) Id() terra.StringValue {
	return terra.ReferenceString(dsd.ref.Append("id"))
}

func (dsd directoryServiceDirectoryAttributes) Name() terra.StringValue {
	return terra.ReferenceString(dsd.ref.Append("name"))
}

func (dsd directoryServiceDirectoryAttributes) Password() terra.StringValue {
	return terra.ReferenceString(dsd.ref.Append("password"))
}

func (dsd directoryServiceDirectoryAttributes) SecurityGroupId() terra.StringValue {
	return terra.ReferenceString(dsd.ref.Append("security_group_id"))
}

func (dsd directoryServiceDirectoryAttributes) ShortName() terra.StringValue {
	return terra.ReferenceString(dsd.ref.Append("short_name"))
}

func (dsd directoryServiceDirectoryAttributes) Size() terra.StringValue {
	return terra.ReferenceString(dsd.ref.Append("size"))
}

func (dsd directoryServiceDirectoryAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](dsd.ref.Append("tags"))
}

func (dsd directoryServiceDirectoryAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](dsd.ref.Append("tags_all"))
}

func (dsd directoryServiceDirectoryAttributes) Type() terra.StringValue {
	return terra.ReferenceString(dsd.ref.Append("type"))
}

func (dsd directoryServiceDirectoryAttributes) ConnectSettings() terra.ListValue[directoryservicedirectory.ConnectSettingsAttributes] {
	return terra.ReferenceList[directoryservicedirectory.ConnectSettingsAttributes](dsd.ref.Append("connect_settings"))
}

func (dsd directoryServiceDirectoryAttributes) Timeouts() directoryservicedirectory.TimeoutsAttributes {
	return terra.ReferenceSingle[directoryservicedirectory.TimeoutsAttributes](dsd.ref.Append("timeouts"))
}

func (dsd directoryServiceDirectoryAttributes) VpcSettings() terra.ListValue[directoryservicedirectory.VpcSettingsAttributes] {
	return terra.ReferenceList[directoryservicedirectory.VpcSettingsAttributes](dsd.ref.Append("vpc_settings"))
}

type directoryServiceDirectoryState struct {
	AccessUrl                        string                                           `json:"access_url"`
	Alias                            string                                           `json:"alias"`
	Description                      string                                           `json:"description"`
	DesiredNumberOfDomainControllers float64                                          `json:"desired_number_of_domain_controllers"`
	DnsIpAddresses                   []string                                         `json:"dns_ip_addresses"`
	Edition                          string                                           `json:"edition"`
	EnableSso                        bool                                             `json:"enable_sso"`
	Id                               string                                           `json:"id"`
	Name                             string                                           `json:"name"`
	Password                         string                                           `json:"password"`
	SecurityGroupId                  string                                           `json:"security_group_id"`
	ShortName                        string                                           `json:"short_name"`
	Size                             string                                           `json:"size"`
	Tags                             map[string]string                                `json:"tags"`
	TagsAll                          map[string]string                                `json:"tags_all"`
	Type                             string                                           `json:"type"`
	ConnectSettings                  []directoryservicedirectory.ConnectSettingsState `json:"connect_settings"`
	Timeouts                         *directoryservicedirectory.TimeoutsState         `json:"timeouts"`
	VpcSettings                      []directoryservicedirectory.VpcSettingsState     `json:"vpc_settings"`
}
