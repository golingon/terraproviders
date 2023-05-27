// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	computeserviceattachment "github.com/golingon/terraproviders/googlebeta/4.66.0/computeserviceattachment"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewComputeServiceAttachment creates a new instance of [ComputeServiceAttachment].
func NewComputeServiceAttachment(name string, args ComputeServiceAttachmentArgs) *ComputeServiceAttachment {
	return &ComputeServiceAttachment{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeServiceAttachment)(nil)

// ComputeServiceAttachment represents the Terraform resource google_compute_service_attachment.
type ComputeServiceAttachment struct {
	Name      string
	Args      ComputeServiceAttachmentArgs
	state     *computeServiceAttachmentState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ComputeServiceAttachment].
func (csa *ComputeServiceAttachment) Type() string {
	return "google_compute_service_attachment"
}

// LocalName returns the local name for [ComputeServiceAttachment].
func (csa *ComputeServiceAttachment) LocalName() string {
	return csa.Name
}

// Configuration returns the configuration (args) for [ComputeServiceAttachment].
func (csa *ComputeServiceAttachment) Configuration() interface{} {
	return csa.Args
}

// DependOn is used for other resources to depend on [ComputeServiceAttachment].
func (csa *ComputeServiceAttachment) DependOn() terra.Reference {
	return terra.ReferenceResource(csa)
}

// Dependencies returns the list of resources [ComputeServiceAttachment] depends_on.
func (csa *ComputeServiceAttachment) Dependencies() terra.Dependencies {
	return csa.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ComputeServiceAttachment].
func (csa *ComputeServiceAttachment) LifecycleManagement() *terra.Lifecycle {
	return csa.Lifecycle
}

// Attributes returns the attributes for [ComputeServiceAttachment].
func (csa *ComputeServiceAttachment) Attributes() computeServiceAttachmentAttributes {
	return computeServiceAttachmentAttributes{ref: terra.ReferenceResource(csa)}
}

// ImportState imports the given attribute values into [ComputeServiceAttachment]'s state.
func (csa *ComputeServiceAttachment) ImportState(av io.Reader) error {
	csa.state = &computeServiceAttachmentState{}
	if err := json.NewDecoder(av).Decode(csa.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", csa.Type(), csa.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ComputeServiceAttachment] has state.
func (csa *ComputeServiceAttachment) State() (*computeServiceAttachmentState, bool) {
	return csa.state, csa.state != nil
}

// StateMust returns the state for [ComputeServiceAttachment]. Panics if the state is nil.
func (csa *ComputeServiceAttachment) StateMust() *computeServiceAttachmentState {
	if csa.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", csa.Type(), csa.LocalName()))
	}
	return csa.state
}

// ComputeServiceAttachmentArgs contains the configurations for google_compute_service_attachment.
type ComputeServiceAttachmentArgs struct {
	// ConnectionPreference: string, required
	ConnectionPreference terra.StringValue `hcl:"connection_preference,attr" validate:"required"`
	// ConsumerRejectLists: list of string, optional
	ConsumerRejectLists terra.ListValue[terra.StringValue] `hcl:"consumer_reject_lists,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// DomainNames: list of string, optional
	DomainNames terra.ListValue[terra.StringValue] `hcl:"domain_names,attr"`
	// EnableProxyProtocol: bool, required
	EnableProxyProtocol terra.BoolValue `hcl:"enable_proxy_protocol,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// NatSubnets: list of string, required
	NatSubnets terra.ListValue[terra.StringValue] `hcl:"nat_subnets,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
	// TargetService: string, required
	TargetService terra.StringValue `hcl:"target_service,attr" validate:"required"`
	// ConnectedEndpoints: min=0
	ConnectedEndpoints []computeserviceattachment.ConnectedEndpoints `hcl:"connected_endpoints,block" validate:"min=0"`
	// ConsumerAcceptLists: min=0
	ConsumerAcceptLists []computeserviceattachment.ConsumerAcceptLists `hcl:"consumer_accept_lists,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *computeserviceattachment.Timeouts `hcl:"timeouts,block"`
}
type computeServiceAttachmentAttributes struct {
	ref terra.Reference
}

// ConnectionPreference returns a reference to field connection_preference of google_compute_service_attachment.
func (csa computeServiceAttachmentAttributes) ConnectionPreference() terra.StringValue {
	return terra.ReferenceAsString(csa.ref.Append("connection_preference"))
}

// ConsumerRejectLists returns a reference to field consumer_reject_lists of google_compute_service_attachment.
func (csa computeServiceAttachmentAttributes) ConsumerRejectLists() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](csa.ref.Append("consumer_reject_lists"))
}

