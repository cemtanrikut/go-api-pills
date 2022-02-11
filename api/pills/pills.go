package pills

import "time"

type pill struct {
	name       string
	amount     string
	reminder   bool
	periodicly bool
	active     bool
	start_date time.Time
	end_date   time.Time
}
