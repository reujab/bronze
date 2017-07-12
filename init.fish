function fish_prompt; env STATUS=$status JOBS=(count (jobs -p)) CMDTIME={$CMD_DURATION}ms bronze print $BRONZE; echo -n ' '; end
