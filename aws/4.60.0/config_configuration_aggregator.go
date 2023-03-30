// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	configconfigurationaggregator "github.com/golingon/terraproviders/aws/4.60.0/configconfigurationaggregator"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewConfigConfigurationAggregator(name string, args ConfigConfigurationAggregatorArgs) *ConfigConfigurationAggregator {
	return &ConfigConfigurationAggregator{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ConfigConfigurationAggregator)(nil)

type ConfigConfigurationAggregator struct {
	Name  string
	Args  ConfigConfigurationAggregatorArgs
	state *configConfigurationAggregatorState
}

func (cca *ConfigConfigurationAggregator) Type() string {
	return "aws_config_configuration_aggregator"
}

func (cca *ConfigConfigurationAggregator) LocalName() string {
	return cca.Name
}

func (cca *ConfigConfigurationAggregator) Configuration() interface{} {
	return cca.Args
}

func (cca *ConfigConfigurationAggregator) Attributes() configConfigurationAggregatorAttributes {
	return configConfigurationAggregatorAttributes{ref: terra.ReferenceResource(cca)}
}

func (cca *ConfigConfigurationAggregator) ImportState(av io.Reader) error {
	cca.state = &configConfigurationAggregatorState{}
	if err := json.NewDecoder(av).Decode(cca.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cca.Type(), cca.LocalName(), err)
	}
	return nil
}

func (cca *ConfigConfigurationAggregator) State() (*configConfigurationAggregatorState, bool) {
	return cca.state, cca.state != nil
}

func (cca *ConfigConfigurationAggregator) StateMust() *configConfigurationAggregatorState {
	if cca.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cca.Type(), cca.LocalName()))
	}
	return cca.state
}

func (cca *ConfigConfigurationAggregator) DependOn() terra.Reference {
	return terra.ReferenceResource(cca)
}

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
	// DependsOn contains resources that ConfigConfigurationAggregator depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type configConfigurationAggregatorAttributes struct {
	ref terra.Reference
}

func (cca configConfigurationAggregatorAttributes) Arn() terra.StringValue {
	return terra.ReferenceString(cca.ref.Append("arn"))
}

func (cca configConfigurationAggregatorAttributes) Id() terra.StringValue {
	return terra.ReferenceString(cca.ref.Append("id"))
}

func (cca configConfigurationAggregatorAttributes) Name() terra.StringValue {
	return terra.ReferenceString(cca.ref.Append("name"))
}

func (cca configConfigurationAggregatorAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](cca.ref.Append("tags"))
}

func (cca configConfigurationAggregatorAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](cca.ref.Append("tags_all"))
}

func (cca configConfigurationAggregatorAttributes) AccountAggregationSource() terra.ListValue[configconfigurationaggregator.AccountAggregationSourceAttributes] {
	return terra.ReferenceList[configconfigurationaggregator.AccountAggregationSourceAttributes](cca.ref.Append("account_aggregation_source"))
}

func (cca configConfigurationAggregatorAttributes) OrganizationAggregationSource() terra.ListValue[configconfigurationaggregator.OrganizationAggregationSourceAttributes] {
	return terra.ReferenceList[configconfigurationaggregator.OrganizationAggregationSourceAttributes](cca.ref.Append("organization_aggregation_source"))
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
