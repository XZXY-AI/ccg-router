# vhs tape: ccg-router dual-CLI demo
# Run: vhs docs/launch/demo-gif.tape
# Output: docs/demo.gif

Output docs/demo.gif

Set Shell "bash"
Set FontFamily "DejaVu Sans Mono"
Set FontSize 16
Set Width 1200
Set Height 700
Set Padding 24
Set Theme "Catppuccin Mocha"
Set TypingSpeed 30ms
Set PlaybackSpeed 1.0

# Scene 1: start the router
Type "ccg-router start &"
Sleep 200ms
Enter
Sleep 1.5s

Type "clear"
Enter
Sleep 200ms

# Scene 2: Claude Code points at it
Type "export ANTHROPIC_BASE_URL=http://127.0.0.1:17180"
Enter
Sleep 400ms
Type "claude code  # any prompt"
Enter
Sleep 1.8s

# Scene 3: same terminal, switch to Codex CLI
Type "export OPENAI_BASE_URL=http://127.0.0.1:17180"
Enter
Sleep 400ms
Type "codex --once 'list TODOs in this dir'"
Enter
Sleep 2s

# Scene 4: query the local ledger
Type "sqlite3 ~/.local/share/ccg-router/ledger.db 'select upstream, model, total_tokens from requests order by id desc limit 5;'"
Enter
Sleep 2s

# End on the ledger output so the still-frame thumbnail shows the value.
Sleep 1.5s
