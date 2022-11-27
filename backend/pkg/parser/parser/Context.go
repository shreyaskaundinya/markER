package parser

const (
	CONTEXT_TABLE         = iota // 0
	CONTEXT_RELN                 // 1
	CONTEXT_RELN_TABLE1          // 2
	CONTEXT_RELN_TABLE2          // 3
	CONTEXT_COMMENT              // 4
	CONTEXT_ATTR                 // 5
	CONTEXT_COMP_ATTR            // 6
	CONTEXT_COMP_SUB_ATTR        // 7
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