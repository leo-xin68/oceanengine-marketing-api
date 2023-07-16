package enums

// InventoryType : 首选投放位置
type InventoryType string

const (
	InventoryTypeIsFeed              InventoryType = "INVENTORY_FEED"
	InventoryTypeIsTextLink          InventoryType = "INVENTORY_TEXT_LINK"
	InventoryTypeIsVideoFeed         InventoryType = "INVENTORY_VIDEO_FEED"
	InventoryTypeIsHotsoonFeed       InventoryType = "INVENTORY_HOTSOON_FEED"
	InventoryTypeIsAwemeFeed         InventoryType = "INVENTORY_AWEME_FEED"
	InventoryTypeIsUnionSlot         InventoryType = "INVENTORY_UNION_SLOT"
	InventoryTypeIsUnionBoutiqueGame InventoryType = "UNION_BOUTIQUE_GAME"
	InventoryTypeIsUnionSplashSlot   InventoryType = "INVENTORY_UNION_SPLASH_SLOT"
	InventoryTypeIsSearch            InventoryType = "INVENTORY_SEARCH"
	InventoryTypeIsUniversal         InventoryType = "INVENTORY_UNIVERSAL"
	InventoryTypeIsBeauty            InventoryType = "INVENTORY_BEAUTY"
	InventoryTypeIsPipixia           InventoryType = "INVENTORY_PIPIXIA"
	InventoryTypeIsAutomobile        InventoryType = "INVENTORY_AUTOMOBILE"
	InventoryTypeIsStudy             InventoryType = "INVENTORY_STUDY"
	InventoryTypeIsFaceU             InventoryType = "INVENTORY_FACE_U"
	InventoryTypeIsTomatoNovel       InventoryType = "INVENTORY_TOMATO_NOVEL"
	InventoryTypeIsHomedAggregate    InventoryType = "INVENTORY_HOMED_AGGREGATE"
)

func (t InventoryType) String() string {
	switch t {
	case InventoryTypeIsFeed:
		return "今日头条"
	case InventoryTypeIsTextLink:
		return "头条文章详情页"
	case InventoryTypeIsVideoFeed:
		return "西瓜信息流"
	case InventoryTypeIsHotsoonFeed:
		return "火山信息流"
	case InventoryTypeIsAwemeFeed:
		return "抖音短视频"
	case InventoryTypeIsUnionSlot:
		return "穿山甲"
	case InventoryTypeIsUnionBoutiqueGame:
		return "ohayoo精品游戏"
	case InventoryTypeIsUnionSplashSlot:
		return "穿山甲开屏广告"
	case InventoryTypeIsSearch:
		return "搜索广告"
	case InventoryTypeIsUniversal:
		return "通投智选"
	case InventoryTypeIsBeauty:
		return "轻颜相机"
	case InventoryTypeIsPipixia:
		return "皮皮虾"
	case InventoryTypeIsAutomobile:
		return "懂车帝"
	case InventoryTypeIsStudy:
		return "好好学习"
	case InventoryTypeIsFaceU:
		return "faceu"
	case InventoryTypeIsTomatoNovel:
		return "番茄小说"
	case InventoryTypeIsHomedAggregate:
		return "住小帮"
	default:
		return ""
	}
}
