// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	opensearchdomain "github.com/golingon/terraproviders/aws/5.10.0/opensearchdomain"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewOpensearchDomain creates a new instance of [OpensearchDomain].
func NewOpensearchDomain(name string, args OpensearchDomainArgs) *OpensearchDomain {
	return &OpensearchDomain{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*OpensearchDomain)(nil)

// OpensearchDomain represents the Terraform resource aws_opensearch_domain.
type OpensearchDomain struct {
	Name      string
	Args      OpensearchDomainArgs
	state     *opensearchDomainState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [OpensearchDomain].
func (od *OpensearchDomain) Type() string {
	return "aws_opensearch_domain"
}

// LocalName returns the local name for [OpensearchDomain].
func (od *OpensearchDomain) LocalName() string {
	return od.Name
}

// Configuration returns the configuration (args) for [OpensearchDomain].
func (od *OpensearchDomain) Configuration() interface{} {
	return od.Args
}

// DependOn is used for other resources to depend on [OpensearchDomain].
func (od *OpensearchDomain) DependOn() terra.Reference {
	return terra.ReferenceResource(od)
}

// Dependencies returns the list of resources [OpensearchDomain] depends_on.
func (od *OpensearchDomain) Dependencies() terra.Dependencies {
	return od.DependsOn
}

// LifecycleManagement returns the lifecycle block for [OpensearchDomain].
func (od *OpensearchDomain) LifecycleManagement() *terra.Lifecycle {
	return od.Lifecycle
}

// Attributes returns the attributes for [OpensearchDomain].
func (od *OpensearchDomain) Attributes() opensearchDomainAttributes {
	return opensearchDomainAttributes{ref: terra.ReferenceResource(od)}
}

// ImportState imports the given attribute values into [OpensearchDomain]'s state.
func (od *OpensearchDomain) ImportState(av io.Reader) error {
	od.state = &opensearchDomainState{}
	if err := json.NewDecoder(av).Decode(od.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", od.Type(), od.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [OpensearchDomain] has state.
func (od *OpensearchDomain) State() (*opensearchDomainState, bool) {
	return od.state, od.state != nil
}

// StateMust returns the state for [OpensearchDomain]. Panics if the state is nil.
func (od *OpensearchDomain) StateMust() *opensearchDomainState {
	if od.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", od.Type(), od.LocalName()))
	}
	return od.state
}

// OpensearchDomainArgs contains the configurations for aws_opensearch_domain.
type OpensearchDomainArgs struct {
	// AccessPolicies: string, optional
	AccessPolicies terra.StringValue `hcl:"access_policies,attr"`
	// AdvancedOptions: map of string, optional
	AdvancedOptions terra.MapValue[terra.StringValue] `hcl:"advanced_options,attr"`
	// DomainName: string, required
	DomainName terra.StringValue `hcl:"domain_name,attr" validate:"required"`
	// EngineVersion: string, optional
	EngineVersion terra.StringValue `hcl:"engine_version,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// AdvancedSecurityOptions: optional
	AdvancedSecurityOptions *opensearchdomain.AdvancedSecurityOptions `hcl:"advanced_security_options,block"`
	// AutoTuneOptions: optional
	AutoTuneOptions *opensearchdomain.AutoTuneOptions `hcl:"auto_tune_options,block"`
	// ClusterConfig: optional
	ClusterConfig *opensearchdomain.ClusterConfig `hcl:"cluster_config,block"`
	// CognitoOptions: optional
	CognitoOptions *opensearchdomain.CognitoOptions `hcl:"cognito_options,block"`
	// DomainEndpointOptions: optional
	DomainEndpointOptions *opensearchdomain.DomainEndpointOptions `hcl:"domain_endpoint_options,block"`
	// EbsOptions: optional
	EbsOptions *opensearchdomain.EbsOptions `hcl:"ebs_options,block"`
	// EncryptAtRest: optional
	EncryptAtRest *opensearchdomain.EncryptAtRest `hcl:"encrypt_at_rest,block"`
	// LogPublishingOptions: min=0
	LogPublishingOptions []opensearchdomain.LogPublishingOptions `hcl:"log_publishing_options,block" validate:"min=0"`
	// NodeToNodeEncryption: optional
	NodeToNodeEncryption *opensearchdomain.NodeToNodeEncryption `hcl:"node_to_node_encryption,block"`
	// OffPeakWindowOptions: optional
	OffPeakWindowOptions *opensearchdomain.OffPeakWindowOptions `hcl:"off_peak_window_options,block"`
	// SnapshotOptions: optional
	SnapshotOptions *opensearchdomain.SnapshotOptions `hcl:"snapshot_options,block"`
	// Timeouts: optional
	Timeouts *opensearchdomain.Timeouts `hcl:"timeouts,block"`
	// VpcOptions: optional
	VpcOptions *opensearchdomain.VpcOptions `hcl:"vpc_options,block"`
}
type opensearchDomainAttributes struct {
	ref terra.Reference
}

// AccessPolicies returns a reference to field access_policies of aws_opensearch_domain.
func (od opensearchDomainAttributes) AccessPolicies() terra.StringValue {
	return terra.ReferenceAsString(od.ref.Append("access_policies"))
}

// AdvancedOptions returns a reference to field advanced_options of aws_opensearch_domain.
func (od opensearchDomainAttributes) AdvancedOptions() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](od.ref.Append("advanced_options"))
}

// Arn returns a reference to field arn of aws_opensearch_domain.
func (od opensearchDomainAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(od.ref.Append("arn"))
}

// DashboardEndpoint returns a reference to field dashboard_endpoint of aws_opensearch_domain.
func (od opensearchDomainAttributes) DashboardEndpoint() terra.StringValue {
	return terra.ReferenceAsString(od.ref.Append("dashboard_endpoint"))
}

// DomainId returns a reference to field domain_id of aws_opensearch_domain.
func (od opensearchDomainAttributes) DomainId() terra.StringValue {
	return terra.ReferenceAsString(od.ref.Append("domain_id"))
}

// DomainName returns a reference to field domain_name of aws_opensearch_domain.
func (od opensearchDomainAttributes) DomainName() terra.StringValue {
	return terra.ReferenceAsString(od.ref.Append("domain_name"))
}

// Endpoint returns a reference to field endpoint of aws_opensearch_domain.
func (od opensearchDomainAttributes) Endpoint() terra.StringValue {
	return terra.ReferenceAsString(od.ref.Append("endpoint"))
}

// EngineVersion returns a reference to field engine_version of aws_opensearch_domain.
func (od opensearchDomainAttributes) EngineVersion() terra.StringValue {
	return terra.ReferenceAsString(od.ref.Append("engine_version"))
}

// Id returns a reference to field id of aws_opensearch_domain.
func (od opensearchDomainAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(od.ref.Append("id"))
}

// KibanaEndpoint returns a reference to field kibana_endpoint of aws_opensearch_domain.
func (od opensearchDomainAttributes) KibanaEndpoint() terra.StringValue {
	return terra.ReferenceAsString(od.ref.Append("kibana_endpoint"))
}

// Tags returns a reference to field tags of aws_opensearch_domain.
func (od opensearchDomainAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](od.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_opensearch_domain.
func (od opensearchDomainAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](od.ref.Append("tags_all"))
}

func (od opensearchDomainAttributes) AdvancedSecurityOptions() terra.ListValue[opensearchdomain.AdvancedSecurityOptionsAttributes] {
	return terra.ReferenceAsList[opensearchdomain.AdvancedSecurityOptionsAttributes](od.ref.Append("advanced_security_options"))
}

func (od opensearchDomainAttributes) AutoTuneOptions() terra.ListValue[opensearchdomain.AutoTuneOptionsAttributes] {
	return terra.ReferenceAsList[opensearchdomain.AutoTuneOptionsAttributes](od.ref.Append("auto_tune_options"))
}

func (od opensearchDomainAttributes) ClusterConfig() terra.ListValue[opensearchdomain.ClusterConfigAttributes] {
	return terra.ReferenceAsList[opensearchdomain.ClusterConfigAttributes](od.ref.Append("cluster_config"))
}

func (od opensearchDomainAttributes) CognitoOptions() terra.ListValue[opensearchdomain.CognitoOptionsAttributes] {
	return terra.ReferenceAsList[opensearchdomain.CognitoOptionsAttributes](od.ref.Append("cognito_options"))
}

func (od opensearchDomainAttributes) DomainEndpointOptions() terra.ListValue[opensearchdomain.DomainEndpointOptionsAttributes] {
	return terra.ReferenceAsList[opensearchdomain.DomainEndpointOptionsAttributes](od.ref.Append("domain_endpoint_options"))
}

func (od opensearchDomainAttributes) EbsOptions() terra.ListValue[opensearchdomain.EbsOptionsAttributes] {
	return terra.ReferenceAsList[opensearchdomain.EbsOptionsAttributes](od.ref.Append("ebs_options"))
}

func (od opensearchDomainAttributes) EncryptAtRest() terra.ListValue[opensearchdomain.EncryptAtRestAttributes] {
	return terra.ReferenceAsList[opensearchdomain.EncryptAtRestAttributes](od.ref.Append("encrypt_at_rest"))
}

func (od opensearchDomainAttributes) LogPublishingOptions() terra.SetValue[opensearchdomain.LogPublishingOptionsAttributes] {
	return terra.ReferenceAsSet[opensearchdomain.LogPublishingOptionsAttributes](od.ref.Append("log_publishing_options"))
}

func (od opensearchDomainAttributes) NodeToNodeEncryption() terra.ListValue[opensearchdomain.NodeToNodeEncryptionAttributes] {
	return terra.ReferenceAsList[opensearchdomain.NodeToNodeEncryptionAttributes](od.ref.Append("node_to_node_encryption"))
}

func (od opensearchDomainAttributes) OffPeakWindowOptions() terra.ListValue[opensearchdomain.OffPeakWindowOptionsAttributes] {
	return terra.ReferenceAsList[opensearchdomain.OffPeakWindowOptionsAttributes](od.ref.Append("off_peak_window_options"))
}

func (od opensearchDomainAttributes) SnapshotOptions() terra.ListValue[opensearchdomain.SnapshotOptionsAttributes] {
	return terra.ReferenceAsList[opensearchdomain.SnapshotOptionsAttributes](od.ref.Append("snapshot_options"))
}

func (od opensearchDomainAttributes) Timeouts() opensearchdomain.TimeoutsAttributes {
	return terra.ReferenceAsSingle[opensearchdomain.TimeoutsAttributes](od.ref.Append("timeouts"))
}

func (od opensearchDomainAttributes) VpcOptions() terra.ListValue[opensearchdomain.VpcOptionsAttributes] {
	return terra.ReferenceAsList[opensearchdomain.VpcOptionsAttributes](od.ref.Append("vpc_options"))
}

type opensearchDomainState struct {
	AccessPolicies          string                                          `json:"access_policies"`
	AdvancedOptions         map[string]string                               `json:"advanced_options"`
	Arn                     string                                          `json:"arn"`
	DashboardEndpoint       string                                          `json:"dashboard_endpoint"`
	DomainId                string                                          `json:"domain_id"`
	DomainName              string                                          `json:"domain_name"`
	Endpoint                string                                          `json:"endpoint"`
	EngineVersion           string                                          `json:"engine_version"`
	Id                      string                                          `json:"id"`
	KibanaEndpoint          string                                          `json:"kibana_endpoint"`
	Tags                    map[string]string                               `json:"tags"`
	TagsAll                 map[string]string                               `json:"tags_all"`
	AdvancedSecurityOptions []opensearchdomain.AdvancedSecurityOptionsState `json:"advanced_security_options"`
	AutoTuneOptions         []opensearchdomain.AutoTuneOptionsState         `json:"auto_tune_options"`
	ClusterConfig           []opensearchdomain.ClusterConfigState           `json:"cluster_config"`
	CognitoOptions          []opensearchdomain.CognitoOptionsState          `json:"cognito_options"`
	DomainEndpointOptions   []opensearchdomain.DomainEndpointOptionsState   `json:"domain_endpoint_options"`
	EbsOptions              []opensearchdomain.EbsOptionsState              `json:"ebs_options"`
	EncryptAtRest           []opensearchdomain.EncryptAtRestState           `json:"encrypt_at_rest"`
	LogPublishingOptions    []opensearchdomain.LogPublishingOptionsState    `json:"log_publishing_options"`
	NodeToNodeEncryption    []opensearchdomain.NodeToNodeEncryptionState    `json:"node_to_node_encryption"`
	OffPeakWindowOptions    []opensearchdomain.OffPeakWindowOptionsState    `json:"off_peak_window_options"`
	SnapshotOptions         []opensearchdomain.SnapshotOptionsState         `json:"snapshot_options"`
	Timeouts                *opensearchdomain.TimeoutsState                 `json:"timeouts"`
	VpcOptions              []opensearchdomain.VpcOptionsState              `json:"vpc_options"`
}
