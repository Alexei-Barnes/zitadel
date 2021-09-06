package restore

import (
	"github.com/caos/orbos/pkg/labels"
	"github.com/caos/zitadel/operator/helpers"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func getJob(
	namespace string,
	nameLabels *labels.Name,
	nodeselector map[string]string,
	tolerations []corev1.Toleration,
	backupSecretName string,
	saSecretKey string,
	assetAKIDKey string,
	assetSAKKey string,
	command string,
	image string,
) *batchv1.Job {
	return &batchv1.Job{
		ObjectMeta: v1.ObjectMeta{
			Name:      nameLabels.Name(),
			Namespace: namespace,
			Labels:    labels.MustK8sMap(nameLabels),
		},
		Spec: batchv1.JobSpec{
			Template: corev1.PodTemplateSpec{
				Spec: corev1.PodSpec{
					NodeSelector:  nodeselector,
					Tolerations:   tolerations,
					RestartPolicy: corev1.RestartPolicyNever,
					Containers: []corev1.Container{{
						Name:  nameLabels.Name(),
						Image: image,
						Command: []string{
							"/bin/bash",
							"-c",
							command,
						},
						VolumeMounts: []corev1.VolumeMount{
							{
								Name:      certsInternalSecretName,
								MountPath: certPath,
							}, {
								Name:      saInternalSecretName,
								SubPath:   saSecretKey,
								MountPath: saSecretPath,
							}, {
								Name:      akidInternalSecretName,
								SubPath:   assetAKIDKey,
								MountPath: akidSecretPath,
							}, {
								Name:      sakInternalSecretName,
								SubPath:   assetSAKKey,
								MountPath: sakSecretPath,
							},
						},
						ImagePullPolicy: corev1.PullAlways,
					}},
					Volumes: []corev1.Volume{
						{
							Name: certsInternalSecretName,
							VolumeSource: corev1.VolumeSource{
								Secret: &corev1.SecretVolumeSource{
									SecretName:  rootSecretName,
									DefaultMode: helpers.PointerInt32(defaultMode),
								},
							},
						}, {
							Name: saInternalSecretName,
							VolumeSource: corev1.VolumeSource{
								Secret: &corev1.SecretVolumeSource{
									SecretName:  backupSecretName,
									DefaultMode: helpers.PointerInt32(defaultMode),
								},
							},
						}, {
							Name: akidInternalSecretName,
							VolumeSource: corev1.VolumeSource{
								Secret: &corev1.SecretVolumeSource{
									SecretName:  backupSecretName,
									DefaultMode: helpers.PointerInt32(defaultMode),
								},
							},
						}, {
							Name: sakInternalSecretName,
							VolumeSource: corev1.VolumeSource{
								Secret: &corev1.SecretVolumeSource{
									SecretName:  backupSecretName,
									DefaultMode: helpers.PointerInt32(defaultMode),
								},
							},
						},
					},
				},
			},
		},
	}
}