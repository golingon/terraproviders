// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_iam_virtual_mfa_device

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// New creates a new instance of [Resource].
func New(name string, args Args) *Resource {
	return &Resource{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Resource)(nil)

// Resource represents the Terraform resource aws_iam_virtual_mfa_device.
type Resource struct {
	Name      string
	Args      Args
	state     *awsIamVirtualMfaDeviceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (aivmd *Resource) Type() string {
	return "aws_iam_virtual_mfa_device"
}

// LocalName returns the local name for [Resource].
func (aivmd *Resource) LocalName() string {
	return aivmd.Name
}

// Configuration returns the configuration (args) for [Resource].
func (aivmd *Resource) Configuration() interface{} {
	return aivmd.Args
}

// DependOn is used for other resources to depend on [Resource].
func (aivmd *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(aivmd)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (aivmd *Resource) Dependencies() terra.Dependencies {
	return aivmd.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (aivmd *Resource) LifecycleManagement() *terra.Lifecycle {
	return aivmd.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (aivmd *Resource) Attributes() awsIamVirtualMfaDeviceAttributes {
	return awsIamVirtualMfaDeviceAttributes{ref: terra.ReferenceResource(aivmd)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (aivmd *Resource) ImportState(state io.Reader) error {
	aivmd.state = &awsIamVirtualMfaDeviceState{}
	if err := json.NewDecoder(state).Decode(aivmd.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", aivmd.Type(), aivmd.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (aivmd *Resource) State() (*awsIamVirtualMfaDeviceState, bool) {
	return aivmd.state, aivmd.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (aivmd *Resource) StateMust() *awsIamVirtualMfaDeviceState {
	if aivmd.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", aivmd.Type(), aivmd.LocalName()))
	}
	return aivmd.state
}

// Args contains the configurations for aws_iam_virtual_mfa_device.
type Args struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Path: string, optional
	Path terra.StringValue `hcl:"path,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// VirtualMfaDeviceName: string, required
	VirtualMfaDeviceName terra.StringValue `hcl:"virtual_mfa_device_name,attr" validate:"required"`
}

type awsIamVirtualMfaDeviceAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_iam_virtual_mfa_device.
func (aivmd awsIamVirtualMfaDeviceAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(aivmd.ref.Append("arn"))
}

// Base32StringSeed returns a reference to field base_32_string_seed of aws_iam_virtual_mfa_device.
func (aivmd awsIamVirtualMfaDeviceAttributes) Base32StringSeed() terra.StringValue {
	return terra.ReferenceAsString(aivmd.ref.Append("base_32_string_seed"))
}

// EnableDate returns a reference to field enable_date of aws_iam_virtual_mfa_device.
func (aivmd awsIamVirtualMfaDeviceAttributes) EnableDate() terra.StringValue {
	return terra.ReferenceAsString(aivmd.ref.Append("enable_date"))
}

// Id returns a reference to field id of aws_iam_virtual_mfa_device.
func (aivmd awsIamVirtualMfaDeviceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aivmd.ref.Append("id"))
}

// Path returns a reference to field path of aws_iam_virtual_mfa_device.
func (aivmd awsIamVirtualMfaDeviceAttributes) Path() terra.StringValue {
	return terra.ReferenceAsString(aivmd.ref.Append("path"))
}

// QrCodePng returns a reference to field qr_code_png of aws_iam_virtual_mfa_device.
func (aivmd awsIamVirtualMfaDeviceAttributes) QrCodePng() terra.StringValue {
	return terra.ReferenceAsString(aivmd.ref.Append("qr_code_png"))
}

// Tags returns a reference to field tags of aws_iam_virtual_mfa_device.
func (aivmd awsIamVirtualMfaDeviceAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](aivmd.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_iam_virtual_mfa_device.
func (aivmd awsIamVirtualMfaDeviceAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](aivmd.ref.Append("tags_all"))
}

// UserName returns a reference to field user_name of aws_iam_virtual_mfa_device.
func (aivmd awsIamVirtualMfaDeviceAttributes) UserName() terra.StringValue {
	return terra.ReferenceAsString(aivmd.ref.Append("user_name"))
}

// VirtualMfaDeviceName returns a reference to field virtual_mfa_device_name of aws_iam_virtual_mfa_device.
func (aivmd awsIamVirtualMfaDeviceAttributes) VirtualMfaDeviceName() terra.StringValue {
	return terra.ReferenceAsString(aivmd.ref.Append("virtual_mfa_device_name"))
}

type awsIamVirtualMfaDeviceState struct {
	Arn                  string            `json:"arn"`
	Base32StringSeed     string            `json:"base_32_string_seed"`
	EnableDate           string            `json:"enable_date"`
	Id                   string            `json:"id"`
	Path                 string            `json:"path"`
	QrCodePng            string            `json:"qr_code_png"`
	Tags                 map[string]string `json:"tags"`
	TagsAll              map[string]string `json:"tags_all"`
	UserName             string            `json:"user_name"`
	VirtualMfaDeviceName string            `json:"virtual_mfa_device_name"`
}
