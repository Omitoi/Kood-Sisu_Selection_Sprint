<<COMMENT
Introduction

Learning Objectives:
* Understand the basics of setting up and managing a code repository.
* Create and execute a basic shell script.
* Synchronize local changes to an online repository using Git.

Instructions
1. Repository Setup:
Initialize a new repository named sprint on Gitea.
Clone the created repository to your local machine.
git clone https://gitea.koodsisu.fi/petrkubec/sprint.git
cd sprint

2. Script Creation:
In the cloned repository, create a file named hello.sh.
The script, when executed, should print: Hello petrkubec!.

$ bash hello.sh
Hello petrkubec!
$

3. Repository Synchronization:
Commit and push the changes to the online repository.
git add hello.sh
git commit -m "My initial script"
git push

4. Submission:
Once you've pushed your script, submit it for review using the provided submission mechanism.

Further resources: https://www.youtube.com/watch?v=CV-ven_rxhw
COMMENT

echo "Hello petrkubec!"