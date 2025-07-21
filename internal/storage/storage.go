package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// ToolConfig는 도구별 설정을 저장하는 구조체입니다
type ToolConfig struct {
	Name        string `json:"name"`        // 도구 이름
	FileName    string `json:"fileName"`    // 생성할 파일명
	Description string `json:"description"` // 파일 설명
	Header      string `json:"header"`      // 파일 헤더 (선택사항)
	Separator   string `json:"separator"`   // 프롬프트 구분자
}

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

// SaveToolConfig는 도구 설정을 저장합니다
func (s *Storage) SaveToolConfig(name, fileName, description, header, separator string) error {
	config := ToolConfig{
		Name:        name,
		FileName:    fileName,
		Description: description,
		Header:      header,
		Separator:   separator,
	}

	// 설정 파일 경로
	configFile := filepath.Join(s.baseDir, "tools", name+".json")
	
	// tools 디렉터리 생성
	toolsDir := filepath.Join(s.baseDir, "tools")
	if err := os.MkdirAll(toolsDir, 0755); err != nil {
		return fmt.Errorf("도구 설정 디렉터리를 생성할 수 없습니다: %w", err)
	}

	// JSON 형태로 설정 저장
	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return fmt.Errorf("도구 설정을 직렬화할 수 없습니다: %w", err)
	}

	if err := os.WriteFile(configFile, data, 0644); err != nil {
		return fmt.Errorf("도구 설정을 저장할 수 없습니다: %w", err)
	}

	return nil
}

// GetToolConfig는 도구 설정을 가져옵니다
func (s *Storage) GetToolConfig(name string) (*ToolConfig, error) {
	configFile := filepath.Join(s.baseDir, "tools", name+".json")
	
	data, err := os.ReadFile(configFile)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("도구 설정을 찾을 수 없습니다: %s", name)
		}
		return nil, fmt.Errorf("도구 설정을 읽을 수 없습니다: %w", err)
	}

	var config ToolConfig
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("도구 설정을 파싱할 수 없습니다: %w", err)
	}

	return &config, nil
}

// ListToolConfigs는 등록된 모든 도구 목록을 반환합니다
func (s *Storage) ListToolConfigs() ([]ToolConfig, error) {
	toolsDir := filepath.Join(s.baseDir, "tools")
	
	entries, err := os.ReadDir(toolsDir)
	if err != nil {
		if os.IsNotExist(err) {
			return []ToolConfig{}, nil // 빈 목록 반환
		}
		return nil, fmt.Errorf("도구 목록을 가져올 수 없습니다: %w", err)
	}

	var configs []ToolConfig
	for _, entry := range entries {
		if !entry.IsDir() && filepath.Ext(entry.Name()) == ".json" {
			name := entry.Name()[:len(entry.Name())-5] // .json 확장자 제거
			
			config, err := s.GetToolConfig(name)
			if err != nil {
				continue // 오류가 있는 설정 파일은 건너뛰기
			}
			configs = append(configs, *config)
		}
	}

	return configs, nil
}