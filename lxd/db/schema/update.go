package schema

import (
	"database/sql"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"

	_ "github.com/mattn/go-sqlite3" // For opening the in-memory database
)

// DotGo writes '<name>.go' source file in the package of the calling function, containing
// SQL statements that match the given schema updates.
//
// The <name>.go file contains a "flattened" render of all given updates and
// can be used to initialize brand new databases using Schema.Fresh().
func DotGo(updates map[int]Update, name string) error {
	// Apply all the updates that we have on a pristine database and dump
	// the resulting schema.
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		return fmt.Errorf("failed to open schema.go for writing: %w", err)
	}

	schema := NewFromMap(updates)

	_, err = schema.Ensure(db)
	if err != nil {
		return err
	}

	dump, err := schema.Dump(db)
	if err != nil {
		return err
	}

	// Passing 1 to runtime.Caller identifies our caller.
	_, filename, _, _ := runtime.Caller(1)

	// runtime.Caller returns the path after "${GOPATH}/src" when used with `go generate`.
	if strings.HasPrefix(filename, "github.com") {
		filename = filepath.Join(os.Getenv("GOPATH"), "src", filename)
	}

	file, err := os.Create(path.Join(path.Dir(filename), name+".go"))
	if err != nil {
		return fmt.Errorf("failed to open Go file for writing: %w", err)
	}

	pkg := path.Base(path.Dir(filename))
	_, err = file.Write([]byte(fmt.Sprintf(dotGoTemplate, pkg, dump)))
	if err != nil {
		return fmt.Errorf("failed to write to Go file: %w", err)
	}

	return nil
}

// Template for schema files (can't use backticks since we need to use backticks
// inside the template itself).
const dotGoTemplate = "package %s\n\n" +
	"// DO NOT EDIT BY HAND\n" +
	"//\n" +
	"// This code was generated by the schema.DotGo function. If you need to\n" +
	"// modify the database schema, please add a new schema update to update.go\n" +
	"// and the run 'make update-schema'.\n" +
	"const freshSchema = `\n" +
	"%s`\n"
