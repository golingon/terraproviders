// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	ivschatloggingconfiguration "github.com/golingon/terraproviders/aws/4.60.0/ivschatloggingconfiguration"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewIvschatLoggingConfiguration creates a new instance of [IvschatLoggingConfiguration].
func NewIvschatLoggingConfiguration(name string, args IvschatLoggingConfigurationArgs) *IvschatLoggingConfiguration {
	return &IvschatLoggingConfiguration{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*IvschatLoggingConfiguration)(nil)

// IvschatLoggingConfiguration represents the Terraform resource aws_ivschat_logging_configuration.
type IvschatLoggingConfiguration struct {
	Name      string
	Args      IvschatLoggingConfigurationArgs
	state     *ivschatLoggingConfigurationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [IvschatLoggingConfiguration].
func (ilc *IvschatLoggingConfiguration) Type() string {
	return "aws_ivschat_logging_configuration"
}

// LocalName returns the local name for [IvschatLoggingConfiguration].
func (ilc *IvschatLoggingConfiguration) LocalName() string {
	return ilc.Name
}

// Configuration returns the configuration (args) for [IvschatLoggingConfiguration].
func (ilc *IvschatLoggingConfiguration) Configuration() interface{} {
	return ilc.Args
}

// DependOn is used for other resources to depend on [IvschatLoggingConfiguration].
func (ilc *IvschatLoggingConfiguration) DependOn() terra.Reference {
	return terra.ReferenceResource(ilc)
}

// Dependencies returns the list of resources [IvschatLoggingConfiguration] depends_on.
func (ilc *IvschatLoggingConfiguration) Dependencies() terra.Dependencies {
	return ilc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [IvschatLoggingConfiguration].
func (ilc *IvschatLoggingConfiguration) LifecycleManagement() *terra.Lifecycle {
	return ilc.Lifecycle
}

// Attributes returns the attributes for [IvschatLoggingConfiguration].
func (ilc *IvschatLoggingConfiguration) Attributes() ivschatLoggingConfigurationAttributes {
	return ivschatLoggingConfigurationAttributes{ref: terra.ReferenceResource(ilc)}
}

// ImportState imports the given attribute values into [IvschatLoggingConfiguration]'s state.
func (ilc *IvschatLoggingConfiguration) ImportState(av io.Reader) error {
	ilc.state = &ivschatLoggingConfigurationState{}
	if err := json.NewDecoder(av).Decode(ilc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ilc.Type(), ilc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [IvschatLoggingConfiguration] has state.
func (ilc *IvschatLoggingConfiguration) State() (*ivschatLoggingConfigurationState, bool) {
	return ilc.state, ilc.state != nil
}

// StateMust returns the state for [IvschatLoggingConfiguration]. Panics if the state is nil.
func (ilc *IvschatLoggingConfiguration) StateMust() *ivschatLoggingConfigurationState {
	if ilc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ilc.Type(), ilc.LocalName()))
	}
	return ilc.state
}

// IvschatLoggingConfigurationArgs contains the configurations for aws_ivschat_logging_configuration.
type IvschatLoggingConfigurationArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// DestinationConfiguration: optional
	DestinationConfiguration *ivschatloggingconfiguration.DestinationConfiguration `hcl:"destination_configuration,block"`
	// Timeouts: optional
	Timeouts *ivschatloggingconfiguration.Timeouts `hcl:"timeouts,block"`
}
type ivschatLoggingConfigurationAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_ivschat_logging_configuration.
func (ilc ivschatLoggingConfigurationAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ilc.ref.Append("arn"))
}

// Id returns a reference to field id of aws_ivschat_logging_configuration.
func (ilc ivschatLoggingConfigurationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ilc.ref.Append("id"))
}

// Name returns a reference to field name of aws_ivschat_logging_configuration.
func (ilc ivschatLoggingConfigurationAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ilc.ref.Append("name"))
}

// State returns a reference to field state of aws_ivschat_logging_configuration.
func (ilc ivschatLoggingConfigurationAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(ilc.ref.Append("state"))
}

// Tags returns a reference to field tags of aws_ivschat_logging_configuration.
func (ilc ivschatLoggingConfigurationAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ilc.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_ivschat_logging_configuration.
func (ilc ivschatLoggingConfigurationAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ilc.ref.Append("tags_all"))
}

func (ilc ivschatLoggingConfigurationAttributes) DestinationConfiguration() terra.ListValue[ivschatloggingconfiguration.DestinationConfigurationAttributes] {
	return terra.ReferenceAsList[ivschatloggingconfiguration.DestinationConfigurationAttributes](ilc.ref.Append("destination_configuration"))
}

func (ilc ivschatLoggingConfigurationAttributes) Timeouts() ivschatloggingconfiguration.TimeoutsAttributes {
	return terra.ReferenceAsSingle[ivschatloggingconfiguration.TimeoutsAttributes](ilc.ref.Append("timeouts"))
}

type ivschatLoggingConfigurationState struct {
	Arn                      string                                                      `json:"arn"`
	Id                       string                                                      `json:"id"`
	Name                     string                                                      `json:"name"`
	State                    string                                                      `json:"state"`
	Tags                     map[string]string                                           `json:"tags"`
	TagsAll                  map[string]string                                           `json:"tags_all"`
	DestinationConfiguration []ivschatloggingconfiguration.DestinationConfigurationState `json:"destination_configuration"`
	Timeouts                 *ivschatloggingconfiguration.TimeoutsState                  `json:"timeouts"`
}
