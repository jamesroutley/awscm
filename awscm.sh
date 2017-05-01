function awscm() {
    local tmpfile=$(mktemp)
    echo $tmpfile
    awscmcore -f $tmpfile
    rm $tmpfile
}
