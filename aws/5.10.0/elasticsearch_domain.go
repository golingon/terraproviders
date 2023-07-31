// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	elasticsearchdomain "github.com/golingon/terraproviders/aws/5.10.0/elasticsearchdomain"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewElasticsearchDomain creates a new instance of [ElasticsearchDomain].
func NewElasticsearchDomain(name string, args ElasticsearchDomainArgs) *ElasticsearchDomain {
	return &ElasticsearchDomain{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ElasticsearchDomain)(nil)

// ElasticsearchDomain represents the Terraform resource aws_elasticsearch_domain.
type ElasticsearchDomain struct {
	Name      string
	Args      ElasticsearchDomainArgs
	state     *elasticsearchDomainState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ElasticsearchDomain].
func (ed *ElasticsearchDomain) Type() string {
	return "aws_elasticsearch_domain"
}

// LocalName returns the local name for [ElasticsearchDomain].
func (ed *ElasticsearchDomain) LocalName() string {
	return ed.Name
}

// Configuration returns the configuration (args) for [ElasticsearchDomain].
func (ed *ElasticsearchDomain) Configuration() interface{} {
	return ed.Args
}

// DependOn is used for other resources to depend on [ElasticsearchDomain].
func (ed *ElasticsearchDomain) DependOn() terra.Reference {
	return terra.ReferenceResource(ed)
}

// Dependencies returns the list of resources [ElasticsearchDomain] depends_on.
func (ed *ElasticsearchDomain) Dependencies() terra.Dependencies {
	return ed.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ElasticsearchDomain].
func (ed *ElasticsearchDomain) LifecycleManagement() *terra.Lifecycle {
	return ed.Lifecycle
}

// Attributes returns the attributes for [ElasticsearchDomain].
func (ed *ElasticsearchDomain) Attributes() elasticsearchDomainAttributes {
	return elasticsearchDomainAttributes{ref: terra.ReferenceResource(ed)}
}

// ImportState imports the given attribute values into [ElasticsearchDomain]'s state.
func (ed *ElasticsearchDomain) ImportState(av io.Reader) error {
	ed.state = &elasticsearchDomainState{}
	if err := json.NewDecoder(av).Decode(ed.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ed.Type(), ed.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ElasticsearchDomain] has state.
func (ed *ElasticsearchDomain) State() (*elasticsearchDomainState, bool) {
	return ed.state, ed.state != nil
}

// StateMust returns the state for [ElasticsearchDomain]. Panics if the state is nil.
func (ed *ElasticsearchDomain) StateMust() *elasticsearchDomainState {
	if ed.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ed.Type(), ed.LocalName()))
	}
	return ed.state
}

