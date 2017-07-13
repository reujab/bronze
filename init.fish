function fish_prompt; env status=$status jobs=(count (jobs -p)) cmdtime={$CMD_DURATION}ms bronze print $BRONZE; echo -n ' '; end
