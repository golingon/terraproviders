// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	apprunnerobservabilityconfiguration "github.com/golingon/terraproviders/aws/5.8.0/apprunnerobservabilityconfiguration"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewApprunnerObservabilityConfiguration creates a new instance of [ApprunnerObservabilityConfiguration].
func NewApprunnerObservabilityConfiguration(name string, args ApprunnerObservabilityConfigurationArgs) *ApprunnerObservabilityConfiguration {
	return &ApprunnerObservabilityConfiguration{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ApprunnerObservabilityConfiguration)(nil)

// ApprunnerObservabilityConfiguration represents the Terraform resource aws_apprunner_observability_configuration.
type ApprunnerObservabilityConfiguration struct {
	Name      string
	Args      ApprunnerObservabilityConfigurationArgs
	state     *apprunnerObservabilityConfigurationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ApprunnerObservabilityConfiguration].
func (aoc *ApprunnerObservabilityConfiguration) Type() string {
	return "aws_apprunner_observability_configuration"
}

// LocalName returns the local name for [ApprunnerObservabilityConfiguration].
func (aoc *ApprunnerObservabilityConfiguration) LocalName() string {
	return aoc.Name
}

// Configuration returns the configuration (args) for [ApprunnerObservabilityConfiguration].
func (aoc *ApprunnerObservabilityConfiguration) Configuration() interface{} {
	return aoc.Args
}

// DependOn is used for other resources to depend on [ApprunnerObservabilityConfiguration].
func (aoc *ApprunnerObservabilityConfiguration) DependOn() terra.Reference {
	return terra.ReferenceResource(aoc)
}

// Dependencies returns the list of resources [ApprunnerObservabilityConfiguration] depends_on.
func (aoc *ApprunnerObservabilityConfiguration) Dependencies() terra.Dependencies {
	return aoc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ApprunnerObservabilityConfiguration].
func (aoc *ApprunnerObservabilityConfiguration) LifecycleManagement() *terra.Lifecycle {
	return aoc.Lifecycle
}

// Attributes returns the attributes for [ApprunnerObservabilityConfiguration].
func (aoc *ApprunnerObservabilityConfiguration) Attributes() apprunnerObservabilityConfigurationAttributes {
	return apprunnerObservabilityConfigurationAttributes{ref: terra.ReferenceResource(aoc)}
}

// ImportState imports the given attribute values into [ApprunnerObservabilityConfiguration]'s state.
func (aoc *ApprunnerObservabilityConfiguration) ImportState(av io.Reader) error {
	aoc.state = &apprunnerObservabilityConfigurationState{}
	if err := json.NewDecoder(av).Decode(aoc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", aoc.Type(), aoc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ApprunnerObservabilityConfiguration] has state.
func (aoc *ApprunnerObservabilityConfiguration) State() (*apprunnerObservabilityConfigurationState, bool) {
	return aoc.state, aoc.state != nil
}

// StateMust returns the state for [ApprunnerObservabilityConfiguration]. Panics if the state is nil.
func (aoc *ApprunnerObservabilityConfiguration) StateMust() *apprunnerObservabilityConfigurationState {
	if aoc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", aoc.Type(), aoc.LocalName()))
	}
	return aoc.state
}

// ApprunnerObservabilityConfigurationArgs contains the configurations for aws_apprunner_observability_configuration.
type ApprunnerObservabilityConfigurationArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ObservabilityConfigurationName: string, required
	ObservabilityConfigurationName terra.StringValue `hcl:"observability_configuration_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// TraceConfiguration: optional
	TraceConfiguration *apprunnerobservabilityconfiguration.TraceConfiguration `hcl:"trace_configuration,block"`
}
type apprunnerObservabilityConfigurationAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_apprunner_observability_configuration.
func (aoc apprunnerObservabilityConfigurationAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(aoc.ref.Append("arn"))
}

// Id returns a reference to field id of aws_apprunner_observability_configuration.
func (aoc apprunnerObservabilityConfigurationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aoc.ref.Append("id"))
}

// Latest returns a reference to field latest of aws_apprunner_observability_configuration.
func (aoc apprunnerObservabilityConfigurationAttributes) Latest() terra.BoolValue {
	return terra.ReferenceAsBool(aoc.ref.Append("latest"))
}

// ObservabilityConfigurationName returns a reference to field observability_configuration_name of aws_apprunner_observability_configuration.
func (aoc apprunnerObservabilityConfigurationAttributes) ObservabilityConfigurationName() terra.StringValue {
	return terra.ReferenceAsString(aoc.ref.Append("observability_configuration_name"))
}

// ObservabilityConfigurationRevision returns a reference to field observability_configuration_revision of aws_apprunner_observability_configuration.
func (aoc apprunnerObservabilityConfigurationAttributes) ObservabilityConfigurationRevision() terra.NumberValue {
	return terra.ReferenceAsNumber(aoc.ref.Append("observability_configuration_revision"))
}

// Status returns a reference to field status of aws_apprunner_observability_configuration.
func (aoc apprunnerObservabilityConfigurationAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(aoc.ref.Append("status"))
}

// Tags returns a reference to field tags of aws_apprunner_observability_configuration.
func (aoc apprunnerObservabilityConfigurationAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](aoc.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_apprunner_observability_configuration.
func (aoc apprunnerObservabilityConfigurationAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](aoc.ref.Append("tags_all"))
}

func (aoc apprunnerObservabilityConfigurationAttributes) TraceConfiguration() terra.ListValue[apprunnerobservabilityconfiguration.TraceConfigurationAttributes] {
	return terra.ReferenceAsList[apprunnerobservabilityconfiguration.TraceConfigurationAttributes](aoc.ref.Append("trace_configuration"))
}

type apprunnerObservabilityConfigurationState struct {
	Arn                                string                                                        `json:"arn"`
	Id                                 string                                                        `json:"id"`
	Latest                             bool                                                          `json:"latest"`
	ObservabilityConfigurationName     string                                                        `json:"observability_configuration_name"`
	ObservabilityConfigurationRevision float64                                                       `json:"observability_configuration_revision"`
	Status                             string                                                        `json:"status"`
	Tags                               map[string]string                                             `json:"tags"`
	TagsAll                            map[string]string                                             `json:"tags_all"`
	TraceConfiguration                 []apprunnerobservabilityconfiguration.TraceConfigurationState `json:"trace_configuration"`
}
