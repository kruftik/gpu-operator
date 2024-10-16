package utils

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

func LabelSelectorPredicate[T client.Object](s metav1.LabelSelector) (predicate.TypedPredicate[T], error) {
	selector, err := metav1.LabelSelectorAsSelector(&s)
	if err != nil {
		return predicate.TypedFuncs[T]{}, err
	}
	return predicate.NewTypedPredicateFuncs[T](func(o T) bool {
		return selector.Matches(labels.Set(o.GetLabels()))
	}), nil
}
