// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	elasticsearchdomainsamloptions "github.com/golingon/terraproviders/aws/4.60.0/elasticsearchdomainsamloptions"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewElasticsearchDomainSamlOptions(name string, args ElasticsearchDomainSamlOptionsArgs) *ElasticsearchDomainSamlOptions {
	return &ElasticsearchDomainSamlOptions{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ElasticsearchDomainSamlOptions)(nil)

type ElasticsearchDomainSamlOptions struct {
	Name  string
	Args  ElasticsearchDomainSamlOptionsArgs
	state *elasticsearchDomainSamlOptionsState
}

func (edso *ElasticsearchDomainSamlOptions) Type() string {
	return "aws_elasticsearch_domain_saml_options"
}

func (edso *ElasticsearchDomainSamlOptions) LocalName() string {
	return edso.Name
}

func (edso *ElasticsearchDomainSamlOptions) Configuration() interface{} {
	return edso.Args
}

func (edso *ElasticsearchDomainSamlOptions) Attributes() elasticsearchDomainSamlOptionsAttributes {
	return elasticsearchDomainSamlOptionsAttributes{ref: terra.ReferenceResource(edso)}
}

func (edso *ElasticsearchDomainSamlOptions) ImportState(av io.Reader) error {
	edso.state = &elasticsearchDomainSamlOptionsState{}
	if err := json.NewDecoder(av).Decode(edso.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", edso.Type(), edso.LocalName(), err)
	}
	return nil
}

func (edso *ElasticsearchDomainSamlOptions) State() (*elasticsearchDomainSamlOptionsState, bool) {
	return edso.state, edso.state != nil
}

func (edso *ElasticsearchDomainSamlOptions) StateMust() *elasticsearchDomainSamlOptionsState {
	if edso.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", edso.Type(), edso.LocalName()))
	}
	return edso.state
}

func (edso *ElasticsearchDomainSamlOptions) DependOn() terra.Reference {
	return terra.ReferenceResource(edso)
}

type ElasticsearchDomainSamlOptionsArgs struct {
	// DomainName: string, required
	DomainName terra.StringValue `hcl:"domain_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// SamlOptions: optional
	SamlOptions *elasticsearchdomainsamloptions.SamlOptions `hcl:"saml_options,block"`
	// Timeouts: optional
	Timeouts *elasticsearchdomainsamloptions.Timeouts `hcl:"timeouts,block"`
	// DependsOn contains resources that ElasticsearchDomainSamlOptions depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type elasticsearchDomainSamlOptionsAttributes struct {
	ref terra.Reference
}

func (edso elasticsearchDomainSamlOptionsAttributes) DomainName() terra.StringValue {
	return terra.ReferenceString(edso.ref.Append("domain_name"))
}

func (edso elasticsearchDomainSamlOptionsAttributes) Id() terra.StringValue {
	return terra.ReferenceString(edso.ref.Append("id"))
}

func (edso elasticsearchDomainSamlOptionsAttributes) SamlOptions() terra.ListValue[elasticsearchdomainsamloptions.SamlOptionsAttributes] {
	return terra.ReferenceList[elasticsearchdomainsamloptions.SamlOptionsAttributes](edso.ref.Append("saml_options"))
}

func (edso elasticsearchDomainSamlOptionsAttributes) Timeouts() elasticsearchdomainsamloptions.TimeoutsAttributes {
	return terra.ReferenceSingle[elasticsearchdomainsamloptions.TimeoutsAttributes](edso.ref.Append("timeouts"))
}

type elasticsearchDomainSamlOptionsState struct {
	DomainName  string                                            `json:"domain_name"`
	Id          string                                            `json:"id"`
	SamlOptions []elasticsearchdomainsamloptions.SamlOptionsState `json:"saml_options"`
	Timeouts    *elasticsearchdomainsamloptions.TimeoutsState     `json:"timeouts"`
}
