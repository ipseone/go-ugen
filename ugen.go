package ugen // import go.ipse.one/ugen

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"strings"

	"github.com/google/uuid"
	"github.com/oklog/ulid/v2"
)

type Generator struct {
	AnyCharacter bool
	Base64       bool
	Hex          bool
	ULID         bool
	UUID         bool
	Lower        bool
	Upper        bool
	Prefix       string
	Suffix       string
	Separator    string
	WithLineFeed bool
}

func (g *Generator) Gen(writer io.Writer, count, length int) error {
	anyCharacter := g.AnyCharacter || g.Base64 || g.Hex

	var genBuf []byte

	switch {
	case g.ULID:
		length = 26
	case g.UUID:
		length = 36
	default:
		genBuf = make([]byte, length)
	}

	var encBuf []byte

	switch {
	case g.Base64:
		encBuf = make([]byte, base64.StdEncoding.EncodedLen(length))
	case g.Hex:
		encBuf = make([]byte, hex.EncodedLen(length))
	}

	builder := new(strings.Builder)

	for i := 0; i < count; i++ {
		switch {
		case g.ULID:
			u := ulid.Make()

			switch {
			case g.Base64, g.Hex:
				genBuf = u[:]
			default:
				genBuf = []byte(u.String())
			}
		case g.UUID:
			u, err := uuid.NewRandom()
			if err != nil {
				return fmt.Errorf("failed to generate UUID: %w", err)
			}

			switch {
			case g.Base64, g.Hex:
				genBuf = u[:]
			default:
				genBuf = []byte(u.String())
			}
		case anyCharacter:
			gen(length, genBuf, nil)
		default:
			gen(length, genBuf, digitLowerUpper)
		}

		builder.Reset()

		if i > 0 {
			builder.WriteString(g.Separator)
		}

		if g.Prefix != "" {
			builder.WriteString(g.Prefix)
		}

		switch {
		case g.Base64:
			base64.StdEncoding.Encode(encBuf, genBuf)
			builder.Write(encBuf)
		case g.Hex:
			hex.Encode(encBuf, genBuf)
			builder.Write(encBuf)
		default:
			builder.Write(genBuf)
		}

		if g.Suffix != "" {
			builder.WriteString(g.Suffix)
		}

		switch {
		case g.Lower:
			fmt.Fprint(writer, strings.ToLower(builder.String()))
		case g.Upper:
			fmt.Fprint(writer, strings.ToUpper(builder.String()))
		default:
			fmt.Fprint(writer, builder.String())
		}
	}

	if g.WithLineFeed {
		fmt.Fprint(writer, "\n")
	}

	return nil
}
