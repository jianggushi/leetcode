package solution

import (
	"reflect"
	"testing"
)

func Test_findRestaurant(t *testing.T) {
	type args struct {
		list1 []string
		list2 []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
		{
			"test01",
			args{
				[]string{"Shogun", "Tapioca Express", "Burger King", "KFC"},
				[]string{"Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"},
			},
			[]string{"Shogun"},
		},
		{
			"test02",
			args{
				[]string{"Shogun", "Tapioca Express", "Burger King", "KFC"},
				[]string{"KFC", "Shogun", "Burger King"},
			},
			[]string{"Shogun"},
		},
		{
			"test03",
			args{
				[]string{"Shogun", "Tapioca Express", "Burger King", "KFC"},
				[]string{"KFC", "Burger King", "Tapioca Express", "Shogun"},
			},
			[]string{"KFC", "Burger King", "Tapioca Express", "Shogun"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findRestaurant(tt.args.list1, tt.args.list2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findRestaurant() = %v, want %v", got, tt.want)
			}
		})
	}
}
