// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	grafanalicenseassociation "github.com/golingon/terraproviders/aws/5.7.0/grafanalicenseassociation"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewGrafanaLicenseAssociation creates a new instance of [GrafanaLicenseAssociation].
func NewGrafanaLicenseAssociation(name string, args GrafanaLicenseAssociationArgs) *GrafanaLicenseAssociation {
	return &GrafanaLicenseAssociation{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*GrafanaLicenseAssociation)(nil)

// GrafanaLicenseAssociation represents the Terraform resource aws_grafana_license_association.
type GrafanaLicenseAssociation struct {
	Name      string
	Args      GrafanaLicenseAssociationArgs
	state     *grafanaLicenseAssociationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [GrafanaLicenseAssociation].
func (gla *GrafanaLicenseAssociation) Type() string {
	return "aws_grafana_license_association"
}

// LocalName returns the local name for [GrafanaLicenseAssociation].
func (gla *GrafanaLicenseAssociation) LocalName() string {
	return gla.Name
}

// Configuration returns the configuration (args) for [GrafanaLicenseAssociation].
func (gla *GrafanaLicenseAssociation) Configuration() interface{} {
	return gla.Args
}

// DependOn is used for other resources to depend on [GrafanaLicenseAssociation].
func (gla *GrafanaLicenseAssociation) DependOn() terra.Reference {
	return terra.ReferenceResource(gla)
}

// Dependencies returns the list of resources [GrafanaLicenseAssociation] depends_on.
func (gla *GrafanaLicenseAssociation) Dependencies() terra.Dependencies {
	return gla.DependsOn
}

// LifecycleManagement returns the lifecycle block for [GrafanaLicenseAssociation].
func (gla *GrafanaLicenseAssociation) LifecycleManagement() *terra.Lifecycle {
	return gla.Lifecycle
}

// Attributes returns the attributes for [GrafanaLicenseAssociation].
func (gla *GrafanaLicenseAssociation) Attributes() grafanaLicenseAssociationAttributes {
	return grafanaLicenseAssociationAttributes{ref: terra.ReferenceResource(gla)}
}

// ImportState imports the given attribute values into [GrafanaLicenseAssociation]'s state.
func (gla *GrafanaLicenseAssociation) ImportState(av io.Reader) error {
	gla.state = &grafanaLicenseAssociationState{}
	if err := json.NewDecoder(av).Decode(gla.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gla.Type(), gla.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [GrafanaLicenseAssociation] has state.
func (gla *GrafanaLicenseAssociation) State() (*grafanaLicenseAssociationState, bool) {
	return gla.state, gla.state != nil
}

// StateMust returns the state for [GrafanaLicenseAssociation]. Panics if the state is nil.
func (gla *GrafanaLicenseAssociation) StateMust() *grafanaLicenseAssociationState {
	if gla.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gla.Type(), gla.LocalName()))
	}
	return gla.state
}

// GrafanaLicenseAssociationArgs contains the configurations for aws_grafana_license_association.
type GrafanaLicenseAssociationArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LicenseType: string, required
	LicenseType terra.StringValue `hcl:"license_type,attr" validate:"required"`
	// WorkspaceId: string, required
	WorkspaceId terra.StringValue `hcl:"workspace_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *grafanalicenseassociation.Timeouts `hcl:"timeouts,block"`
}
type grafanaLicenseAssociationAttributes struct {
	ref terra.Reference
}

// FreeTrialExpiration returns a reference to field free_trial_expiration of aws_grafana_license_association.
func (gla grafanaLicenseAssociationAttributes) FreeTrialExpiration() terra.StringValue {
	return terra.ReferenceAsString(gla.ref.Append("free_trial_expiration"))
}

// Id returns a reference to field id of aws_grafana_license_association.
func (gla grafanaLicenseAssociationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gla.ref.Append("id"))
}

// LicenseExpiration returns a reference to field license_expiration of aws_grafana_license_association.
func (gla grafanaLicenseAssociationAttributes) LicenseExpiration() terra.StringValue {
	return terra.ReferenceAsString(gla.ref.Append("license_expiration"))
}

// LicenseType returns a reference to field license_type of aws_grafana_license_association.
func (gla grafanaLicenseAssociationAttributes) LicenseType() terra.StringValue {
	return terra.ReferenceAsString(gla.ref.Append("license_type"))
}

// WorkspaceId returns a reference to field workspace_id of aws_grafana_license_association.
func (gla grafanaLicenseAssociationAttributes) WorkspaceId() terra.StringValue {
	return terra.ReferenceAsString(gla.ref.Append("workspace_id"))
}

func (gla grafanaLicenseAssociationAttributes) Timeouts() grafanalicenseassociation.TimeoutsAttributes {
	return terra.ReferenceAsSingle[grafanalicenseassociation.TimeoutsAttributes](gla.ref.Append("timeouts"))
}

type grafanaLicenseAssociationState struct {
	FreeTrialExpiration string                                   `json:"free_trial_expiration"`
	Id                  string                                   `json:"id"`
	LicenseExpiration   string                                   `json:"license_expiration"`
	LicenseType         string                                   `json:"license_type"`
	WorkspaceId         string                                   `json:"workspace_id"`
	Timeouts            *grafanalicenseassociation.TimeoutsState `json:"timeouts"`
}
