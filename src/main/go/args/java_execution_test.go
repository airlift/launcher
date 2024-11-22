package args

import "testing"

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
