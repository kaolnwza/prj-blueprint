package schemes

type Wallet struct {
	WalletId string  `gorm:"column:wallet_id; primary_key; type:varchar(36); default:gen_random_uuid()"`
	UserId   string  `gorm:"column:user_id; type:varchar(36); not null; index"`
	Amount   float64 `gorm:"column:amount; type:float; not null; default:0"`
	User     User    `gorm:"references:UserId"`
}
