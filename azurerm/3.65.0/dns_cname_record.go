// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	dnscnamerecord "github.com/golingon/terraproviders/azurerm/3.65.0/dnscnamerecord"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDnsCnameRecord creates a new instance of [DnsCnameRecord].
func NewDnsCnameRecord(name string, args DnsCnameRecordArgs) *DnsCnameRecord {
	return &DnsCnameRecord{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DnsCnameRecord)(nil)

// DnsCnameRecord represents the Terraform resource azurerm_dns_cname_record.
type DnsCnameRecord struct {
	Name      string
	Args      DnsCnameRecordArgs
	state     *dnsCnameRecordState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DnsCnameRecord].
func (dcr *DnsCnameRecord) Type() string {
	return "azurerm_dns_cname_record"
}

// LocalName returns the local name for [DnsCnameRecord].
func (dcr *DnsCnameRecord) LocalName() string {
	return dcr.Name
}

// Configuration returns the configuration (args) for [DnsCnameRecord].
func (dcr *DnsCnameRecord) Configuration() interface{} {
	return dcr.Args
}

// DependOn is used for other resources to depend on [DnsCnameRecord].
func (dcr *DnsCnameRecord) DependOn() terra.Reference {
	return terra.ReferenceResource(dcr)
}

// Dependencies returns the list of resources [DnsCnameRecord] depends_on.
func (dcr *DnsCnameRecord) Dependencies() terra.Dependencies {
	return dcr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DnsCnameRecord].
func (dcr *DnsCnameRecord) LifecycleManagement() *terra.Lifecycle {
	return dcr.Lifecycle
}

// Attributes returns the attributes for [DnsCnameRecord].
func (dcr *DnsCnameRecord) Attributes() dnsCnameRecordAttributes {
	return dnsCnameRecordAttributes{ref: terra.ReferenceResource(dcr)}
}

// ImportState imports the given attribute values into [DnsCnameRecord]'s state.
func (dcr *DnsCnameRecord) ImportState(av io.Reader) error {
	dcr.state = &dnsCnameRecordState{}
	if err := json.NewDecoder(av).Decode(dcr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dcr.Type(), dcr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DnsCnameRecord] has state.
func (dcr *DnsCnameRecord) State() (*dnsCnameRecordState, bool) {
	return dcr.state, dcr.state != nil
}

// StateMust returns the state for [DnsCnameRecord]. Panics if the state is nil.
func (dcr *DnsCnameRecord) StateMust() *dnsCnameRecordState {
	if dcr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dcr.Type(), dcr.LocalName()))
	}
	return dcr.state
}

// DnsCnameRecordArgs contains the configurations for azurerm_dns_cname_record.
type DnsCnameRecordArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Record: string, optional
	Record terra.StringValue `hcl:"record,attr"`
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
	Timeouts *dnscnamerecord.Timeouts `hcl:"timeouts,block"`
}
type dnsCnameRecordAttributes struct {
	ref terra.Reference
}

// Fqdn returns a reference to field fqdn of azurerm_dns_cname_record.
func (dcr dnsCnameRecordAttributes) Fqdn() terra.StringValue {
	return terra.ReferenceAsString(dcr.ref.Append("fqdn"))
}

// Id returns a reference to field id of azurerm_dns_cname_record.
func (dcr dnsCnameRecordAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dcr.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_dns_cname_record.
func (dcr dnsCnameRecordAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dcr.ref.Append("name"))
}

// Record returns a reference to field record of azurerm_dns_cname_record.
func (dcr dnsCnameRecordAttributes) Record() terra.StringValue {
	return terra.ReferenceAsString(dcr.ref.Append("record"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_dns_cname_record.
func (dcr dnsCnameRecordAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(dcr.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_dns_cname_record.
func (dcr dnsCnameRecordAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dcr.ref.Append("tags"))
}

// TargetResourceId returns a reference to field target_resource_id of azurerm_dns_cname_record.
func (dcr dnsCnameRecordAttributes) TargetResourceId() terra.StringValue {
	return terra.ReferenceAsString(dcr.ref.Append("target_resource_id"))
}

// Ttl returns a reference to field ttl of azurerm_dns_cname_record.
func (dcr dnsCnameRecordAttributes) Ttl() terra.NumberValue {
	return terra.ReferenceAsNumber(dcr.ref.Append("ttl"))
}

// ZoneName returns a reference to field zone_name of azurerm_dns_cname_record.
func (dcr dnsCnameRecordAttributes) ZoneName() terra.StringValue {
	return terra.ReferenceAsString(dcr.ref.Append("zone_name"))
}

func (dcr dnsCnameRecordAttributes) Timeouts() dnscnamerecord.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dnscnamerecord.TimeoutsAttributes](dcr.ref.Append("timeouts"))
}

type dnsCnameRecordState struct {
	Fqdn              string                        `json:"fqdn"`
	Id                string                        `json:"id"`
	Name              string                        `json:"name"`
	Record            string                        `json:"record"`
	ResourceGroupName string                        `json:"resource_group_name"`
	Tags              map[string]string             `json:"tags"`
	TargetResourceId  string                        `json:"target_resource_id"`
	Ttl               float64                       `json:"ttl"`
	ZoneName          string                        `json:"zone_name"`
	Timeouts          *dnscnamerecord.TimeoutsState `json:"timeouts"`
}
