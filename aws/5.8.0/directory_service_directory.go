// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	directoryservicedirectory "github.com/golingon/terraproviders/aws/5.8.0/directoryservicedirectory"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDirectoryServiceDirectory creates a new instance of [DirectoryServiceDirectory].
func NewDirectoryServiceDirectory(name string, args DirectoryServiceDirectoryArgs) *DirectoryServiceDirectory {
	return &DirectoryServiceDirectory{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DirectoryServiceDirectory)(nil)

// DirectoryServiceDirectory represents the Terraform resource aws_directory_service_directory.
type DirectoryServiceDirectory struct {
	Name      string
	Args      DirectoryServiceDirectoryArgs
	state     *directoryServiceDirectoryState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DirectoryServiceDirectory].
func (dsd *DirectoryServiceDirectory) Type() string {
	return "aws_directory_service_directory"
}

// LocalName returns the local name for [DirectoryServiceDirectory].
func (dsd *DirectoryServiceDirectory) LocalName() string {
	return dsd.Name
}

// Configuration returns the configuration (args) for [DirectoryServiceDirectory].
func (dsd *DirectoryServiceDirectory) Configuration() interface{} {
	return dsd.Args
}

// DependOn is used for other resources to depend on [DirectoryServiceDirectory].
func (dsd *DirectoryServiceDirectory) DependOn() terra.Reference {
	return terra.ReferenceResource(dsd)
}

// Dependencies returns the list of resources [DirectoryServiceDirectory] depends_on.
func (dsd *DirectoryServiceDirectory) Dependencies() terra.Dependencies {
	return dsd.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DirectoryServiceDirectory].
func (dsd *DirectoryServiceDirectory) LifecycleManagement() *terra.Lifecycle {
	return dsd.Lifecycle
}

// Attributes returns the attributes for [DirectoryServiceDirectory].
func (dsd *DirectoryServiceDirectory) Attributes() directoryServiceDirectoryAttributes {
	return directoryServiceDirectoryAttributes{ref: terra.ReferenceResource(dsd)}
}

// ImportState imports the given attribute values into [DirectoryServiceDirectory]'s state.
func (dsd *DirectoryServiceDirectory) ImportState(av io.Reader) error {
	dsd.state = &directoryServiceDirectoryState{}
	if err := json.NewDecoder(av).Decode(dsd.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dsd.Type(), dsd.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DirectoryServiceDirectory] has state.
func (dsd *DirectoryServiceDirectory) State() (*directoryServiceDirectoryState, bool) {
	return dsd.state, dsd.state != nil
}

// StateMust returns the state for [DirectoryServiceDirectory]. Panics if the state is nil.
func (dsd *DirectoryServiceDirectory) StateMust() *directoryServiceDirectoryState {
	if dsd.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dsd.Type(), dsd.LocalName()))
	}
	return dsd.state
}

// DirectoryServiceDirectoryArgs contains the configurations for aws_directory_service_directory.
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
}
type directoryServiceDirectoryAttributes struct {
	ref terra.Reference
}

// AccessUrl returns a reference to field access_url of aws_directory_service_directory.
func (dsd directoryServiceDirectoryAttributes) AccessUrl() terra.StringValue {
	return terra.ReferenceAsString(dsd.ref.Append("access_url"))
}

// Alias returns a reference to field alias of aws_directory_service_directory.
func (dsd directoryServiceDirectoryAttributes) Alias() terra.StringValue {
	return terra.ReferenceAsString(dsd.ref.Append("alias"))
}

// Description returns a reference to field description of aws_directory_service_directory.
func (dsd directoryServiceDirectoryAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(dsd.ref.Append("description"))
}

// DesiredNumberOfDomainControllers returns a reference to field desired_number_of_domain_controllers of aws_directory_service_directory.
func (dsd directoryServiceDirectoryAttributes) DesiredNumberOfDomainControllers() terra.NumberValue {
	return terra.ReferenceAsNumber(dsd.ref.Append("desired_number_of_domain_controllers"))
}

// DnsIpAddresses returns a reference to field dns_ip_addresses of aws_directory_service_directory.
func (dsd directoryServiceDirectoryAttributes) DnsIpAddresses() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](dsd.ref.Append("dns_ip_addresses"))
}

// Edition returns a reference to field edition of aws_directory_service_directory.
func (dsd directoryServiceDirectoryAttributes) Edition() terra.StringValue {
	return terra.ReferenceAsString(dsd.ref.Append("edition"))
}

// EnableSso returns a reference to field enable_sso of aws_directory_service_directory.
func (dsd directoryServiceDirectoryAttributes) EnableSso() terra.BoolValue {
	return terra.ReferenceAsBool(dsd.ref.Append("enable_sso"))
}

// Id returns a reference to field id of aws_directory_service_directory.
func (dsd directoryServiceDirectoryAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dsd.ref.Append("id"))
}

// Name returns a reference to field name of aws_directory_service_directory.
func (dsd directoryServiceDirectoryAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dsd.ref.Append("name"))
}

// Password returns a reference to field password of aws_directory_service_directory.
func (dsd directoryServiceDirectoryAttributes) Password() terra.StringValue {
	return terra.ReferenceAsString(dsd.ref.Append("password"))
}

// SecurityGroupId returns a reference to field security_group_id of aws_directory_service_directory.
func (dsd directoryServiceDirectoryAttributes) SecurityGroupId() terra.StringValue {
	return terra.ReferenceAsString(dsd.ref.Append("security_group_id"))
}

// ShortName returns a reference to field short_name of aws_directory_service_directory.
func (dsd directoryServiceDirectoryAttributes) ShortName() terra.StringValue {
	return terra.ReferenceAsString(dsd.ref.Append("short_name"))
}

// Size returns a reference to field size of aws_directory_service_directory.
func (dsd directoryServiceDirectoryAttributes) Size() terra.StringValue {
	return terra.ReferenceAsString(dsd.ref.Append("size"))
}

// Tags returns a reference to field tags of aws_directory_service_directory.
func (dsd directoryServiceDirectoryAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dsd.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_directory_service_directory.
func (dsd directoryServiceDirectoryAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dsd.ref.Append("tags_all"))
}

// Type returns a reference to field type of aws_directory_service_directory.
func (dsd directoryServiceDirectoryAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(dsd.ref.Append("type"))
}

func (dsd directoryServiceDirectoryAttributes) ConnectSettings() terra.ListValue[directoryservicedirectory.ConnectSettingsAttributes] {
	return terra.ReferenceAsList[directoryservicedirectory.ConnectSettingsAttributes](dsd.ref.Append("connect_settings"))
}

func (dsd directoryServiceDirectoryAttributes) Timeouts() directoryservicedirectory.TimeoutsAttributes {
	return terra.ReferenceAsSingle[directoryservicedirectory.TimeoutsAttributes](dsd.ref.Append("timeouts"))
}

func (dsd directoryServiceDirectoryAttributes) VpcSettings() terra.ListValue[directoryservicedirectory.VpcSettingsAttributes] {
	return terra.ReferenceAsList[directoryservicedirectory.VpcSettingsAttributes](dsd.ref.Append("vpc_settings"))
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
