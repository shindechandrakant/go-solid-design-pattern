package rate_limiter

type TierType string

var (
	FREE    TierType = "FREE"
	PREMIUM TierType = "PREMIUM"
)

type User struct {
	UserId string
	Tier   TierType
}
