// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewXrayEncryptionConfig creates a new instance of [XrayEncryptionConfig].
func NewXrayEncryptionConfig(name string, args XrayEncryptionConfigArgs) *XrayEncryptionConfig {
	return &XrayEncryptionConfig{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*XrayEncryptionConfig)(nil)

// XrayEncryptionConfig represents the Terraform resource aws_xray_encryption_config.
type XrayEncryptionConfig struct {
	Name      string
	Args      XrayEncryptionConfigArgs
	state     *xrayEncryptionConfigState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [XrayEncryptionConfig].
func (xec *XrayEncryptionConfig) Type() string {
	return "aws_xray_encryption_config"
}

// LocalName returns the local name for [XrayEncryptionConfig].
func (xec *XrayEncryptionConfig) LocalName() string {
	return xec.Name
}

// Configuration returns the configuration (args) for [XrayEncryptionConfig].
func (xec *XrayEncryptionConfig) Configuration() interface{} {
	return xec.Args
}

// DependOn is used for other resources to depend on [XrayEncryptionConfig].
func (xec *XrayEncryptionConfig) DependOn() terra.Reference {
	return terra.ReferenceResource(xec)
}

// Dependencies returns the list of resources [XrayEncryptionConfig] depends_on.
func (xec *XrayEncryptionConfig) Dependencies() terra.Dependencies {
	return xec.DependsOn
}

// LifecycleManagement returns the lifecycle block for [XrayEncryptionConfig].
func (xec *XrayEncryptionConfig) LifecycleManagement() *terra.Lifecycle {
	return xec.Lifecycle
}

// Attributes returns the attributes for [XrayEncryptionConfig].
func (xec *XrayEncryptionConfig) Attributes() xrayEncryptionConfigAttributes {
	return xrayEncryptionConfigAttributes{ref: terra.ReferenceResource(xec)}
}

// ImportState imports the given attribute values into [XrayEncryptionConfig]'s state.
func (xec *XrayEncryptionConfig) ImportState(av io.Reader) error {
	xec.state = &xrayEncryptionConfigState{}
	if err := json.NewDecoder(av).Decode(xec.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", xec.Type(), xec.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [XrayEncryptionConfig] has state.
func (xec *XrayEncryptionConfig) State() (*xrayEncryptionConfigState, bool) {
	return xec.state, xec.state != nil
}

// StateMust returns the state for [XrayEncryptionConfig]. Panics if the state is nil.
func (xec *XrayEncryptionConfig) StateMust() *xrayEncryptionConfigState {
	if xec.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", xec.Type(), xec.LocalName()))
	}
	return xec.state
}

// XrayEncryptionConfigArgs contains the configurations for aws_xray_encryption_config.
type XrayEncryptionConfigArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// KeyId: string, optional
	KeyId terra.StringValue `hcl:"key_id,attr"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
}
type xrayEncryptionConfigAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_xray_encryption_config.
func (xec xrayEncryptionConfigAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(xec.ref.Append("id"))
}

// KeyId returns a reference to field key_id of aws_xray_encryption_config.
func (xec xrayEncryptionConfigAttributes) KeyId() terra.StringValue {
	return terra.ReferenceAsString(xec.ref.Append("key_id"))
}

// Type returns a reference to field type of aws_xray_encryption_config.
func (xec xrayEncryptionConfigAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(xec.ref.Append("type"))
}

type xrayEncryptionConfigState struct {
	Id    string `json:"id"`
	KeyId string `json:"key_id"`
	Type  string `json:"type"`
}