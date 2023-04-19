// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	apprunnercustomdomainassociation "github.com/golingon/terraproviders/aws/4.63.0/apprunnercustomdomainassociation"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewApprunnerCustomDomainAssociation creates a new instance of [ApprunnerCustomDomainAssociation].
func NewApprunnerCustomDomainAssociation(name string, args ApprunnerCustomDomainAssociationArgs) *ApprunnerCustomDomainAssociation {
	return &ApprunnerCustomDomainAssociation{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ApprunnerCustomDomainAssociation)(nil)

// ApprunnerCustomDomainAssociation represents the Terraform resource aws_apprunner_custom_domain_association.
type ApprunnerCustomDomainAssociation struct {
	Name      string
	Args      ApprunnerCustomDomainAssociationArgs
	state     *apprunnerCustomDomainAssociationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ApprunnerCustomDomainAssociation].
func (acda *ApprunnerCustomDomainAssociation) Type() string {
	return "aws_apprunner_custom_domain_association"
}

// LocalName returns the local name for [ApprunnerCustomDomainAssociation].
func (acda *ApprunnerCustomDomainAssociation) LocalName() string {
	return acda.Name
}

// Configuration returns the configuration (args) for [ApprunnerCustomDomainAssociation].
func (acda *ApprunnerCustomDomainAssociation) Configuration() interface{} {
	return acda.Args
}

// DependOn is used for other resources to depend on [ApprunnerCustomDomainAssociation].
func (acda *ApprunnerCustomDomainAssociation) DependOn() terra.Reference {
	return terra.ReferenceResource(acda)
}

// Dependencies returns the list of resources [ApprunnerCustomDomainAssociation] depends_on.
func (acda *ApprunnerCustomDomainAssociation) Dependencies() terra.Dependencies {
	return acda.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ApprunnerCustomDomainAssociation].
func (acda *ApprunnerCustomDomainAssociation) LifecycleManagement() *terra.Lifecycle {
	return acda.Lifecycle
}

// Attributes returns the attributes for [ApprunnerCustomDomainAssociation].
func (acda *ApprunnerCustomDomainAssociation) Attributes() apprunnerCustomDomainAssociationAttributes {
	return apprunnerCustomDomainAssociationAttributes{ref: terra.ReferenceResource(acda)}
}

// ImportState imports the given attribute values into [ApprunnerCustomDomainAssociation]'s state.
func (acda *ApprunnerCustomDomainAssociation) ImportState(av io.Reader) error {
	acda.state = &apprunnerCustomDomainAssociationState{}
	if err := json.NewDecoder(av).Decode(acda.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", acda.Type(), acda.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ApprunnerCustomDomainAssociation] has state.
func (acda *ApprunnerCustomDomainAssociation) State() (*apprunnerCustomDomainAssociationState, bool) {
	return acda.state, acda.state != nil
}

// StateMust returns the state for [ApprunnerCustomDomainAssociation]. Panics if the state is nil.
func (acda *ApprunnerCustomDomainAssociation) StateMust() *apprunnerCustomDomainAssociationState {
	if acda.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", acda.Type(), acda.LocalName()))
	}
	return acda.state
}

// ApprunnerCustomDomainAssociationArgs contains the configurations for aws_apprunner_custom_domain_association.
type ApprunnerCustomDomainAssociationArgs struct {
	// DomainName: string, required
	DomainName terra.StringValue `hcl:"domain_name,attr" validate:"required"`
	// EnableWwwSubdomain: bool, optional
	EnableWwwSubdomain terra.BoolValue `hcl:"enable_www_subdomain,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ServiceArn: string, required
	ServiceArn terra.StringValue `hcl:"service_arn,attr" validate:"required"`
	// CertificateValidationRecords: min=0
	CertificateValidationRecords []apprunnercustomdomainassociation.CertificateValidationRecords `hcl:"certificate_validation_records,block" validate:"min=0"`
}
type apprunnerCustomDomainAssociationAttributes struct {
	ref terra.Reference
}

// DnsTarget returns a reference to field dns_target of aws_apprunner_custom_domain_association.
func (acda apprunnerCustomDomainAssociationAttributes) DnsTarget() terra.StringValue {
	return terra.ReferenceAsString(acda.ref.Append("dns_target"))
}

// DomainName returns a reference to field domain_name of aws_apprunner_custom_domain_association.
func (acda apprunnerCustomDomainAssociationAttributes) DomainName() terra.StringValue {
	return terra.ReferenceAsString(acda.ref.Append("domain_name"))
}

// EnableWwwSubdomain returns a reference to field enable_www_subdomain of aws_apprunner_custom_domain_association.
func (acda apprunnerCustomDomainAssociationAttributes) EnableWwwSubdomain() terra.BoolValue {
	return terra.ReferenceAsBool(acda.ref.Append("enable_www_subdomain"))
}

// Id returns a reference to field id of aws_apprunner_custom_domain_association.
func (acda apprunnerCustomDomainAssociationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(acda.ref.Append("id"))
}

// ServiceArn returns a reference to field service_arn of aws_apprunner_custom_domain_association.
func (acda apprunnerCustomDomainAssociationAttributes) ServiceArn() terra.StringValue {
	return terra.ReferenceAsString(acda.ref.Append("service_arn"))
}

// Status returns a reference to field status of aws_apprunner_custom_domain_association.
func (acda apprunnerCustomDomainAssociationAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(acda.ref.Append("status"))
}

func (acda apprunnerCustomDomainAssociationAttributes) CertificateValidationRecords() terra.SetValue[apprunnercustomdomainassociation.CertificateValidationRecordsAttributes] {
	return terra.ReferenceAsSet[apprunnercustomdomainassociation.CertificateValidationRecordsAttributes](acda.ref.Append("certificate_validation_records"))
}

type apprunnerCustomDomainAssociationState struct {
	DnsTarget                    string                                                               `json:"dns_target"`
	DomainName                   string                                                               `json:"domain_name"`
	EnableWwwSubdomain           bool                                                                 `json:"enable_www_subdomain"`
	Id                           string                                                               `json:"id"`
	ServiceArn                   string                                                               `json:"service_arn"`
	Status                       string                                                               `json:"status"`
	CertificateValidationRecords []apprunnercustomdomainassociation.CertificateValidationRecordsState `json:"certificate_validation_records"`
}
