// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	privatednszone "github.com/golingon/terraproviders/azurerm/3.66.0/privatednszone"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewPrivateDnsZone creates a new instance of [PrivateDnsZone].
func NewPrivateDnsZone(name string, args PrivateDnsZoneArgs) *PrivateDnsZone {
	return &PrivateDnsZone{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*PrivateDnsZone)(nil)

// PrivateDnsZone represents the Terraform resource azurerm_private_dns_zone.
type PrivateDnsZone struct {
	Name      string
	Args      PrivateDnsZoneArgs
	state     *privateDnsZoneState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [PrivateDnsZone].
func (pdz *PrivateDnsZone) Type() string {
	return "azurerm_private_dns_zone"
}

// LocalName returns the local name for [PrivateDnsZone].
func (pdz *PrivateDnsZone) LocalName() string {
	return pdz.Name
}

// Configuration returns the configuration (args) for [PrivateDnsZone].
func (pdz *PrivateDnsZone) Configuration() interface{} {
	return pdz.Args
}

// DependOn is used for other resources to depend on [PrivateDnsZone].
func (pdz *PrivateDnsZone) DependOn() terra.Reference {
	return terra.ReferenceResource(pdz)
}

// Dependencies returns the list of resources [PrivateDnsZone] depends_on.
func (pdz *PrivateDnsZone) Dependencies() terra.Dependencies {
	return pdz.DependsOn
}

// LifecycleManagement returns the lifecycle block for [PrivateDnsZone].
func (pdz *PrivateDnsZone) LifecycleManagement() *terra.Lifecycle {
	return pdz.Lifecycle
}

// Attributes returns the attributes for [PrivateDnsZone].
func (pdz *PrivateDnsZone) Attributes() privateDnsZoneAttributes {
	return privateDnsZoneAttributes{ref: terra.ReferenceResource(pdz)}
}

// ImportState imports the given attribute values into [PrivateDnsZone]'s state.
func (pdz *PrivateDnsZone) ImportState(av io.Reader) error {
	pdz.state = &privateDnsZoneState{}
	if err := json.NewDecoder(av).Decode(pdz.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", pdz.Type(), pdz.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [PrivateDnsZone] has state.
func (pdz *PrivateDnsZone) State() (*privateDnsZoneState, bool) {
	return pdz.state, pdz.state != nil
}

// StateMust returns the state for [PrivateDnsZone]. Panics if the state is nil.
func (pdz *PrivateDnsZone) StateMust() *privateDnsZoneState {
	if pdz.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", pdz.Type(), pdz.LocalName()))
	}
	return pdz.state
}

// PrivateDnsZoneArgs contains the configurations for azurerm_private_dns_zone.
type PrivateDnsZoneArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// SoaRecord: optional
	SoaRecord *privatednszone.SoaRecord `hcl:"soa_record,block"`
	// Timeouts: optional
	Timeouts *privatednszone.Timeouts `hcl:"timeouts,block"`
}
type privateDnsZoneAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_private_dns_zone.
func (pdz privateDnsZoneAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(pdz.ref.Append("id"))
}

// MaxNumberOfRecordSets returns a reference to field max_number_of_record_sets of azurerm_private_dns_zone.
func (pdz privateDnsZoneAttributes) MaxNumberOfRecordSets() terra.NumberValue {
	return terra.ReferenceAsNumber(pdz.ref.Append("max_number_of_record_sets"))
}

// MaxNumberOfVirtualNetworkLinks returns a reference to field max_number_of_virtual_network_links of azurerm_private_dns_zone.
func (pdz privateDnsZoneAttributes) MaxNumberOfVirtualNetworkLinks() terra.NumberValue {
	return terra.ReferenceAsNumber(pdz.ref.Append("max_number_of_virtual_network_links"))
}

// MaxNumberOfVirtualNetworkLinksWithRegistration returns a reference to field max_number_of_virtual_network_links_with_registration of azurerm_private_dns_zone.
func (pdz privateDnsZoneAttributes) MaxNumberOfVirtualNetworkLinksWithRegistration() terra.NumberValue {
	return terra.ReferenceAsNumber(pdz.ref.Append("max_number_of_virtual_network_links_with_registration"))
}

// Name returns a reference to field name of azurerm_private_dns_zone.
func (pdz privateDnsZoneAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(pdz.ref.Append("name"))
}

// NumberOfRecordSets returns a reference to field number_of_record_sets of azurerm_private_dns_zone.
func (pdz privateDnsZoneAttributes) NumberOfRecordSets() terra.NumberValue {
	return terra.ReferenceAsNumber(pdz.ref.Append("number_of_record_sets"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_private_dns_zone.
func (pdz privateDnsZoneAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(pdz.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_private_dns_zone.
func (pdz privateDnsZoneAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](pdz.ref.Append("tags"))
}

func (pdz privateDnsZoneAttributes) SoaRecord() terra.ListValue[privatednszone.SoaRecordAttributes] {
	return terra.ReferenceAsList[privatednszone.SoaRecordAttributes](pdz.ref.Append("soa_record"))
}

func (pdz privateDnsZoneAttributes) Timeouts() privatednszone.TimeoutsAttributes {
	return terra.ReferenceAsSingle[privatednszone.TimeoutsAttributes](pdz.ref.Append("timeouts"))
}

type privateDnsZoneState struct {
	Id                                             string                          `json:"id"`
	MaxNumberOfRecordSets                          float64                         `json:"max_number_of_record_sets"`
	MaxNumberOfVirtualNetworkLinks                 float64                         `json:"max_number_of_virtual_network_links"`
	MaxNumberOfVirtualNetworkLinksWithRegistration float64                         `json:"max_number_of_virtual_network_links_with_registration"`
	Name                                           string                          `json:"name"`
	NumberOfRecordSets                             float64                         `json:"number_of_record_sets"`
	ResourceGroupName                              string                          `json:"resource_group_name"`
	Tags                                           map[string]string               `json:"tags"`
	SoaRecord                                      []privatednszone.SoaRecordState `json:"soa_record"`
	Timeouts                                       *privatednszone.TimeoutsState   `json:"timeouts"`
}