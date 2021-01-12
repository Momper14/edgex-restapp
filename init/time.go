package init

import (
	"os"
	"time"

	timezone "4d63.com/tz"
	"github.com/sirupsen/logrus"
)

func init() {
	// manually set time zone
	if tz := os.Getenv("TZ"); tz != "" {
		var err error
		time.Local, err = timezone.LoadLocation(tz)
		if err != nil {
			logrus.Warnf("error loading location '%s': %v", tz, err)
		}
	}
}
