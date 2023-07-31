// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	ekscluster "github.com/golingon/terraproviders/aws/5.10.0/ekscluster"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewEksCluster creates a new instance of [EksCluster].
func NewEksCluster(name string, args EksClusterArgs) *EksCluster {
	return &EksCluster{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*EksCluster)(nil)

// EksCluster represents the Terraform resource aws_eks_cluster.
type EksCluster struct {
	Name      string
	Args      EksClusterArgs
	state     *eksClusterState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [EksCluster].
func (ec *EksCluster) Type() string {
	return "aws_eks_cluster"
}

// LocalName returns the local name for [EksCluster].
func (ec *EksCluster) LocalName() string {
	return ec.Name
}

// Configuration returns the configuration (args) for [EksCluster].
func (ec *EksCluster) Configuration() interface{} {
	return ec.Args
}

// DependOn is used for other resources to depend on [EksCluster].
func (ec *EksCluster) DependOn() terra.Reference {
	return terra.ReferenceResource(ec)
}

// Dependencies returns the list of resources [EksCluster] depends_on.
func (ec *EksCluster) Dependencies() terra.Dependencies {
	return ec.DependsOn
}

// LifecycleManagement returns the lifecycle block for [EksCluster].
func (ec *EksCluster) LifecycleManagement() *terra.Lifecycle {
	return ec.Lifecycle
}

// Attributes returns the attributes for [EksCluster].
func (ec *EksCluster) Attributes() eksClusterAttributes {
	return eksClusterAttributes{ref: terra.ReferenceResource(ec)}
}

// ImportState imports the given attribute values into [EksCluster]'s state.
func (ec *EksCluster) ImportState(av io.Reader) error {
	ec.state = &eksClusterState{}
	if err := json.NewDecoder(av).Decode(ec.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ec.Type(), ec.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [EksCluster] has state.
func (ec *EksCluster) State() (*eksClusterState, bool) {
	return ec.state, ec.state != nil
}

// StateMust returns the state for [EksCluster]. Panics if the state is nil.
func (ec *EksCluster) StateMust() *eksClusterState {
	if ec.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ec.Type(), ec.LocalName()))
	}
	return ec.state
}

// EksClusterArgs contains the configurations for aws_eks_cluster.
type EksClusterArgs struct {
	// EnabledClusterLogTypes: set of string, optional
	EnabledClusterLogTypes terra.SetValue[terra.StringValue] `hcl:"enabled_cluster_log_types,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// RoleArn: string, required
	RoleArn terra.StringValue `hcl:"role_arn,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// Version: string, optional
	Version terra.StringValue `hcl:"version,attr"`
	// CertificateAuthority: min=0
	CertificateAuthority []ekscluster.CertificateAuthority `hcl:"certificate_authority,block" validate:"min=0"`
	// Identity: min=0
	Identity []ekscluster.Identity `hcl:"identity,block" validate:"min=0"`
	// EncryptionConfig: optional
	EncryptionConfig *ekscluster.EncryptionConfig `hcl:"encryption_config,block"`
	// KubernetesNetworkConfig: optional
	KubernetesNetworkConfig *ekscluster.KubernetesNetworkConfig `hcl:"kubernetes_network_config,block"`
	// OutpostConfig: optional
	OutpostConfig *ekscluster.OutpostConfig `hcl:"outpost_config,block"`
	// Timeouts: optional
	Timeouts *ekscluster.Timeouts `hcl:"timeouts,block"`
	// VpcConfig: required
	VpcConfig *ekscluster.VpcConfig `hcl:"vpc_config,block" validate:"required"`
}
type eksClusterAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_eks_cluster.
func (ec eksClusterAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ec.ref.Append("arn"))
}

// ClusterId returns a reference to field cluster_id of aws_eks_cluster.
func (ec eksClusterAttributes) ClusterId() terra.StringValue {
	return terra.ReferenceAsString(ec.ref.Append("cluster_id"))
}

// CreatedAt returns a reference to field created_at of aws_eks_cluster.
func (ec eksClusterAttributes) CreatedAt() terra.StringValue {
	return terra.ReferenceAsString(ec.ref.Append("created_at"))
}

// EnabledClusterLogTypes returns a reference to field enabled_cluster_log_types of aws_eks_cluster.
func (ec eksClusterAttributes) EnabledClusterLogTypes() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ec.ref.Append("enabled_cluster_log_types"))
}

// Endpoint returns a reference to field endpoint of aws_eks_cluster.
func (ec eksClusterAttributes) Endpoint() terra.StringValue {
	return terra.ReferenceAsString(ec.ref.Append("endpoint"))
}

