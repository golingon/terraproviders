// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	ssmincidentsresponseplan "github.com/golingon/terraproviders/aws/5.6.2/ssmincidentsresponseplan"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSsmincidentsResponsePlan creates a new instance of [SsmincidentsResponsePlan].
func NewSsmincidentsResponsePlan(name string, args SsmincidentsResponsePlanArgs) *SsmincidentsResponsePlan {
	return &SsmincidentsResponsePlan{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SsmincidentsResponsePlan)(nil)

// SsmincidentsResponsePlan represents the Terraform resource aws_ssmincidents_response_plan.
type SsmincidentsResponsePlan struct {
	Name      string
	Args      SsmincidentsResponsePlanArgs
	state     *ssmincidentsResponsePlanState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SsmincidentsResponsePlan].
func (srp *SsmincidentsResponsePlan) Type() string {
	return "aws_ssmincidents_response_plan"
}

// LocalName returns the local name for [SsmincidentsResponsePlan].
func (srp *SsmincidentsResponsePlan) LocalName() string {
	return srp.Name
}

// Configuration returns the configuration (args) for [SsmincidentsResponsePlan].
func (srp *SsmincidentsResponsePlan) Configuration() interface{} {
	return srp.Args
}

// DependOn is used for other resources to depend on [SsmincidentsResponsePlan].
func (srp *SsmincidentsResponsePlan) DependOn() terra.Reference {
	return terra.ReferenceResource(srp)
}

// Dependencies returns the list of resources [SsmincidentsResponsePlan] depends_on.
func (srp *SsmincidentsResponsePlan) Dependencies() terra.Dependencies {
	return srp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SsmincidentsResponsePlan].
func (srp *SsmincidentsResponsePlan) LifecycleManagement() *terra.Lifecycle {
	return srp.Lifecycle
}

// Attributes returns the attributes for [SsmincidentsResponsePlan].
func (srp *SsmincidentsResponsePlan) Attributes() ssmincidentsResponsePlanAttributes {
	return ssmincidentsResponsePlanAttributes{ref: terra.ReferenceResource(srp)}
}

// ImportState imports the given attribute values into [SsmincidentsResponsePlan]'s state.
func (srp *SsmincidentsResponsePlan) ImportState(av io.Reader) error {
	srp.state = &ssmincidentsResponsePlanState{}
	if err := json.NewDecoder(av).Decode(srp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", srp.Type(), srp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SsmincidentsResponsePlan] has state.
func (srp *SsmincidentsResponsePlan) State() (*ssmincidentsResponsePlanState, bool) {
	return srp.state, srp.state != nil
}

// StateMust returns the state for [SsmincidentsResponsePlan]. Panics if the state is nil.
func (srp *SsmincidentsResponsePlan) StateMust() *ssmincidentsResponsePlanState {
	if srp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", srp.Type(), srp.LocalName()))
	}
	return srp.state
}

// SsmincidentsResponsePlanArgs contains the configurations for aws_ssmincidents_response_plan.
type SsmincidentsResponsePlanArgs struct {
	// ChatChannel: set of string, optional
	ChatChannel terra.SetValue[terra.StringValue] `hcl:"chat_channel,attr"`
	// DisplayName: string, optional
	DisplayName terra.StringValue `hcl:"display_name,attr"`
	// Engagements: set of string, optional
	Engagements terra.SetValue[terra.StringValue] `hcl:"engagements,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// Action: optional
	Action *ssmincidentsresponseplan.Action `hcl:"action,block"`
	// IncidentTemplate: required
	IncidentTemplate *ssmincidentsresponseplan.IncidentTemplate `hcl:"incident_template,block" validate:"required"`
	// Integration: optional
	Integration *ssmincidentsresponseplan.Integration `hcl:"integration,block"`
}
type ssmincidentsResponsePlanAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_ssmincidents_response_plan.
func (srp ssmincidentsResponsePlanAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(srp.ref.Append("arn"))
}

// ChatChannel returns a reference to field chat_channel of aws_ssmincidents_response_plan.
func (srp ssmincidentsResponsePlanAttributes) ChatChannel() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](srp.ref.Append("chat_channel"))
}

// DisplayName returns a reference to field display_name of aws_ssmincidents_response_plan.
func (srp ssmincidentsResponsePlanAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(srp.ref.Append("display_name"))
}

// Engagements returns a reference to field engagements of aws_ssmincidents_response_plan.
func (srp ssmincidentsResponsePlanAttributes) Engagements() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](srp.ref.Append("engagements"))
}

// Id returns a reference to field id of aws_ssmincidents_response_plan.
func (srp ssmincidentsResponsePlanAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(srp.ref.Append("id"))
}

// Name returns a reference to field name of aws_ssmincidents_response_plan.
func (srp ssmincidentsResponsePlanAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(srp.ref.Append("name"))
}

// Tags returns a reference to field tags of aws_ssmincidents_response_plan.
func (srp ssmincidentsResponsePlanAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](srp.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_ssmincidents_response_plan.
func (srp ssmincidentsResponsePlanAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](srp.ref.Append("tags_all"))
}

func (srp ssmincidentsResponsePlanAttributes) Action() terra.ListValue[ssmincidentsresponseplan.ActionAttributes] {
	return terra.ReferenceAsList[ssmincidentsresponseplan.ActionAttributes](srp.ref.Append("action"))
}

func (srp ssmincidentsResponsePlanAttributes) IncidentTemplate() terra.ListValue[ssmincidentsresponseplan.IncidentTemplateAttributes] {
	return terra.ReferenceAsList[ssmincidentsresponseplan.IncidentTemplateAttributes](srp.ref.Append("incident_template"))
}

func (srp ssmincidentsResponsePlanAttributes) Integration() terra.ListValue[ssmincidentsresponseplan.IntegrationAttributes] {
	return terra.ReferenceAsList[ssmincidentsresponseplan.IntegrationAttributes](srp.ref.Append("integration"))
}

type ssmincidentsResponsePlanState struct {
	Arn              string                                           `json:"arn"`
	ChatChannel      []string                                         `json:"chat_channel"`
	DisplayName      string                                           `json:"display_name"`
	Engagements      []string                                         `json:"engagements"`
	Id               string                                           `json:"id"`
	Name             string                                           `json:"name"`
	Tags             map[string]string                                `json:"tags"`
	TagsAll          map[string]string                                `json:"tags_all"`
	Action           []ssmincidentsresponseplan.ActionState           `json:"action"`
	IncidentTemplate []ssmincidentsresponseplan.IncidentTemplateState `json:"incident_template"`
	Integration      []ssmincidentsresponseplan.IntegrationState      `json:"integration"`
}
