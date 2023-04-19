// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	vpcipamresourcediscoveryassociation "github.com/golingon/terraproviders/aws/4.63.0/vpcipamresourcediscoveryassociation"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewVpcIpamResourceDiscoveryAssociation creates a new instance of [VpcIpamResourceDiscoveryAssociation].
func NewVpcIpamResourceDiscoveryAssociation(name string, args VpcIpamResourceDiscoveryAssociationArgs) *VpcIpamResourceDiscoveryAssociation {
	return &VpcIpamResourceDiscoveryAssociation{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*VpcIpamResourceDiscoveryAssociation)(nil)

// VpcIpamResourceDiscoveryAssociation represents the Terraform resource aws_vpc_ipam_resource_discovery_association.
type VpcIpamResourceDiscoveryAssociation struct {
	Name      string
	Args      VpcIpamResourceDiscoveryAssociationArgs
	state     *vpcIpamResourceDiscoveryAssociationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [VpcIpamResourceDiscoveryAssociation].
func (virda *VpcIpamResourceDiscoveryAssociation) Type() string {
	return "aws_vpc_ipam_resource_discovery_association"
}

// LocalName returns the local name for [VpcIpamResourceDiscoveryAssociation].
func (virda *VpcIpamResourceDiscoveryAssociation) LocalName() string {
	return virda.Name
}

// Configuration returns the configuration (args) for [VpcIpamResourceDiscoveryAssociation].
func (virda *VpcIpamResourceDiscoveryAssociation) Configuration() interface{} {
	return virda.Args
}

// DependOn is used for other resources to depend on [VpcIpamResourceDiscoveryAssociation].
func (virda *VpcIpamResourceDiscoveryAssociation) DependOn() terra.Reference {
	return terra.ReferenceResource(virda)
}

// Dependencies returns the list of resources [VpcIpamResourceDiscoveryAssociation] depends_on.
func (virda *VpcIpamResourceDiscoveryAssociation) Dependencies() terra.Dependencies {
	return virda.DependsOn
}

// LifecycleManagement returns the lifecycle block for [VpcIpamResourceDiscoveryAssociation].
func (virda *VpcIpamResourceDiscoveryAssociation) LifecycleManagement() *terra.Lifecycle {
	return virda.Lifecycle
}

// Attributes returns the attributes for [VpcIpamResourceDiscoveryAssociation].
func (virda *VpcIpamResourceDiscoveryAssociation) Attributes() vpcIpamResourceDiscoveryAssociationAttributes {
	return vpcIpamResourceDiscoveryAssociationAttributes{ref: terra.ReferenceResource(virda)}
}

// ImportState imports the given attribute values into [VpcIpamResourceDiscoveryAssociation]'s state.
func (virda *VpcIpamResourceDiscoveryAssociation) ImportState(av io.Reader) error {
	virda.state = &vpcIpamResourceDiscoveryAssociationState{}
	if err := json.NewDecoder(av).Decode(virda.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", virda.Type(), virda.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [VpcIpamResourceDiscoveryAssociation] has state.
func (virda *VpcIpamResourceDiscoveryAssociation) State() (*vpcIpamResourceDiscoveryAssociationState, bool) {
	return virda.state, virda.state != nil
}

// StateMust returns the state for [VpcIpamResourceDiscoveryAssociation]. Panics if the state is nil.
func (virda *VpcIpamResourceDiscoveryAssociation) StateMust() *vpcIpamResourceDiscoveryAssociationState {
	if virda.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", virda.Type(), virda.LocalName()))
	}
	return virda.state
}

// VpcIpamResourceDiscoveryAssociationArgs contains the configurations for aws_vpc_ipam_resource_discovery_association.
type VpcIpamResourceDiscoveryAssociationArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IpamId: string, required
	IpamId terra.StringValue `hcl:"ipam_id,attr" validate:"required"`
	// IpamResourceDiscoveryId: string, required
	IpamResourceDiscoveryId terra.StringValue `hcl:"ipam_resource_discovery_id,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// Timeouts: optional
	Timeouts *vpcipamresourcediscoveryassociation.Timeouts `hcl:"timeouts,block"`
}
type vpcIpamResourceDiscoveryAssociationAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_vpc_ipam_resource_discovery_association.
func (virda vpcIpamResourceDiscoveryAssociationAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(virda.ref.Append("arn"))
}

// Id returns a reference to field id of aws_vpc_ipam_resource_discovery_association.
func (virda vpcIpamResourceDiscoveryAssociationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(virda.ref.Append("id"))
}

// IpamArn returns a reference to field ipam_arn of aws_vpc_ipam_resource_discovery_association.
func (virda vpcIpamResourceDiscoveryAssociationAttributes) IpamArn() terra.StringValue {
	return terra.ReferenceAsString(virda.ref.Append("ipam_arn"))
}

// IpamId returns a reference to field ipam_id of aws_vpc_ipam_resource_discovery_association.
func (virda vpcIpamResourceDiscoveryAssociationAttributes) IpamId() terra.StringValue {
	return terra.ReferenceAsString(virda.ref.Append("ipam_id"))
}

// IpamRegion returns a reference to field ipam_region of aws_vpc_ipam_resource_discovery_association.
func (virda vpcIpamResourceDiscoveryAssociationAttributes) IpamRegion() terra.StringValue {
	return terra.ReferenceAsString(virda.ref.Append("ipam_region"))
}

// IpamResourceDiscoveryId returns a reference to field ipam_resource_discovery_id of aws_vpc_ipam_resource_discovery_association.
func (virda vpcIpamResourceDiscoveryAssociationAttributes) IpamResourceDiscoveryId() terra.StringValue {
	return terra.ReferenceAsString(virda.ref.Append("ipam_resource_discovery_id"))
}

// IsDefault returns a reference to field is_default of aws_vpc_ipam_resource_discovery_association.
func (virda vpcIpamResourceDiscoveryAssociationAttributes) IsDefault() terra.BoolValue {
	return terra.ReferenceAsBool(virda.ref.Append("is_default"))
}

// OwnerId returns a reference to field owner_id of aws_vpc_ipam_resource_discovery_association.
func (virda vpcIpamResourceDiscoveryAssociationAttributes) OwnerId() terra.StringValue {
	return terra.ReferenceAsString(virda.ref.Append("owner_id"))
}

// State returns a reference to field state of aws_vpc_ipam_resource_discovery_association.
func (virda vpcIpamResourceDiscoveryAssociationAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(virda.ref.Append("state"))
}

// Tags returns a reference to field tags of aws_vpc_ipam_resource_discovery_association.
func (virda vpcIpamResourceDiscoveryAssociationAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](virda.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_vpc_ipam_resource_discovery_association.
func (virda vpcIpamResourceDiscoveryAssociationAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](virda.ref.Append("tags_all"))
}

func (virda vpcIpamResourceDiscoveryAssociationAttributes) Timeouts() vpcipamresourcediscoveryassociation.TimeoutsAttributes {
	return terra.ReferenceAsSingle[vpcipamresourcediscoveryassociation.TimeoutsAttributes](virda.ref.Append("timeouts"))
}

type vpcIpamResourceDiscoveryAssociationState struct {
	Arn                     string                                             `json:"arn"`
	Id                      string                                             `json:"id"`
	IpamArn                 string                                             `json:"ipam_arn"`
	IpamId                  string                                             `json:"ipam_id"`
	IpamRegion              string                                             `json:"ipam_region"`
	IpamResourceDiscoveryId string                                             `json:"ipam_resource_discovery_id"`
	IsDefault               bool                                               `json:"is_default"`
	OwnerId                 string                                             `json:"owner_id"`
	State                   string                                             `json:"state"`
	Tags                    map[string]string                                  `json:"tags"`
	TagsAll                 map[string]string                                  `json:"tags_all"`
	Timeouts                *vpcipamresourcediscoveryassociation.TimeoutsState `json:"timeouts"`
}