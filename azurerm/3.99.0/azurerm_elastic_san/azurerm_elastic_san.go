// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_elastic_san

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

// Resource represents the Terraform resource azurerm_elastic_san.
type Resource struct {
	Name      string
	Args      Args
	state     *azurermElasticSanState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (aes *Resource) Type() string {
	return "azurerm_elastic_san"
}

// LocalName returns the local name for [Resource].
func (aes *Resource) LocalName() string {
	return aes.Name
}

// Configuration returns the configuration (args) for [Resource].
func (aes *Resource) Configuration() interface{} {
	return aes.Args
}

// DependOn is used for other resources to depend on [Resource].
func (aes *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(aes)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (aes *Resource) Dependencies() terra.Dependencies {
	return aes.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (aes *Resource) LifecycleManagement() *terra.Lifecycle {
	return aes.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (aes *Resource) Attributes() azurermElasticSanAttributes {
	return azurermElasticSanAttributes{ref: terra.ReferenceResource(aes)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (aes *Resource) ImportState(state io.Reader) error {
	aes.state = &azurermElasticSanState{}
	if err := json.NewDecoder(state).Decode(aes.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", aes.Type(), aes.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (aes *Resource) State() (*azurermElasticSanState, bool) {
	return aes.state, aes.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (aes *Resource) StateMust() *azurermElasticSanState {
	if aes.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", aes.Type(), aes.LocalName()))
	}
	return aes.state
}

// Args contains the configurations for azurerm_elastic_san.
type Args struct {
	// BaseSizeInTib: number, required
	BaseSizeInTib terra.NumberValue `hcl:"base_size_in_tib,attr" validate:"required"`
	// ExtendedSizeInTib: number, optional
	ExtendedSizeInTib terra.NumberValue `hcl:"extended_size_in_tib,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Zones: set of string, optional
	Zones terra.SetValue[terra.StringValue] `hcl:"zones,attr"`
	// Sku: required
	Sku *Sku `hcl:"sku,block" validate:"required"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type azurermElasticSanAttributes struct {
	ref terra.Reference
}

// BaseSizeInTib returns a reference to field base_size_in_tib of azurerm_elastic_san.
func (aes azurermElasticSanAttributes) BaseSizeInTib() terra.NumberValue {
	return terra.ReferenceAsNumber(aes.ref.Append("base_size_in_tib"))
}

// ExtendedSizeInTib returns a reference to field extended_size_in_tib of azurerm_elastic_san.
func (aes azurermElasticSanAttributes) ExtendedSizeInTib() terra.NumberValue {
	return terra.ReferenceAsNumber(aes.ref.Append("extended_size_in_tib"))
}

// Id returns a reference to field id of azurerm_elastic_san.
func (aes azurermElasticSanAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aes.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_elastic_san.
func (aes azurermElasticSanAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(aes.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_elastic_san.
func (aes azurermElasticSanAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(aes.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_elastic_san.
func (aes azurermElasticSanAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(aes.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_elastic_san.
func (aes azurermElasticSanAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](aes.ref.Append("tags"))
}

// TotalIops returns a reference to field total_iops of azurerm_elastic_san.
func (aes azurermElasticSanAttributes) TotalIops() terra.NumberValue {
	return terra.ReferenceAsNumber(aes.ref.Append("total_iops"))
}

// TotalMbps returns a reference to field total_mbps of azurerm_elastic_san.
func (aes azurermElasticSanAttributes) TotalMbps() terra.NumberValue {
	return terra.ReferenceAsNumber(aes.ref.Append("total_mbps"))
}

// TotalSizeInTib returns a reference to field total_size_in_tib of azurerm_elastic_san.
func (aes azurermElasticSanAttributes) TotalSizeInTib() terra.NumberValue {
	return terra.ReferenceAsNumber(aes.ref.Append("total_size_in_tib"))
}

// TotalVolumeSizeInGib returns a reference to field total_volume_size_in_gib of azurerm_elastic_san.
func (aes azurermElasticSanAttributes) TotalVolumeSizeInGib() terra.NumberValue {
	return terra.ReferenceAsNumber(aes.ref.Append("total_volume_size_in_gib"))
}

// VolumeGroupCount returns a reference to field volume_group_count of azurerm_elastic_san.
func (aes azurermElasticSanAttributes) VolumeGroupCount() terra.NumberValue {
	return terra.ReferenceAsNumber(aes.ref.Append("volume_group_count"))
}

// Zones returns a reference to field zones of azurerm_elastic_san.
func (aes azurermElasticSanAttributes) Zones() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](aes.ref.Append("zones"))
}

func (aes azurermElasticSanAttributes) Sku() terra.ListValue[SkuAttributes] {
	return terra.ReferenceAsList[SkuAttributes](aes.ref.Append("sku"))
}

func (aes azurermElasticSanAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](aes.ref.Append("timeouts"))
}

type azurermElasticSanState struct {
	BaseSizeInTib        float64           `json:"base_size_in_tib"`
	ExtendedSizeInTib    float64           `json:"extended_size_in_tib"`
	Id                   string            `json:"id"`
	Location             string            `json:"location"`
	Name                 string            `json:"name"`
	ResourceGroupName    string            `json:"resource_group_name"`
	Tags                 map[string]string `json:"tags"`
	TotalIops            float64           `json:"total_iops"`
	TotalMbps            float64           `json:"total_mbps"`
	TotalSizeInTib       float64           `json:"total_size_in_tib"`
	TotalVolumeSizeInGib float64           `json:"total_volume_size_in_gib"`
	VolumeGroupCount     float64           `json:"volume_group_count"`
	Zones                []string          `json:"zones"`
	Sku                  []SkuState        `json:"sku"`
	Timeouts             *TimeoutsState    `json:"timeouts"`
}
