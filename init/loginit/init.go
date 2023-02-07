package loginit

import "log"

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
}