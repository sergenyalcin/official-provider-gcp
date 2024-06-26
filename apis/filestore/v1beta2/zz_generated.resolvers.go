// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0
// Code generated by angryjet. DO NOT EDIT.
// Code transformed by upjet. DO NOT EDIT.

package v1beta2

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	resource "github.com/crossplane/upjet/pkg/resource"
	errors "github.com/pkg/errors"

	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	client "sigs.k8s.io/controller-runtime/pkg/client"

	// ResolveReferences of this Instance.
	apisresolver "github.com/upbound/provider-gcp/internal/apis"
)

func (mg *Instance) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("kms.gcp.upbound.io", "v1beta2", "CryptoKey", "CryptoKeyList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.KMSKeyName),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.KMSKeyNameRef,
			Selector:     mg.Spec.ForProvider.KMSKeyNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.KMSKeyName")
	}
	mg.Spec.ForProvider.KMSKeyName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.KMSKeyNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("kms.gcp.upbound.io", "v1beta2", "CryptoKey", "CryptoKeyList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.KMSKeyName),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.KMSKeyNameRef,
			Selector:     mg.Spec.InitProvider.KMSKeyNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.KMSKeyName")
	}
	mg.Spec.InitProvider.KMSKeyName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.KMSKeyNameRef = rsp.ResolvedReference

	return nil
}
