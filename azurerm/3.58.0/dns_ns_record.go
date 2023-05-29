// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	dnsnsrecord "github.com/golingon/terraproviders/azurerm/3.58.0/dnsnsrecord"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDnsNsRecord creates a new instance of [DnsNsRecord].
func NewDnsNsRecord(name string, args DnsNsRecordArgs) *DnsNsRecord {
	return &DnsNsRecord{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DnsNsRecord)(nil)

// DnsNsRecord represents the Terraform resource azurerm_dns_ns_record.
type DnsNsRecord struct {
	Name      string
	Args      DnsNsRecordArgs
	state     *dnsNsRecordState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DnsNsRecord].
func (dnr *DnsNsRecord) Type() string {
	return "azurerm_dns_ns_record"
}

// LocalName returns the local name for [DnsNsRecord].
func (dnr *DnsNsRecord) LocalName() string {
	return dnr.Name
}

// Configuration returns the configuration (args) for [DnsNsRecord].
func (dnr *DnsNsRecord) Configuration() interface{} {
	return dnr.Args
}

// DependOn is used for other resources to depend on [DnsNsRecord].
func (dnr *DnsNsRecord) DependOn() terra.Reference {
	return terra.ReferenceResource(dnr)
}

// Dependencies returns the list of resources [DnsNsRecord] depends_on.
func (dnr *DnsNsRecord) Dependencies() terra.Dependencies {
	return dnr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DnsNsRecord].
func (dnr *DnsNsRecord) LifecycleManagement() *terra.Lifecycle {
	return dnr.Lifecycle
}

// Attributes returns the attributes for [DnsNsRecord].
func (dnr *DnsNsRecord) Attributes() dnsNsRecordAttributes {
	return dnsNsRecordAttributes{ref: terra.ReferenceResource(dnr)}
}

// ImportState imports the given attribute values into [DnsNsRecord]'s state.
func (dnr *DnsNsRecord) ImportState(av io.Reader) error {
	dnr.state = &dnsNsRecordState{}
	if err := json.NewDecoder(av).Decode(dnr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dnr.Type(), dnr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DnsNsRecord] has state.
func (dnr *DnsNsRecord) State() (*dnsNsRecordState, bool) {
	return dnr.state, dnr.state != nil
}

// StateMust returns the state for [DnsNsRecord]. Panics if the state is nil.
func (dnr *DnsNsRecord) StateMust() *dnsNsRecordState {
	if dnr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dnr.Type(), dnr.LocalName()))
	}
	return dnr.state
}

// DnsNsRecordArgs contains the configurations for azurerm_dns_ns_record.
type DnsNsRecordArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Records: list of string, required
	Records terra.ListValue[terra.StringValue] `hcl:"records,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Ttl: number, required
	Ttl terra.NumberValue `hcl:"ttl,attr" validate:"required"`
	// ZoneName: string, required
	ZoneName terra.StringValue `hcl:"zone_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *dnsnsrecord.Timeouts `hcl:"timeouts,block"`
}
type dnsNsRecordAttributes struct {
	ref terra.Reference
}

// Fqdn returns a reference to field fqdn of azurerm_dns_ns_record.
func (dnr dnsNsRecordAttributes) Fqdn() terra.StringValue {
	return terra.ReferenceAsString(dnr.ref.Append("fqdn"))
}

// Id returns a reference to field id of azurerm_dns_ns_record.
func (dnr dnsNsRecordAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dnr.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_dns_ns_record.
func (dnr dnsNsRecordAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dnr.ref.Append("name"))
}

// Records returns a reference to field records of azurerm_dns_ns_record.
func (dnr dnsNsRecordAttributes) Records() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](dnr.ref.Append("records"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_dns_ns_record.
func (dnr dnsNsRecordAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(dnr.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_dns_ns_record.
func (dnr dnsNsRecordAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dnr.ref.Append("tags"))
}

// Ttl returns a reference to field ttl of azurerm_dns_ns_record.
func (dnr dnsNsRecordAttributes) Ttl() terra.NumberValue {
	return terra.ReferenceAsNumber(dnr.ref.Append("ttl"))
}

// ZoneName returns a reference to field zone_name of azurerm_dns_ns_record.
func (dnr dnsNsRecordAttributes) ZoneName() terra.StringValue {
	return terra.ReferenceAsString(dnr.ref.Append("zone_name"))
}

func (dnr dnsNsRecordAttributes) Timeouts() dnsnsrecord.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dnsnsrecord.TimeoutsAttributes](dnr.ref.Append("timeouts"))
}

type dnsNsRecordState struct {
	Fqdn              string                     `json:"fqdn"`
	Id                string                     `json:"id"`
	Name              string                     `json:"name"`
	Records           []string                   `json:"records"`
	ResourceGroupName string                     `json:"resource_group_name"`
	Tags              map[string]string          `json:"tags"`
	Ttl               float64                    `json:"ttl"`
	ZoneName          string                     `json:"zone_name"`
	Timeouts          *dnsnsrecord.TimeoutsState `json:"timeouts"`
}