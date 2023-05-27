// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	computeinterconnectattachment "github.com/golingon/terraproviders/google/4.66.0/computeinterconnectattachment"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewComputeInterconnectAttachment creates a new instance of [ComputeInterconnectAttachment].
func NewComputeInterconnectAttachment(name string, args ComputeInterconnectAttachmentArgs) *ComputeInterconnectAttachment {
	return &ComputeInterconnectAttachment{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeInterconnectAttachment)(nil)

// ComputeInterconnectAttachment represents the Terraform resource google_compute_interconnect_attachment.
type ComputeInterconnectAttachment struct {
	Name      string
	Args      ComputeInterconnectAttachmentArgs
	state     *computeInterconnectAttachmentState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ComputeInterconnectAttachment].
func (cia *ComputeInterconnectAttachment) Type() string {
	return "google_compute_interconnect_attachment"
}

// LocalName returns the local name for [ComputeInterconnectAttachment].
func (cia *ComputeInterconnectAttachment) LocalName() string {
	return cia.Name
}

// Configuration returns the configuration (args) for [ComputeInterconnectAttachment].
func (cia *ComputeInterconnectAttachment) Configuration() interface{} {
	return cia.Args
}

// DependOn is used for other resources to depend on [ComputeInterconnectAttachment].
func (cia *ComputeInterconnectAttachment) DependOn() terra.Reference {
	return terra.ReferenceResource(cia)
}

// Dependencies returns the list of resources [ComputeInterconnectAttachment] depends_on.
func (cia *ComputeInterconnectAttachment) Dependencies() terra.Dependencies {
	return cia.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ComputeInterconnectAttachment].
func (cia *ComputeInterconnectAttachment) LifecycleManagement() *terra.Lifecycle {
	return cia.Lifecycle
}

// Attributes returns the attributes for [ComputeInterconnectAttachment].
func (cia *ComputeInterconnectAttachment) Attributes() computeInterconnectAttachmentAttributes {
	return computeInterconnectAttachmentAttributes{ref: terra.ReferenceResource(cia)}
}

// ImportState imports the given attribute values into [ComputeInterconnectAttachment]'s state.
func (cia *ComputeInterconnectAttachment) ImportState(av io.Reader) error {
	cia.state = &computeInterconnectAttachmentState{}
	if err := json.NewDecoder(av).Decode(cia.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cia.Type(), cia.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ComputeInterconnectAttachment] has state.
func (cia *ComputeInterconnectAttachment) State() (*computeInterconnectAttachmentState, bool) {
	return cia.state, cia.state != nil
}

// StateMust returns the state for [ComputeInterconnectAttachment]. Panics if the state is nil.
func (cia *ComputeInterconnectAttachment) StateMust() *computeInterconnectAttachmentState {
	if cia.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cia.Type(), cia.LocalName()))
	}
	return cia.state
}

// ComputeInterconnectAttachmentArgs contains the configurations for google_compute_interconnect_attachment.
type ComputeInterconnectAttachmentArgs struct {
	// AdminEnabled: bool, optional
	AdminEnabled terra.BoolValue `hcl:"admin_enabled,attr"`
	// Bandwidth: string, optional
	Bandwidth terra.StringValue `hcl:"bandwidth,attr"`
	// CandidateSubnets: list of string, optional
	CandidateSubnets terra.ListValue[terra.StringValue] `hcl:"candidate_subnets,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// EdgeAvailabilityDomain: string, optional
	EdgeAvailabilityDomain terra.StringValue `hcl:"edge_availability_domain,attr"`
	// Encryption: string, optional
	Encryption terra.StringValue `hcl:"encryption,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Interconnect: string, optional
	Interconnect terra.StringValue `hcl:"interconnect,attr"`
	// IpsecInternalAddresses: list of string, optional
	IpsecInternalAddresses terra.ListValue[terra.StringValue] `hcl:"ipsec_internal_addresses,attr"`
	// Mtu: string, optional
	Mtu terra.StringValue `hcl:"mtu,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
	// Router: string, required
	Router terra.StringValue `hcl:"router,attr" validate:"required"`
	// Type: string, optional
	Type terra.StringValue `hcl:"type,attr"`
	// VlanTag8021Q: number, optional
	VlanTag8021Q terra.NumberValue `hcl:"vlan_tag8021q,attr"`
	// PrivateInterconnectInfo: min=0
	PrivateInterconnectInfo []computeinterconnectattachment.PrivateInterconnectInfo `hcl:"private_interconnect_info,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *computeinterconnectattachment.Timeouts `hcl:"timeouts,block"`
}
type computeInterconnectAttachmentAttributes struct {
	ref terra.Reference
}

