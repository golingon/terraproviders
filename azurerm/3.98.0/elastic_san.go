// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	elasticsan "github.com/golingon/terraproviders/azurerm/3.98.0/elasticsan"
	"io"
)

// NewElasticSan creates a new instance of [ElasticSan].
func NewElasticSan(name string, args ElasticSanArgs) *ElasticSan {
	return &ElasticSan{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ElasticSan)(nil)

// ElasticSan represents the Terraform resource azurerm_elastic_san.
type ElasticSan struct {
	Name      string
	Args      ElasticSanArgs
	state     *elasticSanState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ElasticSan].
func (es *ElasticSan) Type() string {
	return "azurerm_elastic_san"
}

// LocalName returns the local name for [ElasticSan].
func (es *ElasticSan) LocalName() string {
	return es.Name
}

// Configuration returns the configuration (args) for [ElasticSan].
func (es *ElasticSan) Configuration() interface{} {
	return es.Args
}

// DependOn is used for other resources to depend on [ElasticSan].
func (es *ElasticSan) DependOn() terra.Reference {
	return terra.ReferenceResource(es)
}

// Dependencies returns the list of resources [ElasticSan] depends_on.
func (es *ElasticSan) Dependencies() terra.Dependencies {
	return es.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ElasticSan].
func (es *ElasticSan) LifecycleManagement() *terra.Lifecycle {
	return es.Lifecycle
}

// Attributes returns the attributes for [ElasticSan].
func (es *ElasticSan) Attributes() elasticSanAttributes {
	return elasticSanAttributes{ref: terra.ReferenceResource(es)}
}

// ImportState imports the given attribute values into [ElasticSan]'s state.
func (es *ElasticSan) ImportState(av io.Reader) error {
	es.state = &elasticSanState{}
	if err := json.NewDecoder(av).Decode(es.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", es.Type(), es.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ElasticSan] has state.
func (es *ElasticSan) State() (*elasticSanState, bool) {
	return es.state, es.state != nil
}

// StateMust returns the state for [ElasticSan]. Panics if the state is nil.
func (es *ElasticSan) StateMust() *elasticSanState {
	if es.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", es.Type(), es.LocalName()))
	}
	return es.state
}

// ElasticSanArgs contains the configurations for azurerm_elastic_san.
type ElasticSanArgs struct {
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
	Sku *elasticsan.Sku `hcl:"sku,block" validate:"required"`
	// Timeouts: optional
	Timeouts *elasticsan.Timeouts `hcl:"timeouts,block"`
}
type elasticSanAttributes struct {
	ref terra.Reference
}

// BaseSizeInTib returns a reference to field base_size_in_tib of azurerm_elastic_san.
func (es elasticSanAttributes) BaseSizeInTib() terra.NumberValue {
	return terra.ReferenceAsNumber(es.ref.Append("base_size_in_tib"))
}

// ExtendedSizeInTib returns a reference to field extended_size_in_tib of azurerm_elastic_san.
func (es elasticSanAttributes) ExtendedSizeInTib() terra.NumberValue {
	return terra.ReferenceAsNumber(es.ref.Append("extended_size_in_tib"))
}

// Id returns a reference to field id of azurerm_elastic_san.
func (es elasticSanAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(es.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_elastic_san.
func (es elasticSanAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(es.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_elastic_san.
func (es elasticSanAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(es.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_elastic_san.
func (es elasticSanAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(es.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_elastic_san.
func (es elasticSanAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](es.ref.Append("tags"))
}

// TotalIops returns a reference to field total_iops of azurerm_elastic_san.
func (es elasticSanAttributes) TotalIops() terra.NumberValue {
	return terra.ReferenceAsNumber(es.ref.Append("total_iops"))
}

// TotalMbps returns a reference to field total_mbps of azurerm_elastic_san.
func (es elasticSanAttributes) TotalMbps() terra.NumberValue {
	return terra.ReferenceAsNumber(es.ref.Append("total_mbps"))
}

// TotalSizeInTib returns a reference to field total_size_in_tib of azurerm_elastic_san.
func (es elasticSanAttributes) TotalSizeInTib() terra.NumberValue {
	return terra.ReferenceAsNumber(es.ref.Append("total_size_in_tib"))
}

// TotalVolumeSizeInGib returns a reference to field total_volume_size_in_gib of azurerm_elastic_san.
func (es elasticSanAttributes) TotalVolumeSizeInGib() terra.NumberValue {
	return terra.ReferenceAsNumber(es.ref.Append("total_volume_size_in_gib"))
}

// VolumeGroupCount returns a reference to field volume_group_count of azurerm_elastic_san.
func (es elasticSanAttributes) VolumeGroupCount() terra.NumberValue {
	return terra.ReferenceAsNumber(es.ref.Append("volume_group_count"))
}

// Zones returns a reference to field zones of azurerm_elastic_san.
func (es elasticSanAttributes) Zones() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](es.ref.Append("zones"))
}

func (es elasticSanAttributes) Sku() terra.ListValue[elasticsan.SkuAttributes] {
	return terra.ReferenceAsList[elasticsan.SkuAttributes](es.ref.Append("sku"))
}

func (es elasticSanAttributes) Timeouts() elasticsan.TimeoutsAttributes {
	return terra.ReferenceAsSingle[elasticsan.TimeoutsAttributes](es.ref.Append("timeouts"))
}

type elasticSanState struct {
	BaseSizeInTib        float64                   `json:"base_size_in_tib"`
	ExtendedSizeInTib    float64                   `json:"extended_size_in_tib"`
	Id                   string                    `json:"id"`
	Location             string                    `json:"location"`
	Name                 string                    `json:"name"`
	ResourceGroupName    string                    `json:"resource_group_name"`
	Tags                 map[string]string         `json:"tags"`
	TotalIops            float64                   `json:"total_iops"`
	TotalMbps            float64                   `json:"total_mbps"`
	TotalSizeInTib       float64                   `json:"total_size_in_tib"`
	TotalVolumeSizeInGib float64                   `json:"total_volume_size_in_gib"`
	VolumeGroupCount     float64                   `json:"volume_group_count"`
	Zones                []string                  `json:"zones"`
	Sku                  []elasticsan.SkuState     `json:"sku"`
	Timeouts             *elasticsan.TimeoutsState `json:"timeouts"`
}
