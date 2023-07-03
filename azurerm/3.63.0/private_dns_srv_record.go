// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	privatednssrvrecord "github.com/golingon/terraproviders/azurerm/3.63.0/privatednssrvrecord"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewPrivateDnsSrvRecord creates a new instance of [PrivateDnsSrvRecord].
func NewPrivateDnsSrvRecord(name string, args PrivateDnsSrvRecordArgs) *PrivateDnsSrvRecord {
	return &PrivateDnsSrvRecord{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*PrivateDnsSrvRecord)(nil)

// PrivateDnsSrvRecord represents the Terraform resource azurerm_private_dns_srv_record.
type PrivateDnsSrvRecord struct {
	Name      string
	Args      PrivateDnsSrvRecordArgs
	state     *privateDnsSrvRecordState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [PrivateDnsSrvRecord].
func (pdsr *PrivateDnsSrvRecord) Type() string {
	return "azurerm_private_dns_srv_record"
}

// LocalName returns the local name for [PrivateDnsSrvRecord].
func (pdsr *PrivateDnsSrvRecord) LocalName() string {
	return pdsr.Name
}

// Configuration returns the configuration (args) for [PrivateDnsSrvRecord].
func (pdsr *PrivateDnsSrvRecord) Configuration() interface{} {
	return pdsr.Args
}

// DependOn is used for other resources to depend on [PrivateDnsSrvRecord].
func (pdsr *PrivateDnsSrvRecord) DependOn() terra.Reference {
	return terra.ReferenceResource(pdsr)
}

// Dependencies returns the list of resources [PrivateDnsSrvRecord] depends_on.
func (pdsr *PrivateDnsSrvRecord) Dependencies() terra.Dependencies {
	return pdsr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [PrivateDnsSrvRecord].
func (pdsr *PrivateDnsSrvRecord) LifecycleManagement() *terra.Lifecycle {
	return pdsr.Lifecycle
}

// Attributes returns the attributes for [PrivateDnsSrvRecord].
func (pdsr *PrivateDnsSrvRecord) Attributes() privateDnsSrvRecordAttributes {
	return privateDnsSrvRecordAttributes{ref: terra.ReferenceResource(pdsr)}
}

// ImportState imports the given attribute values into [PrivateDnsSrvRecord]'s state.
func (pdsr *PrivateDnsSrvRecord) ImportState(av io.Reader) error {
	pdsr.state = &privateDnsSrvRecordState{}
	if err := json.NewDecoder(av).Decode(pdsr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", pdsr.Type(), pdsr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [PrivateDnsSrvRecord] has state.
func (pdsr *PrivateDnsSrvRecord) State() (*privateDnsSrvRecordState, bool) {
	return pdsr.state, pdsr.state != nil
}

// StateMust returns the state for [PrivateDnsSrvRecord]. Panics if the state is nil.
func (pdsr *PrivateDnsSrvRecord) StateMust() *privateDnsSrvRecordState {
	if pdsr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", pdsr.Type(), pdsr.LocalName()))
	}
	return pdsr.state
}

// PrivateDnsSrvRecordArgs contains the configurations for azurerm_private_dns_srv_record.
type PrivateDnsSrvRecordArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Ttl: number, required
	Ttl terra.NumberValue `hcl:"ttl,attr" validate:"required"`
	// ZoneName: string, required
	ZoneName terra.StringValue `hcl:"zone_name,attr" validate:"required"`
	// Record: min=1
	Record []privatednssrvrecord.Record `hcl:"record,block" validate:"min=1"`
	// Timeouts: optional
	Timeouts *privatednssrvrecord.Timeouts `hcl:"timeouts,block"`
}
type privateDnsSrvRecordAttributes struct {
	ref terra.Reference
}

// Fqdn returns a reference to field fqdn of azurerm_private_dns_srv_record.
func (pdsr privateDnsSrvRecordAttributes) Fqdn() terra.StringValue {
	return terra.ReferenceAsString(pdsr.ref.Append("fqdn"))
}

// Id returns a reference to field id of azurerm_private_dns_srv_record.
func (pdsr privateDnsSrvRecordAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(pdsr.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_private_dns_srv_record.
func (pdsr privateDnsSrvRecordAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(pdsr.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_private_dns_srv_record.
func (pdsr privateDnsSrvRecordAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(pdsr.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_private_dns_srv_record.
func (pdsr privateDnsSrvRecordAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](pdsr.ref.Append("tags"))
}

// Ttl returns a reference to field ttl of azurerm_private_dns_srv_record.
func (pdsr privateDnsSrvRecordAttributes) Ttl() terra.NumberValue {
	return terra.ReferenceAsNumber(pdsr.ref.Append("ttl"))
}

// ZoneName returns a reference to field zone_name of azurerm_private_dns_srv_record.
func (pdsr privateDnsSrvRecordAttributes) ZoneName() terra.StringValue {
	return terra.ReferenceAsString(pdsr.ref.Append("zone_name"))
}

func (pdsr privateDnsSrvRecordAttributes) Record() terra.SetValue[privatednssrvrecord.RecordAttributes] {
	return terra.ReferenceAsSet[privatednssrvrecord.RecordAttributes](pdsr.ref.Append("record"))
}

func (pdsr privateDnsSrvRecordAttributes) Timeouts() privatednssrvrecord.TimeoutsAttributes {
	return terra.ReferenceAsSingle[privatednssrvrecord.TimeoutsAttributes](pdsr.ref.Append("timeouts"))
}

type privateDnsSrvRecordState struct {
	Fqdn              string                             `json:"fqdn"`
	Id                string                             `json:"id"`
	Name              string                             `json:"name"`
	ResourceGroupName string                             `json:"resource_group_name"`
	Tags              map[string]string                  `json:"tags"`
	Ttl               float64                            `json:"ttl"`
	ZoneName          string                             `json:"zone_name"`
	Record            []privatednssrvrecord.RecordState  `json:"record"`
	Timeouts          *privatednssrvrecord.TimeoutsState `json:"timeouts"`
}
