// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	amplifydomainassociation "github.com/golingon/terraproviders/aws/4.66.1/amplifydomainassociation"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewAmplifyDomainAssociation creates a new instance of [AmplifyDomainAssociation].
func NewAmplifyDomainAssociation(name string, args AmplifyDomainAssociationArgs) *AmplifyDomainAssociation {
	return &AmplifyDomainAssociation{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AmplifyDomainAssociation)(nil)

// AmplifyDomainAssociation represents the Terraform resource aws_amplify_domain_association.
type AmplifyDomainAssociation struct {
	Name      string
	Args      AmplifyDomainAssociationArgs
	state     *amplifyDomainAssociationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [AmplifyDomainAssociation].
func (ada *AmplifyDomainAssociation) Type() string {
	return "aws_amplify_domain_association"
}

// LocalName returns the local name for [AmplifyDomainAssociation].
func (ada *AmplifyDomainAssociation) LocalName() string {
	return ada.Name
}

// Configuration returns the configuration (args) for [AmplifyDomainAssociation].
func (ada *AmplifyDomainAssociation) Configuration() interface{} {
	return ada.Args
}

// DependOn is used for other resources to depend on [AmplifyDomainAssociation].
func (ada *AmplifyDomainAssociation) DependOn() terra.Reference {
	return terra.ReferenceResource(ada)
}

// Dependencies returns the list of resources [AmplifyDomainAssociation] depends_on.
func (ada *AmplifyDomainAssociation) Dependencies() terra.Dependencies {
	return ada.DependsOn
}

// LifecycleManagement returns the lifecycle block for [AmplifyDomainAssociation].
func (ada *AmplifyDomainAssociation) LifecycleManagement() *terra.Lifecycle {
	return ada.Lifecycle
}

// Attributes returns the attributes for [AmplifyDomainAssociation].
func (ada *AmplifyDomainAssociation) Attributes() amplifyDomainAssociationAttributes {
	return amplifyDomainAssociationAttributes{ref: terra.ReferenceResource(ada)}
}

// ImportState imports the given attribute values into [AmplifyDomainAssociation]'s state.
func (ada *AmplifyDomainAssociation) ImportState(av io.Reader) error {
	ada.state = &amplifyDomainAssociationState{}
	if err := json.NewDecoder(av).Decode(ada.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ada.Type(), ada.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [AmplifyDomainAssociation] has state.
func (ada *AmplifyDomainAssociation) State() (*amplifyDomainAssociationState, bool) {
	return ada.state, ada.state != nil
}

// StateMust returns the state for [AmplifyDomainAssociation]. Panics if the state is nil.
func (ada *AmplifyDomainAssociation) StateMust() *amplifyDomainAssociationState {
	if ada.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ada.Type(), ada.LocalName()))
	}
	return ada.state
}

// AmplifyDomainAssociationArgs contains the configurations for aws_amplify_domain_association.
type AmplifyDomainAssociationArgs struct {
	// AppId: string, required
	AppId terra.StringValue `hcl:"app_id,attr" validate:"required"`
	// DomainName: string, required
	DomainName terra.StringValue `hcl:"domain_name,attr" validate:"required"`
	// EnableAutoSubDomain: bool, optional
	EnableAutoSubDomain terra.BoolValue `hcl:"enable_auto_sub_domain,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// WaitForVerification: bool, optional
	WaitForVerification terra.BoolValue `hcl:"wait_for_verification,attr"`
	// SubDomain: min=1
	SubDomain []amplifydomainassociation.SubDomain `hcl:"sub_domain,block" validate:"min=1"`
}
type amplifyDomainAssociationAttributes struct {
	ref terra.Reference
}

// AppId returns a reference to field app_id of aws_amplify_domain_association.
func (ada amplifyDomainAssociationAttributes) AppId() terra.StringValue {
	return terra.ReferenceAsString(ada.ref.Append("app_id"))
}

// Arn returns a reference to field arn of aws_amplify_domain_association.
func (ada amplifyDomainAssociationAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ada.ref.Append("arn"))
}

// CertificateVerificationDnsRecord returns a reference to field certificate_verification_dns_record of aws_amplify_domain_association.
func (ada amplifyDomainAssociationAttributes) CertificateVerificationDnsRecord() terra.StringValue {
	return terra.ReferenceAsString(ada.ref.Append("certificate_verification_dns_record"))
}

// DomainName returns a reference to field domain_name of aws_amplify_domain_association.
func (ada amplifyDomainAssociationAttributes) DomainName() terra.StringValue {
	return terra.ReferenceAsString(ada.ref.Append("domain_name"))
}

// EnableAutoSubDomain returns a reference to field enable_auto_sub_domain of aws_amplify_domain_association.
func (ada amplifyDomainAssociationAttributes) EnableAutoSubDomain() terra.BoolValue {
	return terra.ReferenceAsBool(ada.ref.Append("enable_auto_sub_domain"))
}

// Id returns a reference to field id of aws_amplify_domain_association.
func (ada amplifyDomainAssociationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ada.ref.Append("id"))
}

// WaitForVerification returns a reference to field wait_for_verification of aws_amplify_domain_association.
func (ada amplifyDomainAssociationAttributes) WaitForVerification() terra.BoolValue {
	return terra.ReferenceAsBool(ada.ref.Append("wait_for_verification"))
}

func (ada amplifyDomainAssociationAttributes) SubDomain() terra.SetValue[amplifydomainassociation.SubDomainAttributes] {
	return terra.ReferenceAsSet[amplifydomainassociation.SubDomainAttributes](ada.ref.Append("sub_domain"))
}

type amplifyDomainAssociationState struct {
	AppId                            string                                    `json:"app_id"`
	Arn                              string                                    `json:"arn"`
	CertificateVerificationDnsRecord string                                    `json:"certificate_verification_dns_record"`
	DomainName                       string                                    `json:"domain_name"`
	EnableAutoSubDomain              bool                                      `json:"enable_auto_sub_domain"`
	Id                               string                                    `json:"id"`
	WaitForVerification              bool                                      `json:"wait_for_verification"`
	SubDomain                        []amplifydomainassociation.SubDomainState `json:"sub_domain"`
}
