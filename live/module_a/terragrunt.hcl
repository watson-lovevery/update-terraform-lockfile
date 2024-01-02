terraform {
    source = "${get_repo_root()}/test"
}

include {
    path = find_in_parent_folders()
}
