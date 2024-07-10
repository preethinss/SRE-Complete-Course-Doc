# steps
1. Create go application
2. Create Dockerfile with multistage
3. build image from Dockerfile
        - docker build -t <docker_hub_repository_name>/with-multistage:1.0.0 .
          * Note: execute docker build command where Dockerfile is present 
4. push image to docker hub
        - docker push <docker_hub_repository_name>/with-multistage:1.0.0
5. if any permission denied error while pushing image
        - docker login
            * provide username and password of your docker hub account

# Benefits
1. Reduce size of image
2. Provides security at container level