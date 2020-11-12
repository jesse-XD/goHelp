package help

// GetStrCn 获取字符串里的中文
func GetStrCn(str string)(cnStr string){
	r := []rune(str)
	strSlice := []string{}
	for i := 0; i < len(r); i++ {
		if r[i] <= 40869 && r[i] >= 19968 {
			cnStr = cnStr + string(r[i])
			strSlice = append(strSlice, cnStr)

		}
	}
	return
}