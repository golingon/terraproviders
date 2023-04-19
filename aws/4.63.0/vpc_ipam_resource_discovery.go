// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	vpcipamresourcediscovery "github.com/golingon/terraproviders/aws/4.63.0/vpcipamresourcediscovery"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewVpcIpamResourceDiscovery creates a new instance of [VpcIpamResourceDiscovery].
func NewVpcIpamResourceDiscovery(name string, args VpcIpamResourceDiscoveryArgs) *VpcIpamResourceDiscovery {
	return &VpcIpamResourceDiscovery{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*VpcIpamResourceDiscovery)(nil)

// VpcIpamResourceDiscovery represents the Terraform resource aws_vpc_ipam_resource_discovery.
type VpcIpamResourceDiscovery struct {
	Name      string
	Args      VpcIpamResourceDiscoveryArgs
	state     *vpcIpamResourceDiscoveryState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [VpcIpamResourceDiscovery].
func (vird *VpcIpamResourceDiscovery) Type() string {
	return "aws_vpc_ipam_resource_discovery"
}

// LocalName returns the local name for [VpcIpamResourceDiscovery].
func (vird *VpcIpamResourceDiscovery) LocalName() string {
	return vird.Name
}

// Configuration returns the configuration (args) for [VpcIpamResourceDiscovery].
func (vird *VpcIpamResourceDiscovery) Configuration() interface{} {
	return vird.Args
}

// DependOn is used for other resources to depend on [VpcIpamResourceDiscovery].
func (vird *VpcIpamResourceDiscovery) DependOn() terra.Reference {
	return terra.ReferenceResource(vird)
}

// Dependencies returns the list of resources [VpcIpamResourceDiscovery] depends_on.
func (vird *VpcIpamResourceDiscovery) Dependencies() terra.Dependencies {
	return vird.DependsOn
}

// LifecycleManagement returns the lifecycle block for [VpcIpamResourceDiscovery].
func (vird *VpcIpamResourceDiscovery) LifecycleManagement() *terra.Lifecycle {
	return vird.Lifecycle
}

// Attributes returns the attributes for [VpcIpamResourceDiscovery].
func (vird *VpcIpamResourceDiscovery) Attributes() vpcIpamResourceDiscoveryAttributes {
	return vpcIpamResourceDiscoveryAttributes{ref: terra.ReferenceResource(vird)}
}

// ImportState imports the given attribute values into [VpcIpamResourceDiscovery]'s state.
func (vird *VpcIpamResourceDiscovery) ImportState(av io.Reader) error {
	vird.state = &vpcIpamResourceDiscoveryState{}
	if err := json.NewDecoder(av).Decode(vird.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", vird.Type(), vird.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [VpcIpamResourceDiscovery] has state.
func (vird *VpcIpamResourceDiscovery) State() (*vpcIpamResourceDiscoveryState, bool) {
	return vird.state, vird.state != nil
}

// StateMust returns the state for [VpcIpamResourceDiscovery]. Panics if the state is nil.
func (vird *VpcIpamResourceDiscovery) StateMust() *vpcIpamResourceDiscoveryState {
	if vird.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", vird.Type(), vird.LocalName()))
	}
	return vird.state
}

// VpcIpamResourceDiscoveryArgs contains the configurations for aws_vpc_ipam_resource_discovery.
type VpcIpamResourceDiscoveryArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// OperatingRegions: min=1
	OperatingRegions []vpcipamresourcediscovery.OperatingRegions `hcl:"operating_regions,block" validate:"min=1"`
	// Timeouts: optional
	Timeouts *vpcipamresourcediscovery.Timeouts `hcl:"timeouts,block"`
}
type vpcIpamResourceDiscoveryAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_vpc_ipam_resource_discovery.
func (vird vpcIpamResourceDiscoveryAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(vird.ref.Append("arn"))
}

// Description returns a reference to field description of aws_vpc_ipam_resource_discovery.
func (vird vpcIpamResourceDiscoveryAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(vird.ref.Append("description"))
}

// Id returns a reference to field id of aws_vpc_ipam_resource_discovery.
func (vird vpcIpamResourceDiscoveryAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(vird.ref.Append("id"))
}

// IpamResourceDiscoveryRegion returns a reference to field ipam_resource_discovery_region of aws_vpc_ipam_resource_discovery.
func (vird vpcIpamResourceDiscoveryAttributes) IpamResourceDiscoveryRegion() terra.StringValue {
	return terra.ReferenceAsString(vird.ref.Append("ipam_resource_discovery_region"))
}

// IsDefault returns a reference to field is_default of aws_vpc_ipam_resource_discovery.
func (vird vpcIpamResourceDiscoveryAttributes) IsDefault() terra.BoolValue {
	return terra.ReferenceAsBool(vird.ref.Append("is_default"))
}

// OwnerId returns a reference to field owner_id of aws_vpc_ipam_resource_discovery.
func (vird vpcIpamResourceDiscoveryAttributes) OwnerId() terra.StringValue {
	return terra.ReferenceAsString(vird.ref.Append("owner_id"))
}

// Tags returns a reference to field tags of aws_vpc_ipam_resource_discovery.
func (vird vpcIpamResourceDiscoveryAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](vird.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_vpc_ipam_resource_discovery.
func (vird vpcIpamResourceDiscoveryAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](vird.ref.Append("tags_all"))
}

func (vird vpcIpamResourceDiscoveryAttributes) OperatingRegions() terra.SetValue[vpcipamresourcediscovery.OperatingRegionsAttributes] {
	return terra.ReferenceAsSet[vpcipamresourcediscovery.OperatingRegionsAttributes](vird.ref.Append("operating_regions"))
}

func (vird vpcIpamResourceDiscoveryAttributes) Timeouts() vpcipamresourcediscovery.TimeoutsAttributes {
	return terra.ReferenceAsSingle[vpcipamresourcediscovery.TimeoutsAttributes](vird.ref.Append("timeouts"))
}

type vpcIpamResourceDiscoveryState struct {
	Arn                         string                                           `json:"arn"`
	Description                 string                                           `json:"description"`
	Id                          string                                           `json:"id"`
	IpamResourceDiscoveryRegion string                                           `json:"ipam_resource_discovery_region"`
	IsDefault                   bool                                             `json:"is_default"`
	OwnerId                     string                                           `json:"owner_id"`
	Tags                        map[string]string                                `json:"tags"`
	TagsAll                     map[string]string                                `json:"tags_all"`
	OperatingRegions            []vpcipamresourcediscovery.OperatingRegionsState `json:"operating_regions"`
	Timeouts                    *vpcipamresourcediscovery.TimeoutsState          `json:"timeouts"`
}
