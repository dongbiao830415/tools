// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Goyacc is a version of yacc for Go.
It is written in Go and generates parsers written in Go.
Goyacc是yacc的go版本。
它用Go编写，并生成用Go编写的解析器。
Usage:

	goyacc args...

It is largely transliterated from the Inferno version written in Limbo
which in turn was largely transliterated from the Plan 9 version
written in C and documented at
它很大程度上是从Limbo编写的Inferno版本转译而来的，而Inferno版本又是从C语言编写的Plan 9版本转译而来，在
	https://9p.io/magic/man2html/1/yacc

Adepts of the original yacc will have no trouble adapting to this
form of the tool.
原始yacc的专家将可以轻松适应这种形式的工具。

The directory $GOPATH/src/golang.org/x/tools/cmd/goyacc/testdata/expr
is a yacc program for a very simple expression parser. See expr.y and
main.go in that directory for examples of how to write and build
goyacc programs.
$ GOPATH / src / golang.org / x / tools / cmd / goyacc / testdata / expr目录是用于非常简单的表达式解析器的yacc程序。 
有关如何编写和构建goyacc程序的示例，请参见该目录中的expr.y和main.go。

The generated parser is reentrant. The parsing function yyParse expects
to be given an argument that conforms to the following interface:
生成的解析器是可重入的。 解析函数yyParse期望得到一个符合以下接口的参数：

	type yyLexer interface {
		Lex(lval *yySymType) int
		Error(e string)
	}

Lex should return the token identifier, and place other token
information in lval (which replaces the usual yylval).
Error is equivalent to yyerror in the original yacc.
Lex应该返回令牌标识符，并将其他令牌信息放入lval（代替通常的yylval）。 
错误等同于原始yacc中的yyerror。

Code inside the grammar actions may refer to the variable yylex,
which holds the yyLexer passed to yyParse.
语法操作中的代码可能引用变量yylex，该变量保存传递给yyParse的yyLexer。

Clients that need to understand more about the parser state can
create the parser separately from invoking it. The function yyNewParser
returns a yyParser conforming to the following interface:

需要了解有关解析器状态的更多信息的客户端可以与调用它分开创建解析器。
函数yyNewParser返回符合以下接口的yyParser：

	type yyParser interface {
		Parse(yyLex) int
		Lookahead() int
	}

Parse runs the parser; the top-level call yyParse(yylex) is equivalent
to yyNewParser().Parse(yylex).
解析运行解析器； 顶级调用yyParse（yylex）等效于yyNewParser（）。Parse（yylex）。

Lookahead can be called during grammar actions to read (but not consume)
the value of the current lookahead token, as returned by yylex.Lex.
If there is no current lookahead token (because the parser has not called Lex
or has consumed the token returned by the most recent call to Lex),
Lookahead returns -1. Calling Lookahead is equivalent to reading
yychar from within in a grammar action.
可以在语法操作期间调用前瞻，以读取（但不消耗）yalex.Lex返回的当前前瞻标记的值。 
如果没有当前的前瞻令牌（因为解析器未调用Lex或已消耗了对Lex的最近调用返回的令牌），
则Lookahead返回-1。 调用Lookahead等同于从语法操作中读取yychar。


Multiple grammars compiled into a single program should be placed in
distinct packages.  If that is impossible, the "-p prefix" flag to
goyacc sets the prefix, by default yy, that begins the names of
symbols, including types, the parser, and the lexer, generated and
referenced by yacc's generated code.  Setting it to distinct values
allows multiple grammars to be placed in a single package.
编译成一个程序的多个语法应放在不同的程序包中。 如果不可能，则goyacc的“ -p前缀”标志设置前缀（默认为yy），
该前缀以yacc的生成代码生成和引用的符号名称开头，包括类型，解析器和词法分析器。 
将其设置为不同的值可以将多个语法放在一个包中。
*/
package main
