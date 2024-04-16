// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_logic_app_integration_account_batch_configuration

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

// Resource represents the Terraform resource azurerm_logic_app_integration_account_batch_configuration.
type Resource struct {
	Name      string
	Args      Args
	state     *azurermLogicAppIntegrationAccountBatchConfigurationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (alaiabc *Resource) Type() string {
	return "azurerm_logic_app_integration_account_batch_configuration"
}

// LocalName returns the local name for [Resource].
func (alaiabc *Resource) LocalName() string {
	return alaiabc.Name
}

// Configuration returns the configuration (args) for [Resource].
func (alaiabc *Resource) Configuration() interface{} {
	return alaiabc.Args
}

// DependOn is used for other resources to depend on [Resource].
func (alaiabc *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(alaiabc)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (alaiabc *Resource) Dependencies() terra.Dependencies {
	return alaiabc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (alaiabc *Resource) LifecycleManagement() *terra.Lifecycle {
	return alaiabc.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (alaiabc *Resource) Attributes() azurermLogicAppIntegrationAccountBatchConfigurationAttributes {
	return azurermLogicAppIntegrationAccountBatchConfigurationAttributes{ref: terra.ReferenceResource(alaiabc)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (alaiabc *Resource) ImportState(state io.Reader) error {
	alaiabc.state = &azurermLogicAppIntegrationAccountBatchConfigurationState{}
	if err := json.NewDecoder(state).Decode(alaiabc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", alaiabc.Type(), alaiabc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (alaiabc *Resource) State() (*azurermLogicAppIntegrationAccountBatchConfigurationState, bool) {
	return alaiabc.state, alaiabc.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (alaiabc *Resource) StateMust() *azurermLogicAppIntegrationAccountBatchConfigurationState {
	if alaiabc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", alaiabc.Type(), alaiabc.LocalName()))
	}
	return alaiabc.state
}

// Args contains the configurations for azurerm_logic_app_integration_account_batch_configuration.
type Args struct {
	// BatchGroupName: string, required
	BatchGroupName terra.StringValue `hcl:"batch_group_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IntegrationAccountName: string, required
	IntegrationAccountName terra.StringValue `hcl:"integration_account_name,attr" validate:"required"`
	// Metadata: map of string, optional
	Metadata terra.MapValue[terra.StringValue] `hcl:"metadata,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// ReleaseCriteria: required
	ReleaseCriteria *ReleaseCriteria `hcl:"release_criteria,block" validate:"required"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type azurermLogicAppIntegrationAccountBatchConfigurationAttributes struct {
	ref terra.Reference
}

// BatchGroupName returns a reference to field batch_group_name of azurerm_logic_app_integration_account_batch_configuration.
func (alaiabc azurermLogicAppIntegrationAccountBatchConfigurationAttributes) BatchGroupName() terra.StringValue {
	return terra.ReferenceAsString(alaiabc.ref.Append("batch_group_name"))
}

// Id returns a reference to field id of azurerm_logic_app_integration_account_batch_configuration.
func (alaiabc azurermLogicAppIntegrationAccountBatchConfigurationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(alaiabc.ref.Append("id"))
}

// IntegrationAccountName returns a reference to field integration_account_name of azurerm_logic_app_integration_account_batch_configuration.
func (alaiabc azurermLogicAppIntegrationAccountBatchConfigurationAttributes) IntegrationAccountName() terra.StringValue {
	return terra.ReferenceAsString(alaiabc.ref.Append("integration_account_name"))
}

// Metadata returns a reference to field metadata of azurerm_logic_app_integration_account_batch_configuration.
func (alaiabc azurermLogicAppIntegrationAccountBatchConfigurationAttributes) Metadata() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](alaiabc.ref.Append("metadata"))
}

// Name returns a reference to field name of azurerm_logic_app_integration_account_batch_configuration.
func (alaiabc azurermLogicAppIntegrationAccountBatchConfigurationAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(alaiabc.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_logic_app_integration_account_batch_configuration.
func (alaiabc azurermLogicAppIntegrationAccountBatchConfigurationAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(alaiabc.ref.Append("resource_group_name"))
}

func (alaiabc azurermLogicAppIntegrationAccountBatchConfigurationAttributes) ReleaseCriteria() terra.ListValue[ReleaseCriteriaAttributes] {
	return terra.ReferenceAsList[ReleaseCriteriaAttributes](alaiabc.ref.Append("release_criteria"))
}

func (alaiabc azurermLogicAppIntegrationAccountBatchConfigurationAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](alaiabc.ref.Append("timeouts"))
}

type azurermLogicAppIntegrationAccountBatchConfigurationState struct {
	BatchGroupName         string                 `json:"batch_group_name"`
	Id                     string                 `json:"id"`
	IntegrationAccountName string                 `json:"integration_account_name"`
	Metadata               map[string]string      `json:"metadata"`
	Name                   string                 `json:"name"`
	ResourceGroupName      string                 `json:"resource_group_name"`
	ReleaseCriteria        []ReleaseCriteriaState `json:"release_criteria"`
	Timeouts               *TimeoutsState         `json:"timeouts"`
}
