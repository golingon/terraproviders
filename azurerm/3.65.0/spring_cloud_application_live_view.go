// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	springcloudapplicationliveview "github.com/golingon/terraproviders/azurerm/3.65.0/springcloudapplicationliveview"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSpringCloudApplicationLiveView creates a new instance of [SpringCloudApplicationLiveView].
func NewSpringCloudApplicationLiveView(name string, args SpringCloudApplicationLiveViewArgs) *SpringCloudApplicationLiveView {
	return &SpringCloudApplicationLiveView{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SpringCloudApplicationLiveView)(nil)

// SpringCloudApplicationLiveView represents the Terraform resource azurerm_spring_cloud_application_live_view.
type SpringCloudApplicationLiveView struct {
	Name      string
	Args      SpringCloudApplicationLiveViewArgs
	state     *springCloudApplicationLiveViewState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SpringCloudApplicationLiveView].
func (scalv *SpringCloudApplicationLiveView) Type() string {
	return "azurerm_spring_cloud_application_live_view"
}

// LocalName returns the local name for [SpringCloudApplicationLiveView].
func (scalv *SpringCloudApplicationLiveView) LocalName() string {
	return scalv.Name
}

// Configuration returns the configuration (args) for [SpringCloudApplicationLiveView].
func (scalv *SpringCloudApplicationLiveView) Configuration() interface{} {
	return scalv.Args
}

// DependOn is used for other resources to depend on [SpringCloudApplicationLiveView].
func (scalv *SpringCloudApplicationLiveView) DependOn() terra.Reference {
	return terra.ReferenceResource(scalv)
}

// Dependencies returns the list of resources [SpringCloudApplicationLiveView] depends_on.
func (scalv *SpringCloudApplicationLiveView) Dependencies() terra.Dependencies {
	return scalv.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SpringCloudApplicationLiveView].
func (scalv *SpringCloudApplicationLiveView) LifecycleManagement() *terra.Lifecycle {
	return scalv.Lifecycle
}

// Attributes returns the attributes for [SpringCloudApplicationLiveView].
func (scalv *SpringCloudApplicationLiveView) Attributes() springCloudApplicationLiveViewAttributes {
	return springCloudApplicationLiveViewAttributes{ref: terra.ReferenceResource(scalv)}
}

// ImportState imports the given attribute values into [SpringCloudApplicationLiveView]'s state.
func (scalv *SpringCloudApplicationLiveView) ImportState(av io.Reader) error {
	scalv.state = &springCloudApplicationLiveViewState{}
	if err := json.NewDecoder(av).Decode(scalv.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", scalv.Type(), scalv.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SpringCloudApplicationLiveView] has state.
func (scalv *SpringCloudApplicationLiveView) State() (*springCloudApplicationLiveViewState, bool) {
	return scalv.state, scalv.state != nil
}

// StateMust returns the state for [SpringCloudApplicationLiveView]. Panics if the state is nil.
func (scalv *SpringCloudApplicationLiveView) StateMust() *springCloudApplicationLiveViewState {
	if scalv.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", scalv.Type(), scalv.LocalName()))
	}
	return scalv.state
}

// SpringCloudApplicationLiveViewArgs contains the configurations for azurerm_spring_cloud_application_live_view.
type SpringCloudApplicationLiveViewArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// SpringCloudServiceId: string, required
	SpringCloudServiceId terra.StringValue `hcl:"spring_cloud_service_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *springcloudapplicationliveview.Timeouts `hcl:"timeouts,block"`
}
type springCloudApplicationLiveViewAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_spring_cloud_application_live_view.
func (scalv springCloudApplicationLiveViewAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(scalv.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_spring_cloud_application_live_view.
func (scalv springCloudApplicationLiveViewAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(scalv.ref.Append("name"))
}

// SpringCloudServiceId returns a reference to field spring_cloud_service_id of azurerm_spring_cloud_application_live_view.
func (scalv springCloudApplicationLiveViewAttributes) SpringCloudServiceId() terra.StringValue {
	return terra.ReferenceAsString(scalv.ref.Append("spring_cloud_service_id"))
}

func (scalv springCloudApplicationLiveViewAttributes) Timeouts() springcloudapplicationliveview.TimeoutsAttributes {
	return terra.ReferenceAsSingle[springcloudapplicationliveview.TimeoutsAttributes](scalv.ref.Append("timeouts"))
}

type springCloudApplicationLiveViewState struct {
	Id                   string                                        `json:"id"`
	Name                 string                                        `json:"name"`
	SpringCloudServiceId string                                        `json:"spring_cloud_service_id"`
	Timeouts             *springcloudapplicationliveview.TimeoutsState `json:"timeouts"`
}
