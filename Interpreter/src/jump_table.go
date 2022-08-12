package main

type JumpTable [256]*operation

type operation struct {
	// execute is the operation function
	execute     executionFunc
	minStack int
	maxStack int

	jumps   bool // indicates whether the program counter should not increment
	returns bool // determines whether the operations sets the return data content
	halts bool
}

type (
	executionFunc func(pc *uint64, callContext *ScopeContext) ([]byte, error)
)

func minStack(pops, push int) int {
	return pops
}
func newInstructionSet() JumpTable {
	return JumpTable{
		PUSH1: {
			execute:     opPush1,
			minStack:    minStack(0, 1),
		},
		ADD: {
			execute:     opAdd,
			minStack:    minStack(2, 1),
		},
		SUB: {
			execute:     opSub,
			minStack:    minStack(2, 1),
		},
		MUL: {
			execute:     opMul,
			minStack:    minStack(2, 1),
		},
		QUO: {
			execute:     opQuo,
			minStack:    minStack(2, 1),
		},
		REM: {
			execute:     opRem,
			minStack:    minStack(2, 1),
		},
		AND: {
			execute:     opAnd,
			minStack:    minStack(2, 1),
		},
		OR: {
			execute:     opOr,
			minStack:    minStack(2, 1),
		},
		XOR: {
			execute:     opXor,
			minStack:    minStack(2, 1),
		},
		RETURN: {
			execute:    opReturn,
			minStack:   minStack(0, 0),
			halts:      true,
		},

	}
}