// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	sagemakerapp "github.com/golingon/terraproviders/aws/5.8.0/sagemakerapp"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSagemakerApp creates a new instance of [SagemakerApp].
func NewSagemakerApp(name string, args SagemakerAppArgs) *SagemakerApp {
	return &SagemakerApp{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SagemakerApp)(nil)

// SagemakerApp represents the Terraform resource aws_sagemaker_app.
type SagemakerApp struct {
	Name      string
	Args      SagemakerAppArgs
	state     *sagemakerAppState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SagemakerApp].
func (sa *SagemakerApp) Type() string {
	return "aws_sagemaker_app"
}

// LocalName returns the local name for [SagemakerApp].
func (sa *SagemakerApp) LocalName() string {
	return sa.Name
}

// Configuration returns the configuration (args) for [SagemakerApp].
func (sa *SagemakerApp) Configuration() interface{} {
	return sa.Args
}

// DependOn is used for other resources to depend on [SagemakerApp].
func (sa *SagemakerApp) DependOn() terra.Reference {
	return terra.ReferenceResource(sa)
}

// Dependencies returns the list of resources [SagemakerApp] depends_on.
func (sa *SagemakerApp) Dependencies() terra.Dependencies {
	return sa.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SagemakerApp].
func (sa *SagemakerApp) LifecycleManagement() *terra.Lifecycle {
	return sa.Lifecycle
}

// Attributes returns the attributes for [SagemakerApp].
func (sa *SagemakerApp) Attributes() sagemakerAppAttributes {
	return sagemakerAppAttributes{ref: terra.ReferenceResource(sa)}
}

// ImportState imports the given attribute values into [SagemakerApp]'s state.
func (sa *SagemakerApp) ImportState(av io.Reader) error {
	sa.state = &sagemakerAppState{}
	if err := json.NewDecoder(av).Decode(sa.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sa.Type(), sa.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SagemakerApp] has state.
func (sa *SagemakerApp) State() (*sagemakerAppState, bool) {
	return sa.state, sa.state != nil
}

// StateMust returns the state for [SagemakerApp]. Panics if the state is nil.
func (sa *SagemakerApp) StateMust() *sagemakerAppState {
	if sa.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sa.Type(), sa.LocalName()))
	}
	return sa.state
}

// SagemakerAppArgs contains the configurations for aws_sagemaker_app.
type SagemakerAppArgs struct {
	// AppName: string, required
	AppName terra.StringValue `hcl:"app_name,attr" validate:"required"`
	// AppType: string, required
	AppType terra.StringValue `hcl:"app_type,attr" validate:"required"`
	// DomainId: string, required
	DomainId terra.StringValue `hcl:"domain_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// SpaceName: string, optional
	SpaceName terra.StringValue `hcl:"space_name,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// UserProfileName: string, optional
	UserProfileName terra.StringValue `hcl:"user_profile_name,attr"`
	// ResourceSpec: optional
	ResourceSpec *sagemakerapp.ResourceSpec `hcl:"resource_spec,block"`
}
type sagemakerAppAttributes struct {
	ref terra.Reference
}

// AppName returns a reference to field app_name of aws_sagemaker_app.
func (sa sagemakerAppAttributes) AppName() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("app_name"))
}

// AppType returns a reference to field app_type of aws_sagemaker_app.
func (sa sagemakerAppAttributes) AppType() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("app_type"))
}

// Arn returns a reference to field arn of aws_sagemaker_app.
func (sa sagemakerAppAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("arn"))
}

// DomainId returns a reference to field domain_id of aws_sagemaker_app.
func (sa sagemakerAppAttributes) DomainId() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("domain_id"))
}

// Id returns a reference to field id of aws_sagemaker_app.
func (sa sagemakerAppAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("id"))
}

// SpaceName returns a reference to field space_name of aws_sagemaker_app.
func (sa sagemakerAppAttributes) SpaceName() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("space_name"))
}

// Tags returns a reference to field tags of aws_sagemaker_app.
func (sa sagemakerAppAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](sa.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_sagemaker_app.
func (sa sagemakerAppAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](sa.ref.Append("tags_all"))
}

// UserProfileName returns a reference to field user_profile_name of aws_sagemaker_app.
func (sa sagemakerAppAttributes) UserProfileName() terra.StringValue {
	return terra.ReferenceAsString(sa.ref.Append("user_profile_name"))
}

func (sa sagemakerAppAttributes) ResourceSpec() terra.ListValue[sagemakerapp.ResourceSpecAttributes] {
	return terra.ReferenceAsList[sagemakerapp.ResourceSpecAttributes](sa.ref.Append("resource_spec"))
}

type sagemakerAppState struct {
	AppName         string                           `json:"app_name"`
	AppType         string                           `json:"app_type"`
	Arn             string                           `json:"arn"`
	DomainId        string                           `json:"domain_id"`
	Id              string                           `json:"id"`
	SpaceName       string                           `json:"space_name"`
	Tags            map[string]string                `json:"tags"`
	TagsAll         map[string]string                `json:"tags_all"`
	UserProfileName string                           `json:"user_profile_name"`
	ResourceSpec    []sagemakerapp.ResourceSpecState `json:"resource_spec"`
}
