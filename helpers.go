package main

func push(iarg int) IArg {
	return IArg{OP_PUSH, iarg}
}

func plus() IArg {
	return IArg{OP_PLUS, 0}
}

func minus() IArg {
	return IArg{OP_MINUS, 0}
}

func dump() IArg {
	return IArg{OP_DUMP, 0}
}
