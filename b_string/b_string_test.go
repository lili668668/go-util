package b_string

import "testing"
import "fmt"

func TestIndexFrom(t *testing.T) {
	t.Log("#IndexFrom 從字串 str 中的 from 開始尋找 find 的位置，回傳 int")
	str := "我是測試字串，測試用的測試字串"
	find := "測試"
	from := 12

	exp := 21
	result := IndexFrom(str, find, from)

	if result == exp {
		t.Log("通過")
	} else {
		error_str := fmt.Sprintf("錯誤：\nresult: %d", result)
		t.Error(error_str)
	}
}

func TestCapture(t *testing.T) {
	t.Log("#Capture 從字串 str 中，把 start 與 end 這兩個子字串中間的字串取出，回傳取出的字串")
	str := "{我是字串{ㄏㄏ}}"
	str2 := "襪~~~}owo"
	start := "{"
	end := "}"

	exp := "我是字串{ㄏㄏ}"
	exp2 := "襪~~~}owo"
	result := Capture(str, start, end)
	result2 := Capture(str2, start, end)

	if result == exp {
		t.Log("通過")
	} else {
		error_str := fmt.Sprintf("錯誤：\nresult: %s", result)
		t.Error(error_str)
	}

	if result2 == exp2 {
		t.Log("通過")
	} else {
		error_str := fmt.Sprintf("錯誤：\nresult: %s", result2)
		t.Error(error_str)
	}
}

func TestCaptureFrom(t *testing.T) {
	t.Log("#Capture 從字串 str 中第 from 個位置以後，把 start 與 end 這兩個子字串中間的字串取出，回傳取出的字串")
	str := "{asdfsdf}{我是字串}"
	start := "{"
	end := "}"

	exp := "我是字串"
	result := CaptureFrom(str, start, end, 9)

	if result == exp {
		t.Log("通過")
	} else {
		error_str := fmt.Sprintf("錯誤：\nresult: %s", result)
		t.Error(error_str)
	}
}

func TestTrimAll(t *testing.T) {
	t.Log("#TrimAll 把字串 origin 裡的所有空白全部去除，並回傳全部去除空白的字串")
	origin := "襪~~  你好owo \t  \n 襪~~~   本寶寶最美    \n    耶 ~~   owo   "

	exp := "襪~~你好owo襪~~~本寶寶最美耶~~owo"
	result := TrimAll(origin)

	if result == exp {
		t.Log("通過")
	} else {
		error_str := fmt.Sprintf("錯誤：\nresult: %s", result)
		t.Error(error_str)
	}
}

func TestMinIndex(t *testing.T) {
	t.Log("#MinIndex 在 str 中找出兩個不同標記 mark1, mark2 的最小位置，比較兩個位置，回傳比較小的 mark 字串與位置")
	str1 := "{([fjsdl;kfjas])}"
	str2 := "{lkdfgjsdfl;g}"
	str3 := "askdlfaj;sd"

	exp_str1 := "{"
	exp_str2 := "{"
	exp_str3 := ""
	exp_int1 := 0
	exp_int2 := 0
	exp_int3 := -1

	res_str1, res_int1 := MinIndex(str1, "{", "[")
	res_str2, res_int2 := MinIndex(str2, "{", "[")
	res_str3, res_int3 := MinIndex(str3, "{", "[")

	if res_str1 != exp_str1 || res_int1 != exp_int1 {
		t.Error("No Pass: result1")
	}

	if res_str2 != exp_str2 || res_int2 != exp_int2 {
		t.Error("No Pass: result2")
	}

	if res_str3 != exp_str3 || res_int3 != exp_int3 {
		t.Error("No Pass: result3")
	}
}