// Id returns a reference to field id of aws_eks_cluster.
func (ec eksClusterAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ec.ref.Append("id"))
}

// Name returns a reference to field name of aws_eks_cluster.
func (ec eksClusterAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ec.ref.Append("name"))
}

// PlatformVersion returns a reference to field platform_version of aws_eks_cluster.
func (ec eksClusterAttributes) PlatformVersion() terra.StringValue {
	return terra.ReferenceAsString(ec.ref.Append("platform_version"))
}

// RoleArn returns a reference to field role_arn of aws_eks_cluster.
func (ec eksClusterAttributes) RoleArn() terra.StringValue {
	return terra.ReferenceAsString(ec.ref.Append("role_arn"))
}

// Status returns a reference to field status of aws_eks_cluster.
func (ec eksClusterAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(ec.ref.Append("status"))
}

// Tags returns a reference to field tags of aws_eks_cluster.
func (ec eksClusterAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ec.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_eks_cluster.
func (ec eksClusterAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ec.ref.Append("tags_all"))
}

// Version returns a reference to field version of aws_eks_cluster.
func (ec eksClusterAttributes) Version() terra.StringValue {
	return terra.ReferenceAsString(ec.ref.Append("version"))
}

func (ec eksClusterAttributes) CertificateAuthority() terra.ListValue[ekscluster.CertificateAuthorityAttributes] {
	return terra.ReferenceAsList[ekscluster.CertificateAuthorityAttributes](ec.ref.Append("certificate_authority"))
}

func (ec eksClusterAttributes) Identity() terra.ListValue[ekscluster.IdentityAttributes] {
	return terra.ReferenceAsList[ekscluster.IdentityAttributes](ec.ref.Append("identity"))
}

func (ec eksClusterAttributes) EncryptionConfig() terra.ListValue[ekscluster.EncryptionConfigAttributes] {
	return terra.ReferenceAsList[ekscluster.EncryptionConfigAttributes](ec.ref.Append("encryption_config"))
}

func (ec eksClusterAttributes) KubernetesNetworkConfig() terra.ListValue[ekscluster.KubernetesNetworkConfigAttributes] {
	return terra.ReferenceAsList[ekscluster.KubernetesNetworkConfigAttributes](ec.ref.Append("kubernetes_network_config"))
}

func (ec eksClusterAttributes) OutpostConfig() terra.ListValue[ekscluster.OutpostConfigAttributes] {
	return terra.ReferenceAsList[ekscluster.OutpostConfigAttributes](ec.ref.Append("outpost_config"))
}

func (ec eksClusterAttributes) Timeouts() ekscluster.TimeoutsAttributes {
	return terra.ReferenceAsSingle[ekscluster.TimeoutsAttributes](ec.ref.Append("timeouts"))
}

func (ec eksClusterAttributes) VpcConfig() terra.ListValue[ekscluster.VpcConfigAttributes] {
	return terra.ReferenceAsList[ekscluster.VpcConfigAttributes](ec.ref.Append("vpc_config"))
}

type eksClusterState struct {
	Arn                     string                                    `json:"arn"`
	ClusterId               string                                    `json:"cluster_id"`
	CreatedAt               string                                    `json:"created_at"`
	EnabledClusterLogTypes  []string                                  `json:"enabled_cluster_log_types"`
	Endpoint                string                                    `json:"endpoint"`
	Id                      string                                    `json:"id"`
	Name                    string                                    `json:"name"`
	PlatformVersion         string                                    `json:"platform_version"`
	RoleArn                 string                                    `json:"role_arn"`
	Status                  string                                    `json:"status"`
	Tags                    map[string]string                         `json:"tags"`
	TagsAll                 map[string]string                         `json:"tags_all"`
	Version                 string                                    `json:"version"`
	CertificateAuthority    []ekscluster.CertificateAuthorityState    `json:"certificate_authority"`
	Identity                []ekscluster.IdentityState                `json:"identity"`
	EncryptionConfig        []ekscluster.EncryptionConfigState        `json:"encryption_config"`
	KubernetesNetworkConfig []ekscluster.KubernetesNetworkConfigState `json:"kubernetes_network_config"`
	OutpostConfig           []ekscluster.OutpostConfigState           `json:"outpost_config"`
	Timeouts                *ekscluster.TimeoutsState                 `json:"timeouts"`
	VpcConfig               []ekscluster.VpcConfigState               `json:"vpc_config"`
}
