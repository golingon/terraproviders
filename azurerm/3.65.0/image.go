// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	image "github.com/golingon/terraproviders/azurerm/3.65.0/image"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewImage creates a new instance of [Image].
func NewImage(name string, args ImageArgs) *Image {
	return &Image{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Image)(nil)

// Image represents the Terraform resource azurerm_image.
type Image struct {
	Name      string
	Args      ImageArgs
	state     *imageState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Image].
func (i *Image) Type() string {
	return "azurerm_image"
}

// LocalName returns the local name for [Image].
func (i *Image) LocalName() string {
	return i.Name
}

// Configuration returns the configuration (args) for [Image].
func (i *Image) Configuration() interface{} {
	return i.Args
}

// DependOn is used for other resources to depend on [Image].
func (i *Image) DependOn() terra.Reference {
	return terra.ReferenceResource(i)
}

// Dependencies returns the list of resources [Image] depends_on.
func (i *Image) Dependencies() terra.Dependencies {
	return i.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Image].
func (i *Image) LifecycleManagement() *terra.Lifecycle {
	return i.Lifecycle
}

// Attributes returns the attributes for [Image].
func (i *Image) Attributes() imageAttributes {
	return imageAttributes{ref: terra.ReferenceResource(i)}
}

// ImportState imports the given attribute values into [Image]'s state.
func (i *Image) ImportState(av io.Reader) error {
	i.state = &imageState{}
	if err := json.NewDecoder(av).Decode(i.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", i.Type(), i.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Image] has state.
func (i *Image) State() (*imageState, bool) {
	return i.state, i.state != nil
}

// StateMust returns the state for [Image]. Panics if the state is nil.
func (i *Image) StateMust() *imageState {
	if i.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", i.Type(), i.LocalName()))
	}
	return i.state
}

// ImageArgs contains the configurations for azurerm_image.
type ImageArgs struct {
	// HyperVGeneration: string, optional
	HyperVGeneration terra.StringValue `hcl:"hyper_v_generation,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// SourceVirtualMachineId: string, optional
	SourceVirtualMachineId terra.StringValue `hcl:"source_virtual_machine_id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// ZoneResilient: bool, optional
	ZoneResilient terra.BoolValue `hcl:"zone_resilient,attr"`
	// DataDisk: min=0
	DataDisk []image.DataDisk `hcl:"data_disk,block" validate:"min=0"`
	// OsDisk: optional
	OsDisk *image.OsDisk `hcl:"os_disk,block"`
	// Timeouts: optional
	Timeouts *image.Timeouts `hcl:"timeouts,block"`
}
type imageAttributes struct {
	ref terra.Reference
}

// HyperVGeneration returns a reference to field hyper_v_generation of azurerm_image.
func (i imageAttributes) HyperVGeneration() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("hyper_v_generation"))
}

// Id returns a reference to field id of azurerm_image.
func (i imageAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_image.
func (i imageAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_image.
func (i imageAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_image.
func (i imageAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("resource_group_name"))
}

// SourceVirtualMachineId returns a reference to field source_virtual_machine_id of azurerm_image.
func (i imageAttributes) SourceVirtualMachineId() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("source_virtual_machine_id"))
}

// Tags returns a reference to field tags of azurerm_image.
func (i imageAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](i.ref.Append("tags"))
}

// ZoneResilient returns a reference to field zone_resilient of azurerm_image.
func (i imageAttributes) ZoneResilient() terra.BoolValue {
	return terra.ReferenceAsBool(i.ref.Append("zone_resilient"))
}

func (i imageAttributes) DataDisk() terra.ListValue[image.DataDiskAttributes] {
	return terra.ReferenceAsList[image.DataDiskAttributes](i.ref.Append("data_disk"))
}

func (i imageAttributes) OsDisk() terra.ListValue[image.OsDiskAttributes] {
	return terra.ReferenceAsList[image.OsDiskAttributes](i.ref.Append("os_disk"))
}

func (i imageAttributes) Timeouts() image.TimeoutsAttributes {
	return terra.ReferenceAsSingle[image.TimeoutsAttributes](i.ref.Append("timeouts"))
}

type imageState struct {
	HyperVGeneration       string                `json:"hyper_v_generation"`
	Id                     string                `json:"id"`
	Location               string                `json:"location"`
	Name                   string                `json:"name"`
	ResourceGroupName      string                `json:"resource_group_name"`
	SourceVirtualMachineId string                `json:"source_virtual_machine_id"`
	Tags                   map[string]string     `json:"tags"`
	ZoneResilient          bool                  `json:"zone_resilient"`
	DataDisk               []image.DataDiskState `json:"data_disk"`
	OsDisk                 []image.OsDiskState   `json:"os_disk"`
	Timeouts               *image.TimeoutsState  `json:"timeouts"`
}
