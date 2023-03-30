// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewConfigConfigurationRecorderStatus(name string, args ConfigConfigurationRecorderStatusArgs) *ConfigConfigurationRecorderStatus {
	return &ConfigConfigurationRecorderStatus{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ConfigConfigurationRecorderStatus)(nil)

type ConfigConfigurationRecorderStatus struct {
	Name  string
	Args  ConfigConfigurationRecorderStatusArgs
	state *configConfigurationRecorderStatusState
}

func (ccrs *ConfigConfigurationRecorderStatus) Type() string {
	return "aws_config_configuration_recorder_status"
}

func (ccrs *ConfigConfigurationRecorderStatus) LocalName() string {
	return ccrs.Name
}

func (ccrs *ConfigConfigurationRecorderStatus) Configuration() interface{} {
	return ccrs.Args
}

func (ccrs *ConfigConfigurationRecorderStatus) Attributes() configConfigurationRecorderStatusAttributes {
	return configConfigurationRecorderStatusAttributes{ref: terra.ReferenceResource(ccrs)}
}

func (ccrs *ConfigConfigurationRecorderStatus) ImportState(av io.Reader) error {
	ccrs.state = &configConfigurationRecorderStatusState{}
	if err := json.NewDecoder(av).Decode(ccrs.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ccrs.Type(), ccrs.LocalName(), err)
	}
	return nil
}

func (ccrs *ConfigConfigurationRecorderStatus) State() (*configConfigurationRecorderStatusState, bool) {
	return ccrs.state, ccrs.state != nil
}

func (ccrs *ConfigConfigurationRecorderStatus) StateMust() *configConfigurationRecorderStatusState {
	if ccrs.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ccrs.Type(), ccrs.LocalName()))
	}
	return ccrs.state
}

func (ccrs *ConfigConfigurationRecorderStatus) DependOn() terra.Reference {
	return terra.ReferenceResource(ccrs)
}

type ConfigConfigurationRecorderStatusArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IsEnabled: bool, required
	IsEnabled terra.BoolValue `hcl:"is_enabled,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// DependsOn contains resources that ConfigConfigurationRecorderStatus depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type configConfigurationRecorderStatusAttributes struct {
	ref terra.Reference
}

func (ccrs configConfigurationRecorderStatusAttributes) Id() terra.StringValue {
	return terra.ReferenceString(ccrs.ref.Append("id"))
}

func (ccrs configConfigurationRecorderStatusAttributes) IsEnabled() terra.BoolValue {
	return terra.ReferenceBool(ccrs.ref.Append("is_enabled"))
}

func (ccrs configConfigurationRecorderStatusAttributes) Name() terra.StringValue {
	return terra.ReferenceString(ccrs.ref.Append("name"))
}

type configConfigurationRecorderStatusState struct {
	Id        string `json:"id"`
	IsEnabled bool   `json:"is_enabled"`
	Name      string `json:"name"`
}
