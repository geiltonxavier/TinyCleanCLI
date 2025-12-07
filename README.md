# TinyCleanCLI

> Nota pessoal: eu criei esta ferramenta para me ajudar a limpar o meu Mac de tempos em tempos. Uso por minha conta e risco; se você decidir usar, faça o mesmo. Revise cada sugestão antes de apagar qualquer coisa.
>
> Personal note: I built this tool to help me tidy up my Mac from time to time. I use it at my own risk; if you choose to run it, do the same. Review every suggestion before deleting anything.

## Uso rápido / Quick start

### Instalar / Install
```bash
go install github.com/geiltonxavier/TinyCleanCLI/cmd/tinycleancli@latest
```

### Comandos / Commands

- `tinycleancli scan --dry-run`
  - PT: Mostra o que poderia ser limpo sem apagar nada (modo seguro por padrão).
  - EN: Shows what could be cleaned without deleting anything (safe by default).

- `tinycleancli scan --dry-run --verbose`
  - PT: Lista tudo sem abreviações (pode ser longo).
  - EN: Lists everything without abbreviations (can be long).

- `tinycleancli scan --days 45`
  - PT: Considera apps/projetos inativos após 45 dias (placeholder por enquanto).
  - EN: Treats apps/projects as inactive after 45 days (placeholder for now).

- `tinycleancli scan --projects-path ~/Projects --projects-path ~/code`
  - PT: Define pastas adicionais para procurar projetos.
  - EN: Adds extra folders to search for projects.

- `tinycleancli scan --apps=false --projects=false --caches=true`
  - PT: Escaneia só caches/logs/lixo, ignora apps e projetos.
  - EN: Scan caches/logs/trash only, ignore apps and projects.

## Notas importantes / Important notes

- PT: Neste momento a lógica ainda é de rascunho/placeholder (sem apagar nada). Verifique a saída e adapte antes de ligar qualquer remoção real.
- EN: The logic is still placeholder/draft (no deletion yet). Check the output and adapt before enabling any real removal.

- PT: Use sempre `--dry-run` ao testar. Quando implementar deleção, valide diretórios e faça backup.
- EN: Always use `--dry-run` while testing. When you add deletion, validate targets and keep backups.

- PT: O foco é macOS (caches em `~/Library`, `/Library`, `~/.Trash`, etc.).
- EN: Target platform is macOS (caches under `~/Library`, `/Library`, `~/.Trash`, etc.).
