// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	logicappintegrationaccount "github.com/golingon/terraproviders/azurerm/3.63.0/logicappintegrationaccount"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewLogicAppIntegrationAccount creates a new instance of [LogicAppIntegrationAccount].
func NewLogicAppIntegrationAccount(name string, args LogicAppIntegrationAccountArgs) *LogicAppIntegrationAccount {
	return &LogicAppIntegrationAccount{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*LogicAppIntegrationAccount)(nil)

// LogicAppIntegrationAccount represents the Terraform resource azurerm_logic_app_integration_account.
type LogicAppIntegrationAccount struct {
	Name      string
	Args      LogicAppIntegrationAccountArgs
	state     *logicAppIntegrationAccountState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [LogicAppIntegrationAccount].
func (laia *LogicAppIntegrationAccount) Type() string {
	return "azurerm_logic_app_integration_account"
}

// LocalName returns the local name for [LogicAppIntegrationAccount].
func (laia *LogicAppIntegrationAccount) LocalName() string {
	return laia.Name
}

// Configuration returns the configuration (args) for [LogicAppIntegrationAccount].
func (laia *LogicAppIntegrationAccount) Configuration() interface{} {
	return laia.Args
}

// DependOn is used for other resources to depend on [LogicAppIntegrationAccount].
func (laia *LogicAppIntegrationAccount) DependOn() terra.Reference {
	return terra.ReferenceResource(laia)
}

// Dependencies returns the list of resources [LogicAppIntegrationAccount] depends_on.
func (laia *LogicAppIntegrationAccount) Dependencies() terra.Dependencies {
	return laia.DependsOn
}

// LifecycleManagement returns the lifecycle block for [LogicAppIntegrationAccount].
func (laia *LogicAppIntegrationAccount) LifecycleManagement() *terra.Lifecycle {
	return laia.Lifecycle
}

// Attributes returns the attributes for [LogicAppIntegrationAccount].
func (laia *LogicAppIntegrationAccount) Attributes() logicAppIntegrationAccountAttributes {
	return logicAppIntegrationAccountAttributes{ref: terra.ReferenceResource(laia)}
}

// ImportState imports the given attribute values into [LogicAppIntegrationAccount]'s state.
func (laia *LogicAppIntegrationAccount) ImportState(av io.Reader) error {
	laia.state = &logicAppIntegrationAccountState{}
	if err := json.NewDecoder(av).Decode(laia.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", laia.Type(), laia.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [LogicAppIntegrationAccount] has state.
func (laia *LogicAppIntegrationAccount) State() (*logicAppIntegrationAccountState, bool) {
	return laia.state, laia.state != nil
}

// StateMust returns the state for [LogicAppIntegrationAccount]. Panics if the state is nil.
func (laia *LogicAppIntegrationAccount) StateMust() *logicAppIntegrationAccountState {
	if laia.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", laia.Type(), laia.LocalName()))
	}
	return laia.state
}

// LogicAppIntegrationAccountArgs contains the configurations for azurerm_logic_app_integration_account.
type LogicAppIntegrationAccountArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IntegrationServiceEnvironmentId: string, optional
	IntegrationServiceEnvironmentId terra.StringValue `hcl:"integration_service_environment_id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// SkuName: string, required
	SkuName terra.StringValue `hcl:"sku_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Timeouts: optional
	Timeouts *logicappintegrationaccount.Timeouts `hcl:"timeouts,block"`
}
type logicAppIntegrationAccountAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_logic_app_integration_account.
func (laia logicAppIntegrationAccountAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(laia.ref.Append("id"))
}

// IntegrationServiceEnvironmentId returns a reference to field integration_service_environment_id of azurerm_logic_app_integration_account.
func (laia logicAppIntegrationAccountAttributes) IntegrationServiceEnvironmentId() terra.StringValue {
	return terra.ReferenceAsString(laia.ref.Append("integration_service_environment_id"))
}

// Location returns a reference to field location of azurerm_logic_app_integration_account.
func (laia logicAppIntegrationAccountAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(laia.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_logic_app_integration_account.
func (laia logicAppIntegrationAccountAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(laia.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_logic_app_integration_account.
func (laia logicAppIntegrationAccountAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(laia.ref.Append("resource_group_name"))
}

// SkuName returns a reference to field sku_name of azurerm_logic_app_integration_account.
func (laia logicAppIntegrationAccountAttributes) SkuName() terra.StringValue {
	return terra.ReferenceAsString(laia.ref.Append("sku_name"))
}

// Tags returns a reference to field tags of azurerm_logic_app_integration_account.
func (laia logicAppIntegrationAccountAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](laia.ref.Append("tags"))
}

func (laia logicAppIntegrationAccountAttributes) Timeouts() logicappintegrationaccount.TimeoutsAttributes {
	return terra.ReferenceAsSingle[logicappintegrationaccount.TimeoutsAttributes](laia.ref.Append("timeouts"))
}

type logicAppIntegrationAccountState struct {
	Id                              string                                    `json:"id"`
	IntegrationServiceEnvironmentId string                                    `json:"integration_service_environment_id"`
	Location                        string                                    `json:"location"`
	Name                            string                                    `json:"name"`
	ResourceGroupName               string                                    `json:"resource_group_name"`
	SkuName                         string                                    `json:"sku_name"`
	Tags                            map[string]string                         `json:"tags"`
	Timeouts                        *logicappintegrationaccount.TimeoutsState `json:"timeouts"`
}
