package error

//This struct enable the devs to customize the error
type CustomError struct {
	ErrCustom string
}

//This function will return the custom error as string
func(c *CustomError) Error() string {
	return c.ErrCustom
}