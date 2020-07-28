package calc

import "testing"

func TestWorkHours_Calc(t *testing.T) {
	type fields struct {
		StartWork  string
		StartLunch string
		EndLunch   string
		EndWork    string
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			name: "calc hours with success",
			fields: fields{
				"09:00",
				"12:00",
				"13:00",
				"19:00",
			},
			want:    "09:00",
			wantErr: false,
		},
		{
			name: "calc hours and minutes with success",
			fields: fields{
				"09:00",
				"12:00",
				"13:00",
				"18:44",
			},
			want:    "08:44",
			wantErr: false,
		},
		{
			name: "return err when fail to parse startWork",
			fields: fields{
				"09-00",
				"12:00",
				"13:00",
				"18:44",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "return err when fail to parse startLunch",
			fields: fields{
				"09:00",
				"12-00",
				"13:00",
				"18:44",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "return err when fail to parse EndLunch",
			fields: fields{
				"09:00",
				"12:00",
				"13-00",
				"18:44",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "return err when fail to parse EndWork",
			fields: fields{
				"09:00",
				"12:00",
				"13:00",
				"18-44",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := WorkHours{
				StartWork:  tt.fields.StartWork,
				StartLunch: tt.fields.StartLunch,
				EndLunch:   tt.fields.EndLunch,
				EndWork:    tt.fields.EndWork,
			}
			got, err := h.Calc()
			if (err != nil) != tt.wantErr {
				t.Errorf("Calc() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Calc() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWorkHours_CalcPerfectEndWork(t *testing.T) {
	type fields struct {
		StartWork  string
		StartLunch string
		EndLunch   string
		EndWork    string
	}
	type args struct {
		target string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "calc hours with success",
			fields: fields{
				"09:00",
				"12:00",
				"13:00",
				"",
			},
			args: args{
				target: "09:00",
			},
			want:    "19:00",
			wantErr: false,
		},
		{
			name: "calc hours and minutes with success",
			fields: fields{
				"09:00",
				"12:00",
				"13:00",
				"",
			},
			args: args{
				target: "08:44",
			},
			want:    "18:44",
			wantErr: false,
		},
		{
			name: "return err when fail to parse startWork",
			fields: fields{
				"09-00",
				"12:00",
				"13:00",
				"",
			},
			args: args{
				target: "08:44",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "return err when fail to parse startLunch",
			fields: fields{
				"09:00",
				"12-00",
				"13:00",
				"",
			},
			args: args{
				target: "08:44",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "return err when fail to parse EndLunch",
			fields: fields{
				"09:00",
				"12:00",
				"13-00",
				"",
			},
			args: args{
				target: "08:44",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "return err when fail to parse target",
			fields: fields{
				"09:00",
				"12:00",
				"13:00",
				"",
			},
			args: args{
				target: "08-44",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := WorkHours{
				StartWork:  tt.fields.StartWork,
				StartLunch: tt.fields.StartLunch,
				EndLunch:   tt.fields.EndLunch,
				EndWork:    tt.fields.EndWork,
			}
			got, err := h.CalcPerfectEndWork(tt.args.target)
			if (err != nil) != tt.wantErr {
				t.Errorf("CalcPerfectEndWork() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CalcPerfectEndWork() got = %v, want %v", got, tt.want)
			}
		})
	}
}
