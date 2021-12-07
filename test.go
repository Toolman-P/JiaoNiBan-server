package main

import (
	"JiaoNiBan-server/services/tpns"
	"testing"
)

func TestValidateIOS(t *testing.T) {

	req := tpns.NewRequest(
		tpns.WithAudience(tpns.AudienceAll),
		tpns.WithPlatform(tpns.PlatformIOS),
		tpns.WithMessageType(tpns.Notify),
		tpns.WithTitle("this is title"),
		tpns.WithContent("this is content"),
		tpns.WithEnvironment(tpns.Develop),
	)
	if err := req.Validate(); err != nil {
		t.Errorf("validate error: %v", err)
	}
}

func TestValidateAndroid(t *testing.T) {
	req := tpns.NewRequest(
		tpns.WithAudience(tpns.AudienceAll),
		tpns.WithPlatform(tpns.PlatformAndroid),
		tpns.WithMessageType(tpns.Notify),
		tpns.WithTitle("this is title"),
		tpns.WithContent("this is content"),
	)
	if err := req.Validate(); err != nil {
		t.Errorf("validate error: %v", err)
	}
}
