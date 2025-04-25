package args

const vectorApiIncubating = "--add-modules=jdk.incubator.vector"
const allowUnsafeUsage = "--sun-misc-unsafe-memory-access=allow"

var jvmSpecificConfig = map[string][]string{
	"22": {vectorApiIncubating},
	"23": {vectorApiIncubating},
	"24": {vectorApiIncubating, allowUnsafeUsage},
}
