BRONZE_START=$(date +%s%3N)
unsetopt prompt_subst

preexec() {
	BRONZE_START=$(date +%s%3N)
}

precmd() {
	PROMPT="$(env status=$? jobs=$#jobstates cmdtime=$(($(date +%s%3N)-$BRONZE_START))ms bronze print "${BRONZE[@]}") "
}
