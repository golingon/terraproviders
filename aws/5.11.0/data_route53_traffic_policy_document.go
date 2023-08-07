// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataroute53trafficpolicydocument "github.com/golingon/terraproviders/aws/5.11.0/dataroute53trafficpolicydocument"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataRoute53TrafficPolicyDocument creates a new instance of [DataRoute53TrafficPolicyDocument].
func NewDataRoute53TrafficPolicyDocument(name string, args DataRoute53TrafficPolicyDocumentArgs) *DataRoute53TrafficPolicyDocument {
	return &DataRoute53TrafficPolicyDocument{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataRoute53TrafficPolicyDocument)(nil)

// DataRoute53TrafficPolicyDocument represents the Terraform data resource aws_route53_traffic_policy_document.
type DataRoute53TrafficPolicyDocument struct {
	Name string
	Args DataRoute53TrafficPolicyDocumentArgs
}

// DataSource returns the Terraform object type for [DataRoute53TrafficPolicyDocument].
func (rtpd *DataRoute53TrafficPolicyDocument) DataSource() string {
	return "aws_route53_traffic_policy_document"
}

// LocalName returns the local name for [DataRoute53TrafficPolicyDocument].
func (rtpd *DataRoute53TrafficPolicyDocument) LocalName() string {
	return rtpd.Name
}

// Configuration returns the configuration (args) for [DataRoute53TrafficPolicyDocument].
func (rtpd *DataRoute53TrafficPolicyDocument) Configuration() interface{} {
	return rtpd.Args
}

// Attributes returns the attributes for [DataRoute53TrafficPolicyDocument].
func (rtpd *DataRoute53TrafficPolicyDocument) Attributes() dataRoute53TrafficPolicyDocumentAttributes {
	return dataRoute53TrafficPolicyDocumentAttributes{ref: terra.ReferenceDataResource(rtpd)}
}

// DataRoute53TrafficPolicyDocumentArgs contains the configurations for aws_route53_traffic_policy_document.
type DataRoute53TrafficPolicyDocumentArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// RecordType: string, optional
	RecordType terra.StringValue `hcl:"record_type,attr"`
	// StartEndpoint: string, optional
	StartEndpoint terra.StringValue `hcl:"start_endpoint,attr"`
	// StartRule: string, optional
	StartRule terra.StringValue `hcl:"start_rule,attr"`
	// Version: string, optional
	Version terra.StringValue `hcl:"version,attr"`
	// Endpoint: min=0
	Endpoint []dataroute53trafficpolicydocument.Endpoint `hcl:"endpoint,block" validate:"min=0"`
	// Rule: min=0
	Rule []dataroute53trafficpolicydocument.Rule `hcl:"rule,block" validate:"min=0"`
}
type dataRoute53TrafficPolicyDocumentAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_route53_traffic_policy_document.
func (rtpd dataRoute53TrafficPolicyDocumentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(rtpd.ref.Append("id"))
}

// Json returns a reference to field json of aws_route53_traffic_policy_document.
func (rtpd dataRoute53TrafficPolicyDocumentAttributes) Json() terra.StringValue {
	return terra.ReferenceAsString(rtpd.ref.Append("json"))
}

// RecordType returns a reference to field record_type of aws_route53_traffic_policy_document.
func (rtpd dataRoute53TrafficPolicyDocumentAttributes) RecordType() terra.StringValue {
	return terra.ReferenceAsString(rtpd.ref.Append("record_type"))
}

// StartEndpoint returns a reference to field start_endpoint of aws_route53_traffic_policy_document.
func (rtpd dataRoute53TrafficPolicyDocumentAttributes) StartEndpoint() terra.StringValue {
	return terra.ReferenceAsString(rtpd.ref.Append("start_endpoint"))
}

// StartRule returns a reference to field start_rule of aws_route53_traffic_policy_document.
func (rtpd dataRoute53TrafficPolicyDocumentAttributes) StartRule() terra.StringValue {
	return terra.ReferenceAsString(rtpd.ref.Append("start_rule"))
}

// Version returns a reference to field version of aws_route53_traffic_policy_document.
func (rtpd dataRoute53TrafficPolicyDocumentAttributes) Version() terra.StringValue {
	return terra.ReferenceAsString(rtpd.ref.Append("version"))
}

func (rtpd dataRoute53TrafficPolicyDocumentAttributes) Endpoint() terra.SetValue[dataroute53trafficpolicydocument.EndpointAttributes] {
	return terra.ReferenceAsSet[dataroute53trafficpolicydocument.EndpointAttributes](rtpd.ref.Append("endpoint"))
}

func (rtpd dataRoute53TrafficPolicyDocumentAttributes) Rule() terra.SetValue[dataroute53trafficpolicydocument.RuleAttributes] {
	return terra.ReferenceAsSet[dataroute53trafficpolicydocument.RuleAttributes](rtpd.ref.Append("rule"))
}
