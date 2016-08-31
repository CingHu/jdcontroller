package ofp

var XID chan uint32

func init() {
	XID = make(chan uint32)
	go func() {
		i := uint32(0)
		for {
			i += 1
			XID <- i
		}

	}()
}
