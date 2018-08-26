package rule

import (
	. "github.com/henrylee2cn/pholcus/app/spider"
)

func city58_parse(ctx *Context) {
	query := ctx.GetDom()

	var 标题, 价格, 单位, 出租价格, 出租单位, 面积, 类型, 经营状态, 历史经营, 付款方式, 租约方式, 详细地址, 发布时间, 发布人, 联系电话, 详情 string

	标题 = query.Find("div.house-title h1").Text()
	价格 = query.Find(".house_basic_title_money_num").Text()
	单位 = query.Find(".house_basic_title_money_unit").Text()
	出租价格 = query.Find(".house_basic_title_money_num_chuzu").Text()
	出租单位 = query.Find(".house_basic_title_money_unit_chuzu").Text()
	面积 = query.Find(".house-basic-right .house_basic_title_content").Eq(1).Find(".house_basic_title_content_item2").Text()




	ctx.Output(map[int]interface{}{
		0:  标题,
		1:  价格,
		2:  单位,
		3:  出租价格,
		4:  出租单位,
		5:  面积,
		6:  类型,
		7:  经营状态,
		8:  历史经营,
		9:  付款方式,
		10: 租约方式,
		11: 详细地址,
		12: 发布时间,
		13: 发布人,
		14: 联系电话,
		15: 详情,
	})
}
