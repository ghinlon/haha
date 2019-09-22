# [sig-storage-lib-external-provisioner - GoDoc](https://godoc.org/sigs.k8s.io/sig-storage-lib-external-provisioner)

# Links

* [GitHub - kubernetes-sigs/sig-storage-lib-external-provisioner](https://github.com/kubernetes-sigs/sig-storage-lib-external-provisioner)
* [GitHub - kubernetes-incubator/external-storage: External storage plugins, provisioners, and helper libraries](https://github.com/kubernetes-incubator/external-storage)
* [Storage Classes - Kubernetes](https://kubernetes.io/docs/concepts/storage/storage-classes/#provisioner)


# type Provisioner interface 

```go
// Provisioner is an interface that creates templates for PersistentVolumes
// and can create the volume as a new resource in the infrastructure provider.
// It can also remove the volume it created from the underlying storage
// provider.
type Provisioner interface {
	// Provision creates a volume i.e. the storage asset and returns a PV object
	// for the volume
	Provision(ProvisionOptions) (*v1.PersistentVolume, error)
	// Delete removes the storage asset that was created by Provision backing the
	// given PV. Does not delete the PV object itself.
	//
	// May return IgnoredError to indicate that the call has been ignored and no
	// action taken.
	Delete(*v1.PersistentVolume) error
}
```


# type ProvisionController struct

```go
// NewProvisionController creates a new provision controller using
// the given configuration parameters and with private (non-shared) informers.
func NewProvisionController(
	client kubernetes.Interface,
	provisionerName string,
	provisioner Provisioner,
	kubeVersion string,
	options ...func(*ProvisionController) error,
) *ProvisionController {

// processNextClaimWorkItem processes items from claimQueue
func (ctrl *ProvisionController) processNextClaimWorkItem() bool 
func (ctrl *ProvisionController) processNextVolumeWorkItem() bool
```

