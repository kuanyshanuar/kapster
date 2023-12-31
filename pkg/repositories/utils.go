package repositories

import (
	"fmt"
	"reflect"

	resident "github.com/kazakhattila/kapster"
)

func getColumns() (string, string) {
	var a resident.Resident
	var b resident.ResidentSlug
	v1 := reflect.TypeOf(a)
	v2 := reflect.TypeOf(b)
	var return1, return2 []string
	for i := 0; i < v1.NumField(); i++ {
		return1 = append(return1, v1.Field(i).Tag.Get(`json`))
	}
	for i := 0; i < v2.NumField(); i++ {
		return2 = append(return2, v2.Field(i).Tag.Get(`json`))
	}
	var result1 string = `(`
	var result2 string = `(`
	for i := 0; i < len(return1); i++ {
		if i != len(return1)-1 {
			result1 = result1 + return1[i] + `, `

		} else {
			result1 = result1 + return1[i]
		}
	}
	result1 += `)`
	for i := 0; i < len(return2); i++ {
		if i != len(return2)-1 {
			result2 = result2 + return2[i] + `, `

		} else {
			result2 = result2 + return2[i]
		}
	}
	result2 += `)`
	return result1, result2

}

func getFormatted(Data []resident.Resident) string {

	var result string
	for i := 0; i < len(Data); i++ {
		dat := Data[i]
		result = result + ` (`
		v := reflect.ValueOf(dat)
		w := v.Type()
		for j := 0; j < w.NumField(); j++ {
			vv := v.Field(j)
			if j != w.NumField()-1 {
				if vv.Kind() == reflect.Int || vv.Kind() == reflect.Float32 {
					result = result + vv.String() + `, `
				} else {
					result = result + `'` + vv.String() + `'` + `,`

				}

			} else {
				if vv.Kind() == reflect.Int || vv.Kind() == reflect.Float32 {
					result = result + vv.String()
				} else {
					result = result + `'` + vv.String() + `'`

				}

			}
		}
		if i != (len(Data) - 1) {
			result = result + `), `

		} else {
			result = result + `)`
		}
	}
	fmt.Println(result)
	return result
}
func getFormatted2(Data resident.ResidentSlug) string {

	var result string

	result = result + `(`
	v := reflect.ValueOf(Data)
	w := v.Type()
	for j := 0; j < w.NumField(); j++ {
		vv := v.Field(j)
		if j != v.NumField()-1 {

			result = result + vv.String() + `,`

		} else {
			result = result + vv.String()
		}
	}

	result = result + `)`

	return result

}
