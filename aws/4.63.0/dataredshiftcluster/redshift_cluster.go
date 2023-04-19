// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package dataredshiftcluster

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type ClusterNodes struct{}

type ClusterNodesAttributes struct {
	ref terra.Reference
}

func (cn ClusterNodesAttributes) InternalRef() (terra.Reference, error) {
	return cn.ref, nil
}

func (cn ClusterNodesAttributes) InternalWithRef(ref terra.Reference) ClusterNodesAttributes {
	return ClusterNodesAttributes{ref: ref}
}

func (cn ClusterNodesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return cn.ref.InternalTokens()
}

func (cn ClusterNodesAttributes) NodeRole() terra.StringValue {
	return terra.ReferenceAsString(cn.ref.Append("node_role"))
}

func (cn ClusterNodesAttributes) PrivateIpAddress() terra.StringValue {
	return terra.ReferenceAsString(cn.ref.Append("private_ip_address"))
}

func (cn ClusterNodesAttributes) PublicIpAddress() terra.StringValue {
	return terra.ReferenceAsString(cn.ref.Append("public_ip_address"))
}

type ClusterNodesState struct {
	NodeRole         string `json:"node_role"`
	PrivateIpAddress string `json:"private_ip_address"`
	PublicIpAddress  string `json:"public_ip_address"`
}
