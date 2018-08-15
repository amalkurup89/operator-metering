// Code generated by informer-gen. DO NOT EDIT.

// This file was automatically generated by informer-gen

package externalversions

import (
	"fmt"

	v1alpha1 "github.com/operator-framework/operator-metering/pkg/apis/metering/v1alpha1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	cache "k8s.io/client-go/tools/cache"
)

// GenericInformer is type of SharedIndexInformer which will locate and delegate to other
// sharedInformers based on type
type GenericInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() cache.GenericLister
}

type genericInformer struct {
	informer cache.SharedIndexInformer
	resource schema.GroupResource
}

// Informer returns the SharedIndexInformer.
func (f *genericInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

// Lister returns the GenericLister.
func (f *genericInformer) Lister() cache.GenericLister {
	return cache.NewGenericLister(f.Informer().GetIndexer(), f.resource)
}

// ForResource gives generic access to a shared informer of the matching type
// TODO extend this to unknown resources with a client pool
func (f *sharedInformerFactory) ForResource(resource schema.GroupVersionResource) (GenericInformer, error) {
	switch resource {
	// Group=metering.openshift.io, Version=v1alpha1
	case v1alpha1.SchemeGroupVersion.WithResource("prestotables"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Metering().V1alpha1().PrestoTables().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("reports"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Metering().V1alpha1().Reports().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("reportdatasources"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Metering().V1alpha1().ReportDataSources().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("reportgenerationqueries"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Metering().V1alpha1().ReportGenerationQueries().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("reportprometheusqueries"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Metering().V1alpha1().ReportPrometheusQueries().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("scheduledreports"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Metering().V1alpha1().ScheduledReports().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("storagelocations"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Metering().V1alpha1().StorageLocations().Informer()}, nil

	}

	return nil, fmt.Errorf("no informer found for %v", resource)
}
