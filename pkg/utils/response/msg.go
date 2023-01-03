package response

// JSONMessage is a struct to generate json response
type JSONMessage struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
