// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	logicappintegrationaccountbatchconfiguration "github.com/golingon/terraproviders/azurerm/3.63.0/logicappintegrationaccountbatchconfiguration"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewLogicAppIntegrationAccountBatchConfiguration creates a new instance of [LogicAppIntegrationAccountBatchConfiguration].
func NewLogicAppIntegrationAccountBatchConfiguration(name string, args LogicAppIntegrationAccountBatchConfigurationArgs) *LogicAppIntegrationAccountBatchConfiguration {
	return &LogicAppIntegrationAccountBatchConfiguration{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*LogicAppIntegrationAccountBatchConfiguration)(nil)

// LogicAppIntegrationAccountBatchConfiguration represents the Terraform resource azurerm_logic_app_integration_account_batch_configuration.
type LogicAppIntegrationAccountBatchConfiguration struct {
	Name      string
	Args      LogicAppIntegrationAccountBatchConfigurationArgs
	state     *logicAppIntegrationAccountBatchConfigurationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [LogicAppIntegrationAccountBatchConfiguration].
func (laiabc *LogicAppIntegrationAccountBatchConfiguration) Type() string {
	return "azurerm_logic_app_integration_account_batch_configuration"
}

// LocalName returns the local name for [LogicAppIntegrationAccountBatchConfiguration].
func (laiabc *LogicAppIntegrationAccountBatchConfiguration) LocalName() string {
	return laiabc.Name
}

// Configuration returns the configuration (args) for [LogicAppIntegrationAccountBatchConfiguration].
func (laiabc *LogicAppIntegrationAccountBatchConfiguration) Configuration() interface{} {
	return laiabc.Args
}

// DependOn is used for other resources to depend on [LogicAppIntegrationAccountBatchConfiguration].
func (laiabc *LogicAppIntegrationAccountBatchConfiguration) DependOn() terra.Reference {
	return terra.ReferenceResource(laiabc)
}

// Dependencies returns the list of resources [LogicAppIntegrationAccountBatchConfiguration] depends_on.
func (laiabc *LogicAppIntegrationAccountBatchConfiguration) Dependencies() terra.Dependencies {
	return laiabc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [LogicAppIntegrationAccountBatchConfiguration].
func (laiabc *LogicAppIntegrationAccountBatchConfiguration) LifecycleManagement() *terra.Lifecycle {
	return laiabc.Lifecycle
}

// Attributes returns the attributes for [LogicAppIntegrationAccountBatchConfiguration].
func (laiabc *LogicAppIntegrationAccountBatchConfiguration) Attributes() logicAppIntegrationAccountBatchConfigurationAttributes {
	return logicAppIntegrationAccountBatchConfigurationAttributes{ref: terra.ReferenceResource(laiabc)}
}

// ImportState imports the given attribute values into [LogicAppIntegrationAccountBatchConfiguration]'s state.
func (laiabc *LogicAppIntegrationAccountBatchConfiguration) ImportState(av io.Reader) error {
	laiabc.state = &logicAppIntegrationAccountBatchConfigurationState{}
	if err := json.NewDecoder(av).Decode(laiabc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", laiabc.Type(), laiabc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [LogicAppIntegrationAccountBatchConfiguration] has state.
func (laiabc *LogicAppIntegrationAccountBatchConfiguration) State() (*logicAppIntegrationAccountBatchConfigurationState, bool) {
	return laiabc.state, laiabc.state != nil
}

// StateMust returns the state for [LogicAppIntegrationAccountBatchConfiguration]. Panics if the state is nil.
func (laiabc *LogicAppIntegrationAccountBatchConfiguration) StateMust() *logicAppIntegrationAccountBatchConfigurationState {
	if laiabc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", laiabc.Type(), laiabc.LocalName()))
	}
	return laiabc.state
}

// LogicAppIntegrationAccountBatchConfigurationArgs contains the configurations for azurerm_logic_app_integration_account_batch_configuration.
type LogicAppIntegrationAccountBatchConfigurationArgs struct {
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
	ReleaseCriteria *logicappintegrationaccountbatchconfiguration.ReleaseCriteria `hcl:"release_criteria,block" validate:"required"`
	// Timeouts: optional
	Timeouts *logicappintegrationaccountbatchconfiguration.Timeouts `hcl:"timeouts,block"`
}
type logicAppIntegrationAccountBatchConfigurationAttributes struct {
	ref terra.Reference
}

// BatchGroupName returns a reference to field batch_group_name of azurerm_logic_app_integration_account_batch_configuration.
func (laiabc logicAppIntegrationAccountBatchConfigurationAttributes) BatchGroupName() terra.StringValue {
	return terra.ReferenceAsString(laiabc.ref.Append("batch_group_name"))
}

// Id returns a reference to field id of azurerm_logic_app_integration_account_batch_configuration.
func (laiabc logicAppIntegrationAccountBatchConfigurationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(laiabc.ref.Append("id"))
}

// IntegrationAccountName returns a reference to field integration_account_name of azurerm_logic_app_integration_account_batch_configuration.
func (laiabc logicAppIntegrationAccountBatchConfigurationAttributes) IntegrationAccountName() terra.StringValue {
	return terra.ReferenceAsString(laiabc.ref.Append("integration_account_name"))
}

// Metadata returns a reference to field metadata of azurerm_logic_app_integration_account_batch_configuration.
func (laiabc logicAppIntegrationAccountBatchConfigurationAttributes) Metadata() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](laiabc.ref.Append("metadata"))
}

// Name returns a reference to field name of azurerm_logic_app_integration_account_batch_configuration.
func (laiabc logicAppIntegrationAccountBatchConfigurationAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(laiabc.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_logic_app_integration_account_batch_configuration.
func (laiabc logicAppIntegrationAccountBatchConfigurationAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(laiabc.ref.Append("resource_group_name"))
}

func (laiabc logicAppIntegrationAccountBatchConfigurationAttributes) ReleaseCriteria() terra.ListValue[logicappintegrationaccountbatchconfiguration.ReleaseCriteriaAttributes] {
	return terra.ReferenceAsList[logicappintegrationaccountbatchconfiguration.ReleaseCriteriaAttributes](laiabc.ref.Append("release_criteria"))
}

func (laiabc logicAppIntegrationAccountBatchConfigurationAttributes) Timeouts() logicappintegrationaccountbatchconfiguration.TimeoutsAttributes {
	return terra.ReferenceAsSingle[logicappintegrationaccountbatchconfiguration.TimeoutsAttributes](laiabc.ref.Append("timeouts"))
}

type logicAppIntegrationAccountBatchConfigurationState struct {
	BatchGroupName         string                                                              `json:"batch_group_name"`
	Id                     string                                                              `json:"id"`
	IntegrationAccountName string                                                              `json:"integration_account_name"`
	Metadata               map[string]string                                                   `json:"metadata"`
	Name                   string                                                              `json:"name"`
	ResourceGroupName      string                                                              `json:"resource_group_name"`
	ReleaseCriteria        []logicappintegrationaccountbatchconfiguration.ReleaseCriteriaState `json:"release_criteria"`
	Timeouts               *logicappintegrationaccountbatchconfiguration.TimeoutsState         `json:"timeouts"`
}
