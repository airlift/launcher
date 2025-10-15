package args

import "runtime"

var rewriteBytecodes = onArm64Only("-XX:-RewriteBytecodes") // https://bugs.openjdk.org/browse/JDK-8369506
const vectorApiIncubating = "--add-modules=jdk.incubator.vector"
const allowUnsafeUsage = "--sun-misc-unsafe-memory-access=allow"
const allowNativeUsage = "--enable-native-access=ALL-UNNAMED"

var jvmSpecificConfig = map[string][]string{
	"21": {rewriteBytecodes},
	"22": {rewriteBytecodes, vectorApiIncubating},
	"23": {rewriteBytecodes, vectorApiIncubating},
	"24": {rewriteBytecodes, vectorApiIncubating, allowUnsafeUsage, allowNativeUsage},
	"25": {rewriteBytecodes, vectorApiIncubating, allowUnsafeUsage, allowNativeUsage},
	"26": {rewriteBytecodes, vectorApiIncubating, allowUnsafeUsage, allowNativeUsage},
}

func onArm64Only(option string) string {
	if runtime.GOARCH == "arm64" {
		return option
	}
	return ""
}
