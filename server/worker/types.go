package worker

//worker actively reporting information to master for reducing master pressure.
type Reporter struct {
	Name string
}

type StartConfig struct {
	WorkerId       int64
	HealthCheckSec uint16
	MasterHttps    []string
	MasterGRpcs    []string
}
