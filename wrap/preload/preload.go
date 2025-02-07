package preload

import (
	"gin_work/preload"
	"reflect"
)

func Load() {
	var globalPreload *preload.GlobalPreload
	globalValue := reflect.ValueOf(globalPreload)
	for i := 0; i < globalValue.NumMethod(); i++ {
		globalValue.Method(i).Call(nil)
	}
}
