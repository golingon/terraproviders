// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	dnsarecord "github.com/golingon/terraproviders/azurerm/3.63.0/dnsarecord"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDnsARecord creates a new instance of [DnsARecord].
func NewDnsARecord(name string, args DnsARecordArgs) *DnsARecord {
	return &DnsARecord{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DnsARecord)(nil)

// DnsARecord represents the Terraform resource azurerm_dns_a_record.
type DnsARecord struct {
	Name      string
	Args      DnsARecordArgs
	state     *dnsARecordState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DnsARecord].
func (dar *DnsARecord) Type() string {
	return "azurerm_dns_a_record"
}

// LocalName returns the local name for [DnsARecord].
func (dar *DnsARecord) LocalName() string {
	return dar.Name
}

// Configuration returns the configuration (args) for [DnsARecord].
func (dar *DnsARecord) Configuration() interface{} {
	return dar.Args
}

// DependOn is used for other resources to depend on [DnsARecord].
func (dar *DnsARecord) DependOn() terra.Reference {
	return terra.ReferenceResource(dar)
}

// Dependencies returns the list of resources [DnsARecord] depends_on.
func (dar *DnsARecord) Dependencies() terra.Dependencies {
	return dar.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DnsARecord].
func (dar *DnsARecord) LifecycleManagement() *terra.Lifecycle {
	return dar.Lifecycle
}

// Attributes returns the attributes for [DnsARecord].
func (dar *DnsARecord) Attributes() dnsARecordAttributes {
	return dnsARecordAttributes{ref: terra.ReferenceResource(dar)}
}

// ImportState imports the given attribute values into [DnsARecord]'s state.
func (dar *DnsARecord) ImportState(av io.Reader) error {
	dar.state = &dnsARecordState{}
	if err := json.NewDecoder(av).Decode(dar.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dar.Type(), dar.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DnsARecord] has state.
func (dar *DnsARecord) State() (*dnsARecordState, bool) {
	return dar.state, dar.state != nil
}

// StateMust returns the state for [DnsARecord]. Panics if the state is nil.
func (dar *DnsARecord) StateMust() *dnsARecordState {
	if dar.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dar.Type(), dar.LocalName()))
	}
	return dar.state
}

// DnsARecordArgs contains the configurations for azurerm_dns_a_record.
type DnsARecordArgs struct {
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
	Timeouts *dnsarecord.Timeouts `hcl:"timeouts,block"`
}
type dnsARecordAttributes struct {
	ref terra.Reference
}

// Fqdn returns a reference to field fqdn of azurerm_dns_a_record.
func (dar dnsARecordAttributes) Fqdn() terra.StringValue {
	return terra.ReferenceAsString(dar.ref.Append("fqdn"))
}

// Id returns a reference to field id of azurerm_dns_a_record.
func (dar dnsARecordAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dar.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_dns_a_record.
func (dar dnsARecordAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dar.ref.Append("name"))
}

// Records returns a reference to field records of azurerm_dns_a_record.
func (dar dnsARecordAttributes) Records() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](dar.ref.Append("records"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_dns_a_record.
func (dar dnsARecordAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(dar.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_dns_a_record.
func (dar dnsARecordAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dar.ref.Append("tags"))
}

// TargetResourceId returns a reference to field target_resource_id of azurerm_dns_a_record.
func (dar dnsARecordAttributes) TargetResourceId() terra.StringValue {
	return terra.ReferenceAsString(dar.ref.Append("target_resource_id"))
}

// Ttl returns a reference to field ttl of azurerm_dns_a_record.
func (dar dnsARecordAttributes) Ttl() terra.NumberValue {
	return terra.ReferenceAsNumber(dar.ref.Append("ttl"))
}

// ZoneName returns a reference to field zone_name of azurerm_dns_a_record.
func (dar dnsARecordAttributes) ZoneName() terra.StringValue {
	return terra.ReferenceAsString(dar.ref.Append("zone_name"))
}

func (dar dnsARecordAttributes) Timeouts() dnsarecord.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dnsarecord.TimeoutsAttributes](dar.ref.Append("timeouts"))
}

type dnsARecordState struct {
	Fqdn              string                    `json:"fqdn"`
	Id                string                    `json:"id"`
	Name              string                    `json:"name"`
	Records           []string                  `json:"records"`
	ResourceGroupName string                    `json:"resource_group_name"`
	Tags              map[string]string         `json:"tags"`
	TargetResourceId  string                    `json:"target_resource_id"`
	Ttl               float64                   `json:"ttl"`
	ZoneName          string                    `json:"zone_name"`
	Timeouts          *dnsarecord.TimeoutsState `json:"timeouts"`
}
