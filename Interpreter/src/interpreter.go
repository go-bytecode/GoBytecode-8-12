package main

type OpCode byte

type ScopeContext struct {
	Memory   *Memory
	Stack    *Stack
    code []byte
}



func Run(code []byte) (ret []byte, err error) {
	var (
		JT          = newInstructionSet()
		op          OpCode        // current opcode
		mem         = NewMemory() // bound memory
		stack       = newstack()  // local stack
		callContext = &ScopeContext{
			Memory: mem,
			Stack:  stack,
			code:   code,
		}
		pc = uint64(0) // program counter
		_  []byte      // result of the opcode execution function
	)



	for {
		op = OpCode(code[pc])
		operation := JT[op]
		if operation == nil {
			return nil, err
		}
		if sLen := stack.len(); sLen < operation.minStack {
			return nil, err
		}

		ret, err = operation.execute(&pc, callContext)

		switch {
		case err != nil:
			return nil, err
		case operation.halts:
			return ret, nil
		case !operation.jumps:
			pc++
		}
	}

}