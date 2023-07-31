// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	datanetworkmanagercorenetworkpolicydocument "github.com/golingon/terraproviders/aws/5.10.0/datanetworkmanagercorenetworkpolicydocument"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataNetworkmanagerCoreNetworkPolicyDocument creates a new instance of [DataNetworkmanagerCoreNetworkPolicyDocument].
func NewDataNetworkmanagerCoreNetworkPolicyDocument(name string, args DataNetworkmanagerCoreNetworkPolicyDocumentArgs) *DataNetworkmanagerCoreNetworkPolicyDocument {
	return &DataNetworkmanagerCoreNetworkPolicyDocument{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataNetworkmanagerCoreNetworkPolicyDocument)(nil)

// DataNetworkmanagerCoreNetworkPolicyDocument represents the Terraform data resource aws_networkmanager_core_network_policy_document.
type DataNetworkmanagerCoreNetworkPolicyDocument struct {
	Name string
	Args DataNetworkmanagerCoreNetworkPolicyDocumentArgs
}

// DataSource returns the Terraform object type for [DataNetworkmanagerCoreNetworkPolicyDocument].
func (ncnpd *DataNetworkmanagerCoreNetworkPolicyDocument) DataSource() string {
	return "aws_networkmanager_core_network_policy_document"
}

// LocalName returns the local name for [DataNetworkmanagerCoreNetworkPolicyDocument].
func (ncnpd *DataNetworkmanagerCoreNetworkPolicyDocument) LocalName() string {
	return ncnpd.Name
}

// Configuration returns the configuration (args) for [DataNetworkmanagerCoreNetworkPolicyDocument].
func (ncnpd *DataNetworkmanagerCoreNetworkPolicyDocument) Configuration() interface{} {
	return ncnpd.Args
}

// Attributes returns the attributes for [DataNetworkmanagerCoreNetworkPolicyDocument].
func (ncnpd *DataNetworkmanagerCoreNetworkPolicyDocument) Attributes() dataNetworkmanagerCoreNetworkPolicyDocumentAttributes {
	return dataNetworkmanagerCoreNetworkPolicyDocumentAttributes{ref: terra.ReferenceDataResource(ncnpd)}
}

// DataNetworkmanagerCoreNetworkPolicyDocumentArgs contains the configurations for aws_networkmanager_core_network_policy_document.
type DataNetworkmanagerCoreNetworkPolicyDocumentArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Version: string, optional
	Version terra.StringValue `hcl:"version,attr"`
	// AttachmentPolicies: min=0
	AttachmentPolicies []datanetworkmanagercorenetworkpolicydocument.AttachmentPolicies `hcl:"attachment_policies,block" validate:"min=0"`
	// CoreNetworkConfiguration: min=1
	CoreNetworkConfiguration []datanetworkmanagercorenetworkpolicydocument.CoreNetworkConfiguration `hcl:"core_network_configuration,block" validate:"min=1"`
	// SegmentActions: min=0
	SegmentActions []datanetworkmanagercorenetworkpolicydocument.SegmentActions `hcl:"segment_actions,block" validate:"min=0"`
	// Segments: min=1
	Segments []datanetworkmanagercorenetworkpolicydocument.Segments `hcl:"segments,block" validate:"min=1"`
}
type dataNetworkmanagerCoreNetworkPolicyDocumentAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_networkmanager_core_network_policy_document.
func (ncnpd dataNetworkmanagerCoreNetworkPolicyDocumentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ncnpd.ref.Append("id"))
}

// Json returns a reference to field json of aws_networkmanager_core_network_policy_document.
func (ncnpd dataNetworkmanagerCoreNetworkPolicyDocumentAttributes) Json() terra.StringValue {
	return terra.ReferenceAsString(ncnpd.ref.Append("json"))
}

// Version returns a reference to field version of aws_networkmanager_core_network_policy_document.
func (ncnpd dataNetworkmanagerCoreNetworkPolicyDocumentAttributes) Version() terra.StringValue {
	return terra.ReferenceAsString(ncnpd.ref.Append("version"))
}

func (ncnpd dataNetworkmanagerCoreNetworkPolicyDocumentAttributes) AttachmentPolicies() terra.ListValue[datanetworkmanagercorenetworkpolicydocument.AttachmentPoliciesAttributes] {
	return terra.ReferenceAsList[datanetworkmanagercorenetworkpolicydocument.AttachmentPoliciesAttributes](ncnpd.ref.Append("attachment_policies"))
}

func (ncnpd dataNetworkmanagerCoreNetworkPolicyDocumentAttributes) CoreNetworkConfiguration() terra.ListValue[datanetworkmanagercorenetworkpolicydocument.CoreNetworkConfigurationAttributes] {
	return terra.ReferenceAsList[datanetworkmanagercorenetworkpolicydocument.CoreNetworkConfigurationAttributes](ncnpd.ref.Append("core_network_configuration"))
}

func (ncnpd dataNetworkmanagerCoreNetworkPolicyDocumentAttributes) SegmentActions() terra.ListValue[datanetworkmanagercorenetworkpolicydocument.SegmentActionsAttributes] {
	return terra.ReferenceAsList[datanetworkmanagercorenetworkpolicydocument.SegmentActionsAttributes](ncnpd.ref.Append("segment_actions"))
}

func (ncnpd dataNetworkmanagerCoreNetworkPolicyDocumentAttributes) Segments() terra.ListValue[datanetworkmanagercorenetworkpolicydocument.SegmentsAttributes] {
	return terra.ReferenceAsList[datanetworkmanagercorenetworkpolicydocument.SegmentsAttributes](ncnpd.ref.Append("segments"))
}
