package dxgi

import (
	"testing"
)


func TestDXGIOutput_GetDesc(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping Windows-specific test in short mode")
	}

	output := getTestOutput(t)
	defer output.Release()

	desc, err := output.GetDesc()
	if err != nil {
		t.Fatalf("GetDesc failed: %v", err)
	}

	if desc == nil {
		t.Fatal("GetDesc returned nil")
	}

	t.Logf("Output DeviceName: %v", desc.DeviceName)
	t.Logf("AttachedToDesktop: %v", desc.AttachedToDesktop)
	t.Logf("Rotation: %d", desc.Rotation)
}

func TestDXGIOutput_GetDesc1(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping Windows-specific test in short mode")
	}

	output := getTestOutput(t)
	defer output.Release()

	desc, err := output.GetDesc1()
	if err != nil {
		t.Logf("GetDesc1 failed (may not be supported): %v", err)
		return
	}

	if desc == nil {
		t.Fatal("GetDesc1 returned nil")
	}

	t.Logf("Output BitsPerColor: %d", desc.BitsPerColor)
	t.Logf("Output ColorSpace: %d", desc.ColorSpace)
}

func TestDXGIOutput_WaitForVBlank(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping Windows-specific test in short mode")
	}

	output := getTestOutput(t)
	defer output.Release()

	err := output.WaitForVBlank()
	if err != nil {
		t.Logf("WaitForVBlank failed: %v", err)
	} else {
		t.Log("WaitForVBlank succeeded")
	}
}

func TestDXGIOutput_GetDisplayModeList(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping Windows-specific test in short mode")
	}

	output := getTestOutput(t)
	defer output.Release()

	// 首先获取模式数量
	var numModes uint32
	err := output.GetDisplayModeList(0, 0, &numModes, nil)
	if err != nil {
		t.Logf("GetDisplayModeList (query count) failed: %v", err)
		return
	}

	t.Logf("Number of display modes: %d", numModes)

	if numModes > 0 {
		// 获取第一个模式
		modes := make([]DXGIModeDesc, numModes)
		err = output.GetDisplayModeList(0, 0, &numModes, &modes[0])
		if err != nil {
			t.Logf("GetDisplayModeList failed: %v", err)
		} else {
			t.Logf("First mode: %dx%d @ %d/%d Hz",
				modes[0].Width,
				modes[0].Height,
				modes[0].RefreshRate.Numerator,
				modes[0].RefreshRate.Denominator)
		}
	}
}

func TestDXGIOutput_GetGammaControlCapabilities(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping Windows-specific test in short mode")
	}

	output := getTestOutput(t)
	defer output.Release()

	caps, err := output.GetGammaControlCapabilities()
	if err != nil {
		t.Logf("GetGammaControlCapabilities failed: %v", err)
		return
	}

	if caps == nil {
		t.Fatal("GetGammaControlCapabilities returned nil")
	}

	t.Logf("Gamma ScaleAndOffsetSupported: %v", caps.ScaleAndOffsetSupported)
	t.Logf("Gamma NumGammaControlPoints: %d", caps.NumGammaControlPoints)
}

func TestDXGIOutput_GetGammaControl(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping Windows-specific test in short mode")
	}

	output := getTestOutput(t)
	defer output.Release()

	gamma, err := output.GetGammaControl()
	if err != nil {
		t.Logf("GetGammaControl failed: %v", err)
		return
	}

	if gamma == nil {
		t.Fatal("GetGammaControl returned nil")
	}

	t.Log("GetGammaControl succeeded")
}

func TestDXGIOutput_GetFrameStatistics(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping Windows-specific test in short mode")
	}

	output := getTestOutput(t)
	defer output.Release()

	stats, err := output.GetFrameStatistics()
	if err != nil {
		t.Logf("GetFrameStatistics failed: %v", err)
		return
	}

	if stats == nil {
		t.Fatal("GetFrameStatistics returned nil")
	}

	t.Logf("Frame PresentCount: %d", stats.PresentCount)
	t.Logf("Frame SyncRefreshCount: %d", stats.SyncRefreshCount)
}

func TestDXGIOutput_SupportsOverlays(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping Windows-specific test in short mode")
	}

	output := getTestOutput(t)
	defer output.Release()

	supports := output.SupportsOverlays()
	t.Logf("SupportsOverlays: %v", supports)
}

func TestDXGIOutput_CheckOverlaySupport(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping Windows-specific test in short mode")
	}

	output := getTestOutput(t)
	defer output.Release()

	flags, err := output.CheckOverlaySupport(0, nil)
	if err != nil {
		t.Logf("CheckOverlaySupport failed: %v", err)
		return
	}

	t.Logf("Overlay support flags: 0x%x", flags)
}

func TestDXGIOutput_CheckHardwareCompositionSupport(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping Windows-specific test in short mode")
	}

	output := getTestOutput(t)
	defer output.Release()

	flags, err := output.CheckHardwareCompositionSupport()
	if err != nil {
		t.Logf("CheckHardwareCompositionSupport failed (may not be supported): %v", err)
		return
	}

	t.Logf("Hardware composition support flags: 0x%x", flags)
}
