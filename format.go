// 22 october 2015
package main

import (
	"fmt"
	"encoding/hex"
	"strings"

	"github.com/nimogit/reallymine/bridge"
)

func formatSectorPos(pos int64) string {
	return fmt.Sprintf("sector at 0x%X", pos)
}

func formatBridge(b bridge.Bridge) string {
	return fmt.Sprintf("bridge type %s", b.Name())
}

func formatKey(key []byte) string {
	return strings.ToUpper(hex.EncodeToString(key))
}
