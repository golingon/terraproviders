// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	apigeeorganization "github.com/golingon/terraproviders/google/4.59.0/apigeeorganization"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewApigeeOrganization creates a new instance of [ApigeeOrganization].
func NewApigeeOrganization(name string, args ApigeeOrganizationArgs) *ApigeeOrganization {
	return &ApigeeOrganization{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ApigeeOrganization)(nil)

// ApigeeOrganization represents the Terraform resource google_apigee_organization.
type ApigeeOrganization struct {
	Name      string
	Args      ApigeeOrganizationArgs
	state     *apigeeOrganizationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ApigeeOrganization].
func (ao *ApigeeOrganization) Type() string {
	return "google_apigee_organization"
}

// LocalName returns the local name for [ApigeeOrganization].
func (ao *ApigeeOrganization) LocalName() string {
	return ao.Name
}

// Configuration returns the configuration (args) for [ApigeeOrganization].
func (ao *ApigeeOrganization) Configuration() interface{} {
	return ao.Args
}

// DependOn is used for other resources to depend on [ApigeeOrganization].
func (ao *ApigeeOrganization) DependOn() terra.Reference {
	return terra.ReferenceResource(ao)
}

// Dependencies returns the list of resources [ApigeeOrganization] depends_on.
func (ao *ApigeeOrganization) Dependencies() terra.Dependencies {
	return ao.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ApigeeOrganization].
func (ao *ApigeeOrganization) LifecycleManagement() *terra.Lifecycle {
	return ao.Lifecycle
}

// Attributes returns the attributes for [ApigeeOrganization].
func (ao *ApigeeOrganization) Attributes() apigeeOrganizationAttributes {
	return apigeeOrganizationAttributes{ref: terra.ReferenceResource(ao)}
}

// ImportState imports the given attribute values into [ApigeeOrganization]'s state.
func (ao *ApigeeOrganization) ImportState(av io.Reader) error {
	ao.state = &apigeeOrganizationState{}
	if err := json.NewDecoder(av).Decode(ao.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ao.Type(), ao.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ApigeeOrganization] has state.
func (ao *ApigeeOrganization) State() (*apigeeOrganizationState, bool) {
	return ao.state, ao.state != nil
}

// StateMust returns the state for [ApigeeOrganization]. Panics if the state is nil.
func (ao *ApigeeOrganization) StateMust() *apigeeOrganizationState {
	if ao.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ao.Type(), ao.LocalName()))
	}
	return ao.state
}

// ApigeeOrganizationArgs contains the configurations for google_apigee_organization.
type ApigeeOrganizationArgs struct {
	// AnalyticsRegion: string, optional
	AnalyticsRegion terra.StringValue `hcl:"analytics_region,attr"`
	// AuthorizedNetwork: string, optional
	AuthorizedNetwork terra.StringValue `hcl:"authorized_network,attr"`
	// BillingType: string, optional
	BillingType terra.StringValue `hcl:"billing_type,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// DisplayName: string, optional
	DisplayName terra.StringValue `hcl:"display_name,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ProjectId: string, required
	ProjectId terra.StringValue `hcl:"project_id,attr" validate:"required"`
	// Retention: string, optional
	Retention terra.StringValue `hcl:"retention,attr"`
	// RuntimeDatabaseEncryptionKeyName: string, optional
	RuntimeDatabaseEncryptionKeyName terra.StringValue `hcl:"runtime_database_encryption_key_name,attr"`
	// RuntimeType: string, optional
	RuntimeType terra.StringValue `hcl:"runtime_type,attr"`
	// Properties: optional
	Properties *apigeeorganization.Properties `hcl:"properties,block"`
	// Timeouts: optional
	Timeouts *apigeeorganization.Timeouts `hcl:"timeouts,block"`
}
type apigeeOrganizationAttributes struct {
	ref terra.Reference
}

// AnalyticsRegion returns a reference to field analytics_region of google_apigee_organization.
func (ao apigeeOrganizationAttributes) AnalyticsRegion() terra.StringValue {
	return terra.ReferenceAsString(ao.ref.Append("analytics_region"))
}

