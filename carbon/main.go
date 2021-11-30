package carbon

import (
	"fmt"
	"github.com/golang-module/carbon"
)

func main() {
	// Return datetime of today
	fmt.Sprintf("%s", carbon.Now()) // 2020-08-05 13:14:15
	carbon.Now().ToDateTimeString() // 2020-08-05 13:14:15
	// Return date of today
	carbon.Now().ToDateString() // 2020-08-05
	// Return time of today
	carbon.Now().ToTimeString() // 13:14:15
	// Return datetime of today in a given timezone
	carbon.Now(carbon.NewYork).ToDateTimeString() // 2020-08-05 14:14:15
	carbon.SetTimezone(carbon.NewYork).Now().ToDateTimeString() // 2020-08-05 14:14:15
	// Return timestamp with second of today
	carbon.Now().Timestamp() // 1596604455
	carbon.Now().TimestampWithSecond() // 1596604455
	// Return timestamp with millisecond of today
	carbon.Now().TimestampWithMillisecond() // 1596604455000
	// Return timestamp with microsecond of today
	carbon.Now().TimestampWithMicrosecond() // 1596604455000000
	// Return timestamp with nanosecond of today
	carbon.Now().TimestampWithNanosecond() // 1596604455000000000

	// Return datetime of yesterday
	fmt.Sprintf("%s", carbon.Yesterday()) // 2020-08-04 13:14:15
	carbon.Yesterday().ToDateTimeString() // 2020-08-04 13:14:15
	// Return date of yesterday
	carbon.Yesterday().ToDateString() // 2020-08-04
	// Return time of yesterday
	carbon.Yesterday().ToTimeString() // 13:14:15
	// Return datetime of yesterday on a given day
	carbon.Parse("2021-01-28 13:14:15").Yesterday().ToDateTimeString() // 2021-01-27 13:14:15
	// Return datetime of yesterday in a given timezone
	carbon.Yesterday(carbon.NewYork).ToDateTimeString() // 2020-08-04 14:14:15
	carbon.SetTimezone(carbon.NewYork).Yesterday().ToDateTimeString() // 2020-08-04 14:14:15
	// Return timestamp with second of yesterday
	carbon.Yesterday().Timestamp() // 1596518055
	carbon.Yesterday().TimestampWithSecond() // 1596518055
	// Return timestamp with millisecond of yesterday
	carbon.Yesterday().TimestampWithMillisecond() // 1596518055000
	// Return timestamp with microsecond of yesterday
	carbon.Yesterday().TimestampWithMicrosecond() // 1596518055000000
	// Return timestamp with nanosecond of yesterday
	carbon.Yesterday().TimestampWithNanosecond() // 1596518055000000000

	// Return datetime of tomorrow
	fmt.Sprintf("%s", carbon.Tomorrow()) // 2020-08-06 13:14:15
	carbon.Tomorrow().ToDateTimeString() // 2020-08-06 13:14:15
	// Return date of tomorrow
	carbon.Tomorrow().ToDateString() // 2020-08-06
	// Return time of tomorrow
	carbon.Tomorrow().ToTimeString() // 13:14:15
	// Return datetime of tomorrow on a given day
	carbon.Parse("2021-01-28 13:14:15").Tomorrow().ToDateTimeString() // 2021-01-29 13:14:15
	// Return datetime of tomorrow in a given timezone
	carbon.Tomorrow(carbon.NewYork).ToDateTimeString() // 2020-08-06 14:14:15
	carbon.SetTimezone(carbon.NewYork).Tomorrow().ToDateTimeString() // 2020-08-06 14:14:15
	// Return timestamp with second of tomorrow
	carbon.Tomorrow().Timestamp() // 1596690855
	carbon.Tomorrow().TimestampWithSecond() // 1596690855
	// Return timestamp with millisecond of tomorrow
	carbon.Tomorrow().TimestampWithMillisecond() // 1596690855000
	// Return timestamp with microsecond of tomorrow
	carbon.Tomorrow().TimestampWithMicrosecond() // 1596690855000000
	// Return timestamp with nanosecond of tomorrow
	carbon.Tomorrow().TimestampWithNanosecond() // 1596690855000000000

	// Create a Carbon instance from a given timestamp with second
	carbon.CreateFromTimestamp(-1).ToDateTimeString() // 1970-01-01 07:59:59
	carbon.CreateFromTimestamp(-1, carbon.Tokyo).ToDateTimeString() // 1970-01-01 08:59:59
	carbon.CreateFromTimestamp(0).ToDateTimeString() // 1970-01-01 08:00:00
	carbon.CreateFromTimestamp(0, carbon.Tokyo).ToDateTimeString() // 1970-01-01 09:00:00
	carbon.CreateFromTimestamp(1596604455).ToDateTimeString() // 2020-08-05 13:14:15
	carbon.CreateFromTimestamp(1596604455, carbon.Tokyo).ToDateTimeString() // 2020-08-05 14:14:15
	// Create a Carbon instance from a given timestamp with millisecond
	carbon.CreateFromTimestamp(1596604455000).ToDateTimeString() // 2020-08-05 13:14:15
	carbon.CreateFromTimestamp(1596604455000, carbon.Tokyo).ToDateTimeString() // 2020-08-05 14:14:15
	// Create a Carbon instance from a given timestamp with microsecond
	carbon.CreateFromTimestamp(1596604455000000).ToDateTimeString() // 2020-08-05 13:14:15
	carbon.CreateFromTimestamp(1596604455000000, carbon.Tokyo).ToDateTimeString() // 2020-08-05 14:14:15
	// Create a Carbon instance from a given timestamp with nanosecond
	carbon.CreateFromTimestamp(1596604455000000000).ToDateTimeString() // 2020-08-05 13:14:15
	carbon.CreateFromTimestamp(1596604455000000000, carbon.Tokyo).ToDateTimeString() // 2020-08-05 14:14:15

	// Create a Carbon instance from a given year, month, day, hour, minute and second
	carbon.CreateFromDateTime(2020, 8, 5, 13, 14, 15).ToDateTimeString() // 2020-08-05 13:14:15
	carbon.CreateFromDateTime(2020, 8, 5, 13, 14, 15, carbon.Tokyo).ToDateTimeString() // 2020-08-05 14:14:15
	// Create a Carbon instance from a given year, month and day
	carbon.CreateFromDate(2020, 8, 5).ToDateTimeString() // 2020-08-05 13:14:15
	carbon.CreateFromDate(2020, 8, 5, carbon.Tokyo).ToDateTimeString() // 2020-08-05 14:14:15
	// Create a Carbon instance from a given hour, minute and second
	carbon.CreateFromTime(13, 14, 15).ToDateTimeString() // 2020-08-05 13:14:15
	carbon.CreateFromTime(13, 14, 15, carbon.Tokyo).ToDateTimeString() // 2020-08-05 14:14:15

}