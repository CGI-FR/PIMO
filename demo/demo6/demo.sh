# this command apply the masks defined in masking.yml
# to data contained in the file data.jsonl
# and save the result in the file dataout.jsonl
pimo < data.jsonl > dataout.jsonl

# other alternative (same results)
cat data.jsonl | pimo > dataout.jsonl

# to also print the result on screen
cat data.jsonl | pimo | tee dataout.jsonl
