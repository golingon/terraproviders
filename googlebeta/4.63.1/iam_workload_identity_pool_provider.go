// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	iamworkloadidentitypoolprovider "github.com/golingon/terraproviders/googlebeta/4.63.1/iamworkloadidentitypoolprovider"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewIamWorkloadIdentityPoolProvider creates a new instance of [IamWorkloadIdentityPoolProvider].
func NewIamWorkloadIdentityPoolProvider(name string, args IamWorkloadIdentityPoolProviderArgs) *IamWorkloadIdentityPoolProvider {
	return &IamWorkloadIdentityPoolProvider{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*IamWorkloadIdentityPoolProvider)(nil)

// IamWorkloadIdentityPoolProvider represents the Terraform resource google_iam_workload_identity_pool_provider.
type IamWorkloadIdentityPoolProvider struct {
	Name      string
	Args      IamWorkloadIdentityPoolProviderArgs
	state     *iamWorkloadIdentityPoolProviderState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [IamWorkloadIdentityPoolProvider].
func (iwipp *IamWorkloadIdentityPoolProvider) Type() string {
	return "google_iam_workload_identity_pool_provider"
}

// LocalName returns the local name for [IamWorkloadIdentityPoolProvider].
func (iwipp *IamWorkloadIdentityPoolProvider) LocalName() string {
	return iwipp.Name
}

// Configuration returns the configuration (args) for [IamWorkloadIdentityPoolProvider].
func (iwipp *IamWorkloadIdentityPoolProvider) Configuration() interface{} {
	return iwipp.Args
}

// DependOn is used for other resources to depend on [IamWorkloadIdentityPoolProvider].
func (iwipp *IamWorkloadIdentityPoolProvider) DependOn() terra.Reference {
	return terra.ReferenceResource(iwipp)
}

// Dependencies returns the list of resources [IamWorkloadIdentityPoolProvider] depends_on.
func (iwipp *IamWorkloadIdentityPoolProvider) Dependencies() terra.Dependencies {
	return iwipp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [IamWorkloadIdentityPoolProvider].
func (iwipp *IamWorkloadIdentityPoolProvider) LifecycleManagement() *terra.Lifecycle {
	return iwipp.Lifecycle
}

// Attributes returns the attributes for [IamWorkloadIdentityPoolProvider].
func (iwipp *IamWorkloadIdentityPoolProvider) Attributes() iamWorkloadIdentityPoolProviderAttributes {
	return iamWorkloadIdentityPoolProviderAttributes{ref: terra.ReferenceResource(iwipp)}
}

// ImportState imports the given attribute values into [IamWorkloadIdentityPoolProvider]'s state.
func (iwipp *IamWorkloadIdentityPoolProvider) ImportState(av io.Reader) error {
	iwipp.state = &iamWorkloadIdentityPoolProviderState{}
	if err := json.NewDecoder(av).Decode(iwipp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", iwipp.Type(), iwipp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [IamWorkloadIdentityPoolProvider] has state.
func (iwipp *IamWorkloadIdentityPoolProvider) State() (*iamWorkloadIdentityPoolProviderState, bool) {
	return iwipp.state, iwipp.state != nil
}

// StateMust returns the state for [IamWorkloadIdentityPoolProvider]. Panics if the state is nil.
func (iwipp *IamWorkloadIdentityPoolProvider) StateMust() *iamWorkloadIdentityPoolProviderState {
	if iwipp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", iwipp.Type(), iwipp.LocalName()))
	}
	return iwipp.state
}

// IamWorkloadIdentityPoolProviderArgs contains the configurations for google_iam_workload_identity_pool_provider.
type IamWorkloadIdentityPoolProviderArgs struct {
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
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// WorkloadIdentityPoolId: string, required
	WorkloadIdentityPoolId terra.StringValue `hcl:"workload_identity_pool_id,attr" validate:"required"`
	// WorkloadIdentityPoolProviderId: string, required
	WorkloadIdentityPoolProviderId terra.StringValue `hcl:"workload_identity_pool_provider_id,attr" validate:"required"`
	// Aws: optional
	Aws *iamworkloadidentitypoolprovider.Aws `hcl:"aws,block"`
	// Oidc: optional
	Oidc *iamworkloadidentitypoolprovider.Oidc `hcl:"oidc,block"`
	// Timeouts: optional
	Timeouts *iamworkloadidentitypoolprovider.Timeouts `hcl:"timeouts,block"`
}
type iamWorkloadIdentityPoolProviderAttributes struct {
	ref terra.Reference
}

// AttributeCondition returns a reference to field attribute_condition of google_iam_workload_identity_pool_provider.
func (iwipp iamWorkloadIdentityPoolProviderAttributes) AttributeCondition() terra.StringValue {
	return terra.ReferenceAsString(iwipp.ref.Append("attribute_condition"))
}

// AttributeMapping returns a reference to field attribute_mapping of google_iam_workload_identity_pool_provider.
func (iwipp iamWorkloadIdentityPoolProviderAttributes) AttributeMapping() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](iwipp.ref.Append("attribute_mapping"))
}

// Description returns a reference to field description of google_iam_workload_identity_pool_provider.
func (iwipp iamWorkloadIdentityPoolProviderAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(iwipp.ref.Append("description"))
}

// Disabled returns a reference to field disabled of google_iam_workload_identity_pool_provider.
func (iwipp iamWorkloadIdentityPoolProviderAttributes) Disabled() terra.BoolValue {
	return terra.ReferenceAsBool(iwipp.ref.Append("disabled"))
}

// DisplayName returns a reference to field display_name of google_iam_workload_identity_pool_provider.
func (iwipp iamWorkloadIdentityPoolProviderAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(iwipp.ref.Append("display_name"))
}

// Id returns a reference to field id of google_iam_workload_identity_pool_provider.
func (iwipp iamWorkloadIdentityPoolProviderAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(iwipp.ref.Append("id"))
}

// Name returns a reference to field name of google_iam_workload_identity_pool_provider.
func (iwipp iamWorkloadIdentityPoolProviderAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(iwipp.ref.Append("name"))
}

// Project returns a reference to field project of google_iam_workload_identity_pool_provider.
func (iwipp iamWorkloadIdentityPoolProviderAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(iwipp.ref.Append("project"))
}

// State returns a reference to field state of google_iam_workload_identity_pool_provider.
func (iwipp iamWorkloadIdentityPoolProviderAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(iwipp.ref.Append("state"))
}

// WorkloadIdentityPoolId returns a reference to field workload_identity_pool_id of google_iam_workload_identity_pool_provider.
func (iwipp iamWorkloadIdentityPoolProviderAttributes) WorkloadIdentityPoolId() terra.StringValue {
	return terra.ReferenceAsString(iwipp.ref.Append("workload_identity_pool_id"))
}

// WorkloadIdentityPoolProviderId returns a reference to field workload_identity_pool_provider_id of google_iam_workload_identity_pool_provider.
func (iwipp iamWorkloadIdentityPoolProviderAttributes) WorkloadIdentityPoolProviderId() terra.StringValue {
	return terra.ReferenceAsString(iwipp.ref.Append("workload_identity_pool_provider_id"))
}

func (iwipp iamWorkloadIdentityPoolProviderAttributes) Aws() terra.ListValue[iamworkloadidentitypoolprovider.AwsAttributes] {
	return terra.ReferenceAsList[iamworkloadidentitypoolprovider.AwsAttributes](iwipp.ref.Append("aws"))
}

func (iwipp iamWorkloadIdentityPoolProviderAttributes) Oidc() terra.ListValue[iamworkloadidentitypoolprovider.OidcAttributes] {
	return terra.ReferenceAsList[iamworkloadidentitypoolprovider.OidcAttributes](iwipp.ref.Append("oidc"))
}

func (iwipp iamWorkloadIdentityPoolProviderAttributes) Timeouts() iamworkloadidentitypoolprovider.TimeoutsAttributes {
	return terra.ReferenceAsSingle[iamworkloadidentitypoolprovider.TimeoutsAttributes](iwipp.ref.Append("timeouts"))
}

type iamWorkloadIdentityPoolProviderState struct {
	AttributeCondition             string                                         `json:"attribute_condition"`
	AttributeMapping               map[string]string                              `json:"attribute_mapping"`
	Description                    string                                         `json:"description"`
	Disabled                       bool                                           `json:"disabled"`
	DisplayName                    string                                         `json:"display_name"`
	Id                             string                                         `json:"id"`
	Name                           string                                         `json:"name"`
	Project                        string                                         `json:"project"`
	State                          string                                         `json:"state"`
	WorkloadIdentityPoolId         string                                         `json:"workload_identity_pool_id"`
	WorkloadIdentityPoolProviderId string                                         `json:"workload_identity_pool_provider_id"`
	Aws                            []iamworkloadidentitypoolprovider.AwsState     `json:"aws"`
	Oidc                           []iamworkloadidentitypoolprovider.OidcState    `json:"oidc"`
	Timeouts                       *iamworkloadidentitypoolprovider.TimeoutsState `json:"timeouts"`
}
