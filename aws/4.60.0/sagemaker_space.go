// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	sagemakerspace "github.com/golingon/terraproviders/aws/4.60.0/sagemakerspace"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSagemakerSpace creates a new instance of [SagemakerSpace].
func NewSagemakerSpace(name string, args SagemakerSpaceArgs) *SagemakerSpace {
	return &SagemakerSpace{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SagemakerSpace)(nil)

// SagemakerSpace represents the Terraform resource aws_sagemaker_space.
type SagemakerSpace struct {
	Name      string
	Args      SagemakerSpaceArgs
	state     *sagemakerSpaceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SagemakerSpace].
func (ss *SagemakerSpace) Type() string {
	return "aws_sagemaker_space"
}

// LocalName returns the local name for [SagemakerSpace].
func (ss *SagemakerSpace) LocalName() string {
	return ss.Name
}

// Configuration returns the configuration (args) for [SagemakerSpace].
func (ss *SagemakerSpace) Configuration() interface{} {
	return ss.Args
}

// DependOn is used for other resources to depend on [SagemakerSpace].
func (ss *SagemakerSpace) DependOn() terra.Reference {
	return terra.ReferenceResource(ss)
}

// Dependencies returns the list of resources [SagemakerSpace] depends_on.
func (ss *SagemakerSpace) Dependencies() terra.Dependencies {
	return ss.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SagemakerSpace].
func (ss *SagemakerSpace) LifecycleManagement() *terra.Lifecycle {
	return ss.Lifecycle
}

// Attributes returns the attributes for [SagemakerSpace].
func (ss *SagemakerSpace) Attributes() sagemakerSpaceAttributes {
	return sagemakerSpaceAttributes{ref: terra.ReferenceResource(ss)}
}

// ImportState imports the given attribute values into [SagemakerSpace]'s state.
func (ss *SagemakerSpace) ImportState(av io.Reader) error {
	ss.state = &sagemakerSpaceState{}
	if err := json.NewDecoder(av).Decode(ss.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ss.Type(), ss.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SagemakerSpace] has state.
func (ss *SagemakerSpace) State() (*sagemakerSpaceState, bool) {
	return ss.state, ss.state != nil
}

// StateMust returns the state for [SagemakerSpace]. Panics if the state is nil.
func (ss *SagemakerSpace) StateMust() *sagemakerSpaceState {
	if ss.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ss.Type(), ss.LocalName()))
	}
	return ss.state
}

// SagemakerSpaceArgs contains the configurations for aws_sagemaker_space.
type SagemakerSpaceArgs struct {
	// DomainId: string, required
	DomainId terra.StringValue `hcl:"domain_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// SpaceName: string, required
	SpaceName terra.StringValue `hcl:"space_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// SpaceSettings: optional
	SpaceSettings *sagemakerspace.SpaceSettings `hcl:"space_settings,block"`
}
type sagemakerSpaceAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_sagemaker_space.
func (ss sagemakerSpaceAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ss.ref.Append("arn"))
}

// DomainId returns a reference to field domain_id of aws_sagemaker_space.
func (ss sagemakerSpaceAttributes) DomainId() terra.StringValue {
	return terra.ReferenceAsString(ss.ref.Append("domain_id"))
}

// HomeEfsFileSystemUid returns a reference to field home_efs_file_system_uid of aws_sagemaker_space.
func (ss sagemakerSpaceAttributes) HomeEfsFileSystemUid() terra.StringValue {
	return terra.ReferenceAsString(ss.ref.Append("home_efs_file_system_uid"))
}

// Id returns a reference to field id of aws_sagemaker_space.
func (ss sagemakerSpaceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ss.ref.Append("id"))
}

// SpaceName returns a reference to field space_name of aws_sagemaker_space.
func (ss sagemakerSpaceAttributes) SpaceName() terra.StringValue {
	return terra.ReferenceAsString(ss.ref.Append("space_name"))
}

// Tags returns a reference to field tags of aws_sagemaker_space.
func (ss sagemakerSpaceAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ss.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_sagemaker_space.
func (ss sagemakerSpaceAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ss.ref.Append("tags_all"))
}

func (ss sagemakerSpaceAttributes) SpaceSettings() terra.ListValue[sagemakerspace.SpaceSettingsAttributes] {
	return terra.ReferenceAsList[sagemakerspace.SpaceSettingsAttributes](ss.ref.Append("space_settings"))
}

type sagemakerSpaceState struct {
	Arn                  string                              `json:"arn"`
	DomainId             string                              `json:"domain_id"`
	HomeEfsFileSystemUid string                              `json:"home_efs_file_system_uid"`
	Id                   string                              `json:"id"`
	SpaceName            string                              `json:"space_name"`
	Tags                 map[string]string                   `json:"tags"`
	TagsAll              map[string]string                   `json:"tags_all"`
	SpaceSettings        []sagemakerspace.SpaceSettingsState `json:"space_settings"`
}
