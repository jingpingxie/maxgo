package constant

type OrderStatus int

//https://blog.csdn.net/weixin_43364551/article/details/117848680
//https://www.javamall.com.cn/docs/7.2.3_upgrade/achitecture/jiagou/ding-dan/ding-dan-lei-xing-ji-zhuang-tai/
//参考这个
//https://www.pmcaff.com/discuss/1000000000170986?pmc_param%5Bentry_id%5D=2000000000007562
const (
	Unconfirmed   = 0b0000 //未确认
	Confirmed     = 0b0001 //已确认
	Dispatch      = 0b0010 //已派单
	PAID_OFF      = 0b0100 //已付款
	SHIPPED                //已发货
	SHIPROG                //已收货
	COMPLETE               //已完成
	CANCELLED              //已取消
	AFTER_SERVICE          //售后中
	Offline                //线下支付
	cod                    //货到付款
	comment                //评论
	PAY_PARTIAL            //部分支付
)
