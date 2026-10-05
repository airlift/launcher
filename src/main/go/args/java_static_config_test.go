package args

import (
	"slices"
	"testing"
)

func TestGetJvmSpecificConfig(t *testing.T) {
	testCases := []struct {
		majorJavaVersion string
		expectedVersion  string
		expectedError    string
	}{
		{majorJavaVersion: "21", expectedVersion: "21"},
		{majorJavaVersion: "28", expectedVersion: "28"},
		{majorJavaVersion: "29", expectedVersion: "28"},
		{majorJavaVersion: "99", expectedVersion: "28"},
		{majorJavaVersion: "20", expectedError: "no static config for major JVM version: 20"},
		{majorJavaVersion: "1", expectedError: "no static config for major JVM version: 1"},
		{majorJavaVersion: "unknown", expectedError: "unrecognized major JVM version: unknown"},
	}

	for _, testCase := range testCases {
		config, version, err := getJvmSpecificConfig(testCase.majorJavaVersion)
		if testCase.expectedError != "" {
			if err == nil {
				t.Fatalf("Expected error for version %s, got config of version %s", testCase.majorJavaVersion, version)
			}
			if err.Error() != testCase.expectedError {
				t.Fatalf("Expected error %q for version %s, got %q", testCase.expectedError, testCase.majorJavaVersion, err.Error())
			}
			continue
		}
		if err != nil {
			t.Fatalf("Unexpected error for version %s: %v", testCase.majorJavaVersion, err)
		}
		if version != testCase.expectedVersion {
			t.Fatalf("Expected config version %s for version %s, got %s", testCase.expectedVersion, testCase.majorJavaVersion, version)
		}
		if !slices.Equal(config, jvmSpecificConfig[testCase.expectedVersion]) {
			t.Fatalf("Expected config of version %s for version %s, got %v", testCase.expectedVersion, testCase.majorJavaVersion, config)
		}
	}
}
