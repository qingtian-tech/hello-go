package hello

import(
"time"
)

func Say(str string) string {
	now := time.Now().String()
	return str + string(now);
}