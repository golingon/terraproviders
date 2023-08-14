// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	iamrole "github.com/golingon/terraproviders/aws/5.12.0/iamrole"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewIamRole creates a new instance of [IamRole].
func NewIamRole(name string, args IamRoleArgs) *IamRole {
	return &IamRole{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*IamRole)(nil)

// IamRole represents the Terraform resource aws_iam_role.
type IamRole struct {
	Name      string
	Args      IamRoleArgs
	state     *iamRoleState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [IamRole].
func (ir *IamRole) Type() string {
	return "aws_iam_role"
}

// LocalName returns the local name for [IamRole].
func (ir *IamRole) LocalName() string {
	return ir.Name
}

// Configuration returns the configuration (args) for [IamRole].
func (ir *IamRole) Configuration() interface{} {
	return ir.Args
}

// DependOn is used for other resources to depend on [IamRole].
func (ir *IamRole) DependOn() terra.Reference {
	return terra.ReferenceResource(ir)
}

// Dependencies returns the list of resources [IamRole] depends_on.
func (ir *IamRole) Dependencies() terra.Dependencies {
	return ir.DependsOn
}

// LifecycleManagement returns the lifecycle block for [IamRole].
func (ir *IamRole) LifecycleManagement() *terra.Lifecycle {
	return ir.Lifecycle
}

// Attributes returns the attributes for [IamRole].
func (ir *IamRole) Attributes() iamRoleAttributes {
	return iamRoleAttributes{ref: terra.ReferenceResource(ir)}
}

// ImportState imports the given attribute values into [IamRole]'s state.
func (ir *IamRole) ImportState(av io.Reader) error {
	ir.state = &iamRoleState{}
	if err := json.NewDecoder(av).Decode(ir.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ir.Type(), ir.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [IamRole] has state.
func (ir *IamRole) State() (*iamRoleState, bool) {
	return ir.state, ir.state != nil
}

// StateMust returns the state for [IamRole]. Panics if the state is nil.
func (ir *IamRole) StateMust() *iamRoleState {
	if ir.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ir.Type(), ir.LocalName()))
	}
	return ir.state
}

// IamRoleArgs contains the configurations for aws_iam_role.
type IamRoleArgs struct {
	// AssumeRolePolicy: string, required
	AssumeRolePolicy terra.StringValue `hcl:"assume_role_policy,attr" validate:"required"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// ForceDetachPolicies: bool, optional
	ForceDetachPolicies terra.BoolValue `hcl:"force_detach_policies,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ManagedPolicyArns: set of string, optional
	ManagedPolicyArns terra.SetValue[terra.StringValue] `hcl:"managed_policy_arns,attr"`
	// MaxSessionDuration: number, optional
	MaxSessionDuration terra.NumberValue `hcl:"max_session_duration,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// NamePrefix: string, optional
	NamePrefix terra.StringValue `hcl:"name_prefix,attr"`
	// Path: string, optional
	Path terra.StringValue `hcl:"path,attr"`
	// PermissionsBoundary: string, optional
	PermissionsBoundary terra.StringValue `hcl:"permissions_boundary,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// InlinePolicy: min=0
	InlinePolicy []iamrole.InlinePolicy `hcl:"inline_policy,block" validate:"min=0"`
}
type iamRoleAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_iam_role.
func (ir iamRoleAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ir.ref.Append("arn"))
}

// AssumeRolePolicy returns a reference to field assume_role_policy of aws_iam_role.
func (ir iamRoleAttributes) AssumeRolePolicy() terra.StringValue {
	return terra.ReferenceAsString(ir.ref.Append("assume_role_policy"))
}

// CreateDate returns a reference to field create_date of aws_iam_role.
func (ir iamRoleAttributes) CreateDate() terra.StringValue {
	return terra.ReferenceAsString(ir.ref.Append("create_date"))
}

// Description returns a reference to field description of aws_iam_role.
func (ir iamRoleAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(ir.ref.Append("description"))
}

// ForceDetachPolicies returns a reference to field force_detach_policies of aws_iam_role.
func (ir iamRoleAttributes) ForceDetachPolicies() terra.BoolValue {
	return terra.ReferenceAsBool(ir.ref.Append("force_detach_policies"))
}

// Id returns a reference to field id of aws_iam_role.
func (ir iamRoleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ir.ref.Append("id"))
}

// ManagedPolicyArns returns a reference to field managed_policy_arns of aws_iam_role.
func (ir iamRoleAttributes) ManagedPolicyArns() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ir.ref.Append("managed_policy_arns"))
}

// MaxSessionDuration returns a reference to field max_session_duration of aws_iam_role.
func (ir iamRoleAttributes) MaxSessionDuration() terra.NumberValue {
	return terra.ReferenceAsNumber(ir.ref.Append("max_session_duration"))
}

// Name returns a reference to field name of aws_iam_role.
func (ir iamRoleAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ir.ref.Append("name"))
}

// NamePrefix returns a reference to field name_prefix of aws_iam_role.
func (ir iamRoleAttributes) NamePrefix() terra.StringValue {
	return terra.ReferenceAsString(ir.ref.Append("name_prefix"))
}

// Path returns a reference to field path of aws_iam_role.
func (ir iamRoleAttributes) Path() terra.StringValue {
	return terra.ReferenceAsString(ir.ref.Append("path"))
}

// PermissionsBoundary returns a reference to field permissions_boundary of aws_iam_role.
func (ir iamRoleAttributes) PermissionsBoundary() terra.StringValue {
	return terra.ReferenceAsString(ir.ref.Append("permissions_boundary"))
}

// Tags returns a reference to field tags of aws_iam_role.
func (ir iamRoleAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ir.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_iam_role.
func (ir iamRoleAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ir.ref.Append("tags_all"))
}

// UniqueId returns a reference to field unique_id of aws_iam_role.
func (ir iamRoleAttributes) UniqueId() terra.StringValue {
	return terra.ReferenceAsString(ir.ref.Append("unique_id"))
}

func (ir iamRoleAttributes) InlinePolicy() terra.SetValue[iamrole.InlinePolicyAttributes] {
	return terra.ReferenceAsSet[iamrole.InlinePolicyAttributes](ir.ref.Append("inline_policy"))
}

type iamRoleState struct {
	Arn                 string                      `json:"arn"`
	AssumeRolePolicy    string                      `json:"assume_role_policy"`
	CreateDate          string                      `json:"create_date"`
	Description         string                      `json:"description"`
	ForceDetachPolicies bool                        `json:"force_detach_policies"`
	Id                  string                      `json:"id"`
	ManagedPolicyArns   []string                    `json:"managed_policy_arns"`
	MaxSessionDuration  float64                     `json:"max_session_duration"`
	Name                string                      `json:"name"`
	NamePrefix          string                      `json:"name_prefix"`
	Path                string                      `json:"path"`
	PermissionsBoundary string                      `json:"permissions_boundary"`
	Tags                map[string]string           `json:"tags"`
	TagsAll             map[string]string           `json:"tags_all"`
	UniqueId            string                      `json:"unique_id"`
	InlinePolicy        []iamrole.InlinePolicyState `json:"inline_policy"`
}