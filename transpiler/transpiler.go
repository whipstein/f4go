package transpiler

import (
	"bytes"
	"fmt"

	goast "go/ast"
	"go/format"
	"go/token"

	"github.com/Konstantin8105/f4go/ast"
)

func TranspileAST(nss [][]ast.Node) (err error) {
	var file goast.File
	file.Name = goast.NewIdent("main")

	for _, ns := range nss {
		if len(ns) == 0 {
			continue
		}
		var fd goast.FuncDecl
		fd, err = trans(ns)
		if err != nil {
			return
		}
		goast.Print(token.NewFileSet(), fd)
		file.Decls = append(file.Decls, &fd)
	}

	var buf bytes.Buffer
	if err = format.Node(&buf, token.NewFileSet(), &file); err != nil {
		return
	}
	fmt.Println("Code:\n", buf.String())

	return
}

func trans(ns []ast.Node) (fd goast.FuncDecl, err error) {
	if len(ns) < 1 {
		return
	}
	if _, ok := ns[0].(ast.Function_decl); !ok {
		err = fmt.Errorf("Not function_decl : %#v", ns[0])
		return
	}
	n := ns[0].(ast.Function_decl)

	// function name
	if index, ok := ast.IsLink(n.Name); ok {
		var name string
		name, err = getName(ns[index-1], ns)
		if err != nil {
			return
		}
		fd.Name = goast.NewIdent(name)
	} else {
		panic("1")
	}

	// funciton type
	if index, ok := ast.IsLink(n.VarType); ok {
		fmt.Printf("VarType = %#v\n", ns[index-1])

		fd.Type = &goast.FuncType{}

	} else {
		fmt.Println("--- not found var type")
	}

	// function body
	if index, ok := ast.IsLink(n.Body); ok {
		var stmts []goast.Stmt
		stmts, err = transpileStmt(ns[index-1], ns)
		if err != nil {
			return
		}
		fd.Body = &goast.BlockStmt{
			Lbrace: 1,
			List:   stmts,
			Rbrace: 1,
		}
	}

	return
}

func isNull(ds []goast.Stmt) bool {
	for i := range ds {
		if ds[i] == nil {
			return true
		}
	}
	return false
}

func transpileIntegerCast(n ast.Integer_cst, ns []ast.Node) (name, t string, err error) {
	// if index, ok := ast.IsLink(n.VarInt); ok {
	// 	fmt.Printf("Name = %#v\n", ns[index-1])
	// 	name = ns[index-1].(ast.Identifier_node).Strg
	// }
	name = n.VarInt

	if index, ok := ast.IsLink(n.VarType); ok {
		t, err = transpileType(ns[index-1], ns)
		if err != nil {
			return
		}
		fmt.Printf("Type = %v\n", t)
	}

	return
}

func CastToGoType(fortranType string) (goType string, err error) {
	switch fortranType {
	case "integer(kind=4)":
		goType = "int"
	default:
		fmt.Printf("Cannot CastToGoType: %v\n", fortranType)
	}
	return
}

func transpileVarDecl(n ast.Var_decl, ns []ast.Node) (decl goast.Decl, err error) {

	var t, name string

	if index, ok := ast.IsLink(n.Name); ok {
		fmt.Printf("Name = %#v\n", ns[index-1])
		name = ns[index-1].(ast.Identifier_node).Strg
	}

	if index, ok := ast.IsLink(n.TypeD); ok {
		t, err = transpileType(ns[index-1], ns)
		if err != nil {
			return
		}
		t, err = CastToGoType(t)
		if err != nil {
			return
		}
		fmt.Printf("Type = %v\n", t)
	}

	genDecl := goast.GenDecl{
		Tok: token.VAR,
		Specs: []goast.Spec{
			&goast.ValueSpec{
				Names: []*goast.Ident{{Name: name}},
				Type:  goast.NewIdent(t),
			},
		},
	}
	decl = &genDecl
	return
}

func getName(n ast.Node, ns []ast.Node) (name string, err error) {
	// fmt.Printf("getName: %#v\n", n)
	switch n := n.(type) {
	case ast.Identifier_node:
		name = n.Strg
	case ast.Var_decl:
		name = n.Name
	case ast.Integer_cst:
		name = n.VarInt
	default:
		fmt.Printf("Name is not found: %#v\n", n)
	}
	// fmt.Println("name = ", name)
	if index, ok := ast.IsLink(name); ok {
		// fmt.Println(">>> ", index, ok)
		return getName(ns[index-1], ns)
	}
	return
}

func transpileDecl(n ast.Node, ns []ast.Node) (name, t string, err error) {
	switch n := n.(type) {
	// case ast.Var_decl:
	// 	name, t, err = transpileVarDecl(n, ns)
	case ast.Integer_cst:
		name, t, err = transpileIntegerCast(n, ns)
	default:
		fmt.Printf("Cannot transpileDecl: %#v\n", n)
	}
	return
}

