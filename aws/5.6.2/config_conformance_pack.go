// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	configconformancepack "github.com/golingon/terraproviders/aws/5.6.2/configconformancepack"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewConfigConformancePack creates a new instance of [ConfigConformancePack].
func NewConfigConformancePack(name string, args ConfigConformancePackArgs) *ConfigConformancePack {
	return &ConfigConformancePack{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ConfigConformancePack)(nil)

// ConfigConformancePack represents the Terraform resource aws_config_conformance_pack.
type ConfigConformancePack struct {
	Name      string
	Args      ConfigConformancePackArgs
	state     *configConformancePackState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ConfigConformancePack].
func (ccp *ConfigConformancePack) Type() string {
	return "aws_config_conformance_pack"
}

// LocalName returns the local name for [ConfigConformancePack].
func (ccp *ConfigConformancePack) LocalName() string {
	return ccp.Name
}

// Configuration returns the configuration (args) for [ConfigConformancePack].
func (ccp *ConfigConformancePack) Configuration() interface{} {
	return ccp.Args
}

// DependOn is used for other resources to depend on [ConfigConformancePack].
func (ccp *ConfigConformancePack) DependOn() terra.Reference {
	return terra.ReferenceResource(ccp)
}

// Dependencies returns the list of resources [ConfigConformancePack] depends_on.
func (ccp *ConfigConformancePack) Dependencies() terra.Dependencies {
	return ccp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ConfigConformancePack].
func (ccp *ConfigConformancePack) LifecycleManagement() *terra.Lifecycle {
	return ccp.Lifecycle
}

// Attributes returns the attributes for [ConfigConformancePack].
func (ccp *ConfigConformancePack) Attributes() configConformancePackAttributes {
	return configConformancePackAttributes{ref: terra.ReferenceResource(ccp)}
}

// ImportState imports the given attribute values into [ConfigConformancePack]'s state.
func (ccp *ConfigConformancePack) ImportState(av io.Reader) error {
	ccp.state = &configConformancePackState{}
	if err := json.NewDecoder(av).Decode(ccp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ccp.Type(), ccp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ConfigConformancePack] has state.
func (ccp *ConfigConformancePack) State() (*configConformancePackState, bool) {
	return ccp.state, ccp.state != nil
}

// StateMust returns the state for [ConfigConformancePack]. Panics if the state is nil.
func (ccp *ConfigConformancePack) StateMust() *configConformancePackState {
	if ccp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ccp.Type(), ccp.LocalName()))
	}
	return ccp.state
}

// ConfigConformancePackArgs contains the configurations for aws_config_conformance_pack.
type ConfigConformancePackArgs struct {
	// DeliveryS3Bucket: string, optional
	DeliveryS3Bucket terra.StringValue `hcl:"delivery_s3_bucket,attr"`
	// DeliveryS3KeyPrefix: string, optional
	DeliveryS3KeyPrefix terra.StringValue `hcl:"delivery_s3_key_prefix,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// TemplateBody: string, optional
	TemplateBody terra.StringValue `hcl:"template_body,attr"`
	// TemplateS3Uri: string, optional
	TemplateS3Uri terra.StringValue `hcl:"template_s3_uri,attr"`
	// InputParameter: min=0,max=60
	InputParameter []configconformancepack.InputParameter `hcl:"input_parameter,block" validate:"min=0,max=60"`
}
type configConformancePackAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_config_conformance_pack.
func (ccp configConformancePackAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ccp.ref.Append("arn"))
}

// DeliveryS3Bucket returns a reference to field delivery_s3_bucket of aws_config_conformance_pack.
func (ccp configConformancePackAttributes) DeliveryS3Bucket() terra.StringValue {
	return terra.ReferenceAsString(ccp.ref.Append("delivery_s3_bucket"))
}

// DeliveryS3KeyPrefix returns a reference to field delivery_s3_key_prefix of aws_config_conformance_pack.
func (ccp configConformancePackAttributes) DeliveryS3KeyPrefix() terra.StringValue {
	return terra.ReferenceAsString(ccp.ref.Append("delivery_s3_key_prefix"))
}

// Id returns a reference to field id of aws_config_conformance_pack.
func (ccp configConformancePackAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ccp.ref.Append("id"))
}

// Name returns a reference to field name of aws_config_conformance_pack.
func (ccp configConformancePackAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ccp.ref.Append("name"))
}

// TemplateBody returns a reference to field template_body of aws_config_conformance_pack.
func (ccp configConformancePackAttributes) TemplateBody() terra.StringValue {
	return terra.ReferenceAsString(ccp.ref.Append("template_body"))
}

// TemplateS3Uri returns a reference to field template_s3_uri of aws_config_conformance_pack.
func (ccp configConformancePackAttributes) TemplateS3Uri() terra.StringValue {
	return terra.ReferenceAsString(ccp.ref.Append("template_s3_uri"))
}

func (ccp configConformancePackAttributes) InputParameter() terra.SetValue[configconformancepack.InputParameterAttributes] {
	return terra.ReferenceAsSet[configconformancepack.InputParameterAttributes](ccp.ref.Append("input_parameter"))
}

type configConformancePackState struct {
	Arn                 string                                      `json:"arn"`
	DeliveryS3Bucket    string                                      `json:"delivery_s3_bucket"`
	DeliveryS3KeyPrefix string                                      `json:"delivery_s3_key_prefix"`
	Id                  string                                      `json:"id"`
	Name                string                                      `json:"name"`
	TemplateBody        string                                      `json:"template_body"`
	TemplateS3Uri       string                                      `json:"template_s3_uri"`
	InputParameter      []configconformancepack.InputParameterState `json:"input_parameter"`
}
