// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewIamInstanceProfile creates a new instance of [IamInstanceProfile].
func NewIamInstanceProfile(name string, args IamInstanceProfileArgs) *IamInstanceProfile {
	return &IamInstanceProfile{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*IamInstanceProfile)(nil)

// IamInstanceProfile represents the Terraform resource aws_iam_instance_profile.
type IamInstanceProfile struct {
	Name      string
	Args      IamInstanceProfileArgs
	state     *iamInstanceProfileState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [IamInstanceProfile].
func (iip *IamInstanceProfile) Type() string {
	return "aws_iam_instance_profile"
}

// LocalName returns the local name for [IamInstanceProfile].
func (iip *IamInstanceProfile) LocalName() string {
	return iip.Name
}

// Configuration returns the configuration (args) for [IamInstanceProfile].
func (iip *IamInstanceProfile) Configuration() interface{} {
	return iip.Args
}

// DependOn is used for other resources to depend on [IamInstanceProfile].
func (iip *IamInstanceProfile) DependOn() terra.Reference {
	return terra.ReferenceResource(iip)
}

// Dependencies returns the list of resources [IamInstanceProfile] depends_on.
func (iip *IamInstanceProfile) Dependencies() terra.Dependencies {
	return iip.DependsOn
}

// LifecycleManagement returns the lifecycle block for [IamInstanceProfile].
func (iip *IamInstanceProfile) LifecycleManagement() *terra.Lifecycle {
	return iip.Lifecycle
}

// Attributes returns the attributes for [IamInstanceProfile].
func (iip *IamInstanceProfile) Attributes() iamInstanceProfileAttributes {
	return iamInstanceProfileAttributes{ref: terra.ReferenceResource(iip)}
}

// ImportState imports the given attribute values into [IamInstanceProfile]'s state.
func (iip *IamInstanceProfile) ImportState(av io.Reader) error {
	iip.state = &iamInstanceProfileState{}
	if err := json.NewDecoder(av).Decode(iip.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", iip.Type(), iip.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [IamInstanceProfile] has state.
func (iip *IamInstanceProfile) State() (*iamInstanceProfileState, bool) {
	return iip.state, iip.state != nil
}

// StateMust returns the state for [IamInstanceProfile]. Panics if the state is nil.
func (iip *IamInstanceProfile) StateMust() *iamInstanceProfileState {
	if iip.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", iip.Type(), iip.LocalName()))
	}
	return iip.state
}

// IamInstanceProfileArgs contains the configurations for aws_iam_instance_profile.
type IamInstanceProfileArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// NamePrefix: string, optional
	NamePrefix terra.StringValue `hcl:"name_prefix,attr"`
	// Path: string, optional
	Path terra.StringValue `hcl:"path,attr"`
	// Role: string, optional
	Role terra.StringValue `hcl:"role,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
}
type iamInstanceProfileAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_iam_instance_profile.
func (iip iamInstanceProfileAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(iip.ref.Append("arn"))
}

// CreateDate returns a reference to field create_date of aws_iam_instance_profile.
func (iip iamInstanceProfileAttributes) CreateDate() terra.StringValue {
	return terra.ReferenceAsString(iip.ref.Append("create_date"))
}

// Id returns a reference to field id of aws_iam_instance_profile.
func (iip iamInstanceProfileAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(iip.ref.Append("id"))
}

// Name returns a reference to field name of aws_iam_instance_profile.
func (iip iamInstanceProfileAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(iip.ref.Append("name"))
}

// NamePrefix returns a reference to field name_prefix of aws_iam_instance_profile.
func (iip iamInstanceProfileAttributes) NamePrefix() terra.StringValue {
	return terra.ReferenceAsString(iip.ref.Append("name_prefix"))
}

// Path returns a reference to field path of aws_iam_instance_profile.
func (iip iamInstanceProfileAttributes) Path() terra.StringValue {
	return terra.ReferenceAsString(iip.ref.Append("path"))
}

// Role returns a reference to field role of aws_iam_instance_profile.
func (iip iamInstanceProfileAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(iip.ref.Append("role"))
}

// Tags returns a reference to field tags of aws_iam_instance_profile.
func (iip iamInstanceProfileAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](iip.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_iam_instance_profile.
func (iip iamInstanceProfileAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](iip.ref.Append("tags_all"))
}

// UniqueId returns a reference to field unique_id of aws_iam_instance_profile.
func (iip iamInstanceProfileAttributes) UniqueId() terra.StringValue {
	return terra.ReferenceAsString(iip.ref.Append("unique_id"))
}

type iamInstanceProfileState struct {
	Arn        string            `json:"arn"`
	CreateDate string            `json:"create_date"`
	Id         string            `json:"id"`
	Name       string            `json:"name"`
	NamePrefix string            `json:"name_prefix"`
	Path       string            `json:"path"`
	Role       string            `json:"role"`
	Tags       map[string]string `json:"tags"`
	TagsAll    map[string]string `json:"tags_all"`
	UniqueId   string            `json:"unique_id"`
}