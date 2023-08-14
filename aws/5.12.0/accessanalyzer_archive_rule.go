// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	accessanalyzerarchiverule "github.com/golingon/terraproviders/aws/5.12.0/accessanalyzerarchiverule"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewAccessanalyzerArchiveRule creates a new instance of [AccessanalyzerArchiveRule].
func NewAccessanalyzerArchiveRule(name string, args AccessanalyzerArchiveRuleArgs) *AccessanalyzerArchiveRule {
	return &AccessanalyzerArchiveRule{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AccessanalyzerArchiveRule)(nil)

// AccessanalyzerArchiveRule represents the Terraform resource aws_accessanalyzer_archive_rule.
type AccessanalyzerArchiveRule struct {
	Name      string
	Args      AccessanalyzerArchiveRuleArgs
	state     *accessanalyzerArchiveRuleState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [AccessanalyzerArchiveRule].
func (aar *AccessanalyzerArchiveRule) Type() string {
	return "aws_accessanalyzer_archive_rule"
}

// LocalName returns the local name for [AccessanalyzerArchiveRule].
func (aar *AccessanalyzerArchiveRule) LocalName() string {
	return aar.Name
}

// Configuration returns the configuration (args) for [AccessanalyzerArchiveRule].
func (aar *AccessanalyzerArchiveRule) Configuration() interface{} {
	return aar.Args
}

// DependOn is used for other resources to depend on [AccessanalyzerArchiveRule].
func (aar *AccessanalyzerArchiveRule) DependOn() terra.Reference {
	return terra.ReferenceResource(aar)
}

// Dependencies returns the list of resources [AccessanalyzerArchiveRule] depends_on.
func (aar *AccessanalyzerArchiveRule) Dependencies() terra.Dependencies {
	return aar.DependsOn
}

// LifecycleManagement returns the lifecycle block for [AccessanalyzerArchiveRule].
func (aar *AccessanalyzerArchiveRule) LifecycleManagement() *terra.Lifecycle {
	return aar.Lifecycle
}

// Attributes returns the attributes for [AccessanalyzerArchiveRule].
func (aar *AccessanalyzerArchiveRule) Attributes() accessanalyzerArchiveRuleAttributes {
	return accessanalyzerArchiveRuleAttributes{ref: terra.ReferenceResource(aar)}
}

// ImportState imports the given attribute values into [AccessanalyzerArchiveRule]'s state.
func (aar *AccessanalyzerArchiveRule) ImportState(av io.Reader) error {
	aar.state = &accessanalyzerArchiveRuleState{}
	if err := json.NewDecoder(av).Decode(aar.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", aar.Type(), aar.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [AccessanalyzerArchiveRule] has state.
func (aar *AccessanalyzerArchiveRule) State() (*accessanalyzerArchiveRuleState, bool) {
	return aar.state, aar.state != nil
}

// StateMust returns the state for [AccessanalyzerArchiveRule]. Panics if the state is nil.
func (aar *AccessanalyzerArchiveRule) StateMust() *accessanalyzerArchiveRuleState {
	if aar.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", aar.Type(), aar.LocalName()))
	}
	return aar.state
}

// AccessanalyzerArchiveRuleArgs contains the configurations for aws_accessanalyzer_archive_rule.
type AccessanalyzerArchiveRuleArgs struct {
	// AnalyzerName: string, required
	AnalyzerName terra.StringValue `hcl:"analyzer_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// RuleName: string, required
	RuleName terra.StringValue `hcl:"rule_name,attr" validate:"required"`
	// Filter: min=1
	Filter []accessanalyzerarchiverule.Filter `hcl:"filter,block" validate:"min=1"`
}
type accessanalyzerArchiveRuleAttributes struct {
	ref terra.Reference
}

// AnalyzerName returns a reference to field analyzer_name of aws_accessanalyzer_archive_rule.
func (aar accessanalyzerArchiveRuleAttributes) AnalyzerName() terra.StringValue {
	return terra.ReferenceAsString(aar.ref.Append("analyzer_name"))
}

// Id returns a reference to field id of aws_accessanalyzer_archive_rule.
func (aar accessanalyzerArchiveRuleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aar.ref.Append("id"))
}

// RuleName returns a reference to field rule_name of aws_accessanalyzer_archive_rule.
func (aar accessanalyzerArchiveRuleAttributes) RuleName() terra.StringValue {
	return terra.ReferenceAsString(aar.ref.Append("rule_name"))
}

func (aar accessanalyzerArchiveRuleAttributes) Filter() terra.SetValue[accessanalyzerarchiverule.FilterAttributes] {
	return terra.ReferenceAsSet[accessanalyzerarchiverule.FilterAttributes](aar.ref.Append("filter"))
}

type accessanalyzerArchiveRuleState struct {
	AnalyzerName string                                  `json:"analyzer_name"`
	Id           string                                  `json:"id"`
	RuleName     string                                  `json:"rule_name"`
	Filter       []accessanalyzerarchiverule.FilterState `json:"filter"`
}
