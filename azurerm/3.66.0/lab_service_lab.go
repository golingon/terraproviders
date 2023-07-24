// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	labservicelab "github.com/golingon/terraproviders/azurerm/3.66.0/labservicelab"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewLabServiceLab creates a new instance of [LabServiceLab].
func NewLabServiceLab(name string, args LabServiceLabArgs) *LabServiceLab {
	return &LabServiceLab{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*LabServiceLab)(nil)

// LabServiceLab represents the Terraform resource azurerm_lab_service_lab.
type LabServiceLab struct {
	Name      string
	Args      LabServiceLabArgs
	state     *labServiceLabState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [LabServiceLab].
func (lsl *LabServiceLab) Type() string {
	return "azurerm_lab_service_lab"
}

// LocalName returns the local name for [LabServiceLab].
func (lsl *LabServiceLab) LocalName() string {
	return lsl.Name
}

// Configuration returns the configuration (args) for [LabServiceLab].
func (lsl *LabServiceLab) Configuration() interface{} {
	return lsl.Args
}

// DependOn is used for other resources to depend on [LabServiceLab].
func (lsl *LabServiceLab) DependOn() terra.Reference {
	return terra.ReferenceResource(lsl)
}

// Dependencies returns the list of resources [LabServiceLab] depends_on.
func (lsl *LabServiceLab) Dependencies() terra.Dependencies {
	return lsl.DependsOn
}

// LifecycleManagement returns the lifecycle block for [LabServiceLab].
func (lsl *LabServiceLab) LifecycleManagement() *terra.Lifecycle {
	return lsl.Lifecycle
}

// Attributes returns the attributes for [LabServiceLab].
func (lsl *LabServiceLab) Attributes() labServiceLabAttributes {
	return labServiceLabAttributes{ref: terra.ReferenceResource(lsl)}
}

// ImportState imports the given attribute values into [LabServiceLab]'s state.
func (lsl *LabServiceLab) ImportState(av io.Reader) error {
	lsl.state = &labServiceLabState{}
	if err := json.NewDecoder(av).Decode(lsl.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", lsl.Type(), lsl.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [LabServiceLab] has state.
func (lsl *LabServiceLab) State() (*labServiceLabState, bool) {
	return lsl.state, lsl.state != nil
}

// StateMust returns the state for [LabServiceLab]. Panics if the state is nil.
func (lsl *LabServiceLab) StateMust() *labServiceLabState {
	if lsl.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", lsl.Type(), lsl.LocalName()))
	}
	return lsl.state
}

// LabServiceLabArgs contains the configurations for azurerm_lab_service_lab.
type LabServiceLabArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LabPlanId: string, optional
	LabPlanId terra.StringValue `hcl:"lab_plan_id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Title: string, required
	Title terra.StringValue `hcl:"title,attr" validate:"required"`
	// AutoShutdown: optional
	AutoShutdown *labservicelab.AutoShutdown `hcl:"auto_shutdown,block"`
	// ConnectionSetting: optional
	ConnectionSetting *labservicelab.ConnectionSetting `hcl:"connection_setting,block"`
	// Network: optional
	Network *labservicelab.Network `hcl:"network,block"`
	// Roster: optional
	Roster *labservicelab.Roster `hcl:"roster,block"`
	// Security: required
	Security *labservicelab.Security `hcl:"security,block" validate:"required"`
	// Timeouts: optional
	Timeouts *labservicelab.Timeouts `hcl:"timeouts,block"`
	// VirtualMachine: required
	VirtualMachine *labservicelab.VirtualMachine `hcl:"virtual_machine,block" validate:"required"`
}
type labServiceLabAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of azurerm_lab_service_lab.
func (lsl labServiceLabAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(lsl.ref.Append("description"))
}

// Id returns a reference to field id of azurerm_lab_service_lab.
func (lsl labServiceLabAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(lsl.ref.Append("id"))
}

// LabPlanId returns a reference to field lab_plan_id of azurerm_lab_service_lab.
func (lsl labServiceLabAttributes) LabPlanId() terra.StringValue {
	return terra.ReferenceAsString(lsl.ref.Append("lab_plan_id"))
}

// Location returns a reference to field location of azurerm_lab_service_lab.
func (lsl labServiceLabAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(lsl.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_lab_service_lab.
func (lsl labServiceLabAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(lsl.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_lab_service_lab.
func (lsl labServiceLabAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(lsl.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_lab_service_lab.
func (lsl labServiceLabAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](lsl.ref.Append("tags"))
}

// Title returns a reference to field title of azurerm_lab_service_lab.
func (lsl labServiceLabAttributes) Title() terra.StringValue {
	return terra.ReferenceAsString(lsl.ref.Append("title"))
}

func (lsl labServiceLabAttributes) AutoShutdown() terra.ListValue[labservicelab.AutoShutdownAttributes] {
	return terra.ReferenceAsList[labservicelab.AutoShutdownAttributes](lsl.ref.Append("auto_shutdown"))
}

func (lsl labServiceLabAttributes) ConnectionSetting() terra.ListValue[labservicelab.ConnectionSettingAttributes] {
	return terra.ReferenceAsList[labservicelab.ConnectionSettingAttributes](lsl.ref.Append("connection_setting"))
}

func (lsl labServiceLabAttributes) Network() terra.ListValue[labservicelab.NetworkAttributes] {
	return terra.ReferenceAsList[labservicelab.NetworkAttributes](lsl.ref.Append("network"))
}

func (lsl labServiceLabAttributes) Roster() terra.ListValue[labservicelab.RosterAttributes] {
	return terra.ReferenceAsList[labservicelab.RosterAttributes](lsl.ref.Append("roster"))
}

func (lsl labServiceLabAttributes) Security() terra.ListValue[labservicelab.SecurityAttributes] {
	return terra.ReferenceAsList[labservicelab.SecurityAttributes](lsl.ref.Append("security"))
}

func (lsl labServiceLabAttributes) Timeouts() labservicelab.TimeoutsAttributes {
	return terra.ReferenceAsSingle[labservicelab.TimeoutsAttributes](lsl.ref.Append("timeouts"))
}

func (lsl labServiceLabAttributes) VirtualMachine() terra.ListValue[labservicelab.VirtualMachineAttributes] {
	return terra.ReferenceAsList[labservicelab.VirtualMachineAttributes](lsl.ref.Append("virtual_machine"))
}

type labServiceLabState struct {
	Description       string                                 `json:"description"`
	Id                string                                 `json:"id"`
	LabPlanId         string                                 `json:"lab_plan_id"`
	Location          string                                 `json:"location"`
	Name              string                                 `json:"name"`
	ResourceGroupName string                                 `json:"resource_group_name"`
	Tags              map[string]string                      `json:"tags"`
	Title             string                                 `json:"title"`
	AutoShutdown      []labservicelab.AutoShutdownState      `json:"auto_shutdown"`
	ConnectionSetting []labservicelab.ConnectionSettingState `json:"connection_setting"`
	Network           []labservicelab.NetworkState           `json:"network"`
	Roster            []labservicelab.RosterState            `json:"roster"`
	Security          []labservicelab.SecurityState          `json:"security"`
	Timeouts          *labservicelab.TimeoutsState           `json:"timeouts"`
	VirtualMachine    []labservicelab.VirtualMachineState    `json:"virtual_machine"`
}
