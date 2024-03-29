package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type templateVar struct {
	Type1 string
	Type2 string
	Size  int
}

var codeTemplate = `// Code generated by gen-matrix DO NOT EDIT.
package geom

type {{ .Type1 }} [{{ .Size }}][{{ .Size }}]float64

func (m *{{ .Type1 }}) determinant() float64 {
	r := 0.
	for i, v := range m[0] {
		r += v * m.cofactor(0, i)
	}
	return r
}

func (m *{{ .Type1 }}) cofactor(row, column int) float64 {
	return m.minor(row, column) * sign(row, column)
}

func (m *{{ .Type1 }}) minor(row, column int) float64 {
	return m.submatrix(row, column).determinant()
}

func (m *{{ .Type1 }}) submatrix(row, column int) *{{ .Type2 }} {
	r := &{{ .Type2 }}{}
	for ri, mi := 0, 0; mi < {{.Size}}; mi++ {
		if mi == row {
			continue
		}
		for rj, mj := 0, 0; mj < {{.Size}}; mj++ {
			if mj == column {
				continue
			}
			r[ri][rj] = m[mi][mj]
			rj++
		}
		ri++
	}
	return r
}

func (m *{{ .Type1 }}) inverse() *{{ .Type1 }} {
	determinant := m.determinant()
	inverse := &{{ .Type1 }}{}
	for i := 0; i < {{ .Size }}; i++ {
		for j := 0; j < {{ .Size }}; j++ {
			inverse[j][i] = m.cofactor(i, j) / determinant
		}
	}
	return inverse
}


`

func main() {
	vars := []templateVar{
		{
			Type1: "matrix4x4",
			Type2: "matrix3x3",
			Size:  4,
		},
		{
			Type1: "matrix3x3",
			Type2: "matrix2x2",
			Size:  3,
		},
	}

	for _, v := range vars {
		generateCode(v)
	}
}

func generateCode(vars templateVar) {
	filename := fmt.Sprintf("%s.go", strings.ToLower(vars.Type1))
	tmpl := template.New(filename)
	tmpl, err := tmpl.Parse(codeTemplate)
	if err != nil {
		log.Fatal(err)
	}

	generated := new(bytes.Buffer)
	err = tmpl.Execute(generated, vars)
	if err != nil {
		log.Fatalf("Failed to execute template: %v", err)
	}
	generatedFile := filepath.Join(filename)
	err = os.WriteFile(generatedFile, generated.Bytes(), 0644)
	if err != nil {
		log.Fatalf("failed to write formatted file: %v", err)
	}
}
