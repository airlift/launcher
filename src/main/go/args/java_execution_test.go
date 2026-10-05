package args

import (
	"os/exec"
	"slices"
	"testing"
)

func TestParsingJvmVersion(t *testing.T) {
	testCases := []struct {
		output          string
		expectedVersion string
	}{
		{
			output:          "openjdk version \"24-ea\" 2025-03-18\nOpenJDK Runtime Environment (build 24-ea+24-2960)\nOpenJDK 64-Bit Server VM (build 24-ea+24-2960, mixed mode, sharing)",
			expectedVersion: "24-ea",
		},
		{
			output:          "openjdk version \"23\" 2024-09-17\nOpenJDK Runtime Environment Temurin-23+37 (build 23+37)\nOpenJDK 64-Bit Server VM Temurin-23+37 (build 23+37, mixed mode, sharing)",
			expectedVersion: "23",
		},
	}

	for _, testCase := range testCases {
		actualVersion, err := ExtractJvmVersion(testCase.output)
		if err != nil {
			t.Fatalf("Unexpected error while extracting JVM version: %v", err)
		}
		if actualVersion != testCase.expectedVersion {
			t.Fatalf("Expected Java version %s for output %s", testCase.expectedVersion, testCase.output)
		}
	}
}

func TestJavaExecutionMultipleJvmOptions(t *testing.T) {
	if _, err := exec.LookPath("java"); err != nil {
		t.Skip("java not on PATH")
	}

	options := &Options{
		JvmConfig:        []string{},
		JvmOptions:       []string{"-Xmx4g", "-Xms2g"},
		LauncherConfig:   map[string]string{"main-class": "com.example.Main"},
		SystemProperties: map[string]string{},
		ConfigPath:       "/etc/config.properties",
	}

	command, _, err := options.JavaExecution(false)
	if err != nil {
		t.Fatalf("JavaExecution failed: %v", err)
	}

	found := false
	for i := 0; i < len(command)-1; i++ {
		if command[i] == "-Xmx4g" && command[i+1] == "-Xms2g" {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("expected -Xmx4g and -Xms2g as consecutive separate arguments, got: %v", command)
	}
}

func TestJavaExecutionEnvConfig(t *testing.T) {
	if _, err := exec.LookPath("java"); err != nil {
		t.Skip("java not on PATH")
	}
	t.Setenv("MALLOC_ARENA_MAX", "8")
	t.Setenv("LAUNCHER_TEST_INHERITED", "inherited")

	options := &Options{
		JvmConfig:        []string{},
		LauncherConfig:   map[string]string{"main-class": "com.example.Main"},
		SystemProperties: map[string]string{},
		EnvConfig:        map[string]string{"MALLOC_ARENA_MAX": "4", "LAUNCHER_TEST_ADDED": "added"},
		ConfigPath:       "/etc/config.properties",
	}

	_, env, err := options.JavaExecution(false)
	if err != nil {
		t.Fatalf("JavaExecution failed: %v", err)
	}

	for _, expected := range []string{"MALLOC_ARENA_MAX=4", "LAUNCHER_TEST_ADDED=added", "LAUNCHER_TEST_INHERITED=inherited"} {
		if !slices.Contains(env, expected) {
			t.Errorf("expected environment to contain %s, got: %v", expected, env)
		}
	}
	if slices.Contains(env, "MALLOC_ARENA_MAX=8") {
		t.Errorf("expected env config to override inherited MALLOC_ARENA_MAX, got: %v", env)
	}
}
