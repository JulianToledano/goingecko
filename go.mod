module github.com/JulianToledano/goingecko/v3

go 1.22.4

require golang.org/x/time v0.5.0

retract (
	v3.0.1 // Wrong Categories endpoint.
	v3.0.0 // Wrong Categories endpoint.
)
