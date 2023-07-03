// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	privatednsaaaarecord "github.com/golingon/terraproviders/azurerm/3.63.0/privatednsaaaarecord"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewPrivateDnsAaaaRecord creates a new instance of [PrivateDnsAaaaRecord].
func NewPrivateDnsAaaaRecord(name string, args PrivateDnsAaaaRecordArgs) *PrivateDnsAaaaRecord {
	return &PrivateDnsAaaaRecord{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*PrivateDnsAaaaRecord)(nil)

// PrivateDnsAaaaRecord represents the Terraform resource azurerm_private_dns_aaaa_record.
type PrivateDnsAaaaRecord struct {
	Name      string
	Args      PrivateDnsAaaaRecordArgs
	state     *privateDnsAaaaRecordState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [PrivateDnsAaaaRecord].
func (pdar *PrivateDnsAaaaRecord) Type() string {
	return "azurerm_private_dns_aaaa_record"
}

// LocalName returns the local name for [PrivateDnsAaaaRecord].
func (pdar *PrivateDnsAaaaRecord) LocalName() string {
	return pdar.Name
}

// Configuration returns the configuration (args) for [PrivateDnsAaaaRecord].
func (pdar *PrivateDnsAaaaRecord) Configuration() interface{} {
	return pdar.Args
}

// DependOn is used for other resources to depend on [PrivateDnsAaaaRecord].
func (pdar *PrivateDnsAaaaRecord) DependOn() terra.Reference {
	return terra.ReferenceResource(pdar)
}

// Dependencies returns the list of resources [PrivateDnsAaaaRecord] depends_on.
func (pdar *PrivateDnsAaaaRecord) Dependencies() terra.Dependencies {
	return pdar.DependsOn
}

// LifecycleManagement returns the lifecycle block for [PrivateDnsAaaaRecord].
func (pdar *PrivateDnsAaaaRecord) LifecycleManagement() *terra.Lifecycle {
	return pdar.Lifecycle
}

// Attributes returns the attributes for [PrivateDnsAaaaRecord].
func (pdar *PrivateDnsAaaaRecord) Attributes() privateDnsAaaaRecordAttributes {
	return privateDnsAaaaRecordAttributes{ref: terra.ReferenceResource(pdar)}
}

// ImportState imports the given attribute values into [PrivateDnsAaaaRecord]'s state.
func (pdar *PrivateDnsAaaaRecord) ImportState(av io.Reader) error {
	pdar.state = &privateDnsAaaaRecordState{}
	if err := json.NewDecoder(av).Decode(pdar.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", pdar.Type(), pdar.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [PrivateDnsAaaaRecord] has state.
func (pdar *PrivateDnsAaaaRecord) State() (*privateDnsAaaaRecordState, bool) {
	return pdar.state, pdar.state != nil
}

// StateMust returns the state for [PrivateDnsAaaaRecord]. Panics if the state is nil.
func (pdar *PrivateDnsAaaaRecord) StateMust() *privateDnsAaaaRecordState {
	if pdar.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", pdar.Type(), pdar.LocalName()))
	}
	return pdar.state
}

// PrivateDnsAaaaRecordArgs contains the configurations for azurerm_private_dns_aaaa_record.
type PrivateDnsAaaaRecordArgs struct {
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
	Timeouts *privatednsaaaarecord.Timeouts `hcl:"timeouts,block"`
}
type privateDnsAaaaRecordAttributes struct {
	ref terra.Reference
}

// Fqdn returns a reference to field fqdn of azurerm_private_dns_aaaa_record.
func (pdar privateDnsAaaaRecordAttributes) Fqdn() terra.StringValue {
	return terra.ReferenceAsString(pdar.ref.Append("fqdn"))
}

// Id returns a reference to field id of azurerm_private_dns_aaaa_record.
func (pdar privateDnsAaaaRecordAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(pdar.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_private_dns_aaaa_record.
func (pdar privateDnsAaaaRecordAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(pdar.ref.Append("name"))
}

// Records returns a reference to field records of azurerm_private_dns_aaaa_record.
func (pdar privateDnsAaaaRecordAttributes) Records() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](pdar.ref.Append("records"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_private_dns_aaaa_record.
func (pdar privateDnsAaaaRecordAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(pdar.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_private_dns_aaaa_record.
func (pdar privateDnsAaaaRecordAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](pdar.ref.Append("tags"))
}

// Ttl returns a reference to field ttl of azurerm_private_dns_aaaa_record.
func (pdar privateDnsAaaaRecordAttributes) Ttl() terra.NumberValue {
	return terra.ReferenceAsNumber(pdar.ref.Append("ttl"))
}

// ZoneName returns a reference to field zone_name of azurerm_private_dns_aaaa_record.
func (pdar privateDnsAaaaRecordAttributes) ZoneName() terra.StringValue {
	return terra.ReferenceAsString(pdar.ref.Append("zone_name"))
}

func (pdar privateDnsAaaaRecordAttributes) Timeouts() privatednsaaaarecord.TimeoutsAttributes {
	return terra.ReferenceAsSingle[privatednsaaaarecord.TimeoutsAttributes](pdar.ref.Append("timeouts"))
}

type privateDnsAaaaRecordState struct {
	Fqdn              string                              `json:"fqdn"`
	Id                string                              `json:"id"`
	Name              string                              `json:"name"`
	Records           []string                            `json:"records"`
	ResourceGroupName string                              `json:"resource_group_name"`
	Tags              map[string]string                   `json:"tags"`
	Ttl               float64                             `json:"ttl"`
	ZoneName          string                              `json:"zone_name"`
	Timeouts          *privatednsaaaarecord.TimeoutsState `json:"timeouts"`
}
