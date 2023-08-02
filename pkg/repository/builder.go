package repository

import (
	"strings"
)

type commandBuilder struct {
	builder *strings.Builder
}

func newCommandBuilder(localPath string) *commandBuilder {
	builder := &strings.Builder{}
	builder.WriteString("git ")

	if localPath != "" {
		builder.WriteString("-C ")
		builder.WriteString(localPath)
		builder.WriteRune(' ')
	}

	return &commandBuilder{
		builder: builder,
	}
}

func (sb *commandBuilder) Write(strs ...string) {
	for i := 0; i < len(strs); i++ {
		sb.builder.WriteString(strs[i])
		sb.builder.WriteRune(' ')
	}
}

func (sb *commandBuilder) String() string {
	return strings.Trim(sb.builder.String(), " ")
}
