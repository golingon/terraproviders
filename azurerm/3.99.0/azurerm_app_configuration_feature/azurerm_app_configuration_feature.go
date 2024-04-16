// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_app_configuration_feature

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

// Resource represents the Terraform resource azurerm_app_configuration_feature.
type Resource struct {
	Name      string
	Args      Args
	state     *azurermAppConfigurationFeatureState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (aacf *Resource) Type() string {
	return "azurerm_app_configuration_feature"
}

// LocalName returns the local name for [Resource].
func (aacf *Resource) LocalName() string {
	return aacf.Name
}

// Configuration returns the configuration (args) for [Resource].
func (aacf *Resource) Configuration() interface{} {
	return aacf.Args
}

// DependOn is used for other resources to depend on [Resource].
func (aacf *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(aacf)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (aacf *Resource) Dependencies() terra.Dependencies {
	return aacf.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (aacf *Resource) LifecycleManagement() *terra.Lifecycle {
	return aacf.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (aacf *Resource) Attributes() azurermAppConfigurationFeatureAttributes {
	return azurermAppConfigurationFeatureAttributes{ref: terra.ReferenceResource(aacf)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (aacf *Resource) ImportState(state io.Reader) error {
	aacf.state = &azurermAppConfigurationFeatureState{}
	if err := json.NewDecoder(state).Decode(aacf.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", aacf.Type(), aacf.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (aacf *Resource) State() (*azurermAppConfigurationFeatureState, bool) {
	return aacf.state, aacf.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (aacf *Resource) StateMust() *azurermAppConfigurationFeatureState {
	if aacf.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", aacf.Type(), aacf.LocalName()))
	}
	return aacf.state
}

// Args contains the configurations for azurerm_app_configuration_feature.
type Args struct {
	// ConfigurationStoreId: string, required
	ConfigurationStoreId terra.StringValue `hcl:"configuration_store_id,attr" validate:"required"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
	// Etag: string, optional
	Etag terra.StringValue `hcl:"etag,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Key: string, optional
	Key terra.StringValue `hcl:"key,attr"`
	// Label: string, optional
	Label terra.StringValue `hcl:"label,attr"`
	// Locked: bool, optional
	Locked terra.BoolValue `hcl:"locked,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PercentageFilterValue: number, optional
	PercentageFilterValue terra.NumberValue `hcl:"percentage_filter_value,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TargetingFilter: min=0
	TargetingFilter []TargetingFilter `hcl:"targeting_filter,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
	// TimewindowFilter: min=0
	TimewindowFilter []TimewindowFilter `hcl:"timewindow_filter,block" validate:"min=0"`
}

type azurermAppConfigurationFeatureAttributes struct {
	ref terra.Reference
}

// ConfigurationStoreId returns a reference to field configuration_store_id of azurerm_app_configuration_feature.
func (aacf azurermAppConfigurationFeatureAttributes) ConfigurationStoreId() terra.StringValue {
	return terra.ReferenceAsString(aacf.ref.Append("configuration_store_id"))
}

// Description returns a reference to field description of azurerm_app_configuration_feature.
func (aacf azurermAppConfigurationFeatureAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(aacf.ref.Append("description"))
}

// Enabled returns a reference to field enabled of azurerm_app_configuration_feature.
func (aacf azurermAppConfigurationFeatureAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(aacf.ref.Append("enabled"))
}

// Etag returns a reference to field etag of azurerm_app_configuration_feature.
func (aacf azurermAppConfigurationFeatureAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(aacf.ref.Append("etag"))
}

// Id returns a reference to field id of azurerm_app_configuration_feature.
func (aacf azurermAppConfigurationFeatureAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aacf.ref.Append("id"))
}

// Key returns a reference to field key of azurerm_app_configuration_feature.
func (aacf azurermAppConfigurationFeatureAttributes) Key() terra.StringValue {
	return terra.ReferenceAsString(aacf.ref.Append("key"))
}

// Label returns a reference to field label of azurerm_app_configuration_feature.
func (aacf azurermAppConfigurationFeatureAttributes) Label() terra.StringValue {
	return terra.ReferenceAsString(aacf.ref.Append("label"))
}

// Locked returns a reference to field locked of azurerm_app_configuration_feature.
func (aacf azurermAppConfigurationFeatureAttributes) Locked() terra.BoolValue {
	return terra.ReferenceAsBool(aacf.ref.Append("locked"))
}

// Name returns a reference to field name of azurerm_app_configuration_feature.
func (aacf azurermAppConfigurationFeatureAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(aacf.ref.Append("name"))
}

// PercentageFilterValue returns a reference to field percentage_filter_value of azurerm_app_configuration_feature.
func (aacf azurermAppConfigurationFeatureAttributes) PercentageFilterValue() terra.NumberValue {
	return terra.ReferenceAsNumber(aacf.ref.Append("percentage_filter_value"))
}

// Tags returns a reference to field tags of azurerm_app_configuration_feature.
func (aacf azurermAppConfigurationFeatureAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](aacf.ref.Append("tags"))
}

func (aacf azurermAppConfigurationFeatureAttributes) TargetingFilter() terra.ListValue[TargetingFilterAttributes] {
	return terra.ReferenceAsList[TargetingFilterAttributes](aacf.ref.Append("targeting_filter"))
}

func (aacf azurermAppConfigurationFeatureAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](aacf.ref.Append("timeouts"))
}

func (aacf azurermAppConfigurationFeatureAttributes) TimewindowFilter() terra.ListValue[TimewindowFilterAttributes] {
	return terra.ReferenceAsList[TimewindowFilterAttributes](aacf.ref.Append("timewindow_filter"))
}

type azurermAppConfigurationFeatureState struct {
	ConfigurationStoreId  string                  `json:"configuration_store_id"`
	Description           string                  `json:"description"`
	Enabled               bool                    `json:"enabled"`
	Etag                  string                  `json:"etag"`
	Id                    string                  `json:"id"`
	Key                   string                  `json:"key"`
	Label                 string                  `json:"label"`
	Locked                bool                    `json:"locked"`
	Name                  string                  `json:"name"`
	PercentageFilterValue float64                 `json:"percentage_filter_value"`
	Tags                  map[string]string       `json:"tags"`
	TargetingFilter       []TargetingFilterState  `json:"targeting_filter"`
	Timeouts              *TimeoutsState          `json:"timeouts"`
	TimewindowFilter      []TimewindowFilterState `json:"timewindow_filter"`
}
