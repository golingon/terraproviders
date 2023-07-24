// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	devtestpolicy "github.com/golingon/terraproviders/azurerm/3.66.0/devtestpolicy"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDevTestPolicy creates a new instance of [DevTestPolicy].
func NewDevTestPolicy(name string, args DevTestPolicyArgs) *DevTestPolicy {
	return &DevTestPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DevTestPolicy)(nil)

// DevTestPolicy represents the Terraform resource azurerm_dev_test_policy.
type DevTestPolicy struct {
	Name      string
	Args      DevTestPolicyArgs
	state     *devTestPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DevTestPolicy].
func (dtp *DevTestPolicy) Type() string {
	return "azurerm_dev_test_policy"
}

// LocalName returns the local name for [DevTestPolicy].
func (dtp *DevTestPolicy) LocalName() string {
	return dtp.Name
}

// Configuration returns the configuration (args) for [DevTestPolicy].
func (dtp *DevTestPolicy) Configuration() interface{} {
	return dtp.Args
}

// DependOn is used for other resources to depend on [DevTestPolicy].
func (dtp *DevTestPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(dtp)
}

// Dependencies returns the list of resources [DevTestPolicy] depends_on.
func (dtp *DevTestPolicy) Dependencies() terra.Dependencies {
	return dtp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DevTestPolicy].
func (dtp *DevTestPolicy) LifecycleManagement() *terra.Lifecycle {
	return dtp.Lifecycle
}

// Attributes returns the attributes for [DevTestPolicy].
func (dtp *DevTestPolicy) Attributes() devTestPolicyAttributes {
	return devTestPolicyAttributes{ref: terra.ReferenceResource(dtp)}
}

// ImportState imports the given attribute values into [DevTestPolicy]'s state.
func (dtp *DevTestPolicy) ImportState(av io.Reader) error {
	dtp.state = &devTestPolicyState{}
	if err := json.NewDecoder(av).Decode(dtp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dtp.Type(), dtp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DevTestPolicy] has state.
func (dtp *DevTestPolicy) State() (*devTestPolicyState, bool) {
	return dtp.state, dtp.state != nil
}

// StateMust returns the state for [DevTestPolicy]. Panics if the state is nil.
func (dtp *DevTestPolicy) StateMust() *devTestPolicyState {
	if dtp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dtp.Type(), dtp.LocalName()))
	}
	return dtp.state
}

// DevTestPolicyArgs contains the configurations for azurerm_dev_test_policy.
type DevTestPolicyArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// EvaluatorType: string, required
	EvaluatorType terra.StringValue `hcl:"evaluator_type,attr" validate:"required"`
	// FactData: string, optional
	FactData terra.StringValue `hcl:"fact_data,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LabName: string, required
	LabName terra.StringValue `hcl:"lab_name,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PolicySetName: string, required
	PolicySetName terra.StringValue `hcl:"policy_set_name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Threshold: string, required
	Threshold terra.StringValue `hcl:"threshold,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *devtestpolicy.Timeouts `hcl:"timeouts,block"`
}
type devTestPolicyAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of azurerm_dev_test_policy.
func (dtp devTestPolicyAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(dtp.ref.Append("description"))
}

// EvaluatorType returns a reference to field evaluator_type of azurerm_dev_test_policy.
func (dtp devTestPolicyAttributes) EvaluatorType() terra.StringValue {
	return terra.ReferenceAsString(dtp.ref.Append("evaluator_type"))
}

// FactData returns a reference to field fact_data of azurerm_dev_test_policy.
func (dtp devTestPolicyAttributes) FactData() terra.StringValue {
	return terra.ReferenceAsString(dtp.ref.Append("fact_data"))
}

// Id returns a reference to field id of azurerm_dev_test_policy.
func (dtp devTestPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dtp.ref.Append("id"))
}

// LabName returns a reference to field lab_name of azurerm_dev_test_policy.
func (dtp devTestPolicyAttributes) LabName() terra.StringValue {
	return terra.ReferenceAsString(dtp.ref.Append("lab_name"))
}

// Name returns a reference to field name of azurerm_dev_test_policy.
func (dtp devTestPolicyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dtp.ref.Append("name"))
}

// PolicySetName returns a reference to field policy_set_name of azurerm_dev_test_policy.
func (dtp devTestPolicyAttributes) PolicySetName() terra.StringValue {
	return terra.ReferenceAsString(dtp.ref.Append("policy_set_name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_dev_test_policy.
func (dtp devTestPolicyAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(dtp.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_dev_test_policy.
func (dtp devTestPolicyAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dtp.ref.Append("tags"))
}

// Threshold returns a reference to field threshold of azurerm_dev_test_policy.
func (dtp devTestPolicyAttributes) Threshold() terra.StringValue {
	return terra.ReferenceAsString(dtp.ref.Append("threshold"))
}

func (dtp devTestPolicyAttributes) Timeouts() devtestpolicy.TimeoutsAttributes {
	return terra.ReferenceAsSingle[devtestpolicy.TimeoutsAttributes](dtp.ref.Append("timeouts"))
}

type devTestPolicyState struct {
	Description       string                       `json:"description"`
	EvaluatorType     string                       `json:"evaluator_type"`
	FactData          string                       `json:"fact_data"`
	Id                string                       `json:"id"`
	LabName           string                       `json:"lab_name"`
	Name              string                       `json:"name"`
	PolicySetName     string                       `json:"policy_set_name"`
	ResourceGroupName string                       `json:"resource_group_name"`
	Tags              map[string]string            `json:"tags"`
	Threshold         string                       `json:"threshold"`
	Timeouts          *devtestpolicy.TimeoutsState `json:"timeouts"`
}
