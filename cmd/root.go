package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd는 애플리케이션의 기본 명령어를 나타냅니다
var rootCmd = &cobra.Command{
	Use:   "aide",
	Short: "AI 개발 도구 프롬프트 관리 CLI",
	Long: `aide는 Claude Code, Cursor와 같은 AI 개발 도구의 프롬프트를 
관리하고 동기화하는 도구입니다.

프롬프트를 카테고리별로 저장하고, 프로젝트에 즉시 적용할 수 있습니다.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("aide - AI 개발 도구 프롬프트 관리 CLI")
		fmt.Println("도움말을 보려면 'aide --help'를 입력하세요.")
	},
}

// Execute는 모든 하위 명령어를 root 명령어에 추가하고 플래그를 적절히 설정합니다
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}