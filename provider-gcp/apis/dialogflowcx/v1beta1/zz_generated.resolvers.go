/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	resource "github.com/upbound/upjet/pkg/resource"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this EntityType.
func (mg *EntityType) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Parent),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.ParentRef,
		Selector:     mg.Spec.ForProvider.ParentSelector,
		To: reference.To{
			List:    &AgentList{},
			Managed: &Agent{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Parent")
	}
	mg.Spec.ForProvider.Parent = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ParentRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Environment.
func (mg *Environment) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Parent),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.ParentRef,
		Selector:     mg.Spec.ForProvider.ParentSelector,
		To: reference.To{
			List:    &AgentList{},
			Managed: &Agent{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Parent")
	}
	mg.Spec.ForProvider.Parent = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ParentRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.VersionConfigs); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VersionConfigs[i3].Version),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.VersionConfigs[i3].VersionRef,
			Selector:     mg.Spec.ForProvider.VersionConfigs[i3].VersionSelector,
			To: reference.To{
				List:    &VersionList{},
				Managed: &Version{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.VersionConfigs[i3].Version")
		}
		mg.Spec.ForProvider.VersionConfigs[i3].Version = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.VersionConfigs[i3].VersionRef = rsp.ResolvedReference

	}

	return nil
}

// ResolveReferences of this Flow.
func (mg *Flow) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Parent),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.ParentRef,
		Selector:     mg.Spec.ForProvider.ParentSelector,
		To: reference.To{
			List:    &AgentList{},
			Managed: &Agent{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Parent")
	}
	mg.Spec.ForProvider.Parent = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ParentRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Intent.
func (mg *Intent) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Parent),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.ParentRef,
		Selector:     mg.Spec.ForProvider.ParentSelector,
		To: reference.To{
			List:    &AgentList{},
			Managed: &Agent{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Parent")
	}
	mg.Spec.ForProvider.Parent = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ParentRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Page.
func (mg *Page) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Parent),
		Extract:      resource.ExtractParamPath("start_flow", true),
		Reference:    mg.Spec.ForProvider.ParentRef,
		Selector:     mg.Spec.ForProvider.ParentSelector,
		To: reference.To{
			List:    &AgentList{},
			Managed: &Agent{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Parent")
	}
	mg.Spec.ForProvider.Parent = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ParentRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.TransitionRoutes); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.TransitionRoutes[i3].TargetPage),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.TransitionRoutes[i3].TargetPageRef,
			Selector:     mg.Spec.ForProvider.TransitionRoutes[i3].TargetPageSelector,
			To: reference.To{
				List:    &PageList{},
				Managed: &Page{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.TransitionRoutes[i3].TargetPage")
		}
		mg.Spec.ForProvider.TransitionRoutes[i3].TargetPage = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.TransitionRoutes[i3].TargetPageRef = rsp.ResolvedReference

	}

	return nil
}

// ResolveReferences of this Version.
func (mg *Version) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Parent),
		Extract:      resource.ExtractParamPath("start_flow", true),
		Reference:    mg.Spec.ForProvider.ParentRef,
		Selector:     mg.Spec.ForProvider.ParentSelector,
		To: reference.To{
			List:    &AgentList{},
			Managed: &Agent{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Parent")
	}
	mg.Spec.ForProvider.Parent = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ParentRef = rsp.ResolvedReference

	return nil
}
