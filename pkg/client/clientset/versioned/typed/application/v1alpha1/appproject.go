package v1alpha1

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"

	v1alpha1 "github.com/argoproj/argo-cd/pkg/apis/application/v1alpha1"
	scheme "github.com/argoproj/argo-cd/pkg/client/clientset/versioned/scheme"
)

// AppProjectsGetter has a method to return a AppProjectInterface.
// A group's client should implement this interface.
type AppProjectsGetter interface {
	AppProjects(namespace string) AppProjectInterface
}

// AppProjectInterface has methods to work with AppProject resources.
type AppProjectInterface interface {
	Create(*v1alpha1.AppProject) (*v1alpha1.AppProject, error)
	Update(*v1alpha1.AppProject) (*v1alpha1.AppProject, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AppProject, error)
	List(opts v1.ListOptions) (*v1alpha1.AppProjectList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AppProject, err error)
	AppProjectExpansion
}

// appProjects implements AppProjectInterface
type appProjects struct {
	client rest.Interface
	ns     string
}

// newAppProjects returns a AppProjects
func newAppProjects(c *ArgoprojV1alpha1Client, namespace string) *appProjects {
	return &appProjects{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the appProject, and returns the corresponding appProject object, and an error if there is any.
func (c *appProjects) Get(name string, options v1.GetOptions) (result *v1alpha1.AppProject, err error) {
	result = &v1alpha1.AppProject{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("appprojects").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AppProjects that match those selectors.
func (c *appProjects) List(opts v1.ListOptions) (result *v1alpha1.AppProjectList, err error) {
	result = &v1alpha1.AppProjectList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("appprojects").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested appProjects.
func (c *appProjects) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("appprojects").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a appProject and creates it.  Returns the server's representation of the appProject, and an error, if there is any.
func (c *appProjects) Create(appProject *v1alpha1.AppProject) (result *v1alpha1.AppProject, err error) {
	result = &v1alpha1.AppProject{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("appprojects").
		Body(appProject).
		Do().
		Into(result)
	return
}

// Update takes the representation of a appProject and updates it. Returns the server's representation of the appProject, and an error, if there is any.
func (c *appProjects) Update(appProject *v1alpha1.AppProject) (result *v1alpha1.AppProject, err error) {
	result = &v1alpha1.AppProject{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("appprojects").
		Name(appProject.Name).
		Body(appProject).
		Do().
		Into(result)
	return
}

// Delete takes name of the appProject and deletes it. Returns an error if one occurs.
func (c *appProjects) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("appprojects").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *appProjects) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("appprojects").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched appProject.
func (c *appProjects) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AppProject, err error) {
	result = &v1alpha1.AppProject{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("appprojects").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
