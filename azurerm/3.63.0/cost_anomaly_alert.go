// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	costanomalyalert "github.com/golingon/terraproviders/azurerm/3.63.0/costanomalyalert"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCostAnomalyAlert creates a new instance of [CostAnomalyAlert].
func NewCostAnomalyAlert(name string, args CostAnomalyAlertArgs) *CostAnomalyAlert {
	return &CostAnomalyAlert{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CostAnomalyAlert)(nil)

// CostAnomalyAlert represents the Terraform resource azurerm_cost_anomaly_alert.
type CostAnomalyAlert struct {
	Name      string
	Args      CostAnomalyAlertArgs
	state     *costAnomalyAlertState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CostAnomalyAlert].
func (caa *CostAnomalyAlert) Type() string {
	return "azurerm_cost_anomaly_alert"
}

// LocalName returns the local name for [CostAnomalyAlert].
func (caa *CostAnomalyAlert) LocalName() string {
	return caa.Name
}

// Configuration returns the configuration (args) for [CostAnomalyAlert].
func (caa *CostAnomalyAlert) Configuration() interface{} {
	return caa.Args
}

// DependOn is used for other resources to depend on [CostAnomalyAlert].
func (caa *CostAnomalyAlert) DependOn() terra.Reference {
	return terra.ReferenceResource(caa)
}

// Dependencies returns the list of resources [CostAnomalyAlert] depends_on.
func (caa *CostAnomalyAlert) Dependencies() terra.Dependencies {
	return caa.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CostAnomalyAlert].
func (caa *CostAnomalyAlert) LifecycleManagement() *terra.Lifecycle {
	return caa.Lifecycle
}

// Attributes returns the attributes for [CostAnomalyAlert].
func (caa *CostAnomalyAlert) Attributes() costAnomalyAlertAttributes {
	return costAnomalyAlertAttributes{ref: terra.ReferenceResource(caa)}
}

// ImportState imports the given attribute values into [CostAnomalyAlert]'s state.
func (caa *CostAnomalyAlert) ImportState(av io.Reader) error {
	caa.state = &costAnomalyAlertState{}
	if err := json.NewDecoder(av).Decode(caa.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", caa.Type(), caa.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CostAnomalyAlert] has state.
func (caa *CostAnomalyAlert) State() (*costAnomalyAlertState, bool) {
	return caa.state, caa.state != nil
}

// StateMust returns the state for [CostAnomalyAlert]. Panics if the state is nil.
func (caa *CostAnomalyAlert) StateMust() *costAnomalyAlertState {
	if caa.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", caa.Type(), caa.LocalName()))
	}
	return caa.state
}

// CostAnomalyAlertArgs contains the configurations for azurerm_cost_anomaly_alert.
type CostAnomalyAlertArgs struct {
	// DisplayName: string, required
	DisplayName terra.StringValue `hcl:"display_name,attr" validate:"required"`
	// EmailAddresses: set of string, required
	EmailAddresses terra.SetValue[terra.StringValue] `hcl:"email_addresses,attr" validate:"required"`
	// EmailSubject: string, required
	EmailSubject terra.StringValue `hcl:"email_subject,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Message: string, optional
	Message terra.StringValue `hcl:"message,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *costanomalyalert.Timeouts `hcl:"timeouts,block"`
}
type costAnomalyAlertAttributes struct {
	ref terra.Reference
}

// DisplayName returns a reference to field display_name of azurerm_cost_anomaly_alert.
func (caa costAnomalyAlertAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(caa.ref.Append("display_name"))
}

// EmailAddresses returns a reference to field email_addresses of azurerm_cost_anomaly_alert.
func (caa costAnomalyAlertAttributes) EmailAddresses() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](caa.ref.Append("email_addresses"))
}

// EmailSubject returns a reference to field email_subject of azurerm_cost_anomaly_alert.
func (caa costAnomalyAlertAttributes) EmailSubject() terra.StringValue {
	return terra.ReferenceAsString(caa.ref.Append("email_subject"))
}

// Id returns a reference to field id of azurerm_cost_anomaly_alert.
func (caa costAnomalyAlertAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(caa.ref.Append("id"))
}

// Message returns a reference to field message of azurerm_cost_anomaly_alert.
func (caa costAnomalyAlertAttributes) Message() terra.StringValue {
	return terra.ReferenceAsString(caa.ref.Append("message"))
}

// Name returns a reference to field name of azurerm_cost_anomaly_alert.
func (caa costAnomalyAlertAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(caa.ref.Append("name"))
}

func (caa costAnomalyAlertAttributes) Timeouts() costanomalyalert.TimeoutsAttributes {
	return terra.ReferenceAsSingle[costanomalyalert.TimeoutsAttributes](caa.ref.Append("timeouts"))
}

type costAnomalyAlertState struct {
	DisplayName    string                          `json:"display_name"`
	EmailAddresses []string                        `json:"email_addresses"`
	EmailSubject   string                          `json:"email_subject"`
	Id             string                          `json:"id"`
	Message        string                          `json:"message"`
	Name           string                          `json:"name"`
	Timeouts       *costanomalyalert.TimeoutsState `json:"timeouts"`
}
