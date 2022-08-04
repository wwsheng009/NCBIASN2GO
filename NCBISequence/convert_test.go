package NCBISequence

import "testing"

func Test_convert_json_data(t *testing.T) {
	type args struct {
		fname string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "sequence_set",
			args: args{fname: "sequence"},
		}, {
			name: "sequence",
			args: args{fname: "1701950027.asn"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			convert_json_data(tt.args.fname)
		})
	}
}
