// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	appconfigconfigurationprofile "github.com/golingon/terraproviders/aws/5.11.0/appconfigconfigurationprofile"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewAppconfigConfigurationProfile creates a new instance of [AppconfigConfigurationProfile].
func NewAppconfigConfigurationProfile(name string, args AppconfigConfigurationProfileArgs) *AppconfigConfigurationProfile {
	return &AppconfigConfigurationProfile{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AppconfigConfigurationProfile)(nil)

// AppconfigConfigurationProfile represents the Terraform resource aws_appconfig_configuration_profile.
type AppconfigConfigurationProfile struct {
	Name      string
	Args      AppconfigConfigurationProfileArgs
	state     *appconfigConfigurationProfileState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [AppconfigConfigurationProfile].
func (acp *AppconfigConfigurationProfile) Type() string {
	return "aws_appconfig_configuration_profile"
}

// LocalName returns the local name for [AppconfigConfigurationProfile].
func (acp *AppconfigConfigurationProfile) LocalName() string {
	return acp.Name
}

// Configuration returns the configuration (args) for [AppconfigConfigurationProfile].
func (acp *AppconfigConfigurationProfile) Configuration() interface{} {
	return acp.Args
}

// DependOn is used for other resources to depend on [AppconfigConfigurationProfile].
func (acp *AppconfigConfigurationProfile) DependOn() terra.Reference {
	return terra.ReferenceResource(acp)
}

// Dependencies returns the list of resources [AppconfigConfigurationProfile] depends_on.
func (acp *AppconfigConfigurationProfile) Dependencies() terra.Dependencies {
	return acp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [AppconfigConfigurationProfile].
func (acp *AppconfigConfigurationProfile) LifecycleManagement() *terra.Lifecycle {
	return acp.Lifecycle
}

// Attributes returns the attributes for [AppconfigConfigurationProfile].
func (acp *AppconfigConfigurationProfile) Attributes() appconfigConfigurationProfileAttributes {
	return appconfigConfigurationProfileAttributes{ref: terra.ReferenceResource(acp)}
}

// ImportState imports the given attribute values into [AppconfigConfigurationProfile]'s state.
func (acp *AppconfigConfigurationProfile) ImportState(av io.Reader) error {
	acp.state = &appconfigConfigurationProfileState{}
	if err := json.NewDecoder(av).Decode(acp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", acp.Type(), acp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [AppconfigConfigurationProfile] has state.
func (acp *AppconfigConfigurationProfile) State() (*appconfigConfigurationProfileState, bool) {
	return acp.state, acp.state != nil
}

// StateMust returns the state for [AppconfigConfigurationProfile]. Panics if the state is nil.
func (acp *AppconfigConfigurationProfile) StateMust() *appconfigConfigurationProfileState {
	if acp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", acp.Type(), acp.LocalName()))
	}
	return acp.state
}

// AppconfigConfigurationProfileArgs contains the configurations for aws_appconfig_configuration_profile.
type AppconfigConfigurationProfileArgs struct {
	// ApplicationId: string, required
	ApplicationId terra.StringValue `hcl:"application_id,attr" validate:"required"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LocationUri: string, required
	LocationUri terra.StringValue `hcl:"location_uri,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// RetrievalRoleArn: string, optional
	RetrievalRoleArn terra.StringValue `hcl:"retrieval_role_arn,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// Type: string, optional
	Type terra.StringValue `hcl:"type,attr"`
	// Validator: min=0,max=2
	Validator []appconfigconfigurationprofile.Validator `hcl:"validator,block" validate:"min=0,max=2"`
}
type appconfigConfigurationProfileAttributes struct {
	ref terra.Reference
}

// ApplicationId returns a reference to field application_id of aws_appconfig_configuration_profile.
func (acp appconfigConfigurationProfileAttributes) ApplicationId() terra.StringValue {
	return terra.ReferenceAsString(acp.ref.Append("application_id"))
}

// Arn returns a reference to field arn of aws_appconfig_configuration_profile.
func (acp appconfigConfigurationProfileAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(acp.ref.Append("arn"))
}

// ConfigurationProfileId returns a reference to field configuration_profile_id of aws_appconfig_configuration_profile.
func (acp appconfigConfigurationProfileAttributes) ConfigurationProfileId() terra.StringValue {
	return terra.ReferenceAsString(acp.ref.Append("configuration_profile_id"))
}

// Description returns a reference to field description of aws_appconfig_configuration_profile.
func (acp appconfigConfigurationProfileAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(acp.ref.Append("description"))
}

// Id returns a reference to field id of aws_appconfig_configuration_profile.
func (acp appconfigConfigurationProfileAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(acp.ref.Append("id"))
}

// LocationUri returns a reference to field location_uri of aws_appconfig_configuration_profile.
func (acp appconfigConfigurationProfileAttributes) LocationUri() terra.StringValue {
	return terra.ReferenceAsString(acp.ref.Append("location_uri"))
}

// Name returns a reference to field name of aws_appconfig_configuration_profile.
func (acp appconfigConfigurationProfileAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(acp.ref.Append("name"))
}

// RetrievalRoleArn returns a reference to field retrieval_role_arn of aws_appconfig_configuration_profile.
func (acp appconfigConfigurationProfileAttributes) RetrievalRoleArn() terra.StringValue {
	return terra.ReferenceAsString(acp.ref.Append("retrieval_role_arn"))
}

// Tags returns a reference to field tags of aws_appconfig_configuration_profile.
func (acp appconfigConfigurationProfileAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](acp.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_appconfig_configuration_profile.
func (acp appconfigConfigurationProfileAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](acp.ref.Append("tags_all"))
}

// Type returns a reference to field type of aws_appconfig_configuration_profile.
func (acp appconfigConfigurationProfileAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(acp.ref.Append("type"))
}

func (acp appconfigConfigurationProfileAttributes) Validator() terra.SetValue[appconfigconfigurationprofile.ValidatorAttributes] {
	return terra.ReferenceAsSet[appconfigconfigurationprofile.ValidatorAttributes](acp.ref.Append("validator"))
}

type appconfigConfigurationProfileState struct {
	ApplicationId          string                                         `json:"application_id"`
	Arn                    string                                         `json:"arn"`
	ConfigurationProfileId string                                         `json:"configuration_profile_id"`
	Description            string                                         `json:"description"`
	Id                     string                                         `json:"id"`
	LocationUri            string                                         `json:"location_uri"`
	Name                   string                                         `json:"name"`
	RetrievalRoleArn       string                                         `json:"retrieval_role_arn"`
	Tags                   map[string]string                              `json:"tags"`
	TagsAll                map[string]string                              `json:"tags_all"`
	Type                   string                                         `json:"type"`
	Validator              []appconfigconfigurationprofile.ValidatorState `json:"validator"`
}
