module github.com/jensenak/robotgame/store/v2

go 1.14

require (
	github.com/jensenak/robotgame/qwirk v0.0.0-00010101000000-000000000000
	github.com/jinzhu/gorm v1.9.16
	github.com/lib/pq v1.8.0
	github.com/stretchr/testify v1.6.1
)

replace github.com/jensenak/robotgame/qwirk => ../qwirk
