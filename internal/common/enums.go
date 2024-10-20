package common

import "github.com/golang-jwt/jwt/v5"

const (
	TopicCreateCompany = "createCompany"
	TopicUpdateCompany = "updateCompany"
	TopicDeleteCompany = "deleteCompany"
)

var validTypes = map[string]bool{
	"Corporations":        true,
	"NonProfit":           true,
	"Cooperative":         true,
	"Sole Proprietorship": true,
}

func GetValidCompanyTypes() map[string]bool {
	return validTypes
}

type JwtPayload struct {
	Email         string          `json:"email,omitempty"`
	PermissionMap map[string]bool `json:"permissions"`
	jwt.RegisteredClaims
}
