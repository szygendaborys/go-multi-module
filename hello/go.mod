module example.com/hello

go 1.23.4

require (
	example.com/initials v0.0.0-00010101000000-000000000000
	golang.org/x/example/hello v0.0.0-20241216154601-40afcb705d05
)

replace example.com/initials => ../initials
