// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	dnsmanagedzone "github.com/golingon/terraproviders/google/4.77.0/dnsmanagedzone"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDnsManagedZone creates a new instance of [DnsManagedZone].
func NewDnsManagedZone(name string, args DnsManagedZoneArgs) *DnsManagedZone {
	return &DnsManagedZone{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DnsManagedZone)(nil)

// DnsManagedZone represents the Terraform resource google_dns_managed_zone.
type DnsManagedZone struct {
	Name      string
	Args      DnsManagedZoneArgs
	state     *dnsManagedZoneState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DnsManagedZone].
func (dmz *DnsManagedZone) Type() string {
	return "google_dns_managed_zone"
}

// LocalName returns the local name for [DnsManagedZone].
func (dmz *DnsManagedZone) LocalName() string {
	return dmz.Name
}

// Configuration returns the configuration (args) for [DnsManagedZone].
func (dmz *DnsManagedZone) Configuration() interface{} {
	return dmz.Args
}

// DependOn is used for other resources to depend on [DnsManagedZone].
func (dmz *DnsManagedZone) DependOn() terra.Reference {
	return terra.ReferenceResource(dmz)
}

// Dependencies returns the list of resources [DnsManagedZone] depends_on.
func (dmz *DnsManagedZone) Dependencies() terra.Dependencies {
	return dmz.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DnsManagedZone].
func (dmz *DnsManagedZone) LifecycleManagement() *terra.Lifecycle {
	return dmz.Lifecycle
}

// Attributes returns the attributes for [DnsManagedZone].
func (dmz *DnsManagedZone) Attributes() dnsManagedZoneAttributes {
	return dnsManagedZoneAttributes{ref: terra.ReferenceResource(dmz)}
}

// ImportState imports the given attribute values into [DnsManagedZone]'s state.
func (dmz *DnsManagedZone) ImportState(av io.Reader) error {
	dmz.state = &dnsManagedZoneState{}
	if err := json.NewDecoder(av).Decode(dmz.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dmz.Type(), dmz.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DnsManagedZone] has state.
func (dmz *DnsManagedZone) State() (*dnsManagedZoneState, bool) {
	return dmz.state, dmz.state != nil
}

// StateMust returns the state for [DnsManagedZone]. Panics if the state is nil.
func (dmz *DnsManagedZone) StateMust() *dnsManagedZoneState {
	if dmz.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dmz.Type(), dmz.LocalName()))
	}
	return dmz.state
}

