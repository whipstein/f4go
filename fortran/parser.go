package fortran

import (
	"fmt"
	goast "go/ast"
	"go/token"
	"strconv"
)

type node struct {
	tok token.Token
	lit string
}

type parser struct {
	sc    *Scanner
	ast   goast.File
	ident int
	ns    []node

	functionExternalName []string
	initVars             []initialVar

	errs []error
}

type initialVar struct {
	name string
	typ  string
}

func (p *parser) init() {
	p.functionExternalName = make([]string, 0)
}

func (p *parser) prepare() (err error) {
	var last token.Token
	for {
		tok, lit := p.sc.Scan()
		if tok == token.EOF {
			break
		}

		switch tok {
		case token.COMMENT:
			continue
		}

		if last == NEW_LINE && tok == NEW_LINE {
			continue
		}

		p.ns = append(p.ns, node{
			tok: tok,
			lit: lit,
		})
		last = tok
	}

	if len(p.ns) > 0 && p.ns[0].tok == NEW_LINE {
		p.ns = p.ns[1:]
	}

	return
}

func a(ns []node) (out string) {
	for _, n := range ns {
		switch n.tok {
		case NEW_LINE:
			out += fmt.Sprintf("\n")
		default:
			out += fmt.Sprintf(" %v", n.lit)
		}
	}
	return
}

func (p *parser) parse() (err error) {
	err = p.prepare()
	if err != nil {
		return
	}

	p.ast.Name = goast.NewIdent("main")

	var decls []goast.Decl
	p.ident = 0
	decls = p.transpileToNode()
	if len(p.errs) > 0 {
		for _, e := range p.errs {
			err = fmt.Errorf("%v\n%v", err, e)
		}
		fmt.Println("Errors:\n ", err)
		err = nil
		// return
	}

	p.ast.Decls = append(p.ast.Decls, decls...)
	return
}

func (p *parser) transpileToNode() (decls []goast.Decl) {

	if p.ident < 0 || p.ident >= len(p.ns) {
		p.errs = append(p.errs,
			fmt.Errorf("Ident is outside nodes: %d/%d", p.ident, len(p.ns)))
		return
	}

	switch p.ns[p.ident].tok {
	case SUBROUTINE:
		var decl goast.Decl
		decl = p.transpileSubroutine()
		decls = append(decls, decl)

	default:
		// move to next NEW_LINE
		p.addError("Cannot parse line: " + p.getLine())
	}

	return
}

func (p *parser) gotoEndLine() {
	_ = p.getLine()
}

func (p *parser) getLine() (line string) {
	if p.ident < 0 || p.ident >= len(p.ns) {
		p.addError("Cannot get line, ident = " + strconv.Itoa(p.ident))
		return
	}
	for ; p.ident < len(p.ns); p.ident++ {
		if p.ns[p.ident].tok == NEW_LINE {
			break
		}
		line += fmt.Sprintf(" %s", p.ns[p.ident].lit)
	}
	return
}

func (p *parser) transpileSubroutine() (decl goast.Decl) {
	p.init()

	var fd goast.FuncDecl
	fd.Type = &goast.FuncType{
		Params: &goast.FieldList{},
	}

	p.expect(SUBROUTINE)

	p.ident++
	p.expect(token.IDENT)
	name := p.ns[p.ident].lit

	p.ident++
	var hasParens bool = p.ns[p.ident].tok == token.LPAREN
	if hasParens {
		p.expect(token.LPAREN)

		// Parameters
		p.ident++
		fd.Type.Params.List = p.parseParamDecl()

		p.ident++
		p.expect(token.RPAREN)

		p.ident++
		p.expect(NEW_LINE)
	}

	p.ident++
	fd.Name = goast.NewIdent(name)
	fd.Body = &goast.BlockStmt{
		Lbrace: 1,
		List:   p.transpileListStmt(),
	}
	p.ident++

	// delete external function type definition
checkExternalFunction:
	for i := range p.initVars {
		var remove bool
		for _, f := range p.functionExternalName {
			if p.initVars[i].name == f {
				remove = true
				break
			}
		}
		if remove {
			// fmt.Println("Remove external function definition: ", name)
			p.initVars = append(p.initVars[:i], p.initVars[i+1:]...)
			goto checkExternalFunction
		}
	}

	// add correct type of subroutine arguments
checkArguments:
	for i := range fd.Type.Params.List {
		fieldName := fd.Type.Params.List[i].Names[0].Name
		for j := range p.initVars {
			if fieldName == p.initVars[j].name {
				fd.Type.Params.List[i].Type = goast.NewIdent(p.initVars[j].typ)

				// fmt.Println("Remove to arg : ", fieldName)
				p.initVars = append(p.initVars[:j], p.initVars[j+1:]...)
				goto checkArguments
			}
		}
	}

	// init vars
	var vars []goast.Stmt
	for i := range p.initVars {
		vars = append(vars, &goast.DeclStmt{
			Decl: &goast.GenDecl{
				Tok: token.VAR,
				Specs: []goast.Spec{
					&goast.ValueSpec{
						Names: []*goast.Ident{goast.NewIdent(p.initVars[i].name)},
						Type:  goast.NewIdent(p.initVars[i].typ),
					},
				},
			},
		})
	}

	fd.Body.List = append(vars, fd.Body.List...)

	decl = &fd
	return
}

