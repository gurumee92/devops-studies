resource "aws_s3_bucket" "s3" {
    bucket = "gurumee-learn-terraform"
    force_destroy = true
}