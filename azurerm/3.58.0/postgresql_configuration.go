// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	postgresqlconfiguration "github.com/golingon/terraproviders/azurerm/3.58.0/postgresqlconfiguration"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewPostgresqlConfiguration creates a new instance of [PostgresqlConfiguration].
func NewPostgresqlConfiguration(name string, args PostgresqlConfigurationArgs) *PostgresqlConfiguration {
	return &PostgresqlConfiguration{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*PostgresqlConfiguration)(nil)

// PostgresqlConfiguration represents the Terraform resource azurerm_postgresql_configuration.
type PostgresqlConfiguration struct {
	Name      string
	Args      PostgresqlConfigurationArgs
	state     *postgresqlConfigurationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [PostgresqlConfiguration].
func (pc *PostgresqlConfiguration) Type() string {
	return "azurerm_postgresql_configuration"
}

// LocalName returns the local name for [PostgresqlConfiguration].
func (pc *PostgresqlConfiguration) LocalName() string {
	return pc.Name
}

// Configuration returns the configuration (args) for [PostgresqlConfiguration].
func (pc *PostgresqlConfiguration) Configuration() interface{} {
	return pc.Args
}

// DependOn is used for other resources to depend on [PostgresqlConfiguration].
func (pc *PostgresqlConfiguration) DependOn() terra.Reference {
	return terra.ReferenceResource(pc)
}

// Dependencies returns the list of resources [PostgresqlConfiguration] depends_on.
func (pc *PostgresqlConfiguration) Dependencies() terra.Dependencies {
	return pc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [PostgresqlConfiguration].
func (pc *PostgresqlConfiguration) LifecycleManagement() *terra.Lifecycle {
	return pc.Lifecycle
}

// Attributes returns the attributes for [PostgresqlConfiguration].
func (pc *PostgresqlConfiguration) Attributes() postgresqlConfigurationAttributes {
	return postgresqlConfigurationAttributes{ref: terra.ReferenceResource(pc)}
}

// ImportState imports the given attribute values into [PostgresqlConfiguration]'s state.
func (pc *PostgresqlConfiguration) ImportState(av io.Reader) error {
	pc.state = &postgresqlConfigurationState{}
	if err := json.NewDecoder(av).Decode(pc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", pc.Type(), pc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [PostgresqlConfiguration] has state.
func (pc *PostgresqlConfiguration) State() (*postgresqlConfigurationState, bool) {
	return pc.state, pc.state != nil
}

// StateMust returns the state for [PostgresqlConfiguration]. Panics if the state is nil.
func (pc *PostgresqlConfiguration) StateMust() *postgresqlConfigurationState {
	if pc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", pc.Type(), pc.LocalName()))
	}
	return pc.state
}

// PostgresqlConfigurationArgs contains the configurations for azurerm_postgresql_configuration.
type PostgresqlConfigurationArgs struct {
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
	Timeouts *postgresqlconfiguration.Timeouts `hcl:"timeouts,block"`
}
type postgresqlConfigurationAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_postgresql_configuration.
func (pc postgresqlConfigurationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(pc.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_postgresql_configuration.
func (pc postgresqlConfigurationAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(pc.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_postgresql_configuration.
func (pc postgresqlConfigurationAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(pc.ref.Append("resource_group_name"))
}

// ServerName returns a reference to field server_name of azurerm_postgresql_configuration.
func (pc postgresqlConfigurationAttributes) ServerName() terra.StringValue {
	return terra.ReferenceAsString(pc.ref.Append("server_name"))
}

// Value returns a reference to field value of azurerm_postgresql_configuration.
func (pc postgresqlConfigurationAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(pc.ref.Append("value"))
}

func (pc postgresqlConfigurationAttributes) Timeouts() postgresqlconfiguration.TimeoutsAttributes {
	return terra.ReferenceAsSingle[postgresqlconfiguration.TimeoutsAttributes](pc.ref.Append("timeouts"))
}

type postgresqlConfigurationState struct {
	Id                string                                 `json:"id"`
	Name              string                                 `json:"name"`
	ResourceGroupName string                                 `json:"resource_group_name"`
	ServerName        string                                 `json:"server_name"`
	Value             string                                 `json:"value"`
	Timeouts          *postgresqlconfiguration.TimeoutsState `json:"timeouts"`
}
