package main

import (
	"testing"
	"time"
)

func TestGetWordByDate(t *testing.T) {
	type args struct {
		date time.Time
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
				date: time.Date(2022, time.February, 18, 0, 0, 0, 0, time.Now().UTC().Location()),
			},
			want:    "dodge",
			wantErr: false,
		},
		{
			name: "'shake' feb 17, 2022",
			args: args{
				date: time.Date(2022, time.February, 17, 0, 0, 0, 0, time.Now().UTC().Location()),
			},
			want:    "shake",
			wantErr: false,
		},
		{
			name: "out of date range",
			args: args{
				date: time.Date(1950, time.January, 1, 0, 0, 0, 0, time.Now().UTC().Location()),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetWordByDate(tt.args.date)
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
