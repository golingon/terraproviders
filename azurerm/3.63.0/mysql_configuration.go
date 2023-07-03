// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	mysqlconfiguration "github.com/golingon/terraproviders/azurerm/3.63.0/mysqlconfiguration"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewMysqlConfiguration creates a new instance of [MysqlConfiguration].
func NewMysqlConfiguration(name string, args MysqlConfigurationArgs) *MysqlConfiguration {
	return &MysqlConfiguration{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*MysqlConfiguration)(nil)

// MysqlConfiguration represents the Terraform resource azurerm_mysql_configuration.
type MysqlConfiguration struct {
	Name      string
	Args      MysqlConfigurationArgs
	state     *mysqlConfigurationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [MysqlConfiguration].
func (mc *MysqlConfiguration) Type() string {
	return "azurerm_mysql_configuration"
}

// LocalName returns the local name for [MysqlConfiguration].
func (mc *MysqlConfiguration) LocalName() string {
	return mc.Name
}

// Configuration returns the configuration (args) for [MysqlConfiguration].
func (mc *MysqlConfiguration) Configuration() interface{} {
	return mc.Args
}

// DependOn is used for other resources to depend on [MysqlConfiguration].
func (mc *MysqlConfiguration) DependOn() terra.Reference {
	return terra.ReferenceResource(mc)
}

// Dependencies returns the list of resources [MysqlConfiguration] depends_on.
func (mc *MysqlConfiguration) Dependencies() terra.Dependencies {
	return mc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [MysqlConfiguration].
func (mc *MysqlConfiguration) LifecycleManagement() *terra.Lifecycle {
	return mc.Lifecycle
}

// Attributes returns the attributes for [MysqlConfiguration].
func (mc *MysqlConfiguration) Attributes() mysqlConfigurationAttributes {
	return mysqlConfigurationAttributes{ref: terra.ReferenceResource(mc)}
}

// ImportState imports the given attribute values into [MysqlConfiguration]'s state.
func (mc *MysqlConfiguration) ImportState(av io.Reader) error {
	mc.state = &mysqlConfigurationState{}
	if err := json.NewDecoder(av).Decode(mc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", mc.Type(), mc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [MysqlConfiguration] has state.
func (mc *MysqlConfiguration) State() (*mysqlConfigurationState, bool) {
	return mc.state, mc.state != nil
}

// StateMust returns the state for [MysqlConfiguration]. Panics if the state is nil.
func (mc *MysqlConfiguration) StateMust() *mysqlConfigurationState {
	if mc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", mc.Type(), mc.LocalName()))
	}
	return mc.state
}

// MysqlConfigurationArgs contains the configurations for azurerm_mysql_configuration.
type MysqlConfigurationArgs struct {
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
	Timeouts *mysqlconfiguration.Timeouts `hcl:"timeouts,block"`
}
type mysqlConfigurationAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_mysql_configuration.
func (mc mysqlConfigurationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(mc.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_mysql_configuration.
func (mc mysqlConfigurationAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(mc.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_mysql_configuration.
func (mc mysqlConfigurationAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(mc.ref.Append("resource_group_name"))
}

// ServerName returns a reference to field server_name of azurerm_mysql_configuration.
func (mc mysqlConfigurationAttributes) ServerName() terra.StringValue {
	return terra.ReferenceAsString(mc.ref.Append("server_name"))
}

// Value returns a reference to field value of azurerm_mysql_configuration.
func (mc mysqlConfigurationAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(mc.ref.Append("value"))
}

func (mc mysqlConfigurationAttributes) Timeouts() mysqlconfiguration.TimeoutsAttributes {
	return terra.ReferenceAsSingle[mysqlconfiguration.TimeoutsAttributes](mc.ref.Append("timeouts"))
}

type mysqlConfigurationState struct {
	Id                string                            `json:"id"`
	Name              string                            `json:"name"`
	ResourceGroupName string                            `json:"resource_group_name"`
	ServerName        string                            `json:"server_name"`
	Value             string                            `json:"value"`
	Timeouts          *mysqlconfiguration.TimeoutsState `json:"timeouts"`
}
