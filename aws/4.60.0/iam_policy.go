// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewIamPolicy creates a new instance of [IamPolicy].
func NewIamPolicy(name string, args IamPolicyArgs) *IamPolicy {
	return &IamPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*IamPolicy)(nil)

// IamPolicy represents the Terraform resource aws_iam_policy.
type IamPolicy struct {
	Name      string
	Args      IamPolicyArgs
	state     *iamPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [IamPolicy].
func (ip *IamPolicy) Type() string {
	return "aws_iam_policy"
}

// LocalName returns the local name for [IamPolicy].
func (ip *IamPolicy) LocalName() string {
	return ip.Name
}

// Configuration returns the configuration (args) for [IamPolicy].
func (ip *IamPolicy) Configuration() interface{} {
	return ip.Args
}

// DependOn is used for other resources to depend on [IamPolicy].
func (ip *IamPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(ip)
}

// Dependencies returns the list of resources [IamPolicy] depends_on.
func (ip *IamPolicy) Dependencies() terra.Dependencies {
	return ip.DependsOn
}

// LifecycleManagement returns the lifecycle block for [IamPolicy].
func (ip *IamPolicy) LifecycleManagement() *terra.Lifecycle {
	return ip.Lifecycle
}

// Attributes returns the attributes for [IamPolicy].
func (ip *IamPolicy) Attributes() iamPolicyAttributes {
	return iamPolicyAttributes{ref: terra.ReferenceResource(ip)}
}

// ImportState imports the given attribute values into [IamPolicy]'s state.
func (ip *IamPolicy) ImportState(av io.Reader) error {
	ip.state = &iamPolicyState{}
	if err := json.NewDecoder(av).Decode(ip.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ip.Type(), ip.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [IamPolicy] has state.
func (ip *IamPolicy) State() (*iamPolicyState, bool) {
	return ip.state, ip.state != nil
}

// StateMust returns the state for [IamPolicy]. Panics if the state is nil.
func (ip *IamPolicy) StateMust() *iamPolicyState {
	if ip.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ip.Type(), ip.LocalName()))
	}
	return ip.state
}

// IamPolicyArgs contains the configurations for aws_iam_policy.
type IamPolicyArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// NamePrefix: string, optional
	NamePrefix terra.StringValue `hcl:"name_prefix,attr"`
	// Path: string, optional
	Path terra.StringValue `hcl:"path,attr"`
	// Policy: string, required
	Policy terra.StringValue `hcl:"policy,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
}
type iamPolicyAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_iam_policy.
func (ip iamPolicyAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ip.ref.Append("arn"))
}

// Description returns a reference to field description of aws_iam_policy.
func (ip iamPolicyAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(ip.ref.Append("description"))
}

// Id returns a reference to field id of aws_iam_policy.
func (ip iamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ip.ref.Append("id"))
}

// Name returns a reference to field name of aws_iam_policy.
func (ip iamPolicyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ip.ref.Append("name"))
}

// NamePrefix returns a reference to field name_prefix of aws_iam_policy.
func (ip iamPolicyAttributes) NamePrefix() terra.StringValue {
	return terra.ReferenceAsString(ip.ref.Append("name_prefix"))
}

// Path returns a reference to field path of aws_iam_policy.
func (ip iamPolicyAttributes) Path() terra.StringValue {
	return terra.ReferenceAsString(ip.ref.Append("path"))
}

// Policy returns a reference to field policy of aws_iam_policy.
func (ip iamPolicyAttributes) Policy() terra.StringValue {
	return terra.ReferenceAsString(ip.ref.Append("policy"))
}

// PolicyId returns a reference to field policy_id of aws_iam_policy.
func (ip iamPolicyAttributes) PolicyId() terra.StringValue {
	return terra.ReferenceAsString(ip.ref.Append("policy_id"))
}

// Tags returns a reference to field tags of aws_iam_policy.
func (ip iamPolicyAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ip.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_iam_policy.
func (ip iamPolicyAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ip.ref.Append("tags_all"))
}

type iamPolicyState struct {
	Arn         string            `json:"arn"`
	Description string            `json:"description"`
	Id          string            `json:"id"`
	Name        string            `json:"name"`
	NamePrefix  string            `json:"name_prefix"`
	Path        string            `json:"path"`
	Policy      string            `json:"policy"`
	PolicyId    string            `json:"policy_id"`
	Tags        map[string]string `json:"tags"`
	TagsAll     map[string]string `json:"tags_all"`
}
