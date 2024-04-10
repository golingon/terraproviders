// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azuread

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	accesspackageassignmentpolicy "github.com/golingon/terraproviders/azuread/2.47.0/accesspackageassignmentpolicy"
	"io"
)

// NewAccessPackageAssignmentPolicy creates a new instance of [AccessPackageAssignmentPolicy].
func NewAccessPackageAssignmentPolicy(name string, args AccessPackageAssignmentPolicyArgs) *AccessPackageAssignmentPolicy {
	return &AccessPackageAssignmentPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AccessPackageAssignmentPolicy)(nil)

// AccessPackageAssignmentPolicy represents the Terraform resource azuread_access_package_assignment_policy.
type AccessPackageAssignmentPolicy struct {
	Name      string
	Args      AccessPackageAssignmentPolicyArgs
	state     *accessPackageAssignmentPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [AccessPackageAssignmentPolicy].
func (apap *AccessPackageAssignmentPolicy) Type() string {
	return "azuread_access_package_assignment_policy"
}

// LocalName returns the local name for [AccessPackageAssignmentPolicy].
func (apap *AccessPackageAssignmentPolicy) LocalName() string {
	return apap.Name
}

// Configuration returns the configuration (args) for [AccessPackageAssignmentPolicy].
func (apap *AccessPackageAssignmentPolicy) Configuration() interface{} {
	return apap.Args
}

// DependOn is used for other resources to depend on [AccessPackageAssignmentPolicy].
func (apap *AccessPackageAssignmentPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(apap)
}

// Dependencies returns the list of resources [AccessPackageAssignmentPolicy] depends_on.
func (apap *AccessPackageAssignmentPolicy) Dependencies() terra.Dependencies {
	return apap.DependsOn
}

// LifecycleManagement returns the lifecycle block for [AccessPackageAssignmentPolicy].
func (apap *AccessPackageAssignmentPolicy) LifecycleManagement() *terra.Lifecycle {
	return apap.Lifecycle
}

// Attributes returns the attributes for [AccessPackageAssignmentPolicy].
func (apap *AccessPackageAssignmentPolicy) Attributes() accessPackageAssignmentPolicyAttributes {
	return accessPackageAssignmentPolicyAttributes{ref: terra.ReferenceResource(apap)}
}

// ImportState imports the given attribute values into [AccessPackageAssignmentPolicy]'s state.
func (apap *AccessPackageAssignmentPolicy) ImportState(av io.Reader) error {
	apap.state = &accessPackageAssignmentPolicyState{}
	if err := json.NewDecoder(av).Decode(apap.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", apap.Type(), apap.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [AccessPackageAssignmentPolicy] has state.
func (apap *AccessPackageAssignmentPolicy) State() (*accessPackageAssignmentPolicyState, bool) {
	return apap.state, apap.state != nil
}

// StateMust returns the state for [AccessPackageAssignmentPolicy]. Panics if the state is nil.
func (apap *AccessPackageAssignmentPolicy) StateMust() *accessPackageAssignmentPolicyState {
	if apap.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", apap.Type(), apap.LocalName()))
	}
	return apap.state
}

// AccessPackageAssignmentPolicyArgs contains the configurations for azuread_access_package_assignment_policy.
type AccessPackageAssignmentPolicyArgs struct {
	// AccessPackageId: string, required
	AccessPackageId terra.StringValue `hcl:"access_package_id,attr" validate:"required"`
	// Description: string, required
	Description terra.StringValue `hcl:"description,attr" validate:"required"`
	// DisplayName: string, required
	DisplayName terra.StringValue `hcl:"display_name,attr" validate:"required"`
	// DurationInDays: number, optional
	DurationInDays terra.NumberValue `hcl:"duration_in_days,attr"`
	// ExpirationDate: string, optional
	ExpirationDate terra.StringValue `hcl:"expiration_date,attr"`
	// ExtensionEnabled: bool, optional
	ExtensionEnabled terra.BoolValue `hcl:"extension_enabled,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ApprovalSettings: optional
	ApprovalSettings *accesspackageassignmentpolicy.ApprovalSettings `hcl:"approval_settings,block"`
	// AssignmentReviewSettings: optional
	AssignmentReviewSettings *accesspackageassignmentpolicy.AssignmentReviewSettings `hcl:"assignment_review_settings,block"`
	// Question: min=0
	Question []accesspackageassignmentpolicy.Question `hcl:"question,block" validate:"min=0"`
	// RequestorSettings: optional
	RequestorSettings *accesspackageassignmentpolicy.RequestorSettings `hcl:"requestor_settings,block"`
	// Timeouts: optional
	Timeouts *accesspackageassignmentpolicy.Timeouts `hcl:"timeouts,block"`
}
type accessPackageAssignmentPolicyAttributes struct {
	ref terra.Reference
}

// AccessPackageId returns a reference to field access_package_id of azuread_access_package_assignment_policy.
func (apap accessPackageAssignmentPolicyAttributes) AccessPackageId() terra.StringValue {
	return terra.ReferenceAsString(apap.ref.Append("access_package_id"))
}

// Description returns a reference to field description of azuread_access_package_assignment_policy.
func (apap accessPackageAssignmentPolicyAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(apap.ref.Append("description"))
}

// DisplayName returns a reference to field display_name of azuread_access_package_assignment_policy.
func (apap accessPackageAssignmentPolicyAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(apap.ref.Append("display_name"))
}

// DurationInDays returns a reference to field duration_in_days of azuread_access_package_assignment_policy.
func (apap accessPackageAssignmentPolicyAttributes) DurationInDays() terra.NumberValue {
	return terra.ReferenceAsNumber(apap.ref.Append("duration_in_days"))
}

// ExpirationDate returns a reference to field expiration_date of azuread_access_package_assignment_policy.
func (apap accessPackageAssignmentPolicyAttributes) ExpirationDate() terra.StringValue {
	return terra.ReferenceAsString(apap.ref.Append("expiration_date"))
}

// ExtensionEnabled returns a reference to field extension_enabled of azuread_access_package_assignment_policy.
func (apap accessPackageAssignmentPolicyAttributes) ExtensionEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(apap.ref.Append("extension_enabled"))
}

// Id returns a reference to field id of azuread_access_package_assignment_policy.
func (apap accessPackageAssignmentPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(apap.ref.Append("id"))
}

func (apap accessPackageAssignmentPolicyAttributes) ApprovalSettings() terra.ListValue[accesspackageassignmentpolicy.ApprovalSettingsAttributes] {
	return terra.ReferenceAsList[accesspackageassignmentpolicy.ApprovalSettingsAttributes](apap.ref.Append("approval_settings"))
}

func (apap accessPackageAssignmentPolicyAttributes) AssignmentReviewSettings() terra.ListValue[accesspackageassignmentpolicy.AssignmentReviewSettingsAttributes] {
	return terra.ReferenceAsList[accesspackageassignmentpolicy.AssignmentReviewSettingsAttributes](apap.ref.Append("assignment_review_settings"))
}

func (apap accessPackageAssignmentPolicyAttributes) Question() terra.ListValue[accesspackageassignmentpolicy.QuestionAttributes] {
	return terra.ReferenceAsList[accesspackageassignmentpolicy.QuestionAttributes](apap.ref.Append("question"))
}

func (apap accessPackageAssignmentPolicyAttributes) RequestorSettings() terra.ListValue[accesspackageassignmentpolicy.RequestorSettingsAttributes] {
	return terra.ReferenceAsList[accesspackageassignmentpolicy.RequestorSettingsAttributes](apap.ref.Append("requestor_settings"))
}

func (apap accessPackageAssignmentPolicyAttributes) Timeouts() accesspackageassignmentpolicy.TimeoutsAttributes {
	return terra.ReferenceAsSingle[accesspackageassignmentpolicy.TimeoutsAttributes](apap.ref.Append("timeouts"))
}

type accessPackageAssignmentPolicyState struct {
	AccessPackageId          string                                                        `json:"access_package_id"`
	Description              string                                                        `json:"description"`
	DisplayName              string                                                        `json:"display_name"`
	DurationInDays           float64                                                       `json:"duration_in_days"`
	ExpirationDate           string                                                        `json:"expiration_date"`
	ExtensionEnabled         bool                                                          `json:"extension_enabled"`
	Id                       string                                                        `json:"id"`
	ApprovalSettings         []accesspackageassignmentpolicy.ApprovalSettingsState         `json:"approval_settings"`
	AssignmentReviewSettings []accesspackageassignmentpolicy.AssignmentReviewSettingsState `json:"assignment_review_settings"`
	Question                 []accesspackageassignmentpolicy.QuestionState                 `json:"question"`
	RequestorSettings        []accesspackageassignmentpolicy.RequestorSettingsState        `json:"requestor_settings"`
	Timeouts                 *accesspackageassignmentpolicy.TimeoutsState                  `json:"timeouts"`
}
