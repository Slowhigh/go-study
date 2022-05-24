// Concrete command
package command

type onCommand struct {
	device device
}

func (c *onCommand) execute() {
	c.device.on()
}