// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_redshiftserverless_custom_domain_association

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

// Resource represents the Terraform resource aws_redshiftserverless_custom_domain_association.
type Resource struct {
	Name      string
	Args      Args
	state     *awsRedshiftserverlessCustomDomainAssociationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (arcda *Resource) Type() string {
	return "aws_redshiftserverless_custom_domain_association"
}

// LocalName returns the local name for [Resource].
func (arcda *Resource) LocalName() string {
	return arcda.Name
}

// Configuration returns the configuration (args) for [Resource].
func (arcda *Resource) Configuration() interface{} {
	return arcda.Args
}

// DependOn is used for other resources to depend on [Resource].
func (arcda *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(arcda)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (arcda *Resource) Dependencies() terra.Dependencies {
	return arcda.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (arcda *Resource) LifecycleManagement() *terra.Lifecycle {
	return arcda.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (arcda *Resource) Attributes() awsRedshiftserverlessCustomDomainAssociationAttributes {
	return awsRedshiftserverlessCustomDomainAssociationAttributes{ref: terra.ReferenceResource(arcda)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (arcda *Resource) ImportState(state io.Reader) error {
	arcda.state = &awsRedshiftserverlessCustomDomainAssociationState{}
	if err := json.NewDecoder(state).Decode(arcda.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", arcda.Type(), arcda.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (arcda *Resource) State() (*awsRedshiftserverlessCustomDomainAssociationState, bool) {
	return arcda.state, arcda.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (arcda *Resource) StateMust() *awsRedshiftserverlessCustomDomainAssociationState {
	if arcda.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", arcda.Type(), arcda.LocalName()))
	}
	return arcda.state
}

// Args contains the configurations for aws_redshiftserverless_custom_domain_association.
type Args struct {
	// CustomDomainCertificateArn: string, required
	CustomDomainCertificateArn terra.StringValue `hcl:"custom_domain_certificate_arn,attr" validate:"required"`
	// CustomDomainName: string, required
	CustomDomainName terra.StringValue `hcl:"custom_domain_name,attr" validate:"required"`
	// WorkgroupName: string, required
	WorkgroupName terra.StringValue `hcl:"workgroup_name,attr" validate:"required"`
}

type awsRedshiftserverlessCustomDomainAssociationAttributes struct {
	ref terra.Reference
}

// CustomDomainCertificateArn returns a reference to field custom_domain_certificate_arn of aws_redshiftserverless_custom_domain_association.
func (arcda awsRedshiftserverlessCustomDomainAssociationAttributes) CustomDomainCertificateArn() terra.StringValue {
	return terra.ReferenceAsString(arcda.ref.Append("custom_domain_certificate_arn"))
}

// CustomDomainCertificateExpiryTime returns a reference to field custom_domain_certificate_expiry_time of aws_redshiftserverless_custom_domain_association.
func (arcda awsRedshiftserverlessCustomDomainAssociationAttributes) CustomDomainCertificateExpiryTime() terra.StringValue {
	return terra.ReferenceAsString(arcda.ref.Append("custom_domain_certificate_expiry_time"))
}

// CustomDomainName returns a reference to field custom_domain_name of aws_redshiftserverless_custom_domain_association.
func (arcda awsRedshiftserverlessCustomDomainAssociationAttributes) CustomDomainName() terra.StringValue {
	return terra.ReferenceAsString(arcda.ref.Append("custom_domain_name"))
}

// Id returns a reference to field id of aws_redshiftserverless_custom_domain_association.
func (arcda awsRedshiftserverlessCustomDomainAssociationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(arcda.ref.Append("id"))
}

// WorkgroupName returns a reference to field workgroup_name of aws_redshiftserverless_custom_domain_association.
func (arcda awsRedshiftserverlessCustomDomainAssociationAttributes) WorkgroupName() terra.StringValue {
	return terra.ReferenceAsString(arcda.ref.Append("workgroup_name"))
}

type awsRedshiftserverlessCustomDomainAssociationState struct {
	CustomDomainCertificateArn        string `json:"custom_domain_certificate_arn"`
	CustomDomainCertificateExpiryTime string `json:"custom_domain_certificate_expiry_time"`
	CustomDomainName                  string `json:"custom_domain_name"`
	Id                                string `json:"id"`
	WorkgroupName                     string `json:"workgroup_name"`
}
