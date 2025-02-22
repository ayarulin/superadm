package clock

import "time"

var now *time.Time

func Now() time.Time{

  if now == nil {
    return time.Now()
  }

  return *now
}

func Freeze(t time.Time) {
  now = &t
}
