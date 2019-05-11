//利用 reflect 技术把结构体的可 export 值复制到 v 中，v 必须是相似结构体的指针。
//copy the exported value of a struct to v , with reflect.
package gostructcopy

import (
	"errors"
	"reflect"
)

//StructCopy copy the exported value of a struct to a likely struct , with reflect.
func StructCopy(src, dst interface{}) error {
	srcV, err := srcFilter(src)
	if err != nil {
		return err
	}
	dstV, err := dstFilter(dst)
	if err != nil {
		return err
	}
	srcKeys := make(map[string]bool)
	for i := 0; i < srcV.NumField(); i++ {
		srcKeys[srcV.Type().Field(i).Name] = true
	}
	for i := 0; i < dstV.Elem().NumField(); i++ {
		fName := dstV.Elem().Type().Field(i).Name
		if _, ok := srcKeys[fName]; ok {
			v := srcV.FieldByName(dstV.Elem().Type().Field(i).Name)
			if v.CanInterface() {
				dstV.Elem().Field(i).Set(v)
			}
		}
	}

	return nil
}

func srcFilter(src interface{}) (reflect.Value, error) {
	v := reflect.ValueOf(src)
	if v.Type().Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return reflect.Zero(v.Type()), errors.New("src type error: not a struct or a pointer to struct")
	}
	return v, nil
}

func dstFilter(src interface{}) (reflect.Value, error) {
	v := reflect.ValueOf(src)
	if v.Type().Kind() != reflect.Ptr {
		return reflect.Zero(v.Type()), errors.New("src type error: not a pointer to struct")
	}
	if v.Elem().Kind() != reflect.Struct {
		return reflect.Zero(v.Type()), errors.New("src type error: not point to struct")
	}
	return v, nil
}
