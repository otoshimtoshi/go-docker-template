package seeds

import (
	"time"

	"github.com/jinzhu/gorm"
)

func All() []Seed {
	return []Seed{
		{
			Name: "CreateJane",
			Run: func(db *gorm.DB) error {
				return CreateUser(db, "Jane", "pass", "jane@example.com", time.Now(), "https://otot.dev/ogp.png")
			},
		},
		{
			Name: "CreateJohn",
			Run: func(db *gorm.DB) error {
				return CreateUser(db, "John", "pass", "john@example.com", time.Now(), "https://otot.dev/ogp.png")
			},
		},
	}
}
