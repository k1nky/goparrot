package main

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/k1nky/goparrot/internal/config"
	lua "github.com/yuin/gopher-lua"
)

type EchoHTTPHandler func(w http.ResponseWriter, r *http.Request)

func runLuaBlock(code string) string {
	lvm := lua.NewState(lua.Options{
		CallStackSize:       120,
		MinimizeStackMemory: true,
	})
	defer lvm.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	lvm.SetContext(ctx)

	if err := lvm.DoString(code); err != nil {
		return ""
	}
	lv := lvm.Get(-1)
	if str, ok := lv.(lua.LString); ok {
		return string(str)
	}
	return ""
}

func makeHTTPHandler(options config.HandlerConfig) EchoHTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		var buf bytes.Buffer
		if options.DumpHeaders {
			buf.WriteString(fmt.Sprintf("%s %s %s\n", r.Method, r.RequestURI, r.Proto))
			for h, v := range r.Header {
				buf.WriteString(fmt.Sprintf("%s: %s\n", h, v))
			}
		}
		if options.DumpBody {
			buf.ReadFrom(r.Body)
			buf.WriteString("\n")
		}
		if options.ResponseLua != "" {
			luaResponse := runLuaBlock(options.ResponseLua)
			buf.WriteString(luaResponse)
			buf.WriteString("\n")
		}
		buf.WriteTo(w)
		w.WriteHeader(options.Code)
	}
}
