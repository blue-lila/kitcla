package dat

const MessageSuccess = "success"
const MessageFailure = "failure"

type Message struct {
	Title string
	Body  string
	Kind  string
}