func (p *parser) addError(msg string) {
	p.errs = append(p.errs, fmt.Errorf("%s", msg))
}

func (p *parser) expect(t token.Token) {
	if t != p.ns[p.ident].tok {
		panic(fmt.Errorf("Expect %s, but we have {{%s,%s}}",
			view(t), view(p.ns[p.ident].tok), p.ns[p.ident].lit))
	}
}

func (p *parser) transpileListStmt() (stmts []goast.Stmt) {
	for p.ident < len(p.ns) {
		if p.ns[p.ident].tok == END || p.ns[p.ident].tok == token.ELSE {
			// TODO
			break
		}
		stmt := p.parseStmt()
		if stmt == nil {
			// p.addError("stmt is nil in line : " + p.getLine())
			// break
			continue
		}
		stmts = append(stmts, stmt...)
	}
	return
}

func (p *parser) parseInit() (stmts []goast.Stmt) {

	identType := "int"
	switch p.ns[p.ident].tok {
	case LOGICAL:
		identType = "bool"
	case CHARACTER:
		identType = "byte"
	case COMPLEX:
		identType = "complex128"
	case REAL:
		identType = "float64"
	}

	p.ident++
	for ; p.ns[p.ident].tok != NEW_LINE; p.ident++ {
		switch p.ns[p.ident].tok {
		case token.IDENT:
			p.initVars = append(p.initVars, initialVar{
				name: p.ns[p.ident].lit,
				typ:  identType,
			})
		case token.LPAREN:
			// Fortran example: INTEGER A(*)
			p.expect(token.LPAREN)
			p.ident++
			p.expect(token.MUL)
			p.ident++
			p.expect(token.RPAREN)
			p.initVars[len(p.initVars)-1].typ =
				"[]" + p.initVars[len(p.initVars)-1].typ
		case token.COMMA:
			// ignore
		default:
			p.addError("Cannot parseInit value : " + p.ns[p.ident].lit)
		}
	}

	return
}

func (p *parser) parseDo() (sDo goast.ForStmt) {

	p.ident++
	name := p.ns[p.ident].lit

	p.ident++
	p.expect(token.ASSIGN)

	p.ident++
	start := p.ns[p.ident].lit
	sDo.Init = &goast.AssignStmt{
		Lhs: []goast.Expr{
			goast.NewIdent(name),
		},
		Tok: token.DEFINE,
		Rhs: []goast.Expr{
			goast.NewIdent(start),
		},
	}

	p.ident++
	p.expect(token.COMMA)

	p.ident++
	finish := p.ns[p.ident].lit
	sDo.Cond = &goast.BinaryExpr{
		X:  goast.NewIdent(name),
		Op: token.LSS,
		Y:  goast.NewIdent(finish),
	}
	sDo.Post = &goast.IncDecStmt{
		X:   goast.NewIdent(name),
		Tok: token.INC,
	}

	p.ident++
	p.expect(NEW_LINE)

	sDo.Body = &goast.BlockStmt{
		Lbrace: 1,
		List:   p.transpileListStmt(),
	}

	p.gotoEndLine()

	return
}

