package enums

// InventoryCatalog : 广告位大类
type InventoryCatalog string

const (
	ManualInventoryCatalog         InventoryCatalog = "MANUAL"
	UniversalSmartInventoryCatalog InventoryCatalog = "UNIVERSAL_SMART"
)

func (t InventoryCatalog) String() string {
	switch t {
	case ManualInventoryCatalog:
		return "首选媒体"
	case UniversalSmartInventoryCatalog:
		return "通投智选"
	default:
		return ""
	}
}
