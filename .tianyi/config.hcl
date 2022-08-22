job "job1" "POST" {
    query = {
        token = "215d99e6f19335476c5eb212f59cea"
    }
    url = "https://gitlab.com/api/v4/projects/38044358/ref/master/trigger/pipeline"
}

pipeline "pipeline1" {
    job "job1" {}
}

pipeline "pipeline2" {
    job "job1" {}
}