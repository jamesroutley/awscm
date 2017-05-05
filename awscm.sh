# Install awscm by adding the following
# function to your shell startup script:

awscm() {
    tmpfile=$(mktemp)
    echo "$tmpfile"
    awscmcore --file "$tmpfile" "$@"
    # Ignore shellcheck SC1090
    # shellcheck source=/dev/null
    . "$tmpfile"
    rm "$tmpfile"
}
