# URL Shortenner (Enhanced) in GO

* Ref: https://www.educative.io/courses/the-way-to-go
* IaC (terraform) taken from https://github.com/KevinSeghetti/aws-terraform-ecs 

## How to Use Locally
* Clone this repo
* go to the `app` directory
* Assuming you have Go installed, do `go mod init main` from the root directory of this repo
* Run the app using `go run .`
* Open your browser, and go to `localhost:4000/add`
* Type `https://go.dev/tour` in the form and hit enter
* Copy the string returned to the browser
* From your browser, go to `localhost:4000/<copied-string>`, you should be redirected to `https://go.dev/tour`


## Deploying to AWS ECS using Terraform
* Clone this repo
* Go to the `infrastructure` directory
* Assuming you have terraform installed and have aws credentials setup, run the following: 
    * `terraform init`
    * `terraform plan`
    * `terraform apply`
* After the above commands are successful, you should see the load-balancer address in your terminal
* Open your browser, and go to `<load-balancer-address>:4000/add`
* Type `https://go.dev/tour` in the form and hit enter
* Copy the string returned to the browser
* From your browser, go to `<load-balancer-address>:4000/<copied-string>`, you should be redirected to `https://go.dev/tour`
* After your experimentation is done, remove the AWS resources you created with `terraform destroy`


