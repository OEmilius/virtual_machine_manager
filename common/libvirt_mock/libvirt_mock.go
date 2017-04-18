package libvirt

//_ "github.com/rgbkrk/libvirt-go"

type Mashine struct {
	ID     string
	State  string
	Params map[string]string
}

func NewMashine() string {
	return "new_id"
}

func StartMashine() string {
	return "start virt mashine"
}

func StopMashine() string {
	return "stop virt mashine"
}

func DelMashine() string {
	return "delete virt mashine"
}
