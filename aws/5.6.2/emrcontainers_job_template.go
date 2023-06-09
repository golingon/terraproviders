// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	emrcontainersjobtemplate "github.com/golingon/terraproviders/aws/5.6.2/emrcontainersjobtemplate"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewEmrcontainersJobTemplate creates a new instance of [EmrcontainersJobTemplate].
func NewEmrcontainersJobTemplate(name string, args EmrcontainersJobTemplateArgs) *EmrcontainersJobTemplate {
	return &EmrcontainersJobTemplate{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*EmrcontainersJobTemplate)(nil)

// EmrcontainersJobTemplate represents the Terraform resource aws_emrcontainers_job_template.
type EmrcontainersJobTemplate struct {
	Name      string
	Args      EmrcontainersJobTemplateArgs
	state     *emrcontainersJobTemplateState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [EmrcontainersJobTemplate].
func (ejt *EmrcontainersJobTemplate) Type() string {
	return "aws_emrcontainers_job_template"
}

// LocalName returns the local name for [EmrcontainersJobTemplate].
func (ejt *EmrcontainersJobTemplate) LocalName() string {
	return ejt.Name
}

// Configuration returns the configuration (args) for [EmrcontainersJobTemplate].
func (ejt *EmrcontainersJobTemplate) Configuration() interface{} {
	return ejt.Args
}

// DependOn is used for other resources to depend on [EmrcontainersJobTemplate].
func (ejt *EmrcontainersJobTemplate) DependOn() terra.Reference {
	return terra.ReferenceResource(ejt)
}

// Dependencies returns the list of resources [EmrcontainersJobTemplate] depends_on.
func (ejt *EmrcontainersJobTemplate) Dependencies() terra.Dependencies {
	return ejt.DependsOn
}

// LifecycleManagement returns the lifecycle block for [EmrcontainersJobTemplate].
func (ejt *EmrcontainersJobTemplate) LifecycleManagement() *terra.Lifecycle {
	return ejt.Lifecycle
}

// Attributes returns the attributes for [EmrcontainersJobTemplate].
func (ejt *EmrcontainersJobTemplate) Attributes() emrcontainersJobTemplateAttributes {
	return emrcontainersJobTemplateAttributes{ref: terra.ReferenceResource(ejt)}
}

// ImportState imports the given attribute values into [EmrcontainersJobTemplate]'s state.
func (ejt *EmrcontainersJobTemplate) ImportState(av io.Reader) error {
	ejt.state = &emrcontainersJobTemplateState{}
	if err := json.NewDecoder(av).Decode(ejt.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ejt.Type(), ejt.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [EmrcontainersJobTemplate] has state.
func (ejt *EmrcontainersJobTemplate) State() (*emrcontainersJobTemplateState, bool) {
	return ejt.state, ejt.state != nil
}

// StateMust returns the state for [EmrcontainersJobTemplate]. Panics if the state is nil.
func (ejt *EmrcontainersJobTemplate) StateMust() *emrcontainersJobTemplateState {
	if ejt.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ejt.Type(), ejt.LocalName()))
	}
	return ejt.state
}

// EmrcontainersJobTemplateArgs contains the configurations for aws_emrcontainers_job_template.
type EmrcontainersJobTemplateArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// KmsKeyArn: string, optional
	KmsKeyArn terra.StringValue `hcl:"kms_key_arn,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// JobTemplateData: required
	JobTemplateData *emrcontainersjobtemplate.JobTemplateData `hcl:"job_template_data,block" validate:"required"`
	// Timeouts: optional
	Timeouts *emrcontainersjobtemplate.Timeouts `hcl:"timeouts,block"`
}
type emrcontainersJobTemplateAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_emrcontainers_job_template.
func (ejt emrcontainersJobTemplateAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ejt.ref.Append("arn"))
}

// Id returns a reference to field id of aws_emrcontainers_job_template.
func (ejt emrcontainersJobTemplateAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ejt.ref.Append("id"))
}

// KmsKeyArn returns a reference to field kms_key_arn of aws_emrcontainers_job_template.
func (ejt emrcontainersJobTemplateAttributes) KmsKeyArn() terra.StringValue {
	return terra.ReferenceAsString(ejt.ref.Append("kms_key_arn"))
}

// Name returns a reference to field name of aws_emrcontainers_job_template.
func (ejt emrcontainersJobTemplateAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ejt.ref.Append("name"))
}

// Tags returns a reference to field tags of aws_emrcontainers_job_template.
func (ejt emrcontainersJobTemplateAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ejt.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_emrcontainers_job_template.
func (ejt emrcontainersJobTemplateAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ejt.ref.Append("tags_all"))
}

func (ejt emrcontainersJobTemplateAttributes) JobTemplateData() terra.ListValue[emrcontainersjobtemplate.JobTemplateDataAttributes] {
	return terra.ReferenceAsList[emrcontainersjobtemplate.JobTemplateDataAttributes](ejt.ref.Append("job_template_data"))
}

func (ejt emrcontainersJobTemplateAttributes) Timeouts() emrcontainersjobtemplate.TimeoutsAttributes {
	return terra.ReferenceAsSingle[emrcontainersjobtemplate.TimeoutsAttributes](ejt.ref.Append("timeouts"))
}

type emrcontainersJobTemplateState struct {
	Arn             string                                          `json:"arn"`
	Id              string                                          `json:"id"`
	KmsKeyArn       string                                          `json:"kms_key_arn"`
	Name            string                                          `json:"name"`
	Tags            map[string]string                               `json:"tags"`
	TagsAll         map[string]string                               `json:"tags_all"`
	JobTemplateData []emrcontainersjobtemplate.JobTemplateDataState `json:"job_template_data"`
	Timeouts        *emrcontainersjobtemplate.TimeoutsState         `json:"timeouts"`
}
