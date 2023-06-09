// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	labserviceplan "github.com/golingon/terraproviders/azurerm/3.58.0/labserviceplan"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewLabServicePlan creates a new instance of [LabServicePlan].
func NewLabServicePlan(name string, args LabServicePlanArgs) *LabServicePlan {
	return &LabServicePlan{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*LabServicePlan)(nil)

// LabServicePlan represents the Terraform resource azurerm_lab_service_plan.
type LabServicePlan struct {
	Name      string
	Args      LabServicePlanArgs
	state     *labServicePlanState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [LabServicePlan].
func (lsp *LabServicePlan) Type() string {
	return "azurerm_lab_service_plan"
}

// LocalName returns the local name for [LabServicePlan].
func (lsp *LabServicePlan) LocalName() string {
	return lsp.Name
}

// Configuration returns the configuration (args) for [LabServicePlan].
func (lsp *LabServicePlan) Configuration() interface{} {
	return lsp.Args
}

// DependOn is used for other resources to depend on [LabServicePlan].
func (lsp *LabServicePlan) DependOn() terra.Reference {
	return terra.ReferenceResource(lsp)
}

// Dependencies returns the list of resources [LabServicePlan] depends_on.
func (lsp *LabServicePlan) Dependencies() terra.Dependencies {
	return lsp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [LabServicePlan].
func (lsp *LabServicePlan) LifecycleManagement() *terra.Lifecycle {
	return lsp.Lifecycle
}

// Attributes returns the attributes for [LabServicePlan].
func (lsp *LabServicePlan) Attributes() labServicePlanAttributes {
	return labServicePlanAttributes{ref: terra.ReferenceResource(lsp)}
}

// ImportState imports the given attribute values into [LabServicePlan]'s state.
func (lsp *LabServicePlan) ImportState(av io.Reader) error {
	lsp.state = &labServicePlanState{}
	if err := json.NewDecoder(av).Decode(lsp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", lsp.Type(), lsp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [LabServicePlan] has state.
func (lsp *LabServicePlan) State() (*labServicePlanState, bool) {
	return lsp.state, lsp.state != nil
}

// StateMust returns the state for [LabServicePlan]. Panics if the state is nil.
func (lsp *LabServicePlan) StateMust() *labServicePlanState {
	if lsp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", lsp.Type(), lsp.LocalName()))
	}
	return lsp.state
}

// LabServicePlanArgs contains the configurations for azurerm_lab_service_plan.
type LabServicePlanArgs struct {
	// AllowedRegions: list of string, required
	AllowedRegions terra.ListValue[terra.StringValue] `hcl:"allowed_regions,attr" validate:"required"`
	// DefaultNetworkSubnetId: string, optional
	DefaultNetworkSubnetId terra.StringValue `hcl:"default_network_subnet_id,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// SharedGalleryId: string, optional
	SharedGalleryId terra.StringValue `hcl:"shared_gallery_id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// DefaultAutoShutdown: optional
	DefaultAutoShutdown *labserviceplan.DefaultAutoShutdown `hcl:"default_auto_shutdown,block"`
	// DefaultConnection: optional
	DefaultConnection *labserviceplan.DefaultConnection `hcl:"default_connection,block"`
	// Support: optional
	Support *labserviceplan.Support `hcl:"support,block"`
	// Timeouts: optional
	Timeouts *labserviceplan.Timeouts `hcl:"timeouts,block"`
}
type labServicePlanAttributes struct {
	ref terra.Reference
}

// AllowedRegions returns a reference to field allowed_regions of azurerm_lab_service_plan.
func (lsp labServicePlanAttributes) AllowedRegions() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](lsp.ref.Append("allowed_regions"))
}

// DefaultNetworkSubnetId returns a reference to field default_network_subnet_id of azurerm_lab_service_plan.
func (lsp labServicePlanAttributes) DefaultNetworkSubnetId() terra.StringValue {
	return terra.ReferenceAsString(lsp.ref.Append("default_network_subnet_id"))
}

// Id returns a reference to field id of azurerm_lab_service_plan.
func (lsp labServicePlanAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(lsp.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_lab_service_plan.
func (lsp labServicePlanAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(lsp.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_lab_service_plan.
func (lsp labServicePlanAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(lsp.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_lab_service_plan.
func (lsp labServicePlanAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(lsp.ref.Append("resource_group_name"))
}

// SharedGalleryId returns a reference to field shared_gallery_id of azurerm_lab_service_plan.
func (lsp labServicePlanAttributes) SharedGalleryId() terra.StringValue {
	return terra.ReferenceAsString(lsp.ref.Append("shared_gallery_id"))
}

// Tags returns a reference to field tags of azurerm_lab_service_plan.
func (lsp labServicePlanAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](lsp.ref.Append("tags"))
}

func (lsp labServicePlanAttributes) DefaultAutoShutdown() terra.ListValue[labserviceplan.DefaultAutoShutdownAttributes] {
	return terra.ReferenceAsList[labserviceplan.DefaultAutoShutdownAttributes](lsp.ref.Append("default_auto_shutdown"))
}

func (lsp labServicePlanAttributes) DefaultConnection() terra.ListValue[labserviceplan.DefaultConnectionAttributes] {
	return terra.ReferenceAsList[labserviceplan.DefaultConnectionAttributes](lsp.ref.Append("default_connection"))
}

func (lsp labServicePlanAttributes) Support() terra.ListValue[labserviceplan.SupportAttributes] {
	return terra.ReferenceAsList[labserviceplan.SupportAttributes](lsp.ref.Append("support"))
}

func (lsp labServicePlanAttributes) Timeouts() labserviceplan.TimeoutsAttributes {
	return terra.ReferenceAsSingle[labserviceplan.TimeoutsAttributes](lsp.ref.Append("timeouts"))
}

type labServicePlanState struct {
	AllowedRegions         []string                                  `json:"allowed_regions"`
	DefaultNetworkSubnetId string                                    `json:"default_network_subnet_id"`
	Id                     string                                    `json:"id"`
	Location               string                                    `json:"location"`
	Name                   string                                    `json:"name"`
	ResourceGroupName      string                                    `json:"resource_group_name"`
	SharedGalleryId        string                                    `json:"shared_gallery_id"`
	Tags                   map[string]string                         `json:"tags"`
	DefaultAutoShutdown    []labserviceplan.DefaultAutoShutdownState `json:"default_auto_shutdown"`
	DefaultConnection      []labserviceplan.DefaultConnectionState   `json:"default_connection"`
	Support                []labserviceplan.SupportState             `json:"support"`
	Timeouts               *labserviceplan.TimeoutsState             `json:"timeouts"`
}
