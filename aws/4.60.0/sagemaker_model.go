// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	sagemakermodel "github.com/golingon/terraproviders/aws/4.60.0/sagemakermodel"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewSagemakerModel(name string, args SagemakerModelArgs) *SagemakerModel {
	return &SagemakerModel{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SagemakerModel)(nil)

type SagemakerModel struct {
	Name  string
	Args  SagemakerModelArgs
	state *sagemakerModelState
}

func (sm *SagemakerModel) Type() string {
	return "aws_sagemaker_model"
}

func (sm *SagemakerModel) LocalName() string {
	return sm.Name
}

func (sm *SagemakerModel) Configuration() interface{} {
	return sm.Args
}

func (sm *SagemakerModel) Attributes() sagemakerModelAttributes {
	return sagemakerModelAttributes{ref: terra.ReferenceResource(sm)}
}

func (sm *SagemakerModel) ImportState(av io.Reader) error {
	sm.state = &sagemakerModelState{}
	if err := json.NewDecoder(av).Decode(sm.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sm.Type(), sm.LocalName(), err)
	}
	return nil
}

func (sm *SagemakerModel) State() (*sagemakerModelState, bool) {
	return sm.state, sm.state != nil
}

func (sm *SagemakerModel) StateMust() *sagemakerModelState {
	if sm.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sm.Type(), sm.LocalName()))
	}
	return sm.state
}

func (sm *SagemakerModel) DependOn() terra.Reference {
	return terra.ReferenceResource(sm)
}

type SagemakerModelArgs struct {
	// EnableNetworkIsolation: bool, optional
	EnableNetworkIsolation terra.BoolValue `hcl:"enable_network_isolation,attr"`
	// ExecutionRoleArn: string, required
	ExecutionRoleArn terra.StringValue `hcl:"execution_role_arn,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// Container: min=0
	Container []sagemakermodel.Container `hcl:"container,block" validate:"min=0"`
	// InferenceExecutionConfig: optional
	InferenceExecutionConfig *sagemakermodel.InferenceExecutionConfig `hcl:"inference_execution_config,block"`
	// PrimaryContainer: optional
	PrimaryContainer *sagemakermodel.PrimaryContainer `hcl:"primary_container,block"`
	// VpcConfig: optional
	VpcConfig *sagemakermodel.VpcConfig `hcl:"vpc_config,block"`
	// DependsOn contains resources that SagemakerModel depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type sagemakerModelAttributes struct {
	ref terra.Reference
}

func (sm sagemakerModelAttributes) Arn() terra.StringValue {
	return terra.ReferenceString(sm.ref.Append("arn"))
}

func (sm sagemakerModelAttributes) EnableNetworkIsolation() terra.BoolValue {
	return terra.ReferenceBool(sm.ref.Append("enable_network_isolation"))
}

func (sm sagemakerModelAttributes) ExecutionRoleArn() terra.StringValue {
	return terra.ReferenceString(sm.ref.Append("execution_role_arn"))
}

func (sm sagemakerModelAttributes) Id() terra.StringValue {
	return terra.ReferenceString(sm.ref.Append("id"))
}

func (sm sagemakerModelAttributes) Name() terra.StringValue {
	return terra.ReferenceString(sm.ref.Append("name"))
}

func (sm sagemakerModelAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](sm.ref.Append("tags"))
}

func (sm sagemakerModelAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](sm.ref.Append("tags_all"))
}

func (sm sagemakerModelAttributes) Container() terra.ListValue[sagemakermodel.ContainerAttributes] {
	return terra.ReferenceList[sagemakermodel.ContainerAttributes](sm.ref.Append("container"))
}

func (sm sagemakerModelAttributes) InferenceExecutionConfig() terra.ListValue[sagemakermodel.InferenceExecutionConfigAttributes] {
	return terra.ReferenceList[sagemakermodel.InferenceExecutionConfigAttributes](sm.ref.Append("inference_execution_config"))
}

func (sm sagemakerModelAttributes) PrimaryContainer() terra.ListValue[sagemakermodel.PrimaryContainerAttributes] {
	return terra.ReferenceList[sagemakermodel.PrimaryContainerAttributes](sm.ref.Append("primary_container"))
}

func (sm sagemakerModelAttributes) VpcConfig() terra.ListValue[sagemakermodel.VpcConfigAttributes] {
	return terra.ReferenceList[sagemakermodel.VpcConfigAttributes](sm.ref.Append("vpc_config"))
}

type sagemakerModelState struct {
	Arn                      string                                         `json:"arn"`
	EnableNetworkIsolation   bool                                           `json:"enable_network_isolation"`
	ExecutionRoleArn         string                                         `json:"execution_role_arn"`
	Id                       string                                         `json:"id"`
	Name                     string                                         `json:"name"`
	Tags                     map[string]string                              `json:"tags"`
	TagsAll                  map[string]string                              `json:"tags_all"`
	Container                []sagemakermodel.ContainerState                `json:"container"`
	InferenceExecutionConfig []sagemakermodel.InferenceExecutionConfigState `json:"inference_execution_config"`
	PrimaryContainer         []sagemakermodel.PrimaryContainerState         `json:"primary_container"`
	VpcConfig                []sagemakermodel.VpcConfigState                `json:"vpc_config"`
}
