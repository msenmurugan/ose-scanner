// This file was automatically generated by lister-gen

package v1

import (
	v1 "github.com/openshift/api/apps/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// DeploymentConfigLister helps list DeploymentConfigs.
type DeploymentConfigLister interface {
	// List lists all DeploymentConfigs in the indexer.
	List(selector labels.Selector) (ret []*v1.DeploymentConfig, err error)
	// DeploymentConfigs returns an object that can list and get DeploymentConfigs.
	DeploymentConfigs(namespace string) DeploymentConfigNamespaceLister
	DeploymentConfigListerExpansion
}

// deploymentConfigLister implements the DeploymentConfigLister interface.
type deploymentConfigLister struct {
	indexer cache.Indexer
}

// NewDeploymentConfigLister returns a new DeploymentConfigLister.
func NewDeploymentConfigLister(indexer cache.Indexer) DeploymentConfigLister {
	return &deploymentConfigLister{indexer: indexer}
}

// List lists all DeploymentConfigs in the indexer.
func (s *deploymentConfigLister) List(selector labels.Selector) (ret []*v1.DeploymentConfig, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.DeploymentConfig))
	})
	return ret, err
}

// DeploymentConfigs returns an object that can list and get DeploymentConfigs.
func (s *deploymentConfigLister) DeploymentConfigs(namespace string) DeploymentConfigNamespaceLister {
	return deploymentConfigNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DeploymentConfigNamespaceLister helps list and get DeploymentConfigs.
type DeploymentConfigNamespaceLister interface {
	// List lists all DeploymentConfigs in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.DeploymentConfig, err error)
	// Get retrieves the DeploymentConfig from the indexer for a given namespace and name.
	Get(name string) (*v1.DeploymentConfig, error)
	DeploymentConfigNamespaceListerExpansion
}

// deploymentConfigNamespaceLister implements the DeploymentConfigNamespaceLister
// interface.
type deploymentConfigNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all DeploymentConfigs in the indexer for a given namespace.
func (s deploymentConfigNamespaceLister) List(selector labels.Selector) (ret []*v1.DeploymentConfig, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.DeploymentConfig))
	})
	return ret, err
}

// Get retrieves the DeploymentConfig from the indexer for a given namespace and name.
func (s deploymentConfigNamespaceLister) Get(name string) (*v1.DeploymentConfig, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("deploymentconfig"), name)
	}
	return obj.(*v1.DeploymentConfig), nil
}
