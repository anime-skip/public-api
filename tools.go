//go:build tools
// +build tools

package tools

import (
	_ "github.com/99designs/gqlgen"
	_ "github.com/cortesi/modd/cmd/modd"
	_ "github.com/onsi/ginkgo/v2"
	_ "github.com/onsi/ginkgo/v2/ginkgo/generators"
	_ "github.com/onsi/ginkgo/v2/ginkgo/internal"
)
