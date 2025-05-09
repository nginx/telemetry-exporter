//go:build generator

package main

import (
	"bytes"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/nginx/telemetry-exporter/cmd/generator/tests"
)

func TestGenerateScheme(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	parseCfg := parsingConfig{
		pkgName:     "tests",
		typeName:    "Data",
		loadPattern: "github.com/nginx/telemetry-exporter/cmd/generator/tests",
		buildFlags:  []string{"-tags=generator"},
	}

	_ = tests.Data{} // depends on the type being defined

	pResult, err := parse(parseCfg)

	g.Expect(err).ToNot(HaveOccurred())

	var buf bytes.Buffer

	schemeCfg := schemeGenConfig{
		namespace:          "gateway.nginx.org",
		protocol:           "avro",
		dataFabricDataType: "telemetry",
		record:             parseCfg.typeName,
		fields:             pResult.fields,
	}

	g.Expect(generateScheme(&buf, schemeCfg)).To(Succeed())

	g.Expect(buf.Bytes()).ToNot(BeEmpty())
}
