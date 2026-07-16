package main

import (
	"fmt"
	"os"
	"path/filepath"
	"protocol"
	"strings"
)

func printHelp() {
	fmt.Println(`Usage: protocol [options]

Options:
  -i <file>     Input Go source file with struct definitions
  -o <path>     Output file or directory path
  -t <language> Target language:
                  go, csharp, typescript (ts), javascript (js),
                  rust (rs), c, cpp (c++), gdscript, zig, c3, all
  -h            Show this help

Examples:
  protocol -i source/source.go -o export/go/protocol.go -t go
  protocol -i source/source.go -o export/zig/protocol.zig -t zig
  protocol -i source/source.go -o export/c3/protocol.c3 -t c3
  protocol -i source/source.go -o export/ -t all

If no arguments are provided, generates all languages to export/ using source/source.go.`)
}

func exportAll(exporter *protocol.ProtocolExporter, base string) error {
	if base == "" {
		base = "export"
	}
	base = strings.TrimSuffix(base, string(os.PathSeparator))
	base = strings.TrimSuffix(base, "/")
	base = strings.TrimSuffix(base, "\\")

	targets := []struct {
		path string
		lang int
	}{
		{filepath.Join(base, "go", "protocol.go"), protocol.GO},
		{filepath.Join(base, "csharp", "Protocol.cs"), protocol.CSHARP},
		{filepath.Join(base, "gdscript"), protocol.GD_SCRIPT},
		{filepath.Join(base, "typescript", "protocol.ts"), protocol.TYPE_SCRIPT},
		{filepath.Join(base, "c", "protocol.h"), protocol.C},
		{filepath.Join(base, "cpp", "protocol.h"), protocol.CPLUSPLUS},
		{filepath.Join(base, "javascript", "protocol.js"), protocol.JAVA_SCRIPT},
		{filepath.Join(base, "rust", "protocol.rs"), protocol.RUST},
		{filepath.Join(base, "zig", "protocol.zig"), protocol.ZIG},
		{filepath.Join(base, "c3", "protocol.c3"), protocol.C3},
	}

	for _, target := range targets {
		if err := exporter.Export(target.path, target.lang); err != nil {
			return fmt.Errorf("export %s: %w", target.path, err)
		}
	}
	return nil
}

func exportTarget(exporter *protocol.ProtocolExporter, target, output string) error {
	switch target {
	case "go":
		return exporter.Export(output, protocol.GO)
	case "csharp":
		return exporter.Export(output, protocol.CSHARP)
	case "gdscript":
		return exporter.Export(output, protocol.GD_SCRIPT)
	case "typescript", "ts":
		return exporter.Export(output, protocol.TYPE_SCRIPT)
	case "c":
		return exporter.Export(output, protocol.C)
	case "cpp", "c++":
		return exporter.Export(output, protocol.CPLUSPLUS)
	case "javascript", "js":
		return exporter.Export(output, protocol.JAVA_SCRIPT)
	case "rust", "rs":
		return exporter.Export(output, protocol.RUST)
	case "zig":
		return exporter.Export(output, protocol.ZIG)
	case "c3":
		return exporter.Export(output, protocol.C3)
	case "all":
		return exportAll(exporter, output)
	default:
		return fmt.Errorf("unsupported language: %s", target)
	}
}

func main() {
	var input, output, target string
	var help bool
	for idx, arg := range os.Args {
		switch arg {
		case "-i":
			if len(os.Args) > idx+1 {
				input = os.Args[idx+1]
			}
		case "-o":
			if len(os.Args) > idx+1 {
				output = os.Args[idx+1]
			}
		case "-t":
			if len(os.Args) > idx+1 {
				target = os.Args[idx+1]
			}
		case "-h", "--help":
			help = true
		}
	}

	if help {
		printHelp()
		return
	}

	exporter := &protocol.ProtocolExporter{}

	if input != "" && output != "" && target != "" {
		if err := exporter.Load(input); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		if err := exportTarget(exporter, target, output); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		return
	}

	if err := exporter.Load("source/source.go"); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if err := exportAll(exporter, "export"); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
