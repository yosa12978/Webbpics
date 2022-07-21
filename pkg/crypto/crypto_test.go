package crypto_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yosa12978/webbpics/pkg/crypto"
)

func TestNewMD5(t *testing.T) {
	assert.Equal(t, crypto.NewMD5("password"), "5f4dcc3b5aa765d61d8327deb882cf99")
}

func TestNewUserToken(t *testing.T) {
	assert.NotNil(t, crypto.NewToken(32))
}
