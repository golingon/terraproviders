// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	sagemakerappimageconfig "github.com/golingon/terraproviders/aws/4.60.0/sagemakerappimageconfig"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewSagemakerAppImageConfig(name string, args SagemakerAppImageConfigArgs) *SagemakerAppImageConfig {
	return &SagemakerAppImageConfig{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SagemakerAppImageConfig)(nil)

type SagemakerAppImageConfig struct {
	Name  string
	Args  SagemakerAppImageConfigArgs
	state *sagemakerAppImageConfigState
}

func (saic *SagemakerAppImageConfig) Type() string {
	return "aws_sagemaker_app_image_config"
}

func (saic *SagemakerAppImageConfig) LocalName() string {
	return saic.Name
}

func (saic *SagemakerAppImageConfig) Configuration() interface{} {
	return saic.Args
}

func (saic *SagemakerAppImageConfig) Attributes() sagemakerAppImageConfigAttributes {
	return sagemakerAppImageConfigAttributes{ref: terra.ReferenceResource(saic)}
}

func (saic *SagemakerAppImageConfig) ImportState(av io.Reader) error {
	saic.state = &sagemakerAppImageConfigState{}
	if err := json.NewDecoder(av).Decode(saic.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", saic.Type(), saic.LocalName(), err)
	}
	return nil
}

func (saic *SagemakerAppImageConfig) State() (*sagemakerAppImageConfigState, bool) {
	return saic.state, saic.state != nil
}

func (saic *SagemakerAppImageConfig) StateMust() *sagemakerAppImageConfigState {
	if saic.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", saic.Type(), saic.LocalName()))
	}
	return saic.state
}

func (saic *SagemakerAppImageConfig) DependOn() terra.Reference {
	return terra.ReferenceResource(saic)
}

type SagemakerAppImageConfigArgs struct {
	// AppImageConfigName: string, required
	AppImageConfigName terra.StringValue `hcl:"app_image_config_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// KernelGatewayImageConfig: optional
	KernelGatewayImageConfig *sagemakerappimageconfig.KernelGatewayImageConfig `hcl:"kernel_gateway_image_config,block"`
	// DependsOn contains resources that SagemakerAppImageConfig depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type sagemakerAppImageConfigAttributes struct {
	ref terra.Reference
}

func (saic sagemakerAppImageConfigAttributes) AppImageConfigName() terra.StringValue {
	return terra.ReferenceString(saic.ref.Append("app_image_config_name"))
}

func (saic sagemakerAppImageConfigAttributes) Arn() terra.StringValue {
	return terra.ReferenceString(saic.ref.Append("arn"))
}

func (saic sagemakerAppImageConfigAttributes) Id() terra.StringValue {
	return terra.ReferenceString(saic.ref.Append("id"))
}

func (saic sagemakerAppImageConfigAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](saic.ref.Append("tags"))
}

func (saic sagemakerAppImageConfigAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](saic.ref.Append("tags_all"))
}

func (saic sagemakerAppImageConfigAttributes) KernelGatewayImageConfig() terra.ListValue[sagemakerappimageconfig.KernelGatewayImageConfigAttributes] {
	return terra.ReferenceList[sagemakerappimageconfig.KernelGatewayImageConfigAttributes](saic.ref.Append("kernel_gateway_image_config"))
}

type sagemakerAppImageConfigState struct {
	AppImageConfigName       string                                                  `json:"app_image_config_name"`
	Arn                      string                                                  `json:"arn"`
	Id                       string                                                  `json:"id"`
	Tags                     map[string]string                                       `json:"tags"`
	TagsAll                  map[string]string                                       `json:"tags_all"`
	KernelGatewayImageConfig []sagemakerappimageconfig.KernelGatewayImageConfigState `json:"kernel_gateway_image_config"`
}
