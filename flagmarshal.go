package flagmarshal

import (
	"errors"
	"flag"
	"reflect"
	"strconv"
	"strings"
)

func ParseFlags(target interface{}) ([]string, error) {
	targetValue := reflect.ValueOf(target).Elem()
	if targetValue.Kind() == reflect.Ptr {
		targetValue = reflect.Indirect(targetValue)
	}
	if targetValue.Kind() != reflect.Struct {
		return nil, errors.New("struct pointer expected, " + targetValue.Kind().String() + " receved")
	}
	if !targetValue.CanAddr() {
		return nil, errors.New("struct pointer expected, " + targetValue.Kind().String() + " receved")
	}
	for i, I := 0, targetValue.NumField(); i < I; i++ {
		fieldType := targetValue.Type().Field(i)
		if found, ok := fieldType.Tag.Lookup("flag"); ok {
			if !fieldType.IsExported() {
				return nil, errors.New("can't use flags for unexported fieldType " + fieldType.Name)
			}
			fieldValue := targetValue.Field(i)
			for _, flagName := range strings.Split(found, ",") {
				if flagName == "" {
					continue
				}
				help := fieldType.Tag.Get("help")
				switch value := fieldValue.Interface().(type) {
				case string:
					flag.Func(flagName, help, func(strval string) error {
						fieldValue.Set(reflect.ValueOf(strval))
						return nil
					})
				case []string:
					flag.Func(flagName, help, func(strval string) error {
						if fieldValue.Interface() == nil {
							fieldValue.Set(reflect.ValueOf([]string{strval}))
						} else {
							fieldValue.Set(reflect.ValueOf(append(value, strval)))
						}
						return nil
					})
				case uint64:
					flag.Func(flagName, help, func(strval string) error {
						value, err := strconv.ParseUint(strval, 0, 64)
						if err == nil {
							fieldValue.Set(reflect.ValueOf(value))
						}
						return err
					})
				case uint:
					flag.Func(flagName, help, func(strval string) error {
						value, err := strconv.ParseUint(strval, 0, 32)
						if err == nil {
							fieldValue.Set(reflect.ValueOf(uint(value)))
						}
						return err
					})
				case int64:
					flag.Func(flagName, help, func(strval string) error {
						value, err := strconv.ParseInt(strval, 0, 64)
						if err == nil {
							fieldValue.Set(reflect.ValueOf(value))
						}
						return err
					})
				case int:
					flag.Func(flagName, help, func(strval string) error {
						value, err := strconv.ParseInt(strval, 0, 32)
						if err == nil {
							fieldValue.Set(reflect.ValueOf(int(value)))
						}
						return err
					})
				case bool:
					flag.Func(flagName, help, func(strval string) error {
						value, err := strconv.ParseBool(strval)
						if err == nil {
							fieldValue.Set(reflect.ValueOf(value))
						}
						return err
					})
				case float64:
					flag.Func(flagName, help, func(strval string) error {
						value, err := strconv.ParseFloat(strval, 64)
						if err == nil {
							fieldValue.Set(reflect.ValueOf(value))
						}
						return err
					})
				case float32:
					flag.Func(flagName, help, func(strval string) error {
						value, err := strconv.ParseFloat(strval, 32)
						if err == nil {
							fieldValue.Set(reflect.ValueOf(float32(value)))
						}
						return err
					})
				default:
					return nil, errors.New("unsupported type " + fieldType.Type.Kind().String() + " for " + fieldType.Name)
				}
			}
		}
	}
	flag.Parse()
	return flag.Args(), nil
}
