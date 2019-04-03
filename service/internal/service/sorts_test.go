package service_test

import (
	"testing"

	"github.com/sazzer/wyrdwest/service/internal/service"
	"github.com/stretchr/testify/assert"
)

type test struct {
	in  string
	out []service.SortField
}

func buildTests() []test {
	return []test{
		{"", []service.SortField{}},
		{" ", []service.SortField{}},
		{"name", []service.SortField{{Field: "name", Direction: service.SortNatural}}},
		{"name,age", []service.SortField{
			{Field: "name", Direction: service.SortNatural},
			{Field: "age", Direction: service.SortNatural},
		}},
		{"name, age", []service.SortField{
			{Field: "name", Direction: service.SortNatural},
			{Field: "age", Direction: service.SortNatural},
		}},
		{"name,     age", []service.SortField{
			{Field: "name", Direction: service.SortNatural},
			{Field: "age", Direction: service.SortNatural},
		}},
		{"name , age", []service.SortField{
			{Field: "name", Direction: service.SortNatural},
			{Field: "age", Direction: service.SortNatural},
		}},
		{" name , age ", []service.SortField{
			{Field: "name", Direction: service.SortNatural},
			{Field: "age", Direction: service.SortNatural},
		}},
		{"+name", []service.SortField{{Field: "name", Direction: service.SortAscending}}},
		{"-name", []service.SortField{{Field: "name", Direction: service.SortDescending}}},
		{"+name,-age", []service.SortField{
			{Field: "name", Direction: service.SortAscending},
			{Field: "age", Direction: service.SortDescending},
		}},
		{"+name,-age,id", []service.SortField{
			{Field: "name", Direction: service.SortAscending},
			{Field: "age", Direction: service.SortDescending},
			{Field: "id", Direction: service.SortNatural},
		}},
	}
}

func TestParseSorts(t *testing.T) {
	for _, tt := range buildTests() {
		tt := tt // Pin the variable
		t.Run(tt.in, func(t *testing.T) {
			sorts := service.ParseSorts(tt.in)
			assert.Equal(t, tt.out, sorts)
		})
	}
}

func BenchmarkParseSorts(b *testing.B) {
	for _, tt := range buildTests() {
		tt := tt // Pin the variable
		b.Run(tt.in, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				service.ParseSorts(tt.in)
			}
		})
	}
}
