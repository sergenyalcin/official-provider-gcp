// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0
// Code generated by angryjet. DO NOT EDIT.
// Code transformed by upjet. DO NOT EDIT.

package v1beta2

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	errors "github.com/pkg/errors"
	apisresolver "github.com/upbound/provider-gcp/internal/apis"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

func (mg *InstanceIAMMember) ResolveReferences( // ResolveReferences of this InstanceIAMMember.
	ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("notebooks.gcp.upbound.io", "v1beta2", "Instance", "InstanceList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.InstanceName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.InstanceNameRef,
			Selector:     mg.Spec.ForProvider.InstanceNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.InstanceName")
	}
	mg.Spec.ForProvider.InstanceName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.InstanceNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("notebooks.gcp.upbound.io", "v1beta2", "Instance", "InstanceList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.InstanceName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.InstanceNameRef,
			Selector:     mg.Spec.InitProvider.InstanceNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.InstanceName")
	}
	mg.Spec.InitProvider.InstanceName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.InstanceNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this RuntimeIAMMember.
func (mg *RuntimeIAMMember) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("notebooks.gcp.upbound.io", "v1beta2", "Runtime", "RuntimeList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RuntimeName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.RuntimeNameRef,
			Selector:     mg.Spec.ForProvider.RuntimeNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RuntimeName")
	}
	mg.Spec.ForProvider.RuntimeName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RuntimeNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("notebooks.gcp.upbound.io", "v1beta2", "Runtime", "RuntimeList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.RuntimeName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.RuntimeNameRef,
			Selector:     mg.Spec.InitProvider.RuntimeNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.RuntimeName")
	}
	mg.Spec.InitProvider.RuntimeName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.RuntimeNameRef = rsp.ResolvedReference

	return nil
}
