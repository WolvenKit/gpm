module github.com/WolvenKit/gpm/internal/gpm/mod

go 1.15

require (
	github.com/WolvenKit/gpm/internal/gpm/game v0.0.0-00010101000000-000000000000
	github.com/dsnet/compress v0.0.1 // indirect
	github.com/frankban/quicktest v1.11.3 // indirect
	github.com/golang/snappy v0.0.2 // indirect
	github.com/mholt/archiver v3.1.1+incompatible
	github.com/nwaples/rardecode v1.1.0 // indirect
	github.com/pierrec/lz4 v2.6.0+incompatible // indirect
	github.com/spf13/viper v1.7.1
	github.com/ulikunitz/xz v0.5.9 // indirect
	github.com/xi2/xz v0.0.0-20171230120015-48954b6210f8 // indirect
	go.uber.org/zap v1.10.0
)

replace github.com/WolvenKit/gpm/internal/gpm/game => ../game

replace github.com/WolvenKit/gpm/cmd => ../../../cmd
