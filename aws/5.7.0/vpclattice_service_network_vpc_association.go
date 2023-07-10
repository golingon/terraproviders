// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	vpclatticeservicenetworkvpcassociation "github.com/golingon/terraproviders/aws/5.7.0/vpclatticeservicenetworkvpcassociation"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewVpclatticeServiceNetworkVpcAssociation creates a new instance of [VpclatticeServiceNetworkVpcAssociation].
func NewVpclatticeServiceNetworkVpcAssociation(name string, args VpclatticeServiceNetworkVpcAssociationArgs) *VpclatticeServiceNetworkVpcAssociation {
	return &VpclatticeServiceNetworkVpcAssociation{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*VpclatticeServiceNetworkVpcAssociation)(nil)

// VpclatticeServiceNetworkVpcAssociation represents the Terraform resource aws_vpclattice_service_network_vpc_association.
type VpclatticeServiceNetworkVpcAssociation struct {
	Name      string
	Args      VpclatticeServiceNetworkVpcAssociationArgs
	state     *vpclatticeServiceNetworkVpcAssociationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [VpclatticeServiceNetworkVpcAssociation].
func (vsnva *VpclatticeServiceNetworkVpcAssociation) Type() string {
	return "aws_vpclattice_service_network_vpc_association"
}

// LocalName returns the local name for [VpclatticeServiceNetworkVpcAssociation].
func (vsnva *VpclatticeServiceNetworkVpcAssociation) LocalName() string {
	return vsnva.Name
}

// Configuration returns the configuration (args) for [VpclatticeServiceNetworkVpcAssociation].
func (vsnva *VpclatticeServiceNetworkVpcAssociation) Configuration() interface{} {
	return vsnva.Args
}

// DependOn is used for other resources to depend on [VpclatticeServiceNetworkVpcAssociation].
func (vsnva *VpclatticeServiceNetworkVpcAssociation) DependOn() terra.Reference {
	return terra.ReferenceResource(vsnva)
}

// Dependencies returns the list of resources [VpclatticeServiceNetworkVpcAssociation] depends_on.
func (vsnva *VpclatticeServiceNetworkVpcAssociation) Dependencies() terra.Dependencies {
	return vsnva.DependsOn
}

// LifecycleManagement returns the lifecycle block for [VpclatticeServiceNetworkVpcAssociation].
func (vsnva *VpclatticeServiceNetworkVpcAssociation) LifecycleManagement() *terra.Lifecycle {
	return vsnva.Lifecycle
}

// Attributes returns the attributes for [VpclatticeServiceNetworkVpcAssociation].
func (vsnva *VpclatticeServiceNetworkVpcAssociation) Attributes() vpclatticeServiceNetworkVpcAssociationAttributes {
	return vpclatticeServiceNetworkVpcAssociationAttributes{ref: terra.ReferenceResource(vsnva)}
}

// ImportState imports the given attribute values into [VpclatticeServiceNetworkVpcAssociation]'s state.
func (vsnva *VpclatticeServiceNetworkVpcAssociation) ImportState(av io.Reader) error {
	vsnva.state = &vpclatticeServiceNetworkVpcAssociationState{}
	if err := json.NewDecoder(av).Decode(vsnva.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", vsnva.Type(), vsnva.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [VpclatticeServiceNetworkVpcAssociation] has state.
func (vsnva *VpclatticeServiceNetworkVpcAssociation) State() (*vpclatticeServiceNetworkVpcAssociationState, bool) {
	return vsnva.state, vsnva.state != nil
}

// StateMust returns the state for [VpclatticeServiceNetworkVpcAssociation]. Panics if the state is nil.
func (vsnva *VpclatticeServiceNetworkVpcAssociation) StateMust() *vpclatticeServiceNetworkVpcAssociationState {
	if vsnva.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", vsnva.Type(), vsnva.LocalName()))
	}
	return vsnva.state
}

// VpclatticeServiceNetworkVpcAssociationArgs contains the configurations for aws_vpclattice_service_network_vpc_association.
type VpclatticeServiceNetworkVpcAssociationArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// SecurityGroupIds: list of string, optional
	SecurityGroupIds terra.ListValue[terra.StringValue] `hcl:"security_group_ids,attr"`
	// ServiceNetworkIdentifier: string, required
	ServiceNetworkIdentifier terra.StringValue `hcl:"service_network_identifier,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// VpcIdentifier: string, required
	VpcIdentifier terra.StringValue `hcl:"vpc_identifier,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *vpclatticeservicenetworkvpcassociation.Timeouts `hcl:"timeouts,block"`
}
type vpclatticeServiceNetworkVpcAssociationAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_vpclattice_service_network_vpc_association.
func (vsnva vpclatticeServiceNetworkVpcAssociationAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(vsnva.ref.Append("arn"))
}

// CreatedBy returns a reference to field created_by of aws_vpclattice_service_network_vpc_association.
func (vsnva vpclatticeServiceNetworkVpcAssociationAttributes) CreatedBy() terra.StringValue {
	return terra.ReferenceAsString(vsnva.ref.Append("created_by"))
}

// Id returns a reference to field id of aws_vpclattice_service_network_vpc_association.
func (vsnva vpclatticeServiceNetworkVpcAssociationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(vsnva.ref.Append("id"))
}

// SecurityGroupIds returns a reference to field security_group_ids of aws_vpclattice_service_network_vpc_association.
func (vsnva vpclatticeServiceNetworkVpcAssociationAttributes) SecurityGroupIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](vsnva.ref.Append("security_group_ids"))
}

