// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewMskconnectWorkerConfiguration creates a new instance of [MskconnectWorkerConfiguration].
func NewMskconnectWorkerConfiguration(name string, args MskconnectWorkerConfigurationArgs) *MskconnectWorkerConfiguration {
	return &MskconnectWorkerConfiguration{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*MskconnectWorkerConfiguration)(nil)

// MskconnectWorkerConfiguration represents the Terraform resource aws_mskconnect_worker_configuration.
type MskconnectWorkerConfiguration struct {
	Name      string
	Args      MskconnectWorkerConfigurationArgs
	state     *mskconnectWorkerConfigurationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [MskconnectWorkerConfiguration].
func (mwc *MskconnectWorkerConfiguration) Type() string {
	return "aws_mskconnect_worker_configuration"
}

// LocalName returns the local name for [MskconnectWorkerConfiguration].
func (mwc *MskconnectWorkerConfiguration) LocalName() string {
	return mwc.Name
}

// Configuration returns the configuration (args) for [MskconnectWorkerConfiguration].
func (mwc *MskconnectWorkerConfiguration) Configuration() interface{} {
	return mwc.Args
}

// DependOn is used for other resources to depend on [MskconnectWorkerConfiguration].
func (mwc *MskconnectWorkerConfiguration) DependOn() terra.Reference {
	return terra.ReferenceResource(mwc)
}

// Dependencies returns the list of resources [MskconnectWorkerConfiguration] depends_on.
func (mwc *MskconnectWorkerConfiguration) Dependencies() terra.Dependencies {
	return mwc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [MskconnectWorkerConfiguration].
func (mwc *MskconnectWorkerConfiguration) LifecycleManagement() *terra.Lifecycle {
	return mwc.Lifecycle
}

// Attributes returns the attributes for [MskconnectWorkerConfiguration].
func (mwc *MskconnectWorkerConfiguration) Attributes() mskconnectWorkerConfigurationAttributes {
	return mskconnectWorkerConfigurationAttributes{ref: terra.ReferenceResource(mwc)}
}

// ImportState imports the given attribute values into [MskconnectWorkerConfiguration]'s state.
func (mwc *MskconnectWorkerConfiguration) ImportState(av io.Reader) error {
	mwc.state = &mskconnectWorkerConfigurationState{}
	if err := json.NewDecoder(av).Decode(mwc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", mwc.Type(), mwc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [MskconnectWorkerConfiguration] has state.
func (mwc *MskconnectWorkerConfiguration) State() (*mskconnectWorkerConfigurationState, bool) {
	return mwc.state, mwc.state != nil
}

// StateMust returns the state for [MskconnectWorkerConfiguration]. Panics if the state is nil.
func (mwc *MskconnectWorkerConfiguration) StateMust() *mskconnectWorkerConfigurationState {
	if mwc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", mwc.Type(), mwc.LocalName()))
	}
	return mwc.state
}

// MskconnectWorkerConfigurationArgs contains the configurations for aws_mskconnect_worker_configuration.
type MskconnectWorkerConfigurationArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PropertiesFileContent: string, required
	PropertiesFileContent terra.StringValue `hcl:"properties_file_content,attr" validate:"required"`
}
type mskconnectWorkerConfigurationAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_mskconnect_worker_configuration.
func (mwc mskconnectWorkerConfigurationAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(mwc.ref.Append("arn"))
}

// Description returns a reference to field description of aws_mskconnect_worker_configuration.
func (mwc mskconnectWorkerConfigurationAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(mwc.ref.Append("description"))
}

// Id returns a reference to field id of aws_mskconnect_worker_configuration.
func (mwc mskconnectWorkerConfigurationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(mwc.ref.Append("id"))
}

// LatestRevision returns a reference to field latest_revision of aws_mskconnect_worker_configuration.
func (mwc mskconnectWorkerConfigurationAttributes) LatestRevision() terra.NumberValue {
	return terra.ReferenceAsNumber(mwc.ref.Append("latest_revision"))
}

// Name returns a reference to field name of aws_mskconnect_worker_configuration.
func (mwc mskconnectWorkerConfigurationAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(mwc.ref.Append("name"))
}

// PropertiesFileContent returns a reference to field properties_file_content of aws_mskconnect_worker_configuration.
func (mwc mskconnectWorkerConfigurationAttributes) PropertiesFileContent() terra.StringValue {
	return terra.ReferenceAsString(mwc.ref.Append("properties_file_content"))
}

type mskconnectWorkerConfigurationState struct {
	Arn                   string  `json:"arn"`
	Description           string  `json:"description"`
	Id                    string  `json:"id"`
	LatestRevision        float64 `json:"latest_revision"`
	Name                  string  `json:"name"`
	PropertiesFileContent string  `json:"properties_file_content"`
}