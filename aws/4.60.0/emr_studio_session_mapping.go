// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewEmrStudioSessionMapping creates a new instance of [EmrStudioSessionMapping].
func NewEmrStudioSessionMapping(name string, args EmrStudioSessionMappingArgs) *EmrStudioSessionMapping {
	return &EmrStudioSessionMapping{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*EmrStudioSessionMapping)(nil)

// EmrStudioSessionMapping represents the Terraform resource aws_emr_studio_session_mapping.
type EmrStudioSessionMapping struct {
	Name      string
	Args      EmrStudioSessionMappingArgs
	state     *emrStudioSessionMappingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [EmrStudioSessionMapping].
func (essm *EmrStudioSessionMapping) Type() string {
	return "aws_emr_studio_session_mapping"
}

// LocalName returns the local name for [EmrStudioSessionMapping].
func (essm *EmrStudioSessionMapping) LocalName() string {
	return essm.Name
}

// Configuration returns the configuration (args) for [EmrStudioSessionMapping].
func (essm *EmrStudioSessionMapping) Configuration() interface{} {
	return essm.Args
}

// DependOn is used for other resources to depend on [EmrStudioSessionMapping].
func (essm *EmrStudioSessionMapping) DependOn() terra.Reference {
	return terra.ReferenceResource(essm)
}

// Dependencies returns the list of resources [EmrStudioSessionMapping] depends_on.
func (essm *EmrStudioSessionMapping) Dependencies() terra.Dependencies {
	return essm.DependsOn
}

// LifecycleManagement returns the lifecycle block for [EmrStudioSessionMapping].
func (essm *EmrStudioSessionMapping) LifecycleManagement() *terra.Lifecycle {
	return essm.Lifecycle
}

// Attributes returns the attributes for [EmrStudioSessionMapping].
func (essm *EmrStudioSessionMapping) Attributes() emrStudioSessionMappingAttributes {
	return emrStudioSessionMappingAttributes{ref: terra.ReferenceResource(essm)}
}

// ImportState imports the given attribute values into [EmrStudioSessionMapping]'s state.
func (essm *EmrStudioSessionMapping) ImportState(av io.Reader) error {
	essm.state = &emrStudioSessionMappingState{}
	if err := json.NewDecoder(av).Decode(essm.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", essm.Type(), essm.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [EmrStudioSessionMapping] has state.
func (essm *EmrStudioSessionMapping) State() (*emrStudioSessionMappingState, bool) {
	return essm.state, essm.state != nil
}

// StateMust returns the state for [EmrStudioSessionMapping]. Panics if the state is nil.
func (essm *EmrStudioSessionMapping) StateMust() *emrStudioSessionMappingState {
	if essm.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", essm.Type(), essm.LocalName()))
	}
	return essm.state
}

// EmrStudioSessionMappingArgs contains the configurations for aws_emr_studio_session_mapping.
type EmrStudioSessionMappingArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IdentityId: string, optional
	IdentityId terra.StringValue `hcl:"identity_id,attr"`
	// IdentityName: string, optional
	IdentityName terra.StringValue `hcl:"identity_name,attr"`
	// IdentityType: string, required
	IdentityType terra.StringValue `hcl:"identity_type,attr" validate:"required"`
	// SessionPolicyArn: string, required
	SessionPolicyArn terra.StringValue `hcl:"session_policy_arn,attr" validate:"required"`
	// StudioId: string, required
	StudioId terra.StringValue `hcl:"studio_id,attr" validate:"required"`
}
type emrStudioSessionMappingAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_emr_studio_session_mapping.
func (essm emrStudioSessionMappingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(essm.ref.Append("id"))
}

// IdentityId returns a reference to field identity_id of aws_emr_studio_session_mapping.
func (essm emrStudioSessionMappingAttributes) IdentityId() terra.StringValue {
	return terra.ReferenceAsString(essm.ref.Append("identity_id"))
}

// IdentityName returns a reference to field identity_name of aws_emr_studio_session_mapping.
func (essm emrStudioSessionMappingAttributes) IdentityName() terra.StringValue {
	return terra.ReferenceAsString(essm.ref.Append("identity_name"))
}

// IdentityType returns a reference to field identity_type of aws_emr_studio_session_mapping.
func (essm emrStudioSessionMappingAttributes) IdentityType() terra.StringValue {
	return terra.ReferenceAsString(essm.ref.Append("identity_type"))
}

// SessionPolicyArn returns a reference to field session_policy_arn of aws_emr_studio_session_mapping.
func (essm emrStudioSessionMappingAttributes) SessionPolicyArn() terra.StringValue {
	return terra.ReferenceAsString(essm.ref.Append("session_policy_arn"))
}

// StudioId returns a reference to field studio_id of aws_emr_studio_session_mapping.
func (essm emrStudioSessionMappingAttributes) StudioId() terra.StringValue {
	return terra.ReferenceAsString(essm.ref.Append("studio_id"))
}

type emrStudioSessionMappingState struct {
	Id               string `json:"id"`
	IdentityId       string `json:"identity_id"`
	IdentityName     string `json:"identity_name"`
	IdentityType     string `json:"identity_type"`
	SessionPolicyArn string `json:"session_policy_arn"`
	StudioId         string `json:"studio_id"`
}