func (p *parser) parseIf() (sIf goast.IfStmt) {

	p.ident++
	p.expect(token.LPAREN)

	p.ident++
	start := p.ident
	for counter := 1; p.ns[p.ident].tok != token.EOF; p.ident++ {
		var exit bool
		switch p.ns[p.ident].tok {
		case token.LPAREN:
			counter++
		case token.RPAREN:
			counter--
			if counter == 0 {
				exit = true
			}
		}
		if exit {
			break
		}
	}

	sIf.Cond = p.parseExpr(start, p.ident)

	p.expect(token.RPAREN)
	p.ident++

	if p.ns[p.ident].tok == THEN {
		p.gotoEndLine()
		p.ident++
		sIf.Body = &goast.BlockStmt{
			Lbrace: 1,
			List:   p.transpileListStmt(),
		}
	} else {
		sIf.Body = &goast.BlockStmt{
			Lbrace: 1,
			List:   p.parseStmt(),
		}
	}

	if p.ns[p.ident].tok == token.ELSE {
		p.ident++
		sIf.Else = &goast.BlockStmt{
			Lbrace: 1,
			List:   p.transpileListStmt(),
		}
	}

	return
}

func (p *parser) parseExpr(start, end int) (expr goast.Expr) {
	var str string
	for i := start; i < end; i++ {
		str += " " + p.ns[i].lit
	}
	fmt.Println("Expr : ", str)
	//TODO
	return &goast.BinaryExpr{
		X:  goast.NewIdent("temp"),
		Op: token.LSS,
		Y:  goast.NewIdent("4"),
	}
}

func (p *parser) parseExternal() {
	// TODO : remove variable because it is name of funciton
	// TODO : remember definition funciton

	p.expect(EXTERNAL)

	p.ident++
	for ; p.ns[p.ident].tok != token.EOF; p.ident++ {
		if p.ns[p.ident].tok == NEW_LINE {
			p.ident++
			break
		}
		switch p.ns[p.ident].tok {
		case token.IDENT:
			name := p.ns[p.ident].lit
			p.functionExternalName = append(p.functionExternalName, name)
			// fmt.Println("Function external: ", name)
		case token.COMMA:
			// ingore
		default:
			p.addError("Cannot parse External " + p.ns[p.ident].lit)
		}
	}
}

func (p *parser) parseStmt() (stmts []goast.Stmt) {
	switch p.ns[p.ident].tok {
	case INTEGER, CHARACTER, COMPLEX, LOGICAL, REAL:
		stmts = append(stmts, p.parseInit()...)

	case token.RETURN:
		stmts = append(stmts, &goast.ReturnStmt{})
		p.ident++

		p.expect(NEW_LINE)
		p.ident++

	case EXTERNAL:
		p.parseExternal()

	case NEW_LINE:
		// ignore
		p.ident++

	case token.IF:
		sIf := p.parseIf()
		stmts = append(stmts, &sIf)

	case DO:
		sDo := p.parseDo()
		stmts = append(stmts, &sDo)

	// case END:
	// 	// ignore
	// 	p.ident++
	//
	// 	// TODO : p.expect(NEW_LINE)
	// 	// p.ident++
	// 	fmt.Println("Go to end ->", p.getLine())

	default:
		start := p.ident
		for ; p.ident < len(p.ns); p.ident++ {
			if p.ns[p.ident].tok == NEW_LINE {
				stmts = append(stmts, &goast.ExprStmt{
					X: p.parseExpr(start, p.ident),
				})
				break
			}
		}

		p.ident++
	}
	return
}

func (p *parser) parseParamDecl() (fields []*goast.Field) {
	for ; p.ns[p.ident].tok != token.EOF; p.ident++ {
		var exit bool
		switch p.ns[p.ident].tok {
		case token.COMMA:
			// ignore
		case token.IDENT:
			id := p.ns[p.ident].lit
			field := &goast.Field{
				Names: []*goast.Ident{goast.NewIdent(id)},
				Type:  goast.NewIdent("int"),
			}
			fields = append(fields, field)
		case token.RPAREN:
			p.ident--
			exit = true
		default:
			p.addError("Cannot parse parameter decl " + p.ns[p.ident].lit)
			return
		}
		if exit {
			break
		}
	}
	return
}