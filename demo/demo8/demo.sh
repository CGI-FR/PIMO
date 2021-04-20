# wrong approach number 1
cat data.json | jq -c "."  | pimo -c masking-wrong.yml

# wrong approach number 2
cat data.json | jq -c "."  | pimo -c masking-alsowrong.yml | jq

# correct approach - step 2
cat data.json | jq -c "."  | pimo -c masking-pipe-1.yml | jq

# correct approach - step 2
cat data.json | jq -c "."  | pimo -c masking-pipe-2.yml | jq

# advanced usage - externalize
cat data.json | jq -c "."  | pimo -c masking-root.yml | jq

# advanced usage - using caches
cat data-cache.jsonl | jq -c "."  | pimo -c masking-cache.yml
