// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	dnszone "github.com/golingon/terraproviders/azurerm/3.64.0/dnszone"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDnsZone creates a new instance of [DnsZone].
func NewDnsZone(name string, args DnsZoneArgs) *DnsZone {
	return &DnsZone{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DnsZone)(nil)

// DnsZone represents the Terraform resource azurerm_dns_zone.
type DnsZone struct {
	Name      string
	Args      DnsZoneArgs
	state     *dnsZoneState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DnsZone].
func (dz *DnsZone) Type() string {
	return "azurerm_dns_zone"
}

// LocalName returns the local name for [DnsZone].
func (dz *DnsZone) LocalName() string {
	return dz.Name
}

// Configuration returns the configuration (args) for [DnsZone].
func (dz *DnsZone) Configuration() interface{} {
	return dz.Args
}

// DependOn is used for other resources to depend on [DnsZone].
func (dz *DnsZone) DependOn() terra.Reference {
	return terra.ReferenceResource(dz)
}

// Dependencies returns the list of resources [DnsZone] depends_on.
func (dz *DnsZone) Dependencies() terra.Dependencies {
	return dz.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DnsZone].
func (dz *DnsZone) LifecycleManagement() *terra.Lifecycle {
	return dz.Lifecycle
}

// Attributes returns the attributes for [DnsZone].
func (dz *DnsZone) Attributes() dnsZoneAttributes {
	return dnsZoneAttributes{ref: terra.ReferenceResource(dz)}
}

// ImportState imports the given attribute values into [DnsZone]'s state.
func (dz *DnsZone) ImportState(av io.Reader) error {
	dz.state = &dnsZoneState{}
	if err := json.NewDecoder(av).Decode(dz.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dz.Type(), dz.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DnsZone] has state.
func (dz *DnsZone) State() (*dnsZoneState, bool) {
	return dz.state, dz.state != nil
}

// StateMust returns the state for [DnsZone]. Panics if the state is nil.
func (dz *DnsZone) StateMust() *dnsZoneState {
	if dz.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dz.Type(), dz.LocalName()))
	}
	return dz.state
}

// DnsZoneArgs contains the configurations for azurerm_dns_zone.
type DnsZoneArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// SoaRecord: optional
	SoaRecord *dnszone.SoaRecord `hcl:"soa_record,block"`
	// Timeouts: optional
	Timeouts *dnszone.Timeouts `hcl:"timeouts,block"`
}
type dnsZoneAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_dns_zone.
func (dz dnsZoneAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dz.ref.Append("id"))
}

// MaxNumberOfRecordSets returns a reference to field max_number_of_record_sets of azurerm_dns_zone.
func (dz dnsZoneAttributes) MaxNumberOfRecordSets() terra.NumberValue {
	return terra.ReferenceAsNumber(dz.ref.Append("max_number_of_record_sets"))
}

// Name returns a reference to field name of azurerm_dns_zone.
func (dz dnsZoneAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dz.ref.Append("name"))
}

// NameServers returns a reference to field name_servers of azurerm_dns_zone.
func (dz dnsZoneAttributes) NameServers() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](dz.ref.Append("name_servers"))
}

// NumberOfRecordSets returns a reference to field number_of_record_sets of azurerm_dns_zone.
func (dz dnsZoneAttributes) NumberOfRecordSets() terra.NumberValue {
	return terra.ReferenceAsNumber(dz.ref.Append("number_of_record_sets"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_dns_zone.
func (dz dnsZoneAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(dz.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_dns_zone.
func (dz dnsZoneAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dz.ref.Append("tags"))
}

func (dz dnsZoneAttributes) SoaRecord() terra.ListValue[dnszone.SoaRecordAttributes] {
	return terra.ReferenceAsList[dnszone.SoaRecordAttributes](dz.ref.Append("soa_record"))
}

func (dz dnsZoneAttributes) Timeouts() dnszone.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dnszone.TimeoutsAttributes](dz.ref.Append("timeouts"))
}

type dnsZoneState struct {
	Id                    string                   `json:"id"`
	MaxNumberOfRecordSets float64                  `json:"max_number_of_record_sets"`
	Name                  string                   `json:"name"`
	NameServers           []string                 `json:"name_servers"`
	NumberOfRecordSets    float64                  `json:"number_of_record_sets"`
	ResourceGroupName     string                   `json:"resource_group_name"`
	Tags                  map[string]string        `json:"tags"`
	SoaRecord             []dnszone.SoaRecordState `json:"soa_record"`
	Timeouts              *dnszone.TimeoutsState   `json:"timeouts"`
}
