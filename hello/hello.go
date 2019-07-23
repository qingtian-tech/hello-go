package hello

import(
"time"
)

func Say(str string) string {
	now := time.Now().String()
	return "v2:" + str + string(now);
}