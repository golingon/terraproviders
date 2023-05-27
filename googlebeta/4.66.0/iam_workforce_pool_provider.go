// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	iamworkforcepoolprovider "github.com/golingon/terraproviders/googlebeta/4.66.0/iamworkforcepoolprovider"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewIamWorkforcePoolProvider creates a new instance of [IamWorkforcePoolProvider].
func NewIamWorkforcePoolProvider(name string, args IamWorkforcePoolProviderArgs) *IamWorkforcePoolProvider {
	return &IamWorkforcePoolProvider{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*IamWorkforcePoolProvider)(nil)

// IamWorkforcePoolProvider represents the Terraform resource google_iam_workforce_pool_provider.
type IamWorkforcePoolProvider struct {
	Name      string
	Args      IamWorkforcePoolProviderArgs
	state     *iamWorkforcePoolProviderState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [IamWorkforcePoolProvider].
func (iwpp *IamWorkforcePoolProvider) Type() string {
	return "google_iam_workforce_pool_provider"
}

// LocalName returns the local name for [IamWorkforcePoolProvider].
func (iwpp *IamWorkforcePoolProvider) LocalName() string {
	return iwpp.Name
}

// Configuration returns the configuration (args) for [IamWorkforcePoolProvider].
func (iwpp *IamWorkforcePoolProvider) Configuration() interface{} {
	return iwpp.Args
}

// DependOn is used for other resources to depend on [IamWorkforcePoolProvider].
func (iwpp *IamWorkforcePoolProvider) DependOn() terra.Reference {
	return terra.ReferenceResource(iwpp)
}

// Dependencies returns the list of resources [IamWorkforcePoolProvider] depends_on.
func (iwpp *IamWorkforcePoolProvider) Dependencies() terra.Dependencies {
	return iwpp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [IamWorkforcePoolProvider].
func (iwpp *IamWorkforcePoolProvider) LifecycleManagement() *terra.Lifecycle {
	return iwpp.Lifecycle
}

// Attributes returns the attributes for [IamWorkforcePoolProvider].
func (iwpp *IamWorkforcePoolProvider) Attributes() iamWorkforcePoolProviderAttributes {
	return iamWorkforcePoolProviderAttributes{ref: terra.ReferenceResource(iwpp)}
}

// ImportState imports the given attribute values into [IamWorkforcePoolProvider]'s state.
func (iwpp *IamWorkforcePoolProvider) ImportState(av io.Reader) error {
	iwpp.state = &iamWorkforcePoolProviderState{}
	if err := json.NewDecoder(av).Decode(iwpp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", iwpp.Type(), iwpp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [IamWorkforcePoolProvider] has state.
func (iwpp *IamWorkforcePoolProvider) State() (*iamWorkforcePoolProviderState, bool) {
	return iwpp.state, iwpp.state != nil
}

// StateMust returns the state for [IamWorkforcePoolProvider]. Panics if the state is nil.
func (iwpp *IamWorkforcePoolProvider) StateMust() *iamWorkforcePoolProviderState {
	if iwpp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", iwpp.Type(), iwpp.LocalName()))
	}
	return iwpp.state
}

// IamWorkforcePoolProviderArgs contains the configurations for google_iam_workforce_pool_provider.
type IamWorkforcePoolProviderArgs struct {
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
	Oidc *iamworkforcepoolprovider.Oidc `hcl:"oidc,block"`
	// Saml: optional
	Saml *iamworkforcepoolprovider.Saml `hcl:"saml,block"`
	// Timeouts: optional
	Timeouts *iamworkforcepoolprovider.Timeouts `hcl:"timeouts,block"`
}
type iamWorkforcePoolProviderAttributes struct {
	ref terra.Reference
}

// AttributeCondition returns a reference to field attribute_condition of google_iam_workforce_pool_provider.
func (iwpp iamWorkforcePoolProviderAttributes) AttributeCondition() terra.StringValue {
	return terra.ReferenceAsString(iwpp.ref.Append("attribute_condition"))
}

// AttributeMapping returns a reference to field attribute_mapping of google_iam_workforce_pool_provider.
func (iwpp iamWorkforcePoolProviderAttributes) AttributeMapping() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](iwpp.ref.Append("attribute_mapping"))
}

// Description returns a reference to field description of google_iam_workforce_pool_provider.
func (iwpp iamWorkforcePoolProviderAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(iwpp.ref.Append("description"))
}

// Disabled returns a reference to field disabled of google_iam_workforce_pool_provider.
func (iwpp iamWorkforcePoolProviderAttributes) Disabled() terra.BoolValue {
	return terra.ReferenceAsBool(iwpp.ref.Append("disabled"))
}

// DisplayName returns a reference to field display_name of google_iam_workforce_pool_provider.
func (iwpp iamWorkforcePoolProviderAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(iwpp.ref.Append("display_name"))
}

// Id returns a reference to field id of google_iam_workforce_pool_provider.
func (iwpp iamWorkforcePoolProviderAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(iwpp.ref.Append("id"))
}

// Location returns a reference to field location of google_iam_workforce_pool_provider.
func (iwpp iamWorkforcePoolProviderAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(iwpp.ref.Append("location"))
}

// Name returns a reference to field name of google_iam_workforce_pool_provider.
func (iwpp iamWorkforcePoolProviderAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(iwpp.ref.Append("name"))
}

// ProviderId returns a reference to field provider_id of google_iam_workforce_pool_provider.
func (iwpp iamWorkforcePoolProviderAttributes) ProviderId() terra.StringValue {
	return terra.ReferenceAsString(iwpp.ref.Append("provider_id"))
}

// State returns a reference to field state of google_iam_workforce_pool_provider.
func (iwpp iamWorkforcePoolProviderAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(iwpp.ref.Append("state"))
}

// WorkforcePoolId returns a reference to field workforce_pool_id of google_iam_workforce_pool_provider.
func (iwpp iamWorkforcePoolProviderAttributes) WorkforcePoolId() terra.StringValue {
	return terra.ReferenceAsString(iwpp.ref.Append("workforce_pool_id"))
}

func (iwpp iamWorkforcePoolProviderAttributes) Oidc() terra.ListValue[iamworkforcepoolprovider.OidcAttributes] {
	return terra.ReferenceAsList[iamworkforcepoolprovider.OidcAttributes](iwpp.ref.Append("oidc"))
}

func (iwpp iamWorkforcePoolProviderAttributes) Saml() terra.ListValue[iamworkforcepoolprovider.SamlAttributes] {
	return terra.ReferenceAsList[iamworkforcepoolprovider.SamlAttributes](iwpp.ref.Append("saml"))
}

func (iwpp iamWorkforcePoolProviderAttributes) Timeouts() iamworkforcepoolprovider.TimeoutsAttributes {
	return terra.ReferenceAsSingle[iamworkforcepoolprovider.TimeoutsAttributes](iwpp.ref.Append("timeouts"))
}

type iamWorkforcePoolProviderState struct {
	AttributeCondition string                                  `json:"attribute_condition"`
	AttributeMapping   map[string]string                       `json:"attribute_mapping"`
	Description        string                                  `json:"description"`
	Disabled           bool                                    `json:"disabled"`
	DisplayName        string                                  `json:"display_name"`
	Id                 string                                  `json:"id"`
	Location           string                                  `json:"location"`
	Name               string                                  `json:"name"`
	ProviderId         string                                  `json:"provider_id"`
	State              string                                  `json:"state"`
	WorkforcePoolId    string                                  `json:"workforce_pool_id"`
	Oidc               []iamworkforcepoolprovider.OidcState    `json:"oidc"`
	Saml               []iamworkforcepoolprovider.SamlState    `json:"saml"`
	Timeouts           *iamworkforcepoolprovider.TimeoutsState `json:"timeouts"`
}
