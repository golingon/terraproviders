// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	sagemakerworkteam "github.com/golingon/terraproviders/aws/5.7.0/sagemakerworkteam"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSagemakerWorkteam creates a new instance of [SagemakerWorkteam].
func NewSagemakerWorkteam(name string, args SagemakerWorkteamArgs) *SagemakerWorkteam {
	return &SagemakerWorkteam{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SagemakerWorkteam)(nil)

// SagemakerWorkteam represents the Terraform resource aws_sagemaker_workteam.
type SagemakerWorkteam struct {
	Name      string
	Args      SagemakerWorkteamArgs
	state     *sagemakerWorkteamState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SagemakerWorkteam].
func (sw *SagemakerWorkteam) Type() string {
	return "aws_sagemaker_workteam"
}

// LocalName returns the local name for [SagemakerWorkteam].
func (sw *SagemakerWorkteam) LocalName() string {
	return sw.Name
}

// Configuration returns the configuration (args) for [SagemakerWorkteam].
func (sw *SagemakerWorkteam) Configuration() interface{} {
	return sw.Args
}

// DependOn is used for other resources to depend on [SagemakerWorkteam].
func (sw *SagemakerWorkteam) DependOn() terra.Reference {
	return terra.ReferenceResource(sw)
}

// Dependencies returns the list of resources [SagemakerWorkteam] depends_on.
func (sw *SagemakerWorkteam) Dependencies() terra.Dependencies {
	return sw.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SagemakerWorkteam].
func (sw *SagemakerWorkteam) LifecycleManagement() *terra.Lifecycle {
	return sw.Lifecycle
}

// Attributes returns the attributes for [SagemakerWorkteam].
func (sw *SagemakerWorkteam) Attributes() sagemakerWorkteamAttributes {
	return sagemakerWorkteamAttributes{ref: terra.ReferenceResource(sw)}
}

// ImportState imports the given attribute values into [SagemakerWorkteam]'s state.
func (sw *SagemakerWorkteam) ImportState(av io.Reader) error {
	sw.state = &sagemakerWorkteamState{}
	if err := json.NewDecoder(av).Decode(sw.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sw.Type(), sw.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SagemakerWorkteam] has state.
func (sw *SagemakerWorkteam) State() (*sagemakerWorkteamState, bool) {
	return sw.state, sw.state != nil
}

// StateMust returns the state for [SagemakerWorkteam]. Panics if the state is nil.
func (sw *SagemakerWorkteam) StateMust() *sagemakerWorkteamState {
	if sw.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sw.Type(), sw.LocalName()))
	}
	return sw.state
}

// SagemakerWorkteamArgs contains the configurations for aws_sagemaker_workteam.
type SagemakerWorkteamArgs struct {
	// Description: string, required
	Description terra.StringValue `hcl:"description,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// WorkforceName: string, required
	WorkforceName terra.StringValue `hcl:"workforce_name,attr" validate:"required"`
	// WorkteamName: string, required
	WorkteamName terra.StringValue `hcl:"workteam_name,attr" validate:"required"`
	// MemberDefinition: min=1,max=10
	MemberDefinition []sagemakerworkteam.MemberDefinition `hcl:"member_definition,block" validate:"min=1,max=10"`
	// NotificationConfiguration: optional
	NotificationConfiguration *sagemakerworkteam.NotificationConfiguration `hcl:"notification_configuration,block"`
}
type sagemakerWorkteamAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_sagemaker_workteam.
func (sw sagemakerWorkteamAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(sw.ref.Append("arn"))
}

// Description returns a reference to field description of aws_sagemaker_workteam.
func (sw sagemakerWorkteamAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(sw.ref.Append("description"))
}

// Id returns a reference to field id of aws_sagemaker_workteam.
func (sw sagemakerWorkteamAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sw.ref.Append("id"))
}

// Subdomain returns a reference to field subdomain of aws_sagemaker_workteam.
func (sw sagemakerWorkteamAttributes) Subdomain() terra.StringValue {
	return terra.ReferenceAsString(sw.ref.Append("subdomain"))
}

// Tags returns a reference to field tags of aws_sagemaker_workteam.
func (sw sagemakerWorkteamAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](sw.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_sagemaker_workteam.
func (sw sagemakerWorkteamAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](sw.ref.Append("tags_all"))
}

// WorkforceName returns a reference to field workforce_name of aws_sagemaker_workteam.
func (sw sagemakerWorkteamAttributes) WorkforceName() terra.StringValue {
	return terra.ReferenceAsString(sw.ref.Append("workforce_name"))
}

// WorkteamName returns a reference to field workteam_name of aws_sagemaker_workteam.
func (sw sagemakerWorkteamAttributes) WorkteamName() terra.StringValue {
	return terra.ReferenceAsString(sw.ref.Append("workteam_name"))
}

func (sw sagemakerWorkteamAttributes) MemberDefinition() terra.ListValue[sagemakerworkteam.MemberDefinitionAttributes] {
	return terra.ReferenceAsList[sagemakerworkteam.MemberDefinitionAttributes](sw.ref.Append("member_definition"))
}

func (sw sagemakerWorkteamAttributes) NotificationConfiguration() terra.ListValue[sagemakerworkteam.NotificationConfigurationAttributes] {
	return terra.ReferenceAsList[sagemakerworkteam.NotificationConfigurationAttributes](sw.ref.Append("notification_configuration"))
}

type sagemakerWorkteamState struct {
	Arn                       string                                             `json:"arn"`
	Description               string                                             `json:"description"`
	Id                        string                                             `json:"id"`
	Subdomain                 string                                             `json:"subdomain"`
	Tags                      map[string]string                                  `json:"tags"`
	TagsAll                   map[string]string                                  `json:"tags_all"`
	WorkforceName             string                                             `json:"workforce_name"`
	WorkteamName              string                                             `json:"workteam_name"`
	MemberDefinition          []sagemakerworkteam.MemberDefinitionState          `json:"member_definition"`
	NotificationConfiguration []sagemakerworkteam.NotificationConfigurationState `json:"notification_configuration"`
}
