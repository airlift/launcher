package args

const vectorApiIncubating = "--add-modules=jdk.incubator.vector"

var jvmSpecificConfig = map[string][]string{
	"22": {vectorApiIncubating},
	"23": {vectorApiIncubating},
	// 24 is not released yet, but vector API will be still incubating
	"24": {vectorApiIncubating},
}
