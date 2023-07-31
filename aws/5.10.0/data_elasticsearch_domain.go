// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataelasticsearchdomain "github.com/golingon/terraproviders/aws/5.10.0/dataelasticsearchdomain"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataElasticsearchDomain creates a new instance of [DataElasticsearchDomain].
func NewDataElasticsearchDomain(name string, args DataElasticsearchDomainArgs) *DataElasticsearchDomain {
	return &DataElasticsearchDomain{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataElasticsearchDomain)(nil)

// DataElasticsearchDomain represents the Terraform data resource aws_elasticsearch_domain.
type DataElasticsearchDomain struct {
	Name string
	Args DataElasticsearchDomainArgs
}

// DataSource returns the Terraform object type for [DataElasticsearchDomain].
func (ed *DataElasticsearchDomain) DataSource() string {
	return "aws_elasticsearch_domain"
}

// LocalName returns the local name for [DataElasticsearchDomain].
func (ed *DataElasticsearchDomain) LocalName() string {
	return ed.Name
}

// Configuration returns the configuration (args) for [DataElasticsearchDomain].
func (ed *DataElasticsearchDomain) Configuration() interface{} {
	return ed.Args
}

// Attributes returns the attributes for [DataElasticsearchDomain].
func (ed *DataElasticsearchDomain) Attributes() dataElasticsearchDomainAttributes {
	return dataElasticsearchDomainAttributes{ref: terra.ReferenceDataResource(ed)}
}

// DataElasticsearchDomainArgs contains the configurations for aws_elasticsearch_domain.
type DataElasticsearchDomainArgs struct {
	// DomainName: string, required
	DomainName terra.StringValue `hcl:"domain_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// AdvancedSecurityOptions: min=0
	AdvancedSecurityOptions []dataelasticsearchdomain.AdvancedSecurityOptions `hcl:"advanced_security_options,block" validate:"min=0"`
	// AutoTuneOptions: min=0
	AutoTuneOptions []dataelasticsearchdomain.AutoTuneOptions `hcl:"auto_tune_options,block" validate:"min=0"`
	// ClusterConfig: min=0
	ClusterConfig []dataelasticsearchdomain.ClusterConfig `hcl:"cluster_config,block" validate:"min=0"`
	// CognitoOptions: min=0
	CognitoOptions []dataelasticsearchdomain.CognitoOptions `hcl:"cognito_options,block" validate:"min=0"`
	// EbsOptions: min=0
	EbsOptions []dataelasticsearchdomain.EbsOptions `hcl:"ebs_options,block" validate:"min=0"`
	// EncryptionAtRest: min=0
	EncryptionAtRest []dataelasticsearchdomain.EncryptionAtRest `hcl:"encryption_at_rest,block" validate:"min=0"`
	// LogPublishingOptions: min=0
	LogPublishingOptions []dataelasticsearchdomain.LogPublishingOptions `hcl:"log_publishing_options,block" validate:"min=0"`
	// NodeToNodeEncryption: min=0
	NodeToNodeEncryption []dataelasticsearchdomain.NodeToNodeEncryption `hcl:"node_to_node_encryption,block" validate:"min=0"`
	// SnapshotOptions: min=0
	SnapshotOptions []dataelasticsearchdomain.SnapshotOptions `hcl:"snapshot_options,block" validate:"min=0"`
	// VpcOptions: min=0
	VpcOptions []dataelasticsearchdomain.VpcOptions `hcl:"vpc_options,block" validate:"min=0"`
}
type dataElasticsearchDomainAttributes struct {
	ref terra.Reference
}

// AccessPolicies returns a reference to field access_policies of aws_elasticsearch_domain.
func (ed dataElasticsearchDomainAttributes) AccessPolicies() terra.StringValue {
	return terra.ReferenceAsString(ed.ref.Append("access_policies"))
}

// AdvancedOptions returns a reference to field advanced_options of aws_elasticsearch_domain.
func (ed dataElasticsearchDomainAttributes) AdvancedOptions() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ed.ref.Append("advanced_options"))
}

// Arn returns a reference to field arn of aws_elasticsearch_domain.
func (ed dataElasticsearchDomainAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ed.ref.Append("arn"))
}

// Created returns a reference to field created of aws_elasticsearch_domain.
func (ed dataElasticsearchDomainAttributes) Created() terra.BoolValue {
	return terra.ReferenceAsBool(ed.ref.Append("created"))
}

// Deleted returns a reference to field deleted of aws_elasticsearch_domain.
func (ed dataElasticsearchDomainAttributes) Deleted() terra.BoolValue {
	return terra.ReferenceAsBool(ed.ref.Append("deleted"))
}

// DomainId returns a reference to field domain_id of aws_elasticsearch_domain.
func (ed dataElasticsearchDomainAttributes) DomainId() terra.StringValue {
	return terra.ReferenceAsString(ed.ref.Append("domain_id"))
}

// DomainName returns a reference to field domain_name of aws_elasticsearch_domain.
func (ed dataElasticsearchDomainAttributes) DomainName() terra.StringValue {
	return terra.ReferenceAsString(ed.ref.Append("domain_name"))
}

// ElasticsearchVersion returns a reference to field elasticsearch_version of aws_elasticsearch_domain.
func (ed dataElasticsearchDomainAttributes) ElasticsearchVersion() terra.StringValue {
	return terra.ReferenceAsString(ed.ref.Append("elasticsearch_version"))
}

// Endpoint returns a reference to field endpoint of aws_elasticsearch_domain.
func (ed dataElasticsearchDomainAttributes) Endpoint() terra.StringValue {
	return terra.ReferenceAsString(ed.ref.Append("endpoint"))
}

// Id returns a reference to field id of aws_elasticsearch_domain.
func (ed dataElasticsearchDomainAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ed.ref.Append("id"))
}

// KibanaEndpoint returns a reference to field kibana_endpoint of aws_elasticsearch_domain.
func (ed dataElasticsearchDomainAttributes) KibanaEndpoint() terra.StringValue {
	return terra.ReferenceAsString(ed.ref.Append("kibana_endpoint"))
}

// Processing returns a reference to field processing of aws_elasticsearch_domain.
func (ed dataElasticsearchDomainAttributes) Processing() terra.BoolValue {
	return terra.ReferenceAsBool(ed.ref.Append("processing"))
}

// Tags returns a reference to field tags of aws_elasticsearch_domain.
func (ed dataElasticsearchDomainAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ed.ref.Append("tags"))
}

func (ed dataElasticsearchDomainAttributes) AdvancedSecurityOptions() terra.ListValue[dataelasticsearchdomain.AdvancedSecurityOptionsAttributes] {
	return terra.ReferenceAsList[dataelasticsearchdomain.AdvancedSecurityOptionsAttributes](ed.ref.Append("advanced_security_options"))
}

func (ed dataElasticsearchDomainAttributes) AutoTuneOptions() terra.ListValue[dataelasticsearchdomain.AutoTuneOptionsAttributes] {
	return terra.ReferenceAsList[dataelasticsearchdomain.AutoTuneOptionsAttributes](ed.ref.Append("auto_tune_options"))
}

func (ed dataElasticsearchDomainAttributes) ClusterConfig() terra.ListValue[dataelasticsearchdomain.ClusterConfigAttributes] {
	return terra.ReferenceAsList[dataelasticsearchdomain.ClusterConfigAttributes](ed.ref.Append("cluster_config"))
}

func (ed dataElasticsearchDomainAttributes) CognitoOptions() terra.ListValue[dataelasticsearchdomain.CognitoOptionsAttributes] {
	return terra.ReferenceAsList[dataelasticsearchdomain.CognitoOptionsAttributes](ed.ref.Append("cognito_options"))
}

func (ed dataElasticsearchDomainAttributes) EbsOptions() terra.ListValue[dataelasticsearchdomain.EbsOptionsAttributes] {
	return terra.ReferenceAsList[dataelasticsearchdomain.EbsOptionsAttributes](ed.ref.Append("ebs_options"))
}

func (ed dataElasticsearchDomainAttributes) EncryptionAtRest() terra.ListValue[dataelasticsearchdomain.EncryptionAtRestAttributes] {
	return terra.ReferenceAsList[dataelasticsearchdomain.EncryptionAtRestAttributes](ed.ref.Append("encryption_at_rest"))
}

func (ed dataElasticsearchDomainAttributes) LogPublishingOptions() terra.SetValue[dataelasticsearchdomain.LogPublishingOptionsAttributes] {
	return terra.ReferenceAsSet[dataelasticsearchdomain.LogPublishingOptionsAttributes](ed.ref.Append("log_publishing_options"))
}

func (ed dataElasticsearchDomainAttributes) NodeToNodeEncryption() terra.ListValue[dataelasticsearchdomain.NodeToNodeEncryptionAttributes] {
	return terra.ReferenceAsList[dataelasticsearchdomain.NodeToNodeEncryptionAttributes](ed.ref.Append("node_to_node_encryption"))
}

func (ed dataElasticsearchDomainAttributes) SnapshotOptions() terra.ListValue[dataelasticsearchdomain.SnapshotOptionsAttributes] {
	return terra.ReferenceAsList[dataelasticsearchdomain.SnapshotOptionsAttributes](ed.ref.Append("snapshot_options"))
}

func (ed dataElasticsearchDomainAttributes) VpcOptions() terra.ListValue[dataelasticsearchdomain.VpcOptionsAttributes] {
	return terra.ReferenceAsList[dataelasticsearchdomain.VpcOptionsAttributes](ed.ref.Append("vpc_options"))
}
