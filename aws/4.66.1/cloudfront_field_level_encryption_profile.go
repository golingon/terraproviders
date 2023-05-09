// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	cloudfrontfieldlevelencryptionprofile "github.com/golingon/terraproviders/aws/4.66.1/cloudfrontfieldlevelencryptionprofile"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCloudfrontFieldLevelEncryptionProfile creates a new instance of [CloudfrontFieldLevelEncryptionProfile].
func NewCloudfrontFieldLevelEncryptionProfile(name string, args CloudfrontFieldLevelEncryptionProfileArgs) *CloudfrontFieldLevelEncryptionProfile {
	return &CloudfrontFieldLevelEncryptionProfile{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CloudfrontFieldLevelEncryptionProfile)(nil)

// CloudfrontFieldLevelEncryptionProfile represents the Terraform resource aws_cloudfront_field_level_encryption_profile.
type CloudfrontFieldLevelEncryptionProfile struct {
	Name      string
	Args      CloudfrontFieldLevelEncryptionProfileArgs
	state     *cloudfrontFieldLevelEncryptionProfileState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CloudfrontFieldLevelEncryptionProfile].
func (cflep *CloudfrontFieldLevelEncryptionProfile) Type() string {
	return "aws_cloudfront_field_level_encryption_profile"
}

// LocalName returns the local name for [CloudfrontFieldLevelEncryptionProfile].
func (cflep *CloudfrontFieldLevelEncryptionProfile) LocalName() string {
	return cflep.Name
}

// Configuration returns the configuration (args) for [CloudfrontFieldLevelEncryptionProfile].
func (cflep *CloudfrontFieldLevelEncryptionProfile) Configuration() interface{} {
	return cflep.Args
}

// DependOn is used for other resources to depend on [CloudfrontFieldLevelEncryptionProfile].
func (cflep *CloudfrontFieldLevelEncryptionProfile) DependOn() terra.Reference {
	return terra.ReferenceResource(cflep)
}

// Dependencies returns the list of resources [CloudfrontFieldLevelEncryptionProfile] depends_on.
func (cflep *CloudfrontFieldLevelEncryptionProfile) Dependencies() terra.Dependencies {
	return cflep.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CloudfrontFieldLevelEncryptionProfile].
func (cflep *CloudfrontFieldLevelEncryptionProfile) LifecycleManagement() *terra.Lifecycle {
	return cflep.Lifecycle
}

// Attributes returns the attributes for [CloudfrontFieldLevelEncryptionProfile].
func (cflep *CloudfrontFieldLevelEncryptionProfile) Attributes() cloudfrontFieldLevelEncryptionProfileAttributes {
	return cloudfrontFieldLevelEncryptionProfileAttributes{ref: terra.ReferenceResource(cflep)}
}

// ImportState imports the given attribute values into [CloudfrontFieldLevelEncryptionProfile]'s state.
func (cflep *CloudfrontFieldLevelEncryptionProfile) ImportState(av io.Reader) error {
	cflep.state = &cloudfrontFieldLevelEncryptionProfileState{}
	if err := json.NewDecoder(av).Decode(cflep.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cflep.Type(), cflep.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CloudfrontFieldLevelEncryptionProfile] has state.
func (cflep *CloudfrontFieldLevelEncryptionProfile) State() (*cloudfrontFieldLevelEncryptionProfileState, bool) {
	return cflep.state, cflep.state != nil
}

// StateMust returns the state for [CloudfrontFieldLevelEncryptionProfile]. Panics if the state is nil.
func (cflep *CloudfrontFieldLevelEncryptionProfile) StateMust() *cloudfrontFieldLevelEncryptionProfileState {
	if cflep.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cflep.Type(), cflep.LocalName()))
	}
	return cflep.state
}

// CloudfrontFieldLevelEncryptionProfileArgs contains the configurations for aws_cloudfront_field_level_encryption_profile.
type CloudfrontFieldLevelEncryptionProfileArgs struct {
	// Comment: string, optional
	Comment terra.StringValue `hcl:"comment,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// EncryptionEntities: required
	EncryptionEntities *cloudfrontfieldlevelencryptionprofile.EncryptionEntities `hcl:"encryption_entities,block" validate:"required"`
}
type cloudfrontFieldLevelEncryptionProfileAttributes struct {
	ref terra.Reference
}

// CallerReference returns a reference to field caller_reference of aws_cloudfront_field_level_encryption_profile.
func (cflep cloudfrontFieldLevelEncryptionProfileAttributes) CallerReference() terra.StringValue {
	return terra.ReferenceAsString(cflep.ref.Append("caller_reference"))
}

// Comment returns a reference to field comment of aws_cloudfront_field_level_encryption_profile.
func (cflep cloudfrontFieldLevelEncryptionProfileAttributes) Comment() terra.StringValue {
	return terra.ReferenceAsString(cflep.ref.Append("comment"))
}

// Etag returns a reference to field etag of aws_cloudfront_field_level_encryption_profile.
func (cflep cloudfrontFieldLevelEncryptionProfileAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(cflep.ref.Append("etag"))
}

// Id returns a reference to field id of aws_cloudfront_field_level_encryption_profile.
func (cflep cloudfrontFieldLevelEncryptionProfileAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cflep.ref.Append("id"))
}

// Name returns a reference to field name of aws_cloudfront_field_level_encryption_profile.
func (cflep cloudfrontFieldLevelEncryptionProfileAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cflep.ref.Append("name"))
}

func (cflep cloudfrontFieldLevelEncryptionProfileAttributes) EncryptionEntities() terra.ListValue[cloudfrontfieldlevelencryptionprofile.EncryptionEntitiesAttributes] {
	return terra.ReferenceAsList[cloudfrontfieldlevelencryptionprofile.EncryptionEntitiesAttributes](cflep.ref.Append("encryption_entities"))
}

type cloudfrontFieldLevelEncryptionProfileState struct {
	CallerReference    string                                                          `json:"caller_reference"`
	Comment            string                                                          `json:"comment"`
	Etag               string                                                          `json:"etag"`
	Id                 string                                                          `json:"id"`
	Name               string                                                          `json:"name"`
	EncryptionEntities []cloudfrontfieldlevelencryptionprofile.EncryptionEntitiesState `json:"encryption_entities"`
}
