package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/hooneun/aide/internal/storage"
)

// addToolCmd는 새로운 도구를 추가하는 명령어를 나타냅니다
var addToolCmd = &cobra.Command{
	Use:   "add-tool <도구명> <파일명> <파일설명>",
	Short: "새로운 AI 도구 지원을 추가합니다",
	Long: `새로운 AI 도구를 aide에 추가합니다.
도구명, 생성할 파일명, 파일설명을 입력받아 저장합니다.

예시:
  aide add-tool vscode .vscode/settings.json "VS Code 설정 파일"
  aide add-tool windsurf .windsurfrules "Windsurf 규칙 파일"`,
	Args: cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		toolName := args[0]
		fileName := args[1]
		description := args[2]

		storage, err := storage.New()
		if err != nil {
			fmt.Printf("오류: 저장소를 초기화할 수 없습니다: %v\n", err)
			os.Exit(1)
		}

		// 대화형으로 추가 설정 입력받기
		reader := bufio.NewReader(os.Stdin)

		fmt.Printf("새로운 도구 '%s' 설정:\n", toolName)
		fmt.Printf("파일명: %s\n", fileName)
		fmt.Printf("설명: %s\n", description)
		fmt.Print("파일 헤더 (선택사항, 엔터로 건너뛰기): ")
		
		header, _ := reader.ReadString('\n')
		header = strings.TrimSpace(header)

		fmt.Print("파일 구분자 (기본값: '# ---'): ")
		separator, _ := reader.ReadString('\n')
		separator = strings.TrimSpace(separator)
		if separator == "" {
			separator = "# ---"
		}

		// 도구 설정 저장
		err = storage.SaveToolConfig(toolName, fileName, description, header, separator)
		if err != nil {
			fmt.Printf("오류: 도구 설정을 저장할 수 없습니다: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("✅ 도구 '%s'가 성공적으로 추가되었습니다!\n", toolName)
		fmt.Printf("이제 'aide set %s <카테고리> <프롬프트>' 명령을 사용할 수 있습니다.\n", toolName)
	},
}

func init() {
	rootCmd.AddCommand(addToolCmd)
}