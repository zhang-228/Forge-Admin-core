package config

type Config struct {
	//StartTime time.Time `json:",omitempty"`
	User struct {
		DefaultPassword string `json:",default=forge-admin"`
		DefaultStatus   int64  `json:",default=1"`
	}
}
