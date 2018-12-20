## Kubernetes example airports application on GKE

This project is an example of implementing a REST MicroService architecture on Kubernetes which is hosted by GKE.

This project interacts with a CI/CD platform (Jenkins), in order to manage the building & deployment of applications.

### Run Locally

If you want to run this project locally, you will have to implement Minikube functionality. 

You can run the services locally with docker-compose. See docker-compose.yaml.

To build the images locally you will need to run the build.sh file in all the ./docker folders, and build.sh in ./services.

To convert the data you will need to install csvkit "brew install csvkit", and run ./data/convert_data.sh.

### Configuring GKE

You will need to create a key file for your GKE service account.

Go to IAM & Admin -> Service accounts -> SA -> Create Key

Place this file in the root directory of the project and call it "gcp_sa.key.json"

### Running on GKE

Install & configure gcp cli on local machine

You will need to manually build & release the jenkins container to your GKE registry

Change which userspaces you want in ./configs/users.json

Configure values in .env

Run ./setup.sh

After the cluster is setup you can access the cluster as each user with

./run.sh "$USERNAME"

When you load the dashboard you can use the kube.config file in the secrets/${USERNAME} folder to authenticate. If you are authenticating as a user that is not an admin you will not be able to see the default namespace, or even list the available namespaces. So you have to ensure your URL is pointing to your namespace.

After this you can access Jenkins to build & deploy the rest of the images

The Jenkins credentials are: admin/test