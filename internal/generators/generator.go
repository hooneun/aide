package generators

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// Generator는 파일 생성기 인터페이스입니다
type Generator interface {
	Generate(filePath string, prompts []string) error
}

// ClaudeGenerator는 CLAUDE.md 파일을 생성합니다
type ClaudeGenerator struct{}

// CursorGenerator는 .cursorrules 파일을 생성합니다
type CursorGenerator struct{}

// NewGenerator는 도구에 따른 적절한 생성기를 반환합니다
func NewGenerator(tool string) (Generator, error) {
	switch tool {
	case "claude":
		return &ClaudeGenerator{}, nil
	case "cursor":
		return &CursorGenerator{}, nil
	default:
		return nil, fmt.Errorf("지원되지 않는 도구입니다: %s", tool)
	}
}

// Generate는 CLAUDE.md 파일을 생성하거나 업데이트합니다
func (g *ClaudeGenerator) Generate(filePath string, prompts []string) error {
	var content strings.Builder
	
	// 기존 파일이 있는지 확인
	existingContent, err := os.ReadFile(filePath)
	if err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("파일을 읽을 수 없습니다: %w", err)
	}

	// 기존 내용이 있으면 추가, 없으면 새로 생성
	if len(existingContent) > 0 {
		content.WriteString(string(existingContent))
		content.WriteString("\n\n")
	}

	// aide로 추가된 섹션임을 표시
	content.WriteString("# aide 프롬프트\n\n")
	content.WriteString(fmt.Sprintf("다음 프롬프트는 aide에 의해 %s에 추가되었습니다.\n\n", time.Now().Format("2006-01-02 15:04:05")))

	// 각 프롬프트 추가
	for i, prompt := range prompts {
		if i > 0 {
			content.WriteString("\n---\n\n")
		}
		content.WriteString(prompt)
		content.WriteString("\n")
	}

	// 파일에 쓰기
	if err := os.WriteFile(filePath, []byte(content.String()), 0644); err != nil {
		return fmt.Errorf("파일을 저장할 수 없습니다: %w", err)
	}

	return nil
}

// Generate는 .cursorrules 파일을 생성하거나 업데이트합니다
func (g *CursorGenerator) Generate(filePath string, prompts []string) error {
	var content strings.Builder
	
	// 기존 파일이 있는지 확인
	existingContent, err := os.ReadFile(filePath)
	if err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("파일을 읽을 수 없습니다: %w", err)
	}

	// 기존 내용이 있으면 추가, 없으면 새로 생성
	if len(existingContent) > 0 {
		content.WriteString(string(existingContent))
		content.WriteString("\n\n")
	}

	// aide로 추가된 섹션임을 표시
	content.WriteString("# aide 프롬프트\n")
	content.WriteString(fmt.Sprintf("# 다음 규칙은 aide에 의해 %s에 추가되었습니다.\n\n", time.Now().Format("2006-01-02 15:04:05")))

	// 각 프롬프트 추가
	for i, prompt := range prompts {
		if i > 0 {
			content.WriteString("\n---\n\n")
		}
		content.WriteString(prompt)
		content.WriteString("\n")
	}

	// 파일에 쓰기
	if err := os.WriteFile(filePath, []byte(content.String()), 0644); err != nil {
		return fmt.Errorf("파일을 저장할 수 없습니다: %w", err)
	}

	return nil
}

// CheckDuplicatePrompts는 중복된 프롬프트를 확인합니다
func CheckDuplicatePrompts(filePath string, newPrompts []string) ([]string, error) {
	// 기존 파일 내용 읽기
	existingContent, err := os.ReadFile(filePath)
	if err != nil && !os.IsNotExist(err) {
		return nil, fmt.Errorf("파일을 읽을 수 없습니다: %w", err)
	}

	existingText := string(existingContent)
	var uniquePrompts []string

	// 새로운 프롬프트 중 중복되지 않은 것만 추가
	for _, prompt := range newPrompts {
		if !strings.Contains(existingText, strings.TrimSpace(prompt)) {
			uniquePrompts = append(uniquePrompts, prompt)
		}
	}

	return uniquePrompts, nil
}