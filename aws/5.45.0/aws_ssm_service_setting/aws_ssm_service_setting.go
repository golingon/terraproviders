// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_ssm_service_setting

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

// Resource represents the Terraform resource aws_ssm_service_setting.
type Resource struct {
	Name      string
	Args      Args
	state     *awsSsmServiceSettingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (asss *Resource) Type() string {
	return "aws_ssm_service_setting"
}

// LocalName returns the local name for [Resource].
func (asss *Resource) LocalName() string {
	return asss.Name
}

// Configuration returns the configuration (args) for [Resource].
func (asss *Resource) Configuration() interface{} {
	return asss.Args
}

// DependOn is used for other resources to depend on [Resource].
func (asss *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(asss)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (asss *Resource) Dependencies() terra.Dependencies {
	return asss.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (asss *Resource) LifecycleManagement() *terra.Lifecycle {
	return asss.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (asss *Resource) Attributes() awsSsmServiceSettingAttributes {
	return awsSsmServiceSettingAttributes{ref: terra.ReferenceResource(asss)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (asss *Resource) ImportState(state io.Reader) error {
	asss.state = &awsSsmServiceSettingState{}
	if err := json.NewDecoder(state).Decode(asss.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", asss.Type(), asss.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (asss *Resource) State() (*awsSsmServiceSettingState, bool) {
	return asss.state, asss.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (asss *Resource) StateMust() *awsSsmServiceSettingState {
	if asss.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", asss.Type(), asss.LocalName()))
	}
	return asss.state
}

// Args contains the configurations for aws_ssm_service_setting.
type Args struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// SettingId: string, required
	SettingId terra.StringValue `hcl:"setting_id,attr" validate:"required"`
	// SettingValue: string, required
	SettingValue terra.StringValue `hcl:"setting_value,attr" validate:"required"`
}

type awsSsmServiceSettingAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_ssm_service_setting.
func (asss awsSsmServiceSettingAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(asss.ref.Append("arn"))
}

// Id returns a reference to field id of aws_ssm_service_setting.
func (asss awsSsmServiceSettingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(asss.ref.Append("id"))
}

// SettingId returns a reference to field setting_id of aws_ssm_service_setting.
func (asss awsSsmServiceSettingAttributes) SettingId() terra.StringValue {
	return terra.ReferenceAsString(asss.ref.Append("setting_id"))
}

// SettingValue returns a reference to field setting_value of aws_ssm_service_setting.
func (asss awsSsmServiceSettingAttributes) SettingValue() terra.StringValue {
	return terra.ReferenceAsString(asss.ref.Append("setting_value"))
}

// Status returns a reference to field status of aws_ssm_service_setting.
func (asss awsSsmServiceSettingAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(asss.ref.Append("status"))
}

type awsSsmServiceSettingState struct {
	Arn          string `json:"arn"`
	Id           string `json:"id"`
	SettingId    string `json:"setting_id"`
	SettingValue string `json:"setting_value"`
	Status       string `json:"status"`
}
