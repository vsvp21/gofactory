package gofactory

import (
	"gorm.io/gorm"
	"reflect"
)

func Make[T any](factory func() T, count int) []T {
	return MakeOverride[T](factory, count, *new(T))
}

func MakeOverride[T any](factory func() T, count int, overrideModel T) []T {
	m := make([]T, 0, count)
	for i := 0; i < count; i++ {
		m = append(m, override[T](factory(), overrideModel))
	}

	return m
}

func override[T any](model, model1 T) T {
	if !reflect.ValueOf(model1).IsNil() {
		ov := reflect.ValueOf(model).Elem()
		mv := reflect.ValueOf(model1).Elem()

		for i := 0; i < ov.NumField(); i++ {
			oField := ov.Field(i)
			mField := mv.Field(i)

			if mField.IsValid() && !reflect.DeepEqual(mField.Interface(), reflect.Zero(mField.Type()).Interface()) {
				if !reflect.DeepEqual(oField.Interface(), mField.Interface()) {
					oField.Set(mField)
				}
			}
		}
	}

	return model
}

func Create[T any](db *gorm.DB, factory func() T, count int) ([]T, error) {
	return CreateOverride[T](db, factory, count, *new(T))
}

func CreateOverride[T any](db *gorm.DB, factory func() T, count int, overrideModel T) ([]T, error) {
	m := MakeOverride[T](factory, count, overrideModel)

	if err := db.Create(m).Error; err != nil {
		return nil, err
	}

	return m, nil
}
