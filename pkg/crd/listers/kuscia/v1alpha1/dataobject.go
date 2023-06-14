// Copyright 2023 Ant Group Co., Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/secretflow/kuscia/pkg/crd/apis/kuscia/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// DataObjectLister helps list DataObjects.
// All objects returned here must be treated as read-only.
type DataObjectLister interface {
	// List lists all DataObjects in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.DataObject, err error)
	// DataObjects returns an object that can list and get DataObjects.
	DataObjects(namespace string) DataObjectNamespaceLister
	DataObjectListerExpansion
}

// dataObjectLister implements the DataObjectLister interface.
type dataObjectLister struct {
	indexer cache.Indexer
}

// NewDataObjectLister returns a new DataObjectLister.
func NewDataObjectLister(indexer cache.Indexer) DataObjectLister {
	return &dataObjectLister{indexer: indexer}
}

// List lists all DataObjects in the indexer.
func (s *dataObjectLister) List(selector labels.Selector) (ret []*v1alpha1.DataObject, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DataObject))
	})
	return ret, err
}

// DataObjects returns an object that can list and get DataObjects.
func (s *dataObjectLister) DataObjects(namespace string) DataObjectNamespaceLister {
	return dataObjectNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DataObjectNamespaceLister helps list and get DataObjects.
// All objects returned here must be treated as read-only.
type DataObjectNamespaceLister interface {
	// List lists all DataObjects in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.DataObject, err error)
	// Get retrieves the DataObject from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.DataObject, error)
	DataObjectNamespaceListerExpansion
}

// dataObjectNamespaceLister implements the DataObjectNamespaceLister
// interface.
type dataObjectNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all DataObjects in the indexer for a given namespace.
func (s dataObjectNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.DataObject, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DataObject))
	})
	return ret, err
}

// Get retrieves the DataObject from the indexer for a given namespace and name.
func (s dataObjectNamespaceLister) Get(name string) (*v1alpha1.DataObject, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("dataobject"), name)
	}
	return obj.(*v1alpha1.DataObject), nil
}
