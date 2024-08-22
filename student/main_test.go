package main

import (
	"math"
	"testing"
)

// Test for calculateStatistics function
func Test_calculateStatistics(t *testing.T) {
	type args struct {
		numbers []float64
	}
	tests := []struct {
		name       string
		args       args
		wantMean   float64
		wantStddev float64
	}{
		{
			name:       "Single Value",
			args:       args{numbers: []float64{5}},
			wantMean:   5,
			wantStddev: 0,
		},
		{
			name:       "Two Identical Values",
			args:       args{numbers: []float64{5, 5}},
			wantMean:   5,
			wantStddev: 0,
		},
		{
			name:       "Two Different Values",
			args:       args{numbers: []float64{1, 9}},
			wantMean:   5,
			wantStddev: 4,
		},
		{
			name:       "Five Values (Positive Integers)",
			args:       args{numbers: []float64{1, 2, 3, 4, 5}},
			wantMean:   3,
			wantStddev: math.Sqrt(2), // √2 ≈ 1.414
		},
		{
			name:       "Large Numbers",
			args:       args{numbers: []float64{1000000, 1000001, 1000002}},
			wantMean:   1000001,
			wantStddev: 0.816496580927726, // ≈ sqrt(((1000000-1000001)^2 + (1000001-1000001)^2 + (1000002-1000001)^2) / 3) ≈ 0.816
		},
		{
			name:       "High Precision",
			args:       args{numbers: []float64{1.000001, 1.000002, 1.000003}},
			wantMean:   1.000002,
			wantStddev: 0.000000816496580927726, // High precision values
		},
		{
			name:       "Single Large Value",
			args:       args{numbers: []float64{1e6}},
			wantMean:   1e6,
			wantStddev: 0,
		},
		{
			name:       "Negative Values",
			args:       args{numbers: []float64{-1, -2, -3, -4, -5}},
			wantMean:   -3,
			wantStddev: math.Sqrt(2), // Same as positive values with negation
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMean, gotStddev := calculateStatistics(tt.args.numbers)
			if gotMean != tt.wantMean {
				t.Errorf("calculateStatistics() gotMean = %v, want %v", gotMean, tt.wantMean)
			}
			if math.Abs(gotStddev-tt.wantStddev) > 1e-6 { // Allowing a small error margin for floating point comparisons
				t.Errorf("calculateStatistics() gotStddev = %v, want %v", gotStddev, tt.wantStddev)
			}
		})
	}
}

// Test for guess_it_1 function
func Test_guess_it_1(t *testing.T) {
	type args struct {
		mean   float64
		stddev float64
	}
	tests := []struct {
		name           string
		args           args
		wantLowerLimit float64
		wantUpperLimit float64
	}{
		{
			name:           "Zero Mean and Zero Stddev",
			args:           args{mean: 0, stddev: 0},
			wantLowerLimit: 0,
			wantUpperLimit: 0,
		},
		{
			name:           "Zero Stddev",
			args:           args{mean: 10, stddev: 0},
			wantLowerLimit: 10,
			wantUpperLimit: 10,
		},
		{
			name:           "Positive Mean and Positive Stddev",
			args:           args{mean: 10, stddev: 2},
			wantLowerLimit: 4,
			wantUpperLimit: 16,
		},
		{
			name:           "Negative Mean and Positive Stddev",
			args:           args{mean: -5, stddev: 3},
			wantLowerLimit: -14,
			wantUpperLimit: 4,
		},
		{
			name:           "Positive Mean and Negative Stddev",
			args:           args{mean: 7, stddev: -1},
			wantLowerLimit: 10, // In practice, stddev should not be negative; testing edge cases
			wantUpperLimit: 4,
		},
		{
			name:           "Large Mean and Large Stddev",
			args:           args{mean: 1000, stddev: 500},
			wantLowerLimit: -500,
			wantUpperLimit: 2500,
		},
		{
			name:           "Negative Mean and Stddev",
			args:           args{mean: -100, stddev: 10},
			wantLowerLimit: -130,
			wantUpperLimit: -70,
		},
		{
			name:           "Mean and Stddev Both Zero",
			args:           args{mean: 0, stddev: 0},
			wantLowerLimit: 0,
			wantUpperLimit: 0,
		},
		{
			name:           "Very Large Stddev",
			args:           args{mean: 50, stddev: 1000},
			wantLowerLimit: -2950,
			wantUpperLimit: 3050,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotLowerLimit, gotUpperLimit := guess_it_1(tt.args.mean, tt.args.stddev)
			if gotLowerLimit != tt.wantLowerLimit {
				t.Errorf("guess_it_1() gotLowerLimit = %v, want %v", gotLowerLimit, tt.wantLowerLimit)
			}
			if gotUpperLimit != tt.wantUpperLimit {
				t.Errorf("guess_it_1() gotUpperLimit = %v, want %v", gotUpperLimit, tt.wantUpperLimit)
			}
		})
	}
}
