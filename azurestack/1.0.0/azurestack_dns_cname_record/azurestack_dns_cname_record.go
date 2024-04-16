// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurestack_dns_cname_record

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// New creates a new instance of [Resource].
func New(name string, args Args) *Resource {
	return &Resource{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Resource)(nil)

// Resource represents the Terraform resource azurestack_dns_cname_record.
type Resource struct {
	Name      string
	Args      Args
	state     *azurestackDnsCnameRecordState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (adcr *Resource) Type() string {
	return "azurestack_dns_cname_record"
}

// LocalName returns the local name for [Resource].
func (adcr *Resource) LocalName() string {
	return adcr.Name
}

// Configuration returns the configuration (args) for [Resource].
func (adcr *Resource) Configuration() interface{} {
	return adcr.Args
}

// DependOn is used for other resources to depend on [Resource].
func (adcr *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(adcr)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (adcr *Resource) Dependencies() terra.Dependencies {
	return adcr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (adcr *Resource) LifecycleManagement() *terra.Lifecycle {
	return adcr.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (adcr *Resource) Attributes() azurestackDnsCnameRecordAttributes {
	return azurestackDnsCnameRecordAttributes{ref: terra.ReferenceResource(adcr)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (adcr *Resource) ImportState(state io.Reader) error {
	adcr.state = &azurestackDnsCnameRecordState{}
	if err := json.NewDecoder(state).Decode(adcr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", adcr.Type(), adcr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (adcr *Resource) State() (*azurestackDnsCnameRecordState, bool) {
	return adcr.state, adcr.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (adcr *Resource) StateMust() *azurestackDnsCnameRecordState {
	if adcr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", adcr.Type(), adcr.LocalName()))
	}
	return adcr.state
}

// Args contains the configurations for azurestack_dns_cname_record.
type Args struct {
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
	// Ttl: number, required
	Ttl terra.NumberValue `hcl:"ttl,attr" validate:"required"`
	// ZoneName: string, required
	ZoneName terra.StringValue `hcl:"zone_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type azurestackDnsCnameRecordAttributes struct {
	ref terra.Reference
}

// Fqdn returns a reference to field fqdn of azurestack_dns_cname_record.
func (adcr azurestackDnsCnameRecordAttributes) Fqdn() terra.StringValue {
	return terra.ReferenceAsString(adcr.ref.Append("fqdn"))
}

// Id returns a reference to field id of azurestack_dns_cname_record.
func (adcr azurestackDnsCnameRecordAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(adcr.ref.Append("id"))
}

// Name returns a reference to field name of azurestack_dns_cname_record.
func (adcr azurestackDnsCnameRecordAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(adcr.ref.Append("name"))
}

// Record returns a reference to field record of azurestack_dns_cname_record.
func (adcr azurestackDnsCnameRecordAttributes) Record() terra.StringValue {
	return terra.ReferenceAsString(adcr.ref.Append("record"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurestack_dns_cname_record.
func (adcr azurestackDnsCnameRecordAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(adcr.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurestack_dns_cname_record.
func (adcr azurestackDnsCnameRecordAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](adcr.ref.Append("tags"))
}

// Ttl returns a reference to field ttl of azurestack_dns_cname_record.
func (adcr azurestackDnsCnameRecordAttributes) Ttl() terra.NumberValue {
	return terra.ReferenceAsNumber(adcr.ref.Append("ttl"))
}

// ZoneName returns a reference to field zone_name of azurestack_dns_cname_record.
func (adcr azurestackDnsCnameRecordAttributes) ZoneName() terra.StringValue {
	return terra.ReferenceAsString(adcr.ref.Append("zone_name"))
}

func (adcr azurestackDnsCnameRecordAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](adcr.ref.Append("timeouts"))
}

type azurestackDnsCnameRecordState struct {
	Fqdn              string            `json:"fqdn"`
	Id                string            `json:"id"`
	Name              string            `json:"name"`
	Record            string            `json:"record"`
	ResourceGroupName string            `json:"resource_group_name"`
	Tags              map[string]string `json:"tags"`
	Ttl               float64           `json:"ttl"`
	ZoneName          string            `json:"zone_name"`
	Timeouts          *TimeoutsState    `json:"timeouts"`
}
