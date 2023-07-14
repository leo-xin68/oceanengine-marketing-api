package enums

// InventoryType : 首选投放位置
type InventoryType string

const (
	FeedInventoryType              InventoryType = "INVENTORY_FEED"
	TextLinkInventoryType          InventoryType = "INVENTORY_TEXT_LINK"
	VideoFeedInventoryType         InventoryType = "INVENTORY_VIDEO_FEED"
	HotsoonFeedInventoryType       InventoryType = "INVENTORY_HOTSOON_FEED"
	AwemeFeedInventoryType         InventoryType = "INVENTORY_AWEME_FEED"
	UnionSlotInventoryType         InventoryType = "INVENTORY_UNION_SLOT"
	UnionBoutiqueGameInventoryType InventoryType = "UNION_BOUTIQUE_GAME"
	UnionSplashSlotInventoryType   InventoryType = "INVENTORY_UNION_SPLASH_SLOT"
	SearchInventoryType            InventoryType = "INVENTORY_SEARCH"
	UniversalInventoryType         InventoryType = "INVENTORY_UNIVERSAL"
	BeautyInventoryType            InventoryType = "INVENTORY_BEAUTY"
	PipixiaInventoryType           InventoryType = "INVENTORY_PIPIXIA"
	AutomobileInventoryType        InventoryType = "INVENTORY_AUTOMOBILE"
	StudyInventoryType             InventoryType = "INVENTORY_STUDY"
	FaceUInventoryType             InventoryType = "INVENTORY_FACE_U"
	TomatoNovelInventoryType       InventoryType = "INVENTORY_TOMATO_NOVEL"
	HomedAggregateInventoryType    InventoryType = "INVENTORY_HOMED_AGGREGATE"
)

func (t InventoryType) String() string {
	switch t {
	case FeedInventoryType:
		return "今日头条"
	case TextLinkInventoryType:
		return "头条文章详情页"
	case VideoFeedInventoryType:
		return "西瓜信息流"
	case HotsoonFeedInventoryType:
		return "火山信息流"
	case AwemeFeedInventoryType:
		return "抖音短视频"
	case UnionSlotInventoryType:
		return "穿山甲"
	case UnionBoutiqueGameInventoryType:
		return "ohayoo精品游戏"
	case UnionSplashSlotInventoryType:
		return "穿山甲开屏广告"
	case SearchInventoryType:
		return "搜索广告"
	case UniversalInventoryType:
		return "通投智选"
	case BeautyInventoryType:
		return "轻颜相机"
	case PipixiaInventoryType:
		return "皮皮虾"
	case AutomobileInventoryType:
		return "懂车帝"
	case StudyInventoryType:
		return "好好学习"
	case FaceUInventoryType:
		return "faceu"
	case TomatoNovelInventoryType:
		return "番茄小说"
	case HomedAggregateInventoryType:
		return "住小帮"
	default:
		return ""
	}
}