// ElasticsearchDomainArgs contains the configurations for aws_elasticsearch_domain.
type ElasticsearchDomainArgs struct {
	// AccessPolicies: string, optional
	AccessPolicies terra.StringValue `hcl:"access_policies,attr"`
	// AdvancedOptions: map of string, optional
	AdvancedOptions terra.MapValue[terra.StringValue] `hcl:"advanced_options,attr"`
	// DomainName: string, required
	DomainName terra.StringValue `hcl:"domain_name,attr" validate:"required"`
	// ElasticsearchVersion: string, optional
	ElasticsearchVersion terra.StringValue `hcl:"elasticsearch_version,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// AdvancedSecurityOptions: optional
	AdvancedSecurityOptions *elasticsearchdomain.AdvancedSecurityOptions `hcl:"advanced_security_options,block"`
	// AutoTuneOptions: optional
	AutoTuneOptions *elasticsearchdomain.AutoTuneOptions `hcl:"auto_tune_options,block"`
	// ClusterConfig: optional
	ClusterConfig *elasticsearchdomain.ClusterConfig `hcl:"cluster_config,block"`
	// CognitoOptions: optional
	CognitoOptions *elasticsearchdomain.CognitoOptions `hcl:"cognito_options,block"`
	// DomainEndpointOptions: optional
	DomainEndpointOptions *elasticsearchdomain.DomainEndpointOptions `hcl:"domain_endpoint_options,block"`
	// EbsOptions: optional
	EbsOptions *elasticsearchdomain.EbsOptions `hcl:"ebs_options,block"`
	// EncryptAtRest: optional
	EncryptAtRest *elasticsearchdomain.EncryptAtRest `hcl:"encrypt_at_rest,block"`
	// LogPublishingOptions: min=0
	LogPublishingOptions []elasticsearchdomain.LogPublishingOptions `hcl:"log_publishing_options,block" validate:"min=0"`
	// NodeToNodeEncryption: optional
	NodeToNodeEncryption *elasticsearchdomain.NodeToNodeEncryption `hcl:"node_to_node_encryption,block"`
	// SnapshotOptions: optional
	SnapshotOptions *elasticsearchdomain.SnapshotOptions `hcl:"snapshot_options,block"`
	// Timeouts: optional
	Timeouts *elasticsearchdomain.Timeouts `hcl:"timeouts,block"`
	// VpcOptions: optional
	VpcOptions *elasticsearchdomain.VpcOptions `hcl:"vpc_options,block"`
}
type elasticsearchDomainAttributes struct {
	ref terra.Reference
}

// AccessPolicies returns a reference to field access_policies of aws_elasticsearch_domain.
func (ed elasticsearchDomainAttributes) AccessPolicies() terra.StringValue {
	return terra.ReferenceAsString(ed.ref.Append("access_policies"))
}

// AdvancedOptions returns a reference to field advanced_options of aws_elasticsearch_domain.
func (ed elasticsearchDomainAttributes) AdvancedOptions() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ed.ref.Append("advanced_options"))
}

// Arn returns a reference to field arn of aws_elasticsearch_domain.
func (ed elasticsearchDomainAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ed.ref.Append("arn"))
}

// DomainId returns a reference to field domain_id of aws_elasticsearch_domain.
func (ed elasticsearchDomainAttributes) DomainId() terra.StringValue {
	return terra.ReferenceAsString(ed.ref.Append("domain_id"))
}

// DomainName returns a reference to field domain_name of aws_elasticsearch_domain.
func (ed elasticsearchDomainAttributes) DomainName() terra.StringValue {
	return terra.ReferenceAsString(ed.ref.Append("domain_name"))
}

// ElasticsearchVersion returns a reference to field elasticsearch_version of aws_elasticsearch_domain.
func (ed elasticsearchDomainAttributes) ElasticsearchVersion() terra.StringValue {
	return terra.ReferenceAsString(ed.ref.Append("elasticsearch_version"))
}

// Endpoint returns a reference to field endpoint of aws_elasticsearch_domain.
func (ed elasticsearchDomainAttributes) Endpoint() terra.StringValue {
	return terra.ReferenceAsString(ed.ref.Append("endpoint"))
}

// Id returns a reference to field id of aws_elasticsearch_domain.
func (ed elasticsearchDomainAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ed.ref.Append("id"))
}

// KibanaEndpoint returns a reference to field kibana_endpoint of aws_elasticsearch_domain.
func (ed elasticsearchDomainAttributes) KibanaEndpoint() terra.StringValue {
	return terra.ReferenceAsString(ed.ref.Append("kibana_endpoint"))
}

// Tags returns a reference to field tags of aws_elasticsearch_domain.
func (ed elasticsearchDomainAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ed.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_elasticsearch_domain.
func (ed elasticsearchDomainAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ed.ref.Append("tags_all"))
}

func (ed elasticsearchDomainAttributes) AdvancedSecurityOptions() terra.ListValue[elasticsearchdomain.AdvancedSecurityOptionsAttributes] {
	return terra.ReferenceAsList[elasticsearchdomain.AdvancedSecurityOptionsAttributes](ed.ref.Append("advanced_security_options"))
}

func (ed elasticsearchDomainAttributes) AutoTuneOptions() terra.ListValue[elasticsearchdomain.AutoTuneOptionsAttributes] {
	return terra.ReferenceAsList[elasticsearchdomain.AutoTuneOptionsAttributes](ed.ref.Append("auto_tune_options"))
}

func (ed elasticsearchDomainAttributes) ClusterConfig() terra.ListValue[elasticsearchdomain.ClusterConfigAttributes] {
	return terra.ReferenceAsList[elasticsearchdomain.ClusterConfigAttributes](ed.ref.Append("cluster_config"))
}

func (ed elasticsearchDomainAttributes) CognitoOptions() terra.ListValue[elasticsearchdomain.CognitoOptionsAttributes] {
	return terra.ReferenceAsList[elasticsearchdomain.CognitoOptionsAttributes](ed.ref.Append("cognito_options"))
}

func (ed elasticsearchDomainAttributes) DomainEndpointOptions() terra.ListValue[elasticsearchdomain.DomainEndpointOptionsAttributes] {
	return terra.ReferenceAsList[elasticsearchdomain.DomainEndpointOptionsAttributes](ed.ref.Append("domain_endpoint_options"))
}

func (ed elasticsearchDomainAttributes) EbsOptions() terra.ListValue[elasticsearchdomain.EbsOptionsAttributes] {
	return terra.ReferenceAsList[elasticsearchdomain.EbsOptionsAttributes](ed.ref.Append("ebs_options"))
}

func (ed elasticsearchDomainAttributes) EncryptAtRest() terra.ListValue[elasticsearchdomain.EncryptAtRestAttributes] {
	return terra.ReferenceAsList[elasticsearchdomain.EncryptAtRestAttributes](ed.ref.Append("encrypt_at_rest"))
}

func (ed elasticsearchDomainAttributes) LogPublishingOptions() terra.SetValue[elasticsearchdomain.LogPublishingOptionsAttributes] {
	return terra.ReferenceAsSet[elasticsearchdomain.LogPublishingOptionsAttributes](ed.ref.Append("log_publishing_options"))
}

func (ed elasticsearchDomainAttributes) NodeToNodeEncryption() terra.ListValue[elasticsearchdomain.NodeToNodeEncryptionAttributes] {
	return terra.ReferenceAsList[elasticsearchdomain.NodeToNodeEncryptionAttributes](ed.ref.Append("node_to_node_encryption"))
}

func (ed elasticsearchDomainAttributes) SnapshotOptions() terra.ListValue[elasticsearchdomain.SnapshotOptionsAttributes] {
	return terra.ReferenceAsList[elasticsearchdomain.SnapshotOptionsAttributes](ed.ref.Append("snapshot_options"))
}

func (ed elasticsearchDomainAttributes) Timeouts() elasticsearchdomain.TimeoutsAttributes {
	return terra.ReferenceAsSingle[elasticsearchdomain.TimeoutsAttributes](ed.ref.Append("timeouts"))
}

func (ed elasticsearchDomainAttributes) VpcOptions() terra.ListValue[elasticsearchdomain.VpcOptionsAttributes] {
	return terra.ReferenceAsList[elasticsearchdomain.VpcOptionsAttributes](ed.ref.Append("vpc_options"))
}

type elasticsearchDomainState struct {
	AccessPolicies          string                                             `json:"access_policies"`
	AdvancedOptions         map[string]string                                  `json:"advanced_options"`
	Arn                     string                                             `json:"arn"`
	DomainId                string                                             `json:"domain_id"`
	DomainName              string                                             `json:"domain_name"`
	ElasticsearchVersion    string                                             `json:"elasticsearch_version"`
	Endpoint                string                                             `json:"endpoint"`
	Id                      string                                             `json:"id"`
	KibanaEndpoint          string                                             `json:"kibana_endpoint"`
	Tags                    map[string]string                                  `json:"tags"`
	TagsAll                 map[string]string                                  `json:"tags_all"`
	AdvancedSecurityOptions []elasticsearchdomain.AdvancedSecurityOptionsState `json:"advanced_security_options"`
	AutoTuneOptions         []elasticsearchdomain.AutoTuneOptionsState         `json:"auto_tune_options"`
	ClusterConfig           []elasticsearchdomain.ClusterConfigState           `json:"cluster_config"`
	CognitoOptions          []elasticsearchdomain.CognitoOptionsState          `json:"cognito_options"`
	DomainEndpointOptions   []elasticsearchdomain.DomainEndpointOptionsState   `json:"domain_endpoint_options"`
	EbsOptions              []elasticsearchdomain.EbsOptionsState              `json:"ebs_options"`
	EncryptAtRest           []elasticsearchdomain.EncryptAtRestState           `json:"encrypt_at_rest"`
	LogPublishingOptions    []elasticsearchdomain.LogPublishingOptionsState    `json:"log_publishing_options"`
	NodeToNodeEncryption    []elasticsearchdomain.NodeToNodeEncryptionState    `json:"node_to_node_encryption"`
	SnapshotOptions         []elasticsearchdomain.SnapshotOptionsState         `json:"snapshot_options"`
	Timeouts                *elasticsearchdomain.TimeoutsState                 `json:"timeouts"`
	VpcOptions              []elasticsearchdomain.VpcOptionsState              `json:"vpc_options"`
}
