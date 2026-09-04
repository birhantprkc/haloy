package healthcheck

import (
	"testing"
	"time"
)

func TestDefaultConfig(t *testing.T) {
	config := DefaultConfig()

	if !config.Enabled {
		t.Error("DefaultConfig().Enabled should be true (enabled by default)")
	}
	if config.Interval != 15*time.Second {
		t.Errorf("DefaultConfig().Interval = %v, want %v", config.Interval, 15*time.Second)
	}
	if config.Fall != 3 {
		t.Errorf("DefaultConfig().Fall = %d, want 3", config.Fall)
	}
	if config.Rise != 2 {
		t.Errorf("DefaultConfig().Rise = %d, want 2", config.Rise)
	}
	if config.Timeout != 5*time.Second {
		t.Errorf("DefaultConfig().Timeout = %v, want %v", config.Timeout, 5*time.Second)
	}
}

func TestTargetState_String(t *testing.T) {
	tests := []struct {
		state TargetState
		want  string
	}{
		{StateHealthy, "healthy"},
		{StateUnhealthy, "unhealthy"},
		{TargetState(99), "unknown"},
	}

	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := tt.state.String(); got != tt.want {
				t.Errorf("TargetState.String() = %q, want %q", got, tt.want)
			}
		})
	}
}
