// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewRolesanywhereProfile creates a new instance of [RolesanywhereProfile].
func NewRolesanywhereProfile(name string, args RolesanywhereProfileArgs) *RolesanywhereProfile {
	return &RolesanywhereProfile{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*RolesanywhereProfile)(nil)

// RolesanywhereProfile represents the Terraform resource aws_rolesanywhere_profile.
type RolesanywhereProfile struct {
	Name      string
	Args      RolesanywhereProfileArgs
	state     *rolesanywhereProfileState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [RolesanywhereProfile].
func (rp *RolesanywhereProfile) Type() string {
	return "aws_rolesanywhere_profile"
}

// LocalName returns the local name for [RolesanywhereProfile].
func (rp *RolesanywhereProfile) LocalName() string {
	return rp.Name
}

// Configuration returns the configuration (args) for [RolesanywhereProfile].
func (rp *RolesanywhereProfile) Configuration() interface{} {
	return rp.Args
}

// DependOn is used for other resources to depend on [RolesanywhereProfile].
func (rp *RolesanywhereProfile) DependOn() terra.Reference {
	return terra.ReferenceResource(rp)
}

// Dependencies returns the list of resources [RolesanywhereProfile] depends_on.
func (rp *RolesanywhereProfile) Dependencies() terra.Dependencies {
	return rp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [RolesanywhereProfile].
func (rp *RolesanywhereProfile) LifecycleManagement() *terra.Lifecycle {
	return rp.Lifecycle
}

// Attributes returns the attributes for [RolesanywhereProfile].
func (rp *RolesanywhereProfile) Attributes() rolesanywhereProfileAttributes {
	return rolesanywhereProfileAttributes{ref: terra.ReferenceResource(rp)}
}

// ImportState imports the given attribute values into [RolesanywhereProfile]'s state.
func (rp *RolesanywhereProfile) ImportState(av io.Reader) error {
	rp.state = &rolesanywhereProfileState{}
	if err := json.NewDecoder(av).Decode(rp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", rp.Type(), rp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [RolesanywhereProfile] has state.
func (rp *RolesanywhereProfile) State() (*rolesanywhereProfileState, bool) {
	return rp.state, rp.state != nil
}

// StateMust returns the state for [RolesanywhereProfile]. Panics if the state is nil.
func (rp *RolesanywhereProfile) StateMust() *rolesanywhereProfileState {
	if rp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", rp.Type(), rp.LocalName()))
	}
	return rp.state
}

// RolesanywhereProfileArgs contains the configurations for aws_rolesanywhere_profile.
type RolesanywhereProfileArgs struct {
	// DurationSeconds: number, optional
	DurationSeconds terra.NumberValue `hcl:"duration_seconds,attr"`
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ManagedPolicyArns: set of string, optional
	ManagedPolicyArns terra.SetValue[terra.StringValue] `hcl:"managed_policy_arns,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// RequireInstanceProperties: bool, optional
	RequireInstanceProperties terra.BoolValue `hcl:"require_instance_properties,attr"`
	// RoleArns: set of string, required
	RoleArns terra.SetValue[terra.StringValue] `hcl:"role_arns,attr" validate:"required"`
	// SessionPolicy: string, optional
	SessionPolicy terra.StringValue `hcl:"session_policy,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
}
type rolesanywhereProfileAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_rolesanywhere_profile.
func (rp rolesanywhereProfileAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(rp.ref.Append("arn"))
}

// DurationSeconds returns a reference to field duration_seconds of aws_rolesanywhere_profile.
func (rp rolesanywhereProfileAttributes) DurationSeconds() terra.NumberValue {
	return terra.ReferenceAsNumber(rp.ref.Append("duration_seconds"))
}

// Enabled returns a reference to field enabled of aws_rolesanywhere_profile.
func (rp rolesanywhereProfileAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(rp.ref.Append("enabled"))
}

// Id returns a reference to field id of aws_rolesanywhere_profile.
func (rp rolesanywhereProfileAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(rp.ref.Append("id"))
}

// ManagedPolicyArns returns a reference to field managed_policy_arns of aws_rolesanywhere_profile.
func (rp rolesanywhereProfileAttributes) ManagedPolicyArns() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](rp.ref.Append("managed_policy_arns"))
}

// Name returns a reference to field name of aws_rolesanywhere_profile.
func (rp rolesanywhereProfileAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(rp.ref.Append("name"))
}

// RequireInstanceProperties returns a reference to field require_instance_properties of aws_rolesanywhere_profile.
func (rp rolesanywhereProfileAttributes) RequireInstanceProperties() terra.BoolValue {
	return terra.ReferenceAsBool(rp.ref.Append("require_instance_properties"))
}

// RoleArns returns a reference to field role_arns of aws_rolesanywhere_profile.
func (rp rolesanywhereProfileAttributes) RoleArns() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](rp.ref.Append("role_arns"))
}

// SessionPolicy returns a reference to field session_policy of aws_rolesanywhere_profile.
func (rp rolesanywhereProfileAttributes) SessionPolicy() terra.StringValue {
	return terra.ReferenceAsString(rp.ref.Append("session_policy"))
}

// Tags returns a reference to field tags of aws_rolesanywhere_profile.
func (rp rolesanywhereProfileAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](rp.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_rolesanywhere_profile.
func (rp rolesanywhereProfileAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](rp.ref.Append("tags_all"))
}

type rolesanywhereProfileState struct {
	Arn                       string            `json:"arn"`
	DurationSeconds           float64           `json:"duration_seconds"`
	Enabled                   bool              `json:"enabled"`
	Id                        string            `json:"id"`
	ManagedPolicyArns         []string          `json:"managed_policy_arns"`
	Name                      string            `json:"name"`
	RequireInstanceProperties bool              `json:"require_instance_properties"`
	RoleArns                  []string          `json:"role_arns"`
	SessionPolicy             string            `json:"session_policy"`
	Tags                      map[string]string `json:"tags"`
	TagsAll                   map[string]string `json:"tags_all"`
}
