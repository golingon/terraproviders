// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	apimanagementpolicy "github.com/golingon/terraproviders/azurerm/3.55.0/apimanagementpolicy"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewApiManagementPolicy creates a new instance of [ApiManagementPolicy].
func NewApiManagementPolicy(name string, args ApiManagementPolicyArgs) *ApiManagementPolicy {
	return &ApiManagementPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ApiManagementPolicy)(nil)

// ApiManagementPolicy represents the Terraform resource azurerm_api_management_policy.
type ApiManagementPolicy struct {
	Name      string
	Args      ApiManagementPolicyArgs
	state     *apiManagementPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ApiManagementPolicy].
func (amp *ApiManagementPolicy) Type() string {
	return "azurerm_api_management_policy"
}

// LocalName returns the local name for [ApiManagementPolicy].
func (amp *ApiManagementPolicy) LocalName() string {
	return amp.Name
}

// Configuration returns the configuration (args) for [ApiManagementPolicy].
func (amp *ApiManagementPolicy) Configuration() interface{} {
	return amp.Args
}

// DependOn is used for other resources to depend on [ApiManagementPolicy].
func (amp *ApiManagementPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(amp)
}

// Dependencies returns the list of resources [ApiManagementPolicy] depends_on.
func (amp *ApiManagementPolicy) Dependencies() terra.Dependencies {
	return amp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ApiManagementPolicy].
func (amp *ApiManagementPolicy) LifecycleManagement() *terra.Lifecycle {
	return amp.Lifecycle
}

// Attributes returns the attributes for [ApiManagementPolicy].
func (amp *ApiManagementPolicy) Attributes() apiManagementPolicyAttributes {
	return apiManagementPolicyAttributes{ref: terra.ReferenceResource(amp)}
}

// ImportState imports the given attribute values into [ApiManagementPolicy]'s state.
func (amp *ApiManagementPolicy) ImportState(av io.Reader) error {
	amp.state = &apiManagementPolicyState{}
	if err := json.NewDecoder(av).Decode(amp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", amp.Type(), amp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ApiManagementPolicy] has state.
func (amp *ApiManagementPolicy) State() (*apiManagementPolicyState, bool) {
	return amp.state, amp.state != nil
}

// StateMust returns the state for [ApiManagementPolicy]. Panics if the state is nil.
func (amp *ApiManagementPolicy) StateMust() *apiManagementPolicyState {
	if amp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", amp.Type(), amp.LocalName()))
	}
	return amp.state
}

// ApiManagementPolicyArgs contains the configurations for azurerm_api_management_policy.
type ApiManagementPolicyArgs struct {
	// ApiManagementId: string, required
	ApiManagementId terra.StringValue `hcl:"api_management_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// XmlContent: string, optional
	XmlContent terra.StringValue `hcl:"xml_content,attr"`
	// XmlLink: string, optional
	XmlLink terra.StringValue `hcl:"xml_link,attr"`
	// Timeouts: optional
	Timeouts *apimanagementpolicy.Timeouts `hcl:"timeouts,block"`
}
type apiManagementPolicyAttributes struct {
	ref terra.Reference
}

// ApiManagementId returns a reference to field api_management_id of azurerm_api_management_policy.
func (amp apiManagementPolicyAttributes) ApiManagementId() terra.StringValue {
	return terra.ReferenceAsString(amp.ref.Append("api_management_id"))
}

// Id returns a reference to field id of azurerm_api_management_policy.
func (amp apiManagementPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(amp.ref.Append("id"))
}

// XmlContent returns a reference to field xml_content of azurerm_api_management_policy.
func (amp apiManagementPolicyAttributes) XmlContent() terra.StringValue {
	return terra.ReferenceAsString(amp.ref.Append("xml_content"))
}

// XmlLink returns a reference to field xml_link of azurerm_api_management_policy.
func (amp apiManagementPolicyAttributes) XmlLink() terra.StringValue {
	return terra.ReferenceAsString(amp.ref.Append("xml_link"))
}

func (amp apiManagementPolicyAttributes) Timeouts() apimanagementpolicy.TimeoutsAttributes {
	return terra.ReferenceAsSingle[apimanagementpolicy.TimeoutsAttributes](amp.ref.Append("timeouts"))
}

type apiManagementPolicyState struct {
	ApiManagementId string                             `json:"api_management_id"`
	Id              string                             `json:"id"`
	XmlContent      string                             `json:"xml_content"`
	XmlLink         string                             `json:"xml_link"`
	Timeouts        *apimanagementpolicy.TimeoutsState `json:"timeouts"`
}
