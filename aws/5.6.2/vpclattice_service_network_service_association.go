// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	vpclatticeservicenetworkserviceassociation "github.com/golingon/terraproviders/aws/5.6.2/vpclatticeservicenetworkserviceassociation"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewVpclatticeServiceNetworkServiceAssociation creates a new instance of [VpclatticeServiceNetworkServiceAssociation].
func NewVpclatticeServiceNetworkServiceAssociation(name string, args VpclatticeServiceNetworkServiceAssociationArgs) *VpclatticeServiceNetworkServiceAssociation {
	return &VpclatticeServiceNetworkServiceAssociation{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*VpclatticeServiceNetworkServiceAssociation)(nil)

// VpclatticeServiceNetworkServiceAssociation represents the Terraform resource aws_vpclattice_service_network_service_association.
type VpclatticeServiceNetworkServiceAssociation struct {
	Name      string
	Args      VpclatticeServiceNetworkServiceAssociationArgs
	state     *vpclatticeServiceNetworkServiceAssociationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [VpclatticeServiceNetworkServiceAssociation].
func (vsnsa *VpclatticeServiceNetworkServiceAssociation) Type() string {
	return "aws_vpclattice_service_network_service_association"
}

// LocalName returns the local name for [VpclatticeServiceNetworkServiceAssociation].
func (vsnsa *VpclatticeServiceNetworkServiceAssociation) LocalName() string {
	return vsnsa.Name
}

// Configuration returns the configuration (args) for [VpclatticeServiceNetworkServiceAssociation].
func (vsnsa *VpclatticeServiceNetworkServiceAssociation) Configuration() interface{} {
	return vsnsa.Args
}

// DependOn is used for other resources to depend on [VpclatticeServiceNetworkServiceAssociation].
func (vsnsa *VpclatticeServiceNetworkServiceAssociation) DependOn() terra.Reference {
	return terra.ReferenceResource(vsnsa)
}

// Dependencies returns the list of resources [VpclatticeServiceNetworkServiceAssociation] depends_on.
func (vsnsa *VpclatticeServiceNetworkServiceAssociation) Dependencies() terra.Dependencies {
	return vsnsa.DependsOn
}

// LifecycleManagement returns the lifecycle block for [VpclatticeServiceNetworkServiceAssociation].
func (vsnsa *VpclatticeServiceNetworkServiceAssociation) LifecycleManagement() *terra.Lifecycle {
	return vsnsa.Lifecycle
}

// Attributes returns the attributes for [VpclatticeServiceNetworkServiceAssociation].
func (vsnsa *VpclatticeServiceNetworkServiceAssociation) Attributes() vpclatticeServiceNetworkServiceAssociationAttributes {
	return vpclatticeServiceNetworkServiceAssociationAttributes{ref: terra.ReferenceResource(vsnsa)}
}

// ImportState imports the given attribute values into [VpclatticeServiceNetworkServiceAssociation]'s state.
func (vsnsa *VpclatticeServiceNetworkServiceAssociation) ImportState(av io.Reader) error {
	vsnsa.state = &vpclatticeServiceNetworkServiceAssociationState{}
	if err := json.NewDecoder(av).Decode(vsnsa.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", vsnsa.Type(), vsnsa.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [VpclatticeServiceNetworkServiceAssociation] has state.
func (vsnsa *VpclatticeServiceNetworkServiceAssociation) State() (*vpclatticeServiceNetworkServiceAssociationState, bool) {
	return vsnsa.state, vsnsa.state != nil
}

// StateMust returns the state for [VpclatticeServiceNetworkServiceAssociation]. Panics if the state is nil.
func (vsnsa *VpclatticeServiceNetworkServiceAssociation) StateMust() *vpclatticeServiceNetworkServiceAssociationState {
	if vsnsa.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", vsnsa.Type(), vsnsa.LocalName()))
	}
	return vsnsa.state
}

// VpclatticeServiceNetworkServiceAssociationArgs contains the configurations for aws_vpclattice_service_network_service_association.
type VpclatticeServiceNetworkServiceAssociationArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ServiceIdentifier: string, required
	ServiceIdentifier terra.StringValue `hcl:"service_identifier,attr" validate:"required"`
	// ServiceNetworkIdentifier: string, required
	ServiceNetworkIdentifier terra.StringValue `hcl:"service_network_identifier,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// DnsEntry: min=0
	DnsEntry []vpclatticeservicenetworkserviceassociation.DnsEntry `hcl:"dns_entry,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *vpclatticeservicenetworkserviceassociation.Timeouts `hcl:"timeouts,block"`
}
type vpclatticeServiceNetworkServiceAssociationAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_vpclattice_service_network_service_association.
func (vsnsa vpclatticeServiceNetworkServiceAssociationAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(vsnsa.ref.Append("arn"))
}

// CreatedBy returns a reference to field created_by of aws_vpclattice_service_network_service_association.
func (vsnsa vpclatticeServiceNetworkServiceAssociationAttributes) CreatedBy() terra.StringValue {
	return terra.ReferenceAsString(vsnsa.ref.Append("created_by"))
}

// CustomDomainName returns a reference to field custom_domain_name of aws_vpclattice_service_network_service_association.
func (vsnsa vpclatticeServiceNetworkServiceAssociationAttributes) CustomDomainName() terra.StringValue {
	return terra.ReferenceAsString(vsnsa.ref.Append("custom_domain_name"))
}

// Id returns a reference to field id of aws_vpclattice_service_network_service_association.
func (vsnsa vpclatticeServiceNetworkServiceAssociationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(vsnsa.ref.Append("id"))
}

// ServiceIdentifier returns a reference to field service_identifier of aws_vpclattice_service_network_service_association.
func (vsnsa vpclatticeServiceNetworkServiceAssociationAttributes) ServiceIdentifier() terra.StringValue {
	return terra.ReferenceAsString(vsnsa.ref.Append("service_identifier"))
}

// ServiceNetworkIdentifier returns a reference to field service_network_identifier of aws_vpclattice_service_network_service_association.
func (vsnsa vpclatticeServiceNetworkServiceAssociationAttributes) ServiceNetworkIdentifier() terra.StringValue {
	return terra.ReferenceAsString(vsnsa.ref.Append("service_network_identifier"))
}

// Status returns a reference to field status of aws_vpclattice_service_network_service_association.
func (vsnsa vpclatticeServiceNetworkServiceAssociationAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(vsnsa.ref.Append("status"))
}

// Tags returns a reference to field tags of aws_vpclattice_service_network_service_association.
func (vsnsa vpclatticeServiceNetworkServiceAssociationAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](vsnsa.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_vpclattice_service_network_service_association.
func (vsnsa vpclatticeServiceNetworkServiceAssociationAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](vsnsa.ref.Append("tags_all"))
}

func (vsnsa vpclatticeServiceNetworkServiceAssociationAttributes) DnsEntry() terra.ListValue[vpclatticeservicenetworkserviceassociation.DnsEntryAttributes] {
	return terra.ReferenceAsList[vpclatticeservicenetworkserviceassociation.DnsEntryAttributes](vsnsa.ref.Append("dns_entry"))
}

func (vsnsa vpclatticeServiceNetworkServiceAssociationAttributes) Timeouts() vpclatticeservicenetworkserviceassociation.TimeoutsAttributes {
	return terra.ReferenceAsSingle[vpclatticeservicenetworkserviceassociation.TimeoutsAttributes](vsnsa.ref.Append("timeouts"))
}

type vpclatticeServiceNetworkServiceAssociationState struct {
	Arn                      string                                                     `json:"arn"`
	CreatedBy                string                                                     `json:"created_by"`
	CustomDomainName         string                                                     `json:"custom_domain_name"`
	Id                       string                                                     `json:"id"`
	ServiceIdentifier        string                                                     `json:"service_identifier"`
	ServiceNetworkIdentifier string                                                     `json:"service_network_identifier"`
	Status                   string                                                     `json:"status"`
	Tags                     map[string]string                                          `json:"tags"`
	TagsAll                  map[string]string                                          `json:"tags_all"`
	DnsEntry                 []vpclatticeservicenetworkserviceassociation.DnsEntryState `json:"dns_entry"`
	Timeouts                 *vpclatticeservicenetworkserviceassociation.TimeoutsState  `json:"timeouts"`
}
