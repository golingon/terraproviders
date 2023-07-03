// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	sagemakercoderepository "github.com/golingon/terraproviders/aws/5.6.2/sagemakercoderepository"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSagemakerCodeRepository creates a new instance of [SagemakerCodeRepository].
func NewSagemakerCodeRepository(name string, args SagemakerCodeRepositoryArgs) *SagemakerCodeRepository {
	return &SagemakerCodeRepository{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SagemakerCodeRepository)(nil)

// SagemakerCodeRepository represents the Terraform resource aws_sagemaker_code_repository.
type SagemakerCodeRepository struct {
	Name      string
	Args      SagemakerCodeRepositoryArgs
	state     *sagemakerCodeRepositoryState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SagemakerCodeRepository].
func (scr *SagemakerCodeRepository) Type() string {
	return "aws_sagemaker_code_repository"
}

// LocalName returns the local name for [SagemakerCodeRepository].
func (scr *SagemakerCodeRepository) LocalName() string {
	return scr.Name
}

// Configuration returns the configuration (args) for [SagemakerCodeRepository].
func (scr *SagemakerCodeRepository) Configuration() interface{} {
	return scr.Args
}

// DependOn is used for other resources to depend on [SagemakerCodeRepository].
func (scr *SagemakerCodeRepository) DependOn() terra.Reference {
	return terra.ReferenceResource(scr)
}

// Dependencies returns the list of resources [SagemakerCodeRepository] depends_on.
func (scr *SagemakerCodeRepository) Dependencies() terra.Dependencies {
	return scr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SagemakerCodeRepository].
func (scr *SagemakerCodeRepository) LifecycleManagement() *terra.Lifecycle {
	return scr.Lifecycle
}

// Attributes returns the attributes for [SagemakerCodeRepository].
func (scr *SagemakerCodeRepository) Attributes() sagemakerCodeRepositoryAttributes {
	return sagemakerCodeRepositoryAttributes{ref: terra.ReferenceResource(scr)}
}

// ImportState imports the given attribute values into [SagemakerCodeRepository]'s state.
func (scr *SagemakerCodeRepository) ImportState(av io.Reader) error {
	scr.state = &sagemakerCodeRepositoryState{}
	if err := json.NewDecoder(av).Decode(scr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", scr.Type(), scr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SagemakerCodeRepository] has state.
func (scr *SagemakerCodeRepository) State() (*sagemakerCodeRepositoryState, bool) {
	return scr.state, scr.state != nil
}

// StateMust returns the state for [SagemakerCodeRepository]. Panics if the state is nil.
func (scr *SagemakerCodeRepository) StateMust() *sagemakerCodeRepositoryState {
	if scr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", scr.Type(), scr.LocalName()))
	}
	return scr.state
}

// SagemakerCodeRepositoryArgs contains the configurations for aws_sagemaker_code_repository.
type SagemakerCodeRepositoryArgs struct {
	// CodeRepositoryName: string, required
	CodeRepositoryName terra.StringValue `hcl:"code_repository_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// GitConfig: required
	GitConfig *sagemakercoderepository.GitConfig `hcl:"git_config,block" validate:"required"`
}
type sagemakerCodeRepositoryAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_sagemaker_code_repository.
func (scr sagemakerCodeRepositoryAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(scr.ref.Append("arn"))
}

// CodeRepositoryName returns a reference to field code_repository_name of aws_sagemaker_code_repository.
func (scr sagemakerCodeRepositoryAttributes) CodeRepositoryName() terra.StringValue {
	return terra.ReferenceAsString(scr.ref.Append("code_repository_name"))
}

// Id returns a reference to field id of aws_sagemaker_code_repository.
func (scr sagemakerCodeRepositoryAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(scr.ref.Append("id"))
}

// Tags returns a reference to field tags of aws_sagemaker_code_repository.
func (scr sagemakerCodeRepositoryAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](scr.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_sagemaker_code_repository.
func (scr sagemakerCodeRepositoryAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](scr.ref.Append("tags_all"))
}

func (scr sagemakerCodeRepositoryAttributes) GitConfig() terra.ListValue[sagemakercoderepository.GitConfigAttributes] {
	return terra.ReferenceAsList[sagemakercoderepository.GitConfigAttributes](scr.ref.Append("git_config"))
}

type sagemakerCodeRepositoryState struct {
	Arn                string                                   `json:"arn"`
	CodeRepositoryName string                                   `json:"code_repository_name"`
	Id                 string                                   `json:"id"`
	Tags               map[string]string                        `json:"tags"`
	TagsAll            map[string]string                        `json:"tags_all"`
	GitConfig          []sagemakercoderepository.GitConfigState `json:"git_config"`
}
