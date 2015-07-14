package parser

import (
	"github.com/yaozong/encoding/sample"
	"testing"
)

func BenchmarkReflectParse2Parties(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := ReflectParseB2bCreateWorkItem(sample.B2bCreateWorkItemWith2Parties)
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkReflectParse20Parties(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := ReflectParseB2bCreateWorkItem(sample.B2bCreateWorkItemWith20Parties)
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkReflectParse200Parties(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := ReflectParseB2bCreateWorkItem(sample.B2bCreateWorkItemWith200Parties)
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkReflectParse2000Parties(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := ReflectParseB2bCreateWorkItem(sample.B2bCreateWorkItemWith2000Parties)
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkStreamingParse2Parties(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := streamParseB2bCreateWorkItem(sample.B2bCreateWorkItemWith2Parties)
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkStreamingParse20Parties(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := streamParseB2bCreateWorkItem(sample.B2bCreateWorkItemWith20Parties)
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkStreamingParse200Parties(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := streamParseB2bCreateWorkItem(sample.B2bCreateWorkItemWith200Parties)
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkStreamingParse2000Parties(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := streamParseB2bCreateWorkItem(sample.B2bCreateWorkItemWith2000Parties)
		if err != nil {
			panic(err)
		}
	}
}
