// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_iam_workforce_pool_provider

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

// Resource represents the Terraform resource google_iam_workforce_pool_provider.
type Resource struct {
	Name      string
	Args      Args
	state     *googleIamWorkforcePoolProviderState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (giwpp *Resource) Type() string {
	return "google_iam_workforce_pool_provider"
}

// LocalName returns the local name for [Resource].
func (giwpp *Resource) LocalName() string {
	return giwpp.Name
}

// Configuration returns the configuration (args) for [Resource].
func (giwpp *Resource) Configuration() interface{} {
	return giwpp.Args
}

// DependOn is used for other resources to depend on [Resource].
func (giwpp *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(giwpp)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (giwpp *Resource) Dependencies() terra.Dependencies {
	return giwpp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (giwpp *Resource) LifecycleManagement() *terra.Lifecycle {
	return giwpp.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (giwpp *Resource) Attributes() googleIamWorkforcePoolProviderAttributes {
	return googleIamWorkforcePoolProviderAttributes{ref: terra.ReferenceResource(giwpp)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (giwpp *Resource) ImportState(state io.Reader) error {
	giwpp.state = &googleIamWorkforcePoolProviderState{}
	if err := json.NewDecoder(state).Decode(giwpp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", giwpp.Type(), giwpp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (giwpp *Resource) State() (*googleIamWorkforcePoolProviderState, bool) {
	return giwpp.state, giwpp.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (giwpp *Resource) StateMust() *googleIamWorkforcePoolProviderState {
	if giwpp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", giwpp.Type(), giwpp.LocalName()))
	}
	return giwpp.state
}

// Args contains the configurations for google_iam_workforce_pool_provider.
type Args struct {
	// AttributeCondition: string, optional
	AttributeCondition terra.StringValue `hcl:"attribute_condition,attr"`
	// AttributeMapping: map of string, optional
	AttributeMapping terra.MapValue[terra.StringValue] `hcl:"attribute_mapping,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Disabled: bool, optional
	Disabled terra.BoolValue `hcl:"disabled,attr"`
	// DisplayName: string, optional
	DisplayName terra.StringValue `hcl:"display_name,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// ProviderId: string, required
	ProviderId terra.StringValue `hcl:"provider_id,attr" validate:"required"`
	// WorkforcePoolId: string, required
	WorkforcePoolId terra.StringValue `hcl:"workforce_pool_id,attr" validate:"required"`
	// Oidc: optional
	Oidc *Oidc `hcl:"oidc,block"`
	// Saml: optional
	Saml *Saml `hcl:"saml,block"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type googleIamWorkforcePoolProviderAttributes struct {
	ref terra.Reference
}

// AttributeCondition returns a reference to field attribute_condition of google_iam_workforce_pool_provider.
func (giwpp googleIamWorkforcePoolProviderAttributes) AttributeCondition() terra.StringValue {
	return terra.ReferenceAsString(giwpp.ref.Append("attribute_condition"))
}

// AttributeMapping returns a reference to field attribute_mapping of google_iam_workforce_pool_provider.
func (giwpp googleIamWorkforcePoolProviderAttributes) AttributeMapping() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](giwpp.ref.Append("attribute_mapping"))
}

// Description returns a reference to field description of google_iam_workforce_pool_provider.
func (giwpp googleIamWorkforcePoolProviderAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(giwpp.ref.Append("description"))
}

// Disabled returns a reference to field disabled of google_iam_workforce_pool_provider.
func (giwpp googleIamWorkforcePoolProviderAttributes) Disabled() terra.BoolValue {
	return terra.ReferenceAsBool(giwpp.ref.Append("disabled"))
}

// DisplayName returns a reference to field display_name of google_iam_workforce_pool_provider.
func (giwpp googleIamWorkforcePoolProviderAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(giwpp.ref.Append("display_name"))
}

// Id returns a reference to field id of google_iam_workforce_pool_provider.
func (giwpp googleIamWorkforcePoolProviderAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(giwpp.ref.Append("id"))
}

// Location returns a reference to field location of google_iam_workforce_pool_provider.
func (giwpp googleIamWorkforcePoolProviderAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(giwpp.ref.Append("location"))
}

// Name returns a reference to field name of google_iam_workforce_pool_provider.
func (giwpp googleIamWorkforcePoolProviderAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(giwpp.ref.Append("name"))
}

// ProviderId returns a reference to field provider_id of google_iam_workforce_pool_provider.
func (giwpp googleIamWorkforcePoolProviderAttributes) ProviderId() terra.StringValue {
	return terra.ReferenceAsString(giwpp.ref.Append("provider_id"))
}

// State returns a reference to field state of google_iam_workforce_pool_provider.
func (giwpp googleIamWorkforcePoolProviderAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(giwpp.ref.Append("state"))
}

// WorkforcePoolId returns a reference to field workforce_pool_id of google_iam_workforce_pool_provider.
func (giwpp googleIamWorkforcePoolProviderAttributes) WorkforcePoolId() terra.StringValue {
	return terra.ReferenceAsString(giwpp.ref.Append("workforce_pool_id"))
}

func (giwpp googleIamWorkforcePoolProviderAttributes) Oidc() terra.ListValue[OidcAttributes] {
	return terra.ReferenceAsList[OidcAttributes](giwpp.ref.Append("oidc"))
}

func (giwpp googleIamWorkforcePoolProviderAttributes) Saml() terra.ListValue[SamlAttributes] {
	return terra.ReferenceAsList[SamlAttributes](giwpp.ref.Append("saml"))
}

func (giwpp googleIamWorkforcePoolProviderAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](giwpp.ref.Append("timeouts"))
}

type googleIamWorkforcePoolProviderState struct {
	AttributeCondition string            `json:"attribute_condition"`
	AttributeMapping   map[string]string `json:"attribute_mapping"`
	Description        string            `json:"description"`
	Disabled           bool              `json:"disabled"`
	DisplayName        string            `json:"display_name"`
	Id                 string            `json:"id"`
	Location           string            `json:"location"`
	Name               string            `json:"name"`
	ProviderId         string            `json:"provider_id"`
	State              string            `json:"state"`
	WorkforcePoolId    string            `json:"workforce_pool_id"`
	Oidc               []OidcState       `json:"oidc"`
	Saml               []SamlState       `json:"saml"`
	Timeouts           *TimeoutsState    `json:"timeouts"`
}
