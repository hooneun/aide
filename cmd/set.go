package cmd

import (
	"fmt"
	"strings"

	"aide/internal/config"
	"aide/internal/storage"

	"github.com/spf13/cobra"
)

// setCmd는 프롬프트를 저장하는 명령어입니다
var setCmd = &cobra.Command{
	Use:   "set <도구> <카테고리> <프롬프트>",
	Short: "프롬프트를 저장합니다",
	Long: `특정 도구와 카테고리에 프롬프트를 저장합니다.

예시:
  aide set claude review "보안 취약점과 성능 문제를 체크해줘"
  aide set cursor backend "Go 모범 사례와 에러 핸들링에 집중해줘"`,
	Args: cobra.ExactArgs(3),
	RunE: func(cmd *cobra.Command, args []string) error {
		tool := args[0]
		category := args[1]
		prompt := args[2]

		// 설정 초기화
		cfg, err := config.New()
		if err != nil {
			return fmt.Errorf("설정을 초기화할 수 없습니다: %w", err)
		}

		// 지원되는 도구인지 확인
		if err := cfg.ValidateTool(tool); err != nil {
			return err
		}

		// 저장소 초기화
		store, err := storage.New()
		if err != nil {
			return fmt.Errorf("저장소를 초기화할 수 없습니다: %w", err)
		}

		// 카테고리 이름 검증
		if strings.TrimSpace(category) == "" {
			return fmt.Errorf("카테고리 이름은 비어있을 수 없습니다")
		}

		// 프롬프트 검증
		if strings.TrimSpace(prompt) == "" {
			return fmt.Errorf("프롬프트는 비어있을 수 없습니다")
		}

		// 프롬프트 저장
		if err := store.SavePrompt(tool, category, prompt); err != nil {
			return fmt.Errorf("프롬프트를 저장하는 중 오류가 발생했습니다: %w", err)
		}

		fmt.Printf("프롬프트가 저장되었습니다: %s/%s\n", tool, category)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(setCmd)
}