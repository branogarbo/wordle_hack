package main

import "testing"

func TestGetWordByDate(t *testing.T) {
	type args struct {
		month int
		day   int
		year  int
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
			got, err := GetWordByDate(tt.args.month, tt.args.day, tt.args.year)
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
