package code

import "testing"

func TestLevel(t *testing.T) {
	code := "100.01.02"
	level := Level(code)
	if level != 3 {
		t.Errorf("测试失败，active level: %v", level)
	}
}

func TestParent(t *testing.T) {
	hopeCode := "100.01"
	actualCode := Parent("100.01.02")
	if hopeCode != actualCode {
		t.Errorf("测试失败, actual code: %v", actualCode)
	}
}

func TestNext(t *testing.T) {
	hopeCode := "100.02.09"
	actualCode, err := Next("100.02.08")
	if err != nil {
		t.Errorf("error, detail: %v", err)
	}
	if hopeCode != actualCode {
		t.Errorf("test error, actual next code: %v", actualCode)
	}
}
