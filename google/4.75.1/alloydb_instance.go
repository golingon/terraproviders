// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	alloydbinstance "github.com/golingon/terraproviders/google/4.75.1/alloydbinstance"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewAlloydbInstance creates a new instance of [AlloydbInstance].
func NewAlloydbInstance(name string, args AlloydbInstanceArgs) *AlloydbInstance {
	return &AlloydbInstance{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AlloydbInstance)(nil)

// AlloydbInstance represents the Terraform resource google_alloydb_instance.
type AlloydbInstance struct {
	Name      string
	Args      AlloydbInstanceArgs
	state     *alloydbInstanceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [AlloydbInstance].
func (ai *AlloydbInstance) Type() string {
	return "google_alloydb_instance"
}

// LocalName returns the local name for [AlloydbInstance].
func (ai *AlloydbInstance) LocalName() string {
	return ai.Name
}

// Configuration returns the configuration (args) for [AlloydbInstance].
func (ai *AlloydbInstance) Configuration() interface{} {
	return ai.Args
}

// DependOn is used for other resources to depend on [AlloydbInstance].
func (ai *AlloydbInstance) DependOn() terra.Reference {
	return terra.ReferenceResource(ai)
}

// Dependencies returns the list of resources [AlloydbInstance] depends_on.
func (ai *AlloydbInstance) Dependencies() terra.Dependencies {
	return ai.DependsOn
}

// LifecycleManagement returns the lifecycle block for [AlloydbInstance].
func (ai *AlloydbInstance) LifecycleManagement() *terra.Lifecycle {
	return ai.Lifecycle
}

// Attributes returns the attributes for [AlloydbInstance].
func (ai *AlloydbInstance) Attributes() alloydbInstanceAttributes {
	return alloydbInstanceAttributes{ref: terra.ReferenceResource(ai)}
}

// ImportState imports the given attribute values into [AlloydbInstance]'s state.
func (ai *AlloydbInstance) ImportState(av io.Reader) error {
	ai.state = &alloydbInstanceState{}
	if err := json.NewDecoder(av).Decode(ai.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ai.Type(), ai.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [AlloydbInstance] has state.
func (ai *AlloydbInstance) State() (*alloydbInstanceState, bool) {
	return ai.state, ai.state != nil
}

// StateMust returns the state for [AlloydbInstance]. Panics if the state is nil.
func (ai *AlloydbInstance) StateMust() *alloydbInstanceState {
	if ai.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ai.Type(), ai.LocalName()))
	}
	return ai.state
}

// AlloydbInstanceArgs contains the configurations for google_alloydb_instance.
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
}
type alloydbInstanceAttributes struct {
	ref terra.Reference
}

// Annotations returns a reference to field annotations of google_alloydb_instance.
func (ai alloydbInstanceAttributes) Annotations() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ai.ref.Append("annotations"))
}

// AvailabilityType returns a reference to field availability_type of google_alloydb_instance.
func (ai alloydbInstanceAttributes) AvailabilityType() terra.StringValue {
	return terra.ReferenceAsString(ai.ref.Append("availability_type"))
}

// Cluster returns a reference to field cluster of google_alloydb_instance.
func (ai alloydbInstanceAttributes) Cluster() terra.StringValue {
	return terra.ReferenceAsString(ai.ref.Append("cluster"))
}

// CreateTime returns a reference to field create_time of google_alloydb_instance.
func (ai alloydbInstanceAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(ai.ref.Append("create_time"))
}

// DatabaseFlags returns a reference to field database_flags of google_alloydb_instance.
func (ai alloydbInstanceAttributes) DatabaseFlags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ai.ref.Append("database_flags"))
}

// DisplayName returns a reference to field display_name of google_alloydb_instance.
func (ai alloydbInstanceAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(ai.ref.Append("display_name"))
}

// GceZone returns a reference to field gce_zone of google_alloydb_instance.
func (ai alloydbInstanceAttributes) GceZone() terra.StringValue {
	return terra.ReferenceAsString(ai.ref.Append("gce_zone"))
}

// Id returns a reference to field id of google_alloydb_instance.
func (ai alloydbInstanceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ai.ref.Append("id"))
}

// InstanceId returns a reference to field instance_id of google_alloydb_instance.
func (ai alloydbInstanceAttributes) InstanceId() terra.StringValue {
	return terra.ReferenceAsString(ai.ref.Append("instance_id"))
}

// InstanceType returns a reference to field instance_type of google_alloydb_instance.
func (ai alloydbInstanceAttributes) InstanceType() terra.StringValue {
	return terra.ReferenceAsString(ai.ref.Append("instance_type"))
}

// IpAddress returns a reference to field ip_address of google_alloydb_instance.
func (ai alloydbInstanceAttributes) IpAddress() terra.StringValue {
	return terra.ReferenceAsString(ai.ref.Append("ip_address"))
}

// Labels returns a reference to field labels of google_alloydb_instance.
func (ai alloydbInstanceAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ai.ref.Append("labels"))
}

// Name returns a reference to field name of google_alloydb_instance.
func (ai alloydbInstanceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ai.ref.Append("name"))
}

// Reconciling returns a reference to field reconciling of google_alloydb_instance.
func (ai alloydbInstanceAttributes) Reconciling() terra.BoolValue {
	return terra.ReferenceAsBool(ai.ref.Append("reconciling"))
}

// State returns a reference to field state of google_alloydb_instance.
func (ai alloydbInstanceAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(ai.ref.Append("state"))
}

// Uid returns a reference to field uid of google_alloydb_instance.
func (ai alloydbInstanceAttributes) Uid() terra.StringValue {
	return terra.ReferenceAsString(ai.ref.Append("uid"))
}

// UpdateTime returns a reference to field update_time of google_alloydb_instance.
func (ai alloydbInstanceAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(ai.ref.Append("update_time"))
}

func (ai alloydbInstanceAttributes) MachineConfig() terra.ListValue[alloydbinstance.MachineConfigAttributes] {
	return terra.ReferenceAsList[alloydbinstance.MachineConfigAttributes](ai.ref.Append("machine_config"))
}

func (ai alloydbInstanceAttributes) ReadPoolConfig() terra.ListValue[alloydbinstance.ReadPoolConfigAttributes] {
	return terra.ReferenceAsList[alloydbinstance.ReadPoolConfigAttributes](ai.ref.Append("read_pool_config"))
}

func (ai alloydbInstanceAttributes) Timeouts() alloydbinstance.TimeoutsAttributes {
	return terra.ReferenceAsSingle[alloydbinstance.TimeoutsAttributes](ai.ref.Append("timeouts"))
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