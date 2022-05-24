package foo

const (
	// 頭文字が大文字だと外部でも参照が可能
	Max = 100
	// 小文字だと外部からの参照ができない
	min = 1
)

func ReturnMin() int {
	return min
}