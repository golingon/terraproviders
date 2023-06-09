// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataec2networkinsightsanalysis "github.com/golingon/terraproviders/aws/4.60.0/dataec2networkinsightsanalysis"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataEc2NetworkInsightsAnalysis creates a new instance of [DataEc2NetworkInsightsAnalysis].
func NewDataEc2NetworkInsightsAnalysis(name string, args DataEc2NetworkInsightsAnalysisArgs) *DataEc2NetworkInsightsAnalysis {
	return &DataEc2NetworkInsightsAnalysis{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataEc2NetworkInsightsAnalysis)(nil)

// DataEc2NetworkInsightsAnalysis represents the Terraform data resource aws_ec2_network_insights_analysis.
type DataEc2NetworkInsightsAnalysis struct {
	Name string
	Args DataEc2NetworkInsightsAnalysisArgs
}

// DataSource returns the Terraform object type for [DataEc2NetworkInsightsAnalysis].
func (enia *DataEc2NetworkInsightsAnalysis) DataSource() string {
	return "aws_ec2_network_insights_analysis"
}

// LocalName returns the local name for [DataEc2NetworkInsightsAnalysis].
func (enia *DataEc2NetworkInsightsAnalysis) LocalName() string {
	return enia.Name
}

// Configuration returns the configuration (args) for [DataEc2NetworkInsightsAnalysis].
func (enia *DataEc2NetworkInsightsAnalysis) Configuration() interface{} {
	return enia.Args
}

// Attributes returns the attributes for [DataEc2NetworkInsightsAnalysis].
func (enia *DataEc2NetworkInsightsAnalysis) Attributes() dataEc2NetworkInsightsAnalysisAttributes {
	return dataEc2NetworkInsightsAnalysisAttributes{ref: terra.ReferenceDataResource(enia)}
}

// DataEc2NetworkInsightsAnalysisArgs contains the configurations for aws_ec2_network_insights_analysis.
type DataEc2NetworkInsightsAnalysisArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// NetworkInsightsAnalysisId: string, optional
	NetworkInsightsAnalysisId terra.StringValue `hcl:"network_insights_analysis_id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// AlternatePathHints: min=0
	AlternatePathHints []dataec2networkinsightsanalysis.AlternatePathHints `hcl:"alternate_path_hints,block" validate:"min=0"`
	// Explanations: min=0
	Explanations []dataec2networkinsightsanalysis.Explanations `hcl:"explanations,block" validate:"min=0"`
	// ForwardPathComponents: min=0
	ForwardPathComponents []dataec2networkinsightsanalysis.ForwardPathComponents `hcl:"forward_path_components,block" validate:"min=0"`
	// ReturnPathComponents: min=0
	ReturnPathComponents []dataec2networkinsightsanalysis.ReturnPathComponents `hcl:"return_path_components,block" validate:"min=0"`
	// Filter: min=0
	Filter []dataec2networkinsightsanalysis.Filter `hcl:"filter,block" validate:"min=0"`
}
type dataEc2NetworkInsightsAnalysisAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_ec2_network_insights_analysis.
func (enia dataEc2NetworkInsightsAnalysisAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(enia.ref.Append("arn"))
}

// FilterInArns returns a reference to field filter_in_arns of aws_ec2_network_insights_analysis.
func (enia dataEc2NetworkInsightsAnalysisAttributes) FilterInArns() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](enia.ref.Append("filter_in_arns"))
}

// Id returns a reference to field id of aws_ec2_network_insights_analysis.
func (enia dataEc2NetworkInsightsAnalysisAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(enia.ref.Append("id"))
}

// NetworkInsightsAnalysisId returns a reference to field network_insights_analysis_id of aws_ec2_network_insights_analysis.
func (enia dataEc2NetworkInsightsAnalysisAttributes) NetworkInsightsAnalysisId() terra.StringValue {
	return terra.ReferenceAsString(enia.ref.Append("network_insights_analysis_id"))
}

// NetworkInsightsPathId returns a reference to field network_insights_path_id of aws_ec2_network_insights_analysis.
func (enia dataEc2NetworkInsightsAnalysisAttributes) NetworkInsightsPathId() terra.StringValue {
	return terra.ReferenceAsString(enia.ref.Append("network_insights_path_id"))
}

// PathFound returns a reference to field path_found of aws_ec2_network_insights_analysis.
func (enia dataEc2NetworkInsightsAnalysisAttributes) PathFound() terra.BoolValue {
	return terra.ReferenceAsBool(enia.ref.Append("path_found"))
}

// StartDate returns a reference to field start_date of aws_ec2_network_insights_analysis.
func (enia dataEc2NetworkInsightsAnalysisAttributes) StartDate() terra.StringValue {
	return terra.ReferenceAsString(enia.ref.Append("start_date"))
}

// Status returns a reference to field status of aws_ec2_network_insights_analysis.
func (enia dataEc2NetworkInsightsAnalysisAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(enia.ref.Append("status"))
}

// StatusMessage returns a reference to field status_message of aws_ec2_network_insights_analysis.
func (enia dataEc2NetworkInsightsAnalysisAttributes) StatusMessage() terra.StringValue {
	return terra.ReferenceAsString(enia.ref.Append("status_message"))
}

// Tags returns a reference to field tags of aws_ec2_network_insights_analysis.
func (enia dataEc2NetworkInsightsAnalysisAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](enia.ref.Append("tags"))
}

// WarningMessage returns a reference to field warning_message of aws_ec2_network_insights_analysis.
func (enia dataEc2NetworkInsightsAnalysisAttributes) WarningMessage() terra.StringValue {
	return terra.ReferenceAsString(enia.ref.Append("warning_message"))
}

func (enia dataEc2NetworkInsightsAnalysisAttributes) AlternatePathHints() terra.ListValue[dataec2networkinsightsanalysis.AlternatePathHintsAttributes] {
	return terra.ReferenceAsList[dataec2networkinsightsanalysis.AlternatePathHintsAttributes](enia.ref.Append("alternate_path_hints"))
}

func (enia dataEc2NetworkInsightsAnalysisAttributes) Explanations() terra.ListValue[dataec2networkinsightsanalysis.ExplanationsAttributes] {
	return terra.ReferenceAsList[dataec2networkinsightsanalysis.ExplanationsAttributes](enia.ref.Append("explanations"))
}

func (enia dataEc2NetworkInsightsAnalysisAttributes) ForwardPathComponents() terra.ListValue[dataec2networkinsightsanalysis.ForwardPathComponentsAttributes] {
	return terra.ReferenceAsList[dataec2networkinsightsanalysis.ForwardPathComponentsAttributes](enia.ref.Append("forward_path_components"))
}

func (enia dataEc2NetworkInsightsAnalysisAttributes) ReturnPathComponents() terra.ListValue[dataec2networkinsightsanalysis.ReturnPathComponentsAttributes] {
	return terra.ReferenceAsList[dataec2networkinsightsanalysis.ReturnPathComponentsAttributes](enia.ref.Append("return_path_components"))
}

func (enia dataEc2NetworkInsightsAnalysisAttributes) Filter() terra.SetValue[dataec2networkinsightsanalysis.FilterAttributes] {
	return terra.ReferenceAsSet[dataec2networkinsightsanalysis.FilterAttributes](enia.ref.Append("filter"))
}
