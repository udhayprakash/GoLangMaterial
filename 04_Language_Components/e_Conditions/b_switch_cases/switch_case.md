# Switch Statement

    - switch statement is a multiway branch statement.
    - It provides an efficient way to transfer the execution
      of different parts of a code based on the value (also
      called case) of the expression.

    - Go language supports two types of switch statements:

          1. Expression Switch
          2. Type Switch

    Expression Switch:
        - similar to switch statement in c, c++, Java languages.
        - easy way to disptch execution to different parts of code based on value of the expression.

        switch optional_statement; optional_expression {
            case expression_1: Statement ..
            case expression_2: Statement ..
            ...

            default: Statement ..
        }

        - Both "optional_statement" and "optional_expression" are optional, separated by semicolon(;).
        - If switch does not contain any expression, then the compiler assume that expression is true.
        - optional_statement may contain simple statements like variable declarations, increment or assignment statements, or function calls, etc.
        - If a variable is present in optional statement, then the scope of the variable is limited to othat switch statement.
        - In switch statement, the case and default statement does not contain any brek statement But you areallowed to use break and fallthrough statement if your program required.
        - default statement is optional in switch statement.
        - If a case contains multiple values, separate them by a comma(,).

Ref: GeekforGeeks
