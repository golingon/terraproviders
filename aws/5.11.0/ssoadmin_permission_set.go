// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSsoadminPermissionSet creates a new instance of [SsoadminPermissionSet].
func NewSsoadminPermissionSet(name string, args SsoadminPermissionSetArgs) *SsoadminPermissionSet {
	return &SsoadminPermissionSet{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SsoadminPermissionSet)(nil)

// SsoadminPermissionSet represents the Terraform resource aws_ssoadmin_permission_set.
type SsoadminPermissionSet struct {
	Name      string
	Args      SsoadminPermissionSetArgs
	state     *ssoadminPermissionSetState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SsoadminPermissionSet].
func (sps *SsoadminPermissionSet) Type() string {
	return "aws_ssoadmin_permission_set"
}

// LocalName returns the local name for [SsoadminPermissionSet].
func (sps *SsoadminPermissionSet) LocalName() string {
	return sps.Name
}

// Configuration returns the configuration (args) for [SsoadminPermissionSet].
func (sps *SsoadminPermissionSet) Configuration() interface{} {
	return sps.Args
}

// DependOn is used for other resources to depend on [SsoadminPermissionSet].
func (sps *SsoadminPermissionSet) DependOn() terra.Reference {
	return terra.ReferenceResource(sps)
}

// Dependencies returns the list of resources [SsoadminPermissionSet] depends_on.
func (sps *SsoadminPermissionSet) Dependencies() terra.Dependencies {
	return sps.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SsoadminPermissionSet].
func (sps *SsoadminPermissionSet) LifecycleManagement() *terra.Lifecycle {
	return sps.Lifecycle
}

// Attributes returns the attributes for [SsoadminPermissionSet].
func (sps *SsoadminPermissionSet) Attributes() ssoadminPermissionSetAttributes {
	return ssoadminPermissionSetAttributes{ref: terra.ReferenceResource(sps)}
}

// ImportState imports the given attribute values into [SsoadminPermissionSet]'s state.
func (sps *SsoadminPermissionSet) ImportState(av io.Reader) error {
	sps.state = &ssoadminPermissionSetState{}
	if err := json.NewDecoder(av).Decode(sps.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sps.Type(), sps.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SsoadminPermissionSet] has state.
func (sps *SsoadminPermissionSet) State() (*ssoadminPermissionSetState, bool) {
	return sps.state, sps.state != nil
}

// StateMust returns the state for [SsoadminPermissionSet]. Panics if the state is nil.
func (sps *SsoadminPermissionSet) StateMust() *ssoadminPermissionSetState {
	if sps.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sps.Type(), sps.LocalName()))
	}
	return sps.state
}

// SsoadminPermissionSetArgs contains the configurations for aws_ssoadmin_permission_set.
type SsoadminPermissionSetArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InstanceArn: string, required
	InstanceArn terra.StringValue `hcl:"instance_arn,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// RelayState: string, optional
	RelayState terra.StringValue `hcl:"relay_state,attr"`
	// SessionDuration: string, optional
	SessionDuration terra.StringValue `hcl:"session_duration,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
}
type ssoadminPermissionSetAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_ssoadmin_permission_set.
func (sps ssoadminPermissionSetAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(sps.ref.Append("arn"))
}

// CreatedDate returns a reference to field created_date of aws_ssoadmin_permission_set.
func (sps ssoadminPermissionSetAttributes) CreatedDate() terra.StringValue {
	return terra.ReferenceAsString(sps.ref.Append("created_date"))
}

// Description returns a reference to field description of aws_ssoadmin_permission_set.
func (sps ssoadminPermissionSetAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(sps.ref.Append("description"))
}

// Id returns a reference to field id of aws_ssoadmin_permission_set.
func (sps ssoadminPermissionSetAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sps.ref.Append("id"))
}

// InstanceArn returns a reference to field instance_arn of aws_ssoadmin_permission_set.
func (sps ssoadminPermissionSetAttributes) InstanceArn() terra.StringValue {
	return terra.ReferenceAsString(sps.ref.Append("instance_arn"))
}

// Name returns a reference to field name of aws_ssoadmin_permission_set.
func (sps ssoadminPermissionSetAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sps.ref.Append("name"))
}

// RelayState returns a reference to field relay_state of aws_ssoadmin_permission_set.
func (sps ssoadminPermissionSetAttributes) RelayState() terra.StringValue {
	return terra.ReferenceAsString(sps.ref.Append("relay_state"))
}

// SessionDuration returns a reference to field session_duration of aws_ssoadmin_permission_set.
func (sps ssoadminPermissionSetAttributes) SessionDuration() terra.StringValue {
	return terra.ReferenceAsString(sps.ref.Append("session_duration"))
}

// Tags returns a reference to field tags of aws_ssoadmin_permission_set.
func (sps ssoadminPermissionSetAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](sps.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_ssoadmin_permission_set.
func (sps ssoadminPermissionSetAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](sps.ref.Append("tags_all"))
}

type ssoadminPermissionSetState struct {
	Arn             string            `json:"arn"`
	CreatedDate     string            `json:"created_date"`
	Description     string            `json:"description"`
	Id              string            `json:"id"`
	InstanceArn     string            `json:"instance_arn"`
	Name            string            `json:"name"`
	RelayState      string            `json:"relay_state"`
	SessionDuration string            `json:"session_duration"`
	Tags            map[string]string `json:"tags"`
	TagsAll         map[string]string `json:"tags_all"`
}