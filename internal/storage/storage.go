package storage

import (
	"fmt"
	"os"
	"path/filepath"
)

// Storage는 프롬프트 저장소를 관리하는 구조체입니다
type Storage struct {
	baseDir string
}

// New는 새로운 Storage 인스턴스를 생성합니다
func New() (*Storage, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("홈 디렉터리를 찾을 수 없습니다: %w", err)
	}

	baseDir := filepath.Join(homeDir, ".aide")
	
	// .aide 디렉터리가 없으면 생성
	if err := os.MkdirAll(baseDir, 0755); err != nil {
		return nil, fmt.Errorf("저장소 디렉터리를 생성할 수 없습니다: %w", err)
	}

	return &Storage{baseDir: baseDir}, nil
}

// SavePrompt는 프롬프트를 저장합니다
func (s *Storage) SavePrompt(tool, category, prompt string) error {
	// 도구별 디렉터리 생성
	toolDir := filepath.Join(s.baseDir, tool)
	if err := os.MkdirAll(toolDir, 0755); err != nil {
		return fmt.Errorf("도구 디렉터리를 생성할 수 없습니다: %w", err)
	}

	// 프롬프트 파일 경로
	promptFile := filepath.Join(toolDir, category+".txt")
	
	// 프롬프트를 파일에 저장
	if err := os.WriteFile(promptFile, []byte(prompt), 0644); err != nil {
		return fmt.Errorf("프롬프트를 저장할 수 없습니다: %w", err)
	}

	return nil
}

// GetPrompt는 저장된 프롬프트를 가져옵니다
func (s *Storage) GetPrompt(tool, category string) (string, error) {
	promptFile := filepath.Join(s.baseDir, tool, category+".txt")
	
	content, err := os.ReadFile(promptFile)
	if err != nil {
		if os.IsNotExist(err) {
			return "", fmt.Errorf("프롬프트를 찾을 수 없습니다: %s/%s", tool, category)
		}
		return "", fmt.Errorf("프롬프트를 읽을 수 없습니다: %w", err)
	}

	return string(content), nil
}

// ListPrompts는 특정 도구의 모든 프롬프트 카테고리를 나열합니다
func (s *Storage) ListPrompts(tool string) ([]string, error) {
	toolDir := filepath.Join(s.baseDir, tool)
	
	entries, err := os.ReadDir(toolDir)
	if err != nil {
		if os.IsNotExist(err) {
			return []string{}, nil // 빈 목록 반환
		}
		return nil, fmt.Errorf("프롬프트 목록을 가져올 수 없습니다: %w", err)
	}

	var categories []string
	for _, entry := range entries {
		if !entry.IsDir() && filepath.Ext(entry.Name()) == ".txt" {
			// .txt 확장자 제거
			name := entry.Name()
			category := name[:len(name)-4]
			categories = append(categories, category)
		}
	}

	return categories, nil
}

// ListAllPrompts는 모든 도구의 프롬프트를 나열합니다
func (s *Storage) ListAllPrompts() (map[string][]string, error) {
	result := make(map[string][]string)
	
	entries, err := os.ReadDir(s.baseDir)
	if err != nil {
		if os.IsNotExist(err) {
			return result, nil // 빈 맵 반환
		}
		return nil, fmt.Errorf("저장소를 읽을 수 없습니다: %w", err)
	}

	for _, entry := range entries {
		if entry.IsDir() {
			tool := entry.Name()
			categories, err := s.ListPrompts(tool)
			if err != nil {
				return nil, err
			}
			result[tool] = categories
		}
	}

	return result, nil
}