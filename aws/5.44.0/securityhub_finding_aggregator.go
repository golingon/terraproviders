// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// NewSecurityhubFindingAggregator creates a new instance of [SecurityhubFindingAggregator].
func NewSecurityhubFindingAggregator(name string, args SecurityhubFindingAggregatorArgs) *SecurityhubFindingAggregator {
	return &SecurityhubFindingAggregator{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SecurityhubFindingAggregator)(nil)

// SecurityhubFindingAggregator represents the Terraform resource aws_securityhub_finding_aggregator.
type SecurityhubFindingAggregator struct {
	Name      string
	Args      SecurityhubFindingAggregatorArgs
	state     *securityhubFindingAggregatorState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SecurityhubFindingAggregator].
func (sfa *SecurityhubFindingAggregator) Type() string {
	return "aws_securityhub_finding_aggregator"
}

// LocalName returns the local name for [SecurityhubFindingAggregator].
func (sfa *SecurityhubFindingAggregator) LocalName() string {
	return sfa.Name
}

// Configuration returns the configuration (args) for [SecurityhubFindingAggregator].
func (sfa *SecurityhubFindingAggregator) Configuration() interface{} {
	return sfa.Args
}

// DependOn is used for other resources to depend on [SecurityhubFindingAggregator].
func (sfa *SecurityhubFindingAggregator) DependOn() terra.Reference {
	return terra.ReferenceResource(sfa)
}

// Dependencies returns the list of resources [SecurityhubFindingAggregator] depends_on.
func (sfa *SecurityhubFindingAggregator) Dependencies() terra.Dependencies {
	return sfa.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SecurityhubFindingAggregator].
func (sfa *SecurityhubFindingAggregator) LifecycleManagement() *terra.Lifecycle {
	return sfa.Lifecycle
}

// Attributes returns the attributes for [SecurityhubFindingAggregator].
func (sfa *SecurityhubFindingAggregator) Attributes() securityhubFindingAggregatorAttributes {
	return securityhubFindingAggregatorAttributes{ref: terra.ReferenceResource(sfa)}
}

// ImportState imports the given attribute values into [SecurityhubFindingAggregator]'s state.
func (sfa *SecurityhubFindingAggregator) ImportState(av io.Reader) error {
	sfa.state = &securityhubFindingAggregatorState{}
	if err := json.NewDecoder(av).Decode(sfa.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sfa.Type(), sfa.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SecurityhubFindingAggregator] has state.
func (sfa *SecurityhubFindingAggregator) State() (*securityhubFindingAggregatorState, bool) {
	return sfa.state, sfa.state != nil
}

// StateMust returns the state for [SecurityhubFindingAggregator]. Panics if the state is nil.
func (sfa *SecurityhubFindingAggregator) StateMust() *securityhubFindingAggregatorState {
	if sfa.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sfa.Type(), sfa.LocalName()))
	}
	return sfa.state
}

// SecurityhubFindingAggregatorArgs contains the configurations for aws_securityhub_finding_aggregator.
type SecurityhubFindingAggregatorArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LinkingMode: string, required
	LinkingMode terra.StringValue `hcl:"linking_mode,attr" validate:"required"`
	// SpecifiedRegions: set of string, optional
	SpecifiedRegions terra.SetValue[terra.StringValue] `hcl:"specified_regions,attr"`
}
type securityhubFindingAggregatorAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_securityhub_finding_aggregator.
func (sfa securityhubFindingAggregatorAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sfa.ref.Append("id"))
}

// LinkingMode returns a reference to field linking_mode of aws_securityhub_finding_aggregator.
func (sfa securityhubFindingAggregatorAttributes) LinkingMode() terra.StringValue {
	return terra.ReferenceAsString(sfa.ref.Append("linking_mode"))
}

// SpecifiedRegions returns a reference to field specified_regions of aws_securityhub_finding_aggregator.
func (sfa securityhubFindingAggregatorAttributes) SpecifiedRegions() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](sfa.ref.Append("specified_regions"))
}

type securityhubFindingAggregatorState struct {
	Id               string   `json:"id"`
	LinkingMode      string   `json:"linking_mode"`
	SpecifiedRegions []string `json:"specified_regions"`
}
