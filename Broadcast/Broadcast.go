package broadcast

type Broadcast struct {
	Discord    bool
	RocketChat bool
	Slack      bool
}

var myBroadcast = Broadcast{}
