package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/hooneun/aide/internal/storage"
)

// listToolsCmd는 등록된 모든 도구를 나열하는 명령어입니다
var listToolsCmd = &cobra.Command{
	Use:   "list-tools",
	Short: "등록된 모든 AI 도구를 나열합니다",
	Long: `등록된 모든 AI 도구와 설정을 나열합니다.
기본 도구(claude, cursor)와 사용자가 추가한 도구들을 모두 보여줍니다.`,
	Run: func(cmd *cobra.Command, args []string) {
		storage, err := storage.New()
		if err != nil {
			fmt.Printf("오류: 저장소를 초기화할 수 없습니다: %v\n", err)
			return
		}

		fmt.Println("등록된 AI 도구 목록:")
		fmt.Println("==================")

		// 기본 도구들 출력
		fmt.Println("\n📋 기본 도구:")
		fmt.Println("  • claude      - CLAUDE.md 파일 생성 (Claude Code)")
		fmt.Println("  • cursor      - .cursorrules 파일 생성 (Cursor)")

		// 동적으로 추가된 도구들 출력
		configs, err := storage.ListToolConfigs()
		if err != nil {
			fmt.Printf("오류: 도구 목록을 가져올 수 없습니다: %v\n", err)
			return
		}

		if len(configs) > 0 {
			fmt.Println("\n🔧 사용자 추가 도구:")
			for _, config := range configs {
				fmt.Printf("  • %-12s - %s (%s)\n", config.Name, config.Description, config.FileName)
			}
		}

		fmt.Printf("\n총 %d개의 도구가 등록되어 있습니다.\n", 2+len(configs))
		fmt.Println("\n사용법:")
		fmt.Println("  aide set <도구> <카테고리> <프롬프트>")
		fmt.Println("  aide apply <도구> <카테고리>")
		fmt.Println("  aide add-tool <새도구명> <파일명> <설명>")
	},
}

func init() {
	rootCmd.AddCommand(listToolsCmd)
}