package enums

// MaterialStatus : 素材审核状态
type MaterialStatus string

const (
	MaterialStatusIsAdvOfflineBudget        MaterialStatus = "MATERIAL_STATUS_ADV_OFFLINE_BUDGET"
	MaterialStatusIsAdvPreOfflineBudget     MaterialStatus = "MATERIAL_STATUS_ADV_PRE_OFFLINE_BUDGET"
	MaterialStatusIsAudit                   MaterialStatus = "MATERIAL_STATUS_AUDIT"
	MaterialStatusIsAwemeAnchorDisabled     MaterialStatus = "MATERIAL_STATUS_AWEME_ANCHOR_DISABLED"
	MaterialStatusIsDelete                  MaterialStatus = "MATERIAL_STATUS_DELETE"
	MaterialStatusIsDisable                 MaterialStatus = "MATERIAL_STATUS_DISABLE"
	MaterialStatusIsFrozen                  MaterialStatus = "MATERIAL_STATUS_FROZEN"
	MaterialStatusIsLiveRoomOff             MaterialStatus = "MATERIAL_STATUS_LIVE_ROOM_OFF"
	MaterialStatusIsMaterialDelete          MaterialStatus = "MATERIAL_STATUS_MATERIAL_DELETE"
	MaterialStatusIsNoSchedule              MaterialStatus = "MATERIAL_STATUS_NO_SCHEDULE"
	MaterialStatusIsOfflineAudit            MaterialStatus = "MATERIAL_STATUS_OFFLINE_AUDIT"
	MaterialStatusIsOfflineBalance          MaterialStatus = "MATERIAL_STATUS_OFFLINE_BALANCE"
	MaterialStatusIsOfflineBudget           MaterialStatus = "MATERIAL_STATUS_OFFLINE_BUDGET"
	MaterialStatusIsOk                      MaterialStatus = "MATERIAL_STATUS_OK"
	MaterialStatusIsPreofflineBugdet        MaterialStatus = "MATERIAL_STATUS_PREOFFLINE_BUGDET"
	MaterialStatusIsProductOffline          MaterialStatus = "MATERIAL_STATUS_PRODUCT_OFFLINE"
	MaterialStatusIsProjectDisable          MaterialStatus = "MATERIAL_STATUS_PROJECT_DISABLE"
	MaterialStatusIsProjectOfflineBudget    MaterialStatus = "MATERIAL_STATUS_PROJECT_OFFLINE_BUDGET"
	MaterialStatusIsProjectPreofflineBudget MaterialStatus = "MATERIAL_STATUS_PROJECT_PREOFFLINE_BUDGET"
	MaterialStatusIsPromotionDisable        MaterialStatus = "MATERIAL_STATUS_PROMOTION_DISABLE"
	MaterialStatusIsPromotionQuotaLimit     MaterialStatus = "MATERIAL_STATUS_PROMOTION_QUOTA_LIMIT"
	MaterialStatusIsReaudit                 MaterialStatus = "MATERIAL_STATUS_REAUDIT"
	MaterialStatusIsDone                    MaterialStatus = "MATERIAL_STATUS_TIME_DONE"
	MaterialStatusIsTimeNoReach             MaterialStatus = "MATERIAL_STATUS_TIME_NO_REACH"
)

func (t MaterialStatus) String() string {
	switch t {
	case MaterialStatusIsAdvOfflineBudget:
		return "广告主超出预算"
	case MaterialStatusIsAdvPreOfflineBudget:
		return "广告主接近预算"
	case MaterialStatusIsAudit:
		return "新建审核中"
	case MaterialStatusIsAwemeAnchorDisabled:
		return "关联锚点不可投"
	case MaterialStatusIsDelete:
		return "已删除"
	case MaterialStatusIsDisable:
		return "已暂停"
	case MaterialStatusIsFrozen:
		return "已终止"
	case MaterialStatusIsLiveRoomOff:
		return "直播间关闭"
	case MaterialStatusIsMaterialDelete:
		return "素材已删除"
	case MaterialStatusIsNoSchedule:
		return "不在投放时段"
	case MaterialStatusIsOfflineAudit:
		return "审核不通过"
	case MaterialStatusIsOfflineBalance:
		return "账户余额不足"
	case MaterialStatusIsOfflineBudget:
		return "广告超出预算"
	case MaterialStatusIsOk:
		return "投放中"
	case MaterialStatusIsPreofflineBugdet:
		return "广告接近预算"
	case MaterialStatusIsProductOffline:
		return "关联商品不可投"
	case MaterialStatusIsProjectDisable:
		return "已被项目暂停"
	case MaterialStatusIsProjectOfflineBudget:
		return "项目超出预算"
	case MaterialStatusIsProjectPreofflineBudget:
		return "项目接近预算"
	case MaterialStatusIsPromotionDisable:
		return "已被广告暂停"
	case MaterialStatusIsPromotionQuotaLimit:
		return "因为quota限额暂停"
	case MaterialStatusIsReaudit:
		return "修改审核中"
	case MaterialStatusIsDone:
		return "已完成"
	case MaterialStatusIsTimeNoReach:
		return "未到达投放时间"
	default:
		return ""
	}
}
