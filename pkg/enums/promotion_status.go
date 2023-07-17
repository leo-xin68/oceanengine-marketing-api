package enums

// PromotionStatus : 广告状态
type PromotionStatus string

const (
	PromotionStatusIsNotDeleted                 PromotionStatus = "NOT_DELETED"
	PromotionStatusIsAll                        PromotionStatus = "All"
	PromotionStatusIsOk                         PromotionStatus = "OK"
	PromotionStatusIsDeleted                    PromotionStatus = "DELETED"
	PromotionStatusIsProjectOfflineBudget       PromotionStatus = "PROJECT_OFFLINE_BUDGET"
	PromotionStatusIsProjectPreofflineBudget    PromotionStatus = "PROJECT_PREOFFLINE_BUDGET"
	PromotionStatusIsTimeNoReach                PromotionStatus = "TIME_NO_REACH"
	PromotionStatusIsTimeDone                   PromotionStatus = "TIME_DONE"
	PromotionStatusIsNoSchedule                 PromotionStatus = "NO_SCHEDULE"
	PromotionStatusIsAudit                      PromotionStatus = "AUDIT"
	PromotionStatusIsReaudit                    PromotionStatus = "REAUDIT"
	PromotionStatusIsFrozen                     PromotionStatus = "FROZEN"
	PromotionStatusIsAuditDeny                  PromotionStatus = "AUDIT_DENY"
	PromotionStatusIsOfflineBudget              PromotionStatus = "OFFLINE_BUDGET"
	PromotionStatusIsOfflineBalance             PromotionStatus = "OFFLINE_BALANCE"
	PromotionStatusIsPreofflineBudget           PromotionStatus = "PREOFFLINE_BUDGET"
	PromotionStatusIsDisabled                   PromotionStatus = "DISABLED"
	PromotionStatusIsProjectDisabled            PromotionStatus = "PROJECT_DISABLED"
	PromotionStatusIsLiveRoomOff                PromotionStatus = "LIVE_ROOM_OFF"
	PromotionStatusIsProductOffline             PromotionStatus = "PRODUCT_OFFLINE"
	PromotionStatusIsAwemeAccountDisabled       PromotionStatus = "AWEME_ACCOUNT_DISABLED"
	PromotionStatusIsAwemeAnchorDisabled        PromotionStatus = "AWEME_ANCHOR_DISABLED"
	PromotionStatusIsDisableByQuota             PromotionStatus = "DISABLE_BY_QUOTA"
	PromotionStatusIsCreate                     PromotionStatus = "CREATE"
	PromotionStatusIsAdvertiserOfflineBudget    PromotionStatus = "ADVERTISER_OFFLINE_BUDGET"
	PromotionStatusIsAdvertiserPreofflineBudget PromotionStatus = "ADVERTISER_PREOFFLINE_BUDGET"
)

func (t PromotionStatus) String() string {
	switch t {
	case PromotionStatusIsNotDeleted:
		return "不限"
	case PromotionStatusIsAll:
		return "不限（包含已删除）"
	case PromotionStatusIsOk:
		return "投放中"
	case PromotionStatusIsDeleted:
		return "已删除"
	case PromotionStatusIsProjectOfflineBudget:
		return "项目超出预算"
	case PromotionStatusIsProjectPreofflineBudget:
		return "项目接近预算"
	case PromotionStatusIsTimeNoReach:
		return "未到达投放时间"
	case PromotionStatusIsTimeDone:
		return "已完成"
	case PromotionStatusIsNoSchedule:
		return "不在投放时段"
	case PromotionStatusIsAudit:
		return "新建审核中"
	case PromotionStatusIsReaudit:
		return "修改审核中"
	case PromotionStatusIsFrozen:
		return "已终止"
	case PromotionStatusIsAuditDeny:
		return "审核不通过"
	case PromotionStatusIsOfflineBudget:
		return "广告超出预算"
	case PromotionStatusIsOfflineBalance:
		return "账户余额不足"
	case PromotionStatusIsPreofflineBudget:
		return "广告接近预算"
	case PromotionStatusIsDisabled:
		return "已暂停"
	case PromotionStatusIsProjectDisabled:
		return "已被项目暂停"
	case PromotionStatusIsLiveRoomOff:
		return "关联直播间不可投"
	case PromotionStatusIsProductOffline:
		return "关联商品不可投"
	case PromotionStatusIsAwemeAccountDisabled:
		return "关联抖音账号不可投"
	case PromotionStatusIsAwemeAnchorDisabled:
		return "锚点不可投"
	case PromotionStatusIsDisableByQuota:
		return "已暂停（配额达限）"
	case PromotionStatusIsCreate:
		return "新建"
	case PromotionStatusIsAdvertiserOfflineBudget:
		return "账号超出预算"
	case PromotionStatusIsAdvertiserPreofflineBudget:
		return "账号接近预算"
	default:
		return ""
	}
}

