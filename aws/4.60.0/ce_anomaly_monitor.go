// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewCeAnomalyMonitor(name string, args CeAnomalyMonitorArgs) *CeAnomalyMonitor {
	return &CeAnomalyMonitor{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CeAnomalyMonitor)(nil)

type CeAnomalyMonitor struct {
	Name  string
	Args  CeAnomalyMonitorArgs
	state *ceAnomalyMonitorState
}

func (cam *CeAnomalyMonitor) Type() string {
	return "aws_ce_anomaly_monitor"
}

func (cam *CeAnomalyMonitor) LocalName() string {
	return cam.Name
}

func (cam *CeAnomalyMonitor) Configuration() interface{} {
	return cam.Args
}

func (cam *CeAnomalyMonitor) Attributes() ceAnomalyMonitorAttributes {
	return ceAnomalyMonitorAttributes{ref: terra.ReferenceResource(cam)}
}

func (cam *CeAnomalyMonitor) ImportState(av io.Reader) error {
	cam.state = &ceAnomalyMonitorState{}
	if err := json.NewDecoder(av).Decode(cam.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cam.Type(), cam.LocalName(), err)
	}
	return nil
}

func (cam *CeAnomalyMonitor) State() (*ceAnomalyMonitorState, bool) {
	return cam.state, cam.state != nil
}

func (cam *CeAnomalyMonitor) StateMust() *ceAnomalyMonitorState {
	if cam.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cam.Type(), cam.LocalName()))
	}
	return cam.state
}

func (cam *CeAnomalyMonitor) DependOn() terra.Reference {
	return terra.ReferenceResource(cam)
}

type CeAnomalyMonitorArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// MonitorDimension: string, optional
	MonitorDimension terra.StringValue `hcl:"monitor_dimension,attr"`
	// MonitorSpecification: string, optional
	MonitorSpecification terra.StringValue `hcl:"monitor_specification,attr"`
	// MonitorType: string, required
	MonitorType terra.StringValue `hcl:"monitor_type,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// DependsOn contains resources that CeAnomalyMonitor depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type ceAnomalyMonitorAttributes struct {
	ref terra.Reference
}

func (cam ceAnomalyMonitorAttributes) Arn() terra.StringValue {
	return terra.ReferenceString(cam.ref.Append("arn"))
}

func (cam ceAnomalyMonitorAttributes) Id() terra.StringValue {
	return terra.ReferenceString(cam.ref.Append("id"))
}

func (cam ceAnomalyMonitorAttributes) MonitorDimension() terra.StringValue {
	return terra.ReferenceString(cam.ref.Append("monitor_dimension"))
}

func (cam ceAnomalyMonitorAttributes) MonitorSpecification() terra.StringValue {
	return terra.ReferenceString(cam.ref.Append("monitor_specification"))
}

func (cam ceAnomalyMonitorAttributes) MonitorType() terra.StringValue {
	return terra.ReferenceString(cam.ref.Append("monitor_type"))
}

func (cam ceAnomalyMonitorAttributes) Name() terra.StringValue {
	return terra.ReferenceString(cam.ref.Append("name"))
}

func (cam ceAnomalyMonitorAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](cam.ref.Append("tags"))
}

func (cam ceAnomalyMonitorAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](cam.ref.Append("tags_all"))
}

type ceAnomalyMonitorState struct {
	Arn                  string            `json:"arn"`
	Id                   string            `json:"id"`
	MonitorDimension     string            `json:"monitor_dimension"`
	MonitorSpecification string            `json:"monitor_specification"`
	MonitorType          string            `json:"monitor_type"`
	Name                 string            `json:"name"`
	Tags                 map[string]string `json:"tags"`
	TagsAll              map[string]string `json:"tags_all"`
}
