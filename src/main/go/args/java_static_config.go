package args

import (
	"fmt"
	"runtime"
	"strconv"
)

var rewriteBytecodes = onArm64Only("-XX:-RewriteBytecodes") // https://bugs.openjdk.org/browse/JDK-8369506

// https://openjdk.org/jeps/448, enables incubating vector API
const vectorApiIncubating = "--add-modules=jdk.incubator.vector"

// https://openjdk.org/jeps/498, allows deprecated unsafe usage
const allowUnsafeUsage = "--sun-misc-unsafe-memory-access=allow"

// https://openjdk.org/jeps/472, allows native libraries to be loaded from unnamed modules
const allowNativeUsage = "--enable-native-access=ALL-UNNAMED"

// https://openjdk.org/jeps/451, required by JOL to estimate instances size
const allowDynamicAgentLoading = "-XX:+EnableDynamicAgentLoading"

// https://openjdk.org/jeps/519, reduces object header size to 64 bit
const compactObjectHeaders = "-XX:+UseCompactObjectHeaders"

var jvmSpecificConfig = map[string][]string{
	"21": {rewriteBytecodes, allowDynamicAgentLoading},
	"22": {rewriteBytecodes, allowDynamicAgentLoading, vectorApiIncubating},
	"23": {rewriteBytecodes, allowDynamicAgentLoading, vectorApiIncubating},
	"24": {rewriteBytecodes, allowDynamicAgentLoading, vectorApiIncubating, allowUnsafeUsage, allowNativeUsage},
	"25": {rewriteBytecodes, allowDynamicAgentLoading, vectorApiIncubating, allowUnsafeUsage, allowNativeUsage, compactObjectHeaders},
	"26": {allowDynamicAgentLoading, vectorApiIncubating, allowUnsafeUsage, allowNativeUsage, compactObjectHeaders},
	"27": {allowDynamicAgentLoading, vectorApiIncubating, allowUnsafeUsage, allowNativeUsage},
	"28": {allowDynamicAgentLoading, vectorApiIncubating, allowUnsafeUsage, allowNativeUsage},
}

// getJvmSpecificConfig returns the static config for the given major JVM version,
// falling back to the latest known config when the version is newer than any known one.
// The second return value is the version whose config was selected.
func getJvmSpecificConfig(majorJavaVersion string) ([]string, string, error) {
	if config, exists := jvmSpecificConfig[majorJavaVersion]; exists {
		return config, majorJavaVersion, nil
	}

	version, err := strconv.Atoi(majorJavaVersion)
	if err != nil {
		return nil, "", fmt.Errorf("unrecognized major JVM version: %s", majorJavaVersion)
	}

	latestVersion := 0
	for knownVersion := range jvmSpecificConfig {
		known, err := strconv.Atoi(knownVersion)
		if err == nil && known > latestVersion {
			latestVersion = known
		}
	}

	if version > latestVersion {
		latest := strconv.Itoa(latestVersion)
		return jvmSpecificConfig[latest], latest, nil
	}

	return nil, "", fmt.Errorf("no static config for major JVM version: %s", majorJavaVersion)
}

func onArm64Only(option string) string {
	if runtime.GOARCH == "arm64" {
		return option
	}
	return ""
}
