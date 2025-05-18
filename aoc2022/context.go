package aoc2022

import (
	"io"
	"os"
)

type Context struct {
    Day int
    Testing bool
    Input string
}

func ContextDefault() Context {
    return Context{0, false, ""}
}

func (ctx *Context) OnDay(day int) *Context {
    ctx.Day = day
    return ctx
}

func (ctx *Context) WithTesting(testing bool) *Context {
    ctx.Testing = testing
    return ctx
}

func (ctx *Context) WithInput(input string) *Context {
    ctx.Input = input
    return ctx
}

func (ctx *Context) WithInputFromFile(f *os.File) (*Context, error) {
    s, err := io.ReadAll(f)
    if err != nil {
        return nil, err
    }
    ctx.Input = string(s)
    return ctx, nil
}

func (ctx *Context) WithInputFromPath(p string) (*Context, error) {
    f, err := os.Open(p)
    if err != nil {
        return nil, err
    }
    defer f.Close()
    return ctx.WithInputFromFile(f)
}

func (ctx *Context) SplitLines() []string {
    return SplitLines(ctx.Input)
}