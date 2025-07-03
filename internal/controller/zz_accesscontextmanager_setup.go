// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	accesslevel "github.com/upbound/provider-gcp/internal/controller/accesscontextmanager/accesslevel"
	accesslevelcondition "github.com/upbound/provider-gcp/internal/controller/accesscontextmanager/accesslevelcondition"
	accesspolicy "github.com/upbound/provider-gcp/internal/controller/accesscontextmanager/accesspolicy"
	accesspolicyiammember "github.com/upbound/provider-gcp/internal/controller/accesscontextmanager/accesspolicyiammember"
	serviceperimeter "github.com/upbound/provider-gcp/internal/controller/accesscontextmanager/serviceperimeter"
	serviceperimeterresource "github.com/upbound/provider-gcp/internal/controller/accesscontextmanager/serviceperimeterresource"
	"github.com/upbound/provider-gcp/internal/controller/lazyloading"
)

// Setup_accesscontextmanager creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_accesscontextmanager(mgr ctrl.Manager, o controller.Options) error {
	crdToSetupFn := map[schema.GroupKind]func(ctrl.Manager, controller.Options) error{
		schema.GroupKind{Group: "accesscontextmanager.gcp.upbound.io", Kind: "AccessLevel"}:              accesslevel.Setup,
		schema.GroupKind{Group: "accesscontextmanager.gcp.upbound.io", Kind: "AccessLevelCondition"}:     accesslevelcondition.Setup,
		schema.GroupKind{Group: "accesscontextmanager.gcp.upbound.io", Kind: "AccessPolicy"}:             accesspolicy.Setup,
		schema.GroupKind{Group: "accesscontextmanager.gcp.upbound.io", Kind: "AccessPolicyIAMMember"}:    accesspolicyiammember.Setup,
		schema.GroupKind{Group: "accesscontextmanager.gcp.upbound.io", Kind: "ServicePerimeter"}:         serviceperimeter.Setup,
		schema.GroupKind{Group: "accesscontextmanager.gcp.upbound.io", Kind: "ServicePerimeterResource"}: serviceperimeterresource.Setup,
	}
	if err := lazyloading.Setup(mgr, crdToSetupFn, o); err != nil {
		return err
	}
	return nil
}
