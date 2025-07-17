package config

import (
	"fmt"
	"os"
	"path/filepath"
)

// Config는 애플리케이션 설정을 관리하는 구조체입니다
type Config struct {
	AideDir string
}

// New는 새로운 Config 인스턴스를 생성합니다
func New() (*Config, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("홈 디렉터리를 찾을 수 없습니다: %w", err)
	}

	aideDir := filepath.Join(homeDir, ".aide")
	return &Config{AideDir: aideDir}, nil
}

// GetStorageDir는 저장소 디렉터리 경로를 반환합니다
func (c *Config) GetStorageDir() string {
	return c.AideDir
}

// GetToolDir는 특정 도구의 디렉터리 경로를 반환합니다
func (c *Config) GetToolDir(tool string) string {
	return filepath.Join(c.AideDir, tool)
}

// GetPromptFile는 프롬프트 파일 경로를 반환합니다
func (c *Config) GetPromptFile(tool, category string) string {
	return filepath.Join(c.AideDir, tool, category+".txt")
}

// ValidateTool은 지원되는 도구인지 확인합니다
func (c *Config) ValidateTool(tool string) error {
	supportedTools := []string{"claude", "cursor"}
	
	for _, supportedTool := range supportedTools {
		if tool == supportedTool {
			return nil
		}
	}
	
	return fmt.Errorf("지원되지 않는 도구입니다: %s (지원 도구: claude, cursor)", tool)
}

// GetCurrentDir는 현재 작업 디렉터리를 반환합니다
func (c *Config) GetCurrentDir() (string, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("현재 디렉터리를 가져올 수 없습니다: %w", err)
	}
	return currentDir, nil
}

// GetTargetFile은 도구에 따른 대상 파일 경로를 반환합니다
func (c *Config) GetTargetFile(tool string) (string, error) {
	currentDir, err := c.GetCurrentDir()
	if err != nil {
		return "", err
	}

	switch tool {
	case "claude":
		return filepath.Join(currentDir, "CLAUDE.md"), nil
	case "cursor":
		return filepath.Join(currentDir, ".cursorrules"), nil
	default:
		return "", fmt.Errorf("지원되지 않는 도구입니다: %s", tool)
	}
}