// Description returns a reference to field description of google_compute_service_attachment.
func (csa computeServiceAttachmentAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(csa.ref.Append("description"))
}

// DomainNames returns a reference to field domain_names of google_compute_service_attachment.
func (csa computeServiceAttachmentAttributes) DomainNames() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](csa.ref.Append("domain_names"))
}

// EnableProxyProtocol returns a reference to field enable_proxy_protocol of google_compute_service_attachment.
func (csa computeServiceAttachmentAttributes) EnableProxyProtocol() terra.BoolValue {
	return terra.ReferenceAsBool(csa.ref.Append("enable_proxy_protocol"))
}

// Fingerprint returns a reference to field fingerprint of google_compute_service_attachment.
func (csa computeServiceAttachmentAttributes) Fingerprint() terra.StringValue {
	return terra.ReferenceAsString(csa.ref.Append("fingerprint"))
}

// Id returns a reference to field id of google_compute_service_attachment.
func (csa computeServiceAttachmentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(csa.ref.Append("id"))
}

// Name returns a reference to field name of google_compute_service_attachment.
func (csa computeServiceAttachmentAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(csa.ref.Append("name"))
}

// NatSubnets returns a reference to field nat_subnets of google_compute_service_attachment.
func (csa computeServiceAttachmentAttributes) NatSubnets() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](csa.ref.Append("nat_subnets"))
}

// Project returns a reference to field project of google_compute_service_attachment.
func (csa computeServiceAttachmentAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(csa.ref.Append("project"))
}

// Region returns a reference to field region of google_compute_service_attachment.
func (csa computeServiceAttachmentAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(csa.ref.Append("region"))
}

// SelfLink returns a reference to field self_link of google_compute_service_attachment.
func (csa computeServiceAttachmentAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(csa.ref.Append("self_link"))
}

// TargetService returns a reference to field target_service of google_compute_service_attachment.
func (csa computeServiceAttachmentAttributes) TargetService() terra.StringValue {
	return terra.ReferenceAsString(csa.ref.Append("target_service"))
}

func (csa computeServiceAttachmentAttributes) ConnectedEndpoints() terra.ListValue[computeserviceattachment.ConnectedEndpointsAttributes] {
	return terra.ReferenceAsList[computeserviceattachment.ConnectedEndpointsAttributes](csa.ref.Append("connected_endpoints"))
}

func (csa computeServiceAttachmentAttributes) ConsumerAcceptLists() terra.ListValue[computeserviceattachment.ConsumerAcceptListsAttributes] {
	return terra.ReferenceAsList[computeserviceattachment.ConsumerAcceptListsAttributes](csa.ref.Append("consumer_accept_lists"))
}

func (csa computeServiceAttachmentAttributes) Timeouts() computeserviceattachment.TimeoutsAttributes {
	return terra.ReferenceAsSingle[computeserviceattachment.TimeoutsAttributes](csa.ref.Append("timeouts"))
}

type computeServiceAttachmentState struct {
	ConnectionPreference string                                              `json:"connection_preference"`
	ConsumerRejectLists  []string                                            `json:"consumer_reject_lists"`
	Description          string                                              `json:"description"`
	DomainNames          []string                                            `json:"domain_names"`
	EnableProxyProtocol  bool                                                `json:"enable_proxy_protocol"`
	Fingerprint          string                                              `json:"fingerprint"`
	Id                   string                                              `json:"id"`
	Name                 string                                              `json:"name"`
	NatSubnets           []string                                            `json:"nat_subnets"`
	Project              string                                              `json:"project"`
	Region               string                                              `json:"region"`
	SelfLink             string                                              `json:"self_link"`
	TargetService        string                                              `json:"target_service"`
	ConnectedEndpoints   []computeserviceattachment.ConnectedEndpointsState  `json:"connected_endpoints"`
	ConsumerAcceptLists  []computeserviceattachment.ConsumerAcceptListsState `json:"consumer_accept_lists"`
	Timeouts             *computeserviceattachment.TimeoutsState             `json:"timeouts"`
}