func arithOperation(n0, n1 string, tk token.Token, ns []ast.Node) (
	expr goast.Expr, err error) {

	var left, right goast.Expr

	if index, ok := ast.IsLink(n0); ok {
		left, err = transpileExpr(ns[index-1], ns)
		if err != nil {
			return
		}
	}

	if index, ok := ast.IsLink(n1); ok {
		right, err = transpileExpr(ns[index-1], ns)
		if err != nil {
			return
		}
	}

	expr = &goast.BinaryExpr{
		X:  left,
		Op: tk,
		Y:  right,
	}

	return
}

func transpileExpr(n ast.Node, ns []ast.Node) (
	expr goast.Expr, err error) {
	switch n := n.(type) {
	case ast.Integer_cst:
		var name string
		name, err = getName(n, ns)
		if err != nil {
			return
		}
		expr = goast.NewIdent(name)

	case ast.Var_decl:
		var name string
		name, err = getName(n, ns)
		if err != nil {
			return
		}
		expr = goast.NewIdent(name)

	case ast.Plus_expr:
		expr, err = arithOperation(n.Op0, n.Op1, token.ADD, ns)

	case ast.Mult_expr:
		expr, err = arithOperation(n.Op0, n.Op1, token.MUL, ns)

	case ast.Trunc_div_expr:
		expr, err = arithOperation(n.Op0, n.Op1, token.QUO, ns)

	default:
		fmt.Printf("Cannot transpileExpr: %#v\n", n)
	}
	return
}

func transpileStmt(n ast.Node, ns []ast.Node) (
	decls []goast.Stmt, err error) {

	switch n := n.(type) {

	case ast.Bind_expr:
		var b goast.BlockStmt
		b.Rbrace = 1
		b.Lbrace = 1

		// parse var_decl
		fmt.Println("-----------")
		if index, ok := ast.IsLink(n.Vars); ok {
			// fmt.Printf("Var decl= %#v\n", ns[index-1])

			var decl goast.Decl
			decl, err = transpileVarDecl(ns[index-1].(ast.Var_decl), ns)
			if err != nil {
				return
			}

			if decl != nil {
				b.List = append(b.List, &goast.DeclStmt{Decl: decl})
			} else {
				fmt.Println("Decl is null")
			}
		}
		fmt.Println("++++++++++++")

		// parse body
		fmt.Println("-----------")
		if index, ok := ast.IsLink(n.Body); ok {

			// fmt.Printf("Expr = %#v\n", ns[index-1])
			var ds []goast.Stmt
			ds, err = transpileStmt(ns[index-1], ns)
			if err != nil {
				return
			}
			if isNull(ds) {
				fmt.Println("Decl is null")
			} else {
				b.List = append(b.List, ds...)
			}
		}
		fmt.Println("++++++++++++")

		decls = append(decls, &b)

	case ast.Modify_expr:
		// fmt.Printf("%#v\n", n)

		var left goast.Expr
		if index, ok := ast.IsLink(n.Op0); ok {
			left, err = transpileExpr(ns[index-1], ns)
			if err != nil {
				return
			}
		}

		var right goast.Expr
		if index, ok := ast.IsLink(n.Op1); ok {
			right, err = transpileExpr(ns[index-1], ns)
			if err != nil {
				return
			}
		}
		// fmt.Println("Modify_expr: ", left, " = ", right)

		if left == nil {
			fmt.Println("left is null")
		}
		if right == nil {
			fmt.Println("right is null")
		}
		decl := &goast.AssignStmt{
			Lhs: []goast.Expr{left},
			Tok: token.ASSIGN,
			Rhs: []goast.Expr{right},
		}
		decls = append(decls, decl)

	case ast.Statement_list:

		for i := range n.Vals {
			if index, ok := ast.IsLink(n.Vals[i]); ok {
				var decl []goast.Stmt
				decl, err = transpileStmt(ns[index-1], ns)
				if err != nil {
					return
				}
				decls = append(decls, decl...)
			}
		}

	default:
		fmt.Printf("Cannot transpileStmt: %#v\n", n)
	}
	return
}

func transpileType(n ast.Node, ns []ast.Node) (t string, err error) {
	switch n := n.(type) {
	case ast.Integer_type:
		if index, ok := ast.IsLink(n.Name); ok {
			t, err = transpileType(ns[index-1], ns)
			if err != nil {
				return
			}
		}

	case ast.Type_decl:
		if index, ok := ast.IsLink(n.Name); ok {
			t, err = transpileType(ns[index-1], ns)
			if err != nil {
				return
			}
		}

	case ast.Record_type:
		if index, ok := ast.IsLink(n.Name); ok {
			t, err = transpileType(ns[index-1], ns)
			if err != nil {
				return
			}
		}

	case ast.Identifier_node:
		t = n.Strg
	default:
		fmt.Printf("Cannot transpileType : %#v\n", n)
	}
	return
}