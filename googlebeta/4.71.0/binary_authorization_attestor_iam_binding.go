// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	binaryauthorizationattestoriambinding "github.com/golingon/terraproviders/googlebeta/4.71.0/binaryauthorizationattestoriambinding"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewBinaryAuthorizationAttestorIamBinding creates a new instance of [BinaryAuthorizationAttestorIamBinding].
func NewBinaryAuthorizationAttestorIamBinding(name string, args BinaryAuthorizationAttestorIamBindingArgs) *BinaryAuthorizationAttestorIamBinding {
	return &BinaryAuthorizationAttestorIamBinding{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*BinaryAuthorizationAttestorIamBinding)(nil)

// BinaryAuthorizationAttestorIamBinding represents the Terraform resource google_binary_authorization_attestor_iam_binding.
type BinaryAuthorizationAttestorIamBinding struct {
	Name      string
	Args      BinaryAuthorizationAttestorIamBindingArgs
	state     *binaryAuthorizationAttestorIamBindingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [BinaryAuthorizationAttestorIamBinding].
func (baaib *BinaryAuthorizationAttestorIamBinding) Type() string {
	return "google_binary_authorization_attestor_iam_binding"
}

// LocalName returns the local name for [BinaryAuthorizationAttestorIamBinding].
func (baaib *BinaryAuthorizationAttestorIamBinding) LocalName() string {
	return baaib.Name
}

// Configuration returns the configuration (args) for [BinaryAuthorizationAttestorIamBinding].
func (baaib *BinaryAuthorizationAttestorIamBinding) Configuration() interface{} {
	return baaib.Args
}

// DependOn is used for other resources to depend on [BinaryAuthorizationAttestorIamBinding].
func (baaib *BinaryAuthorizationAttestorIamBinding) DependOn() terra.Reference {
	return terra.ReferenceResource(baaib)
}

// Dependencies returns the list of resources [BinaryAuthorizationAttestorIamBinding] depends_on.
func (baaib *BinaryAuthorizationAttestorIamBinding) Dependencies() terra.Dependencies {
	return baaib.DependsOn
}

// LifecycleManagement returns the lifecycle block for [BinaryAuthorizationAttestorIamBinding].
func (baaib *BinaryAuthorizationAttestorIamBinding) LifecycleManagement() *terra.Lifecycle {
	return baaib.Lifecycle
}

// Attributes returns the attributes for [BinaryAuthorizationAttestorIamBinding].
func (baaib *BinaryAuthorizationAttestorIamBinding) Attributes() binaryAuthorizationAttestorIamBindingAttributes {
	return binaryAuthorizationAttestorIamBindingAttributes{ref: terra.ReferenceResource(baaib)}
}

// ImportState imports the given attribute values into [BinaryAuthorizationAttestorIamBinding]'s state.
func (baaib *BinaryAuthorizationAttestorIamBinding) ImportState(av io.Reader) error {
	baaib.state = &binaryAuthorizationAttestorIamBindingState{}
	if err := json.NewDecoder(av).Decode(baaib.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", baaib.Type(), baaib.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [BinaryAuthorizationAttestorIamBinding] has state.
func (baaib *BinaryAuthorizationAttestorIamBinding) State() (*binaryAuthorizationAttestorIamBindingState, bool) {
	return baaib.state, baaib.state != nil
}

// StateMust returns the state for [BinaryAuthorizationAttestorIamBinding]. Panics if the state is nil.
func (baaib *BinaryAuthorizationAttestorIamBinding) StateMust() *binaryAuthorizationAttestorIamBindingState {
	if baaib.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", baaib.Type(), baaib.LocalName()))
	}
	return baaib.state
}

// BinaryAuthorizationAttestorIamBindingArgs contains the configurations for google_binary_authorization_attestor_iam_binding.
type BinaryAuthorizationAttestorIamBindingArgs struct {
	// Attestor: string, required
	Attestor terra.StringValue `hcl:"attestor,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Members: set of string, required
	Members terra.SetValue[terra.StringValue] `hcl:"members,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Condition: optional
	Condition *binaryauthorizationattestoriambinding.Condition `hcl:"condition,block"`
}
type binaryAuthorizationAttestorIamBindingAttributes struct {
	ref terra.Reference
}

// Attestor returns a reference to field attestor of google_binary_authorization_attestor_iam_binding.
func (baaib binaryAuthorizationAttestorIamBindingAttributes) Attestor() terra.StringValue {
	return terra.ReferenceAsString(baaib.ref.Append("attestor"))
}

// Etag returns a reference to field etag of google_binary_authorization_attestor_iam_binding.
func (baaib binaryAuthorizationAttestorIamBindingAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(baaib.ref.Append("etag"))
}

// Id returns a reference to field id of google_binary_authorization_attestor_iam_binding.
func (baaib binaryAuthorizationAttestorIamBindingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(baaib.ref.Append("id"))
}

// Members returns a reference to field members of google_binary_authorization_attestor_iam_binding.
func (baaib binaryAuthorizationAttestorIamBindingAttributes) Members() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](baaib.ref.Append("members"))
}

// Project returns a reference to field project of google_binary_authorization_attestor_iam_binding.
func (baaib binaryAuthorizationAttestorIamBindingAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(baaib.ref.Append("project"))
}

// Role returns a reference to field role of google_binary_authorization_attestor_iam_binding.
func (baaib binaryAuthorizationAttestorIamBindingAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(baaib.ref.Append("role"))
}

func (baaib binaryAuthorizationAttestorIamBindingAttributes) Condition() terra.ListValue[binaryauthorizationattestoriambinding.ConditionAttributes] {
	return terra.ReferenceAsList[binaryauthorizationattestoriambinding.ConditionAttributes](baaib.ref.Append("condition"))
}

type binaryAuthorizationAttestorIamBindingState struct {
	Attestor  string                                                 `json:"attestor"`
	Etag      string                                                 `json:"etag"`
	Id        string                                                 `json:"id"`
	Members   []string                                               `json:"members"`
	Project   string                                                 `json:"project"`
	Role      string                                                 `json:"role"`
	Condition []binaryauthorizationattestoriambinding.ConditionState `json:"condition"`
}
