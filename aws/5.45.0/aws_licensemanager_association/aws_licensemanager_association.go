// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_licensemanager_association

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

// Resource represents the Terraform resource aws_licensemanager_association.
type Resource struct {
	Name      string
	Args      Args
	state     *awsLicensemanagerAssociationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (ala *Resource) Type() string {
	return "aws_licensemanager_association"
}

// LocalName returns the local name for [Resource].
func (ala *Resource) LocalName() string {
	return ala.Name
}

// Configuration returns the configuration (args) for [Resource].
func (ala *Resource) Configuration() interface{} {
	return ala.Args
}

// DependOn is used for other resources to depend on [Resource].
func (ala *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(ala)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (ala *Resource) Dependencies() terra.Dependencies {
	return ala.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (ala *Resource) LifecycleManagement() *terra.Lifecycle {
	return ala.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (ala *Resource) Attributes() awsLicensemanagerAssociationAttributes {
	return awsLicensemanagerAssociationAttributes{ref: terra.ReferenceResource(ala)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (ala *Resource) ImportState(state io.Reader) error {
	ala.state = &awsLicensemanagerAssociationState{}
	if err := json.NewDecoder(state).Decode(ala.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ala.Type(), ala.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (ala *Resource) State() (*awsLicensemanagerAssociationState, bool) {
	return ala.state, ala.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (ala *Resource) StateMust() *awsLicensemanagerAssociationState {
	if ala.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ala.Type(), ala.LocalName()))
	}
	return ala.state
}

// Args contains the configurations for aws_licensemanager_association.
type Args struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LicenseConfigurationArn: string, required
	LicenseConfigurationArn terra.StringValue `hcl:"license_configuration_arn,attr" validate:"required"`
	// ResourceArn: string, required
	ResourceArn terra.StringValue `hcl:"resource_arn,attr" validate:"required"`
}

type awsLicensemanagerAssociationAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_licensemanager_association.
func (ala awsLicensemanagerAssociationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ala.ref.Append("id"))
}

// LicenseConfigurationArn returns a reference to field license_configuration_arn of aws_licensemanager_association.
func (ala awsLicensemanagerAssociationAttributes) LicenseConfigurationArn() terra.StringValue {
	return terra.ReferenceAsString(ala.ref.Append("license_configuration_arn"))
}

// ResourceArn returns a reference to field resource_arn of aws_licensemanager_association.
func (ala awsLicensemanagerAssociationAttributes) ResourceArn() terra.StringValue {
	return terra.ReferenceAsString(ala.ref.Append("resource_arn"))
}

type awsLicensemanagerAssociationState struct {
	Id                      string `json:"id"`
	LicenseConfigurationArn string `json:"license_configuration_arn"`
	ResourceArn             string `json:"resource_arn"`
}
