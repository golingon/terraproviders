// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	vpcipamresourcediscovery "github.com/golingon/terraproviders/aws/4.60.0/vpcipamresourcediscovery"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewVpcIpamResourceDiscovery(name string, args VpcIpamResourceDiscoveryArgs) *VpcIpamResourceDiscovery {
	return &VpcIpamResourceDiscovery{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*VpcIpamResourceDiscovery)(nil)

type VpcIpamResourceDiscovery struct {
	Name  string
	Args  VpcIpamResourceDiscoveryArgs
	state *vpcIpamResourceDiscoveryState
}

func (vird *VpcIpamResourceDiscovery) Type() string {
	return "aws_vpc_ipam_resource_discovery"
}

func (vird *VpcIpamResourceDiscovery) LocalName() string {
	return vird.Name
}

func (vird *VpcIpamResourceDiscovery) Configuration() interface{} {
	return vird.Args
}

func (vird *VpcIpamResourceDiscovery) Attributes() vpcIpamResourceDiscoveryAttributes {
	return vpcIpamResourceDiscoveryAttributes{ref: terra.ReferenceResource(vird)}
}

func (vird *VpcIpamResourceDiscovery) ImportState(av io.Reader) error {
	vird.state = &vpcIpamResourceDiscoveryState{}
	if err := json.NewDecoder(av).Decode(vird.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", vird.Type(), vird.LocalName(), err)
	}
	return nil
}

func (vird *VpcIpamResourceDiscovery) State() (*vpcIpamResourceDiscoveryState, bool) {
	return vird.state, vird.state != nil
}

func (vird *VpcIpamResourceDiscovery) StateMust() *vpcIpamResourceDiscoveryState {
	if vird.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", vird.Type(), vird.LocalName()))
	}
	return vird.state
}

func (vird *VpcIpamResourceDiscovery) DependOn() terra.Reference {
	return terra.ReferenceResource(vird)
}

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
	// DependsOn contains resources that VpcIpamResourceDiscovery depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type vpcIpamResourceDiscoveryAttributes struct {
	ref terra.Reference
}

func (vird vpcIpamResourceDiscoveryAttributes) Arn() terra.StringValue {
	return terra.ReferenceString(vird.ref.Append("arn"))
}

func (vird vpcIpamResourceDiscoveryAttributes) Description() terra.StringValue {
	return terra.ReferenceString(vird.ref.Append("description"))
}

func (vird vpcIpamResourceDiscoveryAttributes) Id() terra.StringValue {
	return terra.ReferenceString(vird.ref.Append("id"))
}

func (vird vpcIpamResourceDiscoveryAttributes) IpamResourceDiscoveryRegion() terra.StringValue {
	return terra.ReferenceString(vird.ref.Append("ipam_resource_discovery_region"))
}

func (vird vpcIpamResourceDiscoveryAttributes) IsDefault() terra.BoolValue {
	return terra.ReferenceBool(vird.ref.Append("is_default"))
}

func (vird vpcIpamResourceDiscoveryAttributes) OwnerId() terra.StringValue {
	return terra.ReferenceString(vird.ref.Append("owner_id"))
}

func (vird vpcIpamResourceDiscoveryAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](vird.ref.Append("tags"))
}

func (vird vpcIpamResourceDiscoveryAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](vird.ref.Append("tags_all"))
}

func (vird vpcIpamResourceDiscoveryAttributes) OperatingRegions() terra.SetValue[vpcipamresourcediscovery.OperatingRegionsAttributes] {
	return terra.ReferenceSet[vpcipamresourcediscovery.OperatingRegionsAttributes](vird.ref.Append("operating_regions"))
}

func (vird vpcIpamResourceDiscoveryAttributes) Timeouts() vpcipamresourcediscovery.TimeoutsAttributes {
	return terra.ReferenceSingle[vpcipamresourcediscovery.TimeoutsAttributes](vird.ref.Append("timeouts"))
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
