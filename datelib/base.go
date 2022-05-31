package datelib

type Date struct {
	CMD *Command
}

func New(admin *Date) *Date {
	opts := &Command{}
	admin.CMD = opts
	return admin
}

func (inst *Date) Uptime() (res *Response) {
	cmd := "uptime"
	inst.CMD.Commands = Builder(cmd)
	res = inst.CMD.RunCommand()
	return
}
