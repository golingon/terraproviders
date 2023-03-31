// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	configconfigurationaggregator "github.com/golingon/terraproviders/aws/4.60.0/configconfigurationaggregator"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewConfigConfigurationAggregator creates a new instance of [ConfigConfigurationAggregator].
func NewConfigConfigurationAggregator(name string, args ConfigConfigurationAggregatorArgs) *ConfigConfigurationAggregator {
	return &ConfigConfigurationAggregator{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ConfigConfigurationAggregator)(nil)

// ConfigConfigurationAggregator represents the Terraform resource aws_config_configuration_aggregator.
type ConfigConfigurationAggregator struct {
	Name      string
	Args      ConfigConfigurationAggregatorArgs
	state     *configConfigurationAggregatorState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ConfigConfigurationAggregator].
func (cca *ConfigConfigurationAggregator) Type() string {
	return "aws_config_configuration_aggregator"
}

// LocalName returns the local name for [ConfigConfigurationAggregator].
func (cca *ConfigConfigurationAggregator) LocalName() string {
	return cca.Name
}

// Configuration returns the configuration (args) for [ConfigConfigurationAggregator].
func (cca *ConfigConfigurationAggregator) Configuration() interface{} {
	return cca.Args
}

// DependOn is used for other resources to depend on [ConfigConfigurationAggregator].
func (cca *ConfigConfigurationAggregator) DependOn() terra.Reference {
	return terra.ReferenceResource(cca)
}

// Dependencies returns the list of resources [ConfigConfigurationAggregator] depends_on.
func (cca *ConfigConfigurationAggregator) Dependencies() terra.Dependencies {
	return cca.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ConfigConfigurationAggregator].
func (cca *ConfigConfigurationAggregator) LifecycleManagement() *terra.Lifecycle {
	return cca.Lifecycle
}

// Attributes returns the attributes for [ConfigConfigurationAggregator].
func (cca *ConfigConfigurationAggregator) Attributes() configConfigurationAggregatorAttributes {
	return configConfigurationAggregatorAttributes{ref: terra.ReferenceResource(cca)}
}

// ImportState imports the given attribute values into [ConfigConfigurationAggregator]'s state.
func (cca *ConfigConfigurationAggregator) ImportState(av io.Reader) error {
	cca.state = &configConfigurationAggregatorState{}
	if err := json.NewDecoder(av).Decode(cca.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cca.Type(), cca.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ConfigConfigurationAggregator] has state.
func (cca *ConfigConfigurationAggregator) State() (*configConfigurationAggregatorState, bool) {
	return cca.state, cca.state != nil
}

// StateMust returns the state for [ConfigConfigurationAggregator]. Panics if the state is nil.
func (cca *ConfigConfigurationAggregator) StateMust() *configConfigurationAggregatorState {
	if cca.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cca.Type(), cca.LocalName()))
	}
	return cca.state
}

// ConfigConfigurationAggregatorArgs contains the configurations for aws_config_configuration_aggregator.
type ConfigConfigurationAggregatorArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// AccountAggregationSource: optional
	AccountAggregationSource *configconfigurationaggregator.AccountAggregationSource `hcl:"account_aggregation_source,block"`
	// OrganizationAggregationSource: optional
	OrganizationAggregationSource *configconfigurationaggregator.OrganizationAggregationSource `hcl:"organization_aggregation_source,block"`
}
type configConfigurationAggregatorAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_config_configuration_aggregator.
func (cca configConfigurationAggregatorAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(cca.ref.Append("arn"))
}

// Id returns a reference to field id of aws_config_configuration_aggregator.
func (cca configConfigurationAggregatorAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cca.ref.Append("id"))
}

// Name returns a reference to field name of aws_config_configuration_aggregator.
func (cca configConfigurationAggregatorAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cca.ref.Append("name"))
}

// Tags returns a reference to field tags of aws_config_configuration_aggregator.
func (cca configConfigurationAggregatorAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cca.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_config_configuration_aggregator.
func (cca configConfigurationAggregatorAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cca.ref.Append("tags_all"))
}

func (cca configConfigurationAggregatorAttributes) AccountAggregationSource() terra.ListValue[configconfigurationaggregator.AccountAggregationSourceAttributes] {
	return terra.ReferenceAsList[configconfigurationaggregator.AccountAggregationSourceAttributes](cca.ref.Append("account_aggregation_source"))
}

func (cca configConfigurationAggregatorAttributes) OrganizationAggregationSource() terra.ListValue[configconfigurationaggregator.OrganizationAggregationSourceAttributes] {
	return terra.ReferenceAsList[configconfigurationaggregator.OrganizationAggregationSourceAttributes](cca.ref.Append("organization_aggregation_source"))
}

type configConfigurationAggregatorState struct {
	Arn                           string                                                             `json:"arn"`
	Id                            string                                                             `json:"id"`
	Name                          string                                                             `json:"name"`
	Tags                          map[string]string                                                  `json:"tags"`
	TagsAll                       map[string]string                                                  `json:"tags_all"`
	AccountAggregationSource      []configconfigurationaggregator.AccountAggregationSourceState      `json:"account_aggregation_source"`
	OrganizationAggregationSource []configconfigurationaggregator.OrganizationAggregationSourceState `json:"organization_aggregation_source"`
}
