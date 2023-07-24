// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	kendraexperience "github.com/golingon/terraproviders/aws/5.9.0/kendraexperience"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewKendraExperience creates a new instance of [KendraExperience].
func NewKendraExperience(name string, args KendraExperienceArgs) *KendraExperience {
	return &KendraExperience{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*KendraExperience)(nil)

// KendraExperience represents the Terraform resource aws_kendra_experience.
type KendraExperience struct {
	Name      string
	Args      KendraExperienceArgs
	state     *kendraExperienceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [KendraExperience].
func (ke *KendraExperience) Type() string {
	return "aws_kendra_experience"
}

// LocalName returns the local name for [KendraExperience].
func (ke *KendraExperience) LocalName() string {
	return ke.Name
}

// Configuration returns the configuration (args) for [KendraExperience].
func (ke *KendraExperience) Configuration() interface{} {
	return ke.Args
}

// DependOn is used for other resources to depend on [KendraExperience].
func (ke *KendraExperience) DependOn() terra.Reference {
	return terra.ReferenceResource(ke)
}

// Dependencies returns the list of resources [KendraExperience] depends_on.
func (ke *KendraExperience) Dependencies() terra.Dependencies {
	return ke.DependsOn
}

// LifecycleManagement returns the lifecycle block for [KendraExperience].
func (ke *KendraExperience) LifecycleManagement() *terra.Lifecycle {
	return ke.Lifecycle
}

// Attributes returns the attributes for [KendraExperience].
func (ke *KendraExperience) Attributes() kendraExperienceAttributes {
	return kendraExperienceAttributes{ref: terra.ReferenceResource(ke)}
}

// ImportState imports the given attribute values into [KendraExperience]'s state.
func (ke *KendraExperience) ImportState(av io.Reader) error {
	ke.state = &kendraExperienceState{}
	if err := json.NewDecoder(av).Decode(ke.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ke.Type(), ke.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [KendraExperience] has state.
func (ke *KendraExperience) State() (*kendraExperienceState, bool) {
	return ke.state, ke.state != nil
}

// StateMust returns the state for [KendraExperience]. Panics if the state is nil.
func (ke *KendraExperience) StateMust() *kendraExperienceState {
	if ke.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ke.Type(), ke.LocalName()))
	}
	return ke.state
}

// KendraExperienceArgs contains the configurations for aws_kendra_experience.
type KendraExperienceArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IndexId: string, required
	IndexId terra.StringValue `hcl:"index_id,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// RoleArn: string, required
	RoleArn terra.StringValue `hcl:"role_arn,attr" validate:"required"`
	// Endpoints: min=0
	Endpoints []kendraexperience.Endpoints `hcl:"endpoints,block" validate:"min=0"`
	// Configuration: optional
	Configuration *kendraexperience.Configuration `hcl:"configuration,block"`
	// Timeouts: optional
	Timeouts *kendraexperience.Timeouts `hcl:"timeouts,block"`
}
type kendraExperienceAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_kendra_experience.
func (ke kendraExperienceAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ke.ref.Append("arn"))
}

// Description returns a reference to field description of aws_kendra_experience.
func (ke kendraExperienceAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(ke.ref.Append("description"))
}

// ExperienceId returns a reference to field experience_id of aws_kendra_experience.
func (ke kendraExperienceAttributes) ExperienceId() terra.StringValue {
	return terra.ReferenceAsString(ke.ref.Append("experience_id"))
}

// Id returns a reference to field id of aws_kendra_experience.
func (ke kendraExperienceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ke.ref.Append("id"))
}

// IndexId returns a reference to field index_id of aws_kendra_experience.
func (ke kendraExperienceAttributes) IndexId() terra.StringValue {
	return terra.ReferenceAsString(ke.ref.Append("index_id"))
}

// Name returns a reference to field name of aws_kendra_experience.
func (ke kendraExperienceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ke.ref.Append("name"))
}

// RoleArn returns a reference to field role_arn of aws_kendra_experience.
func (ke kendraExperienceAttributes) RoleArn() terra.StringValue {
	return terra.ReferenceAsString(ke.ref.Append("role_arn"))
}

// Status returns a reference to field status of aws_kendra_experience.
func (ke kendraExperienceAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(ke.ref.Append("status"))
}

func (ke kendraExperienceAttributes) Endpoints() terra.SetValue[kendraexperience.EndpointsAttributes] {
	return terra.ReferenceAsSet[kendraexperience.EndpointsAttributes](ke.ref.Append("endpoints"))
}

func (ke kendraExperienceAttributes) Configuration() terra.ListValue[kendraexperience.ConfigurationAttributes] {
	return terra.ReferenceAsList[kendraexperience.ConfigurationAttributes](ke.ref.Append("configuration"))
}

func (ke kendraExperienceAttributes) Timeouts() kendraexperience.TimeoutsAttributes {
	return terra.ReferenceAsSingle[kendraexperience.TimeoutsAttributes](ke.ref.Append("timeouts"))
}

type kendraExperienceState struct {
	Arn           string                                `json:"arn"`
	Description   string                                `json:"description"`
	ExperienceId  string                                `json:"experience_id"`
	Id            string                                `json:"id"`
	IndexId       string                                `json:"index_id"`
	Name          string                                `json:"name"`
	RoleArn       string                                `json:"role_arn"`
	Status        string                                `json:"status"`
	Endpoints     []kendraexperience.EndpointsState     `json:"endpoints"`
	Configuration []kendraexperience.ConfigurationState `json:"configuration"`
	Timeouts      *kendraexperience.TimeoutsState       `json:"timeouts"`
}
