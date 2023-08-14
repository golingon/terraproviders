// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	elb "github.com/golingon/terraproviders/aws/5.12.0/elb"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewElb creates a new instance of [Elb].
func NewElb(name string, args ElbArgs) *Elb {
	return &Elb{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Elb)(nil)

// Elb represents the Terraform resource aws_elb.
type Elb struct {
	Name      string
	Args      ElbArgs
	state     *elbState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Elb].
func (e *Elb) Type() string {
	return "aws_elb"
}

// LocalName returns the local name for [Elb].
func (e *Elb) LocalName() string {
	return e.Name
}

// Configuration returns the configuration (args) for [Elb].
func (e *Elb) Configuration() interface{} {
	return e.Args
}

// DependOn is used for other resources to depend on [Elb].
func (e *Elb) DependOn() terra.Reference {
	return terra.ReferenceResource(e)
}

// Dependencies returns the list of resources [Elb] depends_on.
func (e *Elb) Dependencies() terra.Dependencies {
	return e.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Elb].
func (e *Elb) LifecycleManagement() *terra.Lifecycle {
	return e.Lifecycle
}

// Attributes returns the attributes for [Elb].
func (e *Elb) Attributes() elbAttributes {
	return elbAttributes{ref: terra.ReferenceResource(e)}
}

// ImportState imports the given attribute values into [Elb]'s state.
func (e *Elb) ImportState(av io.Reader) error {
	e.state = &elbState{}
	if err := json.NewDecoder(av).Decode(e.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", e.Type(), e.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Elb] has state.
func (e *Elb) State() (*elbState, bool) {
	return e.state, e.state != nil
}

// StateMust returns the state for [Elb]. Panics if the state is nil.
func (e *Elb) StateMust() *elbState {
	if e.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", e.Type(), e.LocalName()))
	}
	return e.state
}

// ElbArgs contains the configurations for aws_elb.
type ElbArgs struct {
	// AvailabilityZones: set of string, optional
	AvailabilityZones terra.SetValue[terra.StringValue] `hcl:"availability_zones,attr"`
	// ConnectionDraining: bool, optional
	ConnectionDraining terra.BoolValue `hcl:"connection_draining,attr"`
	// ConnectionDrainingTimeout: number, optional
	ConnectionDrainingTimeout terra.NumberValue `hcl:"connection_draining_timeout,attr"`
	// CrossZoneLoadBalancing: bool, optional
	CrossZoneLoadBalancing terra.BoolValue `hcl:"cross_zone_load_balancing,attr"`
	// DesyncMitigationMode: string, optional
	DesyncMitigationMode terra.StringValue `hcl:"desync_mitigation_mode,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IdleTimeout: number, optional
	IdleTimeout terra.NumberValue `hcl:"idle_timeout,attr"`
	// Instances: set of string, optional
	Instances terra.SetValue[terra.StringValue] `hcl:"instances,attr"`
	// Internal: bool, optional
	Internal terra.BoolValue `hcl:"internal,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// NamePrefix: string, optional
	NamePrefix terra.StringValue `hcl:"name_prefix,attr"`
	// SecurityGroups: set of string, optional
	SecurityGroups terra.SetValue[terra.StringValue] `hcl:"security_groups,attr"`
	// SourceSecurityGroup: string, optional
	SourceSecurityGroup terra.StringValue `hcl:"source_security_group,attr"`
	// Subnets: set of string, optional
	Subnets terra.SetValue[terra.StringValue] `hcl:"subnets,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// AccessLogs: optional
	AccessLogs *elb.AccessLogs `hcl:"access_logs,block"`
	// HealthCheck: optional
	HealthCheck *elb.HealthCheck `hcl:"health_check,block"`
	// Listener: min=1
	Listener []elb.Listener `hcl:"listener,block" validate:"min=1"`
	// Timeouts: optional
	Timeouts *elb.Timeouts `hcl:"timeouts,block"`
}
type elbAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_elb.
func (e elbAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("arn"))
}

// AvailabilityZones returns a reference to field availability_zones of aws_elb.
func (e elbAttributes) AvailabilityZones() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](e.ref.Append("availability_zones"))
}

// ConnectionDraining returns a reference to field connection_draining of aws_elb.
func (e elbAttributes) ConnectionDraining() terra.BoolValue {
	return terra.ReferenceAsBool(e.ref.Append("connection_draining"))
}

// ConnectionDrainingTimeout returns a reference to field connection_draining_timeout of aws_elb.
func (e elbAttributes) ConnectionDrainingTimeout() terra.NumberValue {
	return terra.ReferenceAsNumber(e.ref.Append("connection_draining_timeout"))
}

// CrossZoneLoadBalancing returns a reference to field cross_zone_load_balancing of aws_elb.
func (e elbAttributes) CrossZoneLoadBalancing() terra.BoolValue {
	return terra.ReferenceAsBool(e.ref.Append("cross_zone_load_balancing"))
}

// DesyncMitigationMode returns a reference to field desync_mitigation_mode of aws_elb.
func (e elbAttributes) DesyncMitigationMode() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("desync_mitigation_mode"))
}

// DnsName returns a reference to field dns_name of aws_elb.
func (e elbAttributes) DnsName() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("dns_name"))
}

// Id returns a reference to field id of aws_elb.
func (e elbAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("id"))
}

// IdleTimeout returns a reference to field idle_timeout of aws_elb.
func (e elbAttributes) IdleTimeout() terra.NumberValue {
	return terra.ReferenceAsNumber(e.ref.Append("idle_timeout"))
}

// Instances returns a reference to field instances of aws_elb.
func (e elbAttributes) Instances() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](e.ref.Append("instances"))
}

// Internal returns a reference to field internal of aws_elb.
func (e elbAttributes) Internal() terra.BoolValue {
	return terra.ReferenceAsBool(e.ref.Append("internal"))
}

// Name returns a reference to field name of aws_elb.
func (e elbAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("name"))
}

// NamePrefix returns a reference to field name_prefix of aws_elb.
func (e elbAttributes) NamePrefix() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("name_prefix"))
}

// SecurityGroups returns a reference to field security_groups of aws_elb.
func (e elbAttributes) SecurityGroups() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](e.ref.Append("security_groups"))
}

// SourceSecurityGroup returns a reference to field source_security_group of aws_elb.
func (e elbAttributes) SourceSecurityGroup() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("source_security_group"))
}

// SourceSecurityGroupId returns a reference to field source_security_group_id of aws_elb.
func (e elbAttributes) SourceSecurityGroupId() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("source_security_group_id"))
}

// Subnets returns a reference to field subnets of aws_elb.
func (e elbAttributes) Subnets() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](e.ref.Append("subnets"))
}

// Tags returns a reference to field tags of aws_elb.
func (e elbAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](e.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_elb.
func (e elbAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](e.ref.Append("tags_all"))
}

// ZoneId returns a reference to field zone_id of aws_elb.
func (e elbAttributes) ZoneId() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("zone_id"))
}

func (e elbAttributes) AccessLogs() terra.ListValue[elb.AccessLogsAttributes] {
	return terra.ReferenceAsList[elb.AccessLogsAttributes](e.ref.Append("access_logs"))
}

func (e elbAttributes) HealthCheck() terra.ListValue[elb.HealthCheckAttributes] {
	return terra.ReferenceAsList[elb.HealthCheckAttributes](e.ref.Append("health_check"))
}

func (e elbAttributes) Listener() terra.SetValue[elb.ListenerAttributes] {
	return terra.ReferenceAsSet[elb.ListenerAttributes](e.ref.Append("listener"))
}

func (e elbAttributes) Timeouts() elb.TimeoutsAttributes {
	return terra.ReferenceAsSingle[elb.TimeoutsAttributes](e.ref.Append("timeouts"))
}

type elbState struct {
	Arn                       string                 `json:"arn"`
	AvailabilityZones         []string               `json:"availability_zones"`
	ConnectionDraining        bool                   `json:"connection_draining"`
	ConnectionDrainingTimeout float64                `json:"connection_draining_timeout"`
	CrossZoneLoadBalancing    bool                   `json:"cross_zone_load_balancing"`
	DesyncMitigationMode      string                 `json:"desync_mitigation_mode"`
	DnsName                   string                 `json:"dns_name"`
	Id                        string                 `json:"id"`
	IdleTimeout               float64                `json:"idle_timeout"`
	Instances                 []string               `json:"instances"`
	Internal                  bool                   `json:"internal"`
	Name                      string                 `json:"name"`
	NamePrefix                string                 `json:"name_prefix"`
	SecurityGroups            []string               `json:"security_groups"`
	SourceSecurityGroup       string                 `json:"source_security_group"`
	SourceSecurityGroupId     string                 `json:"source_security_group_id"`
	Subnets                   []string               `json:"subnets"`
	Tags                      map[string]string      `json:"tags"`
	TagsAll                   map[string]string      `json:"tags_all"`
	ZoneId                    string                 `json:"zone_id"`
	AccessLogs                []elb.AccessLogsState  `json:"access_logs"`
	HealthCheck               []elb.HealthCheckState `json:"health_check"`
	Listener                  []elb.ListenerState    `json:"listener"`
	Timeouts                  *elb.TimeoutsState     `json:"timeouts"`
}
