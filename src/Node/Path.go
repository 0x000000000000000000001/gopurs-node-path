package Node_Path

import (
	"path/filepath"
	"strings"
	"os"

	"gopurs/output/gopurs_runtime"
)

func Normalize(p string) string {
	return filepath.Clean(p)
}

func Concat(paths []interface{}) string {
	var strPaths []string
	for _, p := range paths {
		if v, ok := p.(gopurs_runtime.Value); ok {
			strPaths = append(strPaths, v.StrVal())
		} else {
			strPaths = append(strPaths, p.(string))
		}
	}
	return filepath.Join(strPaths...)
}

func Resolve(from []interface{}, to string, _ interface{}) interface{} {
	var strPaths []string
	for _, p := range from {
		if v, ok := p.(gopurs_runtime.Value); ok {
			strPaths = append(strPaths, v.StrVal())
		} else {
			strPaths = append(strPaths, p.(string))
		}
	}
	strPaths = append(strPaths, to)
	
	absPath, err := filepath.Abs(filepath.Join(strPaths...))
	if err != nil {
		panic(err)
	}
	return absPath
}

func Relative(from string, to string) string {
	rel, err := filepath.Rel(from, to)
	if err != nil {
		panic(err)
	}
	return rel
}

func Dirname(p string) string {
	return filepath.Dir(p)
}

func Basename(p string) string {
	return filepath.Base(p)
}

func BasenameWithoutExt(p string, ext string) string {
	base := filepath.Base(p)
	if strings.HasSuffix(base, ext) {
		return base[:len(base)-len(ext)]
	}
	return base
}

func Extname(p string) string {
	return filepath.Ext(p)
}

func Sep() string {
	return string(os.PathSeparator)
}

func Delimiter() string {
	return string(os.PathListSeparator)
}

func Parse(p string) map[string]interface{} {
	dir := filepath.Dir(p)
	base := filepath.Base(p)
	ext := filepath.Ext(p)
	name := base
	if len(ext) > 0 {
		name = base[:len(base)-len(ext)]
	}
	root := filepath.VolumeName(p)
	if root == "" && filepath.IsAbs(p) {
		root = "/"
	}
	
	return map[string]interface{}{
		"root": root,
		"dir":  dir,
		"base": base,
		"ext":  ext,
		"name": name,
	}
}

func IsAbsolute(p string) bool {
	return filepath.IsAbs(p)
}
