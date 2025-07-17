package storage

import (
	"os"
	"testing"
)

func TestStorage_SaveAndGetPrompt(t *testing.T) {
	// 임시 디렉터리 생성
	tmpDir, err := os.MkdirTemp("", "aide_test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	// 테스트용 Storage 생성
	storage := &Storage{baseDir: tmpDir}

	// 프롬프트 저장 테스트
	tool := "claude"
	category := "test"
	prompt := "테스트 프롬프트입니다"

	err = storage.SavePrompt(tool, category, prompt)
	if err != nil {
		t.Fatalf("프롬프트 저장 실패: %v", err)
	}

	// 저장된 프롬프트 가져오기 테스트
	retrievedPrompt, err := storage.GetPrompt(tool, category)
	if err != nil {
		t.Fatalf("프롬프트 가져오기 실패: %v", err)
	}

	if retrievedPrompt != prompt {
		t.Errorf("프롬프트가 일치하지 않습니다. 예상: %s, 실제: %s", prompt, retrievedPrompt)
	}
}

func TestStorage_ListPrompts(t *testing.T) {
	// 임시 디렉터리 생성
	tmpDir, err := os.MkdirTemp("", "aide_test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	// 테스트용 Storage 생성
	storage := &Storage{baseDir: tmpDir}

	// 여러 프롬프트 저장
	prompts := map[string]string{
		"review":  "리뷰 프롬프트",
		"backend": "백엔드 프롬프트",
		"test":    "테스트 프롬프트",
	}

	tool := "claude"
	for category, prompt := range prompts {
		err = storage.SavePrompt(tool, category, prompt)
		if err != nil {
			t.Fatalf("프롬프트 저장 실패: %v", err)
		}
	}

	// 프롬프트 목록 가져오기
	categories, err := storage.ListPrompts(tool)
	if err != nil {
		t.Fatalf("프롬프트 목록 가져오기 실패: %v", err)
	}

	if len(categories) != len(prompts) {
		t.Errorf("카테고리 개수가 일치하지 않습니다. 예상: %d, 실제: %d", len(prompts), len(categories))
	}

	// 모든 카테고리가 포함되었는지 확인
	categoryMap := make(map[string]bool)
	for _, category := range categories {
		categoryMap[category] = true
	}

	for expectedCategory := range prompts {
		if !categoryMap[expectedCategory] {
			t.Errorf("카테고리 %s가 목록에 없습니다", expectedCategory)
		}
	}
}