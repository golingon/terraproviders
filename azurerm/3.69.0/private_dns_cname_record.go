// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	privatednscnamerecord "github.com/golingon/terraproviders/azurerm/3.69.0/privatednscnamerecord"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewPrivateDnsCnameRecord creates a new instance of [PrivateDnsCnameRecord].
func NewPrivateDnsCnameRecord(name string, args PrivateDnsCnameRecordArgs) *PrivateDnsCnameRecord {
	return &PrivateDnsCnameRecord{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*PrivateDnsCnameRecord)(nil)

// PrivateDnsCnameRecord represents the Terraform resource azurerm_private_dns_cname_record.
type PrivateDnsCnameRecord struct {
	Name      string
	Args      PrivateDnsCnameRecordArgs
	state     *privateDnsCnameRecordState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [PrivateDnsCnameRecord].
func (pdcr *PrivateDnsCnameRecord) Type() string {
	return "azurerm_private_dns_cname_record"
}

// LocalName returns the local name for [PrivateDnsCnameRecord].
func (pdcr *PrivateDnsCnameRecord) LocalName() string {
	return pdcr.Name
}

// Configuration returns the configuration (args) for [PrivateDnsCnameRecord].
func (pdcr *PrivateDnsCnameRecord) Configuration() interface{} {
	return pdcr.Args
}

// DependOn is used for other resources to depend on [PrivateDnsCnameRecord].
func (pdcr *PrivateDnsCnameRecord) DependOn() terra.Reference {
	return terra.ReferenceResource(pdcr)
}

// Dependencies returns the list of resources [PrivateDnsCnameRecord] depends_on.
func (pdcr *PrivateDnsCnameRecord) Dependencies() terra.Dependencies {
	return pdcr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [PrivateDnsCnameRecord].
func (pdcr *PrivateDnsCnameRecord) LifecycleManagement() *terra.Lifecycle {
	return pdcr.Lifecycle
}

// Attributes returns the attributes for [PrivateDnsCnameRecord].
func (pdcr *PrivateDnsCnameRecord) Attributes() privateDnsCnameRecordAttributes {
	return privateDnsCnameRecordAttributes{ref: terra.ReferenceResource(pdcr)}
}

// ImportState imports the given attribute values into [PrivateDnsCnameRecord]'s state.
func (pdcr *PrivateDnsCnameRecord) ImportState(av io.Reader) error {
	pdcr.state = &privateDnsCnameRecordState{}
	if err := json.NewDecoder(av).Decode(pdcr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", pdcr.Type(), pdcr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [PrivateDnsCnameRecord] has state.
func (pdcr *PrivateDnsCnameRecord) State() (*privateDnsCnameRecordState, bool) {
	return pdcr.state, pdcr.state != nil
}

// StateMust returns the state for [PrivateDnsCnameRecord]. Panics if the state is nil.
func (pdcr *PrivateDnsCnameRecord) StateMust() *privateDnsCnameRecordState {
	if pdcr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", pdcr.Type(), pdcr.LocalName()))
	}
	return pdcr.state
}

// PrivateDnsCnameRecordArgs contains the configurations for azurerm_private_dns_cname_record.
type PrivateDnsCnameRecordArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Record: string, required
	Record terra.StringValue `hcl:"record,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Ttl: number, required
	Ttl terra.NumberValue `hcl:"ttl,attr" validate:"required"`
	// ZoneName: string, required
	ZoneName terra.StringValue `hcl:"zone_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *privatednscnamerecord.Timeouts `hcl:"timeouts,block"`
}
type privateDnsCnameRecordAttributes struct {
	ref terra.Reference
}

// Fqdn returns a reference to field fqdn of azurerm_private_dns_cname_record.
func (pdcr privateDnsCnameRecordAttributes) Fqdn() terra.StringValue {
	return terra.ReferenceAsString(pdcr.ref.Append("fqdn"))
}

// Id returns a reference to field id of azurerm_private_dns_cname_record.
func (pdcr privateDnsCnameRecordAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(pdcr.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_private_dns_cname_record.
func (pdcr privateDnsCnameRecordAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(pdcr.ref.Append("name"))
}

// Record returns a reference to field record of azurerm_private_dns_cname_record.
func (pdcr privateDnsCnameRecordAttributes) Record() terra.StringValue {
	return terra.ReferenceAsString(pdcr.ref.Append("record"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_private_dns_cname_record.
func (pdcr privateDnsCnameRecordAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(pdcr.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_private_dns_cname_record.
func (pdcr privateDnsCnameRecordAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](pdcr.ref.Append("tags"))
}

// Ttl returns a reference to field ttl of azurerm_private_dns_cname_record.
func (pdcr privateDnsCnameRecordAttributes) Ttl() terra.NumberValue {
	return terra.ReferenceAsNumber(pdcr.ref.Append("ttl"))
}

// ZoneName returns a reference to field zone_name of azurerm_private_dns_cname_record.
func (pdcr privateDnsCnameRecordAttributes) ZoneName() terra.StringValue {
	return terra.ReferenceAsString(pdcr.ref.Append("zone_name"))
}

func (pdcr privateDnsCnameRecordAttributes) Timeouts() privatednscnamerecord.TimeoutsAttributes {
	return terra.ReferenceAsSingle[privatednscnamerecord.TimeoutsAttributes](pdcr.ref.Append("timeouts"))
}

type privateDnsCnameRecordState struct {
	Fqdn              string                               `json:"fqdn"`
	Id                string                               `json:"id"`
	Name              string                               `json:"name"`
	Record            string                               `json:"record"`
	ResourceGroupName string                               `json:"resource_group_name"`
	Tags              map[string]string                    `json:"tags"`
	Ttl               float64                              `json:"ttl"`
	ZoneName          string                               `json:"zone_name"`
	Timeouts          *privatednscnamerecord.TimeoutsState `json:"timeouts"`
}
