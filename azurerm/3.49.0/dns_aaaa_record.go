// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	dnsaaaarecord "github.com/golingon/terraproviders/azurerm/3.49.0/dnsaaaarecord"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDnsAaaaRecord creates a new instance of [DnsAaaaRecord].
func NewDnsAaaaRecord(name string, args DnsAaaaRecordArgs) *DnsAaaaRecord {
	return &DnsAaaaRecord{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DnsAaaaRecord)(nil)

// DnsAaaaRecord represents the Terraform resource azurerm_dns_aaaa_record.
type DnsAaaaRecord struct {
	Name      string
	Args      DnsAaaaRecordArgs
	state     *dnsAaaaRecordState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DnsAaaaRecord].
func (dar *DnsAaaaRecord) Type() string {
	return "azurerm_dns_aaaa_record"
}

// LocalName returns the local name for [DnsAaaaRecord].
func (dar *DnsAaaaRecord) LocalName() string {
	return dar.Name
}

// Configuration returns the configuration (args) for [DnsAaaaRecord].
func (dar *DnsAaaaRecord) Configuration() interface{} {
	return dar.Args
}

// DependOn is used for other resources to depend on [DnsAaaaRecord].
func (dar *DnsAaaaRecord) DependOn() terra.Reference {
	return terra.ReferenceResource(dar)
}

// Dependencies returns the list of resources [DnsAaaaRecord] depends_on.
func (dar *DnsAaaaRecord) Dependencies() terra.Dependencies {
	return dar.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DnsAaaaRecord].
func (dar *DnsAaaaRecord) LifecycleManagement() *terra.Lifecycle {
	return dar.Lifecycle
}

// Attributes returns the attributes for [DnsAaaaRecord].
func (dar *DnsAaaaRecord) Attributes() dnsAaaaRecordAttributes {
	return dnsAaaaRecordAttributes{ref: terra.ReferenceResource(dar)}
}

// ImportState imports the given attribute values into [DnsAaaaRecord]'s state.
func (dar *DnsAaaaRecord) ImportState(av io.Reader) error {
	dar.state = &dnsAaaaRecordState{}
	if err := json.NewDecoder(av).Decode(dar.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dar.Type(), dar.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DnsAaaaRecord] has state.
func (dar *DnsAaaaRecord) State() (*dnsAaaaRecordState, bool) {
	return dar.state, dar.state != nil
}

// StateMust returns the state for [DnsAaaaRecord]. Panics if the state is nil.
func (dar *DnsAaaaRecord) StateMust() *dnsAaaaRecordState {
	if dar.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dar.Type(), dar.LocalName()))
	}
	return dar.state
}

// DnsAaaaRecordArgs contains the configurations for azurerm_dns_aaaa_record.
type DnsAaaaRecordArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Records: set of string, optional
	Records terra.SetValue[terra.StringValue] `hcl:"records,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TargetResourceId: string, optional
	TargetResourceId terra.StringValue `hcl:"target_resource_id,attr"`
	// Ttl: number, required
	Ttl terra.NumberValue `hcl:"ttl,attr" validate:"required"`
	// ZoneName: string, required
	ZoneName terra.StringValue `hcl:"zone_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *dnsaaaarecord.Timeouts `hcl:"timeouts,block"`
}
type dnsAaaaRecordAttributes struct {
	ref terra.Reference
}

// Fqdn returns a reference to field fqdn of azurerm_dns_aaaa_record.
func (dar dnsAaaaRecordAttributes) Fqdn() terra.StringValue {
	return terra.ReferenceAsString(dar.ref.Append("fqdn"))
}

// Id returns a reference to field id of azurerm_dns_aaaa_record.
func (dar dnsAaaaRecordAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dar.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_dns_aaaa_record.
func (dar dnsAaaaRecordAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dar.ref.Append("name"))
}

// Records returns a reference to field records of azurerm_dns_aaaa_record.
func (dar dnsAaaaRecordAttributes) Records() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](dar.ref.Append("records"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_dns_aaaa_record.
func (dar dnsAaaaRecordAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(dar.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_dns_aaaa_record.
func (dar dnsAaaaRecordAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dar.ref.Append("tags"))
}

// TargetResourceId returns a reference to field target_resource_id of azurerm_dns_aaaa_record.
func (dar dnsAaaaRecordAttributes) TargetResourceId() terra.StringValue {
	return terra.ReferenceAsString(dar.ref.Append("target_resource_id"))
}

// Ttl returns a reference to field ttl of azurerm_dns_aaaa_record.
func (dar dnsAaaaRecordAttributes) Ttl() terra.NumberValue {
	return terra.ReferenceAsNumber(dar.ref.Append("ttl"))
}

// ZoneName returns a reference to field zone_name of azurerm_dns_aaaa_record.
func (dar dnsAaaaRecordAttributes) ZoneName() terra.StringValue {
	return terra.ReferenceAsString(dar.ref.Append("zone_name"))
}

func (dar dnsAaaaRecordAttributes) Timeouts() dnsaaaarecord.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dnsaaaarecord.TimeoutsAttributes](dar.ref.Append("timeouts"))
}

type dnsAaaaRecordState struct {
	Fqdn              string                       `json:"fqdn"`
	Id                string                       `json:"id"`
	Name              string                       `json:"name"`
	Records           []string                     `json:"records"`
	ResourceGroupName string                       `json:"resource_group_name"`
	Tags              map[string]string            `json:"tags"`
	TargetResourceId  string                       `json:"target_resource_id"`
	Ttl               float64                      `json:"ttl"`
	ZoneName          string                       `json:"zone_name"`
	Timeouts          *dnsaaaarecord.TimeoutsState `json:"timeouts"`
}
