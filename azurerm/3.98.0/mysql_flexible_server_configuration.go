// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	mysqlflexibleserverconfiguration "github.com/golingon/terraproviders/azurerm/3.98.0/mysqlflexibleserverconfiguration"
	"io"
)

// NewMysqlFlexibleServerConfiguration creates a new instance of [MysqlFlexibleServerConfiguration].
func NewMysqlFlexibleServerConfiguration(name string, args MysqlFlexibleServerConfigurationArgs) *MysqlFlexibleServerConfiguration {
	return &MysqlFlexibleServerConfiguration{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*MysqlFlexibleServerConfiguration)(nil)

// MysqlFlexibleServerConfiguration represents the Terraform resource azurerm_mysql_flexible_server_configuration.
type MysqlFlexibleServerConfiguration struct {
	Name      string
	Args      MysqlFlexibleServerConfigurationArgs
	state     *mysqlFlexibleServerConfigurationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [MysqlFlexibleServerConfiguration].
func (mfsc *MysqlFlexibleServerConfiguration) Type() string {
	return "azurerm_mysql_flexible_server_configuration"
}

// LocalName returns the local name for [MysqlFlexibleServerConfiguration].
func (mfsc *MysqlFlexibleServerConfiguration) LocalName() string {
	return mfsc.Name
}

// Configuration returns the configuration (args) for [MysqlFlexibleServerConfiguration].
func (mfsc *MysqlFlexibleServerConfiguration) Configuration() interface{} {
	return mfsc.Args
}

// DependOn is used for other resources to depend on [MysqlFlexibleServerConfiguration].
func (mfsc *MysqlFlexibleServerConfiguration) DependOn() terra.Reference {
	return terra.ReferenceResource(mfsc)
}

// Dependencies returns the list of resources [MysqlFlexibleServerConfiguration] depends_on.
func (mfsc *MysqlFlexibleServerConfiguration) Dependencies() terra.Dependencies {
	return mfsc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [MysqlFlexibleServerConfiguration].
func (mfsc *MysqlFlexibleServerConfiguration) LifecycleManagement() *terra.Lifecycle {
	return mfsc.Lifecycle
}

// Attributes returns the attributes for [MysqlFlexibleServerConfiguration].
func (mfsc *MysqlFlexibleServerConfiguration) Attributes() mysqlFlexibleServerConfigurationAttributes {
	return mysqlFlexibleServerConfigurationAttributes{ref: terra.ReferenceResource(mfsc)}
}

// ImportState imports the given attribute values into [MysqlFlexibleServerConfiguration]'s state.
func (mfsc *MysqlFlexibleServerConfiguration) ImportState(av io.Reader) error {
	mfsc.state = &mysqlFlexibleServerConfigurationState{}
	if err := json.NewDecoder(av).Decode(mfsc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", mfsc.Type(), mfsc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [MysqlFlexibleServerConfiguration] has state.
func (mfsc *MysqlFlexibleServerConfiguration) State() (*mysqlFlexibleServerConfigurationState, bool) {
	return mfsc.state, mfsc.state != nil
}

// StateMust returns the state for [MysqlFlexibleServerConfiguration]. Panics if the state is nil.
func (mfsc *MysqlFlexibleServerConfiguration) StateMust() *mysqlFlexibleServerConfigurationState {
	if mfsc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", mfsc.Type(), mfsc.LocalName()))
	}
	return mfsc.state
}

// MysqlFlexibleServerConfigurationArgs contains the configurations for azurerm_mysql_flexible_server_configuration.
type MysqlFlexibleServerConfigurationArgs struct {
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
	Timeouts *mysqlflexibleserverconfiguration.Timeouts `hcl:"timeouts,block"`
}
type mysqlFlexibleServerConfigurationAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_mysql_flexible_server_configuration.
func (mfsc mysqlFlexibleServerConfigurationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(mfsc.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_mysql_flexible_server_configuration.
func (mfsc mysqlFlexibleServerConfigurationAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(mfsc.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_mysql_flexible_server_configuration.
func (mfsc mysqlFlexibleServerConfigurationAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(mfsc.ref.Append("resource_group_name"))
}

// ServerName returns a reference to field server_name of azurerm_mysql_flexible_server_configuration.
func (mfsc mysqlFlexibleServerConfigurationAttributes) ServerName() terra.StringValue {
	return terra.ReferenceAsString(mfsc.ref.Append("server_name"))
}

// Value returns a reference to field value of azurerm_mysql_flexible_server_configuration.
func (mfsc mysqlFlexibleServerConfigurationAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(mfsc.ref.Append("value"))
}

func (mfsc mysqlFlexibleServerConfigurationAttributes) Timeouts() mysqlflexibleserverconfiguration.TimeoutsAttributes {
	return terra.ReferenceAsSingle[mysqlflexibleserverconfiguration.TimeoutsAttributes](mfsc.ref.Append("timeouts"))
}

type mysqlFlexibleServerConfigurationState struct {
	Id                string                                          `json:"id"`
	Name              string                                          `json:"name"`
	ResourceGroupName string                                          `json:"resource_group_name"`
	ServerName        string                                          `json:"server_name"`
	Value             string                                          `json:"value"`
	Timeouts          *mysqlflexibleserverconfiguration.TimeoutsState `json:"timeouts"`
}
