# Debuggers

    1) print/log
    2) GDB - cant understand go architecture completely. some issues
    3) delve

## Delve

    A) Installation:
    ----------------
        go get -u github.com/go-delve/delve/cmd/dlv

        If you are using Go 1.5 you must set GO15VENDOREXPERIMENT=1 before continuing.

    B) Commands:
    ------------
        args - - - - - - - - - - - - Print function arguments.
        break (alias: b) - - - - - - Sets a breakpoint.
        breakpoints (alias: bp) - - - Print out info for active breakpoints.
        call - - - - - - - - - - - - Resumes process, injecting a function call (EXPERIMENTAL!!!)
        clear - - - - - - - - - - - - Deletes breakpoint.
        clearall - - - - - - - - - - Deletes multiple breakpoints.
        condition (alias: cond) - - - Set breakpoint condition.
        config - - - - - - - - - - - Changes configuration parameters.
        continue (alias: c) - - - - - Run until a breakpoint or program termination.
        deferred - - - - - - - - - - Executes command in the context of a deferred call.
        disassemble (alias: disass) - Disassembler.
        down - - - - - - - - - - - - Move the current frame down.
        edit (alias: ed) - - - - - - Open where you are in $DELVE_EDITOR or $EDITOR
        exit (alias: quit | q) - - - Exit the debugger.
        frame - - - - - - - - - - - - Set the current frame, or execute command on a different frame.
        funcs - - - - - - - - - - - - Print list of functions.
        goroutine - - - - - - - - - - Shows or changes current goroutine
        goroutines - - - - - - - - - List program goroutines.
        help (alias: h) - - - - - - - Prints the help message.
        list (alias: ls | l) - - - - Show source code.
        locals - - - - - - - - - - - Print local variables.
        next (alias: n) - - - - - - - Step over to next source line.
        on - - - - - - - - - - - - - Executes a command when a breakpoint is hit.
        print (alias: p) - - - - - - Evaluate an expression.
        regs - - - - - - - - - - - - Print contents of CPU registers.
        restart (alias: r) - - - - - Restart process.
        set - - - - - - - - - - - - - Changes the value of a variable.
        source - - - - - - - - - - - Executes a file containing a list of delve commands
        sources - - - - - - - - - - - Print list of source files.
        stack (alias: bt) - - - - - - Print stack trace.
        step (alias: s) - - - - - - - Single step through program.
        step-instruction (alias: si) Single step a single cpu instruction.
        stepout - - - - - - - - - - - Step out of the current function.
        thread (alias: tr) - - - - - Switch to the specified thread.
        threads - - - - - - - - - - - Print out info for every traced thread.
        trace (alias: t) - - - - - - Set tracepoint.
        types - - - - - - - - - - - - Print list of types
        up - - - - - - - - - - - - - Move the current frame up.
        vars - - - - - - - - - - - - Print package variables.
        whatis - - - - - - - - - - - Prints type of an expression.
