// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	chimesdkvoicesiprule "github.com/golingon/terraproviders/aws/5.6.2/chimesdkvoicesiprule"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewChimesdkvoiceSipRule creates a new instance of [ChimesdkvoiceSipRule].
func NewChimesdkvoiceSipRule(name string, args ChimesdkvoiceSipRuleArgs) *ChimesdkvoiceSipRule {
	return &ChimesdkvoiceSipRule{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ChimesdkvoiceSipRule)(nil)

// ChimesdkvoiceSipRule represents the Terraform resource aws_chimesdkvoice_sip_rule.
type ChimesdkvoiceSipRule struct {
	Name      string
	Args      ChimesdkvoiceSipRuleArgs
	state     *chimesdkvoiceSipRuleState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ChimesdkvoiceSipRule].
func (csr *ChimesdkvoiceSipRule) Type() string {
	return "aws_chimesdkvoice_sip_rule"
}

// LocalName returns the local name for [ChimesdkvoiceSipRule].
func (csr *ChimesdkvoiceSipRule) LocalName() string {
	return csr.Name
}

// Configuration returns the configuration (args) for [ChimesdkvoiceSipRule].
func (csr *ChimesdkvoiceSipRule) Configuration() interface{} {
	return csr.Args
}

// DependOn is used for other resources to depend on [ChimesdkvoiceSipRule].
func (csr *ChimesdkvoiceSipRule) DependOn() terra.Reference {
	return terra.ReferenceResource(csr)
}

// Dependencies returns the list of resources [ChimesdkvoiceSipRule] depends_on.
func (csr *ChimesdkvoiceSipRule) Dependencies() terra.Dependencies {
	return csr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ChimesdkvoiceSipRule].
func (csr *ChimesdkvoiceSipRule) LifecycleManagement() *terra.Lifecycle {
	return csr.Lifecycle
}

// Attributes returns the attributes for [ChimesdkvoiceSipRule].
func (csr *ChimesdkvoiceSipRule) Attributes() chimesdkvoiceSipRuleAttributes {
	return chimesdkvoiceSipRuleAttributes{ref: terra.ReferenceResource(csr)}
}

// ImportState imports the given attribute values into [ChimesdkvoiceSipRule]'s state.
func (csr *ChimesdkvoiceSipRule) ImportState(av io.Reader) error {
	csr.state = &chimesdkvoiceSipRuleState{}
	if err := json.NewDecoder(av).Decode(csr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", csr.Type(), csr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ChimesdkvoiceSipRule] has state.
func (csr *ChimesdkvoiceSipRule) State() (*chimesdkvoiceSipRuleState, bool) {
	return csr.state, csr.state != nil
}

// StateMust returns the state for [ChimesdkvoiceSipRule]. Panics if the state is nil.
func (csr *ChimesdkvoiceSipRule) StateMust() *chimesdkvoiceSipRuleState {
	if csr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", csr.Type(), csr.LocalName()))
	}
	return csr.state
}

// ChimesdkvoiceSipRuleArgs contains the configurations for aws_chimesdkvoice_sip_rule.
type ChimesdkvoiceSipRuleArgs struct {
	// Disabled: bool, optional
	Disabled terra.BoolValue `hcl:"disabled,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// TriggerType: string, required
	TriggerType terra.StringValue `hcl:"trigger_type,attr" validate:"required"`
	// TriggerValue: string, required
	TriggerValue terra.StringValue `hcl:"trigger_value,attr" validate:"required"`
	// TargetApplications: min=1,max=25
	TargetApplications []chimesdkvoicesiprule.TargetApplications `hcl:"target_applications,block" validate:"min=1,max=25"`
}
type chimesdkvoiceSipRuleAttributes struct {
	ref terra.Reference
}

// Disabled returns a reference to field disabled of aws_chimesdkvoice_sip_rule.
func (csr chimesdkvoiceSipRuleAttributes) Disabled() terra.BoolValue {
	return terra.ReferenceAsBool(csr.ref.Append("disabled"))
}

// Id returns a reference to field id of aws_chimesdkvoice_sip_rule.
func (csr chimesdkvoiceSipRuleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(csr.ref.Append("id"))
}

// Name returns a reference to field name of aws_chimesdkvoice_sip_rule.
func (csr chimesdkvoiceSipRuleAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(csr.ref.Append("name"))
}

// TriggerType returns a reference to field trigger_type of aws_chimesdkvoice_sip_rule.
func (csr chimesdkvoiceSipRuleAttributes) TriggerType() terra.StringValue {
	return terra.ReferenceAsString(csr.ref.Append("trigger_type"))
}

// TriggerValue returns a reference to field trigger_value of aws_chimesdkvoice_sip_rule.
func (csr chimesdkvoiceSipRuleAttributes) TriggerValue() terra.StringValue {
	return terra.ReferenceAsString(csr.ref.Append("trigger_value"))
}

func (csr chimesdkvoiceSipRuleAttributes) TargetApplications() terra.SetValue[chimesdkvoicesiprule.TargetApplicationsAttributes] {
	return terra.ReferenceAsSet[chimesdkvoicesiprule.TargetApplicationsAttributes](csr.ref.Append("target_applications"))
}

type chimesdkvoiceSipRuleState struct {
	Disabled           bool                                           `json:"disabled"`
	Id                 string                                         `json:"id"`
	Name               string                                         `json:"name"`
	TriggerType        string                                         `json:"trigger_type"`
	TriggerValue       string                                         `json:"trigger_value"`
	TargetApplications []chimesdkvoicesiprule.TargetApplicationsState `json:"target_applications"`
}
