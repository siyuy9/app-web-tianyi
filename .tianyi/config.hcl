job "job1" "POST" {
    url = https://gitlab.com/api/v4/projects/38044358/ref/REF_NAME/trigger/pipeline

    query = {
        token = 215d99e6f19335476c5eb212f59cea
    }
}

pipeline "pipeline1" {
    job "job1"
}

pipeline "pipeline2" {
    job "job1"
}