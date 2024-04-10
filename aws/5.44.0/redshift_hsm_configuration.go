// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// NewRedshiftHsmConfiguration creates a new instance of [RedshiftHsmConfiguration].
func NewRedshiftHsmConfiguration(name string, args RedshiftHsmConfigurationArgs) *RedshiftHsmConfiguration {
	return &RedshiftHsmConfiguration{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*RedshiftHsmConfiguration)(nil)

// RedshiftHsmConfiguration represents the Terraform resource aws_redshift_hsm_configuration.
type RedshiftHsmConfiguration struct {
	Name      string
	Args      RedshiftHsmConfigurationArgs
	state     *redshiftHsmConfigurationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [RedshiftHsmConfiguration].
func (rhc *RedshiftHsmConfiguration) Type() string {
	return "aws_redshift_hsm_configuration"
}

// LocalName returns the local name for [RedshiftHsmConfiguration].
func (rhc *RedshiftHsmConfiguration) LocalName() string {
	return rhc.Name
}

// Configuration returns the configuration (args) for [RedshiftHsmConfiguration].
func (rhc *RedshiftHsmConfiguration) Configuration() interface{} {
	return rhc.Args
}

// DependOn is used for other resources to depend on [RedshiftHsmConfiguration].
func (rhc *RedshiftHsmConfiguration) DependOn() terra.Reference {
	return terra.ReferenceResource(rhc)
}

// Dependencies returns the list of resources [RedshiftHsmConfiguration] depends_on.
func (rhc *RedshiftHsmConfiguration) Dependencies() terra.Dependencies {
	return rhc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [RedshiftHsmConfiguration].
func (rhc *RedshiftHsmConfiguration) LifecycleManagement() *terra.Lifecycle {
	return rhc.Lifecycle
}

// Attributes returns the attributes for [RedshiftHsmConfiguration].
func (rhc *RedshiftHsmConfiguration) Attributes() redshiftHsmConfigurationAttributes {
	return redshiftHsmConfigurationAttributes{ref: terra.ReferenceResource(rhc)}
}

// ImportState imports the given attribute values into [RedshiftHsmConfiguration]'s state.
func (rhc *RedshiftHsmConfiguration) ImportState(av io.Reader) error {
	rhc.state = &redshiftHsmConfigurationState{}
	if err := json.NewDecoder(av).Decode(rhc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", rhc.Type(), rhc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [RedshiftHsmConfiguration] has state.
func (rhc *RedshiftHsmConfiguration) State() (*redshiftHsmConfigurationState, bool) {
	return rhc.state, rhc.state != nil
}

// StateMust returns the state for [RedshiftHsmConfiguration]. Panics if the state is nil.
func (rhc *RedshiftHsmConfiguration) StateMust() *redshiftHsmConfigurationState {
	if rhc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", rhc.Type(), rhc.LocalName()))
	}
	return rhc.state
}

// RedshiftHsmConfigurationArgs contains the configurations for aws_redshift_hsm_configuration.
type RedshiftHsmConfigurationArgs struct {
	// Description: string, required
	Description terra.StringValue `hcl:"description,attr" validate:"required"`
	// HsmConfigurationIdentifier: string, required
	HsmConfigurationIdentifier terra.StringValue `hcl:"hsm_configuration_identifier,attr" validate:"required"`
	// HsmIpAddress: string, required
	HsmIpAddress terra.StringValue `hcl:"hsm_ip_address,attr" validate:"required"`
	// HsmPartitionName: string, required
	HsmPartitionName terra.StringValue `hcl:"hsm_partition_name,attr" validate:"required"`
	// HsmPartitionPassword: string, required
	HsmPartitionPassword terra.StringValue `hcl:"hsm_partition_password,attr" validate:"required"`
	// HsmServerPublicCertificate: string, required
	HsmServerPublicCertificate terra.StringValue `hcl:"hsm_server_public_certificate,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
}
type redshiftHsmConfigurationAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_redshift_hsm_configuration.
func (rhc redshiftHsmConfigurationAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(rhc.ref.Append("arn"))
}

// Description returns a reference to field description of aws_redshift_hsm_configuration.
func (rhc redshiftHsmConfigurationAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(rhc.ref.Append("description"))
}

// HsmConfigurationIdentifier returns a reference to field hsm_configuration_identifier of aws_redshift_hsm_configuration.
func (rhc redshiftHsmConfigurationAttributes) HsmConfigurationIdentifier() terra.StringValue {
	return terra.ReferenceAsString(rhc.ref.Append("hsm_configuration_identifier"))
}

// HsmIpAddress returns a reference to field hsm_ip_address of aws_redshift_hsm_configuration.
func (rhc redshiftHsmConfigurationAttributes) HsmIpAddress() terra.StringValue {
	return terra.ReferenceAsString(rhc.ref.Append("hsm_ip_address"))
}

// HsmPartitionName returns a reference to field hsm_partition_name of aws_redshift_hsm_configuration.
func (rhc redshiftHsmConfigurationAttributes) HsmPartitionName() terra.StringValue {
	return terra.ReferenceAsString(rhc.ref.Append("hsm_partition_name"))
}

// HsmPartitionPassword returns a reference to field hsm_partition_password of aws_redshift_hsm_configuration.
func (rhc redshiftHsmConfigurationAttributes) HsmPartitionPassword() terra.StringValue {
	return terra.ReferenceAsString(rhc.ref.Append("hsm_partition_password"))
}

// HsmServerPublicCertificate returns a reference to field hsm_server_public_certificate of aws_redshift_hsm_configuration.
func (rhc redshiftHsmConfigurationAttributes) HsmServerPublicCertificate() terra.StringValue {
	return terra.ReferenceAsString(rhc.ref.Append("hsm_server_public_certificate"))
}

// Id returns a reference to field id of aws_redshift_hsm_configuration.
func (rhc redshiftHsmConfigurationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(rhc.ref.Append("id"))
}

// Tags returns a reference to field tags of aws_redshift_hsm_configuration.
func (rhc redshiftHsmConfigurationAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](rhc.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_redshift_hsm_configuration.
func (rhc redshiftHsmConfigurationAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](rhc.ref.Append("tags_all"))
}

type redshiftHsmConfigurationState struct {
	Arn                        string            `json:"arn"`
	Description                string            `json:"description"`
	HsmConfigurationIdentifier string            `json:"hsm_configuration_identifier"`
	HsmIpAddress               string            `json:"hsm_ip_address"`
	HsmPartitionName           string            `json:"hsm_partition_name"`
	HsmPartitionPassword       string            `json:"hsm_partition_password"`
	HsmServerPublicCertificate string            `json:"hsm_server_public_certificate"`
	Id                         string            `json:"id"`
	Tags                       map[string]string `json:"tags"`
	TagsAll                    map[string]string `json:"tags_all"`
}
