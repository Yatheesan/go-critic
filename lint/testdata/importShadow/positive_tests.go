package checker_test

import (
	"fmt"
	"math"

	xmath "math"

	"github.com/go-critic/go-critic/lint"
)

var _ = xmath.E // Make sure xmath is used

func shadowImportedPackages() {
	fmt.Printf("Hello PI=%v, Rule=%v", math.Pi, lint.Rule{})

	/// shadow of imported package 'math'
	math := "some math"
	/// shadow of imported from 'github.com/go-critic/go-critic/lint' package 'lint'
	lint := "some lint"

	fmt.Printf("Hello math=%v, lint=%v", math, lint)
}

func genDeclShadow() {
	/// shadow of imported package 'math'
	const math = 1
	var (
		/// shadow of imported package 'fmt'
		fmt = 2
		/// shadow of imported from 'github.com/go-critic/go-critic/lint' package 'lint'
		lint = 3
	)
	_, _ = fmt, lint
}

/// shadow of imported package 'math'
/// shadow of imported package 'fmt'
func shadowedByParam1(math string, fmt int) {}

/// shadow of imported package 'math'
/// shadow of imported package 'fmt'
func shadowedByParam2() (math string, fmt int) { return }

type shadower struct{}

/// shadow of imported package 'fmt'
func (fmt shadower) f() {}

/// shadow of imported package 'xmath'
func renamedImportShadow(xmath int) {}