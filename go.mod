module github.com/WolvenKit/gpm

go 1.15

require (
	github.com/WolvenKit/gpm/internal/gpm/commands v0.0.0
	github.com/WolvenKit/gpm/internal/gpm/game v0.0.0-00010101000000-000000000000 // indirect
	github.com/mholt/archiver v3.1.1+incompatible
	github.com/spf13/cobra v1.1.1
	github.com/stretchr/testify v1.6.1
	github.com/urfave/cli/v2 v2.3.0
)

replace github.com/WolvenKit/gpm/internal/gpm/commands => ./internal/gpm/commands

replace github.com/WolvenKit/gpm/internal/gpm/game => ./internal/gpm/game

replace github.com/WolvenKit/gpm/internal/gpm/mod => ./internal/gpm/mod
