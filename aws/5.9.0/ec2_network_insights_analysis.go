// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	ec2networkinsightsanalysis "github.com/golingon/terraproviders/aws/5.9.0/ec2networkinsightsanalysis"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewEc2NetworkInsightsAnalysis creates a new instance of [Ec2NetworkInsightsAnalysis].
func NewEc2NetworkInsightsAnalysis(name string, args Ec2NetworkInsightsAnalysisArgs) *Ec2NetworkInsightsAnalysis {
	return &Ec2NetworkInsightsAnalysis{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Ec2NetworkInsightsAnalysis)(nil)

// Ec2NetworkInsightsAnalysis represents the Terraform resource aws_ec2_network_insights_analysis.
type Ec2NetworkInsightsAnalysis struct {
	Name      string
	Args      Ec2NetworkInsightsAnalysisArgs
	state     *ec2NetworkInsightsAnalysisState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Ec2NetworkInsightsAnalysis].
func (enia *Ec2NetworkInsightsAnalysis) Type() string {
	return "aws_ec2_network_insights_analysis"
}

// LocalName returns the local name for [Ec2NetworkInsightsAnalysis].
func (enia *Ec2NetworkInsightsAnalysis) LocalName() string {
	return enia.Name
}

// Configuration returns the configuration (args) for [Ec2NetworkInsightsAnalysis].
func (enia *Ec2NetworkInsightsAnalysis) Configuration() interface{} {
	return enia.Args
}

// DependOn is used for other resources to depend on [Ec2NetworkInsightsAnalysis].
func (enia *Ec2NetworkInsightsAnalysis) DependOn() terra.Reference {
	return terra.ReferenceResource(enia)
}

// Dependencies returns the list of resources [Ec2NetworkInsightsAnalysis] depends_on.
func (enia *Ec2NetworkInsightsAnalysis) Dependencies() terra.Dependencies {
	return enia.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Ec2NetworkInsightsAnalysis].
func (enia *Ec2NetworkInsightsAnalysis) LifecycleManagement() *terra.Lifecycle {
	return enia.Lifecycle
}

// Attributes returns the attributes for [Ec2NetworkInsightsAnalysis].
func (enia *Ec2NetworkInsightsAnalysis) Attributes() ec2NetworkInsightsAnalysisAttributes {
	return ec2NetworkInsightsAnalysisAttributes{ref: terra.ReferenceResource(enia)}
}

// ImportState imports the given attribute values into [Ec2NetworkInsightsAnalysis]'s state.
func (enia *Ec2NetworkInsightsAnalysis) ImportState(av io.Reader) error {
	enia.state = &ec2NetworkInsightsAnalysisState{}
	if err := json.NewDecoder(av).Decode(enia.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", enia.Type(), enia.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Ec2NetworkInsightsAnalysis] has state.
func (enia *Ec2NetworkInsightsAnalysis) State() (*ec2NetworkInsightsAnalysisState, bool) {
	return enia.state, enia.state != nil
}

// StateMust returns the state for [Ec2NetworkInsightsAnalysis]. Panics if the state is nil.
func (enia *Ec2NetworkInsightsAnalysis) StateMust() *ec2NetworkInsightsAnalysisState {
	if enia.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", enia.Type(), enia.LocalName()))
	}
	return enia.state
}

// Ec2NetworkInsightsAnalysisArgs contains the configurations for aws_ec2_network_insights_analysis.
type Ec2NetworkInsightsAnalysisArgs struct {
	// FilterInArns: set of string, optional
	FilterInArns terra.SetValue[terra.StringValue] `hcl:"filter_in_arns,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// NetworkInsightsPathId: string, required
	NetworkInsightsPathId terra.StringValue `hcl:"network_insights_path_id,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// WaitForCompletion: bool, optional
	WaitForCompletion terra.BoolValue `hcl:"wait_for_completion,attr"`
	// AlternatePathHints: min=0
	AlternatePathHints []ec2networkinsightsanalysis.AlternatePathHints `hcl:"alternate_path_hints,block" validate:"min=0"`
	// Explanations: min=0
	Explanations []ec2networkinsightsanalysis.Explanations `hcl:"explanations,block" validate:"min=0"`
	// ForwardPathComponents: min=0
	ForwardPathComponents []ec2networkinsightsanalysis.ForwardPathComponents `hcl:"forward_path_components,block" validate:"min=0"`
	// ReturnPathComponents: min=0
	ReturnPathComponents []ec2networkinsightsanalysis.ReturnPathComponents `hcl:"return_path_components,block" validate:"min=0"`
}
type ec2NetworkInsightsAnalysisAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_ec2_network_insights_analysis.
func (enia ec2NetworkInsightsAnalysisAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(enia.ref.Append("arn"))
}

// FilterInArns returns a reference to field filter_in_arns of aws_ec2_network_insights_analysis.
func (enia ec2NetworkInsightsAnalysisAttributes) FilterInArns() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](enia.ref.Append("filter_in_arns"))
}

