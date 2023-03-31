// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	connectphonenumber "github.com/golingon/terraproviders/aws/4.60.0/connectphonenumber"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewConnectPhoneNumber creates a new instance of [ConnectPhoneNumber].
func NewConnectPhoneNumber(name string, args ConnectPhoneNumberArgs) *ConnectPhoneNumber {
	return &ConnectPhoneNumber{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ConnectPhoneNumber)(nil)

// ConnectPhoneNumber represents the Terraform resource aws_connect_phone_number.
type ConnectPhoneNumber struct {
	Name      string
	Args      ConnectPhoneNumberArgs
	state     *connectPhoneNumberState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ConnectPhoneNumber].
func (cpn *ConnectPhoneNumber) Type() string {
	return "aws_connect_phone_number"
}

// LocalName returns the local name for [ConnectPhoneNumber].
func (cpn *ConnectPhoneNumber) LocalName() string {
	return cpn.Name
}

// Configuration returns the configuration (args) for [ConnectPhoneNumber].
func (cpn *ConnectPhoneNumber) Configuration() interface{} {
	return cpn.Args
}

// DependOn is used for other resources to depend on [ConnectPhoneNumber].
func (cpn *ConnectPhoneNumber) DependOn() terra.Reference {
	return terra.ReferenceResource(cpn)
}

// Dependencies returns the list of resources [ConnectPhoneNumber] depends_on.
func (cpn *ConnectPhoneNumber) Dependencies() terra.Dependencies {
	return cpn.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ConnectPhoneNumber].
func (cpn *ConnectPhoneNumber) LifecycleManagement() *terra.Lifecycle {
	return cpn.Lifecycle
}

// Attributes returns the attributes for [ConnectPhoneNumber].
func (cpn *ConnectPhoneNumber) Attributes() connectPhoneNumberAttributes {
	return connectPhoneNumberAttributes{ref: terra.ReferenceResource(cpn)}
}

// ImportState imports the given attribute values into [ConnectPhoneNumber]'s state.
func (cpn *ConnectPhoneNumber) ImportState(av io.Reader) error {
	cpn.state = &connectPhoneNumberState{}
	if err := json.NewDecoder(av).Decode(cpn.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cpn.Type(), cpn.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ConnectPhoneNumber] has state.
func (cpn *ConnectPhoneNumber) State() (*connectPhoneNumberState, bool) {
	return cpn.state, cpn.state != nil
}

// StateMust returns the state for [ConnectPhoneNumber]. Panics if the state is nil.
func (cpn *ConnectPhoneNumber) StateMust() *connectPhoneNumberState {
	if cpn.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cpn.Type(), cpn.LocalName()))
	}
	return cpn.state
}

// ConnectPhoneNumberArgs contains the configurations for aws_connect_phone_number.
type ConnectPhoneNumberArgs struct {
	// CountryCode: string, required
	CountryCode terra.StringValue `hcl:"country_code,attr" validate:"required"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Prefix: string, optional
	Prefix terra.StringValue `hcl:"prefix,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// TargetArn: string, required
	TargetArn terra.StringValue `hcl:"target_arn,attr" validate:"required"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
	// Status: min=0
	Status []connectphonenumber.Status `hcl:"status,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *connectphonenumber.Timeouts `hcl:"timeouts,block"`
}
type connectPhoneNumberAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_connect_phone_number.
func (cpn connectPhoneNumberAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(cpn.ref.Append("arn"))
}

// CountryCode returns a reference to field country_code of aws_connect_phone_number.
func (cpn connectPhoneNumberAttributes) CountryCode() terra.StringValue {
	return terra.ReferenceAsString(cpn.ref.Append("country_code"))
}

// Description returns a reference to field description of aws_connect_phone_number.
func (cpn connectPhoneNumberAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(cpn.ref.Append("description"))
}

// Id returns a reference to field id of aws_connect_phone_number.
func (cpn connectPhoneNumberAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cpn.ref.Append("id"))
}

// PhoneNumber returns a reference to field phone_number of aws_connect_phone_number.
func (cpn connectPhoneNumberAttributes) PhoneNumber() terra.StringValue {
	return terra.ReferenceAsString(cpn.ref.Append("phone_number"))
}

// Prefix returns a reference to field prefix of aws_connect_phone_number.
func (cpn connectPhoneNumberAttributes) Prefix() terra.StringValue {
	return terra.ReferenceAsString(cpn.ref.Append("prefix"))
}

// Tags returns a reference to field tags of aws_connect_phone_number.
func (cpn connectPhoneNumberAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cpn.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_connect_phone_number.
func (cpn connectPhoneNumberAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cpn.ref.Append("tags_all"))
}

// TargetArn returns a reference to field target_arn of aws_connect_phone_number.
func (cpn connectPhoneNumberAttributes) TargetArn() terra.StringValue {
	return terra.ReferenceAsString(cpn.ref.Append("target_arn"))
}

// Type returns a reference to field type of aws_connect_phone_number.
func (cpn connectPhoneNumberAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(cpn.ref.Append("type"))
}

func (cpn connectPhoneNumberAttributes) Status() terra.ListValue[connectphonenumber.StatusAttributes] {
	return terra.ReferenceAsList[connectphonenumber.StatusAttributes](cpn.ref.Append("status"))
}

func (cpn connectPhoneNumberAttributes) Timeouts() connectphonenumber.TimeoutsAttributes {
	return terra.ReferenceAsSingle[connectphonenumber.TimeoutsAttributes](cpn.ref.Append("timeouts"))
}

type connectPhoneNumberState struct {
	Arn         string                            `json:"arn"`
	CountryCode string                            `json:"country_code"`
	Description string                            `json:"description"`
	Id          string                            `json:"id"`
	PhoneNumber string                            `json:"phone_number"`
	Prefix      string                            `json:"prefix"`
	Tags        map[string]string                 `json:"tags"`
	TagsAll     map[string]string                 `json:"tags_all"`
	TargetArn   string                            `json:"target_arn"`
	Type        string                            `json:"type"`
	Status      []connectphonenumber.StatusState  `json:"status"`
	Timeouts    *connectphonenumber.TimeoutsState `json:"timeouts"`
}
