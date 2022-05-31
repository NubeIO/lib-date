package datelib

type Admin struct {
	CMD *Command
}

func New(admin *Admin) *Admin {
	opts := &Command{}
	admin.CMD = opts
	return admin
}

func (inst *Admin) Uptime() (res *Response) {
	cmd := "uptime"
	inst.CMD.Commands = Builder(cmd)
	res = inst.CMD.RunCommand()
	return
}
