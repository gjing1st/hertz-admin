package slice

// ReverseString
//	@description:	字符串反转
//	@param:
//	@author:	GJing
//	@email:		gjing1st@gmail.com
//	@date:		2022/4/6 12:10
//	@success:
func ReverseString(s string) string {
	sByte := []byte(s)
	for i, j := 0, len(sByte)-1; i < j; i, j = i+1, j-1 {
		sByte[i], sByte[j] = sByte[j], sByte[i]
	}
	return string(sByte)
}
