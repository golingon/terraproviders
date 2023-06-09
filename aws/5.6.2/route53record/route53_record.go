// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package route53record

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Alias struct {
	// EvaluateTargetHealth: bool, required
	EvaluateTargetHealth terra.BoolValue `hcl:"evaluate_target_health,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ZoneId: string, required
	ZoneId terra.StringValue `hcl:"zone_id,attr" validate:"required"`
}

type CidrRoutingPolicy struct {
	// CollectionId: string, required
	CollectionId terra.StringValue `hcl:"collection_id,attr" validate:"required"`
	// LocationName: string, required
	LocationName terra.StringValue `hcl:"location_name,attr" validate:"required"`
}

type FailoverRoutingPolicy struct {
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
}

type GeolocationRoutingPolicy struct {
	// Continent: string, optional
	Continent terra.StringValue `hcl:"continent,attr"`
	// Country: string, optional
	Country terra.StringValue `hcl:"country,attr"`
	// Subdivision: string, optional
	Subdivision terra.StringValue `hcl:"subdivision,attr"`
}

type LatencyRoutingPolicy struct {
	// Region: string, required
	Region terra.StringValue `hcl:"region,attr" validate:"required"`
}

type WeightedRoutingPolicy struct {
	// Weight: number, required
	Weight terra.NumberValue `hcl:"weight,attr" validate:"required"`
}

type AliasAttributes struct {
	ref terra.Reference
}

func (a AliasAttributes) InternalRef() (terra.Reference, error) {
	return a.ref, nil
}

func (a AliasAttributes) InternalWithRef(ref terra.Reference) AliasAttributes {
	return AliasAttributes{ref: ref}
}

func (a AliasAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return a.ref.InternalTokens()
}

func (a AliasAttributes) EvaluateTargetHealth() terra.BoolValue {
	return terra.ReferenceAsBool(a.ref.Append("evaluate_target_health"))
}

func (a AliasAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("name"))
}

func (a AliasAttributes) ZoneId() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("zone_id"))
}

type CidrRoutingPolicyAttributes struct {
	ref terra.Reference
}

func (crp CidrRoutingPolicyAttributes) InternalRef() (terra.Reference, error) {
	return crp.ref, nil
}

func (crp CidrRoutingPolicyAttributes) InternalWithRef(ref terra.Reference) CidrRoutingPolicyAttributes {
	return CidrRoutingPolicyAttributes{ref: ref}
}

func (crp CidrRoutingPolicyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return crp.ref.InternalTokens()
}

func (crp CidrRoutingPolicyAttributes) CollectionId() terra.StringValue {
	return terra.ReferenceAsString(crp.ref.Append("collection_id"))
}

func (crp CidrRoutingPolicyAttributes) LocationName() terra.StringValue {
	return terra.ReferenceAsString(crp.ref.Append("location_name"))
}

type FailoverRoutingPolicyAttributes struct {
	ref terra.Reference
}

func (frp FailoverRoutingPolicyAttributes) InternalRef() (terra.Reference, error) {
	return frp.ref, nil
}

func (frp FailoverRoutingPolicyAttributes) InternalWithRef(ref terra.Reference) FailoverRoutingPolicyAttributes {
	return FailoverRoutingPolicyAttributes{ref: ref}
}

func (frp FailoverRoutingPolicyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return frp.ref.InternalTokens()
}

func (frp FailoverRoutingPolicyAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(frp.ref.Append("type"))
}

type GeolocationRoutingPolicyAttributes struct {
	ref terra.Reference
}

func (grp GeolocationRoutingPolicyAttributes) InternalRef() (terra.Reference, error) {
	return grp.ref, nil
}

func (grp GeolocationRoutingPolicyAttributes) InternalWithRef(ref terra.Reference) GeolocationRoutingPolicyAttributes {
	return GeolocationRoutingPolicyAttributes{ref: ref}
}

func (grp GeolocationRoutingPolicyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return grp.ref.InternalTokens()
}

func (grp GeolocationRoutingPolicyAttributes) Continent() terra.StringValue {
	return terra.ReferenceAsString(grp.ref.Append("continent"))
}

func (grp GeolocationRoutingPolicyAttributes) Country() terra.StringValue {
	return terra.ReferenceAsString(grp.ref.Append("country"))
}

func (grp GeolocationRoutingPolicyAttributes) Subdivision() terra.StringValue {
	return terra.ReferenceAsString(grp.ref.Append("subdivision"))
}

type LatencyRoutingPolicyAttributes struct {
	ref terra.Reference
}

func (lrp LatencyRoutingPolicyAttributes) InternalRef() (terra.Reference, error) {
	return lrp.ref, nil
}

func (lrp LatencyRoutingPolicyAttributes) InternalWithRef(ref terra.Reference) LatencyRoutingPolicyAttributes {
	return LatencyRoutingPolicyAttributes{ref: ref}
}

func (lrp LatencyRoutingPolicyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return lrp.ref.InternalTokens()
}

func (lrp LatencyRoutingPolicyAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(lrp.ref.Append("region"))
}

type WeightedRoutingPolicyAttributes struct {
	ref terra.Reference
}

func (wrp WeightedRoutingPolicyAttributes) InternalRef() (terra.Reference, error) {
	return wrp.ref, nil
}

func (wrp WeightedRoutingPolicyAttributes) InternalWithRef(ref terra.Reference) WeightedRoutingPolicyAttributes {
	return WeightedRoutingPolicyAttributes{ref: ref}
}

func (wrp WeightedRoutingPolicyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return wrp.ref.InternalTokens()
}

func (wrp WeightedRoutingPolicyAttributes) Weight() terra.NumberValue {
	return terra.ReferenceAsNumber(wrp.ref.Append("weight"))
}

type AliasState struct {
	EvaluateTargetHealth bool   `json:"evaluate_target_health"`
	Name                 string `json:"name"`
	ZoneId               string `json:"zone_id"`
}

type CidrRoutingPolicyState struct {
	CollectionId string `json:"collection_id"`
	LocationName string `json:"location_name"`
}

type FailoverRoutingPolicyState struct {
	Type string `json:"type"`
}

type GeolocationRoutingPolicyState struct {
	Continent   string `json:"continent"`
	Country     string `json:"country"`
	Subdivision string `json:"subdivision"`
}

type LatencyRoutingPolicyState struct {
	Region string `json:"region"`
}

type WeightedRoutingPolicyState struct {
	Weight float64 `json:"weight"`
}
