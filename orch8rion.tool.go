// This file was created by `orch8rion pin`, and is used to ensure the
// `go.mod` file contains the necessary entries to ensure repeatable builds when
// using `orch8rion`. It is also used to set up which integrations are enabled.

//go:build tools

//go:generate go run github.com/senforsce/orch8rion pin -generate

package tools

// Imports in this file determine which tracer integrations are enabled in
// orch8rion. New integrations can be automatically discovered by running
// `orch8rion pin` again. You can also manually add new imports here to
// enable additional integrations. When doing so, you can run `orch8rion pin`
// to make sure manually added integrations are valid (i.e, the imported package
// includes a valid `orch8rion.yml` file).
import (
	// Ensures `orch8rion` is present in `go.mod` so that builds are repeatable.
	// Do not remove.
	_ "github.com/senforsce/orch8rion" // integration

	_ "github.com/DataDog/dd-trace-go/orchestrion/all/v2" // integration
)
