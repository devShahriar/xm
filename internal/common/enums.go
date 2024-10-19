package common

var validTypes = map[string]bool{
	"Corporations":        true,
	"NonProfit":           true,
	"Cooperative":         true,
	"Sole Proprietorship": true,
}

func GetValidCompanyTypes() map[string]bool {
	return validTypes
}