// AuthorizedNetwork returns a reference to field authorized_network of google_apigee_organization.
func (ao apigeeOrganizationAttributes) AuthorizedNetwork() terra.StringValue {
	return terra.ReferenceAsString(ao.ref.Append("authorized_network"))
}

// BillingType returns a reference to field billing_type of google_apigee_organization.
func (ao apigeeOrganizationAttributes) BillingType() terra.StringValue {
	return terra.ReferenceAsString(ao.ref.Append("billing_type"))
}

// CaCertificate returns a reference to field ca_certificate of google_apigee_organization.
func (ao apigeeOrganizationAttributes) CaCertificate() terra.StringValue {
	return terra.ReferenceAsString(ao.ref.Append("ca_certificate"))
}

// Description returns a reference to field description of google_apigee_organization.
func (ao apigeeOrganizationAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(ao.ref.Append("description"))
}

// DisplayName returns a reference to field display_name of google_apigee_organization.
func (ao apigeeOrganizationAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(ao.ref.Append("display_name"))
}

// Id returns a reference to field id of google_apigee_organization.
func (ao apigeeOrganizationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ao.ref.Append("id"))
}

// Name returns a reference to field name of google_apigee_organization.
func (ao apigeeOrganizationAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ao.ref.Append("name"))
}

// ProjectId returns a reference to field project_id of google_apigee_organization.
func (ao apigeeOrganizationAttributes) ProjectId() terra.StringValue {
	return terra.ReferenceAsString(ao.ref.Append("project_id"))
}

// Retention returns a reference to field retention of google_apigee_organization.
func (ao apigeeOrganizationAttributes) Retention() terra.StringValue {
	return terra.ReferenceAsString(ao.ref.Append("retention"))
}

// RuntimeDatabaseEncryptionKeyName returns a reference to field runtime_database_encryption_key_name of google_apigee_organization.
func (ao apigeeOrganizationAttributes) RuntimeDatabaseEncryptionKeyName() terra.StringValue {
	return terra.ReferenceAsString(ao.ref.Append("runtime_database_encryption_key_name"))
}

// RuntimeType returns a reference to field runtime_type of google_apigee_organization.
func (ao apigeeOrganizationAttributes) RuntimeType() terra.StringValue {
	return terra.ReferenceAsString(ao.ref.Append("runtime_type"))
}

// SubscriptionType returns a reference to field subscription_type of google_apigee_organization.
func (ao apigeeOrganizationAttributes) SubscriptionType() terra.StringValue {
	return terra.ReferenceAsString(ao.ref.Append("subscription_type"))
}

func (ao apigeeOrganizationAttributes) Properties() terra.ListValue[apigeeorganization.PropertiesAttributes] {
	return terra.ReferenceAsList[apigeeorganization.PropertiesAttributes](ao.ref.Append("properties"))
}

func (ao apigeeOrganizationAttributes) Timeouts() apigeeorganization.TimeoutsAttributes {
	return terra.ReferenceAsSingle[apigeeorganization.TimeoutsAttributes](ao.ref.Append("timeouts"))
}

type apigeeOrganizationState struct {
	AnalyticsRegion                  string                               `json:"analytics_region"`
	AuthorizedNetwork                string                               `json:"authorized_network"`
	BillingType                      string                               `json:"billing_type"`
	CaCertificate                    string                               `json:"ca_certificate"`
	Description                      string                               `json:"description"`
	DisplayName                      string                               `json:"display_name"`
	Id                               string                               `json:"id"`
	Name                             string                               `json:"name"`
	ProjectId                        string                               `json:"project_id"`
	Retention                        string                               `json:"retention"`
	RuntimeDatabaseEncryptionKeyName string                               `json:"runtime_database_encryption_key_name"`
	RuntimeType                      string                               `json:"runtime_type"`
	SubscriptionType                 string                               `json:"subscription_type"`
	Properties                       []apigeeorganization.PropertiesState `json:"properties"`
	Timeouts                         *apigeeorganization.TimeoutsState    `json:"timeouts"`
}
