# aide

AI 개발 환경 - AI 도구 프롬프트를 쉽게 관리하세요.

## aide란?

`aide`는 Claude Code, Cursor 같은 AI 개발 도구들의 프롬프트를 관리하고 동기화해주는 도구입니다. 프로젝트마다 프롬프트를 복사-붙여넣기할 필요가 없어요!

## 주요 기능

- 📝 카테고리별로 프롬프트 저장 및 관리
- 🔄 프로젝트에 즉시 프롬프트 적용
- 🛠️ Claude Code, Cursor 지원
- 🗂️ 도구별 여러 프롬프트 카테고리 지원

## 설치

```bash
go install github.com/yourusername/aide@latest
```

## 빠른 시작

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

## 명령어

### `aide set <도구> <카테고리> <프롬프트>`
특정 도구와 카테고리에 프롬프트를 저장합니다.

### `aide list [도구]`
모든 프롬프트 또는 특정 도구의 프롬프트를 나열합니다.

### `aide apply <도구> <카테고리>[,카테고리2,...]`
현재 프로젝트에 프롬프트를 적용합니다. 해당 파일을 생성하거나 내용을 추가합니다.

## 지원하는 도구

- **Claude Code**: `CLAUDE.md` 파일 생성/업데이트
- **Cursor**: `.cursorrules` 파일 생성/업데이트

## 동작 방식

프롬프트를 적용할 때 `aide`는 다음과 같이 동작합니다:

1. 도구의 설정 파일이 없으면 생성
2. 기존 파일에 구분선과 함께 프롬프트 추가
3. 이미 적용된 프롬프트 중복 방지

## 설정

프롬프트는 `~/.aide/` 폴더에 도구별, 카테고리별로 저장됩니다.

## 라이선스

MIT
