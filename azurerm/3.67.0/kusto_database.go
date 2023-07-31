// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	kustodatabase "github.com/golingon/terraproviders/azurerm/3.67.0/kustodatabase"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewKustoDatabase creates a new instance of [KustoDatabase].
func NewKustoDatabase(name string, args KustoDatabaseArgs) *KustoDatabase {
	return &KustoDatabase{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*KustoDatabase)(nil)

// KustoDatabase represents the Terraform resource azurerm_kusto_database.
type KustoDatabase struct {
	Name      string
	Args      KustoDatabaseArgs
	state     *kustoDatabaseState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [KustoDatabase].
func (kd *KustoDatabase) Type() string {
	return "azurerm_kusto_database"
}

// LocalName returns the local name for [KustoDatabase].
func (kd *KustoDatabase) LocalName() string {
	return kd.Name
}

// Configuration returns the configuration (args) for [KustoDatabase].
func (kd *KustoDatabase) Configuration() interface{} {
	return kd.Args
}

// DependOn is used for other resources to depend on [KustoDatabase].
func (kd *KustoDatabase) DependOn() terra.Reference {
	return terra.ReferenceResource(kd)
}

// Dependencies returns the list of resources [KustoDatabase] depends_on.
func (kd *KustoDatabase) Dependencies() terra.Dependencies {
	return kd.DependsOn
}

// LifecycleManagement returns the lifecycle block for [KustoDatabase].
func (kd *KustoDatabase) LifecycleManagement() *terra.Lifecycle {
	return kd.Lifecycle
}

// Attributes returns the attributes for [KustoDatabase].
func (kd *KustoDatabase) Attributes() kustoDatabaseAttributes {
	return kustoDatabaseAttributes{ref: terra.ReferenceResource(kd)}
}

// ImportState imports the given attribute values into [KustoDatabase]'s state.
func (kd *KustoDatabase) ImportState(av io.Reader) error {
	kd.state = &kustoDatabaseState{}
	if err := json.NewDecoder(av).Decode(kd.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", kd.Type(), kd.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [KustoDatabase] has state.
func (kd *KustoDatabase) State() (*kustoDatabaseState, bool) {
	return kd.state, kd.state != nil
}

// StateMust returns the state for [KustoDatabase]. Panics if the state is nil.
func (kd *KustoDatabase) StateMust() *kustoDatabaseState {
	if kd.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", kd.Type(), kd.LocalName()))
	}
	return kd.state
}

// KustoDatabaseArgs contains the configurations for azurerm_kusto_database.
type KustoDatabaseArgs struct {
	// ClusterName: string, required
	ClusterName terra.StringValue `hcl:"cluster_name,attr" validate:"required"`
	// HotCachePeriod: string, optional
	HotCachePeriod terra.StringValue `hcl:"hot_cache_period,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// SoftDeletePeriod: string, optional
	SoftDeletePeriod terra.StringValue `hcl:"soft_delete_period,attr"`
	// Timeouts: optional
	Timeouts *kustodatabase.Timeouts `hcl:"timeouts,block"`
}
type kustoDatabaseAttributes struct {
	ref terra.Reference
}

// ClusterName returns a reference to field cluster_name of azurerm_kusto_database.
func (kd kustoDatabaseAttributes) ClusterName() terra.StringValue {
	return terra.ReferenceAsString(kd.ref.Append("cluster_name"))
}

// HotCachePeriod returns a reference to field hot_cache_period of azurerm_kusto_database.
func (kd kustoDatabaseAttributes) HotCachePeriod() terra.StringValue {
	return terra.ReferenceAsString(kd.ref.Append("hot_cache_period"))
}

// Id returns a reference to field id of azurerm_kusto_database.
func (kd kustoDatabaseAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(kd.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_kusto_database.
func (kd kustoDatabaseAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(kd.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_kusto_database.
func (kd kustoDatabaseAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(kd.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_kusto_database.
func (kd kustoDatabaseAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(kd.ref.Append("resource_group_name"))
}

// Size returns a reference to field size of azurerm_kusto_database.
func (kd kustoDatabaseAttributes) Size() terra.NumberValue {
	return terra.ReferenceAsNumber(kd.ref.Append("size"))
}

// SoftDeletePeriod returns a reference to field soft_delete_period of azurerm_kusto_database.
func (kd kustoDatabaseAttributes) SoftDeletePeriod() terra.StringValue {
	return terra.ReferenceAsString(kd.ref.Append("soft_delete_period"))
}

func (kd kustoDatabaseAttributes) Timeouts() kustodatabase.TimeoutsAttributes {
	return terra.ReferenceAsSingle[kustodatabase.TimeoutsAttributes](kd.ref.Append("timeouts"))
}

type kustoDatabaseState struct {
	ClusterName       string                       `json:"cluster_name"`
	HotCachePeriod    string                       `json:"hot_cache_period"`
	Id                string                       `json:"id"`
	Location          string                       `json:"location"`
	Name              string                       `json:"name"`
	ResourceGroupName string                       `json:"resource_group_name"`
	Size              float64                      `json:"size"`
	SoftDeletePeriod  string                       `json:"soft_delete_period"`
	Timeouts          *kustodatabase.TimeoutsState `json:"timeouts"`
}
