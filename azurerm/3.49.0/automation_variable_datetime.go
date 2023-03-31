// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	automationvariabledatetime "github.com/golingon/terraproviders/azurerm/3.49.0/automationvariabledatetime"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewAutomationVariableDatetime(name string, args AutomationVariableDatetimeArgs) *AutomationVariableDatetime {
	return &AutomationVariableDatetime{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AutomationVariableDatetime)(nil)

type AutomationVariableDatetime struct {
	Name  string
	Args  AutomationVariableDatetimeArgs
	state *automationVariableDatetimeState
}

func (avd *AutomationVariableDatetime) Type() string {
	return "azurerm_automation_variable_datetime"
}

func (avd *AutomationVariableDatetime) LocalName() string {
	return avd.Name
}

func (avd *AutomationVariableDatetime) Configuration() interface{} {
	return avd.Args
}

func (avd *AutomationVariableDatetime) Attributes() automationVariableDatetimeAttributes {
	return automationVariableDatetimeAttributes{ref: terra.ReferenceResource(avd)}
}

func (avd *AutomationVariableDatetime) ImportState(av io.Reader) error {
	avd.state = &automationVariableDatetimeState{}
	if err := json.NewDecoder(av).Decode(avd.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", avd.Type(), avd.LocalName(), err)
	}
	return nil
}

func (avd *AutomationVariableDatetime) State() (*automationVariableDatetimeState, bool) {
	return avd.state, avd.state != nil
}

func (avd *AutomationVariableDatetime) StateMust() *automationVariableDatetimeState {
	if avd.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", avd.Type(), avd.LocalName()))
	}
	return avd.state
}

func (avd *AutomationVariableDatetime) DependOn() terra.Reference {
	return terra.ReferenceResource(avd)
}

type AutomationVariableDatetimeArgs struct {
	// AutomationAccountName: string, required
	AutomationAccountName terra.StringValue `hcl:"automation_account_name,attr" validate:"required"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Encrypted: bool, optional
	Encrypted terra.BoolValue `hcl:"encrypted,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Value: string, optional
	Value terra.StringValue `hcl:"value,attr"`
	// Timeouts: optional
	Timeouts *automationvariabledatetime.Timeouts `hcl:"timeouts,block"`
	// DependsOn contains resources that AutomationVariableDatetime depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type automationVariableDatetimeAttributes struct {
	ref terra.Reference
}

func (avd automationVariableDatetimeAttributes) AutomationAccountName() terra.StringValue {
	return terra.ReferenceString(avd.ref.Append("automation_account_name"))
}

func (avd automationVariableDatetimeAttributes) Description() terra.StringValue {
	return terra.ReferenceString(avd.ref.Append("description"))
}

func (avd automationVariableDatetimeAttributes) Encrypted() terra.BoolValue {
	return terra.ReferenceBool(avd.ref.Append("encrypted"))
}

func (avd automationVariableDatetimeAttributes) Id() terra.StringValue {
	return terra.ReferenceString(avd.ref.Append("id"))
}

func (avd automationVariableDatetimeAttributes) Name() terra.StringValue {
	return terra.ReferenceString(avd.ref.Append("name"))
}

func (avd automationVariableDatetimeAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceString(avd.ref.Append("resource_group_name"))
}

func (avd automationVariableDatetimeAttributes) Value() terra.StringValue {
	return terra.ReferenceString(avd.ref.Append("value"))
}

func (avd automationVariableDatetimeAttributes) Timeouts() automationvariabledatetime.TimeoutsAttributes {
	return terra.ReferenceSingle[automationvariabledatetime.TimeoutsAttributes](avd.ref.Append("timeouts"))
}

type automationVariableDatetimeState struct {
	AutomationAccountName string                                    `json:"automation_account_name"`
	Description           string                                    `json:"description"`
	Encrypted             bool                                      `json:"encrypted"`
	Id                    string                                    `json:"id"`
	Name                  string                                    `json:"name"`
	ResourceGroupName     string                                    `json:"resource_group_name"`
	Value                 string                                    `json:"value"`
	Timeouts              *automationvariabledatetime.TimeoutsState `json:"timeouts"`
}
