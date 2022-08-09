//
// @File:string
// @Version:1.0.0
// @Description:
// @Author:jingpingxie
// @Date:2022/8/9 11:54
//
package xstring

import "regexp"

//
// @Title:StringStrip
// @Description:remove all invalid tags in the text includes "\t","\r","\n"
// @Author:jingpingxie
// @Date:2022-08-09 11:54:34
// @Param:input
// @Return:string
//
func StringStrip(input string) string {
	if input == "" {
		return ""
	}
	reg := regexp.MustCompile(`[\s\p{Zs}]{1,}`)
	return reg.ReplaceAllString(input, "")
}
