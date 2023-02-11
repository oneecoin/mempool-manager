package lib

func HandleErr(err error) {
	if err != nil {
		panic(err)
	}
}

func HasErr(err error) bool {
	return err != nil
}
