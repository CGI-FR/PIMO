
# define a secret key encoded in base64
export FF1_ENCRYPTION_KEY=$(echo -n "secret12secret12" | base64)

echo "Using masking.yml file to encrypt data.jsonl using FF1 algorithm gives:"

# apply ff1 masks to encrypt the siret column
cat data.jsonl | pimo

echo
echo "Using masking-decrypt.yml file to decrypt previous stream using FF1 algorithm gives:"

# apply ff1 masks to encrypt the siret column then apply the same mask with option decrypt: true
cat data.jsonl | pimo | pimo -c masking-decrypt.yml

echo
echo "Using masking-tweak.yml file to encrypt data.jsonl using FF1 algorithm with a tweak gives:"

cat data.jsonl | pimo -c masking-tweak.yml
