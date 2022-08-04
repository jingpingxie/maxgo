//
// @File:passport
// @Version:1.0.0
// @Description:
// @Author:jingpingxie
// @Date:2022/8/2 17:52
//
package controllers

import "github.com/astaxie/beego"

type PassportController struct {
	beego.Controller
}

//https://passport.baidu.com/v2/getpublickey?token=&tpl=mn&subpro=&apiver=v3&tt=1659432607942&gid=161A899-C771-43EE-8B04-2B0C51178118&loginversion=v4&traceid=&time=1659432608&alg=v3&sig=N3pLbTV0VlVmUExtR1VOTnZrS09wN2NFbWh6S2V5TjMrS1pWRURlUkR6b3hpL0p3Wjh0Z1BCbWFrWFRLcWJwbw%3D%3D&elapsed=17&shaOne=00b836cb57b5f07d4d090423203ed4efbc1cc580&callback=bd__cbs__i09f6d
//
// @Title:publickey
// @Description:获取公钥
// @Author:jingpingxie
// @Date:2022-08-02 17:54:58
// @Receiver:pc
// @Return:string
//
func (pc *PassportController) PublicKey() string {

	return ""
	//({"errno":'0',"msg":'',"pubkey":'-----BEGIN PUBLIC KEY-----\nMIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCo89MpUd6f3evjZh7elj87NYfy\nDmXuyq26Ex3\/KnVEiQMTSGty1PyPG+Gb97iihg2Wn0D7RiOKV\/Vnfs\/+3RhJNx9W\nmW5PGJ\/kae9+wvGGL7esaDiT1tMWNnG6xMjGLdmyMObvaFsrDzUtGb+VT0HyM0GO\nMKnBrPb0KtLk3apQ1QIDAQAB\n-----END PUBLIC KEY-----\n',"key":'ff3El8tDqFVZoT4iR7gEq0elrroJa15Y',    "traceid": ""})

}
