// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_securityhub_finding_aggregator

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// New creates a new instance of [Resource].
func New(name string, args Args) *Resource {
	return &Resource{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Resource)(nil)

// Resource represents the Terraform resource aws_securityhub_finding_aggregator.
type Resource struct {
	Name      string
	Args      Args
	state     *awsSecurityhubFindingAggregatorState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (asfa *Resource) Type() string {
	return "aws_securityhub_finding_aggregator"
}

// LocalName returns the local name for [Resource].
func (asfa *Resource) LocalName() string {
	return asfa.Name
}

// Configuration returns the configuration (args) for [Resource].
func (asfa *Resource) Configuration() interface{} {
	return asfa.Args
}

// DependOn is used for other resources to depend on [Resource].
func (asfa *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(asfa)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (asfa *Resource) Dependencies() terra.Dependencies {
	return asfa.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (asfa *Resource) LifecycleManagement() *terra.Lifecycle {
	return asfa.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (asfa *Resource) Attributes() awsSecurityhubFindingAggregatorAttributes {
	return awsSecurityhubFindingAggregatorAttributes{ref: terra.ReferenceResource(asfa)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (asfa *Resource) ImportState(state io.Reader) error {
	asfa.state = &awsSecurityhubFindingAggregatorState{}
	if err := json.NewDecoder(state).Decode(asfa.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", asfa.Type(), asfa.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (asfa *Resource) State() (*awsSecurityhubFindingAggregatorState, bool) {
	return asfa.state, asfa.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (asfa *Resource) StateMust() *awsSecurityhubFindingAggregatorState {
	if asfa.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", asfa.Type(), asfa.LocalName()))
	}
	return asfa.state
}

// Args contains the configurations for aws_securityhub_finding_aggregator.
type Args struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LinkingMode: string, required
	LinkingMode terra.StringValue `hcl:"linking_mode,attr" validate:"required"`
	// SpecifiedRegions: set of string, optional
	SpecifiedRegions terra.SetValue[terra.StringValue] `hcl:"specified_regions,attr"`
}

type awsSecurityhubFindingAggregatorAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_securityhub_finding_aggregator.
func (asfa awsSecurityhubFindingAggregatorAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(asfa.ref.Append("id"))
}

// LinkingMode returns a reference to field linking_mode of aws_securityhub_finding_aggregator.
func (asfa awsSecurityhubFindingAggregatorAttributes) LinkingMode() terra.StringValue {
	return terra.ReferenceAsString(asfa.ref.Append("linking_mode"))
}

// SpecifiedRegions returns a reference to field specified_regions of aws_securityhub_finding_aggregator.
func (asfa awsSecurityhubFindingAggregatorAttributes) SpecifiedRegions() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](asfa.ref.Append("specified_regions"))
}

type awsSecurityhubFindingAggregatorState struct {
	Id               string   `json:"id"`
	LinkingMode      string   `json:"linking_mode"`
	SpecifiedRegions []string `json:"specified_regions"`
}
