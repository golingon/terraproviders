// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	alloydbinstance "github.com/golingon/terraproviders/google/4.59.0/alloydbinstance"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewAlloydbInstance(name string, args AlloydbInstanceArgs) *AlloydbInstance {
	return &AlloydbInstance{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AlloydbInstance)(nil)

type AlloydbInstance struct {
	Name  string
	Args  AlloydbInstanceArgs
	state *alloydbInstanceState
}

func (ai *AlloydbInstance) Type() string {
	return "google_alloydb_instance"
}

func (ai *AlloydbInstance) LocalName() string {
	return ai.Name
}

func (ai *AlloydbInstance) Configuration() interface{} {
	return ai.Args
}

func (ai *AlloydbInstance) Attributes() alloydbInstanceAttributes {
	return alloydbInstanceAttributes{ref: terra.ReferenceResource(ai)}
}

func (ai *AlloydbInstance) ImportState(av io.Reader) error {
	ai.state = &alloydbInstanceState{}
	if err := json.NewDecoder(av).Decode(ai.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ai.Type(), ai.LocalName(), err)
	}
	return nil
}

func (ai *AlloydbInstance) State() (*alloydbInstanceState, bool) {
	return ai.state, ai.state != nil
}

func (ai *AlloydbInstance) StateMust() *alloydbInstanceState {
	if ai.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ai.Type(), ai.LocalName()))
	}
	return ai.state
}

func (ai *AlloydbInstance) DependOn() terra.Reference {
	return terra.ReferenceResource(ai)
}

type AlloydbInstanceArgs struct {
	// Annotations: map of string, optional
	Annotations terra.MapValue[terra.StringValue] `hcl:"annotations,attr"`
	// AvailabilityType: string, optional
	AvailabilityType terra.StringValue `hcl:"availability_type,attr"`
	// Cluster: string, required
	Cluster terra.StringValue `hcl:"cluster,attr" validate:"required"`
	// DatabaseFlags: map of string, optional
	DatabaseFlags terra.MapValue[terra.StringValue] `hcl:"database_flags,attr"`
	// DisplayName: string, optional
	DisplayName terra.StringValue `hcl:"display_name,attr"`
	// GceZone: string, optional
	GceZone terra.StringValue `hcl:"gce_zone,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InstanceId: string, required
	InstanceId terra.StringValue `hcl:"instance_id,attr" validate:"required"`
	// InstanceType: string, required
	InstanceType terra.StringValue `hcl:"instance_type,attr" validate:"required"`
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// MachineConfig: optional
	MachineConfig *alloydbinstance.MachineConfig `hcl:"machine_config,block"`
	// ReadPoolConfig: optional
	ReadPoolConfig *alloydbinstance.ReadPoolConfig `hcl:"read_pool_config,block"`
	// Timeouts: optional
	Timeouts *alloydbinstance.Timeouts `hcl:"timeouts,block"`
	// DependsOn contains resources that AlloydbInstance depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type alloydbInstanceAttributes struct {
	ref terra.Reference
}

func (ai alloydbInstanceAttributes) Annotations() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](ai.ref.Append("annotations"))
}

func (ai alloydbInstanceAttributes) AvailabilityType() terra.StringValue {
	return terra.ReferenceString(ai.ref.Append("availability_type"))
}

func (ai alloydbInstanceAttributes) Cluster() terra.StringValue {
	return terra.ReferenceString(ai.ref.Append("cluster"))
}

func (ai alloydbInstanceAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceString(ai.ref.Append("create_time"))
}

func (ai alloydbInstanceAttributes) DatabaseFlags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](ai.ref.Append("database_flags"))
}

func (ai alloydbInstanceAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceString(ai.ref.Append("display_name"))
}

func (ai alloydbInstanceAttributes) GceZone() terra.StringValue {
	return terra.ReferenceString(ai.ref.Append("gce_zone"))
}

func (ai alloydbInstanceAttributes) Id() terra.StringValue {
	return terra.ReferenceString(ai.ref.Append("id"))
}

func (ai alloydbInstanceAttributes) InstanceId() terra.StringValue {
	return terra.ReferenceString(ai.ref.Append("instance_id"))
}

func (ai alloydbInstanceAttributes) InstanceType() terra.StringValue {
	return terra.ReferenceString(ai.ref.Append("instance_type"))
}

func (ai alloydbInstanceAttributes) IpAddress() terra.StringValue {
	return terra.ReferenceString(ai.ref.Append("ip_address"))
}

func (ai alloydbInstanceAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](ai.ref.Append("labels"))
}

func (ai alloydbInstanceAttributes) Name() terra.StringValue {
	return terra.ReferenceString(ai.ref.Append("name"))
}

func (ai alloydbInstanceAttributes) Reconciling() terra.BoolValue {
	return terra.ReferenceBool(ai.ref.Append("reconciling"))
}

func (ai alloydbInstanceAttributes) State() terra.StringValue {
	return terra.ReferenceString(ai.ref.Append("state"))
}

func (ai alloydbInstanceAttributes) Uid() terra.StringValue {
	return terra.ReferenceString(ai.ref.Append("uid"))
}

func (ai alloydbInstanceAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceString(ai.ref.Append("update_time"))
}

func (ai alloydbInstanceAttributes) MachineConfig() terra.ListValue[alloydbinstance.MachineConfigAttributes] {
	return terra.ReferenceList[alloydbinstance.MachineConfigAttributes](ai.ref.Append("machine_config"))
}

func (ai alloydbInstanceAttributes) ReadPoolConfig() terra.ListValue[alloydbinstance.ReadPoolConfigAttributes] {
	return terra.ReferenceList[alloydbinstance.ReadPoolConfigAttributes](ai.ref.Append("read_pool_config"))
}

func (ai alloydbInstanceAttributes) Timeouts() alloydbinstance.TimeoutsAttributes {
	return terra.ReferenceSingle[alloydbinstance.TimeoutsAttributes](ai.ref.Append("timeouts"))
}

type alloydbInstanceState struct {
	Annotations      map[string]string                     `json:"annotations"`
	AvailabilityType string                                `json:"availability_type"`
	Cluster          string                                `json:"cluster"`
	CreateTime       string                                `json:"create_time"`
	DatabaseFlags    map[string]string                     `json:"database_flags"`
	DisplayName      string                                `json:"display_name"`
	GceZone          string                                `json:"gce_zone"`
	Id               string                                `json:"id"`
	InstanceId       string                                `json:"instance_id"`
	InstanceType     string                                `json:"instance_type"`
	IpAddress        string                                `json:"ip_address"`
	Labels           map[string]string                     `json:"labels"`
	Name             string                                `json:"name"`
	Reconciling      bool                                  `json:"reconciling"`
	State            string                                `json:"state"`
	Uid              string                                `json:"uid"`
	UpdateTime       string                                `json:"update_time"`
	MachineConfig    []alloydbinstance.MachineConfigState  `json:"machine_config"`
	ReadPoolConfig   []alloydbinstance.ReadPoolConfigState `json:"read_pool_config"`
	Timeouts         *alloydbinstance.TimeoutsState        `json:"timeouts"`
}
