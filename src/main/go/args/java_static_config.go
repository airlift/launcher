package args

const vectorApiIncubating = "--add-modules=jdk.incubator.vector"
const allowUnsafeUsage = "--sun-misc-unsafe-memory-access=allow"
const allowNativeUsage = "--enable-native-access=ALL-UNNAMED"

var jvmSpecificConfig = map[string][]string{
	"22": {vectorApiIncubating},
	"23": {vectorApiIncubating},
	"24": {vectorApiIncubating, allowUnsafeUsage, allowNativeUsage},
	"25": {vectorApiIncubating, allowUnsafeUsage, allowNativeUsage},
	"26": {vectorApiIncubating, allowUnsafeUsage, allowNativeUsage},
}
