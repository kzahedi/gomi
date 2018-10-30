package gomi

import (
	"math"
	"testing"
)

func createParamData(how string) (Parameters, Data) {
	var W [][]float64
	var S [][]float64
	var A [][]float64
	var Wd [][]int
	var Sd [][]int
	var Ad [][]int
	dd := DataDiscretised{W: Wd, S: Sd, A: Ad}
	p := CreateParametersContainer()
	d := Data{W: W, S: S, A: A, Discretised: dd}
	switch how {
	case "uniform":
		d.W = [][]float64{
			{0.0, 1.0},
			{1.0, 0.0},
			{0.0, 0.0},
			{1.0, 1.0},
		}
		d.A = [][]float64{
			{0.0},
			{1.0},
			{0.0},
			{1.0},
		}
		d.S = [][]float64{
			{1.0},
			{0.0},
			{1.0},
			{0.0},
		}
	}

	return p, d
}

func TestContinuousAvgCalculations(t *testing.T) {
	type args struct {
		p Parameters
		d Data
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContinuousAvgCalculations(tt.args.p, tt.args.d); got != tt.want {
				t.Errorf("ContinuousAvgCalculations() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMiWContinuousAvg(t *testing.T) {
	param, data := createParamData("uniform")

	type args struct {
		p    Parameters
		data Data
	}
	tests := []struct {
		name       string
		args       args
		wantResult float64
	}{
		{name: "Random data MI_W", args: args{p: param, data: data}, wantResult: 3.1283},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := MiWContinuousAvg(tt.args.p, tt.args.data); math.Abs(gotResult-tt.wantResult) > 0.0001 {
				t.Errorf("MiWContinuousAvg() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestMiAContinuousAvg(t *testing.T) {
	type args struct {
		p    Parameters
		data Data
	}
	tests := []struct {
		name       string
		args       args
		wantResult float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := MiAContinuousAvg(tt.args.p, tt.args.data); gotResult != tt.wantResult {
				t.Errorf("MiAContinuousAvg() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestMiAPrimeContinuousAvg(t *testing.T) {
	type args struct {
		p    Parameters
		data Data
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MiAPrimeContinuousAvg(tt.args.p, tt.args.data); got != tt.want {
				t.Errorf("MiAPrimeContinuousAvg() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMiMiContinuousAvg(t *testing.T) {
	type args struct {
		p    Parameters
		data Data
	}
	tests := []struct {
		name       string
		args       args
		wantResult float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := MiMiContinuousAvg(tt.args.p, tt.args.data); gotResult != tt.wantResult {
				t.Errorf("MiMiContinuousAvg() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestMiCaContinuousAvg(t *testing.T) {
	type args struct {
		p    Parameters
		data Data
	}
	tests := []struct {
		name       string
		args       args
		wantResult float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := MiCaContinuousAvg(tt.args.p, tt.args.data); gotResult != tt.wantResult {
				t.Errorf("MiCaContinuousAvg() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestMiWaContinuousAvg(t *testing.T) {
	type args struct {
		p    Parameters
		data Data
	}
	tests := []struct {
		name       string
		args       args
		wantResult float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := MiWaContinuousAvg(tt.args.p, tt.args.data); gotResult != tt.wantResult {
				t.Errorf("MiWaContinuousAvg() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestMiWsContinuousAvg(t *testing.T) {
	type args struct {
		p    Parameters
		data Data
	}
	tests := []struct {
		name       string
		args       args
		wantResult float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := MiWsContinuousAvg(tt.args.p, tt.args.data); gotResult != tt.wantResult {
				t.Errorf("MiWsContinuousAvg() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestMiInContinuousAvg(t *testing.T) {
	type args struct {
		p    Parameters
		data Data
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MiInContinuousAvg(tt.args.p, tt.args.data); got != tt.want {
				t.Errorf("MiInContinuousAvg() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMiWpContinuousAvg(t *testing.T) {
	type args struct {
		p    Parameters
		data Data
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MiWpContinuousAvg(tt.args.p, tt.args.data); got != tt.want {
				t.Errorf("MiWpContinuousAvg() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCaContinuousAvg(t *testing.T) {
	type args struct {
		p    Parameters
		data Data
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CaContinuousAvg(tt.args.p, tt.args.data); got != tt.want {
				t.Errorf("CaContinuousAvg() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUIContinuousAvg(t *testing.T) {
	type args struct {
		p    Parameters
		data Data
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UIContinuousAvg(tt.args.p, tt.args.data); got != tt.want {
				t.Errorf("UIContinuousAvg() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCiContinuousAvg(t *testing.T) {
	type args struct {
		p    Parameters
		data Data
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CiContinuousAvg(tt.args.p, tt.args.data); got != tt.want {
				t.Errorf("CiContinuousAvg() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMiSyContinuousAvg(t *testing.T) {
	type args struct {
		p    Parameters
		data Data
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MiSyContinuousAvg(tt.args.p, tt.args.data); got != tt.want {
				t.Errorf("MiSyContinuousAvg() = %v, want %v", got, tt.want)
			}
		})
	}
}
