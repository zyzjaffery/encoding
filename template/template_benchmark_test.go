package template

import (
	"encoding/json"
	"github.com/yaozong/encoding/parser"
	"github.com/yaozong/encoding/sample"
	"github.com/yaozong/encoding/types"
	"testing"
)

var (
	workitemWith2Parties    *types.SoapEnvelope
	workitemWith20Parties   *types.SoapEnvelope
	workitemWith200Parties  *types.SoapEnvelope
	workitemWith2000Parties *types.SoapEnvelope
)

func init() {
	workitemWith2Parties, _ = parser.ReflectParseB2bCreateWorkItem(sample.B2bCreateWorkItemWith2Parties)
	workitemWith20Parties, _ = parser.ReflectParseB2bCreateWorkItem(sample.B2bCreateWorkItemWith20Parties)
	workitemWith200Parties, _ = parser.ReflectParseB2bCreateWorkItem(sample.B2bCreateWorkItemWith200Parties)
	workitemWith2000Parties, _ = parser.ReflectParseB2bCreateWorkItem(sample.B2bCreateWorkItemWith2000Parties)
}

func BenchmarkJsonMarshal2Parties(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := json.Marshal(workitemWith2Parties)
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkJsonMarshal20Parties(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := json.Marshal(workitemWith20Parties)
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkJsonMarshal200Parties(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := json.Marshal(workitemWith200Parties)
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkJsonMarshal2000Parties(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := json.Marshal(workitemWith2000Parties)
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkTemplate2Parties(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := MarshalWithJsonTemplate(workitemWith2Parties)
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkTemplate20Parties(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := MarshalWithJsonTemplate(workitemWith20Parties)
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkTemplate200Parties(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := MarshalWithJsonTemplate(workitemWith200Parties)
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkTemplate2000Parties(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := MarshalWithJsonTemplate(workitemWith2000Parties)
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkMinimisedTemplate2Parties(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := MarshalWithMinimisedJsonTemplate(workitemWith2Parties)
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkMinimisedTemplate20Parties(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := MarshalWithMinimisedJsonTemplate(workitemWith20Parties)
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkMinimisedTemplate200Parties(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := MarshalWithMinimisedJsonTemplate(workitemWith200Parties)
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkMinimisedTemplate2000Parties(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := MarshalWithMinimisedJsonTemplate(workitemWith2000Parties)
		if err != nil {
			panic(err)
		}
	}
}