// DnsManagedZoneArgs contains the configurations for google_dns_managed_zone.
type DnsManagedZoneArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// DnsName: string, required
	DnsName terra.StringValue `hcl:"dns_name,attr" validate:"required"`
	// ForceDestroy: bool, optional
	ForceDestroy terra.BoolValue `hcl:"force_destroy,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Visibility: string, optional
	Visibility terra.StringValue `hcl:"visibility,attr"`
	// CloudLoggingConfig: optional
	CloudLoggingConfig *dnsmanagedzone.CloudLoggingConfig `hcl:"cloud_logging_config,block"`
	// DnssecConfig: optional
	DnssecConfig *dnsmanagedzone.DnssecConfig `hcl:"dnssec_config,block"`
	// ForwardingConfig: optional
	ForwardingConfig *dnsmanagedzone.ForwardingConfig `hcl:"forwarding_config,block"`
	// PeeringConfig: optional
	PeeringConfig *dnsmanagedzone.PeeringConfig `hcl:"peering_config,block"`
	// PrivateVisibilityConfig: optional
	PrivateVisibilityConfig *dnsmanagedzone.PrivateVisibilityConfig `hcl:"private_visibility_config,block"`
	// Timeouts: optional
	Timeouts *dnsmanagedzone.Timeouts `hcl:"timeouts,block"`
}
type dnsManagedZoneAttributes struct {
	ref terra.Reference
}

// CreationTime returns a reference to field creation_time of google_dns_managed_zone.
func (dmz dnsManagedZoneAttributes) CreationTime() terra.StringValue {
	return terra.ReferenceAsString(dmz.ref.Append("creation_time"))
}

// Description returns a reference to field description of google_dns_managed_zone.
func (dmz dnsManagedZoneAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(dmz.ref.Append("description"))
}

// DnsName returns a reference to field dns_name of google_dns_managed_zone.
func (dmz dnsManagedZoneAttributes) DnsName() terra.StringValue {
	return terra.ReferenceAsString(dmz.ref.Append("dns_name"))
}

// ForceDestroy returns a reference to field force_destroy of google_dns_managed_zone.
func (dmz dnsManagedZoneAttributes) ForceDestroy() terra.BoolValue {
	return terra.ReferenceAsBool(dmz.ref.Append("force_destroy"))
}

// Id returns a reference to field id of google_dns_managed_zone.
func (dmz dnsManagedZoneAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dmz.ref.Append("id"))
}

// Labels returns a reference to field labels of google_dns_managed_zone.
func (dmz dnsManagedZoneAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dmz.ref.Append("labels"))
}

// ManagedZoneId returns a reference to field managed_zone_id of google_dns_managed_zone.
func (dmz dnsManagedZoneAttributes) ManagedZoneId() terra.NumberValue {
	return terra.ReferenceAsNumber(dmz.ref.Append("managed_zone_id"))
}

// Name returns a reference to field name of google_dns_managed_zone.
func (dmz dnsManagedZoneAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dmz.ref.Append("name"))
}

// NameServers returns a reference to field name_servers of google_dns_managed_zone.
func (dmz dnsManagedZoneAttributes) NameServers() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](dmz.ref.Append("name_servers"))
}

// Project returns a reference to field project of google_dns_managed_zone.
func (dmz dnsManagedZoneAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(dmz.ref.Append("project"))
}

// Visibility returns a reference to field visibility of google_dns_managed_zone.
func (dmz dnsManagedZoneAttributes) Visibility() terra.StringValue {
	return terra.ReferenceAsString(dmz.ref.Append("visibility"))
}

func (dmz dnsManagedZoneAttributes) CloudLoggingConfig() terra.ListValue[dnsmanagedzone.CloudLoggingConfigAttributes] {
	return terra.ReferenceAsList[dnsmanagedzone.CloudLoggingConfigAttributes](dmz.ref.Append("cloud_logging_config"))
}

func (dmz dnsManagedZoneAttributes) DnssecConfig() terra.ListValue[dnsmanagedzone.DnssecConfigAttributes] {
	return terra.ReferenceAsList[dnsmanagedzone.DnssecConfigAttributes](dmz.ref.Append("dnssec_config"))
}

func (dmz dnsManagedZoneAttributes) ForwardingConfig() terra.ListValue[dnsmanagedzone.ForwardingConfigAttributes] {
	return terra.ReferenceAsList[dnsmanagedzone.ForwardingConfigAttributes](dmz.ref.Append("forwarding_config"))
}

func (dmz dnsManagedZoneAttributes) PeeringConfig() terra.ListValue[dnsmanagedzone.PeeringConfigAttributes] {
	return terra.ReferenceAsList[dnsmanagedzone.PeeringConfigAttributes](dmz.ref.Append("peering_config"))
}

func (dmz dnsManagedZoneAttributes) PrivateVisibilityConfig() terra.ListValue[dnsmanagedzone.PrivateVisibilityConfigAttributes] {
	return terra.ReferenceAsList[dnsmanagedzone.PrivateVisibilityConfigAttributes](dmz.ref.Append("private_visibility_config"))
}

func (dmz dnsManagedZoneAttributes) Timeouts() dnsmanagedzone.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dnsmanagedzone.TimeoutsAttributes](dmz.ref.Append("timeouts"))
}

type dnsManagedZoneState struct {
	CreationTime            string                                        `json:"creation_time"`
	Description             string                                        `json:"description"`
	DnsName                 string                                        `json:"dns_name"`
	ForceDestroy            bool                                          `json:"force_destroy"`
	Id                      string                                        `json:"id"`
	Labels                  map[string]string                             `json:"labels"`
	ManagedZoneId           float64                                       `json:"managed_zone_id"`
	Name                    string                                        `json:"name"`
	NameServers             []string                                      `json:"name_servers"`
	Project                 string                                        `json:"project"`
	Visibility              string                                        `json:"visibility"`
	CloudLoggingConfig      []dnsmanagedzone.CloudLoggingConfigState      `json:"cloud_logging_config"`
	DnssecConfig            []dnsmanagedzone.DnssecConfigState            `json:"dnssec_config"`
	ForwardingConfig        []dnsmanagedzone.ForwardingConfigState        `json:"forwarding_config"`
	PeeringConfig           []dnsmanagedzone.PeeringConfigState           `json:"peering_config"`
	PrivateVisibilityConfig []dnsmanagedzone.PrivateVisibilityConfigState `json:"private_visibility_config"`
	Timeouts                *dnsmanagedzone.TimeoutsState                 `json:"timeouts"`
}
