package solution

import "testing"

func Test_minWindow_01(t *testing.T) {
	ss := "ADOBECODEBANC"
	tt := "ABC"
	expected := "BANC"
	ans := minWindow(ss, tt)
	if ans != expected {
		t.Errorf("expected: %v, ans: %v", expected, ans)
	}
}

func Test_minWindow_02(t *testing.T) {
	ss := "a"
	tt := "aa"
	expected := ""
	ans := minWindow(ss, tt)
	if ans != expected {
		t.Errorf("expected: %v, ans: %v", expected, ans)
	}
}

func Test_minWindow_03(t *testing.T) {
	ss := "AB"
	tt := "a"
	expected := ""
	ans := minWindow(ss, tt)
	if ans != expected {
		t.Errorf("expected: %v, ans: %v", expected, ans)
	}
}

func Test_minWindow_04(t *testing.T) {
	ss := "wyqaalfdtavrmkvrgbrmauenibfxrzjpzzqzlveexayjkvdsyueboloymtvfugvtgnutkyzhaztlvwdvqkwgvmejhbpdimwqbslnrkutkpehnkefwblrprcxvtaffzxitivmssgehklvwqastojihmhcfkhnlexemtrhnxlujxrgpuyiikspycufodubisfwnydaxrwhqqpfkppuzjlzlfhbjbcttkriixkiohpexgjjvafxjqyvyfyjhbccltlvsvdgeumdavoyxtvhmtekzctidxkqsxmlvrrzmefobtmznhizdmlcoemudwkvuyirscqegvsjrfkgoshrgsvvyhrbgdycehtwjlcrjucabpgsjnjqhhnfqeiwhgalptjyflpoiuqjjwdslpiswvxobfljnnwdhgdortezpulysoqddbxbwuqabdjqqhtzpxpjsvkjrvhjmzoralvzhlzkqkbgrwijvzspvcymafymfmfhaaghnfsdrvmlruuntmcqqbdqideprkxrmfbanvfeqrphnlwjxbzqcegmhnczxbslitnvotaemroadkjclksppzeyoiznlsytnopchritiyvlleqypiqgjugxeikpclipzxtgoebxcxkpdaoulecuajueretvpbkqbgwrkaooxbeaduvoaxlaifgblzwdcjtfpsxbsnrrovturokrovtycbcqcytfjomygj"
	tt := "baxtr"
	expected := "blrprcxvta"
	ans := minWindow(ss, tt)
	if ans != expected {
		t.Errorf("expected: %v, ans: %v", expected, ans)
	}
}
