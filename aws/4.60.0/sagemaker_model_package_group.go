// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewSagemakerModelPackageGroup(name string, args SagemakerModelPackageGroupArgs) *SagemakerModelPackageGroup {
	return &SagemakerModelPackageGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SagemakerModelPackageGroup)(nil)

type SagemakerModelPackageGroup struct {
	Name  string
	Args  SagemakerModelPackageGroupArgs
	state *sagemakerModelPackageGroupState
}

func (smpg *SagemakerModelPackageGroup) Type() string {
	return "aws_sagemaker_model_package_group"
}

func (smpg *SagemakerModelPackageGroup) LocalName() string {
	return smpg.Name
}

func (smpg *SagemakerModelPackageGroup) Configuration() interface{} {
	return smpg.Args
}

func (smpg *SagemakerModelPackageGroup) Attributes() sagemakerModelPackageGroupAttributes {
	return sagemakerModelPackageGroupAttributes{ref: terra.ReferenceResource(smpg)}
}

func (smpg *SagemakerModelPackageGroup) ImportState(av io.Reader) error {
	smpg.state = &sagemakerModelPackageGroupState{}
	if err := json.NewDecoder(av).Decode(smpg.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", smpg.Type(), smpg.LocalName(), err)
	}
	return nil
}

func (smpg *SagemakerModelPackageGroup) State() (*sagemakerModelPackageGroupState, bool) {
	return smpg.state, smpg.state != nil
}

func (smpg *SagemakerModelPackageGroup) StateMust() *sagemakerModelPackageGroupState {
	if smpg.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", smpg.Type(), smpg.LocalName()))
	}
	return smpg.state
}

func (smpg *SagemakerModelPackageGroup) DependOn() terra.Reference {
	return terra.ReferenceResource(smpg)
}

type SagemakerModelPackageGroupArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ModelPackageGroupDescription: string, optional
	ModelPackageGroupDescription terra.StringValue `hcl:"model_package_group_description,attr"`
	// ModelPackageGroupName: string, required
	ModelPackageGroupName terra.StringValue `hcl:"model_package_group_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// DependsOn contains resources that SagemakerModelPackageGroup depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type sagemakerModelPackageGroupAttributes struct {
	ref terra.Reference
}

func (smpg sagemakerModelPackageGroupAttributes) Arn() terra.StringValue {
	return terra.ReferenceString(smpg.ref.Append("arn"))
}

func (smpg sagemakerModelPackageGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceString(smpg.ref.Append("id"))
}

func (smpg sagemakerModelPackageGroupAttributes) ModelPackageGroupDescription() terra.StringValue {
	return terra.ReferenceString(smpg.ref.Append("model_package_group_description"))
}

func (smpg sagemakerModelPackageGroupAttributes) ModelPackageGroupName() terra.StringValue {
	return terra.ReferenceString(smpg.ref.Append("model_package_group_name"))
}

func (smpg sagemakerModelPackageGroupAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](smpg.ref.Append("tags"))
}

func (smpg sagemakerModelPackageGroupAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](smpg.ref.Append("tags_all"))
}

type sagemakerModelPackageGroupState struct {
	Arn                          string            `json:"arn"`
	Id                           string            `json:"id"`
	ModelPackageGroupDescription string            `json:"model_package_group_description"`
	ModelPackageGroupName        string            `json:"model_package_group_name"`
	Tags                         map[string]string `json:"tags"`
	TagsAll                      map[string]string `json:"tags_all"`
}
