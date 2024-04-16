// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_iap_brand

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

// Resource represents the Terraform resource google_iap_brand.
type Resource struct {
	Name      string
	Args      Args
	state     *googleIapBrandState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (gib *Resource) Type() string {
	return "google_iap_brand"
}

// LocalName returns the local name for [Resource].
func (gib *Resource) LocalName() string {
	return gib.Name
}

// Configuration returns the configuration (args) for [Resource].
func (gib *Resource) Configuration() interface{} {
	return gib.Args
}

// DependOn is used for other resources to depend on [Resource].
func (gib *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(gib)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (gib *Resource) Dependencies() terra.Dependencies {
	return gib.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (gib *Resource) LifecycleManagement() *terra.Lifecycle {
	return gib.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (gib *Resource) Attributes() googleIapBrandAttributes {
	return googleIapBrandAttributes{ref: terra.ReferenceResource(gib)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (gib *Resource) ImportState(state io.Reader) error {
	gib.state = &googleIapBrandState{}
	if err := json.NewDecoder(state).Decode(gib.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gib.Type(), gib.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (gib *Resource) State() (*googleIapBrandState, bool) {
	return gib.state, gib.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (gib *Resource) StateMust() *googleIapBrandState {
	if gib.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gib.Type(), gib.LocalName()))
	}
	return gib.state
}

// Args contains the configurations for google_iap_brand.
type Args struct {
	// ApplicationTitle: string, required
	ApplicationTitle terra.StringValue `hcl:"application_title,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// SupportEmail: string, required
	SupportEmail terra.StringValue `hcl:"support_email,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type googleIapBrandAttributes struct {
	ref terra.Reference
}

// ApplicationTitle returns a reference to field application_title of google_iap_brand.
func (gib googleIapBrandAttributes) ApplicationTitle() terra.StringValue {
	return terra.ReferenceAsString(gib.ref.Append("application_title"))
}

// Id returns a reference to field id of google_iap_brand.
func (gib googleIapBrandAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gib.ref.Append("id"))
}

// Name returns a reference to field name of google_iap_brand.
func (gib googleIapBrandAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(gib.ref.Append("name"))
}

// OrgInternalOnly returns a reference to field org_internal_only of google_iap_brand.
func (gib googleIapBrandAttributes) OrgInternalOnly() terra.BoolValue {
	return terra.ReferenceAsBool(gib.ref.Append("org_internal_only"))
}

// Project returns a reference to field project of google_iap_brand.
func (gib googleIapBrandAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(gib.ref.Append("project"))
}

// SupportEmail returns a reference to field support_email of google_iap_brand.
func (gib googleIapBrandAttributes) SupportEmail() terra.StringValue {
	return terra.ReferenceAsString(gib.ref.Append("support_email"))
}

func (gib googleIapBrandAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](gib.ref.Append("timeouts"))
}

type googleIapBrandState struct {
	ApplicationTitle string         `json:"application_title"`
	Id               string         `json:"id"`
	Name             string         `json:"name"`
	OrgInternalOnly  bool           `json:"org_internal_only"`
	Project          string         `json:"project"`
	SupportEmail     string         `json:"support_email"`
	Timeouts         *TimeoutsState `json:"timeouts"`
}
