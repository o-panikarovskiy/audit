package di

import (
	"fmt"
	"reflect"
)

var instanseDeps *ServiceLocator

// ServiceLocator struct
type ServiceLocator struct {
	services []interface{}
	types    []reflect.Type
	values   []reflect.Value
}

// Set set locator
func Set(deps *ServiceLocator) {
	if instanseDeps != nil {
		panic(fmt.Errorf("Service Locator must be singleton"))
	}
	instanseDeps = deps
}

// Get ret locator
func Get() *ServiceLocator {
	return instanseDeps
}

// Register name
func (s *ServiceLocator) Register(some interface{}) {
	s.services = append(s.services, some)
	s.types = append(s.types, reflect.TypeOf(some))
	s.values = append(s.values, reflect.ValueOf(some))
}

// Get name
func (s *ServiceLocator) Get(some interface{}) bool {
	k := reflect.TypeOf(some).Elem()
	kind := k.Kind()
	if kind == reflect.Ptr {
		k = k.Elem()
		kind = k.Kind()
	}
	for i, t := range s.types {
		if kind == reflect.Interface && t.Implements(k) {
			reflect.Indirect(
				reflect.ValueOf(some),
			).Set(s.values[i])
			return true
		} else if kind == reflect.Struct && k.AssignableTo(t.Elem()) {
			reflect.ValueOf(some).Elem().Set(s.values[i])
			return true
		}
	}
	return false
}
