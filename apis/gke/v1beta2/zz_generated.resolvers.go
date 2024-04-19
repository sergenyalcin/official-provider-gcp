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
	apisresolver "github.com/upbound/provider-gcp/internal/apis"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

func (mg *BackupBackupPlan) ResolveReferences( // ResolveReferences of this BackupBackupPlan.
	ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	if mg.Spec.ForProvider.BackupConfig != nil {
		if mg.Spec.ForProvider.BackupConfig.EncryptionKey != nil {
			{
				m, l, err = apisresolver.GetManagedResource("kms.gcp.upbound.io", "v1beta2", "CryptoKey", "CryptoKeyList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.BackupConfig.EncryptionKey.GCPKMSEncryptionKey),
					Extract:      resource.ExtractResourceID(),
					Reference:    mg.Spec.ForProvider.BackupConfig.EncryptionKey.GCPKMSEncryptionKeyRef,
					Selector:     mg.Spec.ForProvider.BackupConfig.EncryptionKey.GCPKMSEncryptionKeySelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.BackupConfig.EncryptionKey.GCPKMSEncryptionKey")
			}
			mg.Spec.ForProvider.BackupConfig.EncryptionKey.GCPKMSEncryptionKey = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.BackupConfig.EncryptionKey.GCPKMSEncryptionKeyRef = rsp.ResolvedReference

		}
	}
	{
		m, l, err = apisresolver.GetManagedResource("container.gcp.upbound.io", "v1beta2", "Cluster", "ClusterList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Cluster),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.ClusterRef,
			Selector:     mg.Spec.ForProvider.ClusterSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Cluster")
	}
	mg.Spec.ForProvider.Cluster = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ClusterRef = rsp.ResolvedReference

	if mg.Spec.InitProvider.BackupConfig != nil {
		if mg.Spec.InitProvider.BackupConfig.EncryptionKey != nil {
			{
				m, l, err = apisresolver.GetManagedResource("kms.gcp.upbound.io", "v1beta2", "CryptoKey", "CryptoKeyList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.BackupConfig.EncryptionKey.GCPKMSEncryptionKey),
					Extract:      resource.ExtractResourceID(),
					Reference:    mg.Spec.InitProvider.BackupConfig.EncryptionKey.GCPKMSEncryptionKeyRef,
					Selector:     mg.Spec.InitProvider.BackupConfig.EncryptionKey.GCPKMSEncryptionKeySelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.InitProvider.BackupConfig.EncryptionKey.GCPKMSEncryptionKey")
			}
			mg.Spec.InitProvider.BackupConfig.EncryptionKey.GCPKMSEncryptionKey = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.InitProvider.BackupConfig.EncryptionKey.GCPKMSEncryptionKeyRef = rsp.ResolvedReference

		}
	}
	{
		m, l, err = apisresolver.GetManagedResource("container.gcp.upbound.io", "v1beta2", "Cluster", "ClusterList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Cluster),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.ClusterRef,
			Selector:     mg.Spec.InitProvider.ClusterSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Cluster")
	}
	mg.Spec.InitProvider.Cluster = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ClusterRef = rsp.ResolvedReference

	return nil
}
