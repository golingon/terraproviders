// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	elasticsearchdomain "github.com/golingon/terraproviders/aws/4.60.0/elasticsearchdomain"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewElasticsearchDomain(name string, args ElasticsearchDomainArgs) *ElasticsearchDomain {
	return &ElasticsearchDomain{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ElasticsearchDomain)(nil)

type ElasticsearchDomain struct {
	Name  string
	Args  ElasticsearchDomainArgs
	state *elasticsearchDomainState
}

func (ed *ElasticsearchDomain) Type() string {
	return "aws_elasticsearch_domain"
}

func (ed *ElasticsearchDomain) LocalName() string {
	return ed.Name
}

func (ed *ElasticsearchDomain) Configuration() interface{} {
	return ed.Args
}

func (ed *ElasticsearchDomain) Attributes() elasticsearchDomainAttributes {
	return elasticsearchDomainAttributes{ref: terra.ReferenceResource(ed)}
}

func (ed *ElasticsearchDomain) ImportState(av io.Reader) error {
	ed.state = &elasticsearchDomainState{}
	if err := json.NewDecoder(av).Decode(ed.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ed.Type(), ed.LocalName(), err)
	}
	return nil
}

func (ed *ElasticsearchDomain) State() (*elasticsearchDomainState, bool) {
	return ed.state, ed.state != nil
}

func (ed *ElasticsearchDomain) StateMust() *elasticsearchDomainState {
	if ed.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ed.Type(), ed.LocalName()))
	}
	return ed.state
}

func (ed *ElasticsearchDomain) DependOn() terra.Reference {
	return terra.ReferenceResource(ed)
}

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
	// DependsOn contains resources that ElasticsearchDomain depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type elasticsearchDomainAttributes struct {
	ref terra.Reference
}

func (ed elasticsearchDomainAttributes) AccessPolicies() terra.StringValue {
	return terra.ReferenceString(ed.ref.Append("access_policies"))
}

func (ed elasticsearchDomainAttributes) AdvancedOptions() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](ed.ref.Append("advanced_options"))
}

func (ed elasticsearchDomainAttributes) Arn() terra.StringValue {
	return terra.ReferenceString(ed.ref.Append("arn"))
}

func (ed elasticsearchDomainAttributes) DomainId() terra.StringValue {
	return terra.ReferenceString(ed.ref.Append("domain_id"))
}

func (ed elasticsearchDomainAttributes) DomainName() terra.StringValue {
	return terra.ReferenceString(ed.ref.Append("domain_name"))
}

func (ed elasticsearchDomainAttributes) ElasticsearchVersion() terra.StringValue {
	return terra.ReferenceString(ed.ref.Append("elasticsearch_version"))
}

func (ed elasticsearchDomainAttributes) Endpoint() terra.StringValue {
	return terra.ReferenceString(ed.ref.Append("endpoint"))
}

func (ed elasticsearchDomainAttributes) Id() terra.StringValue {
	return terra.ReferenceString(ed.ref.Append("id"))
}

func (ed elasticsearchDomainAttributes) KibanaEndpoint() terra.StringValue {
	return terra.ReferenceString(ed.ref.Append("kibana_endpoint"))
}

func (ed elasticsearchDomainAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](ed.ref.Append("tags"))
}

func (ed elasticsearchDomainAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](ed.ref.Append("tags_all"))
}

func (ed elasticsearchDomainAttributes) AdvancedSecurityOptions() terra.ListValue[elasticsearchdomain.AdvancedSecurityOptionsAttributes] {
	return terra.ReferenceList[elasticsearchdomain.AdvancedSecurityOptionsAttributes](ed.ref.Append("advanced_security_options"))
}

func (ed elasticsearchDomainAttributes) AutoTuneOptions() terra.ListValue[elasticsearchdomain.AutoTuneOptionsAttributes] {
	return terra.ReferenceList[elasticsearchdomain.AutoTuneOptionsAttributes](ed.ref.Append("auto_tune_options"))
}

func (ed elasticsearchDomainAttributes) ClusterConfig() terra.ListValue[elasticsearchdomain.ClusterConfigAttributes] {
	return terra.ReferenceList[elasticsearchdomain.ClusterConfigAttributes](ed.ref.Append("cluster_config"))
}

func (ed elasticsearchDomainAttributes) CognitoOptions() terra.ListValue[elasticsearchdomain.CognitoOptionsAttributes] {
	return terra.ReferenceList[elasticsearchdomain.CognitoOptionsAttributes](ed.ref.Append("cognito_options"))
}

func (ed elasticsearchDomainAttributes) DomainEndpointOptions() terra.ListValue[elasticsearchdomain.DomainEndpointOptionsAttributes] {
	return terra.ReferenceList[elasticsearchdomain.DomainEndpointOptionsAttributes](ed.ref.Append("domain_endpoint_options"))
}

func (ed elasticsearchDomainAttributes) EbsOptions() terra.ListValue[elasticsearchdomain.EbsOptionsAttributes] {
	return terra.ReferenceList[elasticsearchdomain.EbsOptionsAttributes](ed.ref.Append("ebs_options"))
}

func (ed elasticsearchDomainAttributes) EncryptAtRest() terra.ListValue[elasticsearchdomain.EncryptAtRestAttributes] {
	return terra.ReferenceList[elasticsearchdomain.EncryptAtRestAttributes](ed.ref.Append("encrypt_at_rest"))
}

func (ed elasticsearchDomainAttributes) LogPublishingOptions() terra.SetValue[elasticsearchdomain.LogPublishingOptionsAttributes] {
	return terra.ReferenceSet[elasticsearchdomain.LogPublishingOptionsAttributes](ed.ref.Append("log_publishing_options"))
}

func (ed elasticsearchDomainAttributes) NodeToNodeEncryption() terra.ListValue[elasticsearchdomain.NodeToNodeEncryptionAttributes] {
	return terra.ReferenceList[elasticsearchdomain.NodeToNodeEncryptionAttributes](ed.ref.Append("node_to_node_encryption"))
}

func (ed elasticsearchDomainAttributes) SnapshotOptions() terra.ListValue[elasticsearchdomain.SnapshotOptionsAttributes] {
	return terra.ReferenceList[elasticsearchdomain.SnapshotOptionsAttributes](ed.ref.Append("snapshot_options"))
}

func (ed elasticsearchDomainAttributes) Timeouts() elasticsearchdomain.TimeoutsAttributes {
	return terra.ReferenceSingle[elasticsearchdomain.TimeoutsAttributes](ed.ref.Append("timeouts"))
}

func (ed elasticsearchDomainAttributes) VpcOptions() terra.ListValue[elasticsearchdomain.VpcOptionsAttributes] {
	return terra.ReferenceList[elasticsearchdomain.VpcOptionsAttributes](ed.ref.Append("vpc_options"))
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
