module example

go 1.19

//replace github.com/RickyGrassmuck/lichess-go/lichess => ../lichess

require github.com/RickyGrassmuck/lichess-go/lichess v0.0.0-20220816014250-0efe1011f2ad

require (
	github.com/deepmap/oapi-codegen v1.11.0 // indirect
	github.com/google/uuid v1.3.0 // indirect
)
