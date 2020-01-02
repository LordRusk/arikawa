package discord

import (
	"strconv"
	"time"
)

const DiscordEpoch = 1420070400000 * int64(time.Millisecond)

type Snowflake uint64

func NewSnowflake(t time.Time) Snowflake {
	return Snowflake(TimeToDiscordEpoch(t) << 22)
}

func (s Snowflake) String() string {
	return strconv.FormatUint(uint64(s), 10)
}

func (s Snowflake) Time() time.Time {
	return time.Unix(0, int64(s)>>22*1000000+DiscordEpoch)
}

func (s Snowflake) Worker() uint8 {
	return uint8(s & 0x3E0000)
}

func (s Snowflake) PID() uint8 {
	return uint8(s & 0x1F000 >> 12)
}

func (s Snowflake) Increment() uint16 {
	return uint16(s & 0xFFF)
}

func TimeToDiscordEpoch(t time.Time) int64 {
	return t.UnixNano()/int64(time.Millisecond) - DiscordEpoch
}
