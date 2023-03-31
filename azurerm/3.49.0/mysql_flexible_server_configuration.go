// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	mysqlflexibleserverconfiguration "github.com/golingon/terraproviders/azurerm/3.49.0/mysqlflexibleserverconfiguration"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewMysqlFlexibleServerConfiguration(name string, args MysqlFlexibleServerConfigurationArgs) *MysqlFlexibleServerConfiguration {
	return &MysqlFlexibleServerConfiguration{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*MysqlFlexibleServerConfiguration)(nil)

type MysqlFlexibleServerConfiguration struct {
	Name  string
	Args  MysqlFlexibleServerConfigurationArgs
	state *mysqlFlexibleServerConfigurationState
}

func (mfsc *MysqlFlexibleServerConfiguration) Type() string {
	return "azurerm_mysql_flexible_server_configuration"
}

func (mfsc *MysqlFlexibleServerConfiguration) LocalName() string {
	return mfsc.Name
}

func (mfsc *MysqlFlexibleServerConfiguration) Configuration() interface{} {
	return mfsc.Args
}

func (mfsc *MysqlFlexibleServerConfiguration) Attributes() mysqlFlexibleServerConfigurationAttributes {
	return mysqlFlexibleServerConfigurationAttributes{ref: terra.ReferenceResource(mfsc)}
}

func (mfsc *MysqlFlexibleServerConfiguration) ImportState(av io.Reader) error {
	mfsc.state = &mysqlFlexibleServerConfigurationState{}
	if err := json.NewDecoder(av).Decode(mfsc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", mfsc.Type(), mfsc.LocalName(), err)
	}
	return nil
}

func (mfsc *MysqlFlexibleServerConfiguration) State() (*mysqlFlexibleServerConfigurationState, bool) {
	return mfsc.state, mfsc.state != nil
}

func (mfsc *MysqlFlexibleServerConfiguration) StateMust() *mysqlFlexibleServerConfigurationState {
	if mfsc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", mfsc.Type(), mfsc.LocalName()))
	}
	return mfsc.state
}

func (mfsc *MysqlFlexibleServerConfiguration) DependOn() terra.Reference {
	return terra.ReferenceResource(mfsc)
}

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
	// DependsOn contains resources that MysqlFlexibleServerConfiguration depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type mysqlFlexibleServerConfigurationAttributes struct {
	ref terra.Reference
}

func (mfsc mysqlFlexibleServerConfigurationAttributes) Id() terra.StringValue {
	return terra.ReferenceString(mfsc.ref.Append("id"))
}

func (mfsc mysqlFlexibleServerConfigurationAttributes) Name() terra.StringValue {
	return terra.ReferenceString(mfsc.ref.Append("name"))
}

func (mfsc mysqlFlexibleServerConfigurationAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceString(mfsc.ref.Append("resource_group_name"))
}

func (mfsc mysqlFlexibleServerConfigurationAttributes) ServerName() terra.StringValue {
	return terra.ReferenceString(mfsc.ref.Append("server_name"))
}

func (mfsc mysqlFlexibleServerConfigurationAttributes) Value() terra.StringValue {
	return terra.ReferenceString(mfsc.ref.Append("value"))
}

func (mfsc mysqlFlexibleServerConfigurationAttributes) Timeouts() mysqlflexibleserverconfiguration.TimeoutsAttributes {
	return terra.ReferenceSingle[mysqlflexibleserverconfiguration.TimeoutsAttributes](mfsc.ref.Append("timeouts"))
}

type mysqlFlexibleServerConfigurationState struct {
	Id                string                                          `json:"id"`
	Name              string                                          `json:"name"`
	ResourceGroupName string                                          `json:"resource_group_name"`
	ServerName        string                                          `json:"server_name"`
	Value             string                                          `json:"value"`
	Timeouts          *mysqlflexibleserverconfiguration.TimeoutsState `json:"timeouts"`
}
