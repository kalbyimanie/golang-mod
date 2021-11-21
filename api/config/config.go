package config

// Configurations exported
type Configurations struct {
	Server       ServerConfigurations
	Database     DatabaseConfigurations
	EXAMPLE_PATH string
	EXAMPLE_VAR  string
}

// ServerConfigurations exported
type ServerConfigurations struct {
	Ports int
}

// DatabaseConfigurations exported
type DatabaseConfigurations struct {
	DBName     string
	DBUser     string
	DBPassword string
}
