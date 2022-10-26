package parser

const (
	CONTEXT_TABLE = iota
	CONTEXT_RELN
	CONTEXT_COMMENT
	CONTEXT_ATTR
)

type Context struct {
	Stack []int
}

func (c *Context) Push(CONTEXT_ENUM int) {
	c.Stack = append(c.Stack, CONTEXT_ENUM)
}

func (c *Context) Pop() {
	if len(c.Stack) > 0 {
		c.Stack = c.Stack[:len(c.Stack)-1]
	}
}

func (c *Context) GetCurrentContext() int {
	if len(c.Stack) > 0 {
		return c.Stack[len(c.Stack)-1]
	}
	return -1
}