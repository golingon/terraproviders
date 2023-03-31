// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package cloudhsmv2cluster

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type ClusterCertificates struct{}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type ClusterCertificatesAttributes struct {
	ref terra.Reference
}

func (cc ClusterCertificatesAttributes) InternalRef() terra.Reference {
	return cc.ref
}

func (cc ClusterCertificatesAttributes) InternalWithRef(ref terra.Reference) ClusterCertificatesAttributes {
	return ClusterCertificatesAttributes{ref: ref}
}

func (cc ClusterCertificatesAttributes) InternalTokens() hclwrite.Tokens {
	return cc.ref.InternalTokens()
}

func (cc ClusterCertificatesAttributes) AwsHardwareCertificate() terra.StringValue {
	return terra.ReferenceAsString(cc.ref.Append("aws_hardware_certificate"))
}

func (cc ClusterCertificatesAttributes) ClusterCertificate() terra.StringValue {
	return terra.ReferenceAsString(cc.ref.Append("cluster_certificate"))
}

func (cc ClusterCertificatesAttributes) ClusterCsr() terra.StringValue {
	return terra.ReferenceAsString(cc.ref.Append("cluster_csr"))
}

func (cc ClusterCertificatesAttributes) HsmCertificate() terra.StringValue {
	return terra.ReferenceAsString(cc.ref.Append("hsm_certificate"))
}

func (cc ClusterCertificatesAttributes) ManufacturerHardwareCertificate() terra.StringValue {
	return terra.ReferenceAsString(cc.ref.Append("manufacturer_hardware_certificate"))
}

type TimeoutsAttributes struct {
	ref terra.Reference
}

func (t TimeoutsAttributes) InternalRef() terra.Reference {
	return t.ref
}

func (t TimeoutsAttributes) InternalWithRef(ref terra.Reference) TimeoutsAttributes {
	return TimeoutsAttributes{ref: ref}
}

func (t TimeoutsAttributes) InternalTokens() hclwrite.Tokens {
	return t.ref.InternalTokens()
}

func (t TimeoutsAttributes) Create() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("create"))
}

func (t TimeoutsAttributes) Delete() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("delete"))
}

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("update"))
}

type ClusterCertificatesState struct {
	AwsHardwareCertificate          string `json:"aws_hardware_certificate"`
	ClusterCertificate              string `json:"cluster_certificate"`
	ClusterCsr                      string `json:"cluster_csr"`
	HsmCertificate                  string `json:"hsm_certificate"`
	ManufacturerHardwareCertificate string `json:"manufacturer_hardware_certificate"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
