package cmd

import (
	"fmt"
	"strings"

	"github.com/hooneun/aide/internal/config"
	"github.com/hooneun/aide/internal/generators"
	"github.com/hooneun/aide/internal/storage"

	"github.com/spf13/cobra"
)

// applyCmd는 프롬프트를 현재 프로젝트에 적용하는 명령어입니다
var applyCmd = &cobra.Command{
	Use:   "apply <도구> <카테고리>[,카테고리2,...]",
	Short: "프롬프트를 현재 프로젝트에 적용합니다",
	Long: `저장된 프롬프트를 현재 프로젝트에 적용합니다.
해당 도구의 설정 파일을 생성하거나 내용을 추가합니다.

예시:
  aide apply claude review                    # Claude 리뷰 프롬프트 적용
  aide apply cursor backend                   # Cursor 백엔드 프롬프트 적용
  aide apply cursor backend,frontend          # 여러 프롬프트 동시 적용`,
	Args: cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		tool := args[0]
		categoriesArg := args[1]

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

		// 카테고리 목록 파싱
		categories := strings.Split(categoriesArg, ",")
		for i, category := range categories {
			categories[i] = strings.TrimSpace(category)
		}

		// 각 카테고리에 대해 프롬프트 가져오기
		var prompts []string
		for _, category := range categories {
			if category == "" {
				continue
			}

			prompt, err := store.GetPrompt(tool, category)
			if err != nil {
				return fmt.Errorf("프롬프트를 가져오는 중 오류가 발생했습니다: %w", err)
			}

			prompts = append(prompts, prompt)
		}

		if len(prompts) == 0 {
			return fmt.Errorf("적용할 프롬프트가 없습니다")
		}

		// 대상 파일 경로 가져오기
		targetFile, err := cfg.GetTargetFile(tool)
		if err != nil {
			return fmt.Errorf("대상 파일을 결정할 수 없습니다: %w", err)
		}

		// 파일 생성기 생성
		generator, err := generators.NewGenerator(tool)
		if err != nil {
			return fmt.Errorf("파일 생성기를 초기화할 수 없습니다: %w", err)
		}

		// 중복 프롬프트 확인
		uniquePrompts, err := generators.CheckDuplicatePrompts(targetFile, prompts)
		if err != nil {
			return fmt.Errorf("중복 프롬프트를 확인하는 중 오류가 발생했습니다: %w", err)
		}

		if len(uniquePrompts) == 0 {
			fmt.Printf("모든 프롬프트가 이미 %s에 적용되어 있습니다.\n", targetFile)
			return nil
		}

		// 프롬프트 적용
		if err := generator.Generate(targetFile, uniquePrompts); err != nil {
			return fmt.Errorf("프롬프트를 적용하는 중 오류가 발생했습니다: %w", err)
		}

		fmt.Printf("프롬프트가 %s에 성공적으로 적용되었습니다.\n", targetFile)
		fmt.Printf("적용된 카테고리: %s\n", strings.Join(categories, ", "))

		return nil
	},
}

func init() {
	rootCmd.AddCommand(applyCmd)
}
