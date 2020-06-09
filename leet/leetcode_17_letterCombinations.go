package leet


//17. 电话号码的字母组合
//回溯算法解决
var numMap map[byte]string = map[byte]string{
	'2':"abc",
	'3':"def",
	'4':"ghi",
	'5':"jkl",
	'6':"mno",
	'7':"pqrs",
	'8':"tuv",
	'9':"wxyz",
}

func letterCombinations(digits string) []string {
	var results []string
	if digits == "" {
		return results
	}

	doJoin1(digits,"",0, &results)
	return results
}

func doJoin1(digits string, str string, hight int, results *[]string) {
	value := numMap[digits[hight]]
	if hight == len(digits)-1 {
		for _, v := range value {
			*results = append(*results, str + string(v))
		}
		return
	}
	for _, v := range value {
		doJoin1(digits, str + string(v), hight+1,results)
	}
}