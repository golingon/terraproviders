// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	privatednsarecord "github.com/golingon/terraproviders/azurerm/3.64.0/privatednsarecord"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewPrivateDnsARecord creates a new instance of [PrivateDnsARecord].
func NewPrivateDnsARecord(name string, args PrivateDnsARecordArgs) *PrivateDnsARecord {
	return &PrivateDnsARecord{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*PrivateDnsARecord)(nil)

// PrivateDnsARecord represents the Terraform resource azurerm_private_dns_a_record.
type PrivateDnsARecord struct {
	Name      string
	Args      PrivateDnsARecordArgs
	state     *privateDnsARecordState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [PrivateDnsARecord].
func (pdar *PrivateDnsARecord) Type() string {
	return "azurerm_private_dns_a_record"
}

// LocalName returns the local name for [PrivateDnsARecord].
func (pdar *PrivateDnsARecord) LocalName() string {
	return pdar.Name
}

// Configuration returns the configuration (args) for [PrivateDnsARecord].
func (pdar *PrivateDnsARecord) Configuration() interface{} {
	return pdar.Args
}

// DependOn is used for other resources to depend on [PrivateDnsARecord].
func (pdar *PrivateDnsARecord) DependOn() terra.Reference {
	return terra.ReferenceResource(pdar)
}

// Dependencies returns the list of resources [PrivateDnsARecord] depends_on.
func (pdar *PrivateDnsARecord) Dependencies() terra.Dependencies {
	return pdar.DependsOn
}

// LifecycleManagement returns the lifecycle block for [PrivateDnsARecord].
func (pdar *PrivateDnsARecord) LifecycleManagement() *terra.Lifecycle {
	return pdar.Lifecycle
}

// Attributes returns the attributes for [PrivateDnsARecord].
func (pdar *PrivateDnsARecord) Attributes() privateDnsARecordAttributes {
	return privateDnsARecordAttributes{ref: terra.ReferenceResource(pdar)}
}

// ImportState imports the given attribute values into [PrivateDnsARecord]'s state.
func (pdar *PrivateDnsARecord) ImportState(av io.Reader) error {
	pdar.state = &privateDnsARecordState{}
	if err := json.NewDecoder(av).Decode(pdar.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", pdar.Type(), pdar.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [PrivateDnsARecord] has state.
func (pdar *PrivateDnsARecord) State() (*privateDnsARecordState, bool) {
	return pdar.state, pdar.state != nil
}

// StateMust returns the state for [PrivateDnsARecord]. Panics if the state is nil.
func (pdar *PrivateDnsARecord) StateMust() *privateDnsARecordState {
	if pdar.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", pdar.Type(), pdar.LocalName()))
	}
	return pdar.state
}

// PrivateDnsARecordArgs contains the configurations for azurerm_private_dns_a_record.
type PrivateDnsARecordArgs struct {
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
	Timeouts *privatednsarecord.Timeouts `hcl:"timeouts,block"`
}
type privateDnsARecordAttributes struct {
	ref terra.Reference
}

// Fqdn returns a reference to field fqdn of azurerm_private_dns_a_record.
func (pdar privateDnsARecordAttributes) Fqdn() terra.StringValue {
	return terra.ReferenceAsString(pdar.ref.Append("fqdn"))
}

// Id returns a reference to field id of azurerm_private_dns_a_record.
func (pdar privateDnsARecordAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(pdar.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_private_dns_a_record.
func (pdar privateDnsARecordAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(pdar.ref.Append("name"))
}

// Records returns a reference to field records of azurerm_private_dns_a_record.
func (pdar privateDnsARecordAttributes) Records() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](pdar.ref.Append("records"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_private_dns_a_record.
func (pdar privateDnsARecordAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(pdar.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_private_dns_a_record.
func (pdar privateDnsARecordAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](pdar.ref.Append("tags"))
}

// Ttl returns a reference to field ttl of azurerm_private_dns_a_record.
func (pdar privateDnsARecordAttributes) Ttl() terra.NumberValue {
	return terra.ReferenceAsNumber(pdar.ref.Append("ttl"))
}

// ZoneName returns a reference to field zone_name of azurerm_private_dns_a_record.
func (pdar privateDnsARecordAttributes) ZoneName() terra.StringValue {
	return terra.ReferenceAsString(pdar.ref.Append("zone_name"))
}

func (pdar privateDnsARecordAttributes) Timeouts() privatednsarecord.TimeoutsAttributes {
	return terra.ReferenceAsSingle[privatednsarecord.TimeoutsAttributes](pdar.ref.Append("timeouts"))
}

type privateDnsARecordState struct {
	Fqdn              string                           `json:"fqdn"`
	Id                string                           `json:"id"`
	Name              string                           `json:"name"`
	Records           []string                         `json:"records"`
	ResourceGroupName string                           `json:"resource_group_name"`
	Tags              map[string]string                `json:"tags"`
	Ttl               float64                          `json:"ttl"`
	ZoneName          string                           `json:"zone_name"`
	Timeouts          *privatednsarecord.TimeoutsState `json:"timeouts"`
}
