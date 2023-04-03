// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	essentialcontactscontact "github.com/golingon/terraproviders/google/4.59.0/essentialcontactscontact"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewEssentialContactsContact creates a new instance of [EssentialContactsContact].
func NewEssentialContactsContact(name string, args EssentialContactsContactArgs) *EssentialContactsContact {
	return &EssentialContactsContact{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*EssentialContactsContact)(nil)

// EssentialContactsContact represents the Terraform resource google_essential_contacts_contact.
type EssentialContactsContact struct {
	Name      string
	Args      EssentialContactsContactArgs
	state     *essentialContactsContactState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [EssentialContactsContact].
func (ecc *EssentialContactsContact) Type() string {
	return "google_essential_contacts_contact"
}

// LocalName returns the local name for [EssentialContactsContact].
func (ecc *EssentialContactsContact) LocalName() string {
	return ecc.Name
}

// Configuration returns the configuration (args) for [EssentialContactsContact].
func (ecc *EssentialContactsContact) Configuration() interface{} {
	return ecc.Args
}

// DependOn is used for other resources to depend on [EssentialContactsContact].
func (ecc *EssentialContactsContact) DependOn() terra.Reference {
	return terra.ReferenceResource(ecc)
}

// Dependencies returns the list of resources [EssentialContactsContact] depends_on.
func (ecc *EssentialContactsContact) Dependencies() terra.Dependencies {
	return ecc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [EssentialContactsContact].
func (ecc *EssentialContactsContact) LifecycleManagement() *terra.Lifecycle {
	return ecc.Lifecycle
}

// Attributes returns the attributes for [EssentialContactsContact].
func (ecc *EssentialContactsContact) Attributes() essentialContactsContactAttributes {
	return essentialContactsContactAttributes{ref: terra.ReferenceResource(ecc)}
}

// ImportState imports the given attribute values into [EssentialContactsContact]'s state.
func (ecc *EssentialContactsContact) ImportState(av io.Reader) error {
	ecc.state = &essentialContactsContactState{}
	if err := json.NewDecoder(av).Decode(ecc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ecc.Type(), ecc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [EssentialContactsContact] has state.
func (ecc *EssentialContactsContact) State() (*essentialContactsContactState, bool) {
	return ecc.state, ecc.state != nil
}

// StateMust returns the state for [EssentialContactsContact]. Panics if the state is nil.
func (ecc *EssentialContactsContact) StateMust() *essentialContactsContactState {
	if ecc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ecc.Type(), ecc.LocalName()))
	}
	return ecc.state
}

// EssentialContactsContactArgs contains the configurations for google_essential_contacts_contact.
type EssentialContactsContactArgs struct {
	// Email: string, required
	Email terra.StringValue `hcl:"email,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LanguageTag: string, required
	LanguageTag terra.StringValue `hcl:"language_tag,attr" validate:"required"`
	// NotificationCategorySubscriptions: list of string, required
	NotificationCategorySubscriptions terra.ListValue[terra.StringValue] `hcl:"notification_category_subscriptions,attr" validate:"required"`
	// Parent: string, required
	Parent terra.StringValue `hcl:"parent,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *essentialcontactscontact.Timeouts `hcl:"timeouts,block"`
}
type essentialContactsContactAttributes struct {
	ref terra.Reference
}

// Email returns a reference to field email of google_essential_contacts_contact.
func (ecc essentialContactsContactAttributes) Email() terra.StringValue {
	return terra.ReferenceAsString(ecc.ref.Append("email"))
}

// Id returns a reference to field id of google_essential_contacts_contact.
func (ecc essentialContactsContactAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ecc.ref.Append("id"))
}

// LanguageTag returns a reference to field language_tag of google_essential_contacts_contact.
func (ecc essentialContactsContactAttributes) LanguageTag() terra.StringValue {
	return terra.ReferenceAsString(ecc.ref.Append("language_tag"))
}

// Name returns a reference to field name of google_essential_contacts_contact.
func (ecc essentialContactsContactAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ecc.ref.Append("name"))
}

// NotificationCategorySubscriptions returns a reference to field notification_category_subscriptions of google_essential_contacts_contact.
func (ecc essentialContactsContactAttributes) NotificationCategorySubscriptions() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ecc.ref.Append("notification_category_subscriptions"))
}

// Parent returns a reference to field parent of google_essential_contacts_contact.
func (ecc essentialContactsContactAttributes) Parent() terra.StringValue {
	return terra.ReferenceAsString(ecc.ref.Append("parent"))
}

func (ecc essentialContactsContactAttributes) Timeouts() essentialcontactscontact.TimeoutsAttributes {
	return terra.ReferenceAsSingle[essentialcontactscontact.TimeoutsAttributes](ecc.ref.Append("timeouts"))
}

type essentialContactsContactState struct {
	Email                             string                                  `json:"email"`
	Id                                string                                  `json:"id"`
	LanguageTag                       string                                  `json:"language_tag"`
	Name                              string                                  `json:"name"`
	NotificationCategorySubscriptions []string                                `json:"notification_category_subscriptions"`
	Parent                            string                                  `json:"parent"`
	Timeouts                          *essentialcontactscontact.TimeoutsState `json:"timeouts"`
}
