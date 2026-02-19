package ffi

/*
#cgo LDFLAGS: -L../../rust-lib/target/release -lvecforge
#include "vecforge.h"
*/
import "C"

import (
	"unsafe"
)

func Search(query []float32, data [][]float32, k int) []Hit {
	if len(query) == 0 || len(data) == 0 {
		return nil
	}
	
	result := C.vec_search_multi(
		(*C.float)(&query[0]),
		C.size_t(len(query)),
		(*C.float)(&data[0][0]),
		C.size_t(len(data)),
		C.size_t(len(query)),
		C.size_t(k),
	)
	
	if result == nil {
		return nil
	}
	defer C.vec_search_multi_free(result)
	
	hits := *result
	resultHits := make([]Hit, len(hits))
	for i, h := range hits {
		resultHits[i] = Hit{
			ID:       C.GoString(h.id),
			Score:    float64(h.score),
			Provider: C.GoString(h.provider),
		}
	}
	return resultHits
}

type Hit struct {
	ID       string
	Score   float64
	Provider string
}
