package gomi

import (
	"testing"
)

func createPWBins(wBins []int) Parameters {
	r := CreateParametersContainer()
	r.WBins = wBins
	return r
}

func createPABins(aBins []int) Parameters {
	r := CreateParametersContainer()
	r.ABins = aBins
	return r
}

func createPSBins(aSins []int) Parameters {
	r := CreateParametersContainer()
	r.SBins = aSins
	return r
}

func createPGlobalBins(v int) Parameters {
	r := CreateParametersContainer()
	r.GlobalBins = v
	return r
}

func createData() Data {
	d := Data{W: nil, A: nil, S: nil, Discretised: DataDiscretised{W: nil, S: nil, A: nil}}
	d.W = make([][]float64, 10, 10)
	d.W[0] = make([]float64, 2, 2)
	return d
}

func Test_CalculateWBins(t *testing.T) {
	type args struct {
		p Parameters
		d Data
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "WBins with only one value.",
			args: args{p: createPWBins([]int{10}),
				d: createData()},
			want: 10},
		{name: "WBins with only two values.",
			args: args{p: createPWBins([]int{10, 5}),
				d: createData()},
			want: 5 * 10},
		{name: "WBins with only three values.",
			args: args{p: createPWBins([]int{10, 5, 17}),
				d: createData()},
			want: 10 * 5 * 17},
		{name: "WBins with no values given.",
			args: args{p: createPGlobalBins(123),
				d: createData()},
			want: 15129},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateWBins(tt.args.p, tt.args.d); got != tt.want {
				t.Errorf("calculateWBins() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calculateABins(t *testing.T) {
	type args struct {
		p Parameters
		d Data
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "ABins with only one value.",
			args: args{p: createPABins([]int{10}),
				d: createData()},
			want: 10},
		{name: "ABins with only two values.",
			args: args{p: createPABins([]int{10, 5}),
				d: createData()},
			want: 5 * 10},
		{name: "ABins with only three values.",
			args: args{p: createPABins([]int{10, 5, 17}),
				d: createData()},
			want: 10 * 5 * 17},
		{name: "ABins with no values given.",
			args: args{p: createPGlobalBins(123),
				d: createData()},
			want: 15129},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateABins(tt.args.p, tt.args.d); got != tt.want {
				t.Errorf("calculateABins() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calculateSBins(t *testing.T) {
	type args struct {
		p Parameters
		d Data
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "SBins with only one value.",
			args: args{p: createPSBins([]int{10}),
				d: createData()},
			want: 10},
		{name: "SBins with only two values.",
			args: args{p: createPSBins([]int{10, 5}),
				d: createData()},
			want: 5 * 10},
		{name: "SBins with only three values.",
			args: args{p: createPSBins([]int{10, 5, 17}),
				d: createData()},
			want: 10 * 5 * 17},
		{name: "SBins with no values given.",
			args: args{p: createPGlobalBins(123),
				d: createData()},
			want: 15129},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateSBins(tt.args.p, tt.args.d); got != tt.want {
				t.Errorf("calculateABins() = %v, want %v", got, tt.want)
			}
		})
	}
}
