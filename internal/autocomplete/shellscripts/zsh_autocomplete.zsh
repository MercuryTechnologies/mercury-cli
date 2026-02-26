#!/bin/zsh
compdef ____APPNAME___zsh_autocomplete __APPNAME__

____APPNAME___zsh_autocomplete() {

  local -a opts
  local temp
  local exit_code

  temp=$(COMPLETION_STYLE=zsh "${words[1]}" __complete "${words[@]:1}")
  exit_code=$?

  # Check for custom file completion patterns
  # Patterns can appear anywhere in the word (e.g., inside quotes: 'my file is @file://path')
  local cur="${words[CURRENT]}"

  if [[ "$cur" = *'@'* ]]; then
    # Extract everything after the last @
    local after_last_at="${cur##*@}"

    if [[ $after_last_at =~ ^(file://|data://) ]]; then
      compset -P "*$MATCH"
      _files
    else
      compset -P '*@'
      _files
    fi
    return
  fi

  case $exit_code in
    10)
      # File completion behavior
      _files
      ;;
    11)
      # No completion behavior - return nothing
      return 1
      ;;
    0)
      # Default behavior - show command completions
      opts=("${(@f)temp}")
      _describe 'values' opts
      ;;
  esac
}
