package enums

// InventoryCatalog : 广告位大类
type InventoryCatalog string

const (
	InventoryCatalogIsManual         InventoryCatalog = "MANUAL"
	InventoryCatalogIsUniversalSmart InventoryCatalog = "UNIVERSAL_SMART"
)

func (t InventoryCatalog) String() string {
	switch t {
	case InventoryCatalogIsManual:
		return "首选媒体"
	case InventoryCatalogIsUniversalSmart:
		return "通投智选"
	default:
		return ""
	}
}
