package cmd

import (
	"fmt"
	"sort"

	"aide/internal/config"
	"aide/internal/storage"

	"github.com/spf13/cobra"
)

// listCmd는 저장된 프롬프트를 나열하는 명령어입니다
var listCmd = &cobra.Command{
	Use:   "list [도구]",
	Short: "저장된 프롬프트를 나열합니다",
	Long: `모든 프롬프트 또는 특정 도구의 프롬프트를 나열합니다.

예시:
  aide list          # 모든 프롬프트 나열
  aide list claude   # Claude 프롬프트만 나열
  aide list cursor   # Cursor 프롬프트만 나열`,
	Args: cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		// 설정 초기화
		cfg, err := config.New()
		if err != nil {
			return fmt.Errorf("설정을 초기화할 수 없습니다: %w", err)
		}

		// 저장소 초기화
		store, err := storage.New()
		if err != nil {
			return fmt.Errorf("저장소를 초기화할 수 없습니다: %w", err)
		}

		// 특정 도구가 지정된 경우
		if len(args) == 1 {
			tool := args[0]
			
			// 지원되는 도구인지 확인
			if err := cfg.ValidateTool(tool); err != nil {
				return err
			}

			categories, err := store.ListPrompts(tool)
			if err != nil {
				return fmt.Errorf("프롬프트 목록을 가져오는 중 오류가 발생했습니다: %w", err)
			}

			if len(categories) == 0 {
				fmt.Printf("%s 도구에 저장된 프롬프트가 없습니다.\n", tool)
				return nil
			}

			// 카테고리 정렬
			sort.Strings(categories)

			fmt.Printf("=== %s 프롬프트 목록 ===\n", tool)
			for _, category := range categories {
				fmt.Printf("  - %s\n", category)
			}
			return nil
		}

		// 모든 도구의 프롬프트 나열
		allPrompts, err := store.ListAllPrompts()
		if err != nil {
			return fmt.Errorf("프롬프트 목록을 가져오는 중 오류가 발생했습니다: %w", err)
		}

		if len(allPrompts) == 0 {
			fmt.Println("저장된 프롬프트가 없습니다.")
			return nil
		}

		// 도구별로 정렬하여 출력
		tools := make([]string, 0, len(allPrompts))
		for tool := range allPrompts {
			tools = append(tools, tool)
		}
		sort.Strings(tools)

		fmt.Println("=== 저장된 프롬프트 목록 ===")
		for _, tool := range tools {
			categories := allPrompts[tool]
			if len(categories) == 0 {
				continue
			}

			fmt.Printf("\n%s:\n", tool)
			
			// 카테고리 정렬
			sort.Strings(categories)
			
			for _, category := range categories {
				fmt.Printf("  - %s\n", category)
			}
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}