// ServiceNetworkIdentifier returns a reference to field service_network_identifier of aws_vpclattice_service_network_vpc_association.
func (vsnva vpclatticeServiceNetworkVpcAssociationAttributes) ServiceNetworkIdentifier() terra.StringValue {
	return terra.ReferenceAsString(vsnva.ref.Append("service_network_identifier"))
}

// Status returns a reference to field status of aws_vpclattice_service_network_vpc_association.
func (vsnva vpclatticeServiceNetworkVpcAssociationAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(vsnva.ref.Append("status"))
}

// Tags returns a reference to field tags of aws_vpclattice_service_network_vpc_association.
func (vsnva vpclatticeServiceNetworkVpcAssociationAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](vsnva.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_vpclattice_service_network_vpc_association.
func (vsnva vpclatticeServiceNetworkVpcAssociationAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](vsnva.ref.Append("tags_all"))
}

// VpcIdentifier returns a reference to field vpc_identifier of aws_vpclattice_service_network_vpc_association.
func (vsnva vpclatticeServiceNetworkVpcAssociationAttributes) VpcIdentifier() terra.StringValue {
	return terra.ReferenceAsString(vsnva.ref.Append("vpc_identifier"))
}

func (vsnva vpclatticeServiceNetworkVpcAssociationAttributes) Timeouts() vpclatticeservicenetworkvpcassociation.TimeoutsAttributes {
	return terra.ReferenceAsSingle[vpclatticeservicenetworkvpcassociation.TimeoutsAttributes](vsnva.ref.Append("timeouts"))
}

type vpclatticeServiceNetworkVpcAssociationState struct {
	Arn                      string                                                `json:"arn"`
	CreatedBy                string                                                `json:"created_by"`
	Id                       string                                                `json:"id"`
	SecurityGroupIds         []string                                              `json:"security_group_ids"`
	ServiceNetworkIdentifier string                                                `json:"service_network_identifier"`
	Status                   string                                                `json:"status"`
	Tags                     map[string]string                                     `json:"tags"`
	TagsAll                  map[string]string                                     `json:"tags_all"`
	VpcIdentifier            string                                                `json:"vpc_identifier"`
	Timeouts                 *vpclatticeservicenetworkvpcassociation.TimeoutsState `json:"timeouts"`
}
