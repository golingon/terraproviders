// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	privatednsmxrecord "github.com/golingon/terraproviders/azurerm/3.67.0/privatednsmxrecord"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewPrivateDnsMxRecord creates a new instance of [PrivateDnsMxRecord].
func NewPrivateDnsMxRecord(name string, args PrivateDnsMxRecordArgs) *PrivateDnsMxRecord {
	return &PrivateDnsMxRecord{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*PrivateDnsMxRecord)(nil)

// PrivateDnsMxRecord represents the Terraform resource azurerm_private_dns_mx_record.
type PrivateDnsMxRecord struct {
	Name      string
	Args      PrivateDnsMxRecordArgs
	state     *privateDnsMxRecordState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [PrivateDnsMxRecord].
func (pdmr *PrivateDnsMxRecord) Type() string {
	return "azurerm_private_dns_mx_record"
}

// LocalName returns the local name for [PrivateDnsMxRecord].
func (pdmr *PrivateDnsMxRecord) LocalName() string {
	return pdmr.Name
}

// Configuration returns the configuration (args) for [PrivateDnsMxRecord].
func (pdmr *PrivateDnsMxRecord) Configuration() interface{} {
	return pdmr.Args
}

// DependOn is used for other resources to depend on [PrivateDnsMxRecord].
func (pdmr *PrivateDnsMxRecord) DependOn() terra.Reference {
	return terra.ReferenceResource(pdmr)
}

// Dependencies returns the list of resources [PrivateDnsMxRecord] depends_on.
func (pdmr *PrivateDnsMxRecord) Dependencies() terra.Dependencies {
	return pdmr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [PrivateDnsMxRecord].
func (pdmr *PrivateDnsMxRecord) LifecycleManagement() *terra.Lifecycle {
	return pdmr.Lifecycle
}

// Attributes returns the attributes for [PrivateDnsMxRecord].
func (pdmr *PrivateDnsMxRecord) Attributes() privateDnsMxRecordAttributes {
	return privateDnsMxRecordAttributes{ref: terra.ReferenceResource(pdmr)}
}

// ImportState imports the given attribute values into [PrivateDnsMxRecord]'s state.
func (pdmr *PrivateDnsMxRecord) ImportState(av io.Reader) error {
	pdmr.state = &privateDnsMxRecordState{}
	if err := json.NewDecoder(av).Decode(pdmr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", pdmr.Type(), pdmr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [PrivateDnsMxRecord] has state.
func (pdmr *PrivateDnsMxRecord) State() (*privateDnsMxRecordState, bool) {
	return pdmr.state, pdmr.state != nil
}

// StateMust returns the state for [PrivateDnsMxRecord]. Panics if the state is nil.
func (pdmr *PrivateDnsMxRecord) StateMust() *privateDnsMxRecordState {
	if pdmr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", pdmr.Type(), pdmr.LocalName()))
	}
	return pdmr.state
}

// PrivateDnsMxRecordArgs contains the configurations for azurerm_private_dns_mx_record.
type PrivateDnsMxRecordArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Ttl: number, required
	Ttl terra.NumberValue `hcl:"ttl,attr" validate:"required"`
	// ZoneName: string, required
	ZoneName terra.StringValue `hcl:"zone_name,attr" validate:"required"`
	// Record: min=1
	Record []privatednsmxrecord.Record `hcl:"record,block" validate:"min=1"`
	// Timeouts: optional
	Timeouts *privatednsmxrecord.Timeouts `hcl:"timeouts,block"`
}
type privateDnsMxRecordAttributes struct {
	ref terra.Reference
}

// Fqdn returns a reference to field fqdn of azurerm_private_dns_mx_record.
func (pdmr privateDnsMxRecordAttributes) Fqdn() terra.StringValue {
	return terra.ReferenceAsString(pdmr.ref.Append("fqdn"))
}

// Id returns a reference to field id of azurerm_private_dns_mx_record.
func (pdmr privateDnsMxRecordAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(pdmr.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_private_dns_mx_record.
func (pdmr privateDnsMxRecordAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(pdmr.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_private_dns_mx_record.
func (pdmr privateDnsMxRecordAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(pdmr.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_private_dns_mx_record.
func (pdmr privateDnsMxRecordAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](pdmr.ref.Append("tags"))
}

// Ttl returns a reference to field ttl of azurerm_private_dns_mx_record.
func (pdmr privateDnsMxRecordAttributes) Ttl() terra.NumberValue {
	return terra.ReferenceAsNumber(pdmr.ref.Append("ttl"))
}

// ZoneName returns a reference to field zone_name of azurerm_private_dns_mx_record.
func (pdmr privateDnsMxRecordAttributes) ZoneName() terra.StringValue {
	return terra.ReferenceAsString(pdmr.ref.Append("zone_name"))
}

func (pdmr privateDnsMxRecordAttributes) Record() terra.SetValue[privatednsmxrecord.RecordAttributes] {
	return terra.ReferenceAsSet[privatednsmxrecord.RecordAttributes](pdmr.ref.Append("record"))
}

func (pdmr privateDnsMxRecordAttributes) Timeouts() privatednsmxrecord.TimeoutsAttributes {
	return terra.ReferenceAsSingle[privatednsmxrecord.TimeoutsAttributes](pdmr.ref.Append("timeouts"))
}

type privateDnsMxRecordState struct {
	Fqdn              string                            `json:"fqdn"`
	Id                string                            `json:"id"`
	Name              string                            `json:"name"`
	ResourceGroupName string                            `json:"resource_group_name"`
	Tags              map[string]string                 `json:"tags"`
	Ttl               float64                           `json:"ttl"`
	ZoneName          string                            `json:"zone_name"`
	Record            []privatednsmxrecord.RecordState  `json:"record"`
	Timeouts          *privatednsmxrecord.TimeoutsState `json:"timeouts"`
}
