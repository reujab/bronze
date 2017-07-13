PROMPT_COMMAND=bronze_prompt
bronze_prompt() {
	PS1="$(status=$? jobs=$(jobs -p | wc -l) bronze print "${BRONZE[@]}") "
}
