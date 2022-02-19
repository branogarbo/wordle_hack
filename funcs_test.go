package main

import (
	"testing"
	"time"
)

func TestGetWordByDate(t *testing.T) {
	type args struct {
		year  int
		month time.Month
		day   int
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
				year:  2022,
				month: time.February,
				day:   18,
			},
			want:    "dodge",
			wantErr: false,
		},
		{
			name: "'shake' feb 17, 2022",
			args: args{
				year:  2022,
				month: time.February,
				day:   17,
			},
			want:    "shake",
			wantErr: false,
		},
		{
			name: "out of date range",
			args: args{
				year:  1950,
				month: time.January,
				day:   1,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetWordByDate(tt.args.year, tt.args.month, tt.args.day)
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
