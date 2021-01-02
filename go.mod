module github.com/WolvenKit/gpm

go 1.15

replace github.com/WolvenKit/gpm/internal/gpm/commands => ./internal/gpm/commands

replace github.com/WolvenKit/gpm/internal/gpm/game => ./internal/gpm/game

replace github.com/WolvenKit/gpm/internal/gpm/mod => ./internal/gpm/mod

require github.com/urfave/cli/v2 v2.3.0
