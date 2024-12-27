pipeline {
    agent {
        docker {
            image 'wuzor/jenkins-agent-go'
            args ' --privileged -v /var/run/docker.sock:/var/run/docker.sock'
        }        
    }
    environment {
        DOCKER_IMAGE = 'wuzor/simplebank'
        DOCKER_REGISTRY = 'registry.hub.docker.com' // Replace with actual registry URL
        DOCKER_CREDENTIALS = credentials('docker-cred-id')
        DB_CONTAINER_NAME = 'postgres'
        DB_USER = 'root'
        DB_PASSWORD = 'root'
        DB_NAME = 'simplebank'
        DB_PORT = '5432'
        DB_HOST = 'postgres'
        NETWORK_NAME = 'postgres'
    }
    
    stages {
        stage('Environment Setup') {
            agent {
                docker {
                    image 'wuzor/jenkins-agent-go'
                    args ' --privileged -v /var/run/docker.sock:/var/run/docker.sock'
                }
            }
            stages{
                stage('Checkout') {
                    steps {
                        deleteDir() // Deletes the entire workspace
                        checkout scm
                    }
                }
                stage('create network') {
                    steps{
                        sh 'make create-network'
                        deleteDir()
                    }
                }
            }
        }
        stage('Build Test and Deploy') {
            agent {
                docker {
                    image 'wuzor/jenkins-agent-go'
                    args ' --network=postgres --privileged -v /var/run/docker.sock:/var/run/docker.sock'
                }
            }
            stages {
                stage('Setup Database') {
                    steps {
                        script {
                            // Start the database container
                            sh '''
                                make start
                            '''

                            // Wait for the database to be ready
                            // sh '''
                            // until docker exec $DB_CONTAINER pg_isready -U $DB_USER; do
                            //     echo "Waiting for database to be ready..."
                            //     sleep 2
                            // done
                            // '''
                        }
                    }
                }
                stage('Checkout') {
                    steps {
                        deleteDir() // Deletes the entire workspace
                        checkout scm
                    }
                }
                stage('Linting') {
                    steps {
                        sh 'golangci-lint run --timeout 10m'
                    }
                }
                stage('Test') {
                    steps {
                        sh 'make test'
                    }
                }
                stage('Push') {
                    steps {
                        // sh '''
                            // echo $DOCKER_CREDENTIALS_PSW | docker login $DOCKER_REGISTRY -u $DOCKER_CREDENTIALS_USR --password-stdin
                            // docker build -t $DOCKER_IMAGE:$BUILD_NUMBER .
                            // docker push $DOCKER_IMAGE:$BUILD_NUMBER
                            // docker tag $DOCKER_IMAGE:$BUILD_NUMBER $DOCKER_IMAGE:latest
                            // docker push $DOCKER_IMAGE:latest

                        // '''
                        sh 'echo "yet to be impleted" '
                    }
                }
            }
        }
    }   
    
    post {
        always {
            script {
                docker.image('wuzor/jenkins-agent-go').inside('--network=postgres --privileged -v /var/run/docker.sock:/var/run/docker.sock') {
                    sh '''
                        echo "Logging out of Docker registry..."
                        docker logout $DOCKER_REGISTRY || true
                        echo "Cleaning up workspace..."
                        make stop
                    '''
                }
                cleanWs()
            }
        }
        
    }
}