// PromotionFirstStatus : 广告一级状态
type PromotionFirstStatus string

const (
	PromotionFirstStatusIsEnable  PromotionFirstStatus = "PROMOTION_STATUS_ENABLE"
	PromotionFirstStatusIsDisable PromotionFirstStatus = "PROMOTION_STATUS_DISABLE"
	PromotionFirstStatusIsFrozen  PromotionFirstStatus = "PROMOTION_STATUS_FROZEN"
	PromotionFirstStatusIsDone    PromotionFirstStatus = "PROMOTION_STATUS_DONE"
	PromotionFirstStatusIsDeleted PromotionFirstStatus = "PROMOTION_STATUS_DELETED"
)

func (t PromotionFirstStatus) String() string {
	switch t {
	case PromotionFirstStatusIsEnable:
		return "投放中"
	case PromotionFirstStatusIsDisable:
		return "未投放"
	case PromotionFirstStatusIsFrozen:
		return "已终止"
	case PromotionFirstStatusIsDone:
		return "已完成"
	case PromotionFirstStatusIsDeleted:
		return "已删除"
	default:
		return ""
	}
}

// PromotionSecondStatus : 广告二级状态
type PromotionSecondStatus string

const (
	PromotionSecondStatusIsAuditDeny              PromotionSecondStatus = "AUDIT_DENY"
	PromotionSecondStatusIsAudit                  PromotionSecondStatus = "AUDIT"
	PromotionSecondStatusIsReaudit                PromotionSecondStatus = "REAUDIT"
	PromotionSecondStatusIsDisableByGuota         PromotionSecondStatus = "DISABLE_BY_QUOTA"
	PromotionSecondStatusIsProjectDisabled        PromotionSecondStatus = "PROJECT_DISABLED"
	PromotionSecondStatusIsNoSchedule             PromotionSecondStatus = "NO_SCHEDULE"
	PromotionSecondStatusIsTimeNoReach            PromotionSecondStatus = "TIME_NO_REACH"
	PromotionSecondStatusIsOfflineBalance         PromotionSecondStatus = "OFFLINE_BALANCE"
	PromotionSecondStatusIsBalanceOfflineBudget   PromotionSecondStatus = "BALANCE_OFFLINE_BUDGET"
	PromotionSecondStatusIsProjectOfflineBudget   PromotionSecondStatus = "PROJECT_OFFLINE_BUDGET"
	PromotionSecondStatusIsPromotionOfflineBudget PromotionSecondStatus = "PROMOTION_OFFLINE_BUDGET"
	PromotionSecondStatusIsLiveRoomOff            PromotionSecondStatus = "LIVE_ROOM_OFF"
	PromotionSecondStatusIsProductOffline         PromotionSecondStatus = "PRODUCT_OFFLINE"
	PromotionSecondStatusIsAwemeAccountDisabled   PromotionSecondStatus = "AWEME_ACCOUNT_DISABLED"
	PromotionSecondStatusIsAwemeAnchorDisabled    PromotionSecondStatus = "AWEME_ANCHOR_DISABLED"
)

func (t PromotionSecondStatus) String() string {
	switch t {
	case PromotionSecondStatusIsAuditDeny:
		return "审核不通过"
	case PromotionSecondStatusIsAudit:
		return "新建审核中"
	case PromotionSecondStatusIsReaudit:
		return "修改审核中"
	case PromotionSecondStatusIsDisableByGuota:
		return "已暂停（配额达限）"
	case PromotionSecondStatusIsProjectDisabled:
		return "项目已被暂停"
	case PromotionSecondStatusIsNoSchedule:
		return "不在投放时段"
	case PromotionSecondStatusIsTimeNoReach:
		return "未到达投放时间"
	case PromotionSecondStatusIsOfflineBalance:
		return "未到达投放时间"
	case PromotionSecondStatusIsBalanceOfflineBudget:
		return "账户超出预算"
	case PromotionSecondStatusIsProjectOfflineBudget:
		return "项目超出预算"
	case PromotionSecondStatusIsPromotionOfflineBudget:
		return "广告超出预算"
	case PromotionSecondStatusIsLiveRoomOff:
		return "直播间不可投"
	case PromotionSecondStatusIsProductOffline:
		return "商品不可投"
	case PromotionSecondStatusIsAwemeAccountDisabled:
		return "抖音账号不可投"
	case PromotionSecondStatusIsAwemeAnchorDisabled:
		return "锚点不可投"
	default:
		return ""
	}
}
