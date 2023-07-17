// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	appstreamdirectoryconfig "github.com/golingon/terraproviders/aws/5.8.0/appstreamdirectoryconfig"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewAppstreamDirectoryConfig creates a new instance of [AppstreamDirectoryConfig].
func NewAppstreamDirectoryConfig(name string, args AppstreamDirectoryConfigArgs) *AppstreamDirectoryConfig {
	return &AppstreamDirectoryConfig{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AppstreamDirectoryConfig)(nil)

// AppstreamDirectoryConfig represents the Terraform resource aws_appstream_directory_config.
type AppstreamDirectoryConfig struct {
	Name      string
	Args      AppstreamDirectoryConfigArgs
	state     *appstreamDirectoryConfigState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [AppstreamDirectoryConfig].
func (adc *AppstreamDirectoryConfig) Type() string {
	return "aws_appstream_directory_config"
}

// LocalName returns the local name for [AppstreamDirectoryConfig].
func (adc *AppstreamDirectoryConfig) LocalName() string {
	return adc.Name
}

// Configuration returns the configuration (args) for [AppstreamDirectoryConfig].
func (adc *AppstreamDirectoryConfig) Configuration() interface{} {
	return adc.Args
}

// DependOn is used for other resources to depend on [AppstreamDirectoryConfig].
func (adc *AppstreamDirectoryConfig) DependOn() terra.Reference {
	return terra.ReferenceResource(adc)
}

// Dependencies returns the list of resources [AppstreamDirectoryConfig] depends_on.
func (adc *AppstreamDirectoryConfig) Dependencies() terra.Dependencies {
	return adc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [AppstreamDirectoryConfig].
func (adc *AppstreamDirectoryConfig) LifecycleManagement() *terra.Lifecycle {
	return adc.Lifecycle
}

// Attributes returns the attributes for [AppstreamDirectoryConfig].
func (adc *AppstreamDirectoryConfig) Attributes() appstreamDirectoryConfigAttributes {
	return appstreamDirectoryConfigAttributes{ref: terra.ReferenceResource(adc)}
}

// ImportState imports the given attribute values into [AppstreamDirectoryConfig]'s state.
func (adc *AppstreamDirectoryConfig) ImportState(av io.Reader) error {
	adc.state = &appstreamDirectoryConfigState{}
	if err := json.NewDecoder(av).Decode(adc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", adc.Type(), adc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [AppstreamDirectoryConfig] has state.
func (adc *AppstreamDirectoryConfig) State() (*appstreamDirectoryConfigState, bool) {
	return adc.state, adc.state != nil
}

// StateMust returns the state for [AppstreamDirectoryConfig]. Panics if the state is nil.
func (adc *AppstreamDirectoryConfig) StateMust() *appstreamDirectoryConfigState {
	if adc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", adc.Type(), adc.LocalName()))
	}
	return adc.state
}

// AppstreamDirectoryConfigArgs contains the configurations for aws_appstream_directory_config.
type AppstreamDirectoryConfigArgs struct {
	// DirectoryName: string, required
	DirectoryName terra.StringValue `hcl:"directory_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// OrganizationalUnitDistinguishedNames: set of string, required
	OrganizationalUnitDistinguishedNames terra.SetValue[terra.StringValue] `hcl:"organizational_unit_distinguished_names,attr" validate:"required"`
	// ServiceAccountCredentials: required
	ServiceAccountCredentials *appstreamdirectoryconfig.ServiceAccountCredentials `hcl:"service_account_credentials,block" validate:"required"`
}
type appstreamDirectoryConfigAttributes struct {
	ref terra.Reference
}

// CreatedTime returns a reference to field created_time of aws_appstream_directory_config.
func (adc appstreamDirectoryConfigAttributes) CreatedTime() terra.StringValue {
	return terra.ReferenceAsString(adc.ref.Append("created_time"))
}

// DirectoryName returns a reference to field directory_name of aws_appstream_directory_config.
func (adc appstreamDirectoryConfigAttributes) DirectoryName() terra.StringValue {
	return terra.ReferenceAsString(adc.ref.Append("directory_name"))
}

// Id returns a reference to field id of aws_appstream_directory_config.
func (adc appstreamDirectoryConfigAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(adc.ref.Append("id"))
}

// OrganizationalUnitDistinguishedNames returns a reference to field organizational_unit_distinguished_names of aws_appstream_directory_config.
func (adc appstreamDirectoryConfigAttributes) OrganizationalUnitDistinguishedNames() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](adc.ref.Append("organizational_unit_distinguished_names"))
}

func (adc appstreamDirectoryConfigAttributes) ServiceAccountCredentials() terra.ListValue[appstreamdirectoryconfig.ServiceAccountCredentialsAttributes] {
	return terra.ReferenceAsList[appstreamdirectoryconfig.ServiceAccountCredentialsAttributes](adc.ref.Append("service_account_credentials"))
}

type appstreamDirectoryConfigState struct {
	CreatedTime                          string                                                    `json:"created_time"`
	DirectoryName                        string                                                    `json:"directory_name"`
	Id                                   string                                                    `json:"id"`
	OrganizationalUnitDistinguishedNames []string                                                  `json:"organizational_unit_distinguished_names"`
	ServiceAccountCredentials            []appstreamdirectoryconfig.ServiceAccountCredentialsState `json:"service_account_credentials"`
}
