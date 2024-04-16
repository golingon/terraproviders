// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_postgresql_flexible_server_configuration

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

// Resource represents the Terraform resource azurerm_postgresql_flexible_server_configuration.
type Resource struct {
	Name      string
	Args      Args
	state     *azurermPostgresqlFlexibleServerConfigurationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (apfsc *Resource) Type() string {
	return "azurerm_postgresql_flexible_server_configuration"
}

// LocalName returns the local name for [Resource].
func (apfsc *Resource) LocalName() string {
	return apfsc.Name
}

// Configuration returns the configuration (args) for [Resource].
func (apfsc *Resource) Configuration() interface{} {
	return apfsc.Args
}

// DependOn is used for other resources to depend on [Resource].
func (apfsc *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(apfsc)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (apfsc *Resource) Dependencies() terra.Dependencies {
	return apfsc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (apfsc *Resource) LifecycleManagement() *terra.Lifecycle {
	return apfsc.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (apfsc *Resource) Attributes() azurermPostgresqlFlexibleServerConfigurationAttributes {
	return azurermPostgresqlFlexibleServerConfigurationAttributes{ref: terra.ReferenceResource(apfsc)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (apfsc *Resource) ImportState(state io.Reader) error {
	apfsc.state = &azurermPostgresqlFlexibleServerConfigurationState{}
	if err := json.NewDecoder(state).Decode(apfsc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", apfsc.Type(), apfsc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (apfsc *Resource) State() (*azurermPostgresqlFlexibleServerConfigurationState, bool) {
	return apfsc.state, apfsc.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (apfsc *Resource) StateMust() *azurermPostgresqlFlexibleServerConfigurationState {
	if apfsc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", apfsc.Type(), apfsc.LocalName()))
	}
	return apfsc.state
}

// Args contains the configurations for azurerm_postgresql_flexible_server_configuration.
type Args struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ServerId: string, required
	ServerId terra.StringValue `hcl:"server_id,attr" validate:"required"`
	// Value: string, required
	Value terra.StringValue `hcl:"value,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type azurermPostgresqlFlexibleServerConfigurationAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_postgresql_flexible_server_configuration.
func (apfsc azurermPostgresqlFlexibleServerConfigurationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(apfsc.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_postgresql_flexible_server_configuration.
func (apfsc azurermPostgresqlFlexibleServerConfigurationAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(apfsc.ref.Append("name"))
}

// ServerId returns a reference to field server_id of azurerm_postgresql_flexible_server_configuration.
func (apfsc azurermPostgresqlFlexibleServerConfigurationAttributes) ServerId() terra.StringValue {
	return terra.ReferenceAsString(apfsc.ref.Append("server_id"))
}

// Value returns a reference to field value of azurerm_postgresql_flexible_server_configuration.
func (apfsc azurermPostgresqlFlexibleServerConfigurationAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(apfsc.ref.Append("value"))
}

func (apfsc azurermPostgresqlFlexibleServerConfigurationAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](apfsc.ref.Append("timeouts"))
}

type azurermPostgresqlFlexibleServerConfigurationState struct {
	Id       string         `json:"id"`
	Name     string         `json:"name"`
	ServerId string         `json:"server_id"`
	Value    string         `json:"value"`
	Timeouts *TimeoutsState `json:"timeouts"`
}
