package soap

import (
	"testing"
)

func TestParsing(t *testing.T) {
	createEnv, err := parseCreateWorkItem(payload2Attachments)
	if err != nil {
		t.Fatal("error parsing document")
	}

	witem := createEnv.CreateBody.Create.WorkItem
	if *witem.CorrespondenceCount != 10 {
		t.Fail()
	}
	if *witem.DocumentAttachmentCount != 2 {
		t.Fail()
	}
}

func BenchmarkCreateParse2Attachments(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := parseCreateWorkItem(payload2Attachments)
		if err != nil {
			panic(err)
		}

	}
}

func BenchmarkCreateParse20Attachments(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := parseCreateWorkItem(payload20Attachments)
		if err != nil {
			panic(err)
		}

	}
}

func BenchmarkCreateParse200Attachments(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := parseCreateWorkItem(payload200Attachments)
		if err != nil {
			panic(err)
		}

	}
}
