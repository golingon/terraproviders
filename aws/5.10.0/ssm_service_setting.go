// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSsmServiceSetting creates a new instance of [SsmServiceSetting].
func NewSsmServiceSetting(name string, args SsmServiceSettingArgs) *SsmServiceSetting {
	return &SsmServiceSetting{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SsmServiceSetting)(nil)

// SsmServiceSetting represents the Terraform resource aws_ssm_service_setting.
type SsmServiceSetting struct {
	Name      string
	Args      SsmServiceSettingArgs
	state     *ssmServiceSettingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SsmServiceSetting].
func (sss *SsmServiceSetting) Type() string {
	return "aws_ssm_service_setting"
}

// LocalName returns the local name for [SsmServiceSetting].
func (sss *SsmServiceSetting) LocalName() string {
	return sss.Name
}

// Configuration returns the configuration (args) for [SsmServiceSetting].
func (sss *SsmServiceSetting) Configuration() interface{} {
	return sss.Args
}

// DependOn is used for other resources to depend on [SsmServiceSetting].
func (sss *SsmServiceSetting) DependOn() terra.Reference {
	return terra.ReferenceResource(sss)
}

// Dependencies returns the list of resources [SsmServiceSetting] depends_on.
func (sss *SsmServiceSetting) Dependencies() terra.Dependencies {
	return sss.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SsmServiceSetting].
func (sss *SsmServiceSetting) LifecycleManagement() *terra.Lifecycle {
	return sss.Lifecycle
}

// Attributes returns the attributes for [SsmServiceSetting].
func (sss *SsmServiceSetting) Attributes() ssmServiceSettingAttributes {
	return ssmServiceSettingAttributes{ref: terra.ReferenceResource(sss)}
}

// ImportState imports the given attribute values into [SsmServiceSetting]'s state.
func (sss *SsmServiceSetting) ImportState(av io.Reader) error {
	sss.state = &ssmServiceSettingState{}
	if err := json.NewDecoder(av).Decode(sss.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sss.Type(), sss.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SsmServiceSetting] has state.
func (sss *SsmServiceSetting) State() (*ssmServiceSettingState, bool) {
	return sss.state, sss.state != nil
}

// StateMust returns the state for [SsmServiceSetting]. Panics if the state is nil.
func (sss *SsmServiceSetting) StateMust() *ssmServiceSettingState {
	if sss.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sss.Type(), sss.LocalName()))
	}
	return sss.state
}

// SsmServiceSettingArgs contains the configurations for aws_ssm_service_setting.
type SsmServiceSettingArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// SettingId: string, required
	SettingId terra.StringValue `hcl:"setting_id,attr" validate:"required"`
	// SettingValue: string, required
	SettingValue terra.StringValue `hcl:"setting_value,attr" validate:"required"`
}
type ssmServiceSettingAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_ssm_service_setting.
func (sss ssmServiceSettingAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(sss.ref.Append("arn"))
}

// Id returns a reference to field id of aws_ssm_service_setting.
func (sss ssmServiceSettingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sss.ref.Append("id"))
}

// SettingId returns a reference to field setting_id of aws_ssm_service_setting.
func (sss ssmServiceSettingAttributes) SettingId() terra.StringValue {
	return terra.ReferenceAsString(sss.ref.Append("setting_id"))
}

// SettingValue returns a reference to field setting_value of aws_ssm_service_setting.
func (sss ssmServiceSettingAttributes) SettingValue() terra.StringValue {
	return terra.ReferenceAsString(sss.ref.Append("setting_value"))
}

// Status returns a reference to field status of aws_ssm_service_setting.
func (sss ssmServiceSettingAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(sss.ref.Append("status"))
}

type ssmServiceSettingState struct {
	Arn          string `json:"arn"`
	Id           string `json:"id"`
	SettingId    string `json:"setting_id"`
	SettingValue string `json:"setting_value"`
	Status       string `json:"status"`
}
