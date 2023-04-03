// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	apimanagementapitagdescription "github.com/golingon/terraproviders/azurerm/3.49.0/apimanagementapitagdescription"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewApiManagementApiTagDescription creates a new instance of [ApiManagementApiTagDescription].
func NewApiManagementApiTagDescription(name string, args ApiManagementApiTagDescriptionArgs) *ApiManagementApiTagDescription {
	return &ApiManagementApiTagDescription{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ApiManagementApiTagDescription)(nil)

// ApiManagementApiTagDescription represents the Terraform resource azurerm_api_management_api_tag_description.
type ApiManagementApiTagDescription struct {
	Name      string
	Args      ApiManagementApiTagDescriptionArgs
	state     *apiManagementApiTagDescriptionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ApiManagementApiTagDescription].
func (amatd *ApiManagementApiTagDescription) Type() string {
	return "azurerm_api_management_api_tag_description"
}

// LocalName returns the local name for [ApiManagementApiTagDescription].
func (amatd *ApiManagementApiTagDescription) LocalName() string {
	return amatd.Name
}

// Configuration returns the configuration (args) for [ApiManagementApiTagDescription].
func (amatd *ApiManagementApiTagDescription) Configuration() interface{} {
	return amatd.Args
}

// DependOn is used for other resources to depend on [ApiManagementApiTagDescription].
func (amatd *ApiManagementApiTagDescription) DependOn() terra.Reference {
	return terra.ReferenceResource(amatd)
}

// Dependencies returns the list of resources [ApiManagementApiTagDescription] depends_on.
func (amatd *ApiManagementApiTagDescription) Dependencies() terra.Dependencies {
	return amatd.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ApiManagementApiTagDescription].
func (amatd *ApiManagementApiTagDescription) LifecycleManagement() *terra.Lifecycle {
	return amatd.Lifecycle
}

// Attributes returns the attributes for [ApiManagementApiTagDescription].
func (amatd *ApiManagementApiTagDescription) Attributes() apiManagementApiTagDescriptionAttributes {
	return apiManagementApiTagDescriptionAttributes{ref: terra.ReferenceResource(amatd)}
}

// ImportState imports the given attribute values into [ApiManagementApiTagDescription]'s state.
func (amatd *ApiManagementApiTagDescription) ImportState(av io.Reader) error {
	amatd.state = &apiManagementApiTagDescriptionState{}
	if err := json.NewDecoder(av).Decode(amatd.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", amatd.Type(), amatd.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ApiManagementApiTagDescription] has state.
func (amatd *ApiManagementApiTagDescription) State() (*apiManagementApiTagDescriptionState, bool) {
	return amatd.state, amatd.state != nil
}

// StateMust returns the state for [ApiManagementApiTagDescription]. Panics if the state is nil.
func (amatd *ApiManagementApiTagDescription) StateMust() *apiManagementApiTagDescriptionState {
	if amatd.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", amatd.Type(), amatd.LocalName()))
	}
	return amatd.state
}

// ApiManagementApiTagDescriptionArgs contains the configurations for azurerm_api_management_api_tag_description.
type ApiManagementApiTagDescriptionArgs struct {
	// ApiTagId: string, required
	ApiTagId terra.StringValue `hcl:"api_tag_id,attr" validate:"required"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// ExternalDocumentationDescription: string, optional
	ExternalDocumentationDescription terra.StringValue `hcl:"external_documentation_description,attr"`
	// ExternalDocumentationUrl: string, optional
	ExternalDocumentationUrl terra.StringValue `hcl:"external_documentation_url,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Timeouts: optional
	Timeouts *apimanagementapitagdescription.Timeouts `hcl:"timeouts,block"`
}
type apiManagementApiTagDescriptionAttributes struct {
	ref terra.Reference
}

// ApiTagId returns a reference to field api_tag_id of azurerm_api_management_api_tag_description.
func (amatd apiManagementApiTagDescriptionAttributes) ApiTagId() terra.StringValue {
	return terra.ReferenceAsString(amatd.ref.Append("api_tag_id"))
}

// Description returns a reference to field description of azurerm_api_management_api_tag_description.
func (amatd apiManagementApiTagDescriptionAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(amatd.ref.Append("description"))
}

// ExternalDocumentationDescription returns a reference to field external_documentation_description of azurerm_api_management_api_tag_description.
func (amatd apiManagementApiTagDescriptionAttributes) ExternalDocumentationDescription() terra.StringValue {
	return terra.ReferenceAsString(amatd.ref.Append("external_documentation_description"))
}

// ExternalDocumentationUrl returns a reference to field external_documentation_url of azurerm_api_management_api_tag_description.
func (amatd apiManagementApiTagDescriptionAttributes) ExternalDocumentationUrl() terra.StringValue {
	return terra.ReferenceAsString(amatd.ref.Append("external_documentation_url"))
}

// Id returns a reference to field id of azurerm_api_management_api_tag_description.
func (amatd apiManagementApiTagDescriptionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(amatd.ref.Append("id"))
}

func (amatd apiManagementApiTagDescriptionAttributes) Timeouts() apimanagementapitagdescription.TimeoutsAttributes {
	return terra.ReferenceAsSingle[apimanagementapitagdescription.TimeoutsAttributes](amatd.ref.Append("timeouts"))
}

type apiManagementApiTagDescriptionState struct {
	ApiTagId                         string                                        `json:"api_tag_id"`
	Description                      string                                        `json:"description"`
	ExternalDocumentationDescription string                                        `json:"external_documentation_description"`
	ExternalDocumentationUrl         string                                        `json:"external_documentation_url"`
	Id                               string                                        `json:"id"`
	Timeouts                         *apimanagementapitagdescription.TimeoutsState `json:"timeouts"`
}
