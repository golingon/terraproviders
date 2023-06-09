// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	iotindexingconfiguration "github.com/golingon/terraproviders/aws/5.6.2/iotindexingconfiguration"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewIotIndexingConfiguration creates a new instance of [IotIndexingConfiguration].
func NewIotIndexingConfiguration(name string, args IotIndexingConfigurationArgs) *IotIndexingConfiguration {
	return &IotIndexingConfiguration{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*IotIndexingConfiguration)(nil)

// IotIndexingConfiguration represents the Terraform resource aws_iot_indexing_configuration.
type IotIndexingConfiguration struct {
	Name      string
	Args      IotIndexingConfigurationArgs
	state     *iotIndexingConfigurationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [IotIndexingConfiguration].
func (iic *IotIndexingConfiguration) Type() string {
	return "aws_iot_indexing_configuration"
}

// LocalName returns the local name for [IotIndexingConfiguration].
func (iic *IotIndexingConfiguration) LocalName() string {
	return iic.Name
}

// Configuration returns the configuration (args) for [IotIndexingConfiguration].
func (iic *IotIndexingConfiguration) Configuration() interface{} {
	return iic.Args
}

// DependOn is used for other resources to depend on [IotIndexingConfiguration].
func (iic *IotIndexingConfiguration) DependOn() terra.Reference {
	return terra.ReferenceResource(iic)
}

// Dependencies returns the list of resources [IotIndexingConfiguration] depends_on.
func (iic *IotIndexingConfiguration) Dependencies() terra.Dependencies {
	return iic.DependsOn
}

// LifecycleManagement returns the lifecycle block for [IotIndexingConfiguration].
func (iic *IotIndexingConfiguration) LifecycleManagement() *terra.Lifecycle {
	return iic.Lifecycle
}

// Attributes returns the attributes for [IotIndexingConfiguration].
func (iic *IotIndexingConfiguration) Attributes() iotIndexingConfigurationAttributes {
	return iotIndexingConfigurationAttributes{ref: terra.ReferenceResource(iic)}
}

// ImportState imports the given attribute values into [IotIndexingConfiguration]'s state.
func (iic *IotIndexingConfiguration) ImportState(av io.Reader) error {
	iic.state = &iotIndexingConfigurationState{}
	if err := json.NewDecoder(av).Decode(iic.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", iic.Type(), iic.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [IotIndexingConfiguration] has state.
func (iic *IotIndexingConfiguration) State() (*iotIndexingConfigurationState, bool) {
	return iic.state, iic.state != nil
}

// StateMust returns the state for [IotIndexingConfiguration]. Panics if the state is nil.
func (iic *IotIndexingConfiguration) StateMust() *iotIndexingConfigurationState {
	if iic.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", iic.Type(), iic.LocalName()))
	}
	return iic.state
}

// IotIndexingConfigurationArgs contains the configurations for aws_iot_indexing_configuration.
type IotIndexingConfigurationArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ThingGroupIndexingConfiguration: optional
	ThingGroupIndexingConfiguration *iotindexingconfiguration.ThingGroupIndexingConfiguration `hcl:"thing_group_indexing_configuration,block"`
	// ThingIndexingConfiguration: optional
	ThingIndexingConfiguration *iotindexingconfiguration.ThingIndexingConfiguration `hcl:"thing_indexing_configuration,block"`
}
type iotIndexingConfigurationAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_iot_indexing_configuration.
func (iic iotIndexingConfigurationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(iic.ref.Append("id"))
}

func (iic iotIndexingConfigurationAttributes) ThingGroupIndexingConfiguration() terra.ListValue[iotindexingconfiguration.ThingGroupIndexingConfigurationAttributes] {
	return terra.ReferenceAsList[iotindexingconfiguration.ThingGroupIndexingConfigurationAttributes](iic.ref.Append("thing_group_indexing_configuration"))
}

func (iic iotIndexingConfigurationAttributes) ThingIndexingConfiguration() terra.ListValue[iotindexingconfiguration.ThingIndexingConfigurationAttributes] {
	return terra.ReferenceAsList[iotindexingconfiguration.ThingIndexingConfigurationAttributes](iic.ref.Append("thing_indexing_configuration"))
}

type iotIndexingConfigurationState struct {
	Id                              string                                                          `json:"id"`
	ThingGroupIndexingConfiguration []iotindexingconfiguration.ThingGroupIndexingConfigurationState `json:"thing_group_indexing_configuration"`
	ThingIndexingConfiguration      []iotindexingconfiguration.ThingIndexingConfigurationState      `json:"thing_indexing_configuration"`
}
