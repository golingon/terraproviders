// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	transferaccess "github.com/golingon/terraproviders/aws/5.44.0/transferaccess"
	"io"
)

// NewTransferAccess creates a new instance of [TransferAccess].
func NewTransferAccess(name string, args TransferAccessArgs) *TransferAccess {
	return &TransferAccess{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*TransferAccess)(nil)

// TransferAccess represents the Terraform resource aws_transfer_access.
type TransferAccess struct {
	Name      string
	Args      TransferAccessArgs
	state     *transferAccessState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [TransferAccess].
func (ta *TransferAccess) Type() string {
	return "aws_transfer_access"
}

// LocalName returns the local name for [TransferAccess].
func (ta *TransferAccess) LocalName() string {
	return ta.Name
}

// Configuration returns the configuration (args) for [TransferAccess].
func (ta *TransferAccess) Configuration() interface{} {
	return ta.Args
}

// DependOn is used for other resources to depend on [TransferAccess].
func (ta *TransferAccess) DependOn() terra.Reference {
	return terra.ReferenceResource(ta)
}

// Dependencies returns the list of resources [TransferAccess] depends_on.
func (ta *TransferAccess) Dependencies() terra.Dependencies {
	return ta.DependsOn
}

// LifecycleManagement returns the lifecycle block for [TransferAccess].
func (ta *TransferAccess) LifecycleManagement() *terra.Lifecycle {
	return ta.Lifecycle
}

// Attributes returns the attributes for [TransferAccess].
func (ta *TransferAccess) Attributes() transferAccessAttributes {
	return transferAccessAttributes{ref: terra.ReferenceResource(ta)}
}

// ImportState imports the given attribute values into [TransferAccess]'s state.
func (ta *TransferAccess) ImportState(av io.Reader) error {
	ta.state = &transferAccessState{}
	if err := json.NewDecoder(av).Decode(ta.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ta.Type(), ta.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [TransferAccess] has state.
func (ta *TransferAccess) State() (*transferAccessState, bool) {
	return ta.state, ta.state != nil
}

// StateMust returns the state for [TransferAccess]. Panics if the state is nil.
func (ta *TransferAccess) StateMust() *transferAccessState {
	if ta.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ta.Type(), ta.LocalName()))
	}
	return ta.state
}

// TransferAccessArgs contains the configurations for aws_transfer_access.
type TransferAccessArgs struct {
	// ExternalId: string, required
	ExternalId terra.StringValue `hcl:"external_id,attr" validate:"required"`
	// HomeDirectory: string, optional
	HomeDirectory terra.StringValue `hcl:"home_directory,attr"`
	// HomeDirectoryType: string, optional
	HomeDirectoryType terra.StringValue `hcl:"home_directory_type,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Policy: string, optional
	Policy terra.StringValue `hcl:"policy,attr"`
	// Role: string, optional
	Role terra.StringValue `hcl:"role,attr"`
	// ServerId: string, required
	ServerId terra.StringValue `hcl:"server_id,attr" validate:"required"`
	// HomeDirectoryMappings: min=0,max=50
	HomeDirectoryMappings []transferaccess.HomeDirectoryMappings `hcl:"home_directory_mappings,block" validate:"min=0,max=50"`
	// PosixProfile: optional
	PosixProfile *transferaccess.PosixProfile `hcl:"posix_profile,block"`
}
type transferAccessAttributes struct {
	ref terra.Reference
}

// ExternalId returns a reference to field external_id of aws_transfer_access.
func (ta transferAccessAttributes) ExternalId() terra.StringValue {
	return terra.ReferenceAsString(ta.ref.Append("external_id"))
}

// HomeDirectory returns a reference to field home_directory of aws_transfer_access.
func (ta transferAccessAttributes) HomeDirectory() terra.StringValue {
	return terra.ReferenceAsString(ta.ref.Append("home_directory"))
}

// HomeDirectoryType returns a reference to field home_directory_type of aws_transfer_access.
func (ta transferAccessAttributes) HomeDirectoryType() terra.StringValue {
	return terra.ReferenceAsString(ta.ref.Append("home_directory_type"))
}

// Id returns a reference to field id of aws_transfer_access.
func (ta transferAccessAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ta.ref.Append("id"))
}

// Policy returns a reference to field policy of aws_transfer_access.
func (ta transferAccessAttributes) Policy() terra.StringValue {
	return terra.ReferenceAsString(ta.ref.Append("policy"))
}

// Role returns a reference to field role of aws_transfer_access.
func (ta transferAccessAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(ta.ref.Append("role"))
}

// ServerId returns a reference to field server_id of aws_transfer_access.
func (ta transferAccessAttributes) ServerId() terra.StringValue {
	return terra.ReferenceAsString(ta.ref.Append("server_id"))
}

func (ta transferAccessAttributes) HomeDirectoryMappings() terra.ListValue[transferaccess.HomeDirectoryMappingsAttributes] {
	return terra.ReferenceAsList[transferaccess.HomeDirectoryMappingsAttributes](ta.ref.Append("home_directory_mappings"))
}

func (ta transferAccessAttributes) PosixProfile() terra.ListValue[transferaccess.PosixProfileAttributes] {
	return terra.ReferenceAsList[transferaccess.PosixProfileAttributes](ta.ref.Append("posix_profile"))
}

type transferAccessState struct {
	ExternalId            string                                      `json:"external_id"`
	HomeDirectory         string                                      `json:"home_directory"`
	HomeDirectoryType     string                                      `json:"home_directory_type"`
	Id                    string                                      `json:"id"`
	Policy                string                                      `json:"policy"`
	Role                  string                                      `json:"role"`
	ServerId              string                                      `json:"server_id"`
	HomeDirectoryMappings []transferaccess.HomeDirectoryMappingsState `json:"home_directory_mappings"`
	PosixProfile          []transferaccess.PosixProfileState          `json:"posix_profile"`
}
