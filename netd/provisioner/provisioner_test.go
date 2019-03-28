package provisioner

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBridgeNameIsNotMoreThan15Chars(t *testing.T) {
	got := generateBridgeName("qwertyuiopas")

	assert.LessOrEqualf(t, len(got), 15, "BridgeName length is not less than or equal to 15")
}

func TestBridgeNameIsFormedFromClusterNameWithBRasSuffix(t *testing.T) {
	got := generateBridgeName("qwerty")

	assert.Equal(t, got[len(got)-3:], "_br")
	assert.Equal(t, got, "qwerty_br")
}

func TestClusterNameIsTruncatedIfTooLarge(t *testing.T) {
	got := generateBridgeName("qwertyuiopasdfg")

	assert.Equal(t, got, "qwertyuiopas_br")
}
