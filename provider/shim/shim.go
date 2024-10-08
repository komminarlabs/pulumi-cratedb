package shim

import (
	tfpf "github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/komminarlabs/terraform-provider-cratedb/internal/provider"
)

func NewProvider() tfpf.Provider {
	return provider.New("dev")()
}
