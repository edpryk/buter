package docs

import (
	"fmt"
	"strings"
)

var (
	attackTypeFlag = "a"
	urlFlag        = "u"
	payloadFlag    = "p"
	threadsFlag    = "t"
	headersFlag    = "h"
	delayFlag      = "d"
	methodFlag     = "m"

	urlUsage        = "-u <http://localhost?param1=!abc!&param_N=!efg!> (payloader wrapped into '!' char)"
	payloadUsage    = "-p <payloader-file_1> -p <payloader-file_N>"
	attackTypeUsage = fmt.Sprintf("-a <%s>", strings.Join([]string{ClusterAttack}, "/"))
	threadsUseage   = "-t 5"
	headersUsage    = "'{ \"User-Agent\": \"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/99.1\" }'"
	delayUsage      = "-d 800 (In Milisesonds)"
	methodUsage     = "-m get"

	defaultAttackType = ""
	defaultUrl        = ""
	defaultThreads    = 3
	defaultHeaders    = ""
	defaultDealy      = 800
	defaultMethod     = "get"
)