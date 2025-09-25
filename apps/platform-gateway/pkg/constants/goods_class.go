package constants

const (
	GOODS_CLASS_FOOD        = 1  // 美食
	GOODS_CLASS_CAKE        = 2  // 蛋糕
	GOODS_CLASS_FLOWER      = 3  // 鲜花
	GOODS_CLASS_FRESH       = 4  // 生鲜
	GOODS_CLASS_MEDICINE    = 5  // 药品
	GOODS_CLASS_DOCUMENT    = 6  // 文件
	GOODS_CLASS_DIGITAL     = 7  // 数码
	GOODS_CLASS_CLOTHING    = 8  // 服装
	GOODS_CLASS_CAR         = 9  // 汽配
	GOODS_CLASS_BBQ         = 10 // 烧烤
	GOODS_CLASS_HOTPOT      = 11 // 火锅
	GOODS_CLASS_OTHER       = 12 // 其他
	GOODS_CLASS_CONVENIENCE = 13 // 便利店
	GOODS_CLASS_BAKERY      = 14 // 面包糕点
)

// 商品类型名称映射
var GoodsClassNames = map[int]string{
	GOODS_CLASS_FOOD:        "美食",
	GOODS_CLASS_CAKE:        "蛋糕",
	GOODS_CLASS_FLOWER:      "鲜花",
	GOODS_CLASS_FRESH:       "生鲜",
	GOODS_CLASS_MEDICINE:    "药品",
	GOODS_CLASS_DOCUMENT:    "文件",
	GOODS_CLASS_DIGITAL:     "数码",
	GOODS_CLASS_CLOTHING:    "服装",
	GOODS_CLASS_CAR:         "汽配",
	GOODS_CLASS_BBQ:         "烧烤",
	GOODS_CLASS_HOTPOT:      "火锅",
	GOODS_CLASS_OTHER:       "其他",
	GOODS_CLASS_CONVENIENCE: "便利店",
	GOODS_CLASS_BAKERY:      "面包糕点",
}
