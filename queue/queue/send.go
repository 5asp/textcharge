package queue

type SendQueue struct {
	AppID   int    `json:"appID"`
	To      string `json:"to"`
	Msg     string `json:"msg"`
	Action  int    `json:"action"`
	Created int    `json:"created"`
}
