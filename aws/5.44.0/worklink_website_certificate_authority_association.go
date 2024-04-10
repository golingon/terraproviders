// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// NewWorklinkWebsiteCertificateAuthorityAssociation creates a new instance of [WorklinkWebsiteCertificateAuthorityAssociation].
func NewWorklinkWebsiteCertificateAuthorityAssociation(name string, args WorklinkWebsiteCertificateAuthorityAssociationArgs) *WorklinkWebsiteCertificateAuthorityAssociation {
	return &WorklinkWebsiteCertificateAuthorityAssociation{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*WorklinkWebsiteCertificateAuthorityAssociation)(nil)

// WorklinkWebsiteCertificateAuthorityAssociation represents the Terraform resource aws_worklink_website_certificate_authority_association.
type WorklinkWebsiteCertificateAuthorityAssociation struct {
	Name      string
	Args      WorklinkWebsiteCertificateAuthorityAssociationArgs
	state     *worklinkWebsiteCertificateAuthorityAssociationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [WorklinkWebsiteCertificateAuthorityAssociation].
func (wwcaa *WorklinkWebsiteCertificateAuthorityAssociation) Type() string {
	return "aws_worklink_website_certificate_authority_association"
}

// LocalName returns the local name for [WorklinkWebsiteCertificateAuthorityAssociation].
func (wwcaa *WorklinkWebsiteCertificateAuthorityAssociation) LocalName() string {
	return wwcaa.Name
}

// Configuration returns the configuration (args) for [WorklinkWebsiteCertificateAuthorityAssociation].
func (wwcaa *WorklinkWebsiteCertificateAuthorityAssociation) Configuration() interface{} {
	return wwcaa.Args
}

// DependOn is used for other resources to depend on [WorklinkWebsiteCertificateAuthorityAssociation].
func (wwcaa *WorklinkWebsiteCertificateAuthorityAssociation) DependOn() terra.Reference {
	return terra.ReferenceResource(wwcaa)
}

// Dependencies returns the list of resources [WorklinkWebsiteCertificateAuthorityAssociation] depends_on.
func (wwcaa *WorklinkWebsiteCertificateAuthorityAssociation) Dependencies() terra.Dependencies {
	return wwcaa.DependsOn
}

// LifecycleManagement returns the lifecycle block for [WorklinkWebsiteCertificateAuthorityAssociation].
func (wwcaa *WorklinkWebsiteCertificateAuthorityAssociation) LifecycleManagement() *terra.Lifecycle {
	return wwcaa.Lifecycle
}

// Attributes returns the attributes for [WorklinkWebsiteCertificateAuthorityAssociation].
func (wwcaa *WorklinkWebsiteCertificateAuthorityAssociation) Attributes() worklinkWebsiteCertificateAuthorityAssociationAttributes {
	return worklinkWebsiteCertificateAuthorityAssociationAttributes{ref: terra.ReferenceResource(wwcaa)}
}

// ImportState imports the given attribute values into [WorklinkWebsiteCertificateAuthorityAssociation]'s state.
func (wwcaa *WorklinkWebsiteCertificateAuthorityAssociation) ImportState(av io.Reader) error {
	wwcaa.state = &worklinkWebsiteCertificateAuthorityAssociationState{}
	if err := json.NewDecoder(av).Decode(wwcaa.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", wwcaa.Type(), wwcaa.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [WorklinkWebsiteCertificateAuthorityAssociation] has state.
func (wwcaa *WorklinkWebsiteCertificateAuthorityAssociation) State() (*worklinkWebsiteCertificateAuthorityAssociationState, bool) {
	return wwcaa.state, wwcaa.state != nil
}

// StateMust returns the state for [WorklinkWebsiteCertificateAuthorityAssociation]. Panics if the state is nil.
func (wwcaa *WorklinkWebsiteCertificateAuthorityAssociation) StateMust() *worklinkWebsiteCertificateAuthorityAssociationState {
	if wwcaa.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", wwcaa.Type(), wwcaa.LocalName()))
	}
	return wwcaa.state
}

// WorklinkWebsiteCertificateAuthorityAssociationArgs contains the configurations for aws_worklink_website_certificate_authority_association.
type WorklinkWebsiteCertificateAuthorityAssociationArgs struct {
	// Certificate: string, required
	Certificate terra.StringValue `hcl:"certificate,attr" validate:"required"`
	// DisplayName: string, optional
	DisplayName terra.StringValue `hcl:"display_name,attr"`
	// FleetArn: string, required
	FleetArn terra.StringValue `hcl:"fleet_arn,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
}
type worklinkWebsiteCertificateAuthorityAssociationAttributes struct {
	ref terra.Reference
}

// Certificate returns a reference to field certificate of aws_worklink_website_certificate_authority_association.
func (wwcaa worklinkWebsiteCertificateAuthorityAssociationAttributes) Certificate() terra.StringValue {
	return terra.ReferenceAsString(wwcaa.ref.Append("certificate"))
}

// DisplayName returns a reference to field display_name of aws_worklink_website_certificate_authority_association.
func (wwcaa worklinkWebsiteCertificateAuthorityAssociationAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(wwcaa.ref.Append("display_name"))
}

// FleetArn returns a reference to field fleet_arn of aws_worklink_website_certificate_authority_association.
func (wwcaa worklinkWebsiteCertificateAuthorityAssociationAttributes) FleetArn() terra.StringValue {
	return terra.ReferenceAsString(wwcaa.ref.Append("fleet_arn"))
}

// Id returns a reference to field id of aws_worklink_website_certificate_authority_association.
func (wwcaa worklinkWebsiteCertificateAuthorityAssociationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(wwcaa.ref.Append("id"))
}

// WebsiteCaId returns a reference to field website_ca_id of aws_worklink_website_certificate_authority_association.
func (wwcaa worklinkWebsiteCertificateAuthorityAssociationAttributes) WebsiteCaId() terra.StringValue {
	return terra.ReferenceAsString(wwcaa.ref.Append("website_ca_id"))
}

type worklinkWebsiteCertificateAuthorityAssociationState struct {
	Certificate string `json:"certificate"`
	DisplayName string `json:"display_name"`
	FleetArn    string `json:"fleet_arn"`
	Id          string `json:"id"`
	WebsiteCaId string `json:"website_ca_id"`
}
