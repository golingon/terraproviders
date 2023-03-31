// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	computeregionnetworkendpointgroup "github.com/golingon/terraproviders/google/4.59.0/computeregionnetworkendpointgroup"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewComputeRegionNetworkEndpointGroup(name string, args ComputeRegionNetworkEndpointGroupArgs) *ComputeRegionNetworkEndpointGroup {
	return &ComputeRegionNetworkEndpointGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeRegionNetworkEndpointGroup)(nil)

type ComputeRegionNetworkEndpointGroup struct {
	Name  string
	Args  ComputeRegionNetworkEndpointGroupArgs
	state *computeRegionNetworkEndpointGroupState
}

func (crneg *ComputeRegionNetworkEndpointGroup) Type() string {
	return "google_compute_region_network_endpoint_group"
}

func (crneg *ComputeRegionNetworkEndpointGroup) LocalName() string {
	return crneg.Name
}

func (crneg *ComputeRegionNetworkEndpointGroup) Configuration() interface{} {
	return crneg.Args
}

func (crneg *ComputeRegionNetworkEndpointGroup) Attributes() computeRegionNetworkEndpointGroupAttributes {
	return computeRegionNetworkEndpointGroupAttributes{ref: terra.ReferenceResource(crneg)}
}

func (crneg *ComputeRegionNetworkEndpointGroup) ImportState(av io.Reader) error {
	crneg.state = &computeRegionNetworkEndpointGroupState{}
	if err := json.NewDecoder(av).Decode(crneg.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", crneg.Type(), crneg.LocalName(), err)
	}
	return nil
}

func (crneg *ComputeRegionNetworkEndpointGroup) State() (*computeRegionNetworkEndpointGroupState, bool) {
	return crneg.state, crneg.state != nil
}

func (crneg *ComputeRegionNetworkEndpointGroup) StateMust() *computeRegionNetworkEndpointGroupState {
	if crneg.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", crneg.Type(), crneg.LocalName()))
	}
	return crneg.state
}

func (crneg *ComputeRegionNetworkEndpointGroup) DependOn() terra.Reference {
	return terra.ReferenceResource(crneg)
}

type ComputeRegionNetworkEndpointGroupArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Network: string, optional
	Network terra.StringValue `hcl:"network,attr"`
	// NetworkEndpointType: string, optional
	NetworkEndpointType terra.StringValue `hcl:"network_endpoint_type,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// PscTargetService: string, optional
	PscTargetService terra.StringValue `hcl:"psc_target_service,attr"`
	// Region: string, required
	Region terra.StringValue `hcl:"region,attr" validate:"required"`
	// Subnetwork: string, optional
	Subnetwork terra.StringValue `hcl:"subnetwork,attr"`
	// AppEngine: optional
	AppEngine *computeregionnetworkendpointgroup.AppEngine `hcl:"app_engine,block"`
	// CloudFunction: optional
	CloudFunction *computeregionnetworkendpointgroup.CloudFunction `hcl:"cloud_function,block"`
	// CloudRun: optional
	CloudRun *computeregionnetworkendpointgroup.CloudRun `hcl:"cloud_run,block"`
	// Timeouts: optional
	Timeouts *computeregionnetworkendpointgroup.Timeouts `hcl:"timeouts,block"`
	// DependsOn contains resources that ComputeRegionNetworkEndpointGroup depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type computeRegionNetworkEndpointGroupAttributes struct {
	ref terra.Reference
}

func (crneg computeRegionNetworkEndpointGroupAttributes) Description() terra.StringValue {
	return terra.ReferenceString(crneg.ref.Append("description"))
}

func (crneg computeRegionNetworkEndpointGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceString(crneg.ref.Append("id"))
}

func (crneg computeRegionNetworkEndpointGroupAttributes) Name() terra.StringValue {
	return terra.ReferenceString(crneg.ref.Append("name"))
}

func (crneg computeRegionNetworkEndpointGroupAttributes) Network() terra.StringValue {
	return terra.ReferenceString(crneg.ref.Append("network"))
}

func (crneg computeRegionNetworkEndpointGroupAttributes) NetworkEndpointType() terra.StringValue {
	return terra.ReferenceString(crneg.ref.Append("network_endpoint_type"))
}

func (crneg computeRegionNetworkEndpointGroupAttributes) Project() terra.StringValue {
	return terra.ReferenceString(crneg.ref.Append("project"))
}

func (crneg computeRegionNetworkEndpointGroupAttributes) PscTargetService() terra.StringValue {
	return terra.ReferenceString(crneg.ref.Append("psc_target_service"))
}

func (crneg computeRegionNetworkEndpointGroupAttributes) Region() terra.StringValue {
	return terra.ReferenceString(crneg.ref.Append("region"))
}

func (crneg computeRegionNetworkEndpointGroupAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceString(crneg.ref.Append("self_link"))
}

func (crneg computeRegionNetworkEndpointGroupAttributes) Subnetwork() terra.StringValue {
	return terra.ReferenceString(crneg.ref.Append("subnetwork"))
}

func (crneg computeRegionNetworkEndpointGroupAttributes) AppEngine() terra.ListValue[computeregionnetworkendpointgroup.AppEngineAttributes] {
	return terra.ReferenceList[computeregionnetworkendpointgroup.AppEngineAttributes](crneg.ref.Append("app_engine"))
}

func (crneg computeRegionNetworkEndpointGroupAttributes) CloudFunction() terra.ListValue[computeregionnetworkendpointgroup.CloudFunctionAttributes] {
	return terra.ReferenceList[computeregionnetworkendpointgroup.CloudFunctionAttributes](crneg.ref.Append("cloud_function"))
}

func (crneg computeRegionNetworkEndpointGroupAttributes) CloudRun() terra.ListValue[computeregionnetworkendpointgroup.CloudRunAttributes] {
	return terra.ReferenceList[computeregionnetworkendpointgroup.CloudRunAttributes](crneg.ref.Append("cloud_run"))
}

func (crneg computeRegionNetworkEndpointGroupAttributes) Timeouts() computeregionnetworkendpointgroup.TimeoutsAttributes {
	return terra.ReferenceSingle[computeregionnetworkendpointgroup.TimeoutsAttributes](crneg.ref.Append("timeouts"))
}

type computeRegionNetworkEndpointGroupState struct {
	Description         string                                                 `json:"description"`
	Id                  string                                                 `json:"id"`
	Name                string                                                 `json:"name"`
	Network             string                                                 `json:"network"`
	NetworkEndpointType string                                                 `json:"network_endpoint_type"`
	Project             string                                                 `json:"project"`
	PscTargetService    string                                                 `json:"psc_target_service"`
	Region              string                                                 `json:"region"`
	SelfLink            string                                                 `json:"self_link"`
	Subnetwork          string                                                 `json:"subnetwork"`
	AppEngine           []computeregionnetworkendpointgroup.AppEngineState     `json:"app_engine"`
	CloudFunction       []computeregionnetworkendpointgroup.CloudFunctionState `json:"cloud_function"`
	CloudRun            []computeregionnetworkendpointgroup.CloudRunState      `json:"cloud_run"`
	Timeouts            *computeregionnetworkendpointgroup.TimeoutsState       `json:"timeouts"`
}
