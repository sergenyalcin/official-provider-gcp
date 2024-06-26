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
	common "github.com/upbound/provider-gcp/config/common"
	client "sigs.k8s.io/controller-runtime/pkg/client"

	// ResolveReferences of this Database.
	apisresolver "github.com/upbound/provider-gcp/internal/apis"
)

func (mg *Database) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("spanner.gcp.upbound.io", "v1beta2", "Instance", "InstanceList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Instance),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.InstanceRef,
			Selector:     mg.Spec.ForProvider.InstanceSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Instance")
	}
	mg.Spec.ForProvider.Instance = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.InstanceRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this DatabaseIAMMember.
func (mg *DatabaseIAMMember) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("spanner.gcp.upbound.io", "v1beta2", "Database", "DatabaseList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Database),
			Extract:      common.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.DatabaseRef,
			Selector:     mg.Spec.ForProvider.DatabaseSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Database")
	}
	mg.Spec.ForProvider.Database = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DatabaseRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("spanner.gcp.upbound.io", "v1beta2", "Instance", "InstanceList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Instance),
			Extract:      common.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.InstanceRef,
			Selector:     mg.Spec.ForProvider.InstanceSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Instance")
	}
	mg.Spec.ForProvider.Instance = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.InstanceRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("spanner.gcp.upbound.io", "v1beta2", "Database", "DatabaseList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Database),
			Extract:      common.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.DatabaseRef,
			Selector:     mg.Spec.InitProvider.DatabaseSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Database")
	}
	mg.Spec.InitProvider.Database = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.DatabaseRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("spanner.gcp.upbound.io", "v1beta2", "Instance", "InstanceList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Instance),
			Extract:      common.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.InstanceRef,
			Selector:     mg.Spec.InitProvider.InstanceSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Instance")
	}
	mg.Spec.InitProvider.Instance = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.InstanceRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this InstanceIAMMember.
func (mg *InstanceIAMMember) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("spanner.gcp.upbound.io", "v1beta2", "Instance", "InstanceList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Instance),
			Extract:      common.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.InstanceRef,
			Selector:     mg.Spec.ForProvider.InstanceSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Instance")
	}
	mg.Spec.ForProvider.Instance = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.InstanceRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("spanner.gcp.upbound.io", "v1beta2", "Instance", "InstanceList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Instance),
			Extract:      common.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.InstanceRef,
			Selector:     mg.Spec.InitProvider.InstanceSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Instance")
	}
	mg.Spec.InitProvider.Instance = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.InstanceRef = rsp.ResolvedReference

	return nil
}
