// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewLicensemanagerGrantAccepter creates a new instance of [LicensemanagerGrantAccepter].
func NewLicensemanagerGrantAccepter(name string, args LicensemanagerGrantAccepterArgs) *LicensemanagerGrantAccepter {
	return &LicensemanagerGrantAccepter{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*LicensemanagerGrantAccepter)(nil)

// LicensemanagerGrantAccepter represents the Terraform resource aws_licensemanager_grant_accepter.
type LicensemanagerGrantAccepter struct {
	Name      string
	Args      LicensemanagerGrantAccepterArgs
	state     *licensemanagerGrantAccepterState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [LicensemanagerGrantAccepter].
func (lga *LicensemanagerGrantAccepter) Type() string {
	return "aws_licensemanager_grant_accepter"
}

// LocalName returns the local name for [LicensemanagerGrantAccepter].
func (lga *LicensemanagerGrantAccepter) LocalName() string {
	return lga.Name
}

// Configuration returns the configuration (args) for [LicensemanagerGrantAccepter].
func (lga *LicensemanagerGrantAccepter) Configuration() interface{} {
	return lga.Args
}

// DependOn is used for other resources to depend on [LicensemanagerGrantAccepter].
func (lga *LicensemanagerGrantAccepter) DependOn() terra.Reference {
	return terra.ReferenceResource(lga)
}

// Dependencies returns the list of resources [LicensemanagerGrantAccepter] depends_on.
func (lga *LicensemanagerGrantAccepter) Dependencies() terra.Dependencies {
	return lga.DependsOn
}

// LifecycleManagement returns the lifecycle block for [LicensemanagerGrantAccepter].
func (lga *LicensemanagerGrantAccepter) LifecycleManagement() *terra.Lifecycle {
	return lga.Lifecycle
}

// Attributes returns the attributes for [LicensemanagerGrantAccepter].
func (lga *LicensemanagerGrantAccepter) Attributes() licensemanagerGrantAccepterAttributes {
	return licensemanagerGrantAccepterAttributes{ref: terra.ReferenceResource(lga)}
}

// ImportState imports the given attribute values into [LicensemanagerGrantAccepter]'s state.
func (lga *LicensemanagerGrantAccepter) ImportState(av io.Reader) error {
	lga.state = &licensemanagerGrantAccepterState{}
	if err := json.NewDecoder(av).Decode(lga.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", lga.Type(), lga.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [LicensemanagerGrantAccepter] has state.
func (lga *LicensemanagerGrantAccepter) State() (*licensemanagerGrantAccepterState, bool) {
	return lga.state, lga.state != nil
}

// StateMust returns the state for [LicensemanagerGrantAccepter]. Panics if the state is nil.
func (lga *LicensemanagerGrantAccepter) StateMust() *licensemanagerGrantAccepterState {
	if lga.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", lga.Type(), lga.LocalName()))
	}
	return lga.state
}

// LicensemanagerGrantAccepterArgs contains the configurations for aws_licensemanager_grant_accepter.
type LicensemanagerGrantAccepterArgs struct {
	// GrantArn: string, required
	GrantArn terra.StringValue `hcl:"grant_arn,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
}
type licensemanagerGrantAccepterAttributes struct {
	ref terra.Reference
}

// AllowedOperations returns a reference to field allowed_operations of aws_licensemanager_grant_accepter.
func (lga licensemanagerGrantAccepterAttributes) AllowedOperations() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](lga.ref.Append("allowed_operations"))
}

// GrantArn returns a reference to field grant_arn of aws_licensemanager_grant_accepter.
func (lga licensemanagerGrantAccepterAttributes) GrantArn() terra.StringValue {
	return terra.ReferenceAsString(lga.ref.Append("grant_arn"))
}

// HomeRegion returns a reference to field home_region of aws_licensemanager_grant_accepter.
func (lga licensemanagerGrantAccepterAttributes) HomeRegion() terra.StringValue {
	return terra.ReferenceAsString(lga.ref.Append("home_region"))
}

// Id returns a reference to field id of aws_licensemanager_grant_accepter.
func (lga licensemanagerGrantAccepterAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(lga.ref.Append("id"))
}

// LicenseArn returns a reference to field license_arn of aws_licensemanager_grant_accepter.
func (lga licensemanagerGrantAccepterAttributes) LicenseArn() terra.StringValue {
	return terra.ReferenceAsString(lga.ref.Append("license_arn"))
}

// Name returns a reference to field name of aws_licensemanager_grant_accepter.
func (lga licensemanagerGrantAccepterAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(lga.ref.Append("name"))
}

// ParentArn returns a reference to field parent_arn of aws_licensemanager_grant_accepter.
func (lga licensemanagerGrantAccepterAttributes) ParentArn() terra.StringValue {
	return terra.ReferenceAsString(lga.ref.Append("parent_arn"))
}

// Principal returns a reference to field principal of aws_licensemanager_grant_accepter.
func (lga licensemanagerGrantAccepterAttributes) Principal() terra.StringValue {
	return terra.ReferenceAsString(lga.ref.Append("principal"))
}

// Status returns a reference to field status of aws_licensemanager_grant_accepter.
func (lga licensemanagerGrantAccepterAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(lga.ref.Append("status"))
}

// Version returns a reference to field version of aws_licensemanager_grant_accepter.
func (lga licensemanagerGrantAccepterAttributes) Version() terra.StringValue {
	return terra.ReferenceAsString(lga.ref.Append("version"))
}

type licensemanagerGrantAccepterState struct {
	AllowedOperations []string `json:"allowed_operations"`
	GrantArn          string   `json:"grant_arn"`
	HomeRegion        string   `json:"home_region"`
	Id                string   `json:"id"`
	LicenseArn        string   `json:"license_arn"`
	Name              string   `json:"name"`
	ParentArn         string   `json:"parent_arn"`
	Principal         string   `json:"principal"`
	Status            string   `json:"status"`
	Version           string   `json:"version"`
}