// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	dnstxtrecord "github.com/golingon/terraproviders/azurerm/3.63.0/dnstxtrecord"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDnsTxtRecord creates a new instance of [DnsTxtRecord].
func NewDnsTxtRecord(name string, args DnsTxtRecordArgs) *DnsTxtRecord {
	return &DnsTxtRecord{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DnsTxtRecord)(nil)

// DnsTxtRecord represents the Terraform resource azurerm_dns_txt_record.
type DnsTxtRecord struct {
	Name      string
	Args      DnsTxtRecordArgs
	state     *dnsTxtRecordState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DnsTxtRecord].
func (dtr *DnsTxtRecord) Type() string {
	return "azurerm_dns_txt_record"
}

// LocalName returns the local name for [DnsTxtRecord].
func (dtr *DnsTxtRecord) LocalName() string {
	return dtr.Name
}

// Configuration returns the configuration (args) for [DnsTxtRecord].
func (dtr *DnsTxtRecord) Configuration() interface{} {
	return dtr.Args
}

// DependOn is used for other resources to depend on [DnsTxtRecord].
func (dtr *DnsTxtRecord) DependOn() terra.Reference {
	return terra.ReferenceResource(dtr)
}

// Dependencies returns the list of resources [DnsTxtRecord] depends_on.
func (dtr *DnsTxtRecord) Dependencies() terra.Dependencies {
	return dtr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DnsTxtRecord].
func (dtr *DnsTxtRecord) LifecycleManagement() *terra.Lifecycle {
	return dtr.Lifecycle
}

// Attributes returns the attributes for [DnsTxtRecord].
func (dtr *DnsTxtRecord) Attributes() dnsTxtRecordAttributes {
	return dnsTxtRecordAttributes{ref: terra.ReferenceResource(dtr)}
}

// ImportState imports the given attribute values into [DnsTxtRecord]'s state.
func (dtr *DnsTxtRecord) ImportState(av io.Reader) error {
	dtr.state = &dnsTxtRecordState{}
	if err := json.NewDecoder(av).Decode(dtr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dtr.Type(), dtr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DnsTxtRecord] has state.
func (dtr *DnsTxtRecord) State() (*dnsTxtRecordState, bool) {
	return dtr.state, dtr.state != nil
}

// StateMust returns the state for [DnsTxtRecord]. Panics if the state is nil.
func (dtr *DnsTxtRecord) StateMust() *dnsTxtRecordState {
	if dtr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dtr.Type(), dtr.LocalName()))
	}
	return dtr.state
}

// DnsTxtRecordArgs contains the configurations for azurerm_dns_txt_record.
type DnsTxtRecordArgs struct {
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
	Record []dnstxtrecord.Record `hcl:"record,block" validate:"min=1"`
	// Timeouts: optional
	Timeouts *dnstxtrecord.Timeouts `hcl:"timeouts,block"`
}
type dnsTxtRecordAttributes struct {
	ref terra.Reference
}

// Fqdn returns a reference to field fqdn of azurerm_dns_txt_record.
func (dtr dnsTxtRecordAttributes) Fqdn() terra.StringValue {
	return terra.ReferenceAsString(dtr.ref.Append("fqdn"))
}

// Id returns a reference to field id of azurerm_dns_txt_record.
func (dtr dnsTxtRecordAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dtr.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_dns_txt_record.
func (dtr dnsTxtRecordAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dtr.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_dns_txt_record.
func (dtr dnsTxtRecordAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(dtr.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_dns_txt_record.
func (dtr dnsTxtRecordAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dtr.ref.Append("tags"))
}

// Ttl returns a reference to field ttl of azurerm_dns_txt_record.
func (dtr dnsTxtRecordAttributes) Ttl() terra.NumberValue {
	return terra.ReferenceAsNumber(dtr.ref.Append("ttl"))
}

// ZoneName returns a reference to field zone_name of azurerm_dns_txt_record.
func (dtr dnsTxtRecordAttributes) ZoneName() terra.StringValue {
	return terra.ReferenceAsString(dtr.ref.Append("zone_name"))
}

func (dtr dnsTxtRecordAttributes) Record() terra.SetValue[dnstxtrecord.RecordAttributes] {
	return terra.ReferenceAsSet[dnstxtrecord.RecordAttributes](dtr.ref.Append("record"))
}

func (dtr dnsTxtRecordAttributes) Timeouts() dnstxtrecord.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dnstxtrecord.TimeoutsAttributes](dtr.ref.Append("timeouts"))
}

type dnsTxtRecordState struct {
	Fqdn              string                      `json:"fqdn"`
	Id                string                      `json:"id"`
	Name              string                      `json:"name"`
	ResourceGroupName string                      `json:"resource_group_name"`
	Tags              map[string]string           `json:"tags"`
	Ttl               float64                     `json:"ttl"`
	ZoneName          string                      `json:"zone_name"`
	Record            []dnstxtrecord.RecordState  `json:"record"`
	Timeouts          *dnstxtrecord.TimeoutsState `json:"timeouts"`
}
