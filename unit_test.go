package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenesis_whenTypical(t *testing.T) {
	diff := Diff("https://raw.githubusercontent.com/mihaigalos/sonarqube-diff/main/data_example/demo_baseline.html", "https://raw.githubusercontent.com/mihaigalos/sonarqube-diff/main/data_example/demo_baseline_3additions.html")

	assert.Contains(t, diff, "AWK40HIg-pl6AHs22K6U-manually-added")
	assert.Contains(t, diff, "AWK40IH6-pl6AHs22Mgc-manually-added")
	assert.Contains(t, diff, "AWK40INQ-pl6AHs22Mod-manually-added")

}
