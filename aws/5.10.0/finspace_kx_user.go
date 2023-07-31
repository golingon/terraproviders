// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	finspacekxuser "github.com/golingon/terraproviders/aws/5.10.0/finspacekxuser"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewFinspaceKxUser creates a new instance of [FinspaceKxUser].
func NewFinspaceKxUser(name string, args FinspaceKxUserArgs) *FinspaceKxUser {
	return &FinspaceKxUser{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*FinspaceKxUser)(nil)

// FinspaceKxUser represents the Terraform resource aws_finspace_kx_user.
type FinspaceKxUser struct {
	Name      string
	Args      FinspaceKxUserArgs
	state     *finspaceKxUserState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [FinspaceKxUser].
func (fku *FinspaceKxUser) Type() string {
	return "aws_finspace_kx_user"
}

// LocalName returns the local name for [FinspaceKxUser].
func (fku *FinspaceKxUser) LocalName() string {
	return fku.Name
}

// Configuration returns the configuration (args) for [FinspaceKxUser].
func (fku *FinspaceKxUser) Configuration() interface{} {
	return fku.Args
}

// DependOn is used for other resources to depend on [FinspaceKxUser].
func (fku *FinspaceKxUser) DependOn() terra.Reference {
	return terra.ReferenceResource(fku)
}

// Dependencies returns the list of resources [FinspaceKxUser] depends_on.
func (fku *FinspaceKxUser) Dependencies() terra.Dependencies {
	return fku.DependsOn
}

// LifecycleManagement returns the lifecycle block for [FinspaceKxUser].
func (fku *FinspaceKxUser) LifecycleManagement() *terra.Lifecycle {
	return fku.Lifecycle
}

// Attributes returns the attributes for [FinspaceKxUser].
func (fku *FinspaceKxUser) Attributes() finspaceKxUserAttributes {
	return finspaceKxUserAttributes{ref: terra.ReferenceResource(fku)}
}

// ImportState imports the given attribute values into [FinspaceKxUser]'s state.
func (fku *FinspaceKxUser) ImportState(av io.Reader) error {
	fku.state = &finspaceKxUserState{}
	if err := json.NewDecoder(av).Decode(fku.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", fku.Type(), fku.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [FinspaceKxUser] has state.
func (fku *FinspaceKxUser) State() (*finspaceKxUserState, bool) {
	return fku.state, fku.state != nil
}

// StateMust returns the state for [FinspaceKxUser]. Panics if the state is nil.
func (fku *FinspaceKxUser) StateMust() *finspaceKxUserState {
	if fku.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", fku.Type(), fku.LocalName()))
	}
	return fku.state
}

// FinspaceKxUserArgs contains the configurations for aws_finspace_kx_user.
type FinspaceKxUserArgs struct {
	// EnvironmentId: string, required
	EnvironmentId terra.StringValue `hcl:"environment_id,attr" validate:"required"`
	// IamRole: string, required
	IamRole terra.StringValue `hcl:"iam_role,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// Timeouts: optional
	Timeouts *finspacekxuser.Timeouts `hcl:"timeouts,block"`
}
type finspaceKxUserAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_finspace_kx_user.
func (fku finspaceKxUserAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(fku.ref.Append("arn"))
}

// EnvironmentId returns a reference to field environment_id of aws_finspace_kx_user.
func (fku finspaceKxUserAttributes) EnvironmentId() terra.StringValue {
	return terra.ReferenceAsString(fku.ref.Append("environment_id"))
}

// IamRole returns a reference to field iam_role of aws_finspace_kx_user.
func (fku finspaceKxUserAttributes) IamRole() terra.StringValue {
	return terra.ReferenceAsString(fku.ref.Append("iam_role"))
}

// Id returns a reference to field id of aws_finspace_kx_user.
func (fku finspaceKxUserAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(fku.ref.Append("id"))
}

// Name returns a reference to field name of aws_finspace_kx_user.
func (fku finspaceKxUserAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(fku.ref.Append("name"))
}

// Tags returns a reference to field tags of aws_finspace_kx_user.
func (fku finspaceKxUserAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](fku.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_finspace_kx_user.
func (fku finspaceKxUserAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](fku.ref.Append("tags_all"))
}

func (fku finspaceKxUserAttributes) Timeouts() finspacekxuser.TimeoutsAttributes {
	return terra.ReferenceAsSingle[finspacekxuser.TimeoutsAttributes](fku.ref.Append("timeouts"))
}

type finspaceKxUserState struct {
	Arn           string                        `json:"arn"`
	EnvironmentId string                        `json:"environment_id"`
	IamRole       string                        `json:"iam_role"`
	Id            string                        `json:"id"`
	Name          string                        `json:"name"`
	Tags          map[string]string             `json:"tags"`
	TagsAll       map[string]string             `json:"tags_all"`
	Timeouts      *finspacekxuser.TimeoutsState `json:"timeouts"`
}
