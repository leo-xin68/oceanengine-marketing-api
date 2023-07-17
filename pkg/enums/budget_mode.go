package enums

// BudgetMode : 预算类型
type BudgetMode string

const (
	BudgetModeIsDay   BudgetMode = "BUDGET_MODE_DAY"
	BudgetModeIsTotal BudgetMode = "BUDGET_MODE_TOTAL"
)

func (t BudgetMode) String() string {
	switch t {
	case BudgetModeIsDay:
		return "日预算"
	case BudgetModeIsTotal:
		return "总预算"
	default:
		return ""
	}
}
