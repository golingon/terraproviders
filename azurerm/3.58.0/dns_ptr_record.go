// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	dnsptrrecord "github.com/golingon/terraproviders/azurerm/3.58.0/dnsptrrecord"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDnsPtrRecord creates a new instance of [DnsPtrRecord].
func NewDnsPtrRecord(name string, args DnsPtrRecordArgs) *DnsPtrRecord {
	return &DnsPtrRecord{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DnsPtrRecord)(nil)

// DnsPtrRecord represents the Terraform resource azurerm_dns_ptr_record.
type DnsPtrRecord struct {
	Name      string
	Args      DnsPtrRecordArgs
	state     *dnsPtrRecordState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DnsPtrRecord].
func (dpr *DnsPtrRecord) Type() string {
	return "azurerm_dns_ptr_record"
}

// LocalName returns the local name for [DnsPtrRecord].
func (dpr *DnsPtrRecord) LocalName() string {
	return dpr.Name
}

// Configuration returns the configuration (args) for [DnsPtrRecord].
func (dpr *DnsPtrRecord) Configuration() interface{} {
	return dpr.Args
}

// DependOn is used for other resources to depend on [DnsPtrRecord].
func (dpr *DnsPtrRecord) DependOn() terra.Reference {
	return terra.ReferenceResource(dpr)
}

// Dependencies returns the list of resources [DnsPtrRecord] depends_on.
func (dpr *DnsPtrRecord) Dependencies() terra.Dependencies {
	return dpr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DnsPtrRecord].
func (dpr *DnsPtrRecord) LifecycleManagement() *terra.Lifecycle {
	return dpr.Lifecycle
}

// Attributes returns the attributes for [DnsPtrRecord].
func (dpr *DnsPtrRecord) Attributes() dnsPtrRecordAttributes {
	return dnsPtrRecordAttributes{ref: terra.ReferenceResource(dpr)}
}

// ImportState imports the given attribute values into [DnsPtrRecord]'s state.
func (dpr *DnsPtrRecord) ImportState(av io.Reader) error {
	dpr.state = &dnsPtrRecordState{}
	if err := json.NewDecoder(av).Decode(dpr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dpr.Type(), dpr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DnsPtrRecord] has state.
func (dpr *DnsPtrRecord) State() (*dnsPtrRecordState, bool) {
	return dpr.state, dpr.state != nil
}

// StateMust returns the state for [DnsPtrRecord]. Panics if the state is nil.
func (dpr *DnsPtrRecord) StateMust() *dnsPtrRecordState {
	if dpr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dpr.Type(), dpr.LocalName()))
	}
	return dpr.state
}

// DnsPtrRecordArgs contains the configurations for azurerm_dns_ptr_record.
type DnsPtrRecordArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Records: set of string, required
	Records terra.SetValue[terra.StringValue] `hcl:"records,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Ttl: number, required
	Ttl terra.NumberValue `hcl:"ttl,attr" validate:"required"`
	// ZoneName: string, required
	ZoneName terra.StringValue `hcl:"zone_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *dnsptrrecord.Timeouts `hcl:"timeouts,block"`
}
type dnsPtrRecordAttributes struct {
	ref terra.Reference
}

// Fqdn returns a reference to field fqdn of azurerm_dns_ptr_record.
func (dpr dnsPtrRecordAttributes) Fqdn() terra.StringValue {
	return terra.ReferenceAsString(dpr.ref.Append("fqdn"))
}

// Id returns a reference to field id of azurerm_dns_ptr_record.
func (dpr dnsPtrRecordAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dpr.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_dns_ptr_record.
func (dpr dnsPtrRecordAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dpr.ref.Append("name"))
}

// Records returns a reference to field records of azurerm_dns_ptr_record.
func (dpr dnsPtrRecordAttributes) Records() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](dpr.ref.Append("records"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_dns_ptr_record.
func (dpr dnsPtrRecordAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(dpr.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_dns_ptr_record.
func (dpr dnsPtrRecordAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dpr.ref.Append("tags"))
}

// Ttl returns a reference to field ttl of azurerm_dns_ptr_record.
func (dpr dnsPtrRecordAttributes) Ttl() terra.NumberValue {
	return terra.ReferenceAsNumber(dpr.ref.Append("ttl"))
}

// ZoneName returns a reference to field zone_name of azurerm_dns_ptr_record.
func (dpr dnsPtrRecordAttributes) ZoneName() terra.StringValue {
	return terra.ReferenceAsString(dpr.ref.Append("zone_name"))
}

func (dpr dnsPtrRecordAttributes) Timeouts() dnsptrrecord.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dnsptrrecord.TimeoutsAttributes](dpr.ref.Append("timeouts"))
}

type dnsPtrRecordState struct {
	Fqdn              string                      `json:"fqdn"`
	Id                string                      `json:"id"`
	Name              string                      `json:"name"`
	Records           []string                    `json:"records"`
	ResourceGroupName string                      `json:"resource_group_name"`
	Tags              map[string]string           `json:"tags"`
	Ttl               float64                     `json:"ttl"`
	ZoneName          string                      `json:"zone_name"`
	Timeouts          *dnsptrrecord.TimeoutsState `json:"timeouts"`
}
