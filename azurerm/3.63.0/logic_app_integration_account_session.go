// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	logicappintegrationaccountsession "github.com/golingon/terraproviders/azurerm/3.63.0/logicappintegrationaccountsession"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewLogicAppIntegrationAccountSession creates a new instance of [LogicAppIntegrationAccountSession].
func NewLogicAppIntegrationAccountSession(name string, args LogicAppIntegrationAccountSessionArgs) *LogicAppIntegrationAccountSession {
	return &LogicAppIntegrationAccountSession{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*LogicAppIntegrationAccountSession)(nil)

// LogicAppIntegrationAccountSession represents the Terraform resource azurerm_logic_app_integration_account_session.
type LogicAppIntegrationAccountSession struct {
	Name      string
	Args      LogicAppIntegrationAccountSessionArgs
	state     *logicAppIntegrationAccountSessionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [LogicAppIntegrationAccountSession].
func (laias *LogicAppIntegrationAccountSession) Type() string {
	return "azurerm_logic_app_integration_account_session"
}

// LocalName returns the local name for [LogicAppIntegrationAccountSession].
func (laias *LogicAppIntegrationAccountSession) LocalName() string {
	return laias.Name
}

// Configuration returns the configuration (args) for [LogicAppIntegrationAccountSession].
func (laias *LogicAppIntegrationAccountSession) Configuration() interface{} {
	return laias.Args
}

// DependOn is used for other resources to depend on [LogicAppIntegrationAccountSession].
func (laias *LogicAppIntegrationAccountSession) DependOn() terra.Reference {
	return terra.ReferenceResource(laias)
}

// Dependencies returns the list of resources [LogicAppIntegrationAccountSession] depends_on.
func (laias *LogicAppIntegrationAccountSession) Dependencies() terra.Dependencies {
	return laias.DependsOn
}

// LifecycleManagement returns the lifecycle block for [LogicAppIntegrationAccountSession].
func (laias *LogicAppIntegrationAccountSession) LifecycleManagement() *terra.Lifecycle {
	return laias.Lifecycle
}

// Attributes returns the attributes for [LogicAppIntegrationAccountSession].
func (laias *LogicAppIntegrationAccountSession) Attributes() logicAppIntegrationAccountSessionAttributes {
	return logicAppIntegrationAccountSessionAttributes{ref: terra.ReferenceResource(laias)}
}

// ImportState imports the given attribute values into [LogicAppIntegrationAccountSession]'s state.
func (laias *LogicAppIntegrationAccountSession) ImportState(av io.Reader) error {
	laias.state = &logicAppIntegrationAccountSessionState{}
	if err := json.NewDecoder(av).Decode(laias.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", laias.Type(), laias.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [LogicAppIntegrationAccountSession] has state.
func (laias *LogicAppIntegrationAccountSession) State() (*logicAppIntegrationAccountSessionState, bool) {
	return laias.state, laias.state != nil
}

// StateMust returns the state for [LogicAppIntegrationAccountSession]. Panics if the state is nil.
func (laias *LogicAppIntegrationAccountSession) StateMust() *logicAppIntegrationAccountSessionState {
	if laias.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", laias.Type(), laias.LocalName()))
	}
	return laias.state
}

// LogicAppIntegrationAccountSessionArgs contains the configurations for azurerm_logic_app_integration_account_session.
type LogicAppIntegrationAccountSessionArgs struct {
	// Content: string, required
	Content terra.StringValue `hcl:"content,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IntegrationAccountName: string, required
	IntegrationAccountName terra.StringValue `hcl:"integration_account_name,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *logicappintegrationaccountsession.Timeouts `hcl:"timeouts,block"`
}
type logicAppIntegrationAccountSessionAttributes struct {
	ref terra.Reference
}

// Content returns a reference to field content of azurerm_logic_app_integration_account_session.
func (laias logicAppIntegrationAccountSessionAttributes) Content() terra.StringValue {
	return terra.ReferenceAsString(laias.ref.Append("content"))
}

// Id returns a reference to field id of azurerm_logic_app_integration_account_session.
func (laias logicAppIntegrationAccountSessionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(laias.ref.Append("id"))
}

// IntegrationAccountName returns a reference to field integration_account_name of azurerm_logic_app_integration_account_session.
func (laias logicAppIntegrationAccountSessionAttributes) IntegrationAccountName() terra.StringValue {
	return terra.ReferenceAsString(laias.ref.Append("integration_account_name"))
}

// Name returns a reference to field name of azurerm_logic_app_integration_account_session.
func (laias logicAppIntegrationAccountSessionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(laias.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_logic_app_integration_account_session.
func (laias logicAppIntegrationAccountSessionAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(laias.ref.Append("resource_group_name"))
}

func (laias logicAppIntegrationAccountSessionAttributes) Timeouts() logicappintegrationaccountsession.TimeoutsAttributes {
	return terra.ReferenceAsSingle[logicappintegrationaccountsession.TimeoutsAttributes](laias.ref.Append("timeouts"))
}

type logicAppIntegrationAccountSessionState struct {
	Content                string                                           `json:"content"`
	Id                     string                                           `json:"id"`
	IntegrationAccountName string                                           `json:"integration_account_name"`
	Name                   string                                           `json:"name"`
	ResourceGroupName      string                                           `json:"resource_group_name"`
	Timeouts               *logicappintegrationaccountsession.TimeoutsState `json:"timeouts"`
}
