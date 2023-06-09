// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	binaryauthorizationattestor "github.com/golingon/terraproviders/google/4.66.0/binaryauthorizationattestor"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewBinaryAuthorizationAttestor creates a new instance of [BinaryAuthorizationAttestor].
func NewBinaryAuthorizationAttestor(name string, args BinaryAuthorizationAttestorArgs) *BinaryAuthorizationAttestor {
	return &BinaryAuthorizationAttestor{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*BinaryAuthorizationAttestor)(nil)

// BinaryAuthorizationAttestor represents the Terraform resource google_binary_authorization_attestor.
type BinaryAuthorizationAttestor struct {
	Name      string
	Args      BinaryAuthorizationAttestorArgs
	state     *binaryAuthorizationAttestorState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [BinaryAuthorizationAttestor].
func (baa *BinaryAuthorizationAttestor) Type() string {
	return "google_binary_authorization_attestor"
}

// LocalName returns the local name for [BinaryAuthorizationAttestor].
func (baa *BinaryAuthorizationAttestor) LocalName() string {
	return baa.Name
}

// Configuration returns the configuration (args) for [BinaryAuthorizationAttestor].
func (baa *BinaryAuthorizationAttestor) Configuration() interface{} {
	return baa.Args
}

// DependOn is used for other resources to depend on [BinaryAuthorizationAttestor].
func (baa *BinaryAuthorizationAttestor) DependOn() terra.Reference {
	return terra.ReferenceResource(baa)
}

// Dependencies returns the list of resources [BinaryAuthorizationAttestor] depends_on.
func (baa *BinaryAuthorizationAttestor) Dependencies() terra.Dependencies {
	return baa.DependsOn
}

// LifecycleManagement returns the lifecycle block for [BinaryAuthorizationAttestor].
func (baa *BinaryAuthorizationAttestor) LifecycleManagement() *terra.Lifecycle {
	return baa.Lifecycle
}

// Attributes returns the attributes for [BinaryAuthorizationAttestor].
func (baa *BinaryAuthorizationAttestor) Attributes() binaryAuthorizationAttestorAttributes {
	return binaryAuthorizationAttestorAttributes{ref: terra.ReferenceResource(baa)}
}

// ImportState imports the given attribute values into [BinaryAuthorizationAttestor]'s state.
func (baa *BinaryAuthorizationAttestor) ImportState(av io.Reader) error {
	baa.state = &binaryAuthorizationAttestorState{}
	if err := json.NewDecoder(av).Decode(baa.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", baa.Type(), baa.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [BinaryAuthorizationAttestor] has state.
func (baa *BinaryAuthorizationAttestor) State() (*binaryAuthorizationAttestorState, bool) {
	return baa.state, baa.state != nil
}

// StateMust returns the state for [BinaryAuthorizationAttestor]. Panics if the state is nil.
func (baa *BinaryAuthorizationAttestor) StateMust() *binaryAuthorizationAttestorState {
	if baa.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", baa.Type(), baa.LocalName()))
	}
	return baa.state
}

// BinaryAuthorizationAttestorArgs contains the configurations for google_binary_authorization_attestor.
type BinaryAuthorizationAttestorArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// AttestationAuthorityNote: required
	AttestationAuthorityNote *binaryauthorizationattestor.AttestationAuthorityNote `hcl:"attestation_authority_note,block" validate:"required"`
	// Timeouts: optional
	Timeouts *binaryauthorizationattestor.Timeouts `hcl:"timeouts,block"`
}
type binaryAuthorizationAttestorAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of google_binary_authorization_attestor.
func (baa binaryAuthorizationAttestorAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(baa.ref.Append("description"))
}

// Id returns a reference to field id of google_binary_authorization_attestor.
func (baa binaryAuthorizationAttestorAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(baa.ref.Append("id"))
}

// Name returns a reference to field name of google_binary_authorization_attestor.
func (baa binaryAuthorizationAttestorAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(baa.ref.Append("name"))
}

// Project returns a reference to field project of google_binary_authorization_attestor.
func (baa binaryAuthorizationAttestorAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(baa.ref.Append("project"))
}

func (baa binaryAuthorizationAttestorAttributes) AttestationAuthorityNote() terra.ListValue[binaryauthorizationattestor.AttestationAuthorityNoteAttributes] {
	return terra.ReferenceAsList[binaryauthorizationattestor.AttestationAuthorityNoteAttributes](baa.ref.Append("attestation_authority_note"))
}

func (baa binaryAuthorizationAttestorAttributes) Timeouts() binaryauthorizationattestor.TimeoutsAttributes {
	return terra.ReferenceAsSingle[binaryauthorizationattestor.TimeoutsAttributes](baa.ref.Append("timeouts"))
}

type binaryAuthorizationAttestorState struct {
	Description              string                                                      `json:"description"`
	Id                       string                                                      `json:"id"`
	Name                     string                                                      `json:"name"`
	Project                  string                                                      `json:"project"`
	AttestationAuthorityNote []binaryauthorizationattestor.AttestationAuthorityNoteState `json:"attestation_authority_note"`
	Timeouts                 *binaryauthorizationattestor.TimeoutsState                  `json:"timeouts"`
}