// AdminEnabled returns a reference to field admin_enabled of google_compute_interconnect_attachment.
func (cia computeInterconnectAttachmentAttributes) AdminEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(cia.ref.Append("admin_enabled"))
}

// Bandwidth returns a reference to field bandwidth of google_compute_interconnect_attachment.
func (cia computeInterconnectAttachmentAttributes) Bandwidth() terra.StringValue {
	return terra.ReferenceAsString(cia.ref.Append("bandwidth"))
}

// CandidateSubnets returns a reference to field candidate_subnets of google_compute_interconnect_attachment.
func (cia computeInterconnectAttachmentAttributes) CandidateSubnets() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](cia.ref.Append("candidate_subnets"))
}

// CloudRouterIpAddress returns a reference to field cloud_router_ip_address of google_compute_interconnect_attachment.
func (cia computeInterconnectAttachmentAttributes) CloudRouterIpAddress() terra.StringValue {
	return terra.ReferenceAsString(cia.ref.Append("cloud_router_ip_address"))
}

// CreationTimestamp returns a reference to field creation_timestamp of google_compute_interconnect_attachment.
func (cia computeInterconnectAttachmentAttributes) CreationTimestamp() terra.StringValue {
	return terra.ReferenceAsString(cia.ref.Append("creation_timestamp"))
}

// CustomerRouterIpAddress returns a reference to field customer_router_ip_address of google_compute_interconnect_attachment.
func (cia computeInterconnectAttachmentAttributes) CustomerRouterIpAddress() terra.StringValue {
	return terra.ReferenceAsString(cia.ref.Append("customer_router_ip_address"))
}

// Description returns a reference to field description of google_compute_interconnect_attachment.
func (cia computeInterconnectAttachmentAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(cia.ref.Append("description"))
}

// EdgeAvailabilityDomain returns a reference to field edge_availability_domain of google_compute_interconnect_attachment.
func (cia computeInterconnectAttachmentAttributes) EdgeAvailabilityDomain() terra.StringValue {
	return terra.ReferenceAsString(cia.ref.Append("edge_availability_domain"))
}

// Encryption returns a reference to field encryption of google_compute_interconnect_attachment.
func (cia computeInterconnectAttachmentAttributes) Encryption() terra.StringValue {
	return terra.ReferenceAsString(cia.ref.Append("encryption"))
}

// GoogleReferenceId returns a reference to field google_reference_id of google_compute_interconnect_attachment.
func (cia computeInterconnectAttachmentAttributes) GoogleReferenceId() terra.StringValue {
	return terra.ReferenceAsString(cia.ref.Append("google_reference_id"))
}

// Id returns a reference to field id of google_compute_interconnect_attachment.
func (cia computeInterconnectAttachmentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cia.ref.Append("id"))
}

// Interconnect returns a reference to field interconnect of google_compute_interconnect_attachment.
func (cia computeInterconnectAttachmentAttributes) Interconnect() terra.StringValue {
	return terra.ReferenceAsString(cia.ref.Append("interconnect"))
}

// IpsecInternalAddresses returns a reference to field ipsec_internal_addresses of google_compute_interconnect_attachment.
func (cia computeInterconnectAttachmentAttributes) IpsecInternalAddresses() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](cia.ref.Append("ipsec_internal_addresses"))
}

// Mtu returns a reference to field mtu of google_compute_interconnect_attachment.
func (cia computeInterconnectAttachmentAttributes) Mtu() terra.StringValue {
	return terra.ReferenceAsString(cia.ref.Append("mtu"))
}

// Name returns a reference to field name of google_compute_interconnect_attachment.
func (cia computeInterconnectAttachmentAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cia.ref.Append("name"))
}

