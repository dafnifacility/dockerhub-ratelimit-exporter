package dhubratelimit

import (
	"testing"
	"time"
)

func TestInnerResult_Get(t *testing.T) {
	type fields struct {
		PullLimit     int
		PullRemaining int
		CheckTime     time.Time
		Window        time.Duration
		Authenticated bool
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			fields: fields{
				PullLimit:     66,
				PullRemaining: 99,
				Window:        33 * time.Second,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &InnerResult{
				PullLimit:     tt.fields.PullLimit,
				PullRemaining: tt.fields.PullRemaining,
				CheckTime:     tt.fields.CheckTime,
				Window:        tt.fields.Window,
				Authenticated: tt.fields.Authenticated,
			}
			if got := r.GetLimit(); got != tt.fields.PullLimit {
				t.Errorf("InnerResult.GetLimit() = %v, want %v", got, tt.fields)
			}
			if got := r.GetRemaining(); got != tt.fields.PullRemaining {
				t.Errorf("InnerResult.GetRemaining() = %v, want %v", got, tt.fields)
			}
			if got := r.GetWindow(); got != int(tt.fields.Window) {
				t.Errorf("InnerResult.GetWindow() = %v, want %v", got, tt.fields)
			}
		})
	}
}

func Test_splitRatelimitHeader(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name     string
		inheader string
		want     int
		want1    int
	}{
		{
			name:     "simple",
			inheader: "100;w=21600",
			want:     100,
			want1:    21600,
		},
		{
			name:     "simple2",
			inheader: "98;w=333333",
			want:     98,
			want1:    333333,
		},
		{
			name:     "zero",
			inheader: "0;w=0",
			want:     0,
			want1:    0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := splitRatelimitHeader(tt.inheader)
			if got != tt.want {
				t.Errorf("splitRatelimitHeader() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("splitRatelimitHeader() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
