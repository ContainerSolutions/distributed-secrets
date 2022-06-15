package kubernetes

import (
	"context"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	distributedsecretsv1alpha1 "github.com/DeeAjayi/distributed-secrets/api/v1alpha1"
)

func FetchSecret(ctx context.Context, kubeClient client.Client, disSecret *distributedsecretsv1alpha1.DistributedSecrets) (*corev1.Secret, error) {
	secret := &corev1.Secret{}
	err := kubeClient.Get(ctx, types.NamespacedName{
		Namespace: disSecret.Namespace,
		Name:      disSecret.Spec.SecretRef.Name,
	}, secret)
	return secret, err
}

func CreateSecret(ctx context.Context, kubeClient client.Client, disSecret *distributedsecretsv1alpha1.DistributedSecrets) (*corev1.Secret, error) {
	secret := &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      disSecret.Spec.SecretRef.Name,
			Namespace: disSecret.Namespace,
		},
		Data: disSecret.Spec.SecretRef.Data,
		Type: "Opaque",
	}

	err := kubeClient.Create(ctx, secret)

	return secret, err
}
