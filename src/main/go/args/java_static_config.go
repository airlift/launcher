package args

import "runtime"

var rewriteBytecodes = onArm64Only("-XX:-RewriteBytecodes") // https://bugs.openjdk.org/browse/JDK-8369506
const vectorApiIncubating = "--add-modules=jdk.incubator.vector"

// https://openjdk.org/jeps/498, allows deprecated unsafe usage
const allowUnsafeUsage = "--sun-misc-unsafe-memory-access=allow"

// https://openjdk.org/jeps/472, allows native libraries to be loaded from unnamed modules
const allowNativeUsage = "--enable-native-access=ALL-UNNAMED"

// https://openjdk.org/jeps/451, required by JOL to estimate instances size
const allowDynamicAgentLoading = "-XX:+EnableDynamicAgentLoading"

var jvmSpecificConfig = map[string][]string{
	"21": {rewriteBytecodes, allowDynamicAgentLoading},
	"22": {rewriteBytecodes, allowDynamicAgentLoading, vectorApiIncubating},
	"23": {rewriteBytecodes, allowDynamicAgentLoading, vectorApiIncubating},
	"24": {rewriteBytecodes, allowDynamicAgentLoading, vectorApiIncubating, allowUnsafeUsage, allowNativeUsage},
	"25": {rewriteBytecodes, allowDynamicAgentLoading, vectorApiIncubating, allowUnsafeUsage, allowNativeUsage},
	"26": {rewriteBytecodes, allowDynamicAgentLoading, vectorApiIncubating, allowUnsafeUsage, allowNativeUsage},
}

func onArm64Only(option string) string {
	if runtime.GOARCH == "arm64" {
		return option
	}
	return ""
}
