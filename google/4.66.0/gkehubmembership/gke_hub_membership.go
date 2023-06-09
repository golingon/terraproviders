// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package gkehubmembership

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Authority struct {
	// Issuer: string, required
	Issuer terra.StringValue `hcl:"issuer,attr" validate:"required"`
}

type Endpoint struct {
	// GkeCluster: optional
	GkeCluster *GkeCluster `hcl:"gke_cluster,block"`
}

type GkeCluster struct {
	// ResourceLink: string, required
	ResourceLink terra.StringValue `hcl:"resource_link,attr" validate:"required"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type AuthorityAttributes struct {
	ref terra.Reference
}

func (a AuthorityAttributes) InternalRef() (terra.Reference, error) {
	return a.ref, nil
}

func (a AuthorityAttributes) InternalWithRef(ref terra.Reference) AuthorityAttributes {
	return AuthorityAttributes{ref: ref}
}

func (a AuthorityAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return a.ref.InternalTokens()
}

func (a AuthorityAttributes) Issuer() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("issuer"))
}

type EndpointAttributes struct {
	ref terra.Reference
}

func (e EndpointAttributes) InternalRef() (terra.Reference, error) {
	return e.ref, nil
}

func (e EndpointAttributes) InternalWithRef(ref terra.Reference) EndpointAttributes {
	return EndpointAttributes{ref: ref}
}

func (e EndpointAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return e.ref.InternalTokens()
}

func (e EndpointAttributes) GkeCluster() terra.ListValue[GkeClusterAttributes] {
	return terra.ReferenceAsList[GkeClusterAttributes](e.ref.Append("gke_cluster"))
}

type GkeClusterAttributes struct {
	ref terra.Reference
}

func (gc GkeClusterAttributes) InternalRef() (terra.Reference, error) {
	return gc.ref, nil
}

func (gc GkeClusterAttributes) InternalWithRef(ref terra.Reference) GkeClusterAttributes {
	return GkeClusterAttributes{ref: ref}
}

func (gc GkeClusterAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return gc.ref.InternalTokens()
}

func (gc GkeClusterAttributes) ResourceLink() terra.StringValue {
	return terra.ReferenceAsString(gc.ref.Append("resource_link"))
}

type TimeoutsAttributes struct {
	ref terra.Reference
}

func (t TimeoutsAttributes) InternalRef() (terra.Reference, error) {
	return t.ref, nil
}

func (t TimeoutsAttributes) InternalWithRef(ref terra.Reference) TimeoutsAttributes {
	return TimeoutsAttributes{ref: ref}
}

func (t TimeoutsAttributes) InternalTokens() (hclwrite.Tokens, error) {
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

type AuthorityState struct {
	Issuer string `json:"issuer"`
}

type EndpointState struct {
	GkeCluster []GkeClusterState `json:"gke_cluster"`
}

type GkeClusterState struct {
	ResourceLink string `json:"resource_link"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