// PairingKey returns a reference to field pairing_key of google_compute_interconnect_attachment.
func (cia computeInterconnectAttachmentAttributes) PairingKey() terra.StringValue {
	return terra.ReferenceAsString(cia.ref.Append("pairing_key"))
}

// PartnerAsn returns a reference to field partner_asn of google_compute_interconnect_attachment.
func (cia computeInterconnectAttachmentAttributes) PartnerAsn() terra.StringValue {
	return terra.ReferenceAsString(cia.ref.Append("partner_asn"))
}

// Project returns a reference to field project of google_compute_interconnect_attachment.
func (cia computeInterconnectAttachmentAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(cia.ref.Append("project"))
}

// Region returns a reference to field region of google_compute_interconnect_attachment.
func (cia computeInterconnectAttachmentAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(cia.ref.Append("region"))
}

// Router returns a reference to field router of google_compute_interconnect_attachment.
func (cia computeInterconnectAttachmentAttributes) Router() terra.StringValue {
	return terra.ReferenceAsString(cia.ref.Append("router"))
}

// SelfLink returns a reference to field self_link of google_compute_interconnect_attachment.
func (cia computeInterconnectAttachmentAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(cia.ref.Append("self_link"))
}

// State returns a reference to field state of google_compute_interconnect_attachment.
func (cia computeInterconnectAttachmentAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(cia.ref.Append("state"))
}

// Type returns a reference to field type of google_compute_interconnect_attachment.
func (cia computeInterconnectAttachmentAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(cia.ref.Append("type"))
}

// VlanTag8021Q returns a reference to field vlan_tag8021q of google_compute_interconnect_attachment.
func (cia computeInterconnectAttachmentAttributes) VlanTag8021Q() terra.NumberValue {
	return terra.ReferenceAsNumber(cia.ref.Append("vlan_tag8021q"))
}

func (cia computeInterconnectAttachmentAttributes) PrivateInterconnectInfo() terra.ListValue[computeinterconnectattachment.PrivateInterconnectInfoAttributes] {
	return terra.ReferenceAsList[computeinterconnectattachment.PrivateInterconnectInfoAttributes](cia.ref.Append("private_interconnect_info"))
}

func (cia computeInterconnectAttachmentAttributes) Timeouts() computeinterconnectattachment.TimeoutsAttributes {
	return terra.ReferenceAsSingle[computeinterconnectattachment.TimeoutsAttributes](cia.ref.Append("timeouts"))
}

type computeInterconnectAttachmentState struct {
	AdminEnabled            bool                                                         `json:"admin_enabled"`
	Bandwidth               string                                                       `json:"bandwidth"`
	CandidateSubnets        []string                                                     `json:"candidate_subnets"`
	CloudRouterIpAddress    string                                                       `json:"cloud_router_ip_address"`
	CreationTimestamp       string                                                       `json:"creation_timestamp"`
	CustomerRouterIpAddress string                                                       `json:"customer_router_ip_address"`
	Description             string                                                       `json:"description"`
	EdgeAvailabilityDomain  string                                                       `json:"edge_availability_domain"`
	Encryption              string                                                       `json:"encryption"`
	GoogleReferenceId       string                                                       `json:"google_reference_id"`
	Id                      string                                                       `json:"id"`
	Interconnect            string                                                       `json:"interconnect"`
	IpsecInternalAddresses  []string                                                     `json:"ipsec_internal_addresses"`
	Mtu                     string                                                       `json:"mtu"`
	Name                    string                                                       `json:"name"`
	PairingKey              string                                                       `json:"pairing_key"`
	PartnerAsn              string                                                       `json:"partner_asn"`
	Project                 string                                                       `json:"project"`
	Region                  string                                                       `json:"region"`
	Router                  string                                                       `json:"router"`
	SelfLink                string                                                       `json:"self_link"`
	State                   string                                                       `json:"state"`
	Type                    string                                                       `json:"type"`
	VlanTag8021Q            float64                                                      `json:"vlan_tag8021q"`
	PrivateInterconnectInfo []computeinterconnectattachment.PrivateInterconnectInfoState `json:"private_interconnect_info"`
	Timeouts                *computeinterconnectattachment.TimeoutsState                 `json:"timeouts"`
}
