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

func TestGetWordByString(t *testing.T) {
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
			got, err := GetWordByString(tt.args.dateString)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetWordByString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetWordByString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetWordByInteger(t *testing.T) {
	type args struct {
		year  int
		month int
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
				month: 2,
				day:   18,
				year:  2022,
			},
			want:    "dodge",
			wantErr: false,
		},
		{
			name: "'shake' feb 17, 2022",
			args: args{
				month: 2,
				day:   17,
				year:  2022,
			},
			want:    "shake",
			wantErr: false,
		},
		{
			name: "out of date range",
			args: args{
				month: 1,
				day:   1,
				year:  1950,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetWordByInteger(tt.args.year, tt.args.month, tt.args.day)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetWordByInteger() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetWordByInteger() = %v, want %v", got, tt.want)
			}
		})
	}
}
