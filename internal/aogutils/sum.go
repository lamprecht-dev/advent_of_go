package aogutils

func Sum[T uint8 | int8 | uint16 | int16 | uint32 | int32 | uint64 | int64 | float32 | float64 | int](list []T) (sum T) {
	for _, l := range list {
		sum += l
	}

	return
}
