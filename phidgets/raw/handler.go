package raw

// #include <stddef.h>
// #cgo CFLAGS: -F /Library/Frameworks -framework Phidget22 -I /Library/Frameworks/Phidget22.framework/Headers

// #include "handler.h"
import "C"

func createHandler(f func(h *C.handler) C.int) (*C.handler, error) {
	h := C.newHandler()

	if err := result(f(h)); err != nil {
		C.handlerFree(h)
		return nil, err
	}

	return h, nil
}