// Id returns a reference to field id of aws_ec2_network_insights_analysis.
func (enia ec2NetworkInsightsAnalysisAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(enia.ref.Append("id"))
}

// NetworkInsightsPathId returns a reference to field network_insights_path_id of aws_ec2_network_insights_analysis.
func (enia ec2NetworkInsightsAnalysisAttributes) NetworkInsightsPathId() terra.StringValue {
	return terra.ReferenceAsString(enia.ref.Append("network_insights_path_id"))
}

// PathFound returns a reference to field path_found of aws_ec2_network_insights_analysis.
func (enia ec2NetworkInsightsAnalysisAttributes) PathFound() terra.BoolValue {
	return terra.ReferenceAsBool(enia.ref.Append("path_found"))
}

// StartDate returns a reference to field start_date of aws_ec2_network_insights_analysis.
func (enia ec2NetworkInsightsAnalysisAttributes) StartDate() terra.StringValue {
	return terra.ReferenceAsString(enia.ref.Append("start_date"))
}

// Status returns a reference to field status of aws_ec2_network_insights_analysis.
func (enia ec2NetworkInsightsAnalysisAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(enia.ref.Append("status"))
}

// StatusMessage returns a reference to field status_message of aws_ec2_network_insights_analysis.
func (enia ec2NetworkInsightsAnalysisAttributes) StatusMessage() terra.StringValue {
	return terra.ReferenceAsString(enia.ref.Append("status_message"))
}

// Tags returns a reference to field tags of aws_ec2_network_insights_analysis.
func (enia ec2NetworkInsightsAnalysisAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](enia.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_ec2_network_insights_analysis.
func (enia ec2NetworkInsightsAnalysisAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](enia.ref.Append("tags_all"))
}

// WaitForCompletion returns a reference to field wait_for_completion of aws_ec2_network_insights_analysis.
func (enia ec2NetworkInsightsAnalysisAttributes) WaitForCompletion() terra.BoolValue {
	return terra.ReferenceAsBool(enia.ref.Append("wait_for_completion"))
}

// WarningMessage returns a reference to field warning_message of aws_ec2_network_insights_analysis.
func (enia ec2NetworkInsightsAnalysisAttributes) WarningMessage() terra.StringValue {
	return terra.ReferenceAsString(enia.ref.Append("warning_message"))
}

func (enia ec2NetworkInsightsAnalysisAttributes) AlternatePathHints() terra.ListValue[ec2networkinsightsanalysis.AlternatePathHintsAttributes] {
	return terra.ReferenceAsList[ec2networkinsightsanalysis.AlternatePathHintsAttributes](enia.ref.Append("alternate_path_hints"))
}

func (enia ec2NetworkInsightsAnalysisAttributes) Explanations() terra.ListValue[ec2networkinsightsanalysis.ExplanationsAttributes] {
	return terra.ReferenceAsList[ec2networkinsightsanalysis.ExplanationsAttributes](enia.ref.Append("explanations"))
}

func (enia ec2NetworkInsightsAnalysisAttributes) ForwardPathComponents() terra.ListValue[ec2networkinsightsanalysis.ForwardPathComponentsAttributes] {
	return terra.ReferenceAsList[ec2networkinsightsanalysis.ForwardPathComponentsAttributes](enia.ref.Append("forward_path_components"))
}

func (enia ec2NetworkInsightsAnalysisAttributes) ReturnPathComponents() terra.ListValue[ec2networkinsightsanalysis.ReturnPathComponentsAttributes] {
	return terra.ReferenceAsList[ec2networkinsightsanalysis.ReturnPathComponentsAttributes](enia.ref.Append("return_path_components"))
}

type ec2NetworkInsightsAnalysisState struct {
	Arn                   string                                                  `json:"arn"`
	FilterInArns          []string                                                `json:"filter_in_arns"`
	Id                    string                                                  `json:"id"`
	NetworkInsightsPathId string                                                  `json:"network_insights_path_id"`
	PathFound             bool                                                    `json:"path_found"`
	StartDate             string                                                  `json:"start_date"`
	Status                string                                                  `json:"status"`
	StatusMessage         string                                                  `json:"status_message"`
	Tags                  map[string]string                                       `json:"tags"`
	TagsAll               map[string]string                                       `json:"tags_all"`
	WaitForCompletion     bool                                                    `json:"wait_for_completion"`
	WarningMessage        string                                                  `json:"warning_message"`
	AlternatePathHints    []ec2networkinsightsanalysis.AlternatePathHintsState    `json:"alternate_path_hints"`
	Explanations          []ec2networkinsightsanalysis.ExplanationsState          `json:"explanations"`
	ForwardPathComponents []ec2networkinsightsanalysis.ForwardPathComponentsState `json:"forward_path_components"`
	ReturnPathComponents  []ec2networkinsightsanalysis.ReturnPathComponentsState  `json:"return_path_components"`
}
