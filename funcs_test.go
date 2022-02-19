package main

import (
	"testing"
)

func TestGetWordByDate(t *testing.T) {
	type args struct {
		dateString string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "'dodge' feb 18, 2022",
			args: args{
				dateString: "2022-02-18",
			},
			want:    "dodge",
			wantErr: false,
		},
		{
			name: "'shake' feb 17, 2022",
			args: args{
				dateString: "2022-02-17",
			},
			want:    "shake",
			wantErr: false,
		},
		{
			name: "out of date range",
			args: args{
				dateString: "1950-01-01",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetWordByDate(tt.args.dateString)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetWordByDate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetWordByDate() = %v, want %v", got, tt.want)
			}
		})
	}
}
