// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	mariadbconfiguration "github.com/golingon/terraproviders/azurerm/3.98.0/mariadbconfiguration"
	"io"
)

// NewMariadbConfiguration creates a new instance of [MariadbConfiguration].
func NewMariadbConfiguration(name string, args MariadbConfigurationArgs) *MariadbConfiguration {
	return &MariadbConfiguration{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*MariadbConfiguration)(nil)

// MariadbConfiguration represents the Terraform resource azurerm_mariadb_configuration.
type MariadbConfiguration struct {
	Name      string
	Args      MariadbConfigurationArgs
	state     *mariadbConfigurationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [MariadbConfiguration].
func (mc *MariadbConfiguration) Type() string {
	return "azurerm_mariadb_configuration"
}

// LocalName returns the local name for [MariadbConfiguration].
func (mc *MariadbConfiguration) LocalName() string {
	return mc.Name
}

// Configuration returns the configuration (args) for [MariadbConfiguration].
func (mc *MariadbConfiguration) Configuration() interface{} {
	return mc.Args
}

// DependOn is used for other resources to depend on [MariadbConfiguration].
func (mc *MariadbConfiguration) DependOn() terra.Reference {
	return terra.ReferenceResource(mc)
}

// Dependencies returns the list of resources [MariadbConfiguration] depends_on.
func (mc *MariadbConfiguration) Dependencies() terra.Dependencies {
	return mc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [MariadbConfiguration].
func (mc *MariadbConfiguration) LifecycleManagement() *terra.Lifecycle {
	return mc.Lifecycle
}

// Attributes returns the attributes for [MariadbConfiguration].
func (mc *MariadbConfiguration) Attributes() mariadbConfigurationAttributes {
	return mariadbConfigurationAttributes{ref: terra.ReferenceResource(mc)}
}

// ImportState imports the given attribute values into [MariadbConfiguration]'s state.
func (mc *MariadbConfiguration) ImportState(av io.Reader) error {
	mc.state = &mariadbConfigurationState{}
	if err := json.NewDecoder(av).Decode(mc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", mc.Type(), mc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [MariadbConfiguration] has state.
func (mc *MariadbConfiguration) State() (*mariadbConfigurationState, bool) {
	return mc.state, mc.state != nil
}

// StateMust returns the state for [MariadbConfiguration]. Panics if the state is nil.
func (mc *MariadbConfiguration) StateMust() *mariadbConfigurationState {
	if mc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", mc.Type(), mc.LocalName()))
	}
	return mc.state
}

// MariadbConfigurationArgs contains the configurations for azurerm_mariadb_configuration.
type MariadbConfigurationArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// ServerName: string, required
	ServerName terra.StringValue `hcl:"server_name,attr" validate:"required"`
	// Value: string, required
	Value terra.StringValue `hcl:"value,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *mariadbconfiguration.Timeouts `hcl:"timeouts,block"`
}
type mariadbConfigurationAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_mariadb_configuration.
func (mc mariadbConfigurationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(mc.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_mariadb_configuration.
func (mc mariadbConfigurationAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(mc.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_mariadb_configuration.
func (mc mariadbConfigurationAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(mc.ref.Append("resource_group_name"))
}

// ServerName returns a reference to field server_name of azurerm_mariadb_configuration.
func (mc mariadbConfigurationAttributes) ServerName() terra.StringValue {
	return terra.ReferenceAsString(mc.ref.Append("server_name"))
}

// Value returns a reference to field value of azurerm_mariadb_configuration.
func (mc mariadbConfigurationAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(mc.ref.Append("value"))
}

func (mc mariadbConfigurationAttributes) Timeouts() mariadbconfiguration.TimeoutsAttributes {
	return terra.ReferenceAsSingle[mariadbconfiguration.TimeoutsAttributes](mc.ref.Append("timeouts"))
}

type mariadbConfigurationState struct {
	Id                string                              `json:"id"`
	Name              string                              `json:"name"`
	ResourceGroupName string                              `json:"resource_group_name"`
	ServerName        string                              `json:"server_name"`
	Value             string                              `json:"value"`
	Timeouts          *mariadbconfiguration.TimeoutsState `json:"timeouts"`
}
