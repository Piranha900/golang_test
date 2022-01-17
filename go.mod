module golang_test

go 1.17

replace golang_test/model => ./model

require (
	golang_test/model v0.0.0-00010101000000-000000000000
	golang_test/zipFunctions v0.0.0-00010101000000-000000000000
)

require golang_test/genericFunctions v0.0.0-00010101000000-000000000000

replace golang_test/genericFunctions => ./genericFunctions

replace golang_test/zipFunctions => ./zipFunctions
