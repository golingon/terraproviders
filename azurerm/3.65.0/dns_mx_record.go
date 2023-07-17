// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	dnsmxrecord "github.com/golingon/terraproviders/azurerm/3.65.0/dnsmxrecord"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDnsMxRecord creates a new instance of [DnsMxRecord].
func NewDnsMxRecord(name string, args DnsMxRecordArgs) *DnsMxRecord {
	return &DnsMxRecord{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DnsMxRecord)(nil)

// DnsMxRecord represents the Terraform resource azurerm_dns_mx_record.
type DnsMxRecord struct {
	Name      string
	Args      DnsMxRecordArgs
	state     *dnsMxRecordState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DnsMxRecord].
func (dmr *DnsMxRecord) Type() string {
	return "azurerm_dns_mx_record"
}

// LocalName returns the local name for [DnsMxRecord].
func (dmr *DnsMxRecord) LocalName() string {
	return dmr.Name
}

// Configuration returns the configuration (args) for [DnsMxRecord].
func (dmr *DnsMxRecord) Configuration() interface{} {
	return dmr.Args
}

// DependOn is used for other resources to depend on [DnsMxRecord].
func (dmr *DnsMxRecord) DependOn() terra.Reference {
	return terra.ReferenceResource(dmr)
}

// Dependencies returns the list of resources [DnsMxRecord] depends_on.
func (dmr *DnsMxRecord) Dependencies() terra.Dependencies {
	return dmr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DnsMxRecord].
func (dmr *DnsMxRecord) LifecycleManagement() *terra.Lifecycle {
	return dmr.Lifecycle
}

// Attributes returns the attributes for [DnsMxRecord].
func (dmr *DnsMxRecord) Attributes() dnsMxRecordAttributes {
	return dnsMxRecordAttributes{ref: terra.ReferenceResource(dmr)}
}

// ImportState imports the given attribute values into [DnsMxRecord]'s state.
func (dmr *DnsMxRecord) ImportState(av io.Reader) error {
	dmr.state = &dnsMxRecordState{}
	if err := json.NewDecoder(av).Decode(dmr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dmr.Type(), dmr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DnsMxRecord] has state.
func (dmr *DnsMxRecord) State() (*dnsMxRecordState, bool) {
	return dmr.state, dmr.state != nil
}

// StateMust returns the state for [DnsMxRecord]. Panics if the state is nil.
func (dmr *DnsMxRecord) StateMust() *dnsMxRecordState {
	if dmr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dmr.Type(), dmr.LocalName()))
	}
	return dmr.state
}

// DnsMxRecordArgs contains the configurations for azurerm_dns_mx_record.
type DnsMxRecordArgs struct {
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
	Record []dnsmxrecord.Record `hcl:"record,block" validate:"min=1"`
	// Timeouts: optional
	Timeouts *dnsmxrecord.Timeouts `hcl:"timeouts,block"`
}
type dnsMxRecordAttributes struct {
	ref terra.Reference
}

// Fqdn returns a reference to field fqdn of azurerm_dns_mx_record.
func (dmr dnsMxRecordAttributes) Fqdn() terra.StringValue {
	return terra.ReferenceAsString(dmr.ref.Append("fqdn"))
}

// Id returns a reference to field id of azurerm_dns_mx_record.
func (dmr dnsMxRecordAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dmr.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_dns_mx_record.
func (dmr dnsMxRecordAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dmr.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_dns_mx_record.
func (dmr dnsMxRecordAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(dmr.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_dns_mx_record.
func (dmr dnsMxRecordAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dmr.ref.Append("tags"))
}

// Ttl returns a reference to field ttl of azurerm_dns_mx_record.
func (dmr dnsMxRecordAttributes) Ttl() terra.NumberValue {
	return terra.ReferenceAsNumber(dmr.ref.Append("ttl"))
}

// ZoneName returns a reference to field zone_name of azurerm_dns_mx_record.
func (dmr dnsMxRecordAttributes) ZoneName() terra.StringValue {
	return terra.ReferenceAsString(dmr.ref.Append("zone_name"))
}

func (dmr dnsMxRecordAttributes) Record() terra.SetValue[dnsmxrecord.RecordAttributes] {
	return terra.ReferenceAsSet[dnsmxrecord.RecordAttributes](dmr.ref.Append("record"))
}

func (dmr dnsMxRecordAttributes) Timeouts() dnsmxrecord.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dnsmxrecord.TimeoutsAttributes](dmr.ref.Append("timeouts"))
}

type dnsMxRecordState struct {
	Fqdn              string                     `json:"fqdn"`
	Id                string                     `json:"id"`
	Name              string                     `json:"name"`
	ResourceGroupName string                     `json:"resource_group_name"`
	Tags              map[string]string          `json:"tags"`
	Ttl               float64                    `json:"ttl"`
	ZoneName          string                     `json:"zone_name"`
	Record            []dnsmxrecord.RecordState  `json:"record"`
	Timeouts          *dnsmxrecord.TimeoutsState `json:"timeouts"`
}
