package stub

import (
	"context"

	"github.com/operator-framework/operator-sdk/pkg/sdk"
	"github.com/sirupsen/logrus"
)

func NewHandler() sdk.Handler {
	return &Handler{}
}

type Handler struct {
	// Fill me
}

func (h *Handler) Handle(ctx context.Context, event sdk.Event) error {
	// switch o := event.Object.(type) {
	// case *v1alpha1.Opsler:
	// 	err := sdk.Create(newbusyBoxPod(o))
	// 	if err != nil && !errors.IsAlreadyExists(err) {
	// 		logrus.Errorf("Failed to create busybox pod : %v", err)
	// 		return err
	// 	}
	// }
	logrus.Info("handle")
	return nil
}

// // newbusyBoxPod demonstrates how to create a busybox pod
// func newbusyBoxPod(cr *v1alpha1.Opsler) *corev1.Pod {
// 	labels := map[string]string{
// 		"app": "busy-box",
// 	}
// 	return &corev1.Pod{
// 		TypeMeta: metav1.TypeMeta{
// 			Kind:       "Pod",
// 			APIVersion: "v1",
// 		},
// 		ObjectMeta: metav1.ObjectMeta{
// 			Name:      "busy-box",
// 			Namespace: cr.Namespace,
// 			OwnerReferences: []metav1.OwnerReference{
// 				*metav1.NewControllerRef(cr, schema.GroupVersionKind{
// 					Group:   v1alpha1.SchemeGroupVersion.Group,
// 					Version: v1alpha1.SchemeGroupVersion.Version,
// 					Kind:    "Opsler",
// 				}),
// 			},
// 			Labels: labels,
// 		},
// 		Spec: corev1.PodSpec{
// 			Containers: []corev1.Container{
// 				{
// 					Name:    "busybox",
// 					Image:   "busybox",
// 					Command: []string{"sleep", "3600"},
// 				},
// 			},
// 		},
// 	}
// }
