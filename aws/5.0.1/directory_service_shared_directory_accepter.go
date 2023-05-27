// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	directoryserviceshareddirectoryaccepter "github.com/golingon/terraproviders/aws/5.0.1/directoryserviceshareddirectoryaccepter"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDirectoryServiceSharedDirectoryAccepter creates a new instance of [DirectoryServiceSharedDirectoryAccepter].
func NewDirectoryServiceSharedDirectoryAccepter(name string, args DirectoryServiceSharedDirectoryAccepterArgs) *DirectoryServiceSharedDirectoryAccepter {
	return &DirectoryServiceSharedDirectoryAccepter{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DirectoryServiceSharedDirectoryAccepter)(nil)

// DirectoryServiceSharedDirectoryAccepter represents the Terraform resource aws_directory_service_shared_directory_accepter.
type DirectoryServiceSharedDirectoryAccepter struct {
	Name      string
	Args      DirectoryServiceSharedDirectoryAccepterArgs
	state     *directoryServiceSharedDirectoryAccepterState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DirectoryServiceSharedDirectoryAccepter].
func (dssda *DirectoryServiceSharedDirectoryAccepter) Type() string {
	return "aws_directory_service_shared_directory_accepter"
}

// LocalName returns the local name for [DirectoryServiceSharedDirectoryAccepter].
func (dssda *DirectoryServiceSharedDirectoryAccepter) LocalName() string {
	return dssda.Name
}

// Configuration returns the configuration (args) for [DirectoryServiceSharedDirectoryAccepter].
func (dssda *DirectoryServiceSharedDirectoryAccepter) Configuration() interface{} {
	return dssda.Args
}

// DependOn is used for other resources to depend on [DirectoryServiceSharedDirectoryAccepter].
func (dssda *DirectoryServiceSharedDirectoryAccepter) DependOn() terra.Reference {
	return terra.ReferenceResource(dssda)
}

// Dependencies returns the list of resources [DirectoryServiceSharedDirectoryAccepter] depends_on.
func (dssda *DirectoryServiceSharedDirectoryAccepter) Dependencies() terra.Dependencies {
	return dssda.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DirectoryServiceSharedDirectoryAccepter].
func (dssda *DirectoryServiceSharedDirectoryAccepter) LifecycleManagement() *terra.Lifecycle {
	return dssda.Lifecycle
}

// Attributes returns the attributes for [DirectoryServiceSharedDirectoryAccepter].
func (dssda *DirectoryServiceSharedDirectoryAccepter) Attributes() directoryServiceSharedDirectoryAccepterAttributes {
	return directoryServiceSharedDirectoryAccepterAttributes{ref: terra.ReferenceResource(dssda)}
}

// ImportState imports the given attribute values into [DirectoryServiceSharedDirectoryAccepter]'s state.
func (dssda *DirectoryServiceSharedDirectoryAccepter) ImportState(av io.Reader) error {
	dssda.state = &directoryServiceSharedDirectoryAccepterState{}
	if err := json.NewDecoder(av).Decode(dssda.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dssda.Type(), dssda.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DirectoryServiceSharedDirectoryAccepter] has state.
func (dssda *DirectoryServiceSharedDirectoryAccepter) State() (*directoryServiceSharedDirectoryAccepterState, bool) {
	return dssda.state, dssda.state != nil
}

// StateMust returns the state for [DirectoryServiceSharedDirectoryAccepter]. Panics if the state is nil.
func (dssda *DirectoryServiceSharedDirectoryAccepter) StateMust() *directoryServiceSharedDirectoryAccepterState {
	if dssda.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dssda.Type(), dssda.LocalName()))
	}
	return dssda.state
}

// DirectoryServiceSharedDirectoryAccepterArgs contains the configurations for aws_directory_service_shared_directory_accepter.
type DirectoryServiceSharedDirectoryAccepterArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// SharedDirectoryId: string, required
	SharedDirectoryId terra.StringValue `hcl:"shared_directory_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *directoryserviceshareddirectoryaccepter.Timeouts `hcl:"timeouts,block"`
}
type directoryServiceSharedDirectoryAccepterAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_directory_service_shared_directory_accepter.
func (dssda directoryServiceSharedDirectoryAccepterAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dssda.ref.Append("id"))
}

// Method returns a reference to field method of aws_directory_service_shared_directory_accepter.
func (dssda directoryServiceSharedDirectoryAccepterAttributes) Method() terra.StringValue {
	return terra.ReferenceAsString(dssda.ref.Append("method"))
}

// Notes returns a reference to field notes of aws_directory_service_shared_directory_accepter.
func (dssda directoryServiceSharedDirectoryAccepterAttributes) Notes() terra.StringValue {
	return terra.ReferenceAsString(dssda.ref.Append("notes"))
}

// OwnerAccountId returns a reference to field owner_account_id of aws_directory_service_shared_directory_accepter.
func (dssda directoryServiceSharedDirectoryAccepterAttributes) OwnerAccountId() terra.StringValue {
	return terra.ReferenceAsString(dssda.ref.Append("owner_account_id"))
}

// OwnerDirectoryId returns a reference to field owner_directory_id of aws_directory_service_shared_directory_accepter.
func (dssda directoryServiceSharedDirectoryAccepterAttributes) OwnerDirectoryId() terra.StringValue {
	return terra.ReferenceAsString(dssda.ref.Append("owner_directory_id"))
}

// SharedDirectoryId returns a reference to field shared_directory_id of aws_directory_service_shared_directory_accepter.
func (dssda directoryServiceSharedDirectoryAccepterAttributes) SharedDirectoryId() terra.StringValue {
	return terra.ReferenceAsString(dssda.ref.Append("shared_directory_id"))
}

func (dssda directoryServiceSharedDirectoryAccepterAttributes) Timeouts() directoryserviceshareddirectoryaccepter.TimeoutsAttributes {
	return terra.ReferenceAsSingle[directoryserviceshareddirectoryaccepter.TimeoutsAttributes](dssda.ref.Append("timeouts"))
}

type directoryServiceSharedDirectoryAccepterState struct {
	Id                string                                                 `json:"id"`
	Method            string                                                 `json:"method"`
	Notes             string                                                 `json:"notes"`
	OwnerAccountId    string                                                 `json:"owner_account_id"`
	OwnerDirectoryId  string                                                 `json:"owner_directory_id"`
	SharedDirectoryId string                                                 `json:"shared_directory_id"`
	Timeouts          *directoryserviceshareddirectoryaccepter.TimeoutsState `json:"timeouts"`
}
