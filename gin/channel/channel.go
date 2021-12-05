package channel

var chanLength int32

func init() {
	// adminMessage
	go func() {
		NewAdminMessageChannel()
	}()
}
