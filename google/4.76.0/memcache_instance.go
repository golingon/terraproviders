// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	memcacheinstance "github.com/golingon/terraproviders/google/4.76.0/memcacheinstance"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewMemcacheInstance creates a new instance of [MemcacheInstance].
func NewMemcacheInstance(name string, args MemcacheInstanceArgs) *MemcacheInstance {
	return &MemcacheInstance{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*MemcacheInstance)(nil)

// MemcacheInstance represents the Terraform resource google_memcache_instance.
type MemcacheInstance struct {
	Name      string
	Args      MemcacheInstanceArgs
	state     *memcacheInstanceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [MemcacheInstance].
func (mi *MemcacheInstance) Type() string {
	return "google_memcache_instance"
}

// LocalName returns the local name for [MemcacheInstance].
func (mi *MemcacheInstance) LocalName() string {
	return mi.Name
}

// Configuration returns the configuration (args) for [MemcacheInstance].
func (mi *MemcacheInstance) Configuration() interface{} {
	return mi.Args
}

// DependOn is used for other resources to depend on [MemcacheInstance].
func (mi *MemcacheInstance) DependOn() terra.Reference {
	return terra.ReferenceResource(mi)
}

// Dependencies returns the list of resources [MemcacheInstance] depends_on.
func (mi *MemcacheInstance) Dependencies() terra.Dependencies {
	return mi.DependsOn
}

// LifecycleManagement returns the lifecycle block for [MemcacheInstance].
func (mi *MemcacheInstance) LifecycleManagement() *terra.Lifecycle {
	return mi.Lifecycle
}

// Attributes returns the attributes for [MemcacheInstance].
func (mi *MemcacheInstance) Attributes() memcacheInstanceAttributes {
	return memcacheInstanceAttributes{ref: terra.ReferenceResource(mi)}
}

// ImportState imports the given attribute values into [MemcacheInstance]'s state.
func (mi *MemcacheInstance) ImportState(av io.Reader) error {
	mi.state = &memcacheInstanceState{}
	if err := json.NewDecoder(av).Decode(mi.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", mi.Type(), mi.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [MemcacheInstance] has state.
func (mi *MemcacheInstance) State() (*memcacheInstanceState, bool) {
	return mi.state, mi.state != nil
}

// StateMust returns the state for [MemcacheInstance]. Panics if the state is nil.
func (mi *MemcacheInstance) StateMust() *memcacheInstanceState {
	if mi.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", mi.Type(), mi.LocalName()))
	}
	return mi.state
}

// MemcacheInstanceArgs contains the configurations for google_memcache_instance.
type MemcacheInstanceArgs struct {
	// AuthorizedNetwork: string, optional
	AuthorizedNetwork terra.StringValue `hcl:"authorized_network,attr"`
	// DisplayName: string, optional
	DisplayName terra.StringValue `hcl:"display_name,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// MemcacheVersion: string, optional
	MemcacheVersion terra.StringValue `hcl:"memcache_version,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// NodeCount: number, required
	NodeCount terra.NumberValue `hcl:"node_count,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
	// Zones: set of string, optional
	Zones terra.SetValue[terra.StringValue] `hcl:"zones,attr"`
	// MaintenanceSchedule: min=0
	MaintenanceSchedule []memcacheinstance.MaintenanceSchedule `hcl:"maintenance_schedule,block" validate:"min=0"`
	// MemcacheNodes: min=0
	MemcacheNodes []memcacheinstance.MemcacheNodes `hcl:"memcache_nodes,block" validate:"min=0"`
	// MaintenancePolicy: optional
	MaintenancePolicy *memcacheinstance.MaintenancePolicy `hcl:"maintenance_policy,block"`
	// MemcacheParameters: optional
	MemcacheParameters *memcacheinstance.MemcacheParameters `hcl:"memcache_parameters,block"`
	// NodeConfig: required
	NodeConfig *memcacheinstance.NodeConfig `hcl:"node_config,block" validate:"required"`
	// Timeouts: optional
	Timeouts *memcacheinstance.Timeouts `hcl:"timeouts,block"`
}
type memcacheInstanceAttributes struct {
	ref terra.Reference
}

// AuthorizedNetwork returns a reference to field authorized_network of google_memcache_instance.
func (mi memcacheInstanceAttributes) AuthorizedNetwork() terra.StringValue {
	return terra.ReferenceAsString(mi.ref.Append("authorized_network"))
}

// CreateTime returns a reference to field create_time of google_memcache_instance.
func (mi memcacheInstanceAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(mi.ref.Append("create_time"))
}

// DiscoveryEndpoint returns a reference to field discovery_endpoint of google_memcache_instance.
func (mi memcacheInstanceAttributes) DiscoveryEndpoint() terra.StringValue {
	return terra.ReferenceAsString(mi.ref.Append("discovery_endpoint"))
}

// DisplayName returns a reference to field display_name of google_memcache_instance.
func (mi memcacheInstanceAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(mi.ref.Append("display_name"))
}

// Id returns a reference to field id of google_memcache_instance.
func (mi memcacheInstanceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(mi.ref.Append("id"))
}

// Labels returns a reference to field labels of google_memcache_instance.
func (mi memcacheInstanceAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](mi.ref.Append("labels"))
}

// MemcacheFullVersion returns a reference to field memcache_full_version of google_memcache_instance.
func (mi memcacheInstanceAttributes) MemcacheFullVersion() terra.StringValue {
	return terra.ReferenceAsString(mi.ref.Append("memcache_full_version"))
}

// MemcacheVersion returns a reference to field memcache_version of google_memcache_instance.
func (mi memcacheInstanceAttributes) MemcacheVersion() terra.StringValue {
	return terra.ReferenceAsString(mi.ref.Append("memcache_version"))
}

// Name returns a reference to field name of google_memcache_instance.
func (mi memcacheInstanceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(mi.ref.Append("name"))
}

// NodeCount returns a reference to field node_count of google_memcache_instance.
func (mi memcacheInstanceAttributes) NodeCount() terra.NumberValue {
	return terra.ReferenceAsNumber(mi.ref.Append("node_count"))
}

// Project returns a reference to field project of google_memcache_instance.
func (mi memcacheInstanceAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(mi.ref.Append("project"))
}

// Region returns a reference to field region of google_memcache_instance.
func (mi memcacheInstanceAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(mi.ref.Append("region"))
}

// Zones returns a reference to field zones of google_memcache_instance.
func (mi memcacheInstanceAttributes) Zones() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](mi.ref.Append("zones"))
}

func (mi memcacheInstanceAttributes) MaintenanceSchedule() terra.ListValue[memcacheinstance.MaintenanceScheduleAttributes] {
	return terra.ReferenceAsList[memcacheinstance.MaintenanceScheduleAttributes](mi.ref.Append("maintenance_schedule"))
}

func (mi memcacheInstanceAttributes) MemcacheNodes() terra.ListValue[memcacheinstance.MemcacheNodesAttributes] {
	return terra.ReferenceAsList[memcacheinstance.MemcacheNodesAttributes](mi.ref.Append("memcache_nodes"))
}

func (mi memcacheInstanceAttributes) MaintenancePolicy() terra.ListValue[memcacheinstance.MaintenancePolicyAttributes] {
	return terra.ReferenceAsList[memcacheinstance.MaintenancePolicyAttributes](mi.ref.Append("maintenance_policy"))
}

func (mi memcacheInstanceAttributes) MemcacheParameters() terra.ListValue[memcacheinstance.MemcacheParametersAttributes] {
	return terra.ReferenceAsList[memcacheinstance.MemcacheParametersAttributes](mi.ref.Append("memcache_parameters"))
}

func (mi memcacheInstanceAttributes) NodeConfig() terra.ListValue[memcacheinstance.NodeConfigAttributes] {
	return terra.ReferenceAsList[memcacheinstance.NodeConfigAttributes](mi.ref.Append("node_config"))
}

func (mi memcacheInstanceAttributes) Timeouts() memcacheinstance.TimeoutsAttributes {
	return terra.ReferenceAsSingle[memcacheinstance.TimeoutsAttributes](mi.ref.Append("timeouts"))
}

type memcacheInstanceState struct {
	AuthorizedNetwork   string                                      `json:"authorized_network"`
	CreateTime          string                                      `json:"create_time"`
	DiscoveryEndpoint   string                                      `json:"discovery_endpoint"`
	DisplayName         string                                      `json:"display_name"`
	Id                  string                                      `json:"id"`
	Labels              map[string]string                           `json:"labels"`
	MemcacheFullVersion string                                      `json:"memcache_full_version"`
	MemcacheVersion     string                                      `json:"memcache_version"`
	Name                string                                      `json:"name"`
	NodeCount           float64                                     `json:"node_count"`
	Project             string                                      `json:"project"`
	Region              string                                      `json:"region"`
	Zones               []string                                    `json:"zones"`
	MaintenanceSchedule []memcacheinstance.MaintenanceScheduleState `json:"maintenance_schedule"`
	MemcacheNodes       []memcacheinstance.MemcacheNodesState       `json:"memcache_nodes"`
	MaintenancePolicy   []memcacheinstance.MaintenancePolicyState   `json:"maintenance_policy"`
	MemcacheParameters  []memcacheinstance.MemcacheParametersState  `json:"memcache_parameters"`
	NodeConfig          []memcacheinstance.NodeConfigState          `json:"node_config"`
	Timeouts            *memcacheinstance.TimeoutsState             `json:"timeouts"`
}
