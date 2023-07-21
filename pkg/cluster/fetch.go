package cluster

import (
	"context"
	"fmt"

	eksdv1alpha1 "github.com/aws/eks-distro-build-tooling/release/api/v1alpha1"

	"github.com/aws/eks-anywhere/pkg/api/v1alpha1"
	"github.com/aws/eks-anywhere/pkg/constants"
	v1alpha1release "github.com/aws/eks-anywhere/release/api/v1alpha1"
)

type (
	// BundlesFetcher is used to fetch a bundles resource.
	BundlesFetcher func(string, string) (*v1alpha1release.Bundles, error)

	// EKSAReleaseFetcher is used to fetch a bundles resource.
	EKSAReleaseFetcher func(string) (*v1alpha1release.EKSARelease, error)

	// Fetchers contains fetch functions for resources.
	Fetchers struct {
		BundlesFetcher     BundlesFetcher
		EKSAReleaseFetcher EKSAReleaseFetcher
	}
)

func bundlesNamespacedKey(cluster *v1alpha1.Cluster, release *v1alpha1release.EKSARelease) (name, namespace string) {
	if release != nil && release.Spec.BundlesRef.Name != "" {
		name = release.Spec.BundlesRef.Name
		namespace = release.Spec.BundlesRef.Namespace
	} else if cluster.Spec.BundlesRef != nil {
		name = cluster.Spec.BundlesRef.Name
		namespace = cluster.Spec.BundlesRef.Namespace
	} else {
		// Handles old clusters that don't contain a reference yet to the Bundles
		// For those clusters, the Bundles was created with the same name as the cluster
		// and in the same namespace
		name = cluster.Name
		namespace = cluster.Namespace
	}

	return name, namespace
}

func GetVersionsBundle(clusterConfig *v1alpha1.Cluster, bundles *v1alpha1release.Bundles) (*v1alpha1release.VersionsBundle, error) {
	return getVersionsBundleForKubernetesVersion(clusterConfig.Spec.KubernetesVersion, bundles)
}

func getVersionsBundleForKubernetesVersion(kubernetesVersion v1alpha1.KubernetesVersion, bundles *v1alpha1release.Bundles) (*v1alpha1release.VersionsBundle, error) {
	for _, versionsBundle := range bundles.Spec.VersionsBundles {
		if versionsBundle.KubeVersion == string(kubernetesVersion) {
			return &versionsBundle, nil
		}
	}
	return nil, fmt.Errorf("kubernetes version %s is not supported by bundles manifest %d", kubernetesVersion, bundles.Spec.Number)
}

// BuildSpec constructs a cluster.Spec for an eks-a cluster by retrieving all
// necessary objects from the cluster using a kubernetes client.
func BuildSpec(ctx context.Context, client Client, cluster *v1alpha1.Cluster) (*Spec, error) {
	configBuilder := NewDefaultConfigClientBuilder()
	config, err := configBuilder.Build(ctx, client, cluster)
	if err != nil {
		return nil, err
	}

	return BuildSpecFromConfig(ctx, client, config)
}

// GetBundlesAndEKSAReleaseFromCluster gets the Bundles and EKSARelease resources from the cluster.
func GetBundlesAndEKSAReleaseFromCluster(ctx context.Context, cluster *v1alpha1.Cluster, fetchers Fetchers) (*v1alpha1release.Bundles, *v1alpha1release.EKSARelease, error) {
	var eksaRelease *v1alpha1release.EKSARelease
	var err error
	if cluster.Spec.BundlesRef == nil {
		if cluster.Spec.EksaVersion == nil {
			return nil, nil, fmt.Errorf("either cluster's EksaVersion or BundlesRef need to be set")
		}
		version := string(*cluster.Spec.EksaVersion)
		eksaReleaseName := v1alpha1release.GenerateEKSAReleaseName(version)
		eksaRelease, err = fetchers.EKSAReleaseFetcher(eksaReleaseName)
		if err != nil {
			return nil, nil, fmt.Errorf("error getting EKSARelease %s", eksaReleaseName)
		}
	}

	bundlesName, bundlesNamespace := bundlesNamespacedKey(cluster, eksaRelease)
	bundles, err := fetchers.BundlesFetcher(bundlesName, bundlesNamespace)
	if err != nil {
		return nil, nil, err
	}

	return bundles, eksaRelease, nil
}

// BuildSpecFromConfig constructs a cluster.Spec for an eks-a cluster config by retrieving all dependencies objects from the cluster using a kubernetes client.
func BuildSpecFromConfig(ctx context.Context, client Client, config *Config) (*Spec, error) {
	fetchers := buildFetchers(ctx, client)
	bundles, eksaRelease, err := GetBundlesAndEKSAReleaseFromCluster(ctx, config.Cluster, fetchers)
	if err != nil {
		return nil, err
	}

	versionsBundle, err := GetVersionsBundle(config.Cluster, bundles)
	if err != nil {
		return nil, err
	}

	// Ideally we would use the same namespace as the Bundles, but Bundles can be in any namespace and
	// the eksd release is always in eksa-system
	eksdRelease := &eksdv1alpha1.Release{}
	if err = client.Get(ctx, versionsBundle.EksD.Name, constants.EksaSystemNamespace, eksdRelease); err != nil {
		return nil, err
	}

	return NewSpec(config, bundles, eksdRelease, eksaRelease)
}

func buildFetchers(ctx context.Context, client Client) Fetchers {
	return Fetchers{
		BundlesFetcher:     bundleFetcher(ctx, client),
		EKSAReleaseFetcher: eksaReleaseFetcher(ctx, client),
	}
}

func bundleFetcher(ctx context.Context, client Client) BundlesFetcher {
	return func(name, ns string) (*v1alpha1release.Bundles, error) {
		bundles := &v1alpha1release.Bundles{}
		err := client.Get(ctx, name, ns, bundles)
		if err != nil {
			return nil, err
		}
		return bundles, nil
	}
}

func eksaReleaseFetcher(ctx context.Context, client Client) EKSAReleaseFetcher {
	return func(name string) (*v1alpha1release.EKSARelease, error) {
		eksaRelease := &v1alpha1release.EKSARelease{}
		err := client.Get(ctx, name, constants.EksaSystemNamespace, eksaRelease)
		if err != nil {
			return nil, err
		}
		return eksaRelease, nil
	}
}
