// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewAppconfigHostedConfigurationVersion creates a new instance of [AppconfigHostedConfigurationVersion].
func NewAppconfigHostedConfigurationVersion(name string, args AppconfigHostedConfigurationVersionArgs) *AppconfigHostedConfigurationVersion {
	return &AppconfigHostedConfigurationVersion{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AppconfigHostedConfigurationVersion)(nil)

// AppconfigHostedConfigurationVersion represents the Terraform resource aws_appconfig_hosted_configuration_version.
type AppconfigHostedConfigurationVersion struct {
	Name      string
	Args      AppconfigHostedConfigurationVersionArgs
	state     *appconfigHostedConfigurationVersionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [AppconfigHostedConfigurationVersion].
func (ahcv *AppconfigHostedConfigurationVersion) Type() string {
	return "aws_appconfig_hosted_configuration_version"
}

// LocalName returns the local name for [AppconfigHostedConfigurationVersion].
func (ahcv *AppconfigHostedConfigurationVersion) LocalName() string {
	return ahcv.Name
}

// Configuration returns the configuration (args) for [AppconfigHostedConfigurationVersion].
func (ahcv *AppconfigHostedConfigurationVersion) Configuration() interface{} {
	return ahcv.Args
}

// DependOn is used for other resources to depend on [AppconfigHostedConfigurationVersion].
func (ahcv *AppconfigHostedConfigurationVersion) DependOn() terra.Reference {
	return terra.ReferenceResource(ahcv)
}

// Dependencies returns the list of resources [AppconfigHostedConfigurationVersion] depends_on.
func (ahcv *AppconfigHostedConfigurationVersion) Dependencies() terra.Dependencies {
	return ahcv.DependsOn
}

// LifecycleManagement returns the lifecycle block for [AppconfigHostedConfigurationVersion].
func (ahcv *AppconfigHostedConfigurationVersion) LifecycleManagement() *terra.Lifecycle {
	return ahcv.Lifecycle
}

// Attributes returns the attributes for [AppconfigHostedConfigurationVersion].
func (ahcv *AppconfigHostedConfigurationVersion) Attributes() appconfigHostedConfigurationVersionAttributes {
	return appconfigHostedConfigurationVersionAttributes{ref: terra.ReferenceResource(ahcv)}
}

// ImportState imports the given attribute values into [AppconfigHostedConfigurationVersion]'s state.
func (ahcv *AppconfigHostedConfigurationVersion) ImportState(av io.Reader) error {
	ahcv.state = &appconfigHostedConfigurationVersionState{}
	if err := json.NewDecoder(av).Decode(ahcv.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ahcv.Type(), ahcv.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [AppconfigHostedConfigurationVersion] has state.
func (ahcv *AppconfigHostedConfigurationVersion) State() (*appconfigHostedConfigurationVersionState, bool) {
	return ahcv.state, ahcv.state != nil
}

// StateMust returns the state for [AppconfigHostedConfigurationVersion]. Panics if the state is nil.
func (ahcv *AppconfigHostedConfigurationVersion) StateMust() *appconfigHostedConfigurationVersionState {
	if ahcv.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ahcv.Type(), ahcv.LocalName()))
	}
	return ahcv.state
}

// AppconfigHostedConfigurationVersionArgs contains the configurations for aws_appconfig_hosted_configuration_version.
type AppconfigHostedConfigurationVersionArgs struct {
	// ApplicationId: string, required
	ApplicationId terra.StringValue `hcl:"application_id,attr" validate:"required"`
	// ConfigurationProfileId: string, required
	ConfigurationProfileId terra.StringValue `hcl:"configuration_profile_id,attr" validate:"required"`
	// Content: string, required
	Content terra.StringValue `hcl:"content,attr" validate:"required"`
	// ContentType: string, required
	ContentType terra.StringValue `hcl:"content_type,attr" validate:"required"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
}
type appconfigHostedConfigurationVersionAttributes struct {
	ref terra.Reference
}

// ApplicationId returns a reference to field application_id of aws_appconfig_hosted_configuration_version.
func (ahcv appconfigHostedConfigurationVersionAttributes) ApplicationId() terra.StringValue {
	return terra.ReferenceAsString(ahcv.ref.Append("application_id"))
}

// Arn returns a reference to field arn of aws_appconfig_hosted_configuration_version.
func (ahcv appconfigHostedConfigurationVersionAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ahcv.ref.Append("arn"))
}

// ConfigurationProfileId returns a reference to field configuration_profile_id of aws_appconfig_hosted_configuration_version.
func (ahcv appconfigHostedConfigurationVersionAttributes) ConfigurationProfileId() terra.StringValue {
	return terra.ReferenceAsString(ahcv.ref.Append("configuration_profile_id"))
}

// Content returns a reference to field content of aws_appconfig_hosted_configuration_version.
func (ahcv appconfigHostedConfigurationVersionAttributes) Content() terra.StringValue {
	return terra.ReferenceAsString(ahcv.ref.Append("content"))
}

// ContentType returns a reference to field content_type of aws_appconfig_hosted_configuration_version.
func (ahcv appconfigHostedConfigurationVersionAttributes) ContentType() terra.StringValue {
	return terra.ReferenceAsString(ahcv.ref.Append("content_type"))
}

// Description returns a reference to field description of aws_appconfig_hosted_configuration_version.
func (ahcv appconfigHostedConfigurationVersionAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(ahcv.ref.Append("description"))
}

// Id returns a reference to field id of aws_appconfig_hosted_configuration_version.
func (ahcv appconfigHostedConfigurationVersionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ahcv.ref.Append("id"))
}

// VersionNumber returns a reference to field version_number of aws_appconfig_hosted_configuration_version.
func (ahcv appconfigHostedConfigurationVersionAttributes) VersionNumber() terra.NumberValue {
	return terra.ReferenceAsNumber(ahcv.ref.Append("version_number"))
}

type appconfigHostedConfigurationVersionState struct {
	ApplicationId          string  `json:"application_id"`
	Arn                    string  `json:"arn"`
	ConfigurationProfileId string  `json:"configuration_profile_id"`
	Content                string  `json:"content"`
	ContentType            string  `json:"content_type"`
	Description            string  `json:"description"`
	Id                     string  `json:"id"`
	VersionNumber          float64 `json:"version_number"`
}
