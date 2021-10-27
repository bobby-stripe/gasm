package examples

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/bobby-stripe/gasm/wasi"
	"github.com/bobby-stripe/gasm/wasm"
	"github.com/bobby-stripe/gasm/wasm/naivevm"
)

func Test_trap(t *testing.T) {
	buf, err := os.ReadFile("wasm/trap.wasm")
	require.NoError(t, err)

	mod, err := wasm.DecodeModule((buf))
	require.NoError(t, err)

	store := wasm.NewStore(naivevm.NewEngine())

	err = wasi.NewEnvironment().Register(store)
	require.NoError(t, err)

	err = store.Instantiate(mod, "test")
	require.NoError(t, err)

	_, _, err = store.CallFunction("test", "cause_panic")
	require.Error(t, err)
}
