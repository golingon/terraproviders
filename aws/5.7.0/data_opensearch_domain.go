// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataopensearchdomain "github.com/golingon/terraproviders/aws/5.7.0/dataopensearchdomain"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataOpensearchDomain creates a new instance of [DataOpensearchDomain].
func NewDataOpensearchDomain(name string, args DataOpensearchDomainArgs) *DataOpensearchDomain {
	return &DataOpensearchDomain{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataOpensearchDomain)(nil)

// DataOpensearchDomain represents the Terraform data resource aws_opensearch_domain.
type DataOpensearchDomain struct {
	Name string
	Args DataOpensearchDomainArgs
}

// DataSource returns the Terraform object type for [DataOpensearchDomain].
func (od *DataOpensearchDomain) DataSource() string {
	return "aws_opensearch_domain"
}

// LocalName returns the local name for [DataOpensearchDomain].
func (od *DataOpensearchDomain) LocalName() string {
	return od.Name
}

// Configuration returns the configuration (args) for [DataOpensearchDomain].
func (od *DataOpensearchDomain) Configuration() interface{} {
	return od.Args
}

// Attributes returns the attributes for [DataOpensearchDomain].
func (od *DataOpensearchDomain) Attributes() dataOpensearchDomainAttributes {
	return dataOpensearchDomainAttributes{ref: terra.ReferenceDataResource(od)}
}

// DataOpensearchDomainArgs contains the configurations for aws_opensearch_domain.
type DataOpensearchDomainArgs struct {
	// DomainName: string, required
	DomainName terra.StringValue `hcl:"domain_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// AdvancedSecurityOptions: min=0
	AdvancedSecurityOptions []dataopensearchdomain.AdvancedSecurityOptions `hcl:"advanced_security_options,block" validate:"min=0"`
	// AutoTuneOptions: min=0
	AutoTuneOptions []dataopensearchdomain.AutoTuneOptions `hcl:"auto_tune_options,block" validate:"min=0"`
	// ClusterConfig: min=0
	ClusterConfig []dataopensearchdomain.ClusterConfig `hcl:"cluster_config,block" validate:"min=0"`
	// CognitoOptions: min=0
	CognitoOptions []dataopensearchdomain.CognitoOptions `hcl:"cognito_options,block" validate:"min=0"`
	// EbsOptions: min=0
	EbsOptions []dataopensearchdomain.EbsOptions `hcl:"ebs_options,block" validate:"min=0"`
	// EncryptionAtRest: min=0
	EncryptionAtRest []dataopensearchdomain.EncryptionAtRest `hcl:"encryption_at_rest,block" validate:"min=0"`
	// LogPublishingOptions: min=0
	LogPublishingOptions []dataopensearchdomain.LogPublishingOptions `hcl:"log_publishing_options,block" validate:"min=0"`
	// NodeToNodeEncryption: min=0
	NodeToNodeEncryption []dataopensearchdomain.NodeToNodeEncryption `hcl:"node_to_node_encryption,block" validate:"min=0"`
	// SnapshotOptions: min=0
	SnapshotOptions []dataopensearchdomain.SnapshotOptions `hcl:"snapshot_options,block" validate:"min=0"`
	// VpcOptions: min=0
	VpcOptions []dataopensearchdomain.VpcOptions `hcl:"vpc_options,block" validate:"min=0"`
	// OffPeakWindowOptions: optional
	OffPeakWindowOptions *dataopensearchdomain.OffPeakWindowOptions `hcl:"off_peak_window_options,block"`
}
type dataOpensearchDomainAttributes struct {
	ref terra.Reference
}

// AccessPolicies returns a reference to field access_policies of aws_opensearch_domain.
func (od dataOpensearchDomainAttributes) AccessPolicies() terra.StringValue {
	return terra.ReferenceAsString(od.ref.Append("access_policies"))
}

// AdvancedOptions returns a reference to field advanced_options of aws_opensearch_domain.
func (od dataOpensearchDomainAttributes) AdvancedOptions() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](od.ref.Append("advanced_options"))
}

// Arn returns a reference to field arn of aws_opensearch_domain.
func (od dataOpensearchDomainAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(od.ref.Append("arn"))
}

// Created returns a reference to field created of aws_opensearch_domain.
func (od dataOpensearchDomainAttributes) Created() terra.BoolValue {
	return terra.ReferenceAsBool(od.ref.Append("created"))
}

// DashboardEndpoint returns a reference to field dashboard_endpoint of aws_opensearch_domain.
func (od dataOpensearchDomainAttributes) DashboardEndpoint() terra.StringValue {
	return terra.ReferenceAsString(od.ref.Append("dashboard_endpoint"))
}

// Deleted returns a reference to field deleted of aws_opensearch_domain.
func (od dataOpensearchDomainAttributes) Deleted() terra.BoolValue {
	return terra.ReferenceAsBool(od.ref.Append("deleted"))
}

// DomainId returns a reference to field domain_id of aws_opensearch_domain.
func (od dataOpensearchDomainAttributes) DomainId() terra.StringValue {
	return terra.ReferenceAsString(od.ref.Append("domain_id"))
}

// DomainName returns a reference to field domain_name of aws_opensearch_domain.
func (od dataOpensearchDomainAttributes) DomainName() terra.StringValue {
	return terra.ReferenceAsString(od.ref.Append("domain_name"))
}

// Endpoint returns a reference to field endpoint of aws_opensearch_domain.
func (od dataOpensearchDomainAttributes) Endpoint() terra.StringValue {
	return terra.ReferenceAsString(od.ref.Append("endpoint"))
}

// EngineVersion returns a reference to field engine_version of aws_opensearch_domain.
func (od dataOpensearchDomainAttributes) EngineVersion() terra.StringValue {
	return terra.ReferenceAsString(od.ref.Append("engine_version"))
}

// Id returns a reference to field id of aws_opensearch_domain.
func (od dataOpensearchDomainAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(od.ref.Append("id"))
}

// KibanaEndpoint returns a reference to field kibana_endpoint of aws_opensearch_domain.
func (od dataOpensearchDomainAttributes) KibanaEndpoint() terra.StringValue {
	return terra.ReferenceAsString(od.ref.Append("kibana_endpoint"))
}

// Processing returns a reference to field processing of aws_opensearch_domain.
func (od dataOpensearchDomainAttributes) Processing() terra.BoolValue {
	return terra.ReferenceAsBool(od.ref.Append("processing"))
}

// Tags returns a reference to field tags of aws_opensearch_domain.
func (od dataOpensearchDomainAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](od.ref.Append("tags"))
}

func (od dataOpensearchDomainAttributes) AdvancedSecurityOptions() terra.ListValue[dataopensearchdomain.AdvancedSecurityOptionsAttributes] {
	return terra.ReferenceAsList[dataopensearchdomain.AdvancedSecurityOptionsAttributes](od.ref.Append("advanced_security_options"))
}

func (od dataOpensearchDomainAttributes) AutoTuneOptions() terra.ListValue[dataopensearchdomain.AutoTuneOptionsAttributes] {
	return terra.ReferenceAsList[dataopensearchdomain.AutoTuneOptionsAttributes](od.ref.Append("auto_tune_options"))
}

func (od dataOpensearchDomainAttributes) ClusterConfig() terra.ListValue[dataopensearchdomain.ClusterConfigAttributes] {
	return terra.ReferenceAsList[dataopensearchdomain.ClusterConfigAttributes](od.ref.Append("cluster_config"))
}

func (od dataOpensearchDomainAttributes) CognitoOptions() terra.ListValue[dataopensearchdomain.CognitoOptionsAttributes] {
	return terra.ReferenceAsList[dataopensearchdomain.CognitoOptionsAttributes](od.ref.Append("cognito_options"))
}

func (od dataOpensearchDomainAttributes) EbsOptions() terra.ListValue[dataopensearchdomain.EbsOptionsAttributes] {
	return terra.ReferenceAsList[dataopensearchdomain.EbsOptionsAttributes](od.ref.Append("ebs_options"))
}

func (od dataOpensearchDomainAttributes) EncryptionAtRest() terra.ListValue[dataopensearchdomain.EncryptionAtRestAttributes] {
	return terra.ReferenceAsList[dataopensearchdomain.EncryptionAtRestAttributes](od.ref.Append("encryption_at_rest"))
}

func (od dataOpensearchDomainAttributes) LogPublishingOptions() terra.SetValue[dataopensearchdomain.LogPublishingOptionsAttributes] {
	return terra.ReferenceAsSet[dataopensearchdomain.LogPublishingOptionsAttributes](od.ref.Append("log_publishing_options"))
}

func (od dataOpensearchDomainAttributes) NodeToNodeEncryption() terra.ListValue[dataopensearchdomain.NodeToNodeEncryptionAttributes] {
	return terra.ReferenceAsList[dataopensearchdomain.NodeToNodeEncryptionAttributes](od.ref.Append("node_to_node_encryption"))
}

func (od dataOpensearchDomainAttributes) SnapshotOptions() terra.ListValue[dataopensearchdomain.SnapshotOptionsAttributes] {
	return terra.ReferenceAsList[dataopensearchdomain.SnapshotOptionsAttributes](od.ref.Append("snapshot_options"))
}

func (od dataOpensearchDomainAttributes) VpcOptions() terra.ListValue[dataopensearchdomain.VpcOptionsAttributes] {
	return terra.ReferenceAsList[dataopensearchdomain.VpcOptionsAttributes](od.ref.Append("vpc_options"))
}

func (od dataOpensearchDomainAttributes) OffPeakWindowOptions() terra.ListValue[dataopensearchdomain.OffPeakWindowOptionsAttributes] {
	return terra.ReferenceAsList[dataopensearchdomain.OffPeakWindowOptionsAttributes](od.ref.Append("off_peak_window_options"))
}