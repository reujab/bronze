BRONZE_START=$(date +%s%3N)
unsetopt prompt_subst

preexec() {
	BRONZE_START=$(date +%s%3N)
}

precmd() {
	PROMPT="$(env STATUS=$? JOBS=$#jobstates CMDTIME=$(($(date +%s%3N)-$BRONZE_START))ms bronze print "${BRONZE[@]}") "
}
