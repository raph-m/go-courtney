package shared

import (
	"strings"

	"github.com/dave/patsy"
	"github.com/dave/patsy/vos"
)

// Setup holds globals, environment and command line flags for the courtney
// command
type Setup struct {
	Env      vos.Env
	Paths    *patsy.Cache
	Enforce  bool
	Verbose  bool
	Short    bool
	Timeout  string
	Load     string
	Output   string
	TestArgs []string
	Packages []PackageSpec
}

// PackageSpec identifies a package by dir and path
type PackageSpec struct {
	Dir  string
	Path string
}

// Parse parses a slice of strings into the Packages slice
func (s *Setup) Parse(args []string) error {
	if len(args) == 0 {
		args = []string{"./..."}
	}
	packages := map[string]string{}
	for _, ppath := range args {
		var dir string
		recursive := false
		if strings.HasSuffix(ppath, "/...") {
			recursive = true
		}
		if strings.HasSuffix(ppath, "/") {
			ppath = strings.TrimSuffix(ppath, "/")
		}
		if ppath == "." {
			var err error
			dir, err = s.Env.Getwd()
			if err != nil {
				return err
			}
			ppath, err = s.Paths.Path(dir)
			if err != nil {
				return err
			}
		}

		if recursive {
			paths, err := s.Paths.Dirs(ppath)
			if err != nil {
				return err
			}

			for importPath, dir := range paths {
				packages[importPath] = dir
			}
		} else {
			var err error
			dir, err = s.Paths.Dir(ppath)
			if err != nil {
				return err
			}

			packages[ppath] = dir
		}
	}
	for ppath, dir := range packages {
		s.Packages = append(s.Packages, PackageSpec{Path: ppath, Dir: dir})
	}
	return nil
}
