module github.com/serargz/ptxt

go 1.13

require (
	github.com/mitchellh/go-homedir v1.1.0
	github.com/serargz/ptxt/calendar v0.0.0
	github.com/spf13/cobra v0.0.6
	github.com/spf13/viper v1.6.2
)

replace github.com/serargz/ptxt/calendar v0.0.0 => ./calendar/
