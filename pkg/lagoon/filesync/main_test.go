package filesync

import "testing"

func TestSyncFiles(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"I'm here", "I'm here"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SyncFiles(); got != tt.want {
				t.Errorf("SyncFiles() = %v, want %v", got, tt.want)
			}
		})
	}
}
