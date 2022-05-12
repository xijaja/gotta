package gottaos

import "testing"

func TestGuluOS_PrintOS(t *testing.T) {
	var tests []struct {
		name string
		gos  GottaOS
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.gos.WhatIsOS()
		})
	}
}
