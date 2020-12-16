module github.com/WolvenKit/gpm

go 1.15

require (
	github.com/WolvenKit/gpm/internal/gpm/cli v0.1.0
	github.com/gruntwork-io/gruntwork-cli v0.7.1-0.20200831164626-978768fef544 // https://github.com/gruntwork-io/gruntwork-cli/pull/33
	github.com/stretchr/testify v1.4.0
)

replace github.com/WolvenKit/gpm/internal/gpm/cli => ./internal/gpm/cli
