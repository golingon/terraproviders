// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	privatednsptrrecord "github.com/golingon/terraproviders/azurerm/3.98.0/privatednsptrrecord"
	"io"
)

// NewPrivateDnsPtrRecord creates a new instance of [PrivateDnsPtrRecord].
func NewPrivateDnsPtrRecord(name string, args PrivateDnsPtrRecordArgs) *PrivateDnsPtrRecord {
	return &PrivateDnsPtrRecord{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*PrivateDnsPtrRecord)(nil)

// PrivateDnsPtrRecord represents the Terraform resource azurerm_private_dns_ptr_record.
type PrivateDnsPtrRecord struct {
	Name      string
	Args      PrivateDnsPtrRecordArgs
	state     *privateDnsPtrRecordState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [PrivateDnsPtrRecord].
func (pdpr *PrivateDnsPtrRecord) Type() string {
	return "azurerm_private_dns_ptr_record"
}

// LocalName returns the local name for [PrivateDnsPtrRecord].
func (pdpr *PrivateDnsPtrRecord) LocalName() string {
	return pdpr.Name
}

// Configuration returns the configuration (args) for [PrivateDnsPtrRecord].
func (pdpr *PrivateDnsPtrRecord) Configuration() interface{} {
	return pdpr.Args
}

// DependOn is used for other resources to depend on [PrivateDnsPtrRecord].
func (pdpr *PrivateDnsPtrRecord) DependOn() terra.Reference {
	return terra.ReferenceResource(pdpr)
}

// Dependencies returns the list of resources [PrivateDnsPtrRecord] depends_on.
func (pdpr *PrivateDnsPtrRecord) Dependencies() terra.Dependencies {
	return pdpr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [PrivateDnsPtrRecord].
func (pdpr *PrivateDnsPtrRecord) LifecycleManagement() *terra.Lifecycle {
	return pdpr.Lifecycle
}

// Attributes returns the attributes for [PrivateDnsPtrRecord].
func (pdpr *PrivateDnsPtrRecord) Attributes() privateDnsPtrRecordAttributes {
	return privateDnsPtrRecordAttributes{ref: terra.ReferenceResource(pdpr)}
}

// ImportState imports the given attribute values into [PrivateDnsPtrRecord]'s state.
func (pdpr *PrivateDnsPtrRecord) ImportState(av io.Reader) error {
	pdpr.state = &privateDnsPtrRecordState{}
	if err := json.NewDecoder(av).Decode(pdpr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", pdpr.Type(), pdpr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [PrivateDnsPtrRecord] has state.
func (pdpr *PrivateDnsPtrRecord) State() (*privateDnsPtrRecordState, bool) {
	return pdpr.state, pdpr.state != nil
}

// StateMust returns the state for [PrivateDnsPtrRecord]. Panics if the state is nil.
func (pdpr *PrivateDnsPtrRecord) StateMust() *privateDnsPtrRecordState {
	if pdpr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", pdpr.Type(), pdpr.LocalName()))
	}
	return pdpr.state
}

// PrivateDnsPtrRecordArgs contains the configurations for azurerm_private_dns_ptr_record.
type PrivateDnsPtrRecordArgs struct {
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
	Timeouts *privatednsptrrecord.Timeouts `hcl:"timeouts,block"`
}
type privateDnsPtrRecordAttributes struct {
	ref terra.Reference
}

// Fqdn returns a reference to field fqdn of azurerm_private_dns_ptr_record.
func (pdpr privateDnsPtrRecordAttributes) Fqdn() terra.StringValue {
	return terra.ReferenceAsString(pdpr.ref.Append("fqdn"))
}

// Id returns a reference to field id of azurerm_private_dns_ptr_record.
func (pdpr privateDnsPtrRecordAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(pdpr.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_private_dns_ptr_record.
func (pdpr privateDnsPtrRecordAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(pdpr.ref.Append("name"))
}

// Records returns a reference to field records of azurerm_private_dns_ptr_record.
func (pdpr privateDnsPtrRecordAttributes) Records() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](pdpr.ref.Append("records"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_private_dns_ptr_record.
func (pdpr privateDnsPtrRecordAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(pdpr.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_private_dns_ptr_record.
func (pdpr privateDnsPtrRecordAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](pdpr.ref.Append("tags"))
}

// Ttl returns a reference to field ttl of azurerm_private_dns_ptr_record.
func (pdpr privateDnsPtrRecordAttributes) Ttl() terra.NumberValue {
	return terra.ReferenceAsNumber(pdpr.ref.Append("ttl"))
}

// ZoneName returns a reference to field zone_name of azurerm_private_dns_ptr_record.
func (pdpr privateDnsPtrRecordAttributes) ZoneName() terra.StringValue {
	return terra.ReferenceAsString(pdpr.ref.Append("zone_name"))
}

func (pdpr privateDnsPtrRecordAttributes) Timeouts() privatednsptrrecord.TimeoutsAttributes {
	return terra.ReferenceAsSingle[privatednsptrrecord.TimeoutsAttributes](pdpr.ref.Append("timeouts"))
}

type privateDnsPtrRecordState struct {
	Fqdn              string                             `json:"fqdn"`
	Id                string                             `json:"id"`
	Name              string                             `json:"name"`
	Records           []string                           `json:"records"`
	ResourceGroupName string                             `json:"resource_group_name"`
	Tags              map[string]string                  `json:"tags"`
	Ttl               float64                            `json:"ttl"`
	ZoneName          string                             `json:"zone_name"`
	Timeouts          *privatednsptrrecord.TimeoutsState `json:"timeouts"`
}
