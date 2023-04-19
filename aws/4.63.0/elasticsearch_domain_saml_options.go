// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	elasticsearchdomainsamloptions "github.com/golingon/terraproviders/aws/4.63.0/elasticsearchdomainsamloptions"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewElasticsearchDomainSamlOptions creates a new instance of [ElasticsearchDomainSamlOptions].
func NewElasticsearchDomainSamlOptions(name string, args ElasticsearchDomainSamlOptionsArgs) *ElasticsearchDomainSamlOptions {
	return &ElasticsearchDomainSamlOptions{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ElasticsearchDomainSamlOptions)(nil)

// ElasticsearchDomainSamlOptions represents the Terraform resource aws_elasticsearch_domain_saml_options.
type ElasticsearchDomainSamlOptions struct {
	Name      string
	Args      ElasticsearchDomainSamlOptionsArgs
	state     *elasticsearchDomainSamlOptionsState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ElasticsearchDomainSamlOptions].
func (edso *ElasticsearchDomainSamlOptions) Type() string {
	return "aws_elasticsearch_domain_saml_options"
}

// LocalName returns the local name for [ElasticsearchDomainSamlOptions].
func (edso *ElasticsearchDomainSamlOptions) LocalName() string {
	return edso.Name
}

// Configuration returns the configuration (args) for [ElasticsearchDomainSamlOptions].
func (edso *ElasticsearchDomainSamlOptions) Configuration() interface{} {
	return edso.Args
}

// DependOn is used for other resources to depend on [ElasticsearchDomainSamlOptions].
func (edso *ElasticsearchDomainSamlOptions) DependOn() terra.Reference {
	return terra.ReferenceResource(edso)
}

// Dependencies returns the list of resources [ElasticsearchDomainSamlOptions] depends_on.
func (edso *ElasticsearchDomainSamlOptions) Dependencies() terra.Dependencies {
	return edso.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ElasticsearchDomainSamlOptions].
func (edso *ElasticsearchDomainSamlOptions) LifecycleManagement() *terra.Lifecycle {
	return edso.Lifecycle
}

// Attributes returns the attributes for [ElasticsearchDomainSamlOptions].
func (edso *ElasticsearchDomainSamlOptions) Attributes() elasticsearchDomainSamlOptionsAttributes {
	return elasticsearchDomainSamlOptionsAttributes{ref: terra.ReferenceResource(edso)}
}

// ImportState imports the given attribute values into [ElasticsearchDomainSamlOptions]'s state.
func (edso *ElasticsearchDomainSamlOptions) ImportState(av io.Reader) error {
	edso.state = &elasticsearchDomainSamlOptionsState{}
	if err := json.NewDecoder(av).Decode(edso.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", edso.Type(), edso.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ElasticsearchDomainSamlOptions] has state.
func (edso *ElasticsearchDomainSamlOptions) State() (*elasticsearchDomainSamlOptionsState, bool) {
	return edso.state, edso.state != nil
}

// StateMust returns the state for [ElasticsearchDomainSamlOptions]. Panics if the state is nil.
func (edso *ElasticsearchDomainSamlOptions) StateMust() *elasticsearchDomainSamlOptionsState {
	if edso.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", edso.Type(), edso.LocalName()))
	}
	return edso.state
}

// ElasticsearchDomainSamlOptionsArgs contains the configurations for aws_elasticsearch_domain_saml_options.
type ElasticsearchDomainSamlOptionsArgs struct {
	// DomainName: string, required
	DomainName terra.StringValue `hcl:"domain_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// SamlOptions: optional
	SamlOptions *elasticsearchdomainsamloptions.SamlOptions `hcl:"saml_options,block"`
	// Timeouts: optional
	Timeouts *elasticsearchdomainsamloptions.Timeouts `hcl:"timeouts,block"`
}
type elasticsearchDomainSamlOptionsAttributes struct {
	ref terra.Reference
}

// DomainName returns a reference to field domain_name of aws_elasticsearch_domain_saml_options.
func (edso elasticsearchDomainSamlOptionsAttributes) DomainName() terra.StringValue {
	return terra.ReferenceAsString(edso.ref.Append("domain_name"))
}

// Id returns a reference to field id of aws_elasticsearch_domain_saml_options.
func (edso elasticsearchDomainSamlOptionsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(edso.ref.Append("id"))
}

func (edso elasticsearchDomainSamlOptionsAttributes) SamlOptions() terra.ListValue[elasticsearchdomainsamloptions.SamlOptionsAttributes] {
	return terra.ReferenceAsList[elasticsearchdomainsamloptions.SamlOptionsAttributes](edso.ref.Append("saml_options"))
}

func (edso elasticsearchDomainSamlOptionsAttributes) Timeouts() elasticsearchdomainsamloptions.TimeoutsAttributes {
	return terra.ReferenceAsSingle[elasticsearchdomainsamloptions.TimeoutsAttributes](edso.ref.Append("timeouts"))
}

type elasticsearchDomainSamlOptionsState struct {
	DomainName  string                                            `json:"domain_name"`
	Id          string                                            `json:"id"`
	SamlOptions []elasticsearchdomainsamloptions.SamlOptionsState `json:"saml_options"`
	Timeouts    *elasticsearchdomainsamloptions.TimeoutsState     `json:"timeouts"`
}
