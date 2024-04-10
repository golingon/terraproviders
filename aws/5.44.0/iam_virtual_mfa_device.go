// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// NewIamVirtualMfaDevice creates a new instance of [IamVirtualMfaDevice].
func NewIamVirtualMfaDevice(name string, args IamVirtualMfaDeviceArgs) *IamVirtualMfaDevice {
	return &IamVirtualMfaDevice{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*IamVirtualMfaDevice)(nil)

// IamVirtualMfaDevice represents the Terraform resource aws_iam_virtual_mfa_device.
type IamVirtualMfaDevice struct {
	Name      string
	Args      IamVirtualMfaDeviceArgs
	state     *iamVirtualMfaDeviceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [IamVirtualMfaDevice].
func (ivmd *IamVirtualMfaDevice) Type() string {
	return "aws_iam_virtual_mfa_device"
}

// LocalName returns the local name for [IamVirtualMfaDevice].
func (ivmd *IamVirtualMfaDevice) LocalName() string {
	return ivmd.Name
}

// Configuration returns the configuration (args) for [IamVirtualMfaDevice].
func (ivmd *IamVirtualMfaDevice) Configuration() interface{} {
	return ivmd.Args
}

// DependOn is used for other resources to depend on [IamVirtualMfaDevice].
func (ivmd *IamVirtualMfaDevice) DependOn() terra.Reference {
	return terra.ReferenceResource(ivmd)
}

// Dependencies returns the list of resources [IamVirtualMfaDevice] depends_on.
func (ivmd *IamVirtualMfaDevice) Dependencies() terra.Dependencies {
	return ivmd.DependsOn
}

// LifecycleManagement returns the lifecycle block for [IamVirtualMfaDevice].
func (ivmd *IamVirtualMfaDevice) LifecycleManagement() *terra.Lifecycle {
	return ivmd.Lifecycle
}

// Attributes returns the attributes for [IamVirtualMfaDevice].
func (ivmd *IamVirtualMfaDevice) Attributes() iamVirtualMfaDeviceAttributes {
	return iamVirtualMfaDeviceAttributes{ref: terra.ReferenceResource(ivmd)}
}

// ImportState imports the given attribute values into [IamVirtualMfaDevice]'s state.
func (ivmd *IamVirtualMfaDevice) ImportState(av io.Reader) error {
	ivmd.state = &iamVirtualMfaDeviceState{}
	if err := json.NewDecoder(av).Decode(ivmd.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ivmd.Type(), ivmd.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [IamVirtualMfaDevice] has state.
func (ivmd *IamVirtualMfaDevice) State() (*iamVirtualMfaDeviceState, bool) {
	return ivmd.state, ivmd.state != nil
}

// StateMust returns the state for [IamVirtualMfaDevice]. Panics if the state is nil.
func (ivmd *IamVirtualMfaDevice) StateMust() *iamVirtualMfaDeviceState {
	if ivmd.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ivmd.Type(), ivmd.LocalName()))
	}
	return ivmd.state
}

// IamVirtualMfaDeviceArgs contains the configurations for aws_iam_virtual_mfa_device.
type IamVirtualMfaDeviceArgs struct {
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
type iamVirtualMfaDeviceAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_iam_virtual_mfa_device.
func (ivmd iamVirtualMfaDeviceAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ivmd.ref.Append("arn"))
}

// Base32StringSeed returns a reference to field base_32_string_seed of aws_iam_virtual_mfa_device.
func (ivmd iamVirtualMfaDeviceAttributes) Base32StringSeed() terra.StringValue {
	return terra.ReferenceAsString(ivmd.ref.Append("base_32_string_seed"))
}

// EnableDate returns a reference to field enable_date of aws_iam_virtual_mfa_device.
func (ivmd iamVirtualMfaDeviceAttributes) EnableDate() terra.StringValue {
	return terra.ReferenceAsString(ivmd.ref.Append("enable_date"))
}

// Id returns a reference to field id of aws_iam_virtual_mfa_device.
func (ivmd iamVirtualMfaDeviceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ivmd.ref.Append("id"))
}

// Path returns a reference to field path of aws_iam_virtual_mfa_device.
func (ivmd iamVirtualMfaDeviceAttributes) Path() terra.StringValue {
	return terra.ReferenceAsString(ivmd.ref.Append("path"))
}

// QrCodePng returns a reference to field qr_code_png of aws_iam_virtual_mfa_device.
func (ivmd iamVirtualMfaDeviceAttributes) QrCodePng() terra.StringValue {
	return terra.ReferenceAsString(ivmd.ref.Append("qr_code_png"))
}

// Tags returns a reference to field tags of aws_iam_virtual_mfa_device.
func (ivmd iamVirtualMfaDeviceAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ivmd.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_iam_virtual_mfa_device.
func (ivmd iamVirtualMfaDeviceAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ivmd.ref.Append("tags_all"))
}

// UserName returns a reference to field user_name of aws_iam_virtual_mfa_device.
func (ivmd iamVirtualMfaDeviceAttributes) UserName() terra.StringValue {
	return terra.ReferenceAsString(ivmd.ref.Append("user_name"))
}

// VirtualMfaDeviceName returns a reference to field virtual_mfa_device_name of aws_iam_virtual_mfa_device.
func (ivmd iamVirtualMfaDeviceAttributes) VirtualMfaDeviceName() terra.StringValue {
	return terra.ReferenceAsString(ivmd.ref.Append("virtual_mfa_device_name"))
}

type iamVirtualMfaDeviceState struct {
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
