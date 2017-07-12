PROMPT_COMMAND=bronze_prompt
bronze_prompt() {
	PS1="$(STATUS=$? JOBS=$(jobs -p | wc -l) bronze print "${BRONZE[@]}") "
}
