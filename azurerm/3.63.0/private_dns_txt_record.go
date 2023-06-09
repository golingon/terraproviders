// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	privatednstxtrecord "github.com/golingon/terraproviders/azurerm/3.63.0/privatednstxtrecord"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewPrivateDnsTxtRecord creates a new instance of [PrivateDnsTxtRecord].
func NewPrivateDnsTxtRecord(name string, args PrivateDnsTxtRecordArgs) *PrivateDnsTxtRecord {
	return &PrivateDnsTxtRecord{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*PrivateDnsTxtRecord)(nil)

// PrivateDnsTxtRecord represents the Terraform resource azurerm_private_dns_txt_record.
type PrivateDnsTxtRecord struct {
	Name      string
	Args      PrivateDnsTxtRecordArgs
	state     *privateDnsTxtRecordState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [PrivateDnsTxtRecord].
func (pdtr *PrivateDnsTxtRecord) Type() string {
	return "azurerm_private_dns_txt_record"
}

// LocalName returns the local name for [PrivateDnsTxtRecord].
func (pdtr *PrivateDnsTxtRecord) LocalName() string {
	return pdtr.Name
}

// Configuration returns the configuration (args) for [PrivateDnsTxtRecord].
func (pdtr *PrivateDnsTxtRecord) Configuration() interface{} {
	return pdtr.Args
}

// DependOn is used for other resources to depend on [PrivateDnsTxtRecord].
func (pdtr *PrivateDnsTxtRecord) DependOn() terra.Reference {
	return terra.ReferenceResource(pdtr)
}

// Dependencies returns the list of resources [PrivateDnsTxtRecord] depends_on.
func (pdtr *PrivateDnsTxtRecord) Dependencies() terra.Dependencies {
	return pdtr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [PrivateDnsTxtRecord].
func (pdtr *PrivateDnsTxtRecord) LifecycleManagement() *terra.Lifecycle {
	return pdtr.Lifecycle
}

// Attributes returns the attributes for [PrivateDnsTxtRecord].
func (pdtr *PrivateDnsTxtRecord) Attributes() privateDnsTxtRecordAttributes {
	return privateDnsTxtRecordAttributes{ref: terra.ReferenceResource(pdtr)}
}

// ImportState imports the given attribute values into [PrivateDnsTxtRecord]'s state.
func (pdtr *PrivateDnsTxtRecord) ImportState(av io.Reader) error {
	pdtr.state = &privateDnsTxtRecordState{}
	if err := json.NewDecoder(av).Decode(pdtr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", pdtr.Type(), pdtr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [PrivateDnsTxtRecord] has state.
func (pdtr *PrivateDnsTxtRecord) State() (*privateDnsTxtRecordState, bool) {
	return pdtr.state, pdtr.state != nil
}

// StateMust returns the state for [PrivateDnsTxtRecord]. Panics if the state is nil.
func (pdtr *PrivateDnsTxtRecord) StateMust() *privateDnsTxtRecordState {
	if pdtr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", pdtr.Type(), pdtr.LocalName()))
	}
	return pdtr.state
}

// PrivateDnsTxtRecordArgs contains the configurations for azurerm_private_dns_txt_record.
type PrivateDnsTxtRecordArgs struct {
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
	Record []privatednstxtrecord.Record `hcl:"record,block" validate:"min=1"`
	// Timeouts: optional
	Timeouts *privatednstxtrecord.Timeouts `hcl:"timeouts,block"`
}
type privateDnsTxtRecordAttributes struct {
	ref terra.Reference
}

// Fqdn returns a reference to field fqdn of azurerm_private_dns_txt_record.
func (pdtr privateDnsTxtRecordAttributes) Fqdn() terra.StringValue {
	return terra.ReferenceAsString(pdtr.ref.Append("fqdn"))
}

// Id returns a reference to field id of azurerm_private_dns_txt_record.
func (pdtr privateDnsTxtRecordAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(pdtr.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_private_dns_txt_record.
func (pdtr privateDnsTxtRecordAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(pdtr.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_private_dns_txt_record.
func (pdtr privateDnsTxtRecordAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(pdtr.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_private_dns_txt_record.
func (pdtr privateDnsTxtRecordAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](pdtr.ref.Append("tags"))
}

// Ttl returns a reference to field ttl of azurerm_private_dns_txt_record.
func (pdtr privateDnsTxtRecordAttributes) Ttl() terra.NumberValue {
	return terra.ReferenceAsNumber(pdtr.ref.Append("ttl"))
}

// ZoneName returns a reference to field zone_name of azurerm_private_dns_txt_record.
func (pdtr privateDnsTxtRecordAttributes) ZoneName() terra.StringValue {
	return terra.ReferenceAsString(pdtr.ref.Append("zone_name"))
}

func (pdtr privateDnsTxtRecordAttributes) Record() terra.SetValue[privatednstxtrecord.RecordAttributes] {
	return terra.ReferenceAsSet[privatednstxtrecord.RecordAttributes](pdtr.ref.Append("record"))
}

func (pdtr privateDnsTxtRecordAttributes) Timeouts() privatednstxtrecord.TimeoutsAttributes {
	return terra.ReferenceAsSingle[privatednstxtrecord.TimeoutsAttributes](pdtr.ref.Append("timeouts"))
}

type privateDnsTxtRecordState struct {
	Fqdn              string                             `json:"fqdn"`
	Id                string                             `json:"id"`
	Name              string                             `json:"name"`
	ResourceGroupName string                             `json:"resource_group_name"`
	Tags              map[string]string                  `json:"tags"`
	Ttl               float64                            `json:"ttl"`
	ZoneName          string                             `json:"zone_name"`
	Record            []privatednstxtrecord.RecordState  `json:"record"`
	Timeouts          *privatednstxtrecord.TimeoutsState `json:"timeouts"`
}
