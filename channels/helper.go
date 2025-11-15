package channels

import (
	"fmt"
	"github.com/wuzhixiang0827/goravel-notification/contracts"
	"reflect"
)

// ---- 辅助：反射调用通知类型的 ToXxx 方法 ----

// callToMethod 会尝试在 notification 值上调用 methodName(notifiable)
// 要求被调用方法返回两个值 (interface{}, error) 或者 (map[string]interface{}, error) 或 (map[string]string, error)
func callToMethod(notification interface{}, methodName string, notifiable contracts.Notifiable) (map[string]interface{}, error) {
	v := reflect.ValueOf(notification)
	// attempt pointer receiver methods: if v is not pointer, try pointer
	if !v.IsValid() {
		return nil, fmt.Errorf("invalid notification value")
	}

	// 如果方法存在于指针接收者但 not on value, use Addr() when possible
	method := v.MethodByName(methodName)
	if !method.IsValid() {
		// try pointer
		if v.Kind() != reflect.Ptr && v.CanAddr() {
			method = v.Addr().MethodByName(methodName)
		}
	}
	if !method.IsValid() {
		return nil, fmt.Errorf("method %s not found", methodName)
	}

	// prepare args
	args := []reflect.Value{reflect.ValueOf(notifiable)}
	results := method.Call(args)

	if len(results) == 0 {
		return nil, fmt.Errorf("method %s returned no values", methodName)
	}

	// last result expected to be error
	var err error
	if len(results) >= 2 {
		if !results[1].IsNil() {
			if e, ok := results[1].Interface().(error); ok {
				err = e
			} else {
				// convert non-error second return to error string
				err = fmt.Errorf("second return of %s is not error", methodName)
			}
		}
	}

	if err != nil {
		return nil, err
	}

	// first result parsed into map[string]interface{}
	first := results[0].Interface()
	// if already map[string]interface{}, cast
	if m, ok := first.(map[string]interface{}); ok {
		return m, nil
	}
	// if map[string]string -> convert
	if ms, ok := first.(map[string]string); ok {
		out := make(map[string]interface{}, len(ms))
		for k, v := range ms {
			out[k] = v
		}
		return out, nil
	}
	// if struct or other, try to convert via reflection -> map by exported fields
	rv := reflect.ValueOf(first)
	if rv.Kind() == reflect.Struct {
		out := make(map[string]interface{})
		rt := rv.Type()
		for i := 0; i < rv.NumField(); i++ {
			field := rt.Field(i)
			if field.PkgPath != "" { // unexported
				continue
			}
			out[field.Name] = rv.Field(i).Interface()
		}
		return out, nil
	}

	return nil, fmt.Errorf("unsupported return type from %s", methodName)
}

// Exported helper for channels to call ToX methods
func CallToMethod(notification interface{}, methodName string, notifiable contracts.Notifiable) (map[string]interface{}, error) {
	return callToMethod(notification, methodName, notifiable)
}
