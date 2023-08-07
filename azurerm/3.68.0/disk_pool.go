// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	diskpool "github.com/golingon/terraproviders/azurerm/3.68.0/diskpool"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDiskPool creates a new instance of [DiskPool].
func NewDiskPool(name string, args DiskPoolArgs) *DiskPool {
	return &DiskPool{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DiskPool)(nil)

// DiskPool represents the Terraform resource azurerm_disk_pool.
type DiskPool struct {
	Name      string
	Args      DiskPoolArgs
	state     *diskPoolState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DiskPool].
func (dp *DiskPool) Type() string {
	return "azurerm_disk_pool"
}

// LocalName returns the local name for [DiskPool].
func (dp *DiskPool) LocalName() string {
	return dp.Name
}

// Configuration returns the configuration (args) for [DiskPool].
func (dp *DiskPool) Configuration() interface{} {
	return dp.Args
}

// DependOn is used for other resources to depend on [DiskPool].
func (dp *DiskPool) DependOn() terra.Reference {
	return terra.ReferenceResource(dp)
}

// Dependencies returns the list of resources [DiskPool] depends_on.
func (dp *DiskPool) Dependencies() terra.Dependencies {
	return dp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DiskPool].
func (dp *DiskPool) LifecycleManagement() *terra.Lifecycle {
	return dp.Lifecycle
}

// Attributes returns the attributes for [DiskPool].
func (dp *DiskPool) Attributes() diskPoolAttributes {
	return diskPoolAttributes{ref: terra.ReferenceResource(dp)}
}

// ImportState imports the given attribute values into [DiskPool]'s state.
func (dp *DiskPool) ImportState(av io.Reader) error {
	dp.state = &diskPoolState{}
	if err := json.NewDecoder(av).Decode(dp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dp.Type(), dp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DiskPool] has state.
func (dp *DiskPool) State() (*diskPoolState, bool) {
	return dp.state, dp.state != nil
}

// StateMust returns the state for [DiskPool]. Panics if the state is nil.
func (dp *DiskPool) StateMust() *diskPoolState {
	if dp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dp.Type(), dp.LocalName()))
	}
	return dp.state
}

// DiskPoolArgs contains the configurations for azurerm_disk_pool.
type DiskPoolArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// SkuName: string, required
	SkuName terra.StringValue `hcl:"sku_name,attr" validate:"required"`
	// SubnetId: string, required
	SubnetId terra.StringValue `hcl:"subnet_id,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Zones: set of string, required
	Zones terra.SetValue[terra.StringValue] `hcl:"zones,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *diskpool.Timeouts `hcl:"timeouts,block"`
}
type diskPoolAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_disk_pool.
func (dp diskPoolAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dp.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_disk_pool.
func (dp diskPoolAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(dp.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_disk_pool.
func (dp diskPoolAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dp.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_disk_pool.
func (dp diskPoolAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(dp.ref.Append("resource_group_name"))
}

// SkuName returns a reference to field sku_name of azurerm_disk_pool.
func (dp diskPoolAttributes) SkuName() terra.StringValue {
	return terra.ReferenceAsString(dp.ref.Append("sku_name"))
}

// SubnetId returns a reference to field subnet_id of azurerm_disk_pool.
func (dp diskPoolAttributes) SubnetId() terra.StringValue {
	return terra.ReferenceAsString(dp.ref.Append("subnet_id"))
}

// Tags returns a reference to field tags of azurerm_disk_pool.
func (dp diskPoolAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dp.ref.Append("tags"))
}

// Zones returns a reference to field zones of azurerm_disk_pool.
func (dp diskPoolAttributes) Zones() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](dp.ref.Append("zones"))
}

func (dp diskPoolAttributes) Timeouts() diskpool.TimeoutsAttributes {
	return terra.ReferenceAsSingle[diskpool.TimeoutsAttributes](dp.ref.Append("timeouts"))
}

type diskPoolState struct {
	Id                string                  `json:"id"`
	Location          string                  `json:"location"`
	Name              string                  `json:"name"`
	ResourceGroupName string                  `json:"resource_group_name"`
	SkuName           string                  `json:"sku_name"`
	SubnetId          string                  `json:"subnet_id"`
	Tags              map[string]string       `json:"tags"`
	Zones             []string                `json:"zones"`
	Timeouts          *diskpool.TimeoutsState `json:"timeouts"`
}
