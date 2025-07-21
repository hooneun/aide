# aide

AI 개발 환경 - AI 도구 프롬프트를 쉽게 관리하세요.

## aide란?

`aide`는 Claude Code, Cursor 등 AI 개발 도구들의 프롬프트를 관리하고 동기화해주는 도구입니다.

✨ **특별한 점**: 기본 제공 도구 외에도 VS Code, Windsurf, JetBrains 등 **어떤 AI 도구든 사용자가 직접 추가**하여 통합 관리할 수 있습니다!

## 주요 기능

- 📝 카테고리별로 프롬프트 저장 및 관리
- 🔄 프로젝트에 즉시 프롬프트 적용
- 🛠️ Claude Code, Cursor 기본 지원
- 🗂️ 도구별 여러 프롬프트 카테고리 지원
- ✨ **새 기능**: 사용자 정의 AI 도구 동적 추가 가능

## 설치

```bash
go install github.com/hooneun/aide@latest
```

## 빠른 시작

### 기본 사용법
```bash
# 프롬프트 추가
aide set cursor backend "Go 모범 사례와 에러 핸들링에 집중해줘"
aide set claude review "보안 취약점과 성능 문제를 체크해줘"

# 프롬프트 목록 확인
aide list cursor
aide list claude

# 현재 프로젝트에 적용
aide apply cursor backend
aide apply claude review

# 여러 프롬프트 동시 적용
aide apply cursor backend,frontend
```

### 🆕 새 도구 추가하기
```bash
# 새로운 AI 도구 추가
aide add-tool vscode .vscode/settings.json "VS Code 설정 파일"
aide add-tool windsurf .windsurfrules "Windsurf 규칙 파일"

# 등록된 도구 목록 확인
aide list-tools

# 추가한 도구로 프롬프트 사용
aide set vscode formatting '{"editor.formatOnSave": true, "editor.tabSize": 2}'
aide apply vscode formatting
```

## 명령어

### 기본 명령어

#### `aide set <도구> <카테고리> <프롬프트>`
특정 도구와 카테고리에 프롬프트를 저장합니다.

#### `aide list [도구]`
모든 프롬프트 또는 특정 도구의 프롬프트를 나열합니다.

#### `aide apply <도구> <카테고리>[,카테고리2,...]`
현재 프로젝트에 프롬프트를 적용합니다. 해당 파일을 생성하거나 내용을 추가합니다.

### 🆕 도구 관리 명령어

#### `aide add-tool <도구명> <파일명> <파일설명>`
새로운 AI 도구를 aide에 추가합니다. 대화형으로 파일 헤더와 구분자를 설정할 수 있습니다.

**예시:**
```bash
aide add-tool jetbrains .idea/aide-prompts.txt "JetBrains IDE 프롬프트 파일"
```

#### `aide list-tools`
등록된 모든 AI 도구(기본 + 사용자 추가)를 나열합니다.

## 지원하는 도구

### 📋 기본 제공 도구
- **Claude Code**: `CLAUDE.md` 파일 생성/업데이트
- **Cursor**: `.cursorrules` 파일 생성/업데이트

### 🔧 사용자 정의 도구
`aide add-tool` 명령어로 어떤 AI 도구든 추가할 수 있습니다!

**추가 가능한 도구 예시:**
- **VS Code**: `.vscode/settings.json` 설정 파일
- **Windsurf**: `.windsurfrules` 규칙 파일
- **JetBrains IDE**: `.idea/aide-prompts.txt` 프롬프트 파일
- **Vim/Neovim**: `.aide-prompts` 설정 파일
- **기타**: 프롬프트를 파일로 관리하는 모든 도구

## 동작 방식

프롬프트를 적용할 때 `aide`는 다음과 같이 동작합니다:

1. 도구의 설정 파일이 없으면 생성
2. 기존 파일에 구분선과 함께 프롬프트 추가
3. 이미 적용된 프롬프트 중복 방지

## 설정

### 📁 폴더 구조
프롬프트와 도구 설정은 `~/.aide/` 폴더에 저장됩니다:

```
~/.aide/
├── claude/          # Claude 프롬프트들
├── cursor/          # Cursor 프롬프트들
├── tools/           # 🆕 도구 설정 파일들 (JSON)
│   ├── vscode.json  # VS Code 도구 설정
│   └── windsurf.json # Windsurf 도구 설정
├── vscode/          # 🆕 VS Code 프롬프트들
└── windsurf/        # 🆕 Windsurf 프롬프트들
```

### ⚙️ 도구 설정 파일 형식
사용자 정의 도구는 JSON 형태로 설정이 저장됩니다:

```json
{
  "name": "vscode",
  "fileName": ".vscode/settings.json",
  "description": "VS Code 설정 파일",
  "header": "// VS Code 설정",
  "separator": "// ---"
}
```

## 라이선스

MIT
