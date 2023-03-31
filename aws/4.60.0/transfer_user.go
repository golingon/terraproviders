// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	transferuser "github.com/golingon/terraproviders/aws/4.60.0/transferuser"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewTransferUser creates a new instance of [TransferUser].
func NewTransferUser(name string, args TransferUserArgs) *TransferUser {
	return &TransferUser{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*TransferUser)(nil)

// TransferUser represents the Terraform resource aws_transfer_user.
type TransferUser struct {
	Name      string
	Args      TransferUserArgs
	state     *transferUserState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [TransferUser].
func (tu *TransferUser) Type() string {
	return "aws_transfer_user"
}

// LocalName returns the local name for [TransferUser].
func (tu *TransferUser) LocalName() string {
	return tu.Name
}

// Configuration returns the configuration (args) for [TransferUser].
func (tu *TransferUser) Configuration() interface{} {
	return tu.Args
}

// DependOn is used for other resources to depend on [TransferUser].
func (tu *TransferUser) DependOn() terra.Reference {
	return terra.ReferenceResource(tu)
}

// Dependencies returns the list of resources [TransferUser] depends_on.
func (tu *TransferUser) Dependencies() terra.Dependencies {
	return tu.DependsOn
}

// LifecycleManagement returns the lifecycle block for [TransferUser].
func (tu *TransferUser) LifecycleManagement() *terra.Lifecycle {
	return tu.Lifecycle
}

// Attributes returns the attributes for [TransferUser].
func (tu *TransferUser) Attributes() transferUserAttributes {
	return transferUserAttributes{ref: terra.ReferenceResource(tu)}
}

// ImportState imports the given attribute values into [TransferUser]'s state.
func (tu *TransferUser) ImportState(av io.Reader) error {
	tu.state = &transferUserState{}
	if err := json.NewDecoder(av).Decode(tu.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", tu.Type(), tu.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [TransferUser] has state.
func (tu *TransferUser) State() (*transferUserState, bool) {
	return tu.state, tu.state != nil
}

// StateMust returns the state for [TransferUser]. Panics if the state is nil.
func (tu *TransferUser) StateMust() *transferUserState {
	if tu.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", tu.Type(), tu.LocalName()))
	}
	return tu.state
}

// TransferUserArgs contains the configurations for aws_transfer_user.
type TransferUserArgs struct {
	// HomeDirectory: string, optional
	HomeDirectory terra.StringValue `hcl:"home_directory,attr"`
	// HomeDirectoryType: string, optional
	HomeDirectoryType terra.StringValue `hcl:"home_directory_type,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Policy: string, optional
	Policy terra.StringValue `hcl:"policy,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// ServerId: string, required
	ServerId terra.StringValue `hcl:"server_id,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// UserName: string, required
	UserName terra.StringValue `hcl:"user_name,attr" validate:"required"`
	// HomeDirectoryMappings: min=0
	HomeDirectoryMappings []transferuser.HomeDirectoryMappings `hcl:"home_directory_mappings,block" validate:"min=0"`
	// PosixProfile: optional
	PosixProfile *transferuser.PosixProfile `hcl:"posix_profile,block"`
	// Timeouts: optional
	Timeouts *transferuser.Timeouts `hcl:"timeouts,block"`
}
type transferUserAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_transfer_user.
func (tu transferUserAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(tu.ref.Append("arn"))
}

// HomeDirectory returns a reference to field home_directory of aws_transfer_user.
func (tu transferUserAttributes) HomeDirectory() terra.StringValue {
	return terra.ReferenceAsString(tu.ref.Append("home_directory"))
}

// HomeDirectoryType returns a reference to field home_directory_type of aws_transfer_user.
func (tu transferUserAttributes) HomeDirectoryType() terra.StringValue {
	return terra.ReferenceAsString(tu.ref.Append("home_directory_type"))
}

// Id returns a reference to field id of aws_transfer_user.
func (tu transferUserAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(tu.ref.Append("id"))
}

// Policy returns a reference to field policy of aws_transfer_user.
func (tu transferUserAttributes) Policy() terra.StringValue {
	return terra.ReferenceAsString(tu.ref.Append("policy"))
}

// Role returns a reference to field role of aws_transfer_user.
func (tu transferUserAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(tu.ref.Append("role"))
}

// ServerId returns a reference to field server_id of aws_transfer_user.
func (tu transferUserAttributes) ServerId() terra.StringValue {
	return terra.ReferenceAsString(tu.ref.Append("server_id"))
}

// Tags returns a reference to field tags of aws_transfer_user.
func (tu transferUserAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](tu.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_transfer_user.
func (tu transferUserAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](tu.ref.Append("tags_all"))
}

// UserName returns a reference to field user_name of aws_transfer_user.
func (tu transferUserAttributes) UserName() terra.StringValue {
	return terra.ReferenceAsString(tu.ref.Append("user_name"))
}

func (tu transferUserAttributes) HomeDirectoryMappings() terra.ListValue[transferuser.HomeDirectoryMappingsAttributes] {
	return terra.ReferenceAsList[transferuser.HomeDirectoryMappingsAttributes](tu.ref.Append("home_directory_mappings"))
}

func (tu transferUserAttributes) PosixProfile() terra.ListValue[transferuser.PosixProfileAttributes] {
	return terra.ReferenceAsList[transferuser.PosixProfileAttributes](tu.ref.Append("posix_profile"))
}

func (tu transferUserAttributes) Timeouts() transferuser.TimeoutsAttributes {
	return terra.ReferenceAsSingle[transferuser.TimeoutsAttributes](tu.ref.Append("timeouts"))
}

type transferUserState struct {
	Arn                   string                                    `json:"arn"`
	HomeDirectory         string                                    `json:"home_directory"`
	HomeDirectoryType     string                                    `json:"home_directory_type"`
	Id                    string                                    `json:"id"`
	Policy                string                                    `json:"policy"`
	Role                  string                                    `json:"role"`
	ServerId              string                                    `json:"server_id"`
	Tags                  map[string]string                         `json:"tags"`
	TagsAll               map[string]string                         `json:"tags_all"`
	UserName              string                                    `json:"user_name"`
	HomeDirectoryMappings []transferuser.HomeDirectoryMappingsState `json:"home_directory_mappings"`
	PosixProfile          []transferuser.PosixProfileState          `json:"posix_profile"`
	Timeouts              *transferuser.TimeoutsState               `json:"timeouts"`
}
