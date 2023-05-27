// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	datassmincidentsresponseplan "github.com/golingon/terraproviders/aws/5.0.1/datassmincidentsresponseplan"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataSsmincidentsResponsePlan creates a new instance of [DataSsmincidentsResponsePlan].
func NewDataSsmincidentsResponsePlan(name string, args DataSsmincidentsResponsePlanArgs) *DataSsmincidentsResponsePlan {
	return &DataSsmincidentsResponsePlan{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataSsmincidentsResponsePlan)(nil)

// DataSsmincidentsResponsePlan represents the Terraform data resource aws_ssmincidents_response_plan.
type DataSsmincidentsResponsePlan struct {
	Name string
	Args DataSsmincidentsResponsePlanArgs
}

// DataSource returns the Terraform object type for [DataSsmincidentsResponsePlan].
func (srp *DataSsmincidentsResponsePlan) DataSource() string {
	return "aws_ssmincidents_response_plan"
}

// LocalName returns the local name for [DataSsmincidentsResponsePlan].
func (srp *DataSsmincidentsResponsePlan) LocalName() string {
	return srp.Name
}

// Configuration returns the configuration (args) for [DataSsmincidentsResponsePlan].
func (srp *DataSsmincidentsResponsePlan) Configuration() interface{} {
	return srp.Args
}

// Attributes returns the attributes for [DataSsmincidentsResponsePlan].
func (srp *DataSsmincidentsResponsePlan) Attributes() dataSsmincidentsResponsePlanAttributes {
	return dataSsmincidentsResponsePlanAttributes{ref: terra.ReferenceDataResource(srp)}
}

// DataSsmincidentsResponsePlanArgs contains the configurations for aws_ssmincidents_response_plan.
type DataSsmincidentsResponsePlanArgs struct {
	// Arn: string, required
	Arn terra.StringValue `hcl:"arn,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Action: min=0
	Action []datassmincidentsresponseplan.Action `hcl:"action,block" validate:"min=0"`
	// IncidentTemplate: min=0
	IncidentTemplate []datassmincidentsresponseplan.IncidentTemplate `hcl:"incident_template,block" validate:"min=0"`
	// Integration: min=0
	Integration []datassmincidentsresponseplan.Integration `hcl:"integration,block" validate:"min=0"`
}
type dataSsmincidentsResponsePlanAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_ssmincidents_response_plan.
func (srp dataSsmincidentsResponsePlanAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(srp.ref.Append("arn"))
}

// ChatChannel returns a reference to field chat_channel of aws_ssmincidents_response_plan.
func (srp dataSsmincidentsResponsePlanAttributes) ChatChannel() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](srp.ref.Append("chat_channel"))
}

// DisplayName returns a reference to field display_name of aws_ssmincidents_response_plan.
func (srp dataSsmincidentsResponsePlanAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(srp.ref.Append("display_name"))
}

// Engagements returns a reference to field engagements of aws_ssmincidents_response_plan.
func (srp dataSsmincidentsResponsePlanAttributes) Engagements() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](srp.ref.Append("engagements"))
}

// Id returns a reference to field id of aws_ssmincidents_response_plan.
func (srp dataSsmincidentsResponsePlanAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(srp.ref.Append("id"))
}

// Name returns a reference to field name of aws_ssmincidents_response_plan.
func (srp dataSsmincidentsResponsePlanAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(srp.ref.Append("name"))
}

// Tags returns a reference to field tags of aws_ssmincidents_response_plan.
func (srp dataSsmincidentsResponsePlanAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](srp.ref.Append("tags"))
}

func (srp dataSsmincidentsResponsePlanAttributes) Action() terra.ListValue[datassmincidentsresponseplan.ActionAttributes] {
	return terra.ReferenceAsList[datassmincidentsresponseplan.ActionAttributes](srp.ref.Append("action"))
}

func (srp dataSsmincidentsResponsePlanAttributes) IncidentTemplate() terra.ListValue[datassmincidentsresponseplan.IncidentTemplateAttributes] {
	return terra.ReferenceAsList[datassmincidentsresponseplan.IncidentTemplateAttributes](srp.ref.Append("incident_template"))
}

func (srp dataSsmincidentsResponsePlanAttributes) Integration() terra.ListValue[datassmincidentsresponseplan.IntegrationAttributes] {
	return terra.ReferenceAsList[datassmincidentsresponseplan.IntegrationAttributes](srp.ref.Append("integration"))
}
