module github.com/WolvenKit/gpm/internal/gpm/mod

go 1.15

require (
	github.com/spf13/viper v1.7.1
	go.uber.org/zap v1.10.0
)

replace github.com/WolvenKit/gpm/internal/gpm/game => ../game
