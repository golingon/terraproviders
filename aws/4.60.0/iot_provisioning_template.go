// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	iotprovisioningtemplate "github.com/golingon/terraproviders/aws/4.60.0/iotprovisioningtemplate"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewIotProvisioningTemplate creates a new instance of [IotProvisioningTemplate].
func NewIotProvisioningTemplate(name string, args IotProvisioningTemplateArgs) *IotProvisioningTemplate {
	return &IotProvisioningTemplate{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*IotProvisioningTemplate)(nil)

// IotProvisioningTemplate represents the Terraform resource aws_iot_provisioning_template.
type IotProvisioningTemplate struct {
	Name      string
	Args      IotProvisioningTemplateArgs
	state     *iotProvisioningTemplateState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [IotProvisioningTemplate].
func (ipt *IotProvisioningTemplate) Type() string {
	return "aws_iot_provisioning_template"
}

// LocalName returns the local name for [IotProvisioningTemplate].
func (ipt *IotProvisioningTemplate) LocalName() string {
	return ipt.Name
}

// Configuration returns the configuration (args) for [IotProvisioningTemplate].
func (ipt *IotProvisioningTemplate) Configuration() interface{} {
	return ipt.Args
}

// DependOn is used for other resources to depend on [IotProvisioningTemplate].
func (ipt *IotProvisioningTemplate) DependOn() terra.Reference {
	return terra.ReferenceResource(ipt)
}

// Dependencies returns the list of resources [IotProvisioningTemplate] depends_on.
func (ipt *IotProvisioningTemplate) Dependencies() terra.Dependencies {
	return ipt.DependsOn
}

// LifecycleManagement returns the lifecycle block for [IotProvisioningTemplate].
func (ipt *IotProvisioningTemplate) LifecycleManagement() *terra.Lifecycle {
	return ipt.Lifecycle
}

// Attributes returns the attributes for [IotProvisioningTemplate].
func (ipt *IotProvisioningTemplate) Attributes() iotProvisioningTemplateAttributes {
	return iotProvisioningTemplateAttributes{ref: terra.ReferenceResource(ipt)}
}

// ImportState imports the given attribute values into [IotProvisioningTemplate]'s state.
func (ipt *IotProvisioningTemplate) ImportState(av io.Reader) error {
	ipt.state = &iotProvisioningTemplateState{}
	if err := json.NewDecoder(av).Decode(ipt.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ipt.Type(), ipt.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [IotProvisioningTemplate] has state.
func (ipt *IotProvisioningTemplate) State() (*iotProvisioningTemplateState, bool) {
	return ipt.state, ipt.state != nil
}

// StateMust returns the state for [IotProvisioningTemplate]. Panics if the state is nil.
func (ipt *IotProvisioningTemplate) StateMust() *iotProvisioningTemplateState {
	if ipt.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ipt.Type(), ipt.LocalName()))
	}
	return ipt.state
}

// IotProvisioningTemplateArgs contains the configurations for aws_iot_provisioning_template.
type IotProvisioningTemplateArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ProvisioningRoleArn: string, required
	ProvisioningRoleArn terra.StringValue `hcl:"provisioning_role_arn,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// TemplateBody: string, required
	TemplateBody terra.StringValue `hcl:"template_body,attr" validate:"required"`
	// PreProvisioningHook: optional
	PreProvisioningHook *iotprovisioningtemplate.PreProvisioningHook `hcl:"pre_provisioning_hook,block"`
}
type iotProvisioningTemplateAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_iot_provisioning_template.
func (ipt iotProvisioningTemplateAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ipt.ref.Append("arn"))
}

// DefaultVersionId returns a reference to field default_version_id of aws_iot_provisioning_template.
func (ipt iotProvisioningTemplateAttributes) DefaultVersionId() terra.NumberValue {
	return terra.ReferenceAsNumber(ipt.ref.Append("default_version_id"))
}

// Description returns a reference to field description of aws_iot_provisioning_template.
func (ipt iotProvisioningTemplateAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(ipt.ref.Append("description"))
}

// Enabled returns a reference to field enabled of aws_iot_provisioning_template.
func (ipt iotProvisioningTemplateAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(ipt.ref.Append("enabled"))
}

// Id returns a reference to field id of aws_iot_provisioning_template.
func (ipt iotProvisioningTemplateAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ipt.ref.Append("id"))
}

// Name returns a reference to field name of aws_iot_provisioning_template.
func (ipt iotProvisioningTemplateAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ipt.ref.Append("name"))
}

// ProvisioningRoleArn returns a reference to field provisioning_role_arn of aws_iot_provisioning_template.
func (ipt iotProvisioningTemplateAttributes) ProvisioningRoleArn() terra.StringValue {
	return terra.ReferenceAsString(ipt.ref.Append("provisioning_role_arn"))
}

// Tags returns a reference to field tags of aws_iot_provisioning_template.
func (ipt iotProvisioningTemplateAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ipt.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_iot_provisioning_template.
func (ipt iotProvisioningTemplateAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ipt.ref.Append("tags_all"))
}

// TemplateBody returns a reference to field template_body of aws_iot_provisioning_template.
func (ipt iotProvisioningTemplateAttributes) TemplateBody() terra.StringValue {
	return terra.ReferenceAsString(ipt.ref.Append("template_body"))
}

func (ipt iotProvisioningTemplateAttributes) PreProvisioningHook() terra.ListValue[iotprovisioningtemplate.PreProvisioningHookAttributes] {
	return terra.ReferenceAsList[iotprovisioningtemplate.PreProvisioningHookAttributes](ipt.ref.Append("pre_provisioning_hook"))
}

type iotProvisioningTemplateState struct {
	Arn                 string                                             `json:"arn"`
	DefaultVersionId    float64                                            `json:"default_version_id"`
	Description         string                                             `json:"description"`
	Enabled             bool                                               `json:"enabled"`
	Id                  string                                             `json:"id"`
	Name                string                                             `json:"name"`
	ProvisioningRoleArn string                                             `json:"provisioning_role_arn"`
	Tags                map[string]string                                  `json:"tags"`
	TagsAll             map[string]string                                  `json:"tags_all"`
	TemplateBody        string                                             `json:"template_body"`
	PreProvisioningHook []iotprovisioningtemplate.PreProvisioningHookState `json:"pre_provisioning_hook"`
}
