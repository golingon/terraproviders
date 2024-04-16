// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_lab_service_lab

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

// Resource represents the Terraform resource azurerm_lab_service_lab.
type Resource struct {
	Name      string
	Args      Args
	state     *azurermLabServiceLabState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (alsl *Resource) Type() string {
	return "azurerm_lab_service_lab"
}

// LocalName returns the local name for [Resource].
func (alsl *Resource) LocalName() string {
	return alsl.Name
}

// Configuration returns the configuration (args) for [Resource].
func (alsl *Resource) Configuration() interface{} {
	return alsl.Args
}

// DependOn is used for other resources to depend on [Resource].
func (alsl *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(alsl)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (alsl *Resource) Dependencies() terra.Dependencies {
	return alsl.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (alsl *Resource) LifecycleManagement() *terra.Lifecycle {
	return alsl.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (alsl *Resource) Attributes() azurermLabServiceLabAttributes {
	return azurermLabServiceLabAttributes{ref: terra.ReferenceResource(alsl)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (alsl *Resource) ImportState(state io.Reader) error {
	alsl.state = &azurermLabServiceLabState{}
	if err := json.NewDecoder(state).Decode(alsl.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", alsl.Type(), alsl.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (alsl *Resource) State() (*azurermLabServiceLabState, bool) {
	return alsl.state, alsl.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (alsl *Resource) StateMust() *azurermLabServiceLabState {
	if alsl.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", alsl.Type(), alsl.LocalName()))
	}
	return alsl.state
}

// Args contains the configurations for azurerm_lab_service_lab.
type Args struct {
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
	AutoShutdown *AutoShutdown `hcl:"auto_shutdown,block"`
	// ConnectionSetting: required
	ConnectionSetting *ConnectionSetting `hcl:"connection_setting,block" validate:"required"`
	// Network: optional
	Network *Network `hcl:"network,block"`
	// Roster: optional
	Roster *Roster `hcl:"roster,block"`
	// Security: required
	Security *Security `hcl:"security,block" validate:"required"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
	// VirtualMachine: required
	VirtualMachine *VirtualMachine `hcl:"virtual_machine,block" validate:"required"`
}

type azurermLabServiceLabAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of azurerm_lab_service_lab.
func (alsl azurermLabServiceLabAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(alsl.ref.Append("description"))
}

// Id returns a reference to field id of azurerm_lab_service_lab.
func (alsl azurermLabServiceLabAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(alsl.ref.Append("id"))
}

// LabPlanId returns a reference to field lab_plan_id of azurerm_lab_service_lab.
func (alsl azurermLabServiceLabAttributes) LabPlanId() terra.StringValue {
	return terra.ReferenceAsString(alsl.ref.Append("lab_plan_id"))
}

// Location returns a reference to field location of azurerm_lab_service_lab.
func (alsl azurermLabServiceLabAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(alsl.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_lab_service_lab.
func (alsl azurermLabServiceLabAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(alsl.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_lab_service_lab.
func (alsl azurermLabServiceLabAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(alsl.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_lab_service_lab.
func (alsl azurermLabServiceLabAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](alsl.ref.Append("tags"))
}

// Title returns a reference to field title of azurerm_lab_service_lab.
func (alsl azurermLabServiceLabAttributes) Title() terra.StringValue {
	return terra.ReferenceAsString(alsl.ref.Append("title"))
}

func (alsl azurermLabServiceLabAttributes) AutoShutdown() terra.ListValue[AutoShutdownAttributes] {
	return terra.ReferenceAsList[AutoShutdownAttributes](alsl.ref.Append("auto_shutdown"))
}

func (alsl azurermLabServiceLabAttributes) ConnectionSetting() terra.ListValue[ConnectionSettingAttributes] {
	return terra.ReferenceAsList[ConnectionSettingAttributes](alsl.ref.Append("connection_setting"))
}

func (alsl azurermLabServiceLabAttributes) Network() terra.ListValue[NetworkAttributes] {
	return terra.ReferenceAsList[NetworkAttributes](alsl.ref.Append("network"))
}

func (alsl azurermLabServiceLabAttributes) Roster() terra.ListValue[RosterAttributes] {
	return terra.ReferenceAsList[RosterAttributes](alsl.ref.Append("roster"))
}

func (alsl azurermLabServiceLabAttributes) Security() terra.ListValue[SecurityAttributes] {
	return terra.ReferenceAsList[SecurityAttributes](alsl.ref.Append("security"))
}

func (alsl azurermLabServiceLabAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](alsl.ref.Append("timeouts"))
}

func (alsl azurermLabServiceLabAttributes) VirtualMachine() terra.ListValue[VirtualMachineAttributes] {
	return terra.ReferenceAsList[VirtualMachineAttributes](alsl.ref.Append("virtual_machine"))
}

type azurermLabServiceLabState struct {
	Description       string                   `json:"description"`
	Id                string                   `json:"id"`
	LabPlanId         string                   `json:"lab_plan_id"`
	Location          string                   `json:"location"`
	Name              string                   `json:"name"`
	ResourceGroupName string                   `json:"resource_group_name"`
	Tags              map[string]string        `json:"tags"`
	Title             string                   `json:"title"`
	AutoShutdown      []AutoShutdownState      `json:"auto_shutdown"`
	ConnectionSetting []ConnectionSettingState `json:"connection_setting"`
	Network           []NetworkState           `json:"network"`
	Roster            []RosterState            `json:"roster"`
	Security          []SecurityState          `json:"security"`
	Timeouts          *TimeoutsState           `json:"timeouts"`
	VirtualMachine    []VirtualMachineState    `json:"virtual_machine"`
}
