package score

import "testing"

func Test_calcNameScore(t *testing.T) {
	type args struct {
		lastName  string
		firstName string
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{"张杰", args{"张", "杰"}, 92.5, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CalcNameScore(tt.args.lastName, tt.args.firstName)
			if (err != nil) != tt.wantErr {
				t.Errorf("CalcNameScore() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CalcNameScore() got = %v, want %v", got, tt.want)
			}
		})
	}
}
