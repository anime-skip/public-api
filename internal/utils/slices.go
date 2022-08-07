package utils

import (
	"github.com/samber/lo"
)

func CopySlice[T any](original []T) []T {
	if original == nil {
		return nil
	}
	copied := make([]T, len(original))
	copy(copied, original)
	return copied
}

// Based on https://github.com/anime-skip/player/blob/d98c41a659cfee12757b209bd68292352bc8a6bd/packages/common/src/utils/GeneralUtils.ts#L171
func ComputeSliceDiffs[T any](
	newItems []T,
	oldItems []T,
	getID func(t1 T) string,
	needsUpdated func(t1 T, t2 T) bool,
) (toLeave []T, toCreate []T, toUpdate []T, toDelete []T) {
	toCreate = []T{}
	toDelete = []T{}

	getItemMap := func(items []T) map[string]T {
		return lo.Reduce(items, func(res map[string]T, item T, _ int) map[string]T {
			res[getID(item)] = item
			return res
		}, map[string]T{})
	}
	oldItemsMap := getItemMap(oldItems)
	newItemsMap := getItemMap(newItems)

	intersection := lo.Filter[T](newItems, func(newItem T, _ int) bool {
		_, ok := oldItemsMap[getID(newItem)]
		return ok
	})
	toLeave = []T{}
	toUpdate = []T{}
	for _, newItem := range intersection {
		oldItem := oldItemsMap[getID(newItem)]
		if needsUpdated(newItem, oldItem) {
			toUpdate = append(toUpdate, newItem)
		} else {
			toLeave = append(toLeave, oldItem)
		}
	}

	toCreate = lo.Filter(newItems, func(newItem T, _ int) bool {
		_, ok := oldItemsMap[getID(newItem)]
		return !ok
	})
	toDelete = lo.Filter(oldItems, func(oldItem T, _ int) bool {
		_, ok := newItemsMap[getID(oldItem)]
		return !ok
	})
	return
}
