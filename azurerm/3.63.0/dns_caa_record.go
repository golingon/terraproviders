// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	dnscaarecord "github.com/golingon/terraproviders/azurerm/3.63.0/dnscaarecord"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDnsCaaRecord creates a new instance of [DnsCaaRecord].
func NewDnsCaaRecord(name string, args DnsCaaRecordArgs) *DnsCaaRecord {
	return &DnsCaaRecord{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DnsCaaRecord)(nil)

// DnsCaaRecord represents the Terraform resource azurerm_dns_caa_record.
type DnsCaaRecord struct {
	Name      string
	Args      DnsCaaRecordArgs
	state     *dnsCaaRecordState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DnsCaaRecord].
func (dcr *DnsCaaRecord) Type() string {
	return "azurerm_dns_caa_record"
}

// LocalName returns the local name for [DnsCaaRecord].
func (dcr *DnsCaaRecord) LocalName() string {
	return dcr.Name
}

// Configuration returns the configuration (args) for [DnsCaaRecord].
func (dcr *DnsCaaRecord) Configuration() interface{} {
	return dcr.Args
}

// DependOn is used for other resources to depend on [DnsCaaRecord].
func (dcr *DnsCaaRecord) DependOn() terra.Reference {
	return terra.ReferenceResource(dcr)
}

// Dependencies returns the list of resources [DnsCaaRecord] depends_on.
func (dcr *DnsCaaRecord) Dependencies() terra.Dependencies {
	return dcr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DnsCaaRecord].
func (dcr *DnsCaaRecord) LifecycleManagement() *terra.Lifecycle {
	return dcr.Lifecycle
}

// Attributes returns the attributes for [DnsCaaRecord].
func (dcr *DnsCaaRecord) Attributes() dnsCaaRecordAttributes {
	return dnsCaaRecordAttributes{ref: terra.ReferenceResource(dcr)}
}

// ImportState imports the given attribute values into [DnsCaaRecord]'s state.
func (dcr *DnsCaaRecord) ImportState(av io.Reader) error {
	dcr.state = &dnsCaaRecordState{}
	if err := json.NewDecoder(av).Decode(dcr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dcr.Type(), dcr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DnsCaaRecord] has state.
func (dcr *DnsCaaRecord) State() (*dnsCaaRecordState, bool) {
	return dcr.state, dcr.state != nil
}

// StateMust returns the state for [DnsCaaRecord]. Panics if the state is nil.
func (dcr *DnsCaaRecord) StateMust() *dnsCaaRecordState {
	if dcr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dcr.Type(), dcr.LocalName()))
	}
	return dcr.state
}

// DnsCaaRecordArgs contains the configurations for azurerm_dns_caa_record.
type DnsCaaRecordArgs struct {
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
	Record []dnscaarecord.Record `hcl:"record,block" validate:"min=1"`
	// Timeouts: optional
	Timeouts *dnscaarecord.Timeouts `hcl:"timeouts,block"`
}
type dnsCaaRecordAttributes struct {
	ref terra.Reference
}

// Fqdn returns a reference to field fqdn of azurerm_dns_caa_record.
func (dcr dnsCaaRecordAttributes) Fqdn() terra.StringValue {
	return terra.ReferenceAsString(dcr.ref.Append("fqdn"))
}

// Id returns a reference to field id of azurerm_dns_caa_record.
func (dcr dnsCaaRecordAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dcr.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_dns_caa_record.
func (dcr dnsCaaRecordAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dcr.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_dns_caa_record.
func (dcr dnsCaaRecordAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(dcr.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_dns_caa_record.
func (dcr dnsCaaRecordAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dcr.ref.Append("tags"))
}

// Ttl returns a reference to field ttl of azurerm_dns_caa_record.
func (dcr dnsCaaRecordAttributes) Ttl() terra.NumberValue {
	return terra.ReferenceAsNumber(dcr.ref.Append("ttl"))
}

// ZoneName returns a reference to field zone_name of azurerm_dns_caa_record.
func (dcr dnsCaaRecordAttributes) ZoneName() terra.StringValue {
	return terra.ReferenceAsString(dcr.ref.Append("zone_name"))
}

func (dcr dnsCaaRecordAttributes) Record() terra.SetValue[dnscaarecord.RecordAttributes] {
	return terra.ReferenceAsSet[dnscaarecord.RecordAttributes](dcr.ref.Append("record"))
}

func (dcr dnsCaaRecordAttributes) Timeouts() dnscaarecord.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dnscaarecord.TimeoutsAttributes](dcr.ref.Append("timeouts"))
}

type dnsCaaRecordState struct {
	Fqdn              string                      `json:"fqdn"`
	Id                string                      `json:"id"`
	Name              string                      `json:"name"`
	ResourceGroupName string                      `json:"resource_group_name"`
	Tags              map[string]string           `json:"tags"`
	Ttl               float64                     `json:"ttl"`
	ZoneName          string                      `json:"zone_name"`
	Record            []dnscaarecord.RecordState  `json:"record"`
	Timeouts          *dnscaarecord.TimeoutsState `json:"timeouts"`
}
