package channel

func init(){
	// adminMessage
	go func() {
		NewAdminMessageChannel()
	}()
}