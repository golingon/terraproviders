// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	apimanagementtag "github.com/golingon/terraproviders/azurerm/3.64.0/apimanagementtag"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewApiManagementTag creates a new instance of [ApiManagementTag].
func NewApiManagementTag(name string, args ApiManagementTagArgs) *ApiManagementTag {
	return &ApiManagementTag{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ApiManagementTag)(nil)

// ApiManagementTag represents the Terraform resource azurerm_api_management_tag.
type ApiManagementTag struct {
	Name      string
	Args      ApiManagementTagArgs
	state     *apiManagementTagState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ApiManagementTag].
func (amt *ApiManagementTag) Type() string {
	return "azurerm_api_management_tag"
}

// LocalName returns the local name for [ApiManagementTag].
func (amt *ApiManagementTag) LocalName() string {
	return amt.Name
}

// Configuration returns the configuration (args) for [ApiManagementTag].
func (amt *ApiManagementTag) Configuration() interface{} {
	return amt.Args
}

// DependOn is used for other resources to depend on [ApiManagementTag].
func (amt *ApiManagementTag) DependOn() terra.Reference {
	return terra.ReferenceResource(amt)
}

// Dependencies returns the list of resources [ApiManagementTag] depends_on.
func (amt *ApiManagementTag) Dependencies() terra.Dependencies {
	return amt.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ApiManagementTag].
func (amt *ApiManagementTag) LifecycleManagement() *terra.Lifecycle {
	return amt.Lifecycle
}

// Attributes returns the attributes for [ApiManagementTag].
func (amt *ApiManagementTag) Attributes() apiManagementTagAttributes {
	return apiManagementTagAttributes{ref: terra.ReferenceResource(amt)}
}

// ImportState imports the given attribute values into [ApiManagementTag]'s state.
func (amt *ApiManagementTag) ImportState(av io.Reader) error {
	amt.state = &apiManagementTagState{}
	if err := json.NewDecoder(av).Decode(amt.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", amt.Type(), amt.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ApiManagementTag] has state.
func (amt *ApiManagementTag) State() (*apiManagementTagState, bool) {
	return amt.state, amt.state != nil
}

// StateMust returns the state for [ApiManagementTag]. Panics if the state is nil.
func (amt *ApiManagementTag) StateMust() *apiManagementTagState {
	if amt.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", amt.Type(), amt.LocalName()))
	}
	return amt.state
}

// ApiManagementTagArgs contains the configurations for azurerm_api_management_tag.
type ApiManagementTagArgs struct {
	// ApiManagementId: string, required
	ApiManagementId terra.StringValue `hcl:"api_management_id,attr" validate:"required"`
	// DisplayName: string, optional
	DisplayName terra.StringValue `hcl:"display_name,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *apimanagementtag.Timeouts `hcl:"timeouts,block"`
}
type apiManagementTagAttributes struct {
	ref terra.Reference
}

// ApiManagementId returns a reference to field api_management_id of azurerm_api_management_tag.
func (amt apiManagementTagAttributes) ApiManagementId() terra.StringValue {
	return terra.ReferenceAsString(amt.ref.Append("api_management_id"))
}

// DisplayName returns a reference to field display_name of azurerm_api_management_tag.
func (amt apiManagementTagAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(amt.ref.Append("display_name"))
}

// Id returns a reference to field id of azurerm_api_management_tag.
func (amt apiManagementTagAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(amt.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_api_management_tag.
func (amt apiManagementTagAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(amt.ref.Append("name"))
}

func (amt apiManagementTagAttributes) Timeouts() apimanagementtag.TimeoutsAttributes {
	return terra.ReferenceAsSingle[apimanagementtag.TimeoutsAttributes](amt.ref.Append("timeouts"))
}

type apiManagementTagState struct {
	ApiManagementId string                          `json:"api_management_id"`
	DisplayName     string                          `json:"display_name"`
	Id              string                          `json:"id"`
	Name            string                          `json:"name"`
	Timeouts        *apimanagementtag.TimeoutsState `json:"timeouts"`
}
