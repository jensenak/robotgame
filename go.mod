module github.com/jensenak/robotgame/v2

go 1.14

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/hashicorp/go-uuid v1.0.2
	github.com/jensenak/robotgame/qwirk v0.0.0-00010101000000-000000000000
	github.com/jensenak/robotgame/store v0.0.0-00010101000000-000000000000
)

replace github.com/jensenak/robotgame/qwirk => ./qwirk

replace github.com/jensenak/robotgame/store => ./